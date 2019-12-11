// +build !ignore_autogenerated

/*
Copyright 2019 Pivotal.

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

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqCluster) DeepCopyInto(out *RabbitmqCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqCluster.
func (in *RabbitmqCluster) DeepCopy() *RabbitmqCluster {
	if in == nil {
		return nil
	}
	out := new(RabbitmqCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterComputeResource) DeepCopyInto(out *RabbitmqClusterComputeResource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterComputeResource.
func (in *RabbitmqClusterComputeResource) DeepCopy() *RabbitmqClusterComputeResource {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterComputeResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterList) DeepCopyInto(out *RabbitmqClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RabbitmqCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterList.
func (in *RabbitmqClusterList) DeepCopy() *RabbitmqClusterList {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterPersistenceSpec) DeepCopyInto(out *RabbitmqClusterPersistenceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterPersistenceSpec.
func (in *RabbitmqClusterPersistenceSpec) DeepCopy() *RabbitmqClusterPersistenceSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterPersistenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterResourceSpec) DeepCopyInto(out *RabbitmqClusterResourceSpec) {
	*out = *in
	out.Request = in.Request
	out.Limit = in.Limit
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterResourceSpec.
func (in *RabbitmqClusterResourceSpec) DeepCopy() *RabbitmqClusterResourceSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterServiceSpec) DeepCopyInto(out *RabbitmqClusterServiceSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterServiceSpec.
func (in *RabbitmqClusterServiceSpec) DeepCopy() *RabbitmqClusterServiceSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterSpec) DeepCopyInto(out *RabbitmqClusterSpec) {
	*out = *in
	in.Service.DeepCopyInto(&out.Service)
	out.Persistence = in.Persistence
	out.Resource = in.Resource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterSpec.
func (in *RabbitmqClusterSpec) DeepCopy() *RabbitmqClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterStatus) DeepCopyInto(out *RabbitmqClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterStatus.
func (in *RabbitmqClusterStatus) DeepCopy() *RabbitmqClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterStatus)
	in.DeepCopyInto(out)
	return out
}
