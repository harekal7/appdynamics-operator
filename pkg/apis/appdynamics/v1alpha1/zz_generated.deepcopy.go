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
func (in *AgentRequest) DeepCopyInto(out *AgentRequest) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchString != nil {
		in, out := &in.MatchString, &out.MatchString
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentRequest.
func (in *AgentRequest) DeepCopy() *AgentRequest {
	if in == nil {
		return nil
	}
	out := new(AgentRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDBag) DeepCopyInto(out *AppDBag) {
	*out = *in
	if in.NsToMonitor != nil {
		in, out := &in.NsToMonitor, &out.NsToMonitor
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NsToMonitorExclude != nil {
		in, out := &in.NsToMonitorExclude, &out.NsToMonitorExclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeploysToDashboard != nil {
		in, out := &in.DeploysToDashboard, &out.DeploysToDashboard
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodesToMonitor != nil {
		in, out := &in.NodesToMonitor, &out.NodesToMonitor
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodesToMonitorExclude != nil {
		in, out := &in.NodesToMonitorExclude, &out.NodesToMonitorExclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NsToInstrument != nil {
		in, out := &in.NsToInstrument, &out.NsToInstrument
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NsToInstrumentExclude != nil {
		in, out := &in.NsToInstrumentExclude, &out.NsToInstrumentExclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NSInstrumentRule != nil {
		in, out := &in.NSInstrumentRule, &out.NSInstrumentRule
		*out = make([]AgentRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InstrumentMatchString != nil {
		in, out := &in.InstrumentMatchString, &out.InstrumentMatchString
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SchemaUpdateCache != nil {
		in, out := &in.SchemaUpdateCache, &out.SchemaUpdateCache
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDBag.
func (in *AppDBag) DeepCopy() *AppDBag {
	if in == nil {
		return nil
	}
	out := new(AppDBag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAgent) DeepCopyInto(out *ClusterAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAgent.
func (in *ClusterAgent) DeepCopy() *ClusterAgent {
	if in == nil {
		return nil
	}
	out := new(ClusterAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAgentList) DeepCopyInto(out *ClusterAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAgentList.
func (in *ClusterAgentList) DeepCopy() *ClusterAgentList {
	if in == nil {
		return nil
	}
	out := new(ClusterAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAgentSpec) DeepCopyInto(out *ClusterAgentSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.InstrumentMatchString != nil {
		in, out := &in.InstrumentMatchString, &out.InstrumentMatchString
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NsToInstrument != nil {
		in, out := &in.NsToInstrument, &out.NsToInstrument
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NsToInstrumentExclude != nil {
		in, out := &in.NsToInstrumentExclude, &out.NsToInstrumentExclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NsToMonitor != nil {
		in, out := &in.NsToMonitor, &out.NsToMonitor
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NsToMonitorExclude != nil {
		in, out := &in.NsToMonitorExclude, &out.NsToMonitorExclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodesToMonitor != nil {
		in, out := &in.NodesToMonitor, &out.NodesToMonitor
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodesToMonitorExclude != nil {
		in, out := &in.NodesToMonitorExclude, &out.NodesToMonitorExclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InstrumentRule != nil {
		in, out := &in.InstrumentRule, &out.InstrumentRule
		*out = make([]AgentRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DashboardTiers != nil {
		in, out := &in.DashboardTiers, &out.DashboardTiers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAgentSpec.
func (in *ClusterAgentSpec) DeepCopy() *ClusterAgentSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAgentStatus) DeepCopyInto(out *ClusterAgentStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAgentStatus.
func (in *ClusterAgentStatus) DeepCopy() *ClusterAgentStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterAgentStatus)
	in.DeepCopyInto(out)
	return out
}