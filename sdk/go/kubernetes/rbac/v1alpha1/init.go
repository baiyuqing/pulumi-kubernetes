// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRole":
		r = &ClusterRole{}
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleBinding":
		r = &ClusterRoleBinding{}
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleBindingList":
		r = &ClusterRoleBindingList{}
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleBindingPatch":
		r = &ClusterRoleBindingPatch{}
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleList":
		r = &ClusterRoleList{}
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRolePatch":
		r = &ClusterRolePatch{}
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:Role":
		r = &Role{}
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBinding":
		r = &RoleBinding{}
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBindingList":
		r = &RoleBindingList{}
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBindingPatch":
		r = &RoleBindingPatch{}
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleList":
		r = &RoleList{}
	case "kubernetes:rbac.authorization.k8s.io/v1alpha1:RolePatch":
		r = &RolePatch{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"kubernetes",
		"rbac.authorization.k8s.io/v1alpha1",
		&module{version},
	)
}
