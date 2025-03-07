package user

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_user", func(r *config.Resource) {
		r.ShortGroup = "user"
	})

	p.AddResourceConfigurator("keycloak_user_groups", func(r *config.Resource) {
		r.ShortGroup = "user"

		r.References["user_id"] = config.Reference{
			Type: "User",
		}

		r.References["group_ids"] = config.Reference{
			Type: "github.com/stakater/provider-keycloak/apis/group/v1alpha1.Group",
		}
	})
}
