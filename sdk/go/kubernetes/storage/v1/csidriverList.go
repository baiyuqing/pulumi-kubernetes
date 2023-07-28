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

// CSIDriverList is a collection of CSIDriver objects.
type CSIDriverList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// items is the list of CSIDriver
	Items CSIDriverTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewCSIDriverList registers a new resource with the given unique name, arguments, and options.
func NewCSIDriverList(ctx *pulumi.Context,
	name string, args *CSIDriverListArgs, opts ...pulumi.ResourceOption) (*CSIDriverList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CSIDriverList")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CSIDriverList
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1:CSIDriverList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSIDriverList gets an existing CSIDriverList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSIDriverList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSIDriverListState, opts ...pulumi.ResourceOption) (*CSIDriverList, error) {
	var resource CSIDriverList
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1:CSIDriverList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSIDriverList resources.
type csidriverListState struct {
}

type CSIDriverListState struct {
}

func (CSIDriverListState) ElementType() reflect.Type {
	return reflect.TypeOf((*csidriverListState)(nil)).Elem()
}

type csidriverListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of CSIDriver
	Items []CSIDriverType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a CSIDriverList resource.
type CSIDriverListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of CSIDriver
	Items CSIDriverTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (CSIDriverListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csidriverListArgs)(nil)).Elem()
}

type CSIDriverListInput interface {
	pulumi.Input

	ToCSIDriverListOutput() CSIDriverListOutput
	ToCSIDriverListOutputWithContext(ctx context.Context) CSIDriverListOutput
}

func (*CSIDriverList) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIDriverList)(nil)).Elem()
}

func (i *CSIDriverList) ToCSIDriverListOutput() CSIDriverListOutput {
	return i.ToCSIDriverListOutputWithContext(context.Background())
}

func (i *CSIDriverList) ToCSIDriverListOutputWithContext(ctx context.Context) CSIDriverListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIDriverListOutput)
}

// CSIDriverListArrayInput is an input type that accepts CSIDriverListArray and CSIDriverListArrayOutput values.
// You can construct a concrete instance of `CSIDriverListArrayInput` via:
//
//	CSIDriverListArray{ CSIDriverListArgs{...} }
type CSIDriverListArrayInput interface {
	pulumi.Input

	ToCSIDriverListArrayOutput() CSIDriverListArrayOutput
	ToCSIDriverListArrayOutputWithContext(context.Context) CSIDriverListArrayOutput
}

type CSIDriverListArray []CSIDriverListInput

func (CSIDriverListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIDriverList)(nil)).Elem()
}

func (i CSIDriverListArray) ToCSIDriverListArrayOutput() CSIDriverListArrayOutput {
	return i.ToCSIDriverListArrayOutputWithContext(context.Background())
}

func (i CSIDriverListArray) ToCSIDriverListArrayOutputWithContext(ctx context.Context) CSIDriverListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIDriverListArrayOutput)
}

// CSIDriverListMapInput is an input type that accepts CSIDriverListMap and CSIDriverListMapOutput values.
// You can construct a concrete instance of `CSIDriverListMapInput` via:
//
//	CSIDriverListMap{ "key": CSIDriverListArgs{...} }
type CSIDriverListMapInput interface {
	pulumi.Input

	ToCSIDriverListMapOutput() CSIDriverListMapOutput
	ToCSIDriverListMapOutputWithContext(context.Context) CSIDriverListMapOutput
}

type CSIDriverListMap map[string]CSIDriverListInput

func (CSIDriverListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIDriverList)(nil)).Elem()
}

func (i CSIDriverListMap) ToCSIDriverListMapOutput() CSIDriverListMapOutput {
	return i.ToCSIDriverListMapOutputWithContext(context.Background())
}

func (i CSIDriverListMap) ToCSIDriverListMapOutputWithContext(ctx context.Context) CSIDriverListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIDriverListMapOutput)
}

type CSIDriverListOutput struct{ *pulumi.OutputState }

func (CSIDriverListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIDriverList)(nil)).Elem()
}

func (o CSIDriverListOutput) ToCSIDriverListOutput() CSIDriverListOutput {
	return o
}

func (o CSIDriverListOutput) ToCSIDriverListOutputWithContext(ctx context.Context) CSIDriverListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CSIDriverListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CSIDriverList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// items is the list of CSIDriver
func (o CSIDriverListOutput) Items() CSIDriverTypeArrayOutput {
	return o.ApplyT(func(v *CSIDriverList) CSIDriverTypeArrayOutput { return v.Items }).(CSIDriverTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CSIDriverListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CSIDriverList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o CSIDriverListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *CSIDriverList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type CSIDriverListArrayOutput struct{ *pulumi.OutputState }

func (CSIDriverListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIDriverList)(nil)).Elem()
}

func (o CSIDriverListArrayOutput) ToCSIDriverListArrayOutput() CSIDriverListArrayOutput {
	return o
}

func (o CSIDriverListArrayOutput) ToCSIDriverListArrayOutputWithContext(ctx context.Context) CSIDriverListArrayOutput {
	return o
}

func (o CSIDriverListArrayOutput) Index(i pulumi.IntInput) CSIDriverListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CSIDriverList {
		return vs[0].([]*CSIDriverList)[vs[1].(int)]
	}).(CSIDriverListOutput)
}

type CSIDriverListMapOutput struct{ *pulumi.OutputState }

func (CSIDriverListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIDriverList)(nil)).Elem()
}

func (o CSIDriverListMapOutput) ToCSIDriverListMapOutput() CSIDriverListMapOutput {
	return o
}

func (o CSIDriverListMapOutput) ToCSIDriverListMapOutputWithContext(ctx context.Context) CSIDriverListMapOutput {
	return o
}

func (o CSIDriverListMapOutput) MapIndex(k pulumi.StringInput) CSIDriverListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CSIDriverList {
		return vs[0].(map[string]*CSIDriverList)[vs[1].(string)]
	}).(CSIDriverListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CSIDriverListInput)(nil)).Elem(), &CSIDriverList{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIDriverListArrayInput)(nil)).Elem(), CSIDriverListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIDriverListMapInput)(nil)).Elem(), CSIDriverListMap{})
	pulumi.RegisterOutputType(CSIDriverListOutput{})
	pulumi.RegisterOutputType(CSIDriverListArrayOutput{})
	pulumi.RegisterOutputType(CSIDriverListMapOutput{})
}
