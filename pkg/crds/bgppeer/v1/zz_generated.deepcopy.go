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
func (in *BGPNeigh) DeepCopyInto(out *BGPNeigh) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPNeigh.
func (in *BGPNeigh) DeepCopy() *BGPNeigh {
	if in == nil {
		return nil
	}
	out := new(BGPNeigh)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeerModel) DeepCopyInto(out *BGPPeerModel) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeerModel.
func (in *BGPPeerModel) DeepCopy() *BGPPeerModel {
	if in == nil {
		return nil
	}
	out := new(BGPPeerModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeerService) DeepCopyInto(out *BGPPeerService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeerService.
func (in *BGPPeerService) DeepCopy() *BGPPeerService {
	if in == nil {
		return nil
	}
	out := new(BGPPeerService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPPeerService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeerServiceList) DeepCopyInto(out *BGPPeerServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*BGPPeerService, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BGPPeerService)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeerServiceList.
func (in *BGPPeerServiceList) DeepCopy() *BGPPeerServiceList {
	if in == nil {
		return nil
	}
	out := new(BGPPeerServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPPeerServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeerServiceSpec) DeepCopyInto(out *BGPPeerServiceSpec) {
	*out = *in
	out.Model = in.Model
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeerServiceSpec.
func (in *BGPPeerServiceSpec) DeepCopy() *BGPPeerServiceSpec {
	if in == nil {
		return nil
	}
	out := new(BGPPeerServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeerServiceStatus) DeepCopyInto(out *BGPPeerServiceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeerServiceStatus.
func (in *BGPPeerServiceStatus) DeepCopy() *BGPPeerServiceStatus {
	if in == nil {
		return nil
	}
	out := new(BGPPeerServiceStatus)
	in.DeepCopyInto(out)
	return out
}
