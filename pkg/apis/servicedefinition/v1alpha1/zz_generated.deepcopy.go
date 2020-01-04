// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CMSSpec) DeepCopyInto(out *CMSSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CMSSpec.
func (in *CMSSpec) DeepCopy() *CMSSpec {
	if in == nil {
		return nil
	}
	out := new(CMSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSpec) DeepCopyInto(out *DatabaseSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSpec.
func (in *DatabaseSpec) DeepCopy() *DatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Servicedefinition) DeepCopyInto(out *Servicedefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Servicedefinition.
func (in *Servicedefinition) DeepCopy() *Servicedefinition {
	if in == nil {
		return nil
	}
	out := new(Servicedefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Servicedefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicedefinitionList) DeepCopyInto(out *ServicedefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Servicedefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicedefinitionList.
func (in *ServicedefinitionList) DeepCopy() *ServicedefinitionList {
	if in == nil {
		return nil
	}
	out := new(ServicedefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicedefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicedefinitionSpec) DeepCopyInto(out *ServicedefinitionSpec) {
	*out = *in
	out.Webserver = in.Webserver
	out.Database = in.Database
	out.CMS = in.CMS
	out.Service = in.Service
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicedefinitionSpec.
func (in *ServicedefinitionSpec) DeepCopy() *ServicedefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(ServicedefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicedefinitionStatus) DeepCopyInto(out *ServicedefinitionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicedefinitionStatus.
func (in *ServicedefinitionStatus) DeepCopy() *ServicedefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(ServicedefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebserverSpec) DeepCopyInto(out *WebserverSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebserverSpec.
func (in *WebserverSpec) DeepCopy() *WebserverSpec {
	if in == nil {
		return nil
	}
	out := new(WebserverSpec)
	in.DeepCopyInto(out)
	return out
}
