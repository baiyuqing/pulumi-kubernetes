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
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ResourceClaimTemplate is used to produce ResourceClaim objects.
type ResourceClaimTemplate struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// Describes the ResourceClaim that is to be generated.
	//
	// This field is immutable. A ResourceClaim will get created by the control plane for a Pod when needed and then not get updated anymore.
	Spec ResourceClaimTemplateSpecOutput `pulumi:"spec"`
}

// NewResourceClaimTemplate registers a new resource with the given unique name, arguments, and options.
func NewResourceClaimTemplate(ctx *pulumi.Context,
	name string, args *ResourceClaimTemplateArgs, opts ...pulumi.ResourceOption) (*ResourceClaimTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("resource.k8s.io/v1alpha2")
	args.Kind = pulumi.StringPtr("ResourceClaimTemplate")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:resource.k8s.io/v1alpha1:ResourceClaimTemplate"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceClaimTemplate
	err := ctx.RegisterResource("kubernetes:resource.k8s.io/v1alpha2:ResourceClaimTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceClaimTemplate gets an existing ResourceClaimTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceClaimTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceClaimTemplateState, opts ...pulumi.ResourceOption) (*ResourceClaimTemplate, error) {
	var resource ResourceClaimTemplate
	err := ctx.ReadResource("kubernetes:resource.k8s.io/v1alpha2:ResourceClaimTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceClaimTemplate resources.
type resourceClaimTemplateState struct {
}

type ResourceClaimTemplateState struct {
}

func (ResourceClaimTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceClaimTemplateState)(nil)).Elem()
}

type resourceClaimTemplateArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Describes the ResourceClaim that is to be generated.
	//
	// This field is immutable. A ResourceClaim will get created by the control plane for a Pod when needed and then not get updated anymore.
	Spec ResourceClaimTemplateSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ResourceClaimTemplate resource.
type ResourceClaimTemplateArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata
	Metadata metav1.ObjectMetaPtrInput
	// Describes the ResourceClaim that is to be generated.
	//
	// This field is immutable. A ResourceClaim will get created by the control plane for a Pod when needed and then not get updated anymore.
	Spec ResourceClaimTemplateSpecInput
}

func (ResourceClaimTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceClaimTemplateArgs)(nil)).Elem()
}

type ResourceClaimTemplateInput interface {
	pulumi.Input

	ToResourceClaimTemplateOutput() ResourceClaimTemplateOutput
	ToResourceClaimTemplateOutputWithContext(ctx context.Context) ResourceClaimTemplateOutput
}

func (*ResourceClaimTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceClaimTemplate)(nil)).Elem()
}

func (i *ResourceClaimTemplate) ToResourceClaimTemplateOutput() ResourceClaimTemplateOutput {
	return i.ToResourceClaimTemplateOutputWithContext(context.Background())
}

func (i *ResourceClaimTemplate) ToResourceClaimTemplateOutputWithContext(ctx context.Context) ResourceClaimTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimTemplateOutput)
}

func (i *ResourceClaimTemplate) ToOutput(ctx context.Context) pulumix.Output[*ResourceClaimTemplate] {
	return pulumix.Output[*ResourceClaimTemplate]{
		OutputState: i.ToResourceClaimTemplateOutputWithContext(ctx).OutputState,
	}
}

// ResourceClaimTemplateArrayInput is an input type that accepts ResourceClaimTemplateArray and ResourceClaimTemplateArrayOutput values.
// You can construct a concrete instance of `ResourceClaimTemplateArrayInput` via:
//
//	ResourceClaimTemplateArray{ ResourceClaimTemplateArgs{...} }
type ResourceClaimTemplateArrayInput interface {
	pulumi.Input

	ToResourceClaimTemplateArrayOutput() ResourceClaimTemplateArrayOutput
	ToResourceClaimTemplateArrayOutputWithContext(context.Context) ResourceClaimTemplateArrayOutput
}

type ResourceClaimTemplateArray []ResourceClaimTemplateInput

func (ResourceClaimTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceClaimTemplate)(nil)).Elem()
}

func (i ResourceClaimTemplateArray) ToResourceClaimTemplateArrayOutput() ResourceClaimTemplateArrayOutput {
	return i.ToResourceClaimTemplateArrayOutputWithContext(context.Background())
}

func (i ResourceClaimTemplateArray) ToResourceClaimTemplateArrayOutputWithContext(ctx context.Context) ResourceClaimTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimTemplateArrayOutput)
}

func (i ResourceClaimTemplateArray) ToOutput(ctx context.Context) pulumix.Output[[]*ResourceClaimTemplate] {
	return pulumix.Output[[]*ResourceClaimTemplate]{
		OutputState: i.ToResourceClaimTemplateArrayOutputWithContext(ctx).OutputState,
	}
}

// ResourceClaimTemplateMapInput is an input type that accepts ResourceClaimTemplateMap and ResourceClaimTemplateMapOutput values.
// You can construct a concrete instance of `ResourceClaimTemplateMapInput` via:
//
//	ResourceClaimTemplateMap{ "key": ResourceClaimTemplateArgs{...} }
type ResourceClaimTemplateMapInput interface {
	pulumi.Input

	ToResourceClaimTemplateMapOutput() ResourceClaimTemplateMapOutput
	ToResourceClaimTemplateMapOutputWithContext(context.Context) ResourceClaimTemplateMapOutput
}

type ResourceClaimTemplateMap map[string]ResourceClaimTemplateInput

func (ResourceClaimTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceClaimTemplate)(nil)).Elem()
}

func (i ResourceClaimTemplateMap) ToResourceClaimTemplateMapOutput() ResourceClaimTemplateMapOutput {
	return i.ToResourceClaimTemplateMapOutputWithContext(context.Background())
}

func (i ResourceClaimTemplateMap) ToResourceClaimTemplateMapOutputWithContext(ctx context.Context) ResourceClaimTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimTemplateMapOutput)
}

func (i ResourceClaimTemplateMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ResourceClaimTemplate] {
	return pulumix.Output[map[string]*ResourceClaimTemplate]{
		OutputState: i.ToResourceClaimTemplateMapOutputWithContext(ctx).OutputState,
	}
}

type ResourceClaimTemplateOutput struct{ *pulumi.OutputState }

func (ResourceClaimTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceClaimTemplate)(nil)).Elem()
}

func (o ResourceClaimTemplateOutput) ToResourceClaimTemplateOutput() ResourceClaimTemplateOutput {
	return o
}

func (o ResourceClaimTemplateOutput) ToResourceClaimTemplateOutputWithContext(ctx context.Context) ResourceClaimTemplateOutput {
	return o
}

func (o ResourceClaimTemplateOutput) ToOutput(ctx context.Context) pulumix.Output[*ResourceClaimTemplate] {
	return pulumix.Output[*ResourceClaimTemplate]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ResourceClaimTemplateOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceClaimTemplate) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ResourceClaimTemplateOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceClaimTemplate) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object metadata
func (o ResourceClaimTemplateOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *ResourceClaimTemplate) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// Describes the ResourceClaim that is to be generated.
//
// This field is immutable. A ResourceClaim will get created by the control plane for a Pod when needed and then not get updated anymore.
func (o ResourceClaimTemplateOutput) Spec() ResourceClaimTemplateSpecOutput {
	return o.ApplyT(func(v *ResourceClaimTemplate) ResourceClaimTemplateSpecOutput { return v.Spec }).(ResourceClaimTemplateSpecOutput)
}

type ResourceClaimTemplateArrayOutput struct{ *pulumi.OutputState }

func (ResourceClaimTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceClaimTemplate)(nil)).Elem()
}

func (o ResourceClaimTemplateArrayOutput) ToResourceClaimTemplateArrayOutput() ResourceClaimTemplateArrayOutput {
	return o
}

func (o ResourceClaimTemplateArrayOutput) ToResourceClaimTemplateArrayOutputWithContext(ctx context.Context) ResourceClaimTemplateArrayOutput {
	return o
}

func (o ResourceClaimTemplateArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ResourceClaimTemplate] {
	return pulumix.Output[[]*ResourceClaimTemplate]{
		OutputState: o.OutputState,
	}
}

func (o ResourceClaimTemplateArrayOutput) Index(i pulumi.IntInput) ResourceClaimTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceClaimTemplate {
		return vs[0].([]*ResourceClaimTemplate)[vs[1].(int)]
	}).(ResourceClaimTemplateOutput)
}

type ResourceClaimTemplateMapOutput struct{ *pulumi.OutputState }

func (ResourceClaimTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceClaimTemplate)(nil)).Elem()
}

func (o ResourceClaimTemplateMapOutput) ToResourceClaimTemplateMapOutput() ResourceClaimTemplateMapOutput {
	return o
}

func (o ResourceClaimTemplateMapOutput) ToResourceClaimTemplateMapOutputWithContext(ctx context.Context) ResourceClaimTemplateMapOutput {
	return o
}

func (o ResourceClaimTemplateMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ResourceClaimTemplate] {
	return pulumix.Output[map[string]*ResourceClaimTemplate]{
		OutputState: o.OutputState,
	}
}

func (o ResourceClaimTemplateMapOutput) MapIndex(k pulumi.StringInput) ResourceClaimTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceClaimTemplate {
		return vs[0].(map[string]*ResourceClaimTemplate)[vs[1].(string)]
	}).(ResourceClaimTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimTemplateInput)(nil)).Elem(), &ResourceClaimTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimTemplateArrayInput)(nil)).Elem(), ResourceClaimTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimTemplateMapInput)(nil)).Elem(), ResourceClaimTemplateMap{})
	pulumi.RegisterOutputType(ResourceClaimTemplateOutput{})
	pulumi.RegisterOutputType(ResourceClaimTemplateArrayOutput{})
	pulumi.RegisterOutputType(ResourceClaimTemplateMapOutput{})
}
