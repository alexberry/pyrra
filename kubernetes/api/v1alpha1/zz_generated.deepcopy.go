//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 Pyrra Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatencyIndicator) DeepCopyInto(out *LatencyIndicator) {
	*out = *in
	out.Success = in.Success
	out.Total = in.Total
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatencyIndicator.
func (in *LatencyIndicator) DeepCopy() *LatencyIndicator {
	if in == nil {
		return nil
	}
	out := new(LatencyIndicator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Query) DeepCopyInto(out *Query) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Query.
func (in *Query) DeepCopy() *Query {
	if in == nil {
		return nil
	}
	out := new(Query)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RatioIndicator) DeepCopyInto(out *RatioIndicator) {
	*out = *in
	out.Errors = in.Errors
	out.Total = in.Total
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RatioIndicator.
func (in *RatioIndicator) DeepCopy() *RatioIndicator {
	if in == nil {
		return nil
	}
	out := new(RatioIndicator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelIndicator) DeepCopyInto(out *ServiceLevelIndicator) {
	*out = *in
	if in.Ratio != nil {
		in, out := &in.Ratio, &out.Ratio
		*out = new(RatioIndicator)
		**out = **in
	}
	if in.Latency != nil {
		in, out := &in.Latency, &out.Latency
		*out = new(LatencyIndicator)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelIndicator.
func (in *ServiceLevelIndicator) DeepCopy() *ServiceLevelIndicator {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelIndicator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelObjective) DeepCopyInto(out *ServiceLevelObjective) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelObjective.
func (in *ServiceLevelObjective) DeepCopy() *ServiceLevelObjective {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelObjective)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceLevelObjective) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelObjectiveList) DeepCopyInto(out *ServiceLevelObjectiveList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceLevelObjective, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelObjectiveList.
func (in *ServiceLevelObjectiveList) DeepCopy() *ServiceLevelObjectiveList {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelObjectiveList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceLevelObjectiveList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelObjectiveSpec) DeepCopyInto(out *ServiceLevelObjectiveSpec) {
	*out = *in
	in.ServiceLevelIndicator.DeepCopyInto(&out.ServiceLevelIndicator)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelObjectiveSpec.
func (in *ServiceLevelObjectiveSpec) DeepCopy() *ServiceLevelObjectiveSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelObjectiveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelObjectiveStatus) DeepCopyInto(out *ServiceLevelObjectiveStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelObjectiveStatus.
func (in *ServiceLevelObjectiveStatus) DeepCopy() *ServiceLevelObjectiveStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelObjectiveStatus)
	in.DeepCopyInto(out)
	return out
}
