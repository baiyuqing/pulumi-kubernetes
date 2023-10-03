// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// Describes a certificate signing request
type CertificateSigningRequestPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput          `pulumi:"kind"`
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// The certificate request itself and any additional information.
	Spec CertificateSigningRequestSpecPatchPtrOutput `pulumi:"spec"`
	// Derived information about the request.
	Status CertificateSigningRequestStatusPatchPtrOutput `pulumi:"status"`
}

// NewCertificateSigningRequestPatch registers a new resource with the given unique name, arguments, and options.
func NewCertificateSigningRequestPatch(ctx *pulumi.Context,
	name string, args *CertificateSigningRequestPatchArgs, opts ...pulumi.ResourceOption) (*CertificateSigningRequestPatch, error) {
	if args == nil {
		args = &CertificateSigningRequestPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("certificates.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("CertificateSigningRequest")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:certificates.k8s.io/v1:CertificateSigningRequestPatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CertificateSigningRequestPatch
	err := ctx.RegisterResource("kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequestPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateSigningRequestPatch gets an existing CertificateSigningRequestPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateSigningRequestPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateSigningRequestPatchState, opts ...pulumi.ResourceOption) (*CertificateSigningRequestPatch, error) {
	var resource CertificateSigningRequestPatch
	err := ctx.ReadResource("kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequestPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateSigningRequestPatch resources.
type certificateSigningRequestPatchState struct {
}

type CertificateSigningRequestPatchState struct {
}

func (CertificateSigningRequestPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestPatchState)(nil)).Elem()
}

type certificateSigningRequestPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string                 `pulumi:"kind"`
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// The certificate request itself and any additional information.
	Spec *CertificateSigningRequestSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a CertificateSigningRequestPatch resource.
type CertificateSigningRequestPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPatchPtrInput
	// The certificate request itself and any additional information.
	Spec CertificateSigningRequestSpecPatchPtrInput
}

func (CertificateSigningRequestPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestPatchArgs)(nil)).Elem()
}

type CertificateSigningRequestPatchInput interface {
	pulumi.Input

	ToCertificateSigningRequestPatchOutput() CertificateSigningRequestPatchOutput
	ToCertificateSigningRequestPatchOutputWithContext(ctx context.Context) CertificateSigningRequestPatchOutput
}

func (*CertificateSigningRequestPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateSigningRequestPatch)(nil)).Elem()
}

func (i *CertificateSigningRequestPatch) ToCertificateSigningRequestPatchOutput() CertificateSigningRequestPatchOutput {
	return i.ToCertificateSigningRequestPatchOutputWithContext(context.Background())
}

func (i *CertificateSigningRequestPatch) ToCertificateSigningRequestPatchOutputWithContext(ctx context.Context) CertificateSigningRequestPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestPatchOutput)
}

func (i *CertificateSigningRequestPatch) ToOutput(ctx context.Context) pulumix.Output[*CertificateSigningRequestPatch] {
	return pulumix.Output[*CertificateSigningRequestPatch]{
		OutputState: i.ToCertificateSigningRequestPatchOutputWithContext(ctx).OutputState,
	}
}

// CertificateSigningRequestPatchArrayInput is an input type that accepts CertificateSigningRequestPatchArray and CertificateSigningRequestPatchArrayOutput values.
// You can construct a concrete instance of `CertificateSigningRequestPatchArrayInput` via:
//
//	CertificateSigningRequestPatchArray{ CertificateSigningRequestPatchArgs{...} }
type CertificateSigningRequestPatchArrayInput interface {
	pulumi.Input

	ToCertificateSigningRequestPatchArrayOutput() CertificateSigningRequestPatchArrayOutput
	ToCertificateSigningRequestPatchArrayOutputWithContext(context.Context) CertificateSigningRequestPatchArrayOutput
}

type CertificateSigningRequestPatchArray []CertificateSigningRequestPatchInput

func (CertificateSigningRequestPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateSigningRequestPatch)(nil)).Elem()
}

func (i CertificateSigningRequestPatchArray) ToCertificateSigningRequestPatchArrayOutput() CertificateSigningRequestPatchArrayOutput {
	return i.ToCertificateSigningRequestPatchArrayOutputWithContext(context.Background())
}

func (i CertificateSigningRequestPatchArray) ToCertificateSigningRequestPatchArrayOutputWithContext(ctx context.Context) CertificateSigningRequestPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestPatchArrayOutput)
}

func (i CertificateSigningRequestPatchArray) ToOutput(ctx context.Context) pulumix.Output[[]*CertificateSigningRequestPatch] {
	return pulumix.Output[[]*CertificateSigningRequestPatch]{
		OutputState: i.ToCertificateSigningRequestPatchArrayOutputWithContext(ctx).OutputState,
	}
}

// CertificateSigningRequestPatchMapInput is an input type that accepts CertificateSigningRequestPatchMap and CertificateSigningRequestPatchMapOutput values.
// You can construct a concrete instance of `CertificateSigningRequestPatchMapInput` via:
//
//	CertificateSigningRequestPatchMap{ "key": CertificateSigningRequestPatchArgs{...} }
type CertificateSigningRequestPatchMapInput interface {
	pulumi.Input

	ToCertificateSigningRequestPatchMapOutput() CertificateSigningRequestPatchMapOutput
	ToCertificateSigningRequestPatchMapOutputWithContext(context.Context) CertificateSigningRequestPatchMapOutput
}

type CertificateSigningRequestPatchMap map[string]CertificateSigningRequestPatchInput

func (CertificateSigningRequestPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateSigningRequestPatch)(nil)).Elem()
}

func (i CertificateSigningRequestPatchMap) ToCertificateSigningRequestPatchMapOutput() CertificateSigningRequestPatchMapOutput {
	return i.ToCertificateSigningRequestPatchMapOutputWithContext(context.Background())
}

func (i CertificateSigningRequestPatchMap) ToCertificateSigningRequestPatchMapOutputWithContext(ctx context.Context) CertificateSigningRequestPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestPatchMapOutput)
}

func (i CertificateSigningRequestPatchMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CertificateSigningRequestPatch] {
	return pulumix.Output[map[string]*CertificateSigningRequestPatch]{
		OutputState: i.ToCertificateSigningRequestPatchMapOutputWithContext(ctx).OutputState,
	}
}

type CertificateSigningRequestPatchOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateSigningRequestPatch)(nil)).Elem()
}

func (o CertificateSigningRequestPatchOutput) ToCertificateSigningRequestPatchOutput() CertificateSigningRequestPatchOutput {
	return o
}

func (o CertificateSigningRequestPatchOutput) ToCertificateSigningRequestPatchOutputWithContext(ctx context.Context) CertificateSigningRequestPatchOutput {
	return o
}

func (o CertificateSigningRequestPatchOutput) ToOutput(ctx context.Context) pulumix.Output[*CertificateSigningRequestPatch] {
	return pulumix.Output[*CertificateSigningRequestPatch]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CertificateSigningRequestPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateSigningRequestPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CertificateSigningRequestPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateSigningRequestPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o CertificateSigningRequestPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *CertificateSigningRequestPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// The certificate request itself and any additional information.
func (o CertificateSigningRequestPatchOutput) Spec() CertificateSigningRequestSpecPatchPtrOutput {
	return o.ApplyT(func(v *CertificateSigningRequestPatch) CertificateSigningRequestSpecPatchPtrOutput { return v.Spec }).(CertificateSigningRequestSpecPatchPtrOutput)
}

// Derived information about the request.
func (o CertificateSigningRequestPatchOutput) Status() CertificateSigningRequestStatusPatchPtrOutput {
	return o.ApplyT(func(v *CertificateSigningRequestPatch) CertificateSigningRequestStatusPatchPtrOutput { return v.Status }).(CertificateSigningRequestStatusPatchPtrOutput)
}

type CertificateSigningRequestPatchArrayOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateSigningRequestPatch)(nil)).Elem()
}

func (o CertificateSigningRequestPatchArrayOutput) ToCertificateSigningRequestPatchArrayOutput() CertificateSigningRequestPatchArrayOutput {
	return o
}

func (o CertificateSigningRequestPatchArrayOutput) ToCertificateSigningRequestPatchArrayOutputWithContext(ctx context.Context) CertificateSigningRequestPatchArrayOutput {
	return o
}

func (o CertificateSigningRequestPatchArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CertificateSigningRequestPatch] {
	return pulumix.Output[[]*CertificateSigningRequestPatch]{
		OutputState: o.OutputState,
	}
}

func (o CertificateSigningRequestPatchArrayOutput) Index(i pulumi.IntInput) CertificateSigningRequestPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertificateSigningRequestPatch {
		return vs[0].([]*CertificateSigningRequestPatch)[vs[1].(int)]
	}).(CertificateSigningRequestPatchOutput)
}

type CertificateSigningRequestPatchMapOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateSigningRequestPatch)(nil)).Elem()
}

func (o CertificateSigningRequestPatchMapOutput) ToCertificateSigningRequestPatchMapOutput() CertificateSigningRequestPatchMapOutput {
	return o
}

func (o CertificateSigningRequestPatchMapOutput) ToCertificateSigningRequestPatchMapOutputWithContext(ctx context.Context) CertificateSigningRequestPatchMapOutput {
	return o
}

func (o CertificateSigningRequestPatchMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CertificateSigningRequestPatch] {
	return pulumix.Output[map[string]*CertificateSigningRequestPatch]{
		OutputState: o.OutputState,
	}
}

func (o CertificateSigningRequestPatchMapOutput) MapIndex(k pulumi.StringInput) CertificateSigningRequestPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertificateSigningRequestPatch {
		return vs[0].(map[string]*CertificateSigningRequestPatch)[vs[1].(string)]
	}).(CertificateSigningRequestPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestPatchInput)(nil)).Elem(), &CertificateSigningRequestPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestPatchArrayInput)(nil)).Elem(), CertificateSigningRequestPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestPatchMapInput)(nil)).Elem(), CertificateSigningRequestPatchMap{})
	pulumi.RegisterOutputType(CertificateSigningRequestPatchOutput{})
	pulumi.RegisterOutputType(CertificateSigningRequestPatchArrayOutput{})
	pulumi.RegisterOutputType(CertificateSigningRequestPatchMapOutput{})
}
