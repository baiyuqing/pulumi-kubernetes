// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/internal"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// PodSchedulingContext objects hold information that is needed to schedule a Pod with ResourceClaims that use "WaitForFirstConsumer" allocation mode.
//
// This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
type PodSchedulingContextPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Spec describes where resources for the Pod are needed.
	Spec PodSchedulingContextSpecPatchPtrOutput `pulumi:"spec"`
	// Status describes where resources for the Pod can be allocated.
	Status PodSchedulingContextStatusPatchPtrOutput `pulumi:"status"`
}

// NewPodSchedulingContextPatch registers a new resource with the given unique name, arguments, and options.
func NewPodSchedulingContextPatch(ctx *pulumi.Context,
	name string, args *PodSchedulingContextPatchArgs, opts ...pulumi.ResourceOption) (*PodSchedulingContextPatch, error) {
	if args == nil {
		args = &PodSchedulingContextPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("resource.k8s.io/v1alpha2")
	args.Kind = pulumi.StringPtr("PodSchedulingContext")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PodSchedulingContextPatch
	err := ctx.RegisterResource("kubernetes:resource.k8s.io/v1alpha2:PodSchedulingContextPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodSchedulingContextPatch gets an existing PodSchedulingContextPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodSchedulingContextPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodSchedulingContextPatchState, opts ...pulumi.ResourceOption) (*PodSchedulingContextPatch, error) {
	var resource PodSchedulingContextPatch
	err := ctx.ReadResource("kubernetes:resource.k8s.io/v1alpha2:PodSchedulingContextPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodSchedulingContextPatch resources.
type podSchedulingContextPatchState struct {
}

type PodSchedulingContextPatchState struct {
}

func (PodSchedulingContextPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*podSchedulingContextPatchState)(nil)).Elem()
}

type podSchedulingContextPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Spec describes where resources for the Pod are needed.
	Spec *PodSchedulingContextSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a PodSchedulingContextPatch resource.
type PodSchedulingContextPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// Spec describes where resources for the Pod are needed.
	Spec PodSchedulingContextSpecPatchPtrInput
}

func (PodSchedulingContextPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podSchedulingContextPatchArgs)(nil)).Elem()
}

type PodSchedulingContextPatchInput interface {
	pulumi.Input

	ToPodSchedulingContextPatchOutput() PodSchedulingContextPatchOutput
	ToPodSchedulingContextPatchOutputWithContext(ctx context.Context) PodSchedulingContextPatchOutput
}

func (*PodSchedulingContextPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**PodSchedulingContextPatch)(nil)).Elem()
}

func (i *PodSchedulingContextPatch) ToPodSchedulingContextPatchOutput() PodSchedulingContextPatchOutput {
	return i.ToPodSchedulingContextPatchOutputWithContext(context.Background())
}

func (i *PodSchedulingContextPatch) ToPodSchedulingContextPatchOutputWithContext(ctx context.Context) PodSchedulingContextPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodSchedulingContextPatchOutput)
}

func (i *PodSchedulingContextPatch) ToOutput(ctx context.Context) pulumix.Output[*PodSchedulingContextPatch] {
	return pulumix.Output[*PodSchedulingContextPatch]{
		OutputState: i.ToPodSchedulingContextPatchOutputWithContext(ctx).OutputState,
	}
}

// PodSchedulingContextPatchArrayInput is an input type that accepts PodSchedulingContextPatchArray and PodSchedulingContextPatchArrayOutput values.
// You can construct a concrete instance of `PodSchedulingContextPatchArrayInput` via:
//
//	PodSchedulingContextPatchArray{ PodSchedulingContextPatchArgs{...} }
type PodSchedulingContextPatchArrayInput interface {
	pulumi.Input

	ToPodSchedulingContextPatchArrayOutput() PodSchedulingContextPatchArrayOutput
	ToPodSchedulingContextPatchArrayOutputWithContext(context.Context) PodSchedulingContextPatchArrayOutput
}

type PodSchedulingContextPatchArray []PodSchedulingContextPatchInput

func (PodSchedulingContextPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodSchedulingContextPatch)(nil)).Elem()
}

func (i PodSchedulingContextPatchArray) ToPodSchedulingContextPatchArrayOutput() PodSchedulingContextPatchArrayOutput {
	return i.ToPodSchedulingContextPatchArrayOutputWithContext(context.Background())
}

func (i PodSchedulingContextPatchArray) ToPodSchedulingContextPatchArrayOutputWithContext(ctx context.Context) PodSchedulingContextPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodSchedulingContextPatchArrayOutput)
}

func (i PodSchedulingContextPatchArray) ToOutput(ctx context.Context) pulumix.Output[[]*PodSchedulingContextPatch] {
	return pulumix.Output[[]*PodSchedulingContextPatch]{
		OutputState: i.ToPodSchedulingContextPatchArrayOutputWithContext(ctx).OutputState,
	}
}

// PodSchedulingContextPatchMapInput is an input type that accepts PodSchedulingContextPatchMap and PodSchedulingContextPatchMapOutput values.
// You can construct a concrete instance of `PodSchedulingContextPatchMapInput` via:
//
//	PodSchedulingContextPatchMap{ "key": PodSchedulingContextPatchArgs{...} }
type PodSchedulingContextPatchMapInput interface {
	pulumi.Input

	ToPodSchedulingContextPatchMapOutput() PodSchedulingContextPatchMapOutput
	ToPodSchedulingContextPatchMapOutputWithContext(context.Context) PodSchedulingContextPatchMapOutput
}

type PodSchedulingContextPatchMap map[string]PodSchedulingContextPatchInput

func (PodSchedulingContextPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodSchedulingContextPatch)(nil)).Elem()
}

func (i PodSchedulingContextPatchMap) ToPodSchedulingContextPatchMapOutput() PodSchedulingContextPatchMapOutput {
	return i.ToPodSchedulingContextPatchMapOutputWithContext(context.Background())
}

func (i PodSchedulingContextPatchMap) ToPodSchedulingContextPatchMapOutputWithContext(ctx context.Context) PodSchedulingContextPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodSchedulingContextPatchMapOutput)
}

func (i PodSchedulingContextPatchMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PodSchedulingContextPatch] {
	return pulumix.Output[map[string]*PodSchedulingContextPatch]{
		OutputState: i.ToPodSchedulingContextPatchMapOutputWithContext(ctx).OutputState,
	}
}

type PodSchedulingContextPatchOutput struct{ *pulumi.OutputState }

func (PodSchedulingContextPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PodSchedulingContextPatch)(nil)).Elem()
}

func (o PodSchedulingContextPatchOutput) ToPodSchedulingContextPatchOutput() PodSchedulingContextPatchOutput {
	return o
}

func (o PodSchedulingContextPatchOutput) ToPodSchedulingContextPatchOutputWithContext(ctx context.Context) PodSchedulingContextPatchOutput {
	return o
}

func (o PodSchedulingContextPatchOutput) ToOutput(ctx context.Context) pulumix.Output[*PodSchedulingContextPatch] {
	return pulumix.Output[*PodSchedulingContextPatch]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PodSchedulingContextPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PodSchedulingContextPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PodSchedulingContextPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PodSchedulingContextPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object metadata
func (o PodSchedulingContextPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *PodSchedulingContextPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Spec describes where resources for the Pod are needed.
func (o PodSchedulingContextPatchOutput) Spec() PodSchedulingContextSpecPatchPtrOutput {
	return o.ApplyT(func(v *PodSchedulingContextPatch) PodSchedulingContextSpecPatchPtrOutput { return v.Spec }).(PodSchedulingContextSpecPatchPtrOutput)
}

// Status describes where resources for the Pod can be allocated.
func (o PodSchedulingContextPatchOutput) Status() PodSchedulingContextStatusPatchPtrOutput {
	return o.ApplyT(func(v *PodSchedulingContextPatch) PodSchedulingContextStatusPatchPtrOutput { return v.Status }).(PodSchedulingContextStatusPatchPtrOutput)
}

type PodSchedulingContextPatchArrayOutput struct{ *pulumi.OutputState }

func (PodSchedulingContextPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodSchedulingContextPatch)(nil)).Elem()
}

func (o PodSchedulingContextPatchArrayOutput) ToPodSchedulingContextPatchArrayOutput() PodSchedulingContextPatchArrayOutput {
	return o
}

func (o PodSchedulingContextPatchArrayOutput) ToPodSchedulingContextPatchArrayOutputWithContext(ctx context.Context) PodSchedulingContextPatchArrayOutput {
	return o
}

func (o PodSchedulingContextPatchArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PodSchedulingContextPatch] {
	return pulumix.Output[[]*PodSchedulingContextPatch]{
		OutputState: o.OutputState,
	}
}

func (o PodSchedulingContextPatchArrayOutput) Index(i pulumi.IntInput) PodSchedulingContextPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PodSchedulingContextPatch {
		return vs[0].([]*PodSchedulingContextPatch)[vs[1].(int)]
	}).(PodSchedulingContextPatchOutput)
}

type PodSchedulingContextPatchMapOutput struct{ *pulumi.OutputState }

func (PodSchedulingContextPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodSchedulingContextPatch)(nil)).Elem()
}

func (o PodSchedulingContextPatchMapOutput) ToPodSchedulingContextPatchMapOutput() PodSchedulingContextPatchMapOutput {
	return o
}

func (o PodSchedulingContextPatchMapOutput) ToPodSchedulingContextPatchMapOutputWithContext(ctx context.Context) PodSchedulingContextPatchMapOutput {
	return o
}

func (o PodSchedulingContextPatchMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PodSchedulingContextPatch] {
	return pulumix.Output[map[string]*PodSchedulingContextPatch]{
		OutputState: o.OutputState,
	}
}

func (o PodSchedulingContextPatchMapOutput) MapIndex(k pulumi.StringInput) PodSchedulingContextPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PodSchedulingContextPatch {
		return vs[0].(map[string]*PodSchedulingContextPatch)[vs[1].(string)]
	}).(PodSchedulingContextPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PodSchedulingContextPatchInput)(nil)).Elem(), &PodSchedulingContextPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodSchedulingContextPatchArrayInput)(nil)).Elem(), PodSchedulingContextPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodSchedulingContextPatchMapInput)(nil)).Elem(), PodSchedulingContextPatchMap{})
	pulumi.RegisterOutputType(PodSchedulingContextPatchOutput{})
	pulumi.RegisterOutputType(PodSchedulingContextPatchArrayOutput{})
	pulumi.RegisterOutputType(PodSchedulingContextPatchMapOutput{})
}
