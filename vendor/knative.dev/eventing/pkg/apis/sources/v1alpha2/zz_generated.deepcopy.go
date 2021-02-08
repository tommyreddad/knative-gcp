// +build !ignore_autogenerated

/*
Copyright 2021 The Knative Authors

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

package v1alpha2

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIVersionKind) DeepCopyInto(out *APIVersionKind) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIVersionKind.
func (in *APIVersionKind) DeepCopy() *APIVersionKind {
	if in == nil {
		return nil
	}
	out := new(APIVersionKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIVersionKindSelector) DeepCopyInto(out *APIVersionKindSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIVersionKindSelector.
func (in *APIVersionKindSelector) DeepCopy() *APIVersionKindSelector {
	if in == nil {
		return nil
	}
	out := new(APIVersionKindSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiServerSource) DeepCopyInto(out *ApiServerSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiServerSource.
func (in *ApiServerSource) DeepCopy() *ApiServerSource {
	if in == nil {
		return nil
	}
	out := new(ApiServerSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiServerSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiServerSourceList) DeepCopyInto(out *ApiServerSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApiServerSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiServerSourceList.
func (in *ApiServerSourceList) DeepCopy() *ApiServerSourceList {
	if in == nil {
		return nil
	}
	out := new(ApiServerSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiServerSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiServerSourceSpec) DeepCopyInto(out *ApiServerSourceSpec) {
	*out = *in
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]APIVersionKindSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceOwner != nil {
		in, out := &in.ResourceOwner, &out.ResourceOwner
		*out = new(APIVersionKind)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiServerSourceSpec.
func (in *ApiServerSourceSpec) DeepCopy() *ApiServerSourceSpec {
	if in == nil {
		return nil
	}
	out := new(ApiServerSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiServerSourceStatus) DeepCopyInto(out *ApiServerSourceStatus) {
	*out = *in
	in.SourceStatus.DeepCopyInto(&out.SourceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiServerSourceStatus.
func (in *ApiServerSourceStatus) DeepCopy() *ApiServerSourceStatus {
	if in == nil {
		return nil
	}
	out := new(ApiServerSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSource) DeepCopyInto(out *ContainerSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSource.
func (in *ContainerSource) DeepCopy() *ContainerSource {
	if in == nil {
		return nil
	}
	out := new(ContainerSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSourceList) DeepCopyInto(out *ContainerSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContainerSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSourceList.
func (in *ContainerSourceList) DeepCopy() *ContainerSourceList {
	if in == nil {
		return nil
	}
	out := new(ContainerSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSourceSpec) DeepCopyInto(out *ContainerSourceSpec) {
	*out = *in
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSourceSpec.
func (in *ContainerSourceSpec) DeepCopy() *ContainerSourceSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSourceStatus) DeepCopyInto(out *ContainerSourceStatus) {
	*out = *in
	in.SourceStatus.DeepCopyInto(&out.SourceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSourceStatus.
func (in *ContainerSourceStatus) DeepCopy() *ContainerSourceStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PingSource) DeepCopyInto(out *PingSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PingSource.
func (in *PingSource) DeepCopy() *PingSource {
	if in == nil {
		return nil
	}
	out := new(PingSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PingSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PingSourceList) DeepCopyInto(out *PingSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PingSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PingSourceList.
func (in *PingSourceList) DeepCopy() *PingSourceList {
	if in == nil {
		return nil
	}
	out := new(PingSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PingSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PingSourceSpec) DeepCopyInto(out *PingSourceSpec) {
	*out = *in
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PingSourceSpec.
func (in *PingSourceSpec) DeepCopy() *PingSourceSpec {
	if in == nil {
		return nil
	}
	out := new(PingSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PingSourceStatus) DeepCopyInto(out *PingSourceStatus) {
	*out = *in
	in.SourceStatus.DeepCopyInto(&out.SourceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PingSourceStatus.
func (in *PingSourceStatus) DeepCopy() *PingSourceStatus {
	if in == nil {
		return nil
	}
	out := new(PingSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkBinding) DeepCopyInto(out *SinkBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkBinding.
func (in *SinkBinding) DeepCopy() *SinkBinding {
	if in == nil {
		return nil
	}
	out := new(SinkBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SinkBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkBindingList) DeepCopyInto(out *SinkBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SinkBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkBindingList.
func (in *SinkBindingList) DeepCopy() *SinkBindingList {
	if in == nil {
		return nil
	}
	out := new(SinkBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SinkBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkBindingSpec) DeepCopyInto(out *SinkBindingSpec) {
	*out = *in
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	in.BindingSpec.DeepCopyInto(&out.BindingSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkBindingSpec.
func (in *SinkBindingSpec) DeepCopy() *SinkBindingSpec {
	if in == nil {
		return nil
	}
	out := new(SinkBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkBindingStatus) DeepCopyInto(out *SinkBindingStatus) {
	*out = *in
	in.SourceStatus.DeepCopyInto(&out.SourceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkBindingStatus.
func (in *SinkBindingStatus) DeepCopy() *SinkBindingStatus {
	if in == nil {
		return nil
	}
	out := new(SinkBindingStatus)
	in.DeepCopyInto(out)
	return out
}
