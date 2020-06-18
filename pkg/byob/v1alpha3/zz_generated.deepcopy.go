// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServer) DeepCopyInto(out *APIServer) {
	*out = *in
	if in.AdditionalSANs != nil {
		in, out := &in.AdditionalSANs, &out.AdditionalSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraArguments != nil {
		in, out := &in.ExtraArguments, &out.ExtraArguments
		*out = make([]ServerArgument, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServer.
func (in *APIServer) DeepCopy() *APIServer {
	if in == nil {
		return nil
	}
	out := new(APIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Deps != nil {
		in, out := &in.Deps, &out.Deps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationWebhook) DeepCopyInto(out *AuthenticationWebhook) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationWebhook.
func (in *AuthenticationWebhook) DeepCopy() *AuthenticationWebhook {
	if in == nil {
		return nil
	}
	out := new(AuthenticationWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationWebhook) DeepCopyInto(out *AuthorizationWebhook) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationWebhook.
func (in *AuthorizationWebhook) DeepCopy() *AuthorizationWebhook {
	if in == nil {
		return nil
	}
	out := new(AuthorizationWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOBCluster) DeepCopyInto(out *BYOBCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOBCluster.
func (in *BYOBCluster) DeepCopy() *BYOBCluster {
	if in == nil {
		return nil
	}
	out := new(BYOBCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BYOBCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOBClusterList) DeepCopyInto(out *BYOBClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BYOBCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOBClusterList.
func (in *BYOBClusterList) DeepCopy() *BYOBClusterList {
	if in == nil {
		return nil
	}
	out := new(BYOBClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BYOBClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOBClusterSpec) DeepCopyInto(out *BYOBClusterSpec) {
	*out = *in
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(AuthenticationWebhook)
		**out = **in
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(AuthorizationWebhook)
		**out = **in
	}
	in.OS.DeepCopyInto(&out.OS)
	out.CRI = in.CRI
	in.APIServer.DeepCopyInto(&out.APIServer)
	if in.KubeletArguments != nil {
		in, out := &in.KubeletArguments, &out.KubeletArguments
		*out = make([]ServerArgument, len(*in))
		copy(*out, *in)
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]Addon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOBClusterSpec.
func (in *BYOBClusterSpec) DeepCopy() *BYOBClusterSpec {
	if in == nil {
		return nil
	}
	out := new(BYOBClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOBClusterStatus) DeepCopyInto(out *BYOBClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOBClusterStatus.
func (in *BYOBClusterStatus) DeepCopy() *BYOBClusterStatus {
	if in == nil {
		return nil
	}
	out := new(BYOBClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOBMachine) DeepCopyInto(out *BYOBMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOBMachine.
func (in *BYOBMachine) DeepCopy() *BYOBMachine {
	if in == nil {
		return nil
	}
	out := new(BYOBMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BYOBMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOBMachineList) DeepCopyInto(out *BYOBMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BYOBMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOBMachineList.
func (in *BYOBMachineList) DeepCopy() *BYOBMachineList {
	if in == nil {
		return nil
	}
	out := new(BYOBMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BYOBMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOBMachineSpec) DeepCopyInto(out *BYOBMachineSpec) {
	*out = *in
	out.Private = in.Private
	out.Public = in.Public
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOBMachineSpec.
func (in *BYOBMachineSpec) DeepCopy() *BYOBMachineSpec {
	if in == nil {
		return nil
	}
	out := new(BYOBMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOBMachineStatus) DeepCopyInto(out *BYOBMachineStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOBMachineStatus.
func (in *BYOBMachineStatus) DeepCopy() *BYOBMachineStatus {
	if in == nil {
		return nil
	}
	out := new(BYOBMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntime) DeepCopyInto(out *ContainerRuntime) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntime.
func (in *ContainerRuntime) DeepCopy() *ContainerRuntime {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndPoint) DeepCopyInto(out *EndPoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndPoint.
func (in *EndPoint) DeepCopy() *EndPoint {
	if in == nil {
		return nil
	}
	out := new(EndPoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSpec) DeepCopyInto(out *FileSpec) {
	*out = *in
	out.Source = in.Source
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSpec.
func (in *FileSpec) DeepCopy() *FileSpec {
	if in == nil {
		return nil
	}
	out := new(FileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSConfig) DeepCopyInto(out *OSConfig) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]FileSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSConfig.
func (in *OSConfig) DeepCopy() *OSConfig {
	if in == nil {
		return nil
	}
	out := new(OSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerArgument) DeepCopyInto(out *ServerArgument) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerArgument.
func (in *ServerArgument) DeepCopy() *ServerArgument {
	if in == nil {
		return nil
	}
	out := new(ServerArgument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceSpec) DeepCopyInto(out *SourceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceSpec.
func (in *SourceSpec) DeepCopy() *SourceSpec {
	if in == nil {
		return nil
	}
	out := new(SourceSpec)
	in.DeepCopyInto(out)
	return out
}