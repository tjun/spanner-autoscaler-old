/*

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CPUTargetUtilization defines the utilization of Cloud Spanner CPU
type CPUTargetUtilization struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// fraction of the requested CPU that should be utilized/used,
	// e.g. 70 means that 70% of the requested CPU should be in use.
	TargetPercentage int `json:"targetPercentage"`
}

// SpannerAutoscalerSpec defines the desired state of SpannerAutoscaler
type SpannerAutoscalerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Minimum=1
	// lower limit for the number of nodes that can be set by the autoscaler, default 1.
	MinNodes *int `json:"minNodes"`

	// +kubebuilder:validation:Minimum=1
	// upper limit for the number of nodes that can be set by the autoscaler.
	// It cannot be smaller than MinNodes.
	MaxNodes int `json:"maxNodes"`

	// target average CPU utilization for Spanner with priority = high.
	HighPriorityCPUUtilization CPUTargetUtilization `json:"highPriorityCPUUtilization"`
}

// SpannerAutoscalerStatus defines the observed state of SpannerAutoscaler
type SpannerAutoscalerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// last time the SpannerAutoscaler scaled the number of Spanner nodes.
	// used by the autoscaler to controle how often the number of nodes is changed.
	// +optional
	LastScaleTime *metav1.Time `json:"lastScaleTime,omitempty"`

	// current number of nodes of Spanner managed by this autoscaler.
	CurrentNodes int `json:"currentNodes,omitempty"`

	// desired number of nodes of Spanner managed by this autoscaler.
	DesiredNodes int `json:"desiredNodes,omitempty"`

	// current average CPU utilization for high proority task, represented as a percentage.
	CurrentHighPriorityCPUUtilizationPercentage *int `json:"currentHighPriorityCPUUtilizationPercentage,omitempty"`
}

// +kubebuilder:object:root=true

// SpannerAutoscaler is the Schema for the spannerautoscalers API
type SpannerAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpannerAutoscalerSpec   `json:"spec,omitempty"`
	Status SpannerAutoscalerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpannerAutoscalerList contains a list of SpannerAutoscaler
type SpannerAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpannerAutoscaler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpannerAutoscaler{}, &SpannerAutoscalerList{})
}
