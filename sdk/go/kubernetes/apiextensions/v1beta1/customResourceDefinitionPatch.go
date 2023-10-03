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
// CustomResourceDefinition represents a resource that should be exposed on the API server.  Its name MUST be in the format <.spec.name>.<.spec.group>. Deprecated in v1.16, planned for removal in v1.19. Use apiextensions.k8s.io/v1 CustomResourceDefinition instead.
type CustomResourceDefinitionPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput          `pulumi:"kind"`
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpecPatchPtrOutput `pulumi:"spec"`
	// status indicates the actual state of the CustomResourceDefinition
	Status CustomResourceDefinitionStatusPatchPtrOutput `pulumi:"status"`
}

// NewCustomResourceDefinitionPatch registers a new resource with the given unique name, arguments, and options.
func NewCustomResourceDefinitionPatch(ctx *pulumi.Context,
	name string, args *CustomResourceDefinitionPatchArgs, opts ...pulumi.ResourceOption) (*CustomResourceDefinitionPatch, error) {
	if args == nil {
		args = &CustomResourceDefinitionPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("apiextensions.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("CustomResourceDefinition")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apiextensions.k8s.io/v1:CustomResourceDefinitionPatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CustomResourceDefinitionPatch
	err := ctx.RegisterResource("kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinitionPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomResourceDefinitionPatch gets an existing CustomResourceDefinitionPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomResourceDefinitionPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomResourceDefinitionPatchState, opts ...pulumi.ResourceOption) (*CustomResourceDefinitionPatch, error) {
	var resource CustomResourceDefinitionPatch
	err := ctx.ReadResource("kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinitionPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomResourceDefinitionPatch resources.
type customResourceDefinitionPatchState struct {
}

type CustomResourceDefinitionPatchState struct {
}

func (CustomResourceDefinitionPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionPatchState)(nil)).Elem()
}

type customResourceDefinitionPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string                 `pulumi:"kind"`
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// spec describes how the user wants the resources to appear
	Spec *CustomResourceDefinitionSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a CustomResourceDefinitionPatch resource.
type CustomResourceDefinitionPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPatchPtrInput
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpecPatchPtrInput
}

func (CustomResourceDefinitionPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionPatchArgs)(nil)).Elem()
}

type CustomResourceDefinitionPatchInput interface {
	pulumi.Input

	ToCustomResourceDefinitionPatchOutput() CustomResourceDefinitionPatchOutput
	ToCustomResourceDefinitionPatchOutputWithContext(ctx context.Context) CustomResourceDefinitionPatchOutput
}

func (*CustomResourceDefinitionPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomResourceDefinitionPatch)(nil)).Elem()
}

func (i *CustomResourceDefinitionPatch) ToCustomResourceDefinitionPatchOutput() CustomResourceDefinitionPatchOutput {
	return i.ToCustomResourceDefinitionPatchOutputWithContext(context.Background())
}

func (i *CustomResourceDefinitionPatch) ToCustomResourceDefinitionPatchOutputWithContext(ctx context.Context) CustomResourceDefinitionPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceDefinitionPatchOutput)
}

func (i *CustomResourceDefinitionPatch) ToOutput(ctx context.Context) pulumix.Output[*CustomResourceDefinitionPatch] {
	return pulumix.Output[*CustomResourceDefinitionPatch]{
		OutputState: i.ToCustomResourceDefinitionPatchOutputWithContext(ctx).OutputState,
	}
}

// CustomResourceDefinitionPatchArrayInput is an input type that accepts CustomResourceDefinitionPatchArray and CustomResourceDefinitionPatchArrayOutput values.
// You can construct a concrete instance of `CustomResourceDefinitionPatchArrayInput` via:
//
//	CustomResourceDefinitionPatchArray{ CustomResourceDefinitionPatchArgs{...} }
type CustomResourceDefinitionPatchArrayInput interface {
	pulumi.Input

	ToCustomResourceDefinitionPatchArrayOutput() CustomResourceDefinitionPatchArrayOutput
	ToCustomResourceDefinitionPatchArrayOutputWithContext(context.Context) CustomResourceDefinitionPatchArrayOutput
}

type CustomResourceDefinitionPatchArray []CustomResourceDefinitionPatchInput

func (CustomResourceDefinitionPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomResourceDefinitionPatch)(nil)).Elem()
}

func (i CustomResourceDefinitionPatchArray) ToCustomResourceDefinitionPatchArrayOutput() CustomResourceDefinitionPatchArrayOutput {
	return i.ToCustomResourceDefinitionPatchArrayOutputWithContext(context.Background())
}

func (i CustomResourceDefinitionPatchArray) ToCustomResourceDefinitionPatchArrayOutputWithContext(ctx context.Context) CustomResourceDefinitionPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceDefinitionPatchArrayOutput)
}

func (i CustomResourceDefinitionPatchArray) ToOutput(ctx context.Context) pulumix.Output[[]*CustomResourceDefinitionPatch] {
	return pulumix.Output[[]*CustomResourceDefinitionPatch]{
		OutputState: i.ToCustomResourceDefinitionPatchArrayOutputWithContext(ctx).OutputState,
	}
}

// CustomResourceDefinitionPatchMapInput is an input type that accepts CustomResourceDefinitionPatchMap and CustomResourceDefinitionPatchMapOutput values.
// You can construct a concrete instance of `CustomResourceDefinitionPatchMapInput` via:
//
//	CustomResourceDefinitionPatchMap{ "key": CustomResourceDefinitionPatchArgs{...} }
type CustomResourceDefinitionPatchMapInput interface {
	pulumi.Input

	ToCustomResourceDefinitionPatchMapOutput() CustomResourceDefinitionPatchMapOutput
	ToCustomResourceDefinitionPatchMapOutputWithContext(context.Context) CustomResourceDefinitionPatchMapOutput
}

type CustomResourceDefinitionPatchMap map[string]CustomResourceDefinitionPatchInput

func (CustomResourceDefinitionPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomResourceDefinitionPatch)(nil)).Elem()
}

func (i CustomResourceDefinitionPatchMap) ToCustomResourceDefinitionPatchMapOutput() CustomResourceDefinitionPatchMapOutput {
	return i.ToCustomResourceDefinitionPatchMapOutputWithContext(context.Background())
}

func (i CustomResourceDefinitionPatchMap) ToCustomResourceDefinitionPatchMapOutputWithContext(ctx context.Context) CustomResourceDefinitionPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceDefinitionPatchMapOutput)
}

func (i CustomResourceDefinitionPatchMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CustomResourceDefinitionPatch] {
	return pulumix.Output[map[string]*CustomResourceDefinitionPatch]{
		OutputState: i.ToCustomResourceDefinitionPatchMapOutputWithContext(ctx).OutputState,
	}
}

type CustomResourceDefinitionPatchOutput struct{ *pulumi.OutputState }

func (CustomResourceDefinitionPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomResourceDefinitionPatch)(nil)).Elem()
}

func (o CustomResourceDefinitionPatchOutput) ToCustomResourceDefinitionPatchOutput() CustomResourceDefinitionPatchOutput {
	return o
}

func (o CustomResourceDefinitionPatchOutput) ToCustomResourceDefinitionPatchOutputWithContext(ctx context.Context) CustomResourceDefinitionPatchOutput {
	return o
}

func (o CustomResourceDefinitionPatchOutput) ToOutput(ctx context.Context) pulumix.Output[*CustomResourceDefinitionPatch] {
	return pulumix.Output[*CustomResourceDefinitionPatch]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CustomResourceDefinitionPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinitionPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CustomResourceDefinitionPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinitionPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o CustomResourceDefinitionPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinitionPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// spec describes how the user wants the resources to appear
func (o CustomResourceDefinitionPatchOutput) Spec() CustomResourceDefinitionSpecPatchPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinitionPatch) CustomResourceDefinitionSpecPatchPtrOutput { return v.Spec }).(CustomResourceDefinitionSpecPatchPtrOutput)
}

// status indicates the actual state of the CustomResourceDefinition
func (o CustomResourceDefinitionPatchOutput) Status() CustomResourceDefinitionStatusPatchPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinitionPatch) CustomResourceDefinitionStatusPatchPtrOutput { return v.Status }).(CustomResourceDefinitionStatusPatchPtrOutput)
}

type CustomResourceDefinitionPatchArrayOutput struct{ *pulumi.OutputState }

func (CustomResourceDefinitionPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomResourceDefinitionPatch)(nil)).Elem()
}

func (o CustomResourceDefinitionPatchArrayOutput) ToCustomResourceDefinitionPatchArrayOutput() CustomResourceDefinitionPatchArrayOutput {
	return o
}

func (o CustomResourceDefinitionPatchArrayOutput) ToCustomResourceDefinitionPatchArrayOutputWithContext(ctx context.Context) CustomResourceDefinitionPatchArrayOutput {
	return o
}

func (o CustomResourceDefinitionPatchArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CustomResourceDefinitionPatch] {
	return pulumix.Output[[]*CustomResourceDefinitionPatch]{
		OutputState: o.OutputState,
	}
}

func (o CustomResourceDefinitionPatchArrayOutput) Index(i pulumi.IntInput) CustomResourceDefinitionPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomResourceDefinitionPatch {
		return vs[0].([]*CustomResourceDefinitionPatch)[vs[1].(int)]
	}).(CustomResourceDefinitionPatchOutput)
}

type CustomResourceDefinitionPatchMapOutput struct{ *pulumi.OutputState }

func (CustomResourceDefinitionPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomResourceDefinitionPatch)(nil)).Elem()
}

func (o CustomResourceDefinitionPatchMapOutput) ToCustomResourceDefinitionPatchMapOutput() CustomResourceDefinitionPatchMapOutput {
	return o
}

func (o CustomResourceDefinitionPatchMapOutput) ToCustomResourceDefinitionPatchMapOutputWithContext(ctx context.Context) CustomResourceDefinitionPatchMapOutput {
	return o
}

func (o CustomResourceDefinitionPatchMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CustomResourceDefinitionPatch] {
	return pulumix.Output[map[string]*CustomResourceDefinitionPatch]{
		OutputState: o.OutputState,
	}
}

func (o CustomResourceDefinitionPatchMapOutput) MapIndex(k pulumi.StringInput) CustomResourceDefinitionPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomResourceDefinitionPatch {
		return vs[0].(map[string]*CustomResourceDefinitionPatch)[vs[1].(string)]
	}).(CustomResourceDefinitionPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomResourceDefinitionPatchInput)(nil)).Elem(), &CustomResourceDefinitionPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomResourceDefinitionPatchArrayInput)(nil)).Elem(), CustomResourceDefinitionPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomResourceDefinitionPatchMapInput)(nil)).Elem(), CustomResourceDefinitionPatchMap{})
	pulumi.RegisterOutputType(CustomResourceDefinitionPatchOutput{})
	pulumi.RegisterOutputType(CustomResourceDefinitionPatchArrayOutput{})
	pulumi.RegisterOutputType(CustomResourceDefinitionPatchMapOutput{})
}
