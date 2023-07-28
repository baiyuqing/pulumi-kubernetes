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

// PersistentVolumeList is a list of PersistentVolume items.
type PersistentVolumeList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// items is a list of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
	Items PersistentVolumeTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewPersistentVolumeList registers a new resource with the given unique name, arguments, and options.
func NewPersistentVolumeList(ctx *pulumi.Context,
	name string, args *PersistentVolumeListArgs, opts ...pulumi.ResourceOption) (*PersistentVolumeList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("PersistentVolumeList")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PersistentVolumeList
	err := ctx.RegisterResource("kubernetes:core/v1:PersistentVolumeList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistentVolumeList gets an existing PersistentVolumeList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistentVolumeList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersistentVolumeListState, opts ...pulumi.ResourceOption) (*PersistentVolumeList, error) {
	var resource PersistentVolumeList
	err := ctx.ReadResource("kubernetes:core/v1:PersistentVolumeList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersistentVolumeList resources.
type persistentVolumeListState struct {
}

type PersistentVolumeListState struct {
}

func (PersistentVolumeListState) ElementType() reflect.Type {
	return reflect.TypeOf((*persistentVolumeListState)(nil)).Elem()
}

type persistentVolumeListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is a list of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
	Items []PersistentVolumeType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a PersistentVolumeList resource.
type PersistentVolumeListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is a list of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
	Items PersistentVolumeTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (PersistentVolumeListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*persistentVolumeListArgs)(nil)).Elem()
}

type PersistentVolumeListInput interface {
	pulumi.Input

	ToPersistentVolumeListOutput() PersistentVolumeListOutput
	ToPersistentVolumeListOutputWithContext(ctx context.Context) PersistentVolumeListOutput
}

func (*PersistentVolumeList) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentVolumeList)(nil)).Elem()
}

func (i *PersistentVolumeList) ToPersistentVolumeListOutput() PersistentVolumeListOutput {
	return i.ToPersistentVolumeListOutputWithContext(context.Background())
}

func (i *PersistentVolumeList) ToPersistentVolumeListOutputWithContext(ctx context.Context) PersistentVolumeListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeListOutput)
}

// PersistentVolumeListArrayInput is an input type that accepts PersistentVolumeListArray and PersistentVolumeListArrayOutput values.
// You can construct a concrete instance of `PersistentVolumeListArrayInput` via:
//
//	PersistentVolumeListArray{ PersistentVolumeListArgs{...} }
type PersistentVolumeListArrayInput interface {
	pulumi.Input

	ToPersistentVolumeListArrayOutput() PersistentVolumeListArrayOutput
	ToPersistentVolumeListArrayOutputWithContext(context.Context) PersistentVolumeListArrayOutput
}

type PersistentVolumeListArray []PersistentVolumeListInput

func (PersistentVolumeListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistentVolumeList)(nil)).Elem()
}

func (i PersistentVolumeListArray) ToPersistentVolumeListArrayOutput() PersistentVolumeListArrayOutput {
	return i.ToPersistentVolumeListArrayOutputWithContext(context.Background())
}

func (i PersistentVolumeListArray) ToPersistentVolumeListArrayOutputWithContext(ctx context.Context) PersistentVolumeListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeListArrayOutput)
}

// PersistentVolumeListMapInput is an input type that accepts PersistentVolumeListMap and PersistentVolumeListMapOutput values.
// You can construct a concrete instance of `PersistentVolumeListMapInput` via:
//
//	PersistentVolumeListMap{ "key": PersistentVolumeListArgs{...} }
type PersistentVolumeListMapInput interface {
	pulumi.Input

	ToPersistentVolumeListMapOutput() PersistentVolumeListMapOutput
	ToPersistentVolumeListMapOutputWithContext(context.Context) PersistentVolumeListMapOutput
}

type PersistentVolumeListMap map[string]PersistentVolumeListInput

func (PersistentVolumeListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistentVolumeList)(nil)).Elem()
}

func (i PersistentVolumeListMap) ToPersistentVolumeListMapOutput() PersistentVolumeListMapOutput {
	return i.ToPersistentVolumeListMapOutputWithContext(context.Background())
}

func (i PersistentVolumeListMap) ToPersistentVolumeListMapOutputWithContext(ctx context.Context) PersistentVolumeListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeListMapOutput)
}

type PersistentVolumeListOutput struct{ *pulumi.OutputState }

func (PersistentVolumeListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentVolumeList)(nil)).Elem()
}

func (o PersistentVolumeListOutput) ToPersistentVolumeListOutput() PersistentVolumeListOutput {
	return o
}

func (o PersistentVolumeListOutput) ToPersistentVolumeListOutputWithContext(ctx context.Context) PersistentVolumeListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PersistentVolumeListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistentVolumeList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// items is a list of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
func (o PersistentVolumeListOutput) Items() PersistentVolumeTypeArrayOutput {
	return o.ApplyT(func(v *PersistentVolumeList) PersistentVolumeTypeArrayOutput { return v.Items }).(PersistentVolumeTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PersistentVolumeListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistentVolumeList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PersistentVolumeListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *PersistentVolumeList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type PersistentVolumeListArrayOutput struct{ *pulumi.OutputState }

func (PersistentVolumeListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistentVolumeList)(nil)).Elem()
}

func (o PersistentVolumeListArrayOutput) ToPersistentVolumeListArrayOutput() PersistentVolumeListArrayOutput {
	return o
}

func (o PersistentVolumeListArrayOutput) ToPersistentVolumeListArrayOutputWithContext(ctx context.Context) PersistentVolumeListArrayOutput {
	return o
}

func (o PersistentVolumeListArrayOutput) Index(i pulumi.IntInput) PersistentVolumeListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PersistentVolumeList {
		return vs[0].([]*PersistentVolumeList)[vs[1].(int)]
	}).(PersistentVolumeListOutput)
}

type PersistentVolumeListMapOutput struct{ *pulumi.OutputState }

func (PersistentVolumeListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistentVolumeList)(nil)).Elem()
}

func (o PersistentVolumeListMapOutput) ToPersistentVolumeListMapOutput() PersistentVolumeListMapOutput {
	return o
}

func (o PersistentVolumeListMapOutput) ToPersistentVolumeListMapOutputWithContext(ctx context.Context) PersistentVolumeListMapOutput {
	return o
}

func (o PersistentVolumeListMapOutput) MapIndex(k pulumi.StringInput) PersistentVolumeListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PersistentVolumeList {
		return vs[0].(map[string]*PersistentVolumeList)[vs[1].(string)]
	}).(PersistentVolumeListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PersistentVolumeListInput)(nil)).Elem(), &PersistentVolumeList{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistentVolumeListArrayInput)(nil)).Elem(), PersistentVolumeListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistentVolumeListMapInput)(nil)).Elem(), PersistentVolumeListMap{})
	pulumi.RegisterOutputType(PersistentVolumeListOutput{})
	pulumi.RegisterOutputType(PersistentVolumeListArrayOutput{})
	pulumi.RegisterOutputType(PersistentVolumeListMapOutput{})
}
