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

// ServiceList holds a list of services.
type ServiceList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of services
	Items ServiceTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewServiceList registers a new resource with the given unique name, arguments, and options.
func NewServiceList(ctx *pulumi.Context,
	name string, args *ServiceListArgs, opts ...pulumi.ResourceOption) (*ServiceList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ServiceList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServiceList
	err := ctx.RegisterResource("kubernetes:core/v1:ServiceList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceList gets an existing ServiceList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceListState, opts ...pulumi.ResourceOption) (*ServiceList, error) {
	var resource ServiceList
	err := ctx.ReadResource("kubernetes:core/v1:ServiceList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceList resources.
type serviceListState struct {
}

type ServiceListState struct {
}

func (ServiceListState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceListState)(nil)).Elem()
}

type serviceListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of services
	Items []ServiceType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ServiceList resource.
type ServiceListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of services
	Items ServiceTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ServiceListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceListArgs)(nil)).Elem()
}

type ServiceListInput interface {
	pulumi.Input

	ToServiceListOutput() ServiceListOutput
	ToServiceListOutputWithContext(ctx context.Context) ServiceListOutput
}

func (*ServiceList) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceList)(nil)).Elem()
}

func (i *ServiceList) ToServiceListOutput() ServiceListOutput {
	return i.ToServiceListOutputWithContext(context.Background())
}

func (i *ServiceList) ToServiceListOutputWithContext(ctx context.Context) ServiceListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceListOutput)
}

func (i *ServiceList) ToOutput(ctx context.Context) pulumix.Output[*ServiceList] {
	return pulumix.Output[*ServiceList]{
		OutputState: i.ToServiceListOutputWithContext(ctx).OutputState,
	}
}

// ServiceListArrayInput is an input type that accepts ServiceListArray and ServiceListArrayOutput values.
// You can construct a concrete instance of `ServiceListArrayInput` via:
//
//	ServiceListArray{ ServiceListArgs{...} }
type ServiceListArrayInput interface {
	pulumi.Input

	ToServiceListArrayOutput() ServiceListArrayOutput
	ToServiceListArrayOutputWithContext(context.Context) ServiceListArrayOutput
}

type ServiceListArray []ServiceListInput

func (ServiceListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceList)(nil)).Elem()
}

func (i ServiceListArray) ToServiceListArrayOutput() ServiceListArrayOutput {
	return i.ToServiceListArrayOutputWithContext(context.Background())
}

func (i ServiceListArray) ToServiceListArrayOutputWithContext(ctx context.Context) ServiceListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceListArrayOutput)
}

func (i ServiceListArray) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceList] {
	return pulumix.Output[[]*ServiceList]{
		OutputState: i.ToServiceListArrayOutputWithContext(ctx).OutputState,
	}
}

// ServiceListMapInput is an input type that accepts ServiceListMap and ServiceListMapOutput values.
// You can construct a concrete instance of `ServiceListMapInput` via:
//
//	ServiceListMap{ "key": ServiceListArgs{...} }
type ServiceListMapInput interface {
	pulumi.Input

	ToServiceListMapOutput() ServiceListMapOutput
	ToServiceListMapOutputWithContext(context.Context) ServiceListMapOutput
}

type ServiceListMap map[string]ServiceListInput

func (ServiceListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceList)(nil)).Elem()
}

func (i ServiceListMap) ToServiceListMapOutput() ServiceListMapOutput {
	return i.ToServiceListMapOutputWithContext(context.Background())
}

func (i ServiceListMap) ToServiceListMapOutputWithContext(ctx context.Context) ServiceListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceListMapOutput)
}

func (i ServiceListMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceList] {
	return pulumix.Output[map[string]*ServiceList]{
		OutputState: i.ToServiceListMapOutputWithContext(ctx).OutputState,
	}
}

type ServiceListOutput struct{ *pulumi.OutputState }

func (ServiceListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceList)(nil)).Elem()
}

func (o ServiceListOutput) ToServiceListOutput() ServiceListOutput {
	return o
}

func (o ServiceListOutput) ToServiceListOutputWithContext(ctx context.Context) ServiceListOutput {
	return o
}

func (o ServiceListOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceList] {
	return pulumix.Output[*ServiceList]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ServiceListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of services
func (o ServiceListOutput) Items() ServiceTypeArrayOutput {
	return o.ApplyT(func(v *ServiceList) ServiceTypeArrayOutput { return v.Items }).(ServiceTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ServiceListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ServiceListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ServiceList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ServiceListArrayOutput struct{ *pulumi.OutputState }

func (ServiceListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceList)(nil)).Elem()
}

func (o ServiceListArrayOutput) ToServiceListArrayOutput() ServiceListArrayOutput {
	return o
}

func (o ServiceListArrayOutput) ToServiceListArrayOutputWithContext(ctx context.Context) ServiceListArrayOutput {
	return o
}

func (o ServiceListArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceList] {
	return pulumix.Output[[]*ServiceList]{
		OutputState: o.OutputState,
	}
}

func (o ServiceListArrayOutput) Index(i pulumi.IntInput) ServiceListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceList {
		return vs[0].([]*ServiceList)[vs[1].(int)]
	}).(ServiceListOutput)
}

type ServiceListMapOutput struct{ *pulumi.OutputState }

func (ServiceListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceList)(nil)).Elem()
}

func (o ServiceListMapOutput) ToServiceListMapOutput() ServiceListMapOutput {
	return o
}

func (o ServiceListMapOutput) ToServiceListMapOutputWithContext(ctx context.Context) ServiceListMapOutput {
	return o
}

func (o ServiceListMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceList] {
	return pulumix.Output[map[string]*ServiceList]{
		OutputState: o.OutputState,
	}
}

func (o ServiceListMapOutput) MapIndex(k pulumi.StringInput) ServiceListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceList {
		return vs[0].(map[string]*ServiceList)[vs[1].(string)]
	}).(ServiceListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceListInput)(nil)).Elem(), &ServiceList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceListArrayInput)(nil)).Elem(), ServiceListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceListMapInput)(nil)).Elem(), ServiceListMap{})
	pulumi.RegisterOutputType(ServiceListOutput{})
	pulumi.RegisterOutputType(ServiceListArrayOutput{})
	pulumi.RegisterOutputType(ServiceListMapOutput{})
}
