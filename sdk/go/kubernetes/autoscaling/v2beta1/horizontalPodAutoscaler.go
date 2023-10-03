// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.
type HorizontalPodAutoscaler struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecOutput `pulumi:"spec"`
	// status is the current information about the autoscaler.
	Status HorizontalPodAutoscalerStatusPtrOutput `pulumi:"status"`
}

// NewHorizontalPodAutoscaler registers a new resource with the given unique name, arguments, and options.
func NewHorizontalPodAutoscaler(ctx *pulumi.Context,
	name string, args *HorizontalPodAutoscalerArgs, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscaler, error) {
	if args == nil {
		args = &HorizontalPodAutoscalerArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("autoscaling/v2beta1")
	args.Kind = pulumi.StringPtr("HorizontalPodAutoscaler")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:autoscaling/v1:HorizontalPodAutoscaler"),
		},
		{
			Type: pulumi.String("kubernetes:autoscaling/v2:HorizontalPodAutoscaler"),
		},
		{
			Type: pulumi.String("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscaler"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HorizontalPodAutoscaler
	err := ctx.RegisterResource("kubernetes:autoscaling/v2beta1:HorizontalPodAutoscaler", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHorizontalPodAutoscaler gets an existing HorizontalPodAutoscaler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHorizontalPodAutoscaler(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HorizontalPodAutoscalerState, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscaler, error) {
	var resource HorizontalPodAutoscaler
	err := ctx.ReadResource("kubernetes:autoscaling/v2beta1:HorizontalPodAutoscaler", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HorizontalPodAutoscaler resources.
type horizontalPodAutoscalerState struct {
}

type HorizontalPodAutoscalerState struct {
}

func (HorizontalPodAutoscalerState) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerState)(nil)).Elem()
}

type horizontalPodAutoscalerArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec *HorizontalPodAutoscalerSpec `pulumi:"spec"`
}

// The set of arguments for constructing a HorizontalPodAutoscaler resource.
type HorizontalPodAutoscalerArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecPtrInput
}

func (HorizontalPodAutoscalerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerArgs)(nil)).Elem()
}

type HorizontalPodAutoscalerInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerOutput() HorizontalPodAutoscalerOutput
	ToHorizontalPodAutoscalerOutputWithContext(ctx context.Context) HorizontalPodAutoscalerOutput
}

func (*HorizontalPodAutoscaler) ElementType() reflect.Type {
	return reflect.TypeOf((**HorizontalPodAutoscaler)(nil)).Elem()
}

func (i *HorizontalPodAutoscaler) ToHorizontalPodAutoscalerOutput() HorizontalPodAutoscalerOutput {
	return i.ToHorizontalPodAutoscalerOutputWithContext(context.Background())
}

func (i *HorizontalPodAutoscaler) ToHorizontalPodAutoscalerOutputWithContext(ctx context.Context) HorizontalPodAutoscalerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerOutput)
}

func (i *HorizontalPodAutoscaler) ToOutput(ctx context.Context) pulumix.Output[*HorizontalPodAutoscaler] {
	return pulumix.Output[*HorizontalPodAutoscaler]{
		OutputState: i.ToHorizontalPodAutoscalerOutputWithContext(ctx).OutputState,
	}
}

// HorizontalPodAutoscalerArrayInput is an input type that accepts HorizontalPodAutoscalerArray and HorizontalPodAutoscalerArrayOutput values.
// You can construct a concrete instance of `HorizontalPodAutoscalerArrayInput` via:
//
//	HorizontalPodAutoscalerArray{ HorizontalPodAutoscalerArgs{...} }
type HorizontalPodAutoscalerArrayInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerArrayOutput() HorizontalPodAutoscalerArrayOutput
	ToHorizontalPodAutoscalerArrayOutputWithContext(context.Context) HorizontalPodAutoscalerArrayOutput
}

type HorizontalPodAutoscalerArray []HorizontalPodAutoscalerInput

func (HorizontalPodAutoscalerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HorizontalPodAutoscaler)(nil)).Elem()
}

func (i HorizontalPodAutoscalerArray) ToHorizontalPodAutoscalerArrayOutput() HorizontalPodAutoscalerArrayOutput {
	return i.ToHorizontalPodAutoscalerArrayOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscalerArray) ToHorizontalPodAutoscalerArrayOutputWithContext(ctx context.Context) HorizontalPodAutoscalerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerArrayOutput)
}

func (i HorizontalPodAutoscalerArray) ToOutput(ctx context.Context) pulumix.Output[[]*HorizontalPodAutoscaler] {
	return pulumix.Output[[]*HorizontalPodAutoscaler]{
		OutputState: i.ToHorizontalPodAutoscalerArrayOutputWithContext(ctx).OutputState,
	}
}

// HorizontalPodAutoscalerMapInput is an input type that accepts HorizontalPodAutoscalerMap and HorizontalPodAutoscalerMapOutput values.
// You can construct a concrete instance of `HorizontalPodAutoscalerMapInput` via:
//
//	HorizontalPodAutoscalerMap{ "key": HorizontalPodAutoscalerArgs{...} }
type HorizontalPodAutoscalerMapInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerMapOutput() HorizontalPodAutoscalerMapOutput
	ToHorizontalPodAutoscalerMapOutputWithContext(context.Context) HorizontalPodAutoscalerMapOutput
}

type HorizontalPodAutoscalerMap map[string]HorizontalPodAutoscalerInput

func (HorizontalPodAutoscalerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HorizontalPodAutoscaler)(nil)).Elem()
}

func (i HorizontalPodAutoscalerMap) ToHorizontalPodAutoscalerMapOutput() HorizontalPodAutoscalerMapOutput {
	return i.ToHorizontalPodAutoscalerMapOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscalerMap) ToHorizontalPodAutoscalerMapOutputWithContext(ctx context.Context) HorizontalPodAutoscalerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerMapOutput)
}

func (i HorizontalPodAutoscalerMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*HorizontalPodAutoscaler] {
	return pulumix.Output[map[string]*HorizontalPodAutoscaler]{
		OutputState: i.ToHorizontalPodAutoscalerMapOutputWithContext(ctx).OutputState,
	}
}

type HorizontalPodAutoscalerOutput struct{ *pulumi.OutputState }

func (HorizontalPodAutoscalerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HorizontalPodAutoscaler)(nil)).Elem()
}

func (o HorizontalPodAutoscalerOutput) ToHorizontalPodAutoscalerOutput() HorizontalPodAutoscalerOutput {
	return o
}

func (o HorizontalPodAutoscalerOutput) ToHorizontalPodAutoscalerOutputWithContext(ctx context.Context) HorizontalPodAutoscalerOutput {
	return o
}

func (o HorizontalPodAutoscalerOutput) ToOutput(ctx context.Context) pulumix.Output[*HorizontalPodAutoscaler] {
	return pulumix.Output[*HorizontalPodAutoscaler]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o HorizontalPodAutoscalerOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *HorizontalPodAutoscaler) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o HorizontalPodAutoscalerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *HorizontalPodAutoscaler) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o HorizontalPodAutoscalerOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *HorizontalPodAutoscaler) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
func (o HorizontalPodAutoscalerOutput) Spec() HorizontalPodAutoscalerSpecOutput {
	return o.ApplyT(func(v *HorizontalPodAutoscaler) HorizontalPodAutoscalerSpecOutput { return v.Spec }).(HorizontalPodAutoscalerSpecOutput)
}

// status is the current information about the autoscaler.
func (o HorizontalPodAutoscalerOutput) Status() HorizontalPodAutoscalerStatusPtrOutput {
	return o.ApplyT(func(v *HorizontalPodAutoscaler) HorizontalPodAutoscalerStatusPtrOutput { return v.Status }).(HorizontalPodAutoscalerStatusPtrOutput)
}

type HorizontalPodAutoscalerArrayOutput struct{ *pulumi.OutputState }

func (HorizontalPodAutoscalerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HorizontalPodAutoscaler)(nil)).Elem()
}

func (o HorizontalPodAutoscalerArrayOutput) ToHorizontalPodAutoscalerArrayOutput() HorizontalPodAutoscalerArrayOutput {
	return o
}

func (o HorizontalPodAutoscalerArrayOutput) ToHorizontalPodAutoscalerArrayOutputWithContext(ctx context.Context) HorizontalPodAutoscalerArrayOutput {
	return o
}

func (o HorizontalPodAutoscalerArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*HorizontalPodAutoscaler] {
	return pulumix.Output[[]*HorizontalPodAutoscaler]{
		OutputState: o.OutputState,
	}
}

func (o HorizontalPodAutoscalerArrayOutput) Index(i pulumi.IntInput) HorizontalPodAutoscalerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HorizontalPodAutoscaler {
		return vs[0].([]*HorizontalPodAutoscaler)[vs[1].(int)]
	}).(HorizontalPodAutoscalerOutput)
}

type HorizontalPodAutoscalerMapOutput struct{ *pulumi.OutputState }

func (HorizontalPodAutoscalerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HorizontalPodAutoscaler)(nil)).Elem()
}

func (o HorizontalPodAutoscalerMapOutput) ToHorizontalPodAutoscalerMapOutput() HorizontalPodAutoscalerMapOutput {
	return o
}

func (o HorizontalPodAutoscalerMapOutput) ToHorizontalPodAutoscalerMapOutputWithContext(ctx context.Context) HorizontalPodAutoscalerMapOutput {
	return o
}

func (o HorizontalPodAutoscalerMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*HorizontalPodAutoscaler] {
	return pulumix.Output[map[string]*HorizontalPodAutoscaler]{
		OutputState: o.OutputState,
	}
}

func (o HorizontalPodAutoscalerMapOutput) MapIndex(k pulumi.StringInput) HorizontalPodAutoscalerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HorizontalPodAutoscaler {
		return vs[0].(map[string]*HorizontalPodAutoscaler)[vs[1].(string)]
	}).(HorizontalPodAutoscalerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HorizontalPodAutoscalerInput)(nil)).Elem(), &HorizontalPodAutoscaler{})
	pulumi.RegisterInputType(reflect.TypeOf((*HorizontalPodAutoscalerArrayInput)(nil)).Elem(), HorizontalPodAutoscalerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HorizontalPodAutoscalerMapInput)(nil)).Elem(), HorizontalPodAutoscalerMap{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerArrayOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerMapOutput{})
}
