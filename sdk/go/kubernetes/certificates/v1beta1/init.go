// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
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
	case "kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequest":
		r = &CertificateSigningRequest{}
	case "kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequestList":
		r = &CertificateSigningRequestList{}
	case "kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequestPatch":
		r = &CertificateSigningRequestPatch{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"kubernetes",
		"certificates.k8s.io/v1beta1",
		&module{version},
	)
}
