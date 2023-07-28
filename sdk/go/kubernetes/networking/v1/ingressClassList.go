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
)

// IngressClassList is a collection of IngressClasses.
type IngressClassList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// items is the list of IngressClasses.
	Items IngressClassTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata.
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewIngressClassList registers a new resource with the given unique name, arguments, and options.
func NewIngressClassList(ctx *pulumi.Context,
	name string, args *IngressClassListArgs, opts ...pulumi.ResourceOption) (*IngressClassList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("networking.k8s.io/v1")
	args.Kind = pulumi.StringPtr("IngressClassList")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IngressClassList
	err := ctx.RegisterResource("kubernetes:networking.k8s.io/v1:IngressClassList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIngressClassList gets an existing IngressClassList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIngressClassList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IngressClassListState, opts ...pulumi.ResourceOption) (*IngressClassList, error) {
	var resource IngressClassList
	err := ctx.ReadResource("kubernetes:networking.k8s.io/v1:IngressClassList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IngressClassList resources.
type ingressClassListState struct {
}

type IngressClassListState struct {
}

func (IngressClassListState) ElementType() reflect.Type {
	return reflect.TypeOf((*ingressClassListState)(nil)).Elem()
}

type ingressClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of IngressClasses.
	Items []IngressClassType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a IngressClassList resource.
type IngressClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of IngressClasses.
	Items IngressClassTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata.
	Metadata metav1.ListMetaPtrInput
}

func (IngressClassListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ingressClassListArgs)(nil)).Elem()
}

type IngressClassListInput interface {
	pulumi.Input

	ToIngressClassListOutput() IngressClassListOutput
	ToIngressClassListOutputWithContext(ctx context.Context) IngressClassListOutput
}

func (*IngressClassList) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressClassList)(nil)).Elem()
}

func (i *IngressClassList) ToIngressClassListOutput() IngressClassListOutput {
	return i.ToIngressClassListOutputWithContext(context.Background())
}

func (i *IngressClassList) ToIngressClassListOutputWithContext(ctx context.Context) IngressClassListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressClassListOutput)
}

// IngressClassListArrayInput is an input type that accepts IngressClassListArray and IngressClassListArrayOutput values.
// You can construct a concrete instance of `IngressClassListArrayInput` via:
//
//	IngressClassListArray{ IngressClassListArgs{...} }
type IngressClassListArrayInput interface {
	pulumi.Input

	ToIngressClassListArrayOutput() IngressClassListArrayOutput
	ToIngressClassListArrayOutputWithContext(context.Context) IngressClassListArrayOutput
}

type IngressClassListArray []IngressClassListInput

func (IngressClassListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IngressClassList)(nil)).Elem()
}

func (i IngressClassListArray) ToIngressClassListArrayOutput() IngressClassListArrayOutput {
	return i.ToIngressClassListArrayOutputWithContext(context.Background())
}

func (i IngressClassListArray) ToIngressClassListArrayOutputWithContext(ctx context.Context) IngressClassListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressClassListArrayOutput)
}

// IngressClassListMapInput is an input type that accepts IngressClassListMap and IngressClassListMapOutput values.
// You can construct a concrete instance of `IngressClassListMapInput` via:
//
//	IngressClassListMap{ "key": IngressClassListArgs{...} }
type IngressClassListMapInput interface {
	pulumi.Input

	ToIngressClassListMapOutput() IngressClassListMapOutput
	ToIngressClassListMapOutputWithContext(context.Context) IngressClassListMapOutput
}

type IngressClassListMap map[string]IngressClassListInput

func (IngressClassListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IngressClassList)(nil)).Elem()
}

func (i IngressClassListMap) ToIngressClassListMapOutput() IngressClassListMapOutput {
	return i.ToIngressClassListMapOutputWithContext(context.Background())
}

func (i IngressClassListMap) ToIngressClassListMapOutputWithContext(ctx context.Context) IngressClassListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressClassListMapOutput)
}

type IngressClassListOutput struct{ *pulumi.OutputState }

func (IngressClassListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressClassList)(nil)).Elem()
}

func (o IngressClassListOutput) ToIngressClassListOutput() IngressClassListOutput {
	return o
}

func (o IngressClassListOutput) ToIngressClassListOutputWithContext(ctx context.Context) IngressClassListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o IngressClassListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *IngressClassList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// items is the list of IngressClasses.
func (o IngressClassListOutput) Items() IngressClassTypeArrayOutput {
	return o.ApplyT(func(v *IngressClassList) IngressClassTypeArrayOutput { return v.Items }).(IngressClassTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o IngressClassListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *IngressClassList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata.
func (o IngressClassListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *IngressClassList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type IngressClassListArrayOutput struct{ *pulumi.OutputState }

func (IngressClassListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IngressClassList)(nil)).Elem()
}

func (o IngressClassListArrayOutput) ToIngressClassListArrayOutput() IngressClassListArrayOutput {
	return o
}

func (o IngressClassListArrayOutput) ToIngressClassListArrayOutputWithContext(ctx context.Context) IngressClassListArrayOutput {
	return o
}

func (o IngressClassListArrayOutput) Index(i pulumi.IntInput) IngressClassListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IngressClassList {
		return vs[0].([]*IngressClassList)[vs[1].(int)]
	}).(IngressClassListOutput)
}

type IngressClassListMapOutput struct{ *pulumi.OutputState }

func (IngressClassListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IngressClassList)(nil)).Elem()
}

func (o IngressClassListMapOutput) ToIngressClassListMapOutput() IngressClassListMapOutput {
	return o
}

func (o IngressClassListMapOutput) ToIngressClassListMapOutputWithContext(ctx context.Context) IngressClassListMapOutput {
	return o
}

func (o IngressClassListMapOutput) MapIndex(k pulumi.StringInput) IngressClassListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IngressClassList {
		return vs[0].(map[string]*IngressClassList)[vs[1].(string)]
	}).(IngressClassListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IngressClassListInput)(nil)).Elem(), &IngressClassList{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressClassListArrayInput)(nil)).Elem(), IngressClassListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressClassListMapInput)(nil)).Elem(), IngressClassListMap{})
	pulumi.RegisterOutputType(IngressClassListOutput{})
	pulumi.RegisterOutputType(IngressClassListArrayOutput{})
	pulumi.RegisterOutputType(IngressClassListMapOutput{})
}
