// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/internal"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ControllerRevisionList is a resource containing a list of ControllerRevision objects.
type ControllerRevisionList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Items is the list of ControllerRevisions
	Items ControllerRevisionTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewControllerRevisionList registers a new resource with the given unique name, arguments, and options.
func NewControllerRevisionList(ctx *pulumi.Context,
	name string, args *ControllerRevisionListArgs, opts ...pulumi.ResourceOption) (*ControllerRevisionList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("apps/v1beta1")
	args.Kind = pulumi.StringPtr("ControllerRevisionList")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ControllerRevisionList
	err := ctx.RegisterResource("kubernetes:apps/v1beta1:ControllerRevisionList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetControllerRevisionList gets an existing ControllerRevisionList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetControllerRevisionList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControllerRevisionListState, opts ...pulumi.ResourceOption) (*ControllerRevisionList, error) {
	var resource ControllerRevisionList
	err := ctx.ReadResource("kubernetes:apps/v1beta1:ControllerRevisionList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ControllerRevisionList resources.
type controllerRevisionListState struct {
}

type ControllerRevisionListState struct {
}

func (ControllerRevisionListState) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerRevisionListState)(nil)).Elem()
}

type controllerRevisionListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of ControllerRevisions
	Items []ControllerRevisionType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ControllerRevisionList resource.
type ControllerRevisionListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is the list of ControllerRevisions
	Items ControllerRevisionTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (ControllerRevisionListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerRevisionListArgs)(nil)).Elem()
}

type ControllerRevisionListInput interface {
	pulumi.Input

	ToControllerRevisionListOutput() ControllerRevisionListOutput
	ToControllerRevisionListOutputWithContext(ctx context.Context) ControllerRevisionListOutput
}

func (*ControllerRevisionList) ElementType() reflect.Type {
	return reflect.TypeOf((**ControllerRevisionList)(nil)).Elem()
}

func (i *ControllerRevisionList) ToControllerRevisionListOutput() ControllerRevisionListOutput {
	return i.ToControllerRevisionListOutputWithContext(context.Background())
}

func (i *ControllerRevisionList) ToControllerRevisionListOutputWithContext(ctx context.Context) ControllerRevisionListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControllerRevisionListOutput)
}

func (i *ControllerRevisionList) ToOutput(ctx context.Context) pulumix.Output[*ControllerRevisionList] {
	return pulumix.Output[*ControllerRevisionList]{
		OutputState: i.ToControllerRevisionListOutputWithContext(ctx).OutputState,
	}
}

// ControllerRevisionListArrayInput is an input type that accepts ControllerRevisionListArray and ControllerRevisionListArrayOutput values.
// You can construct a concrete instance of `ControllerRevisionListArrayInput` via:
//
//	ControllerRevisionListArray{ ControllerRevisionListArgs{...} }
type ControllerRevisionListArrayInput interface {
	pulumi.Input

	ToControllerRevisionListArrayOutput() ControllerRevisionListArrayOutput
	ToControllerRevisionListArrayOutputWithContext(context.Context) ControllerRevisionListArrayOutput
}

type ControllerRevisionListArray []ControllerRevisionListInput

func (ControllerRevisionListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ControllerRevisionList)(nil)).Elem()
}

func (i ControllerRevisionListArray) ToControllerRevisionListArrayOutput() ControllerRevisionListArrayOutput {
	return i.ToControllerRevisionListArrayOutputWithContext(context.Background())
}

func (i ControllerRevisionListArray) ToControllerRevisionListArrayOutputWithContext(ctx context.Context) ControllerRevisionListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControllerRevisionListArrayOutput)
}

func (i ControllerRevisionListArray) ToOutput(ctx context.Context) pulumix.Output[[]*ControllerRevisionList] {
	return pulumix.Output[[]*ControllerRevisionList]{
		OutputState: i.ToControllerRevisionListArrayOutputWithContext(ctx).OutputState,
	}
}

// ControllerRevisionListMapInput is an input type that accepts ControllerRevisionListMap and ControllerRevisionListMapOutput values.
// You can construct a concrete instance of `ControllerRevisionListMapInput` via:
//
//	ControllerRevisionListMap{ "key": ControllerRevisionListArgs{...} }
type ControllerRevisionListMapInput interface {
	pulumi.Input

	ToControllerRevisionListMapOutput() ControllerRevisionListMapOutput
	ToControllerRevisionListMapOutputWithContext(context.Context) ControllerRevisionListMapOutput
}

type ControllerRevisionListMap map[string]ControllerRevisionListInput

func (ControllerRevisionListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ControllerRevisionList)(nil)).Elem()
}

func (i ControllerRevisionListMap) ToControllerRevisionListMapOutput() ControllerRevisionListMapOutput {
	return i.ToControllerRevisionListMapOutputWithContext(context.Background())
}

func (i ControllerRevisionListMap) ToControllerRevisionListMapOutputWithContext(ctx context.Context) ControllerRevisionListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControllerRevisionListMapOutput)
}

func (i ControllerRevisionListMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ControllerRevisionList] {
	return pulumix.Output[map[string]*ControllerRevisionList]{
		OutputState: i.ToControllerRevisionListMapOutputWithContext(ctx).OutputState,
	}
}

type ControllerRevisionListOutput struct{ *pulumi.OutputState }

func (ControllerRevisionListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControllerRevisionList)(nil)).Elem()
}

func (o ControllerRevisionListOutput) ToControllerRevisionListOutput() ControllerRevisionListOutput {
	return o
}

func (o ControllerRevisionListOutput) ToControllerRevisionListOutputWithContext(ctx context.Context) ControllerRevisionListOutput {
	return o
}

func (o ControllerRevisionListOutput) ToOutput(ctx context.Context) pulumix.Output[*ControllerRevisionList] {
	return pulumix.Output[*ControllerRevisionList]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ControllerRevisionListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ControllerRevisionList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Items is the list of ControllerRevisions
func (o ControllerRevisionListOutput) Items() ControllerRevisionTypeArrayOutput {
	return o.ApplyT(func(v *ControllerRevisionList) ControllerRevisionTypeArrayOutput { return v.Items }).(ControllerRevisionTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ControllerRevisionListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ControllerRevisionList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ControllerRevisionListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ControllerRevisionList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ControllerRevisionListArrayOutput struct{ *pulumi.OutputState }

func (ControllerRevisionListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ControllerRevisionList)(nil)).Elem()
}

func (o ControllerRevisionListArrayOutput) ToControllerRevisionListArrayOutput() ControllerRevisionListArrayOutput {
	return o
}

func (o ControllerRevisionListArrayOutput) ToControllerRevisionListArrayOutputWithContext(ctx context.Context) ControllerRevisionListArrayOutput {
	return o
}

func (o ControllerRevisionListArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ControllerRevisionList] {
	return pulumix.Output[[]*ControllerRevisionList]{
		OutputState: o.OutputState,
	}
}

func (o ControllerRevisionListArrayOutput) Index(i pulumi.IntInput) ControllerRevisionListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ControllerRevisionList {
		return vs[0].([]*ControllerRevisionList)[vs[1].(int)]
	}).(ControllerRevisionListOutput)
}

type ControllerRevisionListMapOutput struct{ *pulumi.OutputState }

func (ControllerRevisionListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ControllerRevisionList)(nil)).Elem()
}

func (o ControllerRevisionListMapOutput) ToControllerRevisionListMapOutput() ControllerRevisionListMapOutput {
	return o
}

func (o ControllerRevisionListMapOutput) ToControllerRevisionListMapOutputWithContext(ctx context.Context) ControllerRevisionListMapOutput {
	return o
}

func (o ControllerRevisionListMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ControllerRevisionList] {
	return pulumix.Output[map[string]*ControllerRevisionList]{
		OutputState: o.OutputState,
	}
}

func (o ControllerRevisionListMapOutput) MapIndex(k pulumi.StringInput) ControllerRevisionListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ControllerRevisionList {
		return vs[0].(map[string]*ControllerRevisionList)[vs[1].(string)]
	}).(ControllerRevisionListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ControllerRevisionListInput)(nil)).Elem(), &ControllerRevisionList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControllerRevisionListArrayInput)(nil)).Elem(), ControllerRevisionListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControllerRevisionListMapInput)(nil)).Elem(), ControllerRevisionListMap{})
	pulumi.RegisterOutputType(ControllerRevisionListOutput{})
	pulumi.RegisterOutputType(ControllerRevisionListArrayOutput{})
	pulumi.RegisterOutputType(ControllerRevisionListMapOutput{})
}
