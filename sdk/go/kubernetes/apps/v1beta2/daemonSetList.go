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

// DaemonSetList is a collection of daemon sets.
type DaemonSetList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// A list of daemon sets.
	Items DaemonSetTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewDaemonSetList registers a new resource with the given unique name, arguments, and options.
func NewDaemonSetList(ctx *pulumi.Context,
	name string, args *DaemonSetListArgs, opts ...pulumi.ResourceOption) (*DaemonSetList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("apps/v1beta2")
	args.Kind = pulumi.StringPtr("DaemonSetList")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DaemonSetList
	err := ctx.RegisterResource("kubernetes:apps/v1beta2:DaemonSetList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDaemonSetList gets an existing DaemonSetList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDaemonSetList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DaemonSetListState, opts ...pulumi.ResourceOption) (*DaemonSetList, error) {
	var resource DaemonSetList
	err := ctx.ReadResource("kubernetes:apps/v1beta2:DaemonSetList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DaemonSetList resources.
type daemonSetListState struct {
}

type DaemonSetListState struct {
}

func (DaemonSetListState) ElementType() reflect.Type {
	return reflect.TypeOf((*daemonSetListState)(nil)).Elem()
}

type daemonSetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// A list of daemon sets.
	Items []DaemonSetType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a DaemonSetList resource.
type DaemonSetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// A list of daemon sets.
	Items DaemonSetTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (DaemonSetListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*daemonSetListArgs)(nil)).Elem()
}

type DaemonSetListInput interface {
	pulumi.Input

	ToDaemonSetListOutput() DaemonSetListOutput
	ToDaemonSetListOutputWithContext(ctx context.Context) DaemonSetListOutput
}

func (*DaemonSetList) ElementType() reflect.Type {
	return reflect.TypeOf((**DaemonSetList)(nil)).Elem()
}

func (i *DaemonSetList) ToDaemonSetListOutput() DaemonSetListOutput {
	return i.ToDaemonSetListOutputWithContext(context.Background())
}

func (i *DaemonSetList) ToDaemonSetListOutputWithContext(ctx context.Context) DaemonSetListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaemonSetListOutput)
}

func (i *DaemonSetList) ToOutput(ctx context.Context) pulumix.Output[*DaemonSetList] {
	return pulumix.Output[*DaemonSetList]{
		OutputState: i.ToDaemonSetListOutputWithContext(ctx).OutputState,
	}
}

// DaemonSetListArrayInput is an input type that accepts DaemonSetListArray and DaemonSetListArrayOutput values.
// You can construct a concrete instance of `DaemonSetListArrayInput` via:
//
//	DaemonSetListArray{ DaemonSetListArgs{...} }
type DaemonSetListArrayInput interface {
	pulumi.Input

	ToDaemonSetListArrayOutput() DaemonSetListArrayOutput
	ToDaemonSetListArrayOutputWithContext(context.Context) DaemonSetListArrayOutput
}

type DaemonSetListArray []DaemonSetListInput

func (DaemonSetListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DaemonSetList)(nil)).Elem()
}

func (i DaemonSetListArray) ToDaemonSetListArrayOutput() DaemonSetListArrayOutput {
	return i.ToDaemonSetListArrayOutputWithContext(context.Background())
}

func (i DaemonSetListArray) ToDaemonSetListArrayOutputWithContext(ctx context.Context) DaemonSetListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaemonSetListArrayOutput)
}

func (i DaemonSetListArray) ToOutput(ctx context.Context) pulumix.Output[[]*DaemonSetList] {
	return pulumix.Output[[]*DaemonSetList]{
		OutputState: i.ToDaemonSetListArrayOutputWithContext(ctx).OutputState,
	}
}

// DaemonSetListMapInput is an input type that accepts DaemonSetListMap and DaemonSetListMapOutput values.
// You can construct a concrete instance of `DaemonSetListMapInput` via:
//
//	DaemonSetListMap{ "key": DaemonSetListArgs{...} }
type DaemonSetListMapInput interface {
	pulumi.Input

	ToDaemonSetListMapOutput() DaemonSetListMapOutput
	ToDaemonSetListMapOutputWithContext(context.Context) DaemonSetListMapOutput
}

type DaemonSetListMap map[string]DaemonSetListInput

func (DaemonSetListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DaemonSetList)(nil)).Elem()
}

func (i DaemonSetListMap) ToDaemonSetListMapOutput() DaemonSetListMapOutput {
	return i.ToDaemonSetListMapOutputWithContext(context.Background())
}

func (i DaemonSetListMap) ToDaemonSetListMapOutputWithContext(ctx context.Context) DaemonSetListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaemonSetListMapOutput)
}

func (i DaemonSetListMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DaemonSetList] {
	return pulumix.Output[map[string]*DaemonSetList]{
		OutputState: i.ToDaemonSetListMapOutputWithContext(ctx).OutputState,
	}
}

type DaemonSetListOutput struct{ *pulumi.OutputState }

func (DaemonSetListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DaemonSetList)(nil)).Elem()
}

func (o DaemonSetListOutput) ToDaemonSetListOutput() DaemonSetListOutput {
	return o
}

func (o DaemonSetListOutput) ToDaemonSetListOutputWithContext(ctx context.Context) DaemonSetListOutput {
	return o
}

func (o DaemonSetListOutput) ToOutput(ctx context.Context) pulumix.Output[*DaemonSetList] {
	return pulumix.Output[*DaemonSetList]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o DaemonSetListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DaemonSetList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// A list of daemon sets.
func (o DaemonSetListOutput) Items() DaemonSetTypeArrayOutput {
	return o.ApplyT(func(v *DaemonSetList) DaemonSetTypeArrayOutput { return v.Items }).(DaemonSetTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o DaemonSetListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DaemonSetList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o DaemonSetListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *DaemonSetList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type DaemonSetListArrayOutput struct{ *pulumi.OutputState }

func (DaemonSetListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DaemonSetList)(nil)).Elem()
}

func (o DaemonSetListArrayOutput) ToDaemonSetListArrayOutput() DaemonSetListArrayOutput {
	return o
}

func (o DaemonSetListArrayOutput) ToDaemonSetListArrayOutputWithContext(ctx context.Context) DaemonSetListArrayOutput {
	return o
}

func (o DaemonSetListArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DaemonSetList] {
	return pulumix.Output[[]*DaemonSetList]{
		OutputState: o.OutputState,
	}
}

func (o DaemonSetListArrayOutput) Index(i pulumi.IntInput) DaemonSetListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DaemonSetList {
		return vs[0].([]*DaemonSetList)[vs[1].(int)]
	}).(DaemonSetListOutput)
}

type DaemonSetListMapOutput struct{ *pulumi.OutputState }

func (DaemonSetListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DaemonSetList)(nil)).Elem()
}

func (o DaemonSetListMapOutput) ToDaemonSetListMapOutput() DaemonSetListMapOutput {
	return o
}

func (o DaemonSetListMapOutput) ToDaemonSetListMapOutputWithContext(ctx context.Context) DaemonSetListMapOutput {
	return o
}

func (o DaemonSetListMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DaemonSetList] {
	return pulumix.Output[map[string]*DaemonSetList]{
		OutputState: o.OutputState,
	}
}

func (o DaemonSetListMapOutput) MapIndex(k pulumi.StringInput) DaemonSetListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DaemonSetList {
		return vs[0].(map[string]*DaemonSetList)[vs[1].(string)]
	}).(DaemonSetListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DaemonSetListInput)(nil)).Elem(), &DaemonSetList{})
	pulumi.RegisterInputType(reflect.TypeOf((*DaemonSetListArrayInput)(nil)).Elem(), DaemonSetListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DaemonSetListMapInput)(nil)).Elem(), DaemonSetListMap{})
	pulumi.RegisterOutputType(DaemonSetListOutput{})
	pulumi.RegisterOutputType(DaemonSetListArrayOutput{})
	pulumi.RegisterOutputType(DaemonSetListMapOutput{})
}
