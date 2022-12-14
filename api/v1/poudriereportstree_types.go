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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PoudrierePortsTreeSpec defines the desired state of PoudrierePortsTree
type PoudrierePortsTreeSpec struct {
	FetchMethod string `json:"fetchmethod,omitempty"`
	Branch      string `json:"branch,omitempty"`
}

// PoudrierePortsTreeStatus defines the observed state of PoudrierePortsTree
type PoudrierePortsTreeStatus struct {
	FetchMethod  string `json:"fetchmethod,omitempty"`
	CreationDate string `json:"creationdate,omitempty"`
	CreationTime string `json:"creationtime,omitempty"`
	Mountpoint   string `json:"mountpoint,omitempty"`
	Ready        bool   `json:"ready,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PoudrierePortsTree is the Schema for the poudriereportstrees API
type PoudrierePortsTree struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PoudrierePortsTreeSpec   `json:"spec,omitempty"`
	Status PoudrierePortsTreeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PoudrierePortsTreeList contains a list of PoudrierePortsTree
type PoudrierePortsTreeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PoudrierePortsTree `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PoudrierePortsTree{}, &PoudrierePortsTreeList{})
}
