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

type RoleMapperInitParameters struct {

	// The ID of the client this role mapper should be added to. Conflicts with client_scope_id. This argument is required if client_scope_id is not set.
	// The destination client of the role. Cannot be used at the same time as client_scope_id.
	// +crossplane:generate:reference:type=github.com/stakater/provider-keycloak/apis/openidclient/v1alpha1.Client
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The ID of the client scope this role mapper should be added to. Conflicts with client_id. This argument is required if client_id is not set.
	// The destination client scope of the role. Cannot be used at the same time as client_id.
	ClientScopeID *string `json:"clientScopeId,omitempty" tf:"client_scope_id,omitempty"`

	// The realm this role mapper exists within.
	// The realm id where the associated client or client scope exists.
	// +crossplane:generate:reference:type=github.com/stakater/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// The ID of the role to be added to this role mapper.
	// Id of the role to assign
	// +crossplane:generate:reference:type=github.com/stakater/provider-keycloak/apis/role/v1alpha1.Role
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Reference to a Role in role to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDRef *v1.Reference `json:"roleIdRef,omitempty" tf:"-"`

	// Selector for a Role in role to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDSelector *v1.Selector `json:"roleIdSelector,omitempty" tf:"-"`
}

type RoleMapperObservation struct {

	// The ID of the client this role mapper should be added to. Conflicts with client_scope_id. This argument is required if client_scope_id is not set.
	// The destination client of the role. Cannot be used at the same time as client_scope_id.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The ID of the client scope this role mapper should be added to. Conflicts with client_id. This argument is required if client_id is not set.
	// The destination client scope of the role. Cannot be used at the same time as client_id.
	ClientScopeID *string `json:"clientScopeId,omitempty" tf:"client_scope_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The realm this role mapper exists within.
	// The realm id where the associated client or client scope exists.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// The ID of the role to be added to this role mapper.
	// Id of the role to assign
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`
}

type RoleMapperParameters struct {

	// The ID of the client this role mapper should be added to. Conflicts with client_scope_id. This argument is required if client_scope_id is not set.
	// The destination client of the role. Cannot be used at the same time as client_scope_id.
	// +crossplane:generate:reference:type=github.com/stakater/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The ID of the client scope this role mapper should be added to. Conflicts with client_id. This argument is required if client_id is not set.
	// The destination client scope of the role. Cannot be used at the same time as client_id.
	// +kubebuilder:validation:Optional
	ClientScopeID *string `json:"clientScopeId,omitempty" tf:"client_scope_id,omitempty"`

	// The realm this role mapper exists within.
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

	// The ID of the role to be added to this role mapper.
	// Id of the role to assign
	// +crossplane:generate:reference:type=github.com/stakater/provider-keycloak/apis/role/v1alpha1.Role
	// +kubebuilder:validation:Optional
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Reference to a Role in role to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDRef *v1.Reference `json:"roleIdRef,omitempty" tf:"-"`

	// Selector for a Role in role to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDSelector *v1.Selector `json:"roleIdSelector,omitempty" tf:"-"`
}

// RoleMapperSpec defines the desired state of RoleMapper
type RoleMapperSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleMapperParameters `json:"forProvider"`
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
	InitProvider RoleMapperInitParameters `json:"initProvider,omitempty"`
}

// RoleMapperStatus defines the observed state of RoleMapper.
type RoleMapperStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleMapperObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RoleMapper is the Schema for the RoleMappers API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type RoleMapper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleMapperSpec   `json:"spec"`
	Status            RoleMapperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleMapperList contains a list of RoleMappers
type RoleMapperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleMapper `json:"items"`
}

// Repository type metadata.
var (
	RoleMapper_Kind             = "RoleMapper"
	RoleMapper_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleMapper_Kind}.String()
	RoleMapper_KindAPIVersion   = RoleMapper_Kind + "." + CRDGroupVersion.String()
	RoleMapper_GroupVersionKind = CRDGroupVersion.WithKind(RoleMapper_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleMapper{}, &RoleMapperList{})
}
