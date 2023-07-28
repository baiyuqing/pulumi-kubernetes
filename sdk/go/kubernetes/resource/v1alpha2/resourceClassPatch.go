// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/internal"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// ResourceClass is used by administrators to influence how resources are allocated.
//
// This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
type ResourceClassPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// DriverName defines the name of the dynamic resource driver that is used for allocation of a ResourceClaim that uses this class.
	//
	// Resource drivers have a unique name in forward domain order (acme.example.com).
	DriverName pulumi.StringPtrOutput `pulumi:"driverName"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// ParametersRef references an arbitrary separate object that may hold parameters that will be used by the driver when allocating a resource that uses this class. A dynamic resource driver can distinguish between parameters stored here and and those stored in ResourceClaimSpec.
	ParametersRef ResourceClassParametersReferencePatchPtrOutput `pulumi:"parametersRef"`
	// Only nodes matching the selector will be considered by the scheduler when trying to find a Node that fits a Pod when that Pod uses a ResourceClaim that has not been allocated yet.
	//
	// Setting this field is optional. If null, all nodes are candidates.
	SuitableNodes corev1.NodeSelectorPatchPtrOutput `pulumi:"suitableNodes"`
}

// NewResourceClassPatch registers a new resource with the given unique name, arguments, and options.
func NewResourceClassPatch(ctx *pulumi.Context,
	name string, args *ResourceClassPatchArgs, opts ...pulumi.ResourceOption) (*ResourceClassPatch, error) {
	if args == nil {
		args = &ResourceClassPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("resource.k8s.io/v1alpha2")
	args.Kind = pulumi.StringPtr("ResourceClass")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:resource.k8s.io/v1alpha1:ResourceClassPatch"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceClassPatch
	err := ctx.RegisterResource("kubernetes:resource.k8s.io/v1alpha2:ResourceClassPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceClassPatch gets an existing ResourceClassPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceClassPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceClassPatchState, opts ...pulumi.ResourceOption) (*ResourceClassPatch, error) {
	var resource ResourceClassPatch
	err := ctx.ReadResource("kubernetes:resource.k8s.io/v1alpha2:ResourceClassPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceClassPatch resources.
type resourceClassPatchState struct {
}

type ResourceClassPatchState struct {
}

func (ResourceClassPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceClassPatchState)(nil)).Elem()
}

type resourceClassPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// DriverName defines the name of the dynamic resource driver that is used for allocation of a ResourceClaim that uses this class.
	//
	// Resource drivers have a unique name in forward domain order (acme.example.com).
	DriverName *string `pulumi:"driverName"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// ParametersRef references an arbitrary separate object that may hold parameters that will be used by the driver when allocating a resource that uses this class. A dynamic resource driver can distinguish between parameters stored here and and those stored in ResourceClaimSpec.
	ParametersRef *ResourceClassParametersReferencePatch `pulumi:"parametersRef"`
	// Only nodes matching the selector will be considered by the scheduler when trying to find a Node that fits a Pod when that Pod uses a ResourceClaim that has not been allocated yet.
	//
	// Setting this field is optional. If null, all nodes are candidates.
	SuitableNodes *corev1.NodeSelectorPatch `pulumi:"suitableNodes"`
}

// The set of arguments for constructing a ResourceClassPatch resource.
type ResourceClassPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// DriverName defines the name of the dynamic resource driver that is used for allocation of a ResourceClaim that uses this class.
	//
	// Resource drivers have a unique name in forward domain order (acme.example.com).
	DriverName pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// ParametersRef references an arbitrary separate object that may hold parameters that will be used by the driver when allocating a resource that uses this class. A dynamic resource driver can distinguish between parameters stored here and and those stored in ResourceClaimSpec.
	ParametersRef ResourceClassParametersReferencePatchPtrInput
	// Only nodes matching the selector will be considered by the scheduler when trying to find a Node that fits a Pod when that Pod uses a ResourceClaim that has not been allocated yet.
	//
	// Setting this field is optional. If null, all nodes are candidates.
	SuitableNodes corev1.NodeSelectorPatchPtrInput
}

func (ResourceClassPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceClassPatchArgs)(nil)).Elem()
}

type ResourceClassPatchInput interface {
	pulumi.Input

	ToResourceClassPatchOutput() ResourceClassPatchOutput
	ToResourceClassPatchOutputWithContext(ctx context.Context) ResourceClassPatchOutput
}

func (*ResourceClassPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceClassPatch)(nil)).Elem()
}

func (i *ResourceClassPatch) ToResourceClassPatchOutput() ResourceClassPatchOutput {
	return i.ToResourceClassPatchOutputWithContext(context.Background())
}

func (i *ResourceClassPatch) ToResourceClassPatchOutputWithContext(ctx context.Context) ResourceClassPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClassPatchOutput)
}

// ResourceClassPatchArrayInput is an input type that accepts ResourceClassPatchArray and ResourceClassPatchArrayOutput values.
// You can construct a concrete instance of `ResourceClassPatchArrayInput` via:
//
//	ResourceClassPatchArray{ ResourceClassPatchArgs{...} }
type ResourceClassPatchArrayInput interface {
	pulumi.Input

	ToResourceClassPatchArrayOutput() ResourceClassPatchArrayOutput
	ToResourceClassPatchArrayOutputWithContext(context.Context) ResourceClassPatchArrayOutput
}

type ResourceClassPatchArray []ResourceClassPatchInput

func (ResourceClassPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceClassPatch)(nil)).Elem()
}

func (i ResourceClassPatchArray) ToResourceClassPatchArrayOutput() ResourceClassPatchArrayOutput {
	return i.ToResourceClassPatchArrayOutputWithContext(context.Background())
}

func (i ResourceClassPatchArray) ToResourceClassPatchArrayOutputWithContext(ctx context.Context) ResourceClassPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClassPatchArrayOutput)
}

// ResourceClassPatchMapInput is an input type that accepts ResourceClassPatchMap and ResourceClassPatchMapOutput values.
// You can construct a concrete instance of `ResourceClassPatchMapInput` via:
//
//	ResourceClassPatchMap{ "key": ResourceClassPatchArgs{...} }
type ResourceClassPatchMapInput interface {
	pulumi.Input

	ToResourceClassPatchMapOutput() ResourceClassPatchMapOutput
	ToResourceClassPatchMapOutputWithContext(context.Context) ResourceClassPatchMapOutput
}

type ResourceClassPatchMap map[string]ResourceClassPatchInput

func (ResourceClassPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceClassPatch)(nil)).Elem()
}

func (i ResourceClassPatchMap) ToResourceClassPatchMapOutput() ResourceClassPatchMapOutput {
	return i.ToResourceClassPatchMapOutputWithContext(context.Background())
}

func (i ResourceClassPatchMap) ToResourceClassPatchMapOutputWithContext(ctx context.Context) ResourceClassPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClassPatchMapOutput)
}

type ResourceClassPatchOutput struct{ *pulumi.OutputState }

func (ResourceClassPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceClassPatch)(nil)).Elem()
}

func (o ResourceClassPatchOutput) ToResourceClassPatchOutput() ResourceClassPatchOutput {
	return o
}

func (o ResourceClassPatchOutput) ToResourceClassPatchOutputWithContext(ctx context.Context) ResourceClassPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ResourceClassPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceClassPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// DriverName defines the name of the dynamic resource driver that is used for allocation of a ResourceClaim that uses this class.
//
// Resource drivers have a unique name in forward domain order (acme.example.com).
func (o ResourceClassPatchOutput) DriverName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceClassPatch) pulumi.StringPtrOutput { return v.DriverName }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ResourceClassPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceClassPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object metadata
func (o ResourceClassPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ResourceClassPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// ParametersRef references an arbitrary separate object that may hold parameters that will be used by the driver when allocating a resource that uses this class. A dynamic resource driver can distinguish between parameters stored here and and those stored in ResourceClaimSpec.
func (o ResourceClassPatchOutput) ParametersRef() ResourceClassParametersReferencePatchPtrOutput {
	return o.ApplyT(func(v *ResourceClassPatch) ResourceClassParametersReferencePatchPtrOutput { return v.ParametersRef }).(ResourceClassParametersReferencePatchPtrOutput)
}

// Only nodes matching the selector will be considered by the scheduler when trying to find a Node that fits a Pod when that Pod uses a ResourceClaim that has not been allocated yet.
//
// Setting this field is optional. If null, all nodes are candidates.
func (o ResourceClassPatchOutput) SuitableNodes() corev1.NodeSelectorPatchPtrOutput {
	return o.ApplyT(func(v *ResourceClassPatch) corev1.NodeSelectorPatchPtrOutput { return v.SuitableNodes }).(corev1.NodeSelectorPatchPtrOutput)
}

type ResourceClassPatchArrayOutput struct{ *pulumi.OutputState }

func (ResourceClassPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceClassPatch)(nil)).Elem()
}

func (o ResourceClassPatchArrayOutput) ToResourceClassPatchArrayOutput() ResourceClassPatchArrayOutput {
	return o
}

func (o ResourceClassPatchArrayOutput) ToResourceClassPatchArrayOutputWithContext(ctx context.Context) ResourceClassPatchArrayOutput {
	return o
}

func (o ResourceClassPatchArrayOutput) Index(i pulumi.IntInput) ResourceClassPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceClassPatch {
		return vs[0].([]*ResourceClassPatch)[vs[1].(int)]
	}).(ResourceClassPatchOutput)
}

type ResourceClassPatchMapOutput struct{ *pulumi.OutputState }

func (ResourceClassPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceClassPatch)(nil)).Elem()
}

func (o ResourceClassPatchMapOutput) ToResourceClassPatchMapOutput() ResourceClassPatchMapOutput {
	return o
}

func (o ResourceClassPatchMapOutput) ToResourceClassPatchMapOutputWithContext(ctx context.Context) ResourceClassPatchMapOutput {
	return o
}

func (o ResourceClassPatchMapOutput) MapIndex(k pulumi.StringInput) ResourceClassPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceClassPatch {
		return vs[0].(map[string]*ResourceClassPatch)[vs[1].(string)]
	}).(ResourceClassPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClassPatchInput)(nil)).Elem(), &ResourceClassPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClassPatchArrayInput)(nil)).Elem(), ResourceClassPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClassPatchMapInput)(nil)).Elem(), ResourceClassPatchMap{})
	pulumi.RegisterOutputType(ResourceClassPatchOutput{})
	pulumi.RegisterOutputType(ResourceClassPatchArrayOutput{})
	pulumi.RegisterOutputType(ResourceClassPatchMapOutput{})
}
