//go:build !ignore_autogenerated

/*
Copyright 2022.

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

package parser

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSON) DeepCopyInto(out *JSON) {
	*out = *in
	if in.TimeKeep != nil {
		in, out := &in.TimeKeep, &out.TimeKeep
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSON.
func (in *JSON) DeepCopy() *JSON {
	if in == nil {
		return nil
	}
	out := new(JSON)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LSTV) DeepCopyInto(out *LSTV) {
	*out = *in
	if in.TimeKeep != nil {
		in, out := &in.TimeKeep, &out.TimeKeep
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LSTV.
func (in *LSTV) DeepCopy() *LSTV {
	if in == nil {
		return nil
	}
	out := new(LSTV)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logfmt) DeepCopyInto(out *Logfmt) {
	*out = *in
	if in.TimeKeep != nil {
		in, out := &in.TimeKeep, &out.TimeKeep
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logfmt.
func (in *Logfmt) DeepCopy() *Logfmt {
	if in == nil {
		return nil
	}
	out := new(Logfmt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Regex) DeepCopyInto(out *Regex) {
	*out = *in
	if in.TimeKeep != nil {
		in, out := &in.TimeKeep, &out.TimeKeep
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Regex.
func (in *Regex) DeepCopy() *Regex {
	if in == nil {
		return nil
	}
	out := new(Regex)
	in.DeepCopyInto(out)
	return out
}
