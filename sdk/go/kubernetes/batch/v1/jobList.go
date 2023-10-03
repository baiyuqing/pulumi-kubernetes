// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// JobList is a collection of jobs.
type JobList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// items is the list of Jobs.
	Items JobTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewJobList registers a new resource with the given unique name, arguments, and options.
func NewJobList(ctx *pulumi.Context,
	name string, args *JobListArgs, opts ...pulumi.ResourceOption) (*JobList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("batch/v1")
	args.Kind = pulumi.StringPtr("JobList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource JobList
	err := ctx.RegisterResource("kubernetes:batch/v1:JobList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobList gets an existing JobList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobListState, opts ...pulumi.ResourceOption) (*JobList, error) {
	var resource JobList
	err := ctx.ReadResource("kubernetes:batch/v1:JobList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobList resources.
type jobListState struct {
}

type JobListState struct {
}

func (JobListState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobListState)(nil)).Elem()
}

type jobListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of Jobs.
	Items []JobType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a JobList resource.
type JobListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of Jobs.
	Items JobTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (JobListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobListArgs)(nil)).Elem()
}

type JobListInput interface {
	pulumi.Input

	ToJobListOutput() JobListOutput
	ToJobListOutputWithContext(ctx context.Context) JobListOutput
}

func (*JobList) ElementType() reflect.Type {
	return reflect.TypeOf((**JobList)(nil)).Elem()
}

func (i *JobList) ToJobListOutput() JobListOutput {
	return i.ToJobListOutputWithContext(context.Background())
}

func (i *JobList) ToJobListOutputWithContext(ctx context.Context) JobListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobListOutput)
}

func (i *JobList) ToOutput(ctx context.Context) pulumix.Output[*JobList] {
	return pulumix.Output[*JobList]{
		OutputState: i.ToJobListOutputWithContext(ctx).OutputState,
	}
}

// JobListArrayInput is an input type that accepts JobListArray and JobListArrayOutput values.
// You can construct a concrete instance of `JobListArrayInput` via:
//
//	JobListArray{ JobListArgs{...} }
type JobListArrayInput interface {
	pulumi.Input

	ToJobListArrayOutput() JobListArrayOutput
	ToJobListArrayOutputWithContext(context.Context) JobListArrayOutput
}

type JobListArray []JobListInput

func (JobListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JobList)(nil)).Elem()
}

func (i JobListArray) ToJobListArrayOutput() JobListArrayOutput {
	return i.ToJobListArrayOutputWithContext(context.Background())
}

func (i JobListArray) ToJobListArrayOutputWithContext(ctx context.Context) JobListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobListArrayOutput)
}

func (i JobListArray) ToOutput(ctx context.Context) pulumix.Output[[]*JobList] {
	return pulumix.Output[[]*JobList]{
		OutputState: i.ToJobListArrayOutputWithContext(ctx).OutputState,
	}
}

// JobListMapInput is an input type that accepts JobListMap and JobListMapOutput values.
// You can construct a concrete instance of `JobListMapInput` via:
//
//	JobListMap{ "key": JobListArgs{...} }
type JobListMapInput interface {
	pulumi.Input

	ToJobListMapOutput() JobListMapOutput
	ToJobListMapOutputWithContext(context.Context) JobListMapOutput
}

type JobListMap map[string]JobListInput

func (JobListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JobList)(nil)).Elem()
}

func (i JobListMap) ToJobListMapOutput() JobListMapOutput {
	return i.ToJobListMapOutputWithContext(context.Background())
}

func (i JobListMap) ToJobListMapOutputWithContext(ctx context.Context) JobListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobListMapOutput)
}

func (i JobListMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*JobList] {
	return pulumix.Output[map[string]*JobList]{
		OutputState: i.ToJobListMapOutputWithContext(ctx).OutputState,
	}
}

type JobListOutput struct{ *pulumi.OutputState }

func (JobListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobList)(nil)).Elem()
}

func (o JobListOutput) ToJobListOutput() JobListOutput {
	return o
}

func (o JobListOutput) ToJobListOutputWithContext(ctx context.Context) JobListOutput {
	return o
}

func (o JobListOutput) ToOutput(ctx context.Context) pulumix.Output[*JobList] {
	return pulumix.Output[*JobList]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o JobListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *JobList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// items is the list of Jobs.
func (o JobListOutput) Items() JobTypeArrayOutput {
	return o.ApplyT(func(v *JobList) JobTypeArrayOutput { return v.Items }).(JobTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o JobListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *JobList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o JobListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *JobList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type JobListArrayOutput struct{ *pulumi.OutputState }

func (JobListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JobList)(nil)).Elem()
}

func (o JobListArrayOutput) ToJobListArrayOutput() JobListArrayOutput {
	return o
}

func (o JobListArrayOutput) ToJobListArrayOutputWithContext(ctx context.Context) JobListArrayOutput {
	return o
}

func (o JobListArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*JobList] {
	return pulumix.Output[[]*JobList]{
		OutputState: o.OutputState,
	}
}

func (o JobListArrayOutput) Index(i pulumi.IntInput) JobListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *JobList {
		return vs[0].([]*JobList)[vs[1].(int)]
	}).(JobListOutput)
}

type JobListMapOutput struct{ *pulumi.OutputState }

func (JobListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JobList)(nil)).Elem()
}

func (o JobListMapOutput) ToJobListMapOutput() JobListMapOutput {
	return o
}

func (o JobListMapOutput) ToJobListMapOutputWithContext(ctx context.Context) JobListMapOutput {
	return o
}

func (o JobListMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*JobList] {
	return pulumix.Output[map[string]*JobList]{
		OutputState: o.OutputState,
	}
}

func (o JobListMapOutput) MapIndex(k pulumi.StringInput) JobListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *JobList {
		return vs[0].(map[string]*JobList)[vs[1].(string)]
	}).(JobListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobListInput)(nil)).Elem(), &JobList{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobListArrayInput)(nil)).Elem(), JobListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobListMapInput)(nil)).Elem(), JobListMap{})
	pulumi.RegisterOutputType(JobListOutput{})
	pulumi.RegisterOutputType(JobListArrayOutput{})
	pulumi.RegisterOutputType(JobListMapOutput{})
}
