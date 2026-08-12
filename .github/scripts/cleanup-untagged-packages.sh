#!/usr/bin/env bash
#
# Delete untagged GHCR container versions of a package.
#
# Usage: cleanup-untagged-packages.sh <owner> <package-name>
#
# GHCR never garbage-collects manifests that lost their tag, so a workflow that
# repeatedly pushes the same tag (":latest", ":preflight", ...) leaves one dead
# version behind on every run. This script removes those.
#
# Multi-arch images need care: "docker buildx" pushes an OCI index plus one
# child manifest per platform (and one attestation manifest per platform), and
# GHCR lists every child as its own *untagged* version. Deleting those breaks
# the tag that points at them, so children referenced by a tagged index are
# preserved here. This is why plain "delete-only-untagged-versions" tooling is
# not safe for the multi-arch kube-loxilb package.
#
# Requires: gh (authenticated via GH_TOKEN), jq, curl.
#
set -euo pipefail

OWNER="${1:?usage: $0 <owner> <package-name>}"
PACKAGE="${2:?usage: $0 <owner> <package-name>}"

# Versions younger than this are left alone so a concurrent build cannot have
# its freshly pushed children deleted out from under it.
MIN_AGE_SECONDS="${MIN_AGE_SECONDS:-1800}"

# Set DRY_RUN=1 to list what would be removed without deleting anything.
DRY_RUN="${DRY_RUN:-0}"

VERSIONS_API="/orgs/${OWNER}/packages/container/${PACKAGE}/versions"

versions="$(gh api --paginate "${VERSIONS_API}?per_page=100")"

total="$(jq 'length' <<<"$versions")"
echo "${PACKAGE}: ${total} version(s) found"

# Collect the digests that tagged manifest lists point at, so multi-arch tags
# keep working after the sweep.
registry_token="$(curl -fsSL -u "${GITHUB_ACTOR:-github-actions}:${GH_TOKEN}" \
  "https://ghcr.io/token?scope=repository:${OWNER}/${PACKAGE}:pull&service=ghcr.io" | jq -r '.token')"

accept='application/vnd.oci.image.index.v1+json'
accept+=',application/vnd.docker.distribution.manifest.list.v2+json'
accept+=',application/vnd.oci.image.manifest.v1+json'
accept+=',application/vnd.docker.distribution.manifest.v2+json'

referenced="$(mktemp)"
trap 'rm -f "$referenced"' EXIT

while read -r tag; do
  [ -n "$tag" ] || continue
  curl -fsSL -H "Authorization: Bearer ${registry_token}" -H "Accept: ${accept}" \
    "https://ghcr.io/v2/${OWNER}/${PACKAGE}/manifests/${tag}" \
    | jq -r '.manifests[]?.digest // empty' >>"$referenced"
done < <(jq -r '.[].metadata.container.tags[]?' <<<"$versions")

sort -u -o "$referenced" "$referenced"
echo "${PACKAGE}: $(wc -l <"$referenced") child digest(s) referenced by tagged indexes"

now="$(date +%s)"
deleted=0

while read -r id digest created; do
  if grep -qxF "$digest" "$referenced"; then
    continue
  fi
  if [ $(( now - $(date -d "$created" +%s) )) -lt "$MIN_AGE_SECONDS" ]; then
    echo "skip   ${digest} (younger than ${MIN_AGE_SECONDS}s)"
    continue
  fi
  if [ "$DRY_RUN" = "1" ]; then
    echo "would delete ${digest}"
  else
    echo "delete ${digest}"
    gh api -X DELETE "${VERSIONS_API}/${id}"
  fi
  deleted=$(( deleted + 1 ))
done < <(jq -r '.[] | select((.metadata.container.tags | length) == 0)
                    | "\(.id) \(.name) \(.created_at)"' <<<"$versions")

echo "${PACKAGE}: deleted ${deleted} untagged version(s)"
