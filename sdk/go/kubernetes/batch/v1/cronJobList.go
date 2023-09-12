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

// CronJobList is a collection of cron jobs.
type CronJobList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// items is the list of CronJobs.
	Items CronJobTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewCronJobList registers a new resource with the given unique name, arguments, and options.
func NewCronJobList(ctx *pulumi.Context,
	name string, args *CronJobListArgs, opts ...pulumi.ResourceOption) (*CronJobList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("batch/v1")
	args.Kind = pulumi.StringPtr("CronJobList")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CronJobList
	err := ctx.RegisterResource("kubernetes:batch/v1:CronJobList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCronJobList gets an existing CronJobList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCronJobList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CronJobListState, opts ...pulumi.ResourceOption) (*CronJobList, error) {
	var resource CronJobList
	err := ctx.ReadResource("kubernetes:batch/v1:CronJobList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CronJobList resources.
type cronJobListState struct {
}

type CronJobListState struct {
}

func (CronJobListState) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobListState)(nil)).Elem()
}

type cronJobListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of CronJobs.
	Items []CronJobType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a CronJobList resource.
type CronJobListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of CronJobs.
	Items CronJobTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (CronJobListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobListArgs)(nil)).Elem()
}

type CronJobListInput interface {
	pulumi.Input

	ToCronJobListOutput() CronJobListOutput
	ToCronJobListOutputWithContext(ctx context.Context) CronJobListOutput
}

func (*CronJobList) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJobList)(nil)).Elem()
}

func (i *CronJobList) ToCronJobListOutput() CronJobListOutput {
	return i.ToCronJobListOutputWithContext(context.Background())
}

func (i *CronJobList) ToCronJobListOutputWithContext(ctx context.Context) CronJobListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobListOutput)
}

func (i *CronJobList) ToOutput(ctx context.Context) pulumix.Output[*CronJobList] {
	return pulumix.Output[*CronJobList]{
		OutputState: i.ToCronJobListOutputWithContext(ctx).OutputState,
	}
}

// CronJobListArrayInput is an input type that accepts CronJobListArray and CronJobListArrayOutput values.
// You can construct a concrete instance of `CronJobListArrayInput` via:
//
//	CronJobListArray{ CronJobListArgs{...} }
type CronJobListArrayInput interface {
	pulumi.Input

	ToCronJobListArrayOutput() CronJobListArrayOutput
	ToCronJobListArrayOutputWithContext(context.Context) CronJobListArrayOutput
}

type CronJobListArray []CronJobListInput

func (CronJobListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CronJobList)(nil)).Elem()
}

func (i CronJobListArray) ToCronJobListArrayOutput() CronJobListArrayOutput {
	return i.ToCronJobListArrayOutputWithContext(context.Background())
}

func (i CronJobListArray) ToCronJobListArrayOutputWithContext(ctx context.Context) CronJobListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobListArrayOutput)
}

func (i CronJobListArray) ToOutput(ctx context.Context) pulumix.Output[[]*CronJobList] {
	return pulumix.Output[[]*CronJobList]{
		OutputState: i.ToCronJobListArrayOutputWithContext(ctx).OutputState,
	}
}

// CronJobListMapInput is an input type that accepts CronJobListMap and CronJobListMapOutput values.
// You can construct a concrete instance of `CronJobListMapInput` via:
//
//	CronJobListMap{ "key": CronJobListArgs{...} }
type CronJobListMapInput interface {
	pulumi.Input

	ToCronJobListMapOutput() CronJobListMapOutput
	ToCronJobListMapOutputWithContext(context.Context) CronJobListMapOutput
}

type CronJobListMap map[string]CronJobListInput

func (CronJobListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CronJobList)(nil)).Elem()
}

func (i CronJobListMap) ToCronJobListMapOutput() CronJobListMapOutput {
	return i.ToCronJobListMapOutputWithContext(context.Background())
}

func (i CronJobListMap) ToCronJobListMapOutputWithContext(ctx context.Context) CronJobListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobListMapOutput)
}

func (i CronJobListMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CronJobList] {
	return pulumix.Output[map[string]*CronJobList]{
		OutputState: i.ToCronJobListMapOutputWithContext(ctx).OutputState,
	}
}

type CronJobListOutput struct{ *pulumi.OutputState }

func (CronJobListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJobList)(nil)).Elem()
}

func (o CronJobListOutput) ToCronJobListOutput() CronJobListOutput {
	return o
}

func (o CronJobListOutput) ToCronJobListOutputWithContext(ctx context.Context) CronJobListOutput {
	return o
}

func (o CronJobListOutput) ToOutput(ctx context.Context) pulumix.Output[*CronJobList] {
	return pulumix.Output[*CronJobList]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CronJobListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJobList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// items is the list of CronJobs.
func (o CronJobListOutput) Items() CronJobTypeArrayOutput {
	return o.ApplyT(func(v *CronJobList) CronJobTypeArrayOutput { return v.Items }).(CronJobTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CronJobListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJobList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o CronJobListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *CronJobList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type CronJobListArrayOutput struct{ *pulumi.OutputState }

func (CronJobListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CronJobList)(nil)).Elem()
}

func (o CronJobListArrayOutput) ToCronJobListArrayOutput() CronJobListArrayOutput {
	return o
}

func (o CronJobListArrayOutput) ToCronJobListArrayOutputWithContext(ctx context.Context) CronJobListArrayOutput {
	return o
}

func (o CronJobListArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CronJobList] {
	return pulumix.Output[[]*CronJobList]{
		OutputState: o.OutputState,
	}
}

func (o CronJobListArrayOutput) Index(i pulumi.IntInput) CronJobListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CronJobList {
		return vs[0].([]*CronJobList)[vs[1].(int)]
	}).(CronJobListOutput)
}

type CronJobListMapOutput struct{ *pulumi.OutputState }

func (CronJobListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CronJobList)(nil)).Elem()
}

func (o CronJobListMapOutput) ToCronJobListMapOutput() CronJobListMapOutput {
	return o
}

func (o CronJobListMapOutput) ToCronJobListMapOutputWithContext(ctx context.Context) CronJobListMapOutput {
	return o
}

func (o CronJobListMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CronJobList] {
	return pulumix.Output[map[string]*CronJobList]{
		OutputState: o.OutputState,
	}
}

func (o CronJobListMapOutput) MapIndex(k pulumi.StringInput) CronJobListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CronJobList {
		return vs[0].(map[string]*CronJobList)[vs[1].(string)]
	}).(CronJobListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CronJobListInput)(nil)).Elem(), &CronJobList{})
	pulumi.RegisterInputType(reflect.TypeOf((*CronJobListArrayInput)(nil)).Elem(), CronJobListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CronJobListMapInput)(nil)).Elem(), CronJobListMap{})
	pulumi.RegisterOutputType(CronJobListOutput{})
	pulumi.RegisterOutputType(CronJobListArrayOutput{})
	pulumi.RegisterOutputType(CronJobListMapOutput{})
}
