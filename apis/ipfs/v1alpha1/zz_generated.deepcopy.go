//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPeer) DeepCopyInto(out *ClusterPeer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPeer.
func (in *ClusterPeer) DeepCopy() *ClusterPeer {
	if in == nil {
		return nil
	}
	out := new(ClusterPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPeer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPeerList) DeepCopyInto(out *ClusterPeerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPeerList.
func (in *ClusterPeerList) DeepCopy() *ClusterPeerList {
	if in == nil {
		return nil
	}
	out := new(ClusterPeerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPeerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPeerSpec) DeepCopyInto(out *ClusterPeerSpec) {
	*out = *in
	if in.TrustedPeers != nil {
		in, out := &in.TrustedPeers, &out.TrustedPeers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BootstrapPeers != nil {
		in, out := &in.BootstrapPeers, &out.BootstrapPeers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPeerSpec.
func (in *ClusterPeerSpec) DeepCopy() *ClusterPeerSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterPeerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPeerStatus) DeepCopyInto(out *ClusterPeerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPeerStatus.
func (in *ClusterPeerStatus) DeepCopy() *ClusterPeerStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterPeerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Peer) DeepCopyInto(out *Peer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Peer.
func (in *Peer) DeepCopy() *Peer {
	if in == nil {
		return nil
	}
	out := new(Peer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Peer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerList) DeepCopyInto(out *PeerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Peer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerList.
func (in *PeerList) DeepCopy() *PeerList {
	if in == nil {
		return nil
	}
	out := new(PeerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PeerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerSpec) DeepCopyInto(out *PeerSpec) {
	*out = *in
	if in.InitProfiles != nil {
		in, out := &in.InitProfiles, &out.InitProfiles
		*out = make([]Profile, len(*in))
		copy(*out, *in)
	}
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]Profile, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerSpec.
func (in *PeerSpec) DeepCopy() *PeerSpec {
	if in == nil {
		return nil
	}
	out := new(PeerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerStatus) DeepCopyInto(out *PeerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerStatus.
func (in *PeerStatus) DeepCopy() *PeerStatus {
	if in == nil {
		return nil
	}
	out := new(PeerStatus)
	in.DeepCopyInto(out)
	return out
}
