// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/internal"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// FlowSchemaList is a list of FlowSchema objects.
type FlowSchemaList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// `items` is a list of FlowSchemas.
	Items FlowSchemaTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// `metadata` is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewFlowSchemaList registers a new resource with the given unique name, arguments, and options.
func NewFlowSchemaList(ctx *pulumi.Context,
	name string, args *FlowSchemaListArgs, opts ...pulumi.ResourceOption) (*FlowSchemaList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("flowcontrol.apiserver.k8s.io/v1beta2")
	args.Kind = pulumi.StringPtr("FlowSchemaList")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FlowSchemaList
	err := ctx.RegisterResource("kubernetes:flowcontrol.apiserver.k8s.io/v1beta2:FlowSchemaList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowSchemaList gets an existing FlowSchemaList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowSchemaList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowSchemaListState, opts ...pulumi.ResourceOption) (*FlowSchemaList, error) {
	var resource FlowSchemaList
	err := ctx.ReadResource("kubernetes:flowcontrol.apiserver.k8s.io/v1beta2:FlowSchemaList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowSchemaList resources.
type flowSchemaListState struct {
}

type FlowSchemaListState struct {
}

func (FlowSchemaListState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowSchemaListState)(nil)).Elem()
}

type flowSchemaListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// `items` is a list of FlowSchemas.
	Items []FlowSchemaType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// `metadata` is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a FlowSchemaList resource.
type FlowSchemaListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// `items` is a list of FlowSchemas.
	Items FlowSchemaTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// `metadata` is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (FlowSchemaListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowSchemaListArgs)(nil)).Elem()
}

type FlowSchemaListInput interface {
	pulumi.Input

	ToFlowSchemaListOutput() FlowSchemaListOutput
	ToFlowSchemaListOutputWithContext(ctx context.Context) FlowSchemaListOutput
}

func (*FlowSchemaList) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowSchemaList)(nil)).Elem()
}

func (i *FlowSchemaList) ToFlowSchemaListOutput() FlowSchemaListOutput {
	return i.ToFlowSchemaListOutputWithContext(context.Background())
}

func (i *FlowSchemaList) ToFlowSchemaListOutputWithContext(ctx context.Context) FlowSchemaListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowSchemaListOutput)
}

func (i *FlowSchemaList) ToOutput(ctx context.Context) pulumix.Output[*FlowSchemaList] {
	return pulumix.Output[*FlowSchemaList]{
		OutputState: i.ToFlowSchemaListOutputWithContext(ctx).OutputState,
	}
}

// FlowSchemaListArrayInput is an input type that accepts FlowSchemaListArray and FlowSchemaListArrayOutput values.
// You can construct a concrete instance of `FlowSchemaListArrayInput` via:
//
//	FlowSchemaListArray{ FlowSchemaListArgs{...} }
type FlowSchemaListArrayInput interface {
	pulumi.Input

	ToFlowSchemaListArrayOutput() FlowSchemaListArrayOutput
	ToFlowSchemaListArrayOutputWithContext(context.Context) FlowSchemaListArrayOutput
}

type FlowSchemaListArray []FlowSchemaListInput

func (FlowSchemaListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowSchemaList)(nil)).Elem()
}

func (i FlowSchemaListArray) ToFlowSchemaListArrayOutput() FlowSchemaListArrayOutput {
	return i.ToFlowSchemaListArrayOutputWithContext(context.Background())
}

func (i FlowSchemaListArray) ToFlowSchemaListArrayOutputWithContext(ctx context.Context) FlowSchemaListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowSchemaListArrayOutput)
}

func (i FlowSchemaListArray) ToOutput(ctx context.Context) pulumix.Output[[]*FlowSchemaList] {
	return pulumix.Output[[]*FlowSchemaList]{
		OutputState: i.ToFlowSchemaListArrayOutputWithContext(ctx).OutputState,
	}
}

// FlowSchemaListMapInput is an input type that accepts FlowSchemaListMap and FlowSchemaListMapOutput values.
// You can construct a concrete instance of `FlowSchemaListMapInput` via:
//
//	FlowSchemaListMap{ "key": FlowSchemaListArgs{...} }
type FlowSchemaListMapInput interface {
	pulumi.Input

	ToFlowSchemaListMapOutput() FlowSchemaListMapOutput
	ToFlowSchemaListMapOutputWithContext(context.Context) FlowSchemaListMapOutput
}

type FlowSchemaListMap map[string]FlowSchemaListInput

func (FlowSchemaListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowSchemaList)(nil)).Elem()
}

func (i FlowSchemaListMap) ToFlowSchemaListMapOutput() FlowSchemaListMapOutput {
	return i.ToFlowSchemaListMapOutputWithContext(context.Background())
}

func (i FlowSchemaListMap) ToFlowSchemaListMapOutputWithContext(ctx context.Context) FlowSchemaListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowSchemaListMapOutput)
}

func (i FlowSchemaListMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FlowSchemaList] {
	return pulumix.Output[map[string]*FlowSchemaList]{
		OutputState: i.ToFlowSchemaListMapOutputWithContext(ctx).OutputState,
	}
}

type FlowSchemaListOutput struct{ *pulumi.OutputState }

func (FlowSchemaListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowSchemaList)(nil)).Elem()
}

func (o FlowSchemaListOutput) ToFlowSchemaListOutput() FlowSchemaListOutput {
	return o
}

func (o FlowSchemaListOutput) ToFlowSchemaListOutputWithContext(ctx context.Context) FlowSchemaListOutput {
	return o
}

func (o FlowSchemaListOutput) ToOutput(ctx context.Context) pulumix.Output[*FlowSchemaList] {
	return pulumix.Output[*FlowSchemaList]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o FlowSchemaListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowSchemaList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// `items` is a list of FlowSchemas.
func (o FlowSchemaListOutput) Items() FlowSchemaTypeArrayOutput {
	return o.ApplyT(func(v *FlowSchemaList) FlowSchemaTypeArrayOutput { return v.Items }).(FlowSchemaTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o FlowSchemaListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowSchemaList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// `metadata` is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o FlowSchemaListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *FlowSchemaList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type FlowSchemaListArrayOutput struct{ *pulumi.OutputState }

func (FlowSchemaListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowSchemaList)(nil)).Elem()
}

func (o FlowSchemaListArrayOutput) ToFlowSchemaListArrayOutput() FlowSchemaListArrayOutput {
	return o
}

func (o FlowSchemaListArrayOutput) ToFlowSchemaListArrayOutputWithContext(ctx context.Context) FlowSchemaListArrayOutput {
	return o
}

func (o FlowSchemaListArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FlowSchemaList] {
	return pulumix.Output[[]*FlowSchemaList]{
		OutputState: o.OutputState,
	}
}

func (o FlowSchemaListArrayOutput) Index(i pulumi.IntInput) FlowSchemaListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlowSchemaList {
		return vs[0].([]*FlowSchemaList)[vs[1].(int)]
	}).(FlowSchemaListOutput)
}

type FlowSchemaListMapOutput struct{ *pulumi.OutputState }

func (FlowSchemaListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowSchemaList)(nil)).Elem()
}

func (o FlowSchemaListMapOutput) ToFlowSchemaListMapOutput() FlowSchemaListMapOutput {
	return o
}

func (o FlowSchemaListMapOutput) ToFlowSchemaListMapOutputWithContext(ctx context.Context) FlowSchemaListMapOutput {
	return o
}

func (o FlowSchemaListMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FlowSchemaList] {
	return pulumix.Output[map[string]*FlowSchemaList]{
		OutputState: o.OutputState,
	}
}

func (o FlowSchemaListMapOutput) MapIndex(k pulumi.StringInput) FlowSchemaListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlowSchemaList {
		return vs[0].(map[string]*FlowSchemaList)[vs[1].(string)]
	}).(FlowSchemaListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowSchemaListInput)(nil)).Elem(), &FlowSchemaList{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowSchemaListArrayInput)(nil)).Elem(), FlowSchemaListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowSchemaListMapInput)(nil)).Elem(), FlowSchemaListMap{})
	pulumi.RegisterOutputType(FlowSchemaListOutput{})
	pulumi.RegisterOutputType(FlowSchemaListArrayOutput{})
	pulumi.RegisterOutputType(FlowSchemaListMapOutput{})
}
