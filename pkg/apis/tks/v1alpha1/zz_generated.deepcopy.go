// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TksComplianceMonitor) DeepCopyInto(out *TksComplianceMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TksComplianceMonitor.
func (in *TksComplianceMonitor) DeepCopy() *TksComplianceMonitor {
	if in == nil {
		return nil
	}
	out := new(TksComplianceMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TksComplianceMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TksComplianceMonitorList) DeepCopyInto(out *TksComplianceMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TksComplianceMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TksComplianceMonitorList.
func (in *TksComplianceMonitorList) DeepCopy() *TksComplianceMonitorList {
	if in == nil {
		return nil
	}
	out := new(TksComplianceMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TksComplianceMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TksComplianceMonitorSpec) DeepCopyInto(out *TksComplianceMonitorSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TksComplianceMonitorSpec.
func (in *TksComplianceMonitorSpec) DeepCopy() *TksComplianceMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(TksComplianceMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TksComplianceMonitorStatus) DeepCopyInto(out *TksComplianceMonitorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TksComplianceMonitorStatus.
func (in *TksComplianceMonitorStatus) DeepCopy() *TksComplianceMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(TksComplianceMonitorStatus)
	in.DeepCopyInto(out)
	return out
}
