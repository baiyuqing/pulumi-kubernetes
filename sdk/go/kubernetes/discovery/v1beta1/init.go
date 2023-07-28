// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

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
	case "kubernetes:discovery.k8s.io/v1beta1:EndpointSlice":
		r = &EndpointSlice{}
	case "kubernetes:discovery.k8s.io/v1beta1:EndpointSliceList":
		r = &EndpointSliceList{}
	case "kubernetes:discovery.k8s.io/v1beta1:EndpointSlicePatch":
		r = &EndpointSlicePatch{}
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
		"discovery.k8s.io/v1beta1",
		&module{version},
	)
}
