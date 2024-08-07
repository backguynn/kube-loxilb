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

package v1

// BGPActionsApplyConfiguration represents an declarative configuration of the BGPActions type for use
// with apply.
type BGPActionsApplyConfiguration struct {
	SetMed            *string                             `json:"setMed,omitempty"`
	SetCommunity      *SetCommunityApplyConfiguration     `json:"setCommunity,omitempty"`
	SetExtCommunity   *SetCommunityApplyConfiguration     `json:"setExtCommunity,omitempty"`
	SetLargeCommunity *SetCommunityApplyConfiguration     `json:"setLargeCommunity,omitempty"`
	SetNextHop        *string                             `json:"setNextHop,omitempty"`
	SetLocalPerf      *int                                `json:"setLocalPerf,omitempty"`
	SetAsPathPrepend  *SetAsPathPrependApplyConfiguration `json:"setAsPathPrepend,omitempty"`
}

// BGPActionsApplyConfiguration constructs an declarative configuration of the BGPActions type for use with
// apply.
func BGPActions() *BGPActionsApplyConfiguration {
	return &BGPActionsApplyConfiguration{}
}

// WithSetMed sets the SetMed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SetMed field is set to the value of the last call.
func (b *BGPActionsApplyConfiguration) WithSetMed(value string) *BGPActionsApplyConfiguration {
	b.SetMed = &value
	return b
}

// WithSetCommunity sets the SetCommunity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SetCommunity field is set to the value of the last call.
func (b *BGPActionsApplyConfiguration) WithSetCommunity(value *SetCommunityApplyConfiguration) *BGPActionsApplyConfiguration {
	b.SetCommunity = value
	return b
}

// WithSetExtCommunity sets the SetExtCommunity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SetExtCommunity field is set to the value of the last call.
func (b *BGPActionsApplyConfiguration) WithSetExtCommunity(value *SetCommunityApplyConfiguration) *BGPActionsApplyConfiguration {
	b.SetExtCommunity = value
	return b
}

// WithSetLargeCommunity sets the SetLargeCommunity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SetLargeCommunity field is set to the value of the last call.
func (b *BGPActionsApplyConfiguration) WithSetLargeCommunity(value *SetCommunityApplyConfiguration) *BGPActionsApplyConfiguration {
	b.SetLargeCommunity = value
	return b
}

// WithSetNextHop sets the SetNextHop field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SetNextHop field is set to the value of the last call.
func (b *BGPActionsApplyConfiguration) WithSetNextHop(value string) *BGPActionsApplyConfiguration {
	b.SetNextHop = &value
	return b
}

// WithSetLocalPerf sets the SetLocalPerf field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SetLocalPerf field is set to the value of the last call.
func (b *BGPActionsApplyConfiguration) WithSetLocalPerf(value int) *BGPActionsApplyConfiguration {
	b.SetLocalPerf = &value
	return b
}

// WithSetAsPathPrepend sets the SetAsPathPrepend field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SetAsPathPrepend field is set to the value of the last call.
func (b *BGPActionsApplyConfiguration) WithSetAsPathPrepend(value *SetAsPathPrependApplyConfiguration) *BGPActionsApplyConfiguration {
	b.SetAsPathPrepend = value
	return b
}
