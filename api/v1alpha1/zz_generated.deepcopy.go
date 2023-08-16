//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalStorage) DeepCopyInto(out *AdditionalStorage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalStorage.
func (in *AdditionalStorage) DeepCopy() *AdditionalStorage {
	if in == nil {
		return nil
	}
	out := new(AdditionalStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackEnd) DeepCopyInto(out *BackEnd) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackEnd.
func (in *BackEnd) DeepCopy() *BackEnd {
	if in == nil {
		return nil
	}
	out := new(BackEnd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.Backends = in.Backends
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]Domain, len(*in))
		copy(*out, *in)
	}
	out.HealthCheck = in.HealthCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	out.VPC = in.VPC
	in.APIServerLoadbalancers.DeepCopyInto(&out.APIServerLoadbalancers)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]BackEnd, len(*in))
		copy(*out, *in)
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]Domain, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCSpec) DeepCopyInto(out *VPCSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCSpec.
func (in *VPCSpec) DeepCopy() *VPCSpec {
	if in == nil {
		return nil
	}
	out := new(VPCSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpsieCluster) DeepCopyInto(out *VpsieCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpsieCluster.
func (in *VpsieCluster) DeepCopy() *VpsieCluster {
	if in == nil {
		return nil
	}
	out := new(VpsieCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpsieCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpsieClusterList) DeepCopyInto(out *VpsieClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpsieCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpsieClusterList.
func (in *VpsieClusterList) DeepCopy() *VpsieClusterList {
	if in == nil {
		return nil
	}
	out := new(VpsieClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpsieClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpsieClusterSpec) DeepCopyInto(out *VpsieClusterSpec) {
	*out = *in
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	in.Network.DeepCopyInto(&out.Network)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpsieClusterSpec.
func (in *VpsieClusterSpec) DeepCopy() *VpsieClusterSpec {
	if in == nil {
		return nil
	}
	out := new(VpsieClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpsieClusterStatus) DeepCopyInto(out *VpsieClusterStatus) {
	*out = *in
	out.Network = in.Network
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpsieClusterStatus.
func (in *VpsieClusterStatus) DeepCopy() *VpsieClusterStatus {
	if in == nil {
		return nil
	}
	out := new(VpsieClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpsieLoadBalancerHealthCheck) DeepCopyInto(out *VpsieLoadBalancerHealthCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpsieLoadBalancerHealthCheck.
func (in *VpsieLoadBalancerHealthCheck) DeepCopy() *VpsieLoadBalancerHealthCheck {
	if in == nil {
		return nil
	}
	out := new(VpsieLoadBalancerHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpsieMachine) DeepCopyInto(out *VpsieMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpsieMachine.
func (in *VpsieMachine) DeepCopy() *VpsieMachine {
	if in == nil {
		return nil
	}
	out := new(VpsieMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpsieMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpsieMachineList) DeepCopyInto(out *VpsieMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpsieMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpsieMachineList.
func (in *VpsieMachineList) DeepCopy() *VpsieMachineList {
	if in == nil {
		return nil
	}
	out := new(VpsieMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpsieMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpsieMachineSpec) DeepCopyInto(out *VpsieMachineSpec) {
	*out = *in
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.VpsiePlan != nil {
		in, out := &in.VpsiePlan, &out.VpsiePlan
		*out = new(string)
		**out = **in
	}
	if in.OsIdentifier != nil {
		in, out := &in.OsIdentifier, &out.OsIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]AdditionalStorage, len(*in))
		copy(*out, *in)
	}
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpsieMachineSpec.
func (in *VpsieMachineSpec) DeepCopy() *VpsieMachineSpec {
	if in == nil {
		return nil
	}
	out := new(VpsieMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpsieMachineStatus) DeepCopyInto(out *VpsieMachineStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.InstanceStatus != nil {
		in, out := &in.InstanceStatus, &out.InstanceStatus
		*out = new(VpsieInstanceStatus)
		**out = **in
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpsieMachineStatus.
func (in *VpsieMachineStatus) DeepCopy() *VpsieMachineStatus {
	if in == nil {
		return nil
	}
	out := new(VpsieMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpsieNetworkResource) DeepCopyInto(out *VpsieNetworkResource) {
	*out = *in
	out.APIServerLoadbalancersRef = in.APIServerLoadbalancersRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpsieNetworkResource.
func (in *VpsieNetworkResource) DeepCopy() *VpsieNetworkResource {
	if in == nil {
		return nil
	}
	out := new(VpsieNetworkResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpsieResourceReference) DeepCopyInto(out *VpsieResourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpsieResourceReference.
func (in *VpsieResourceReference) DeepCopy() *VpsieResourceReference {
	if in == nil {
		return nil
	}
	out := new(VpsieResourceReference)
	in.DeepCopyInto(out)
	return out
}
