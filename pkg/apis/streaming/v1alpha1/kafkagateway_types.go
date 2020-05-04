/*
Copyright 2019 the original author or authors.

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
	"github.com/projectriff/reconciler-runtime/apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	sapis "github.com/projectriff/system/pkg/apis"
	duckv1 "github.com/projectriff/system/pkg/apis/duck/v1"
	"github.com/projectriff/system/pkg/refs"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

var (
	KafkaGatewayLabelKey = GroupVersion.Group + "/kafka-gateway"
	KafkaGatewayType     = "kafka"
)

var (
	_ sapis.Resource = (*KafkaGateway)(nil)
)

// KafkaGatewaySpec defines the desired state of KafkaGateway
type KafkaGatewaySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// BootstrapServers is a comma-separated list of host and port pairs that are the
	// addresses of the Kafka brokers in a "bootstrap" Kafka cluster that a Kafka client
	// connects to initially to bootstrap itself.
	//
	// A host and port pair uses `:` as the separator.
	BootstrapServers string `json:"bootstrapServers"`
}

// KafkaGatewayStatus defines the observed state of KafkaGateway
type KafkaGatewayStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	apis.Status      `json:",inline"`
	Address          *duckv1.Addressable             `json:"address,omitempty"`
	GatewayRef       *refs.TypedLocalObjectReference `json:"gatewayRef,omitempty"`
	GatewayImage     string                          `json:"gatewayImage,omitempty"`
	ProvisionerImage string                          `json:"provisionerImage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories="riff"
// +kubebuilder:printcolumn:name="Bootstrap Servers",type=string,JSONPath=`.spec.bootstrapServers`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Reason",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +genclient

// KafkaGateway is the Schema for the gateways API
type KafkaGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KafkaGatewaySpec   `json:"spec,omitempty"`
	Status KafkaGatewayStatus `json:"status,omitempty"`
}

func (*KafkaGateway) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("KafkaGateway")
}

func (p *KafkaGateway) GetStatus() sapis.ResourceStatus {
	return &p.Status
}

// +kubebuilder:object:root=true

// KafkaGatewayList contains a list of KafkaGateway
type KafkaGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KafkaGateway{}, &KafkaGatewayList{})
}
