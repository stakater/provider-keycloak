// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProtocolMapperInitParameters struct {

	// The ID of the client this protocol mapper should be added to. Conflicts with client_scope_id. This argument is required if client_scope_id is not set.
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	// +crossplane:generate:reference:type=github.com/stakater/provider-keycloak/apis/openidclient/v1alpha1.Client
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The ID of the client scope this protocol mapper should be added to. Conflicts with client_id. This argument is required if client_id is not set.
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeID *string `json:"clientScopeId,omitempty" tf:"client_scope_id,omitempty"`

	// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// The display name of this protocol mapper in the GUI.
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of client (either openid-connect or saml). The type must match the type of the client.
	// The protocol of the client (openid-connect / saml).
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
	// The type of the protocol mapper.
	ProtocolMapper *string `json:"protocolMapper,omitempty" tf:"protocol_mapper,omitempty"`

	// The realm this protocol mapper exists within.
	// The realm id where the associated client or client scope exists.
	// +crossplane:generate:reference:type=github.com/stakater/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

type ProtocolMapperObservation struct {

	// The ID of the client this protocol mapper should be added to. Conflicts with client_scope_id. This argument is required if client_scope_id is not set.
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The ID of the client scope this protocol mapper should be added to. Conflicts with client_id. This argument is required if client_id is not set.
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeID *string `json:"clientScopeId,omitempty" tf:"client_scope_id,omitempty"`

	// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The display name of this protocol mapper in the GUI.
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of client (either openid-connect or saml). The type must match the type of the client.
	// The protocol of the client (openid-connect / saml).
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
	// The type of the protocol mapper.
	ProtocolMapper *string `json:"protocolMapper,omitempty" tf:"protocol_mapper,omitempty"`

	// The realm this protocol mapper exists within.
	// The realm id where the associated client or client scope exists.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`
}

type ProtocolMapperParameters struct {

	// The ID of the client this protocol mapper should be added to. Conflicts with client_scope_id. This argument is required if client_scope_id is not set.
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	// +crossplane:generate:reference:type=github.com/stakater/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The ID of the client scope this protocol mapper should be added to. Conflicts with client_id. This argument is required if client_id is not set.
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	// +kubebuilder:validation:Optional
	ClientScopeID *string `json:"clientScopeId,omitempty" tf:"client_scope_id,omitempty"`

	// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// The display name of this protocol mapper in the GUI.
	// A human-friendly name that will appear in the Keycloak console.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of client (either openid-connect or saml). The type must match the type of the client.
	// The protocol of the client (openid-connect / saml).
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
	// The type of the protocol mapper.
	// +kubebuilder:validation:Optional
	ProtocolMapper *string `json:"protocolMapper,omitempty" tf:"protocol_mapper,omitempty"`

	// The realm this protocol mapper exists within.
	// The realm id where the associated client or client scope exists.
	// +crossplane:generate:reference:type=github.com/stakater/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

// ProtocolMapperSpec defines the desired state of ProtocolMapper
type ProtocolMapperSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtocolMapperParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ProtocolMapperInitParameters `json:"initProvider,omitempty"`
}

// ProtocolMapperStatus defines the observed state of ProtocolMapper.
type ProtocolMapperStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtocolMapperObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ProtocolMapper is the Schema for the ProtocolMappers API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type ProtocolMapper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || (has(self.initProvider) && has(self.initProvider.config))",message="spec.forProvider.config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocolMapper) || (has(self.initProvider) && has(self.initProvider.protocolMapper))",message="spec.forProvider.protocolMapper is a required parameter"
	Spec   ProtocolMapperSpec   `json:"spec"`
	Status ProtocolMapperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtocolMapperList contains a list of ProtocolMappers
type ProtocolMapperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtocolMapper `json:"items"`
}

// Repository type metadata.
var (
	ProtocolMapper_Kind             = "ProtocolMapper"
	ProtocolMapper_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProtocolMapper_Kind}.String()
	ProtocolMapper_KindAPIVersion   = ProtocolMapper_Kind + "." + CRDGroupVersion.String()
	ProtocolMapper_GroupVersionKind = CRDGroupVersion.WithKind(ProtocolMapper_Kind)
)

func init() {
	SchemeBuilder.Register(&ProtocolMapper{}, &ProtocolMapperList{})
}
