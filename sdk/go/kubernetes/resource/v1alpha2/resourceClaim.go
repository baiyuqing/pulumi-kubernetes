// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/internal"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ResourceClaim describes which resources are needed by a resource consumer. Its status tracks whether the resource has been allocated and what the resulting attributes are.
//
// This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
type ResourceClaim struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// Spec describes the desired attributes of a resource that then needs to be allocated. It can only be set once when creating the ResourceClaim.
	Spec ResourceClaimSpecOutput `pulumi:"spec"`
	// Status describes whether the resource is available and with which attributes.
	Status ResourceClaimStatusPtrOutput `pulumi:"status"`
}

// NewResourceClaim registers a new resource with the given unique name, arguments, and options.
func NewResourceClaim(ctx *pulumi.Context,
	name string, args *ResourceClaimArgs, opts ...pulumi.ResourceOption) (*ResourceClaim, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("resource.k8s.io/v1alpha2")
	args.Kind = pulumi.StringPtr("ResourceClaim")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:resource.k8s.io/v1alpha1:ResourceClaim"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceClaim
	err := ctx.RegisterResource("kubernetes:resource.k8s.io/v1alpha2:ResourceClaim", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceClaim gets an existing ResourceClaim resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceClaim(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceClaimState, opts ...pulumi.ResourceOption) (*ResourceClaim, error) {
	var resource ResourceClaim
	err := ctx.ReadResource("kubernetes:resource.k8s.io/v1alpha2:ResourceClaim", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceClaim resources.
type resourceClaimState struct {
}

type ResourceClaimState struct {
}

func (ResourceClaimState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceClaimState)(nil)).Elem()
}

type resourceClaimArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec describes the desired attributes of a resource that then needs to be allocated. It can only be set once when creating the ResourceClaim.
	Spec ResourceClaimSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ResourceClaim resource.
type ResourceClaimArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec describes the desired attributes of a resource that then needs to be allocated. It can only be set once when creating the ResourceClaim.
	Spec ResourceClaimSpecInput
}

func (ResourceClaimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceClaimArgs)(nil)).Elem()
}

type ResourceClaimInput interface {
	pulumi.Input

	ToResourceClaimOutput() ResourceClaimOutput
	ToResourceClaimOutputWithContext(ctx context.Context) ResourceClaimOutput
}

func (*ResourceClaim) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceClaim)(nil)).Elem()
}

func (i *ResourceClaim) ToResourceClaimOutput() ResourceClaimOutput {
	return i.ToResourceClaimOutputWithContext(context.Background())
}

func (i *ResourceClaim) ToResourceClaimOutputWithContext(ctx context.Context) ResourceClaimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimOutput)
}

// ResourceClaimArrayInput is an input type that accepts ResourceClaimArray and ResourceClaimArrayOutput values.
// You can construct a concrete instance of `ResourceClaimArrayInput` via:
//
//	ResourceClaimArray{ ResourceClaimArgs{...} }
type ResourceClaimArrayInput interface {
	pulumi.Input

	ToResourceClaimArrayOutput() ResourceClaimArrayOutput
	ToResourceClaimArrayOutputWithContext(context.Context) ResourceClaimArrayOutput
}

type ResourceClaimArray []ResourceClaimInput

func (ResourceClaimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceClaim)(nil)).Elem()
}

func (i ResourceClaimArray) ToResourceClaimArrayOutput() ResourceClaimArrayOutput {
	return i.ToResourceClaimArrayOutputWithContext(context.Background())
}

func (i ResourceClaimArray) ToResourceClaimArrayOutputWithContext(ctx context.Context) ResourceClaimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimArrayOutput)
}

// ResourceClaimMapInput is an input type that accepts ResourceClaimMap and ResourceClaimMapOutput values.
// You can construct a concrete instance of `ResourceClaimMapInput` via:
//
//	ResourceClaimMap{ "key": ResourceClaimArgs{...} }
type ResourceClaimMapInput interface {
	pulumi.Input

	ToResourceClaimMapOutput() ResourceClaimMapOutput
	ToResourceClaimMapOutputWithContext(context.Context) ResourceClaimMapOutput
}

type ResourceClaimMap map[string]ResourceClaimInput

func (ResourceClaimMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceClaim)(nil)).Elem()
}

func (i ResourceClaimMap) ToResourceClaimMapOutput() ResourceClaimMapOutput {
	return i.ToResourceClaimMapOutputWithContext(context.Background())
}

func (i ResourceClaimMap) ToResourceClaimMapOutputWithContext(ctx context.Context) ResourceClaimMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimMapOutput)
}

type ResourceClaimOutput struct{ *pulumi.OutputState }

func (ResourceClaimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceClaim)(nil)).Elem()
}

func (o ResourceClaimOutput) ToResourceClaimOutput() ResourceClaimOutput {
	return o
}

func (o ResourceClaimOutput) ToResourceClaimOutputWithContext(ctx context.Context) ResourceClaimOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ResourceClaimOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceClaim) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ResourceClaimOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceClaim) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object metadata
func (o ResourceClaimOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *ResourceClaim) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// Spec describes the desired attributes of a resource that then needs to be allocated. It can only be set once when creating the ResourceClaim.
func (o ResourceClaimOutput) Spec() ResourceClaimSpecOutput {
	return o.ApplyT(func(v *ResourceClaim) ResourceClaimSpecOutput { return v.Spec }).(ResourceClaimSpecOutput)
}

// Status describes whether the resource is available and with which attributes.
func (o ResourceClaimOutput) Status() ResourceClaimStatusPtrOutput {
	return o.ApplyT(func(v *ResourceClaim) ResourceClaimStatusPtrOutput { return v.Status }).(ResourceClaimStatusPtrOutput)
}

type ResourceClaimArrayOutput struct{ *pulumi.OutputState }

func (ResourceClaimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceClaim)(nil)).Elem()
}

func (o ResourceClaimArrayOutput) ToResourceClaimArrayOutput() ResourceClaimArrayOutput {
	return o
}

func (o ResourceClaimArrayOutput) ToResourceClaimArrayOutputWithContext(ctx context.Context) ResourceClaimArrayOutput {
	return o
}

func (o ResourceClaimArrayOutput) Index(i pulumi.IntInput) ResourceClaimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceClaim {
		return vs[0].([]*ResourceClaim)[vs[1].(int)]
	}).(ResourceClaimOutput)
}

type ResourceClaimMapOutput struct{ *pulumi.OutputState }

func (ResourceClaimMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceClaim)(nil)).Elem()
}

func (o ResourceClaimMapOutput) ToResourceClaimMapOutput() ResourceClaimMapOutput {
	return o
}

func (o ResourceClaimMapOutput) ToResourceClaimMapOutputWithContext(ctx context.Context) ResourceClaimMapOutput {
	return o
}

func (o ResourceClaimMapOutput) MapIndex(k pulumi.StringInput) ResourceClaimOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceClaim {
		return vs[0].(map[string]*ResourceClaim)[vs[1].(string)]
	}).(ResourceClaimOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimInput)(nil)).Elem(), &ResourceClaim{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimArrayInput)(nil)).Elem(), ResourceClaimArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimMapInput)(nil)).Elem(), ResourceClaimMap{})
	pulumi.RegisterOutputType(ResourceClaimOutput{})
	pulumi.RegisterOutputType(ResourceClaimArrayOutput{})
	pulumi.RegisterOutputType(ResourceClaimMapOutput{})
}
