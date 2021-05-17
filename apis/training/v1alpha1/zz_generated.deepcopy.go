// +build !ignore_autogenerated

/*
Copyright 2021 The Alibaba Authors.

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
	modelv1alpha1 "github.com/alibaba/kubedl/apis/model/v1alpha1"
	"github.com/alibaba/kubedl/pkg/job_controller/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticDLJob) DeepCopyInto(out *ElasticDLJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticDLJob.
func (in *ElasticDLJob) DeepCopy() *ElasticDLJob {
	if in == nil {
		return nil
	}
	out := new(ElasticDLJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticDLJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticDLJobList) DeepCopyInto(out *ElasticDLJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticDLJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticDLJobList.
func (in *ElasticDLJobList) DeepCopy() *ElasticDLJobList {
	if in == nil {
		return nil
	}
	out := new(ElasticDLJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticDLJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticDLJobSpec) DeepCopyInto(out *ElasticDLJobSpec) {
	*out = *in
	in.RunPolicy.DeepCopyInto(&out.RunPolicy)
	if in.ElasticDLReplicaSpecs != nil {
		in, out := &in.ElasticDLReplicaSpecs, &out.ElasticDLReplicaSpecs
		*out = make(map[v1.ReplicaType]*v1.ReplicaSpec, len(*in))
		for key, val := range *in {
			var outVal *v1.ReplicaSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1.ReplicaSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticDLJobSpec.
func (in *ElasticDLJobSpec) DeepCopy() *ElasticDLJobSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticDLJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LegacyV1Alpha1) DeepCopyInto(out *LegacyV1Alpha1) {
	*out = *in
	if in.DeprecatedGPUs != nil {
		in, out := &in.DeprecatedGPUs, &out.DeprecatedGPUs
		*out = new(int32)
		**out = **in
	}
	if in.GPUsPerNode != nil {
		in, out := &in.GPUsPerNode, &out.GPUsPerNode
		*out = new(int32)
		**out = **in
	}
	if in.ProcessingUnits != nil {
		in, out := &in.ProcessingUnits, &out.ProcessingUnits
		*out = new(int32)
		**out = **in
	}
	if in.ProcessingUnitsPerNode != nil {
		in, out := &in.ProcessingUnitsPerNode, &out.ProcessingUnitsPerNode
		*out = new(int32)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LegacyV1Alpha1.
func (in *LegacyV1Alpha1) DeepCopy() *LegacyV1Alpha1 {
	if in == nil {
		return nil
	}
	out := new(LegacyV1Alpha1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LegacyV1Alpha2) DeepCopyInto(out *LegacyV1Alpha2) {
	*out = *in
	if in.MPIDistribution != nil {
		in, out := &in.MPIDistribution, &out.MPIDistribution
		*out = new(MPIDistributionType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LegacyV1Alpha2.
func (in *LegacyV1Alpha2) DeepCopy() *LegacyV1Alpha2 {
	if in == nil {
		return nil
	}
	out := new(LegacyV1Alpha2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MPIJob) DeepCopyInto(out *MPIJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MPIJob.
func (in *MPIJob) DeepCopy() *MPIJob {
	if in == nil {
		return nil
	}
	out := new(MPIJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MPIJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MPIJobLegacySpec) DeepCopyInto(out *MPIJobLegacySpec) {
	*out = *in
	if in.RunPolicy != nil {
		in, out := &in.RunPolicy, &out.RunPolicy
		*out = new(v1.RunPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.LegacyV1Alpha1 != nil {
		in, out := &in.LegacyV1Alpha1, &out.LegacyV1Alpha1
		*out = new(LegacyV1Alpha1)
		(*in).DeepCopyInto(*out)
	}
	if in.LegacyV1Alpha2 != nil {
		in, out := &in.LegacyV1Alpha2, &out.LegacyV1Alpha2
		*out = new(LegacyV1Alpha2)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MPIJobLegacySpec.
func (in *MPIJobLegacySpec) DeepCopy() *MPIJobLegacySpec {
	if in == nil {
		return nil
	}
	out := new(MPIJobLegacySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MPIJobList) DeepCopyInto(out *MPIJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MPIJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MPIJobList.
func (in *MPIJobList) DeepCopy() *MPIJobList {
	if in == nil {
		return nil
	}
	out := new(MPIJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MPIJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MPIJobSpec) DeepCopyInto(out *MPIJobSpec) {
	*out = *in
	if in.SlotsPerWorker != nil {
		in, out := &in.SlotsPerWorker, &out.SlotsPerWorker
		*out = new(int32)
		**out = **in
	}
	if in.MPIReplicaSpecs != nil {
		in, out := &in.MPIReplicaSpecs, &out.MPIReplicaSpecs
		*out = make(map[v1.ReplicaType]*v1.ReplicaSpec, len(*in))
		for key, val := range *in {
			var outVal *v1.ReplicaSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1.ReplicaSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	in.RunPolicy.DeepCopyInto(&out.RunPolicy)
	if in.MPIJobLegacySpec != nil {
		in, out := &in.MPIJobLegacySpec, &out.MPIJobLegacySpec
		*out = new(MPIJobLegacySpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MPIJobSpec.
func (in *MPIJobSpec) DeepCopy() *MPIJobSpec {
	if in == nil {
		return nil
	}
	out := new(MPIJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MarsJob) DeepCopyInto(out *MarsJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MarsJob.
func (in *MarsJob) DeepCopy() *MarsJob {
	if in == nil {
		return nil
	}
	out := new(MarsJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MarsJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MarsJobList) DeepCopyInto(out *MarsJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MarsJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MarsJobList.
func (in *MarsJobList) DeepCopy() *MarsJobList {
	if in == nil {
		return nil
	}
	out := new(MarsJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MarsJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MarsJobSpec) DeepCopyInto(out *MarsJobSpec) {
	*out = *in
	in.RunPolicy.DeepCopyInto(&out.RunPolicy)
	if in.WorkerMemoryTuningPolicy != nil {
		in, out := &in.WorkerMemoryTuningPolicy, &out.WorkerMemoryTuningPolicy
		*out = new(MarsWorkerMemoryTuningPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.WebHost != nil {
		in, out := &in.WebHost, &out.WebHost
		*out = new(string)
		**out = **in
	}
	if in.MarsReplicaSpecs != nil {
		in, out := &in.MarsReplicaSpecs, &out.MarsReplicaSpecs
		*out = make(map[v1.ReplicaType]*v1.ReplicaSpec, len(*in))
		for key, val := range *in {
			var outVal *v1.ReplicaSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1.ReplicaSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MarsJobSpec.
func (in *MarsJobSpec) DeepCopy() *MarsJobSpec {
	if in == nil {
		return nil
	}
	out := new(MarsJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MarsJobStatus) DeepCopyInto(out *MarsJobStatus) {
	*out = *in
	in.JobStatus.DeepCopyInto(&out.JobStatus)
	if in.WebServiceAddresses != nil {
		in, out := &in.WebServiceAddresses, &out.WebServiceAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MarsJobStatus.
func (in *MarsJobStatus) DeepCopy() *MarsJobStatus {
	if in == nil {
		return nil
	}
	out := new(MarsJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MarsWorkerMemoryTuningPolicy) DeepCopyInto(out *MarsWorkerMemoryTuningPolicy) {
	*out = *in
	if in.PlasmaStore != nil {
		in, out := &in.PlasmaStore, &out.PlasmaStore
		*out = new(string)
		**out = **in
	}
	if in.LockFreeFileIO != nil {
		in, out := &in.LockFreeFileIO, &out.LockFreeFileIO
		*out = new(bool)
		**out = **in
	}
	if in.SpillDirs != nil {
		in, out := &in.SpillDirs, &out.SpillDirs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WorkerCachePercentage != nil {
		in, out := &in.WorkerCachePercentage, &out.WorkerCachePercentage
		*out = new(int32)
		**out = **in
	}
	if in.WorkerCacheSize != nil {
		in, out := &in.WorkerCacheSize, &out.WorkerCacheSize
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MarsWorkerMemoryTuningPolicy.
func (in *MarsWorkerMemoryTuningPolicy) DeepCopy() *MarsWorkerMemoryTuningPolicy {
	if in == nil {
		return nil
	}
	out := new(MarsWorkerMemoryTuningPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PyTorchJob) DeepCopyInto(out *PyTorchJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PyTorchJob.
func (in *PyTorchJob) DeepCopy() *PyTorchJob {
	if in == nil {
		return nil
	}
	out := new(PyTorchJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PyTorchJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PyTorchJobList) DeepCopyInto(out *PyTorchJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PyTorchJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PyTorchJobList.
func (in *PyTorchJobList) DeepCopy() *PyTorchJobList {
	if in == nil {
		return nil
	}
	out := new(PyTorchJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PyTorchJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PyTorchJobSpec) DeepCopyInto(out *PyTorchJobSpec) {
	*out = *in
	in.RunPolicy.DeepCopyInto(&out.RunPolicy)
	if in.PyTorchReplicaSpecs != nil {
		in, out := &in.PyTorchReplicaSpecs, &out.PyTorchReplicaSpecs
		*out = make(map[v1.ReplicaType]*v1.ReplicaSpec, len(*in))
		for key, val := range *in {
			var outVal *v1.ReplicaSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1.ReplicaSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.ModelVersion != nil {
		in, out := &in.ModelVersion, &out.ModelVersion
		*out = new(modelv1alpha1.ModelVersionSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PyTorchJobSpec.
func (in *PyTorchJobSpec) DeepCopy() *PyTorchJobSpec {
	if in == nil {
		return nil
	}
	out := new(PyTorchJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PyTorchJobStatus) DeepCopyInto(out *PyTorchJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PyTorchJobStatus.
func (in *PyTorchJobStatus) DeepCopy() *PyTorchJobStatus {
	if in == nil {
		return nil
	}
	out := new(PyTorchJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFJob) DeepCopyInto(out *TFJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFJob.
func (in *TFJob) DeepCopy() *TFJob {
	if in == nil {
		return nil
	}
	out := new(TFJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TFJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFJobList) DeepCopyInto(out *TFJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TFJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFJobList.
func (in *TFJobList) DeepCopy() *TFJobList {
	if in == nil {
		return nil
	}
	out := new(TFJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TFJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFJobSpec) DeepCopyInto(out *TFJobSpec) {
	*out = *in
	in.RunPolicy.DeepCopyInto(&out.RunPolicy)
	if in.SuccessPolicy != nil {
		in, out := &in.SuccessPolicy, &out.SuccessPolicy
		*out = new(v1.SuccessPolicy)
		**out = **in
	}
	if in.TFReplicaSpecs != nil {
		in, out := &in.TFReplicaSpecs, &out.TFReplicaSpecs
		*out = make(map[v1.ReplicaType]*v1.ReplicaSpec, len(*in))
		for key, val := range *in {
			var outVal *v1.ReplicaSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1.ReplicaSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.ModelVersion != nil {
		in, out := &in.ModelVersion, &out.ModelVersion
		*out = new(modelv1alpha1.ModelVersionSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFJobSpec.
func (in *TFJobSpec) DeepCopy() *TFJobSpec {
	if in == nil {
		return nil
	}
	out := new(TFJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XDLJob) DeepCopyInto(out *XDLJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XDLJob.
func (in *XDLJob) DeepCopy() *XDLJob {
	if in == nil {
		return nil
	}
	out := new(XDLJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XDLJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XDLJobList) DeepCopyInto(out *XDLJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XDLJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XDLJobList.
func (in *XDLJobList) DeepCopy() *XDLJobList {
	if in == nil {
		return nil
	}
	out := new(XDLJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XDLJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XDLJobSpec) DeepCopyInto(out *XDLJobSpec) {
	*out = *in
	in.RunPolicy.DeepCopyInto(&out.RunPolicy)
	if in.XDLReplicaSpecs != nil {
		in, out := &in.XDLReplicaSpecs, &out.XDLReplicaSpecs
		*out = make(map[v1.ReplicaType]*v1.ReplicaSpec, len(*in))
		for key, val := range *in {
			var outVal *v1.ReplicaSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1.ReplicaSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.MinFinishWorkerNum != nil {
		in, out := &in.MinFinishWorkerNum, &out.MinFinishWorkerNum
		*out = new(int32)
		**out = **in
	}
	if in.MinFinishWorkerPercentage != nil {
		in, out := &in.MinFinishWorkerPercentage, &out.MinFinishWorkerPercentage
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XDLJobSpec.
func (in *XDLJobSpec) DeepCopy() *XDLJobSpec {
	if in == nil {
		return nil
	}
	out := new(XDLJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XDLJobStatus) DeepCopyInto(out *XDLJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XDLJobStatus.
func (in *XDLJobStatus) DeepCopy() *XDLJobStatus {
	if in == nil {
		return nil
	}
	out := new(XDLJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XGBoostJob) DeepCopyInto(out *XGBoostJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XGBoostJob.
func (in *XGBoostJob) DeepCopy() *XGBoostJob {
	if in == nil {
		return nil
	}
	out := new(XGBoostJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XGBoostJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XGBoostJobList) DeepCopyInto(out *XGBoostJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XGBoostJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XGBoostJobList.
func (in *XGBoostJobList) DeepCopy() *XGBoostJobList {
	if in == nil {
		return nil
	}
	out := new(XGBoostJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XGBoostJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XGBoostJobSpec) DeepCopyInto(out *XGBoostJobSpec) {
	*out = *in
	in.RunPolicy.DeepCopyInto(&out.RunPolicy)
	if in.XGBReplicaSpecs != nil {
		in, out := &in.XGBReplicaSpecs, &out.XGBReplicaSpecs
		*out = make(map[v1.ReplicaType]*v1.ReplicaSpec, len(*in))
		for key, val := range *in {
			var outVal *v1.ReplicaSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1.ReplicaSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XGBoostJobSpec.
func (in *XGBoostJobSpec) DeepCopy() *XGBoostJobSpec {
	if in == nil {
		return nil
	}
	out := new(XGBoostJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XGBoostJobStatus) DeepCopyInto(out *XGBoostJobStatus) {
	*out = *in
	in.JobStatus.DeepCopyInto(&out.JobStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XGBoostJobStatus.
func (in *XGBoostJobStatus) DeepCopy() *XGBoostJobStatus {
	if in == nil {
		return nil
	}
	out := new(XGBoostJobStatus)
	in.DeepCopyInto(out)
	return out
}
