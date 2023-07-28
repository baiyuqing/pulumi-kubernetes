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
	case "kubernetes:resource.k8s.io/v1alpha1:PodScheduling":
		r = &PodScheduling{}
	case "kubernetes:resource.k8s.io/v1alpha1:PodSchedulingList":
		r = &PodSchedulingList{}
	case "kubernetes:resource.k8s.io/v1alpha1:PodSchedulingPatch":
		r = &PodSchedulingPatch{}
	case "kubernetes:resource.k8s.io/v1alpha1:ResourceClaim":
		r = &ResourceClaim{}
	case "kubernetes:resource.k8s.io/v1alpha1:ResourceClaimList":
		r = &ResourceClaimList{}
	case "kubernetes:resource.k8s.io/v1alpha1:ResourceClaimPatch":
		r = &ResourceClaimPatch{}
	case "kubernetes:resource.k8s.io/v1alpha1:ResourceClaimTemplate":
		r = &ResourceClaimTemplate{}
	case "kubernetes:resource.k8s.io/v1alpha1:ResourceClaimTemplateList":
		r = &ResourceClaimTemplateList{}
	case "kubernetes:resource.k8s.io/v1alpha1:ResourceClaimTemplatePatch":
		r = &ResourceClaimTemplatePatch{}
	case "kubernetes:resource.k8s.io/v1alpha1:ResourceClass":
		r = &ResourceClass{}
	case "kubernetes:resource.k8s.io/v1alpha1:ResourceClassList":
		r = &ResourceClassList{}
	case "kubernetes:resource.k8s.io/v1alpha1:ResourceClassPatch":
		r = &ResourceClassPatch{}
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
		"resource.k8s.io/v1alpha1",
		&module{version},
	)
}
