//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCenter) DeepCopyInto(out *DataCenter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCenter.
func (in *DataCenter) DeepCopy() *DataCenter {
	if in == nil {
		return nil
	}
	out := new(DataCenter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataCenter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCenterList) DeepCopyInto(out *DataCenterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataCenter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCenterList.
func (in *DataCenterList) DeepCopy() *DataCenterList {
	if in == nil {
		return nil
	}
	out := new(DataCenterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataCenterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCenterSpec) DeepCopyInto(out *DataCenterSpec) {
	*out = *in
	in.ResourceInfo.DeepCopyInto(&out.ResourceInfo)
	out.Scheduler = in.Scheduler
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCenterSpec.
func (in *DataCenterSpec) DeepCopy() *DataCenterSpec {
	if in == nil {
		return nil
	}
	out := new(DataCenterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCenterStatus) DeepCopyInto(out *DataCenterStatus) {
	*out = *in
	out.Idle = in.Idle
	in.UpdatedTime.DeepCopyInto(&out.UpdatedTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCenterStatus.
func (in *DataCenterStatus) DeepCopy() *DataCenterStatus {
	if in == nil {
		return nil
	}
	out := new(DataCenterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Idle) DeepCopyInto(out *Idle) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Idle.
func (in *Idle) DeepCopy() *Idle {
	if in == nil {
		return nil
	}
	out := new(Idle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInfo) DeepCopyInto(out *ResourceInfo) {
	*out = *in
	if in.ReserveResources != nil {
		in, out := &in.ReserveResources, &out.ReserveResources
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInfo.
func (in *ResourceInfo) DeepCopy() *ResourceInfo {
	if in == nil {
		return nil
	}
	out := new(ResourceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerParams) DeepCopyInto(out *SchedulerParams) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerParams.
func (in *SchedulerParams) DeepCopy() *SchedulerParams {
	if in == nil {
		return nil
	}
	out := new(SchedulerParams)
	in.DeepCopyInto(out)
	return out
}
