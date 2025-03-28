/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	bgppeerv1 "github.com/loxilb-io/kube-loxilb/pkg/bgp-client/applyconfiguration/bgppeer/v1"
	applyconfigurationbgppolicyapplyv1 "github.com/loxilb-io/kube-loxilb/pkg/bgp-client/applyconfiguration/bgppolicyapply/v1"
	applyconfigurationbgppolicydefinedsetsv1 "github.com/loxilb-io/kube-loxilb/pkg/bgp-client/applyconfiguration/bgppolicydefinedsets/v1"
	applyconfigurationbgppolicydefinitionv1 "github.com/loxilb-io/kube-loxilb/pkg/bgp-client/applyconfiguration/bgppolicydefinition/v1"
	v1 "github.com/loxilb-io/kube-loxilb/pkg/crds/bgppeer/v1"
	bgppolicyapplyv1 "github.com/loxilb-io/kube-loxilb/pkg/crds/bgppolicyapply/v1"
	bgppolicydefinedsetsv1 "github.com/loxilb-io/kube-loxilb/pkg/crds/bgppolicydefinedsets/v1"
	bgppolicydefinitionv1 "github.com/loxilb-io/kube-loxilb/pkg/crds/bgppolicydefinition/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=bgppeer.loxilb.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("BGPPeerModel"):
		return &bgppeerv1.BGPPeerModelApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BGPPeerService"):
		return &bgppeerv1.BGPPeerServiceApplyConfiguration{}

		// Group=bgppolicyapply.loxilb.io, Version=v1
	case bgppolicyapplyv1.SchemeGroupVersion.WithKind("BGPPolicyApplyModel"):
		return &applyconfigurationbgppolicyapplyv1.BGPPolicyApplyModelApplyConfiguration{}
	case bgppolicyapplyv1.SchemeGroupVersion.WithKind("BGPPolicyApplyService"):
		return &applyconfigurationbgppolicyapplyv1.BGPPolicyApplyServiceApplyConfiguration{}

		// Group=bgppolicydefinedsets.loxilb.io, Version=v1
	case bgppolicydefinedsetsv1.SchemeGroupVersion.WithKind("BGPPolicyDefinedSetsModel"):
		return &applyconfigurationbgppolicydefinedsetsv1.BGPPolicyDefinedSetsModelApplyConfiguration{}
	case bgppolicydefinedsetsv1.SchemeGroupVersion.WithKind("BGPPolicyDefinedSetsService"):
		return &applyconfigurationbgppolicydefinedsetsv1.BGPPolicyDefinedSetsServiceApplyConfiguration{}

		// Group=bgppolicydefinition.loxilb.io, Version=v1
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("Actions"):
		return &applyconfigurationbgppolicydefinitionv1.ActionsApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("BGPActions"):
		return &applyconfigurationbgppolicydefinitionv1.BGPActionsApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("BGPAsPathLength"):
		return &applyconfigurationbgppolicydefinitionv1.BGPAsPathLengthApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("BGPAsPathSet"):
		return &applyconfigurationbgppolicydefinitionv1.BGPAsPathSetApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("BGPCommunitySet"):
		return &applyconfigurationbgppolicydefinitionv1.BGPCommunitySetApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("BGPConditions"):
		return &applyconfigurationbgppolicydefinitionv1.BGPConditionsApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("BGPPolicyDefinition"):
		return &applyconfigurationbgppolicydefinitionv1.BGPPolicyDefinitionApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("BGPPolicyDefinitionService"):
		return &applyconfigurationbgppolicydefinitionv1.BGPPolicyDefinitionServiceApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("Conditions"):
		return &applyconfigurationbgppolicydefinitionv1.ConditionsApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("MatchNeighborSet"):
		return &applyconfigurationbgppolicydefinitionv1.MatchNeighborSetApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("MatchPrefixSet"):
		return &applyconfigurationbgppolicydefinitionv1.MatchPrefixSetApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("SetAsPathPrepend"):
		return &applyconfigurationbgppolicydefinitionv1.SetAsPathPrependApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("SetCommunity"):
		return &applyconfigurationbgppolicydefinitionv1.SetCommunityApplyConfiguration{}
	case bgppolicydefinitionv1.SchemeGroupVersion.WithKind("Statement"):
		return &applyconfigurationbgppolicydefinitionv1.StatementApplyConfiguration{}

	}
	return nil
}
