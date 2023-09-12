// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/internal"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// NamespaceList is a list of Namespaces.
type NamespaceList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Items NamespaceTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewNamespaceList registers a new resource with the given unique name, arguments, and options.
func NewNamespaceList(ctx *pulumi.Context,
	name string, args *NamespaceListArgs, opts ...pulumi.ResourceOption) (*NamespaceList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("NamespaceList")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NamespaceList
	err := ctx.RegisterResource("kubernetes:core/v1:NamespaceList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceList gets an existing NamespaceList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceListState, opts ...pulumi.ResourceOption) (*NamespaceList, error) {
	var resource NamespaceList
	err := ctx.ReadResource("kubernetes:core/v1:NamespaceList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceList resources.
type namespaceListState struct {
}

type NamespaceListState struct {
}

func (NamespaceListState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceListState)(nil)).Elem()
}

type namespaceListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Items []NamespaceType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a NamespaceList resource.
type NamespaceListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Items NamespaceTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (NamespaceListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceListArgs)(nil)).Elem()
}

type NamespaceListInput interface {
	pulumi.Input

	ToNamespaceListOutput() NamespaceListOutput
	ToNamespaceListOutputWithContext(ctx context.Context) NamespaceListOutput
}

func (*NamespaceList) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceList)(nil)).Elem()
}

func (i *NamespaceList) ToNamespaceListOutput() NamespaceListOutput {
	return i.ToNamespaceListOutputWithContext(context.Background())
}

func (i *NamespaceList) ToNamespaceListOutputWithContext(ctx context.Context) NamespaceListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceListOutput)
}

func (i *NamespaceList) ToOutput(ctx context.Context) pulumix.Output[*NamespaceList] {
	return pulumix.Output[*NamespaceList]{
		OutputState: i.ToNamespaceListOutputWithContext(ctx).OutputState,
	}
}

// NamespaceListArrayInput is an input type that accepts NamespaceListArray and NamespaceListArrayOutput values.
// You can construct a concrete instance of `NamespaceListArrayInput` via:
//
//	NamespaceListArray{ NamespaceListArgs{...} }
type NamespaceListArrayInput interface {
	pulumi.Input

	ToNamespaceListArrayOutput() NamespaceListArrayOutput
	ToNamespaceListArrayOutputWithContext(context.Context) NamespaceListArrayOutput
}

type NamespaceListArray []NamespaceListInput

func (NamespaceListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamespaceList)(nil)).Elem()
}

func (i NamespaceListArray) ToNamespaceListArrayOutput() NamespaceListArrayOutput {
	return i.ToNamespaceListArrayOutputWithContext(context.Background())
}

func (i NamespaceListArray) ToNamespaceListArrayOutputWithContext(ctx context.Context) NamespaceListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceListArrayOutput)
}

func (i NamespaceListArray) ToOutput(ctx context.Context) pulumix.Output[[]*NamespaceList] {
	return pulumix.Output[[]*NamespaceList]{
		OutputState: i.ToNamespaceListArrayOutputWithContext(ctx).OutputState,
	}
}

// NamespaceListMapInput is an input type that accepts NamespaceListMap and NamespaceListMapOutput values.
// You can construct a concrete instance of `NamespaceListMapInput` via:
//
//	NamespaceListMap{ "key": NamespaceListArgs{...} }
type NamespaceListMapInput interface {
	pulumi.Input

	ToNamespaceListMapOutput() NamespaceListMapOutput
	ToNamespaceListMapOutputWithContext(context.Context) NamespaceListMapOutput
}

type NamespaceListMap map[string]NamespaceListInput

func (NamespaceListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamespaceList)(nil)).Elem()
}

func (i NamespaceListMap) ToNamespaceListMapOutput() NamespaceListMapOutput {
	return i.ToNamespaceListMapOutputWithContext(context.Background())
}

func (i NamespaceListMap) ToNamespaceListMapOutputWithContext(ctx context.Context) NamespaceListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceListMapOutput)
}

func (i NamespaceListMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*NamespaceList] {
	return pulumix.Output[map[string]*NamespaceList]{
		OutputState: i.ToNamespaceListMapOutputWithContext(ctx).OutputState,
	}
}

type NamespaceListOutput struct{ *pulumi.OutputState }

func (NamespaceListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceList)(nil)).Elem()
}

func (o NamespaceListOutput) ToNamespaceListOutput() NamespaceListOutput {
	return o
}

func (o NamespaceListOutput) ToNamespaceListOutputWithContext(ctx context.Context) NamespaceListOutput {
	return o
}

func (o NamespaceListOutput) ToOutput(ctx context.Context) pulumix.Output[*NamespaceList] {
	return pulumix.Output[*NamespaceList]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o NamespaceListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
func (o NamespaceListOutput) Items() NamespaceTypeArrayOutput {
	return o.ApplyT(func(v *NamespaceList) NamespaceTypeArrayOutput { return v.Items }).(NamespaceTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o NamespaceListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o NamespaceListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *NamespaceList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type NamespaceListArrayOutput struct{ *pulumi.OutputState }

func (NamespaceListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamespaceList)(nil)).Elem()
}

func (o NamespaceListArrayOutput) ToNamespaceListArrayOutput() NamespaceListArrayOutput {
	return o
}

func (o NamespaceListArrayOutput) ToNamespaceListArrayOutputWithContext(ctx context.Context) NamespaceListArrayOutput {
	return o
}

func (o NamespaceListArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*NamespaceList] {
	return pulumix.Output[[]*NamespaceList]{
		OutputState: o.OutputState,
	}
}

func (o NamespaceListArrayOutput) Index(i pulumi.IntInput) NamespaceListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NamespaceList {
		return vs[0].([]*NamespaceList)[vs[1].(int)]
	}).(NamespaceListOutput)
}

type NamespaceListMapOutput struct{ *pulumi.OutputState }

func (NamespaceListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamespaceList)(nil)).Elem()
}

func (o NamespaceListMapOutput) ToNamespaceListMapOutput() NamespaceListMapOutput {
	return o
}

func (o NamespaceListMapOutput) ToNamespaceListMapOutputWithContext(ctx context.Context) NamespaceListMapOutput {
	return o
}

func (o NamespaceListMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*NamespaceList] {
	return pulumix.Output[map[string]*NamespaceList]{
		OutputState: o.OutputState,
	}
}

func (o NamespaceListMapOutput) MapIndex(k pulumi.StringInput) NamespaceListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NamespaceList {
		return vs[0].(map[string]*NamespaceList)[vs[1].(string)]
	}).(NamespaceListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceListInput)(nil)).Elem(), &NamespaceList{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceListArrayInput)(nil)).Elem(), NamespaceListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceListMapInput)(nil)).Elem(), NamespaceListMap{})
	pulumi.RegisterOutputType(NamespaceListOutput{})
	pulumi.RegisterOutputType(NamespaceListArrayOutput{})
	pulumi.RegisterOutputType(NamespaceListMapOutput{})
}
