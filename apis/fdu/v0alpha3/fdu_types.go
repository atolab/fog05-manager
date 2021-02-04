/*
* Copyright (c) 2018,2021 ADLINK Technology Inc.
*
* This program and the accompanying materials are made available under the
* terms of the Eclipse Public License 2.0 which is available at
* http://www.eclipse.org/legal/epl-2.0, or the Apache Software License 2.0
* which is available at https://www.apache.org/licenses/LICENSE-2.0.
*
* SPDX-License-Identifier: EPL-2.0 OR Apache-2.0
* Contributors:
*   ADLINK fog05 team, <fog05@adlink-labs.tech>
 */

package v0alpha3

import (
	"github.com/eclipse-fog05/fog05-go/pkg/types/fdu"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FDUStatus defines the observed state of FDU
type FDUStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Record fdu.FDURecord `json:"record,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// FDU is the Schema for the fdus API
type FDU struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   fdu.FDUDescriptor `json:"spec,omitempty"`
	Status FDUStatus         `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FDUList contains a list of FDU
type FDUList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FDU `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FDU{}, &FDUList{})
}
