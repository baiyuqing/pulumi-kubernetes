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

// CSIStorageCapacityList is a collection of CSIStorageCapacity objects.
type CSIStorageCapacityList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// items is the list of CSIStorageCapacity objects.
	Items CSIStorageCapacityTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewCSIStorageCapacityList registers a new resource with the given unique name, arguments, and options.
func NewCSIStorageCapacityList(ctx *pulumi.Context,
	name string, args *CSIStorageCapacityListArgs, opts ...pulumi.ResourceOption) (*CSIStorageCapacityList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CSIStorageCapacityList")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CSIStorageCapacityList
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1:CSIStorageCapacityList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSIStorageCapacityList gets an existing CSIStorageCapacityList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSIStorageCapacityList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSIStorageCapacityListState, opts ...pulumi.ResourceOption) (*CSIStorageCapacityList, error) {
	var resource CSIStorageCapacityList
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1:CSIStorageCapacityList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSIStorageCapacityList resources.
type csistorageCapacityListState struct {
}

type CSIStorageCapacityListState struct {
}

func (CSIStorageCapacityListState) ElementType() reflect.Type {
	return reflect.TypeOf((*csistorageCapacityListState)(nil)).Elem()
}

type csistorageCapacityListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of CSIStorageCapacity objects.
	Items []CSIStorageCapacityType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a CSIStorageCapacityList resource.
type CSIStorageCapacityListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of CSIStorageCapacity objects.
	Items CSIStorageCapacityTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (CSIStorageCapacityListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csistorageCapacityListArgs)(nil)).Elem()
}

type CSIStorageCapacityListInput interface {
	pulumi.Input

	ToCSIStorageCapacityListOutput() CSIStorageCapacityListOutput
	ToCSIStorageCapacityListOutputWithContext(ctx context.Context) CSIStorageCapacityListOutput
}

func (*CSIStorageCapacityList) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIStorageCapacityList)(nil)).Elem()
}

func (i *CSIStorageCapacityList) ToCSIStorageCapacityListOutput() CSIStorageCapacityListOutput {
	return i.ToCSIStorageCapacityListOutputWithContext(context.Background())
}

func (i *CSIStorageCapacityList) ToCSIStorageCapacityListOutputWithContext(ctx context.Context) CSIStorageCapacityListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityListOutput)
}

// CSIStorageCapacityListArrayInput is an input type that accepts CSIStorageCapacityListArray and CSIStorageCapacityListArrayOutput values.
// You can construct a concrete instance of `CSIStorageCapacityListArrayInput` via:
//
//	CSIStorageCapacityListArray{ CSIStorageCapacityListArgs{...} }
type CSIStorageCapacityListArrayInput interface {
	pulumi.Input

	ToCSIStorageCapacityListArrayOutput() CSIStorageCapacityListArrayOutput
	ToCSIStorageCapacityListArrayOutputWithContext(context.Context) CSIStorageCapacityListArrayOutput
}

type CSIStorageCapacityListArray []CSIStorageCapacityListInput

func (CSIStorageCapacityListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIStorageCapacityList)(nil)).Elem()
}

func (i CSIStorageCapacityListArray) ToCSIStorageCapacityListArrayOutput() CSIStorageCapacityListArrayOutput {
	return i.ToCSIStorageCapacityListArrayOutputWithContext(context.Background())
}

func (i CSIStorageCapacityListArray) ToCSIStorageCapacityListArrayOutputWithContext(ctx context.Context) CSIStorageCapacityListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityListArrayOutput)
}

// CSIStorageCapacityListMapInput is an input type that accepts CSIStorageCapacityListMap and CSIStorageCapacityListMapOutput values.
// You can construct a concrete instance of `CSIStorageCapacityListMapInput` via:
//
//	CSIStorageCapacityListMap{ "key": CSIStorageCapacityListArgs{...} }
type CSIStorageCapacityListMapInput interface {
	pulumi.Input

	ToCSIStorageCapacityListMapOutput() CSIStorageCapacityListMapOutput
	ToCSIStorageCapacityListMapOutputWithContext(context.Context) CSIStorageCapacityListMapOutput
}

type CSIStorageCapacityListMap map[string]CSIStorageCapacityListInput

func (CSIStorageCapacityListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIStorageCapacityList)(nil)).Elem()
}

func (i CSIStorageCapacityListMap) ToCSIStorageCapacityListMapOutput() CSIStorageCapacityListMapOutput {
	return i.ToCSIStorageCapacityListMapOutputWithContext(context.Background())
}

func (i CSIStorageCapacityListMap) ToCSIStorageCapacityListMapOutputWithContext(ctx context.Context) CSIStorageCapacityListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityListMapOutput)
}

type CSIStorageCapacityListOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIStorageCapacityList)(nil)).Elem()
}

func (o CSIStorageCapacityListOutput) ToCSIStorageCapacityListOutput() CSIStorageCapacityListOutput {
	return o
}

func (o CSIStorageCapacityListOutput) ToCSIStorageCapacityListOutputWithContext(ctx context.Context) CSIStorageCapacityListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CSIStorageCapacityListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CSIStorageCapacityList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// items is the list of CSIStorageCapacity objects.
func (o CSIStorageCapacityListOutput) Items() CSIStorageCapacityTypeArrayOutput {
	return o.ApplyT(func(v *CSIStorageCapacityList) CSIStorageCapacityTypeArrayOutput { return v.Items }).(CSIStorageCapacityTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CSIStorageCapacityListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CSIStorageCapacityList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o CSIStorageCapacityListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *CSIStorageCapacityList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type CSIStorageCapacityListArrayOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIStorageCapacityList)(nil)).Elem()
}

func (o CSIStorageCapacityListArrayOutput) ToCSIStorageCapacityListArrayOutput() CSIStorageCapacityListArrayOutput {
	return o
}

func (o CSIStorageCapacityListArrayOutput) ToCSIStorageCapacityListArrayOutputWithContext(ctx context.Context) CSIStorageCapacityListArrayOutput {
	return o
}

func (o CSIStorageCapacityListArrayOutput) Index(i pulumi.IntInput) CSIStorageCapacityListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CSIStorageCapacityList {
		return vs[0].([]*CSIStorageCapacityList)[vs[1].(int)]
	}).(CSIStorageCapacityListOutput)
}

type CSIStorageCapacityListMapOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIStorageCapacityList)(nil)).Elem()
}

func (o CSIStorageCapacityListMapOutput) ToCSIStorageCapacityListMapOutput() CSIStorageCapacityListMapOutput {
	return o
}

func (o CSIStorageCapacityListMapOutput) ToCSIStorageCapacityListMapOutputWithContext(ctx context.Context) CSIStorageCapacityListMapOutput {
	return o
}

func (o CSIStorageCapacityListMapOutput) MapIndex(k pulumi.StringInput) CSIStorageCapacityListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CSIStorageCapacityList {
		return vs[0].(map[string]*CSIStorageCapacityList)[vs[1].(string)]
	}).(CSIStorageCapacityListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityListInput)(nil)).Elem(), &CSIStorageCapacityList{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityListArrayInput)(nil)).Elem(), CSIStorageCapacityListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityListMapInput)(nil)).Elem(), CSIStorageCapacityListMap{})
	pulumi.RegisterOutputType(CSIStorageCapacityListOutput{})
	pulumi.RegisterOutputType(CSIStorageCapacityListArrayOutput{})
	pulumi.RegisterOutputType(CSIStorageCapacityListMapOutput{})
}
