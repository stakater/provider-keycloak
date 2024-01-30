/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/stakater/provider-keycloak/apis/realm/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this IdentityProvider.
func (mg *IdentityProvider) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Realm),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmRef,
		Selector:     mg.Spec.ForProvider.RealmSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Realm")
	}
	mg.Spec.ForProvider.Realm = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Realm),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmRef,
		Selector:     mg.Spec.InitProvider.RealmSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Realm")
	}
	mg.Spec.InitProvider.Realm = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmRef = rsp.ResolvedReference

	return nil
}
