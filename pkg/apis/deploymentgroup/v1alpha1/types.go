/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DeploymentGroup is a specification for a DeploymentGroup resource
type DeploymentGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentGroupSpec   `json:"spec"`
	Status DeploymentGroupStatus `json:"status"`
}

// DeploymentGroupSpec is the spec for a DeploymentGroup resource
type DeploymentGroupSpec struct {
	DeploymentName string `json:"deploymentName"`
        Replicas       *int32 `json:"replicas"`
        Template corev1.PodTemplateSpec `json:"template"`
}

// DeploymentGroupStatus is the status for a DeploymentGroup resource
type DeploymentGroupStatus struct {
        AvailableReplicas int32 `json:"AvailableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DeploymentGroupList is a list of DeploymentGroup resources
type DeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DeploymentGroup `json:"items"`
}
