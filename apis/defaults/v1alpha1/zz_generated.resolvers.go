/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha11 "github.com/stakater/provider-keycloak/apis/realm/v1alpha1"
	v1alpha1 "github.com/stakater/provider-keycloak/apis/role/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Roles.
func (mg *Roles) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.DefaultRoles),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.DefaultRolesRefs,
		Selector:      mg.Spec.ForProvider.DefaultRolesSelector,
		To: reference.To{
			List:    &v1alpha1.RoleList{},
			Managed: &v1alpha1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DefaultRoles")
	}
	mg.Spec.ForProvider.DefaultRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.DefaultRolesRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha11.RealmList{},
			Managed: &v1alpha11.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.DefaultRoles),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.DefaultRolesRefs,
		Selector:      mg.Spec.InitProvider.DefaultRolesSelector,
		To: reference.To{
			List:    &v1alpha1.RoleList{},
			Managed: &v1alpha1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DefaultRoles")
	}
	mg.Spec.InitProvider.DefaultRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.DefaultRolesRefs = mrsp.ResolvedReferences

	return nil
}
