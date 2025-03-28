//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) NetLOX Inc

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPolicyApplyModel) DeepCopyInto(out *BGPPolicyApplyModel) {
	*out = *in
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPolicyApplyModel.
func (in *BGPPolicyApplyModel) DeepCopy() *BGPPolicyApplyModel {
	if in == nil {
		return nil
	}
	out := new(BGPPolicyApplyModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPolicyApplyService) DeepCopyInto(out *BGPPolicyApplyService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPolicyApplyService.
func (in *BGPPolicyApplyService) DeepCopy() *BGPPolicyApplyService {
	if in == nil {
		return nil
	}
	out := new(BGPPolicyApplyService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPPolicyApplyService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPolicyApplyServiceList) DeepCopyInto(out *BGPPolicyApplyServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*BGPPolicyApplyService, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BGPPolicyApplyService)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPolicyApplyServiceList.
func (in *BGPPolicyApplyServiceList) DeepCopy() *BGPPolicyApplyServiceList {
	if in == nil {
		return nil
	}
	out := new(BGPPolicyApplyServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPPolicyApplyServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPolicyApplyServiceStatus) DeepCopyInto(out *BGPPolicyApplyServiceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPolicyApplyServiceStatus.
func (in *BGPPolicyApplyServiceStatus) DeepCopy() *BGPPolicyApplyServiceStatus {
	if in == nil {
		return nil
	}
	out := new(BGPPolicyApplyServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPolicyApplySpec) DeepCopyInto(out *BGPPolicyApplySpec) {
	*out = *in
	in.Model.DeepCopyInto(&out.Model)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPolicyApplySpec.
func (in *BGPPolicyApplySpec) DeepCopy() *BGPPolicyApplySpec {
	if in == nil {
		return nil
	}
	out := new(BGPPolicyApplySpec)
	in.DeepCopyInto(out)
	return out
}
