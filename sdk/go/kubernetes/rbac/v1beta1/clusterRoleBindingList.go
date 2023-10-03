// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ClusterRoleBindingList is a collection of ClusterRoleBindings. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRoleBindingList, and will no longer be served in v1.20.
type ClusterRoleBindingList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Items is a list of ClusterRoleBindings
	Items ClusterRoleBindingTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewClusterRoleBindingList registers a new resource with the given unique name, arguments, and options.
func NewClusterRoleBindingList(ctx *pulumi.Context,
	name string, args *ClusterRoleBindingListArgs, opts ...pulumi.ResourceOption) (*ClusterRoleBindingList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("ClusterRoleBindingList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ClusterRoleBindingList
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBindingList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterRoleBindingList gets an existing ClusterRoleBindingList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterRoleBindingList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterRoleBindingListState, opts ...pulumi.ResourceOption) (*ClusterRoleBindingList, error) {
	var resource ClusterRoleBindingList
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBindingList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterRoleBindingList resources.
type clusterRoleBindingListState struct {
}

type ClusterRoleBindingListState struct {
}

func (ClusterRoleBindingListState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleBindingListState)(nil)).Elem()
}

type clusterRoleBindingListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of ClusterRoleBindings
	Items []ClusterRoleBindingType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ClusterRoleBindingList resource.
type ClusterRoleBindingListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of ClusterRoleBindings
	Items ClusterRoleBindingTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ListMetaPtrInput
}

func (ClusterRoleBindingListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleBindingListArgs)(nil)).Elem()
}

type ClusterRoleBindingListInput interface {
	pulumi.Input

	ToClusterRoleBindingListOutput() ClusterRoleBindingListOutput
	ToClusterRoleBindingListOutputWithContext(ctx context.Context) ClusterRoleBindingListOutput
}

func (*ClusterRoleBindingList) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRoleBindingList)(nil)).Elem()
}

func (i *ClusterRoleBindingList) ToClusterRoleBindingListOutput() ClusterRoleBindingListOutput {
	return i.ToClusterRoleBindingListOutputWithContext(context.Background())
}

func (i *ClusterRoleBindingList) ToClusterRoleBindingListOutputWithContext(ctx context.Context) ClusterRoleBindingListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleBindingListOutput)
}

func (i *ClusterRoleBindingList) ToOutput(ctx context.Context) pulumix.Output[*ClusterRoleBindingList] {
	return pulumix.Output[*ClusterRoleBindingList]{
		OutputState: i.ToClusterRoleBindingListOutputWithContext(ctx).OutputState,
	}
}

// ClusterRoleBindingListArrayInput is an input type that accepts ClusterRoleBindingListArray and ClusterRoleBindingListArrayOutput values.
// You can construct a concrete instance of `ClusterRoleBindingListArrayInput` via:
//
//	ClusterRoleBindingListArray{ ClusterRoleBindingListArgs{...} }
type ClusterRoleBindingListArrayInput interface {
	pulumi.Input

	ToClusterRoleBindingListArrayOutput() ClusterRoleBindingListArrayOutput
	ToClusterRoleBindingListArrayOutputWithContext(context.Context) ClusterRoleBindingListArrayOutput
}

type ClusterRoleBindingListArray []ClusterRoleBindingListInput

func (ClusterRoleBindingListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterRoleBindingList)(nil)).Elem()
}

func (i ClusterRoleBindingListArray) ToClusterRoleBindingListArrayOutput() ClusterRoleBindingListArrayOutput {
	return i.ToClusterRoleBindingListArrayOutputWithContext(context.Background())
}

func (i ClusterRoleBindingListArray) ToClusterRoleBindingListArrayOutputWithContext(ctx context.Context) ClusterRoleBindingListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleBindingListArrayOutput)
}

func (i ClusterRoleBindingListArray) ToOutput(ctx context.Context) pulumix.Output[[]*ClusterRoleBindingList] {
	return pulumix.Output[[]*ClusterRoleBindingList]{
		OutputState: i.ToClusterRoleBindingListArrayOutputWithContext(ctx).OutputState,
	}
}

// ClusterRoleBindingListMapInput is an input type that accepts ClusterRoleBindingListMap and ClusterRoleBindingListMapOutput values.
// You can construct a concrete instance of `ClusterRoleBindingListMapInput` via:
//
//	ClusterRoleBindingListMap{ "key": ClusterRoleBindingListArgs{...} }
type ClusterRoleBindingListMapInput interface {
	pulumi.Input

	ToClusterRoleBindingListMapOutput() ClusterRoleBindingListMapOutput
	ToClusterRoleBindingListMapOutputWithContext(context.Context) ClusterRoleBindingListMapOutput
}

type ClusterRoleBindingListMap map[string]ClusterRoleBindingListInput

func (ClusterRoleBindingListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterRoleBindingList)(nil)).Elem()
}

func (i ClusterRoleBindingListMap) ToClusterRoleBindingListMapOutput() ClusterRoleBindingListMapOutput {
	return i.ToClusterRoleBindingListMapOutputWithContext(context.Background())
}

func (i ClusterRoleBindingListMap) ToClusterRoleBindingListMapOutputWithContext(ctx context.Context) ClusterRoleBindingListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleBindingListMapOutput)
}

func (i ClusterRoleBindingListMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ClusterRoleBindingList] {
	return pulumix.Output[map[string]*ClusterRoleBindingList]{
		OutputState: i.ToClusterRoleBindingListMapOutputWithContext(ctx).OutputState,
	}
}

type ClusterRoleBindingListOutput struct{ *pulumi.OutputState }

func (ClusterRoleBindingListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRoleBindingList)(nil)).Elem()
}

func (o ClusterRoleBindingListOutput) ToClusterRoleBindingListOutput() ClusterRoleBindingListOutput {
	return o
}

func (o ClusterRoleBindingListOutput) ToClusterRoleBindingListOutputWithContext(ctx context.Context) ClusterRoleBindingListOutput {
	return o
}

func (o ClusterRoleBindingListOutput) ToOutput(ctx context.Context) pulumix.Output[*ClusterRoleBindingList] {
	return pulumix.Output[*ClusterRoleBindingList]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ClusterRoleBindingListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterRoleBindingList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Items is a list of ClusterRoleBindings
func (o ClusterRoleBindingListOutput) Items() ClusterRoleBindingTypeArrayOutput {
	return o.ApplyT(func(v *ClusterRoleBindingList) ClusterRoleBindingTypeArrayOutput { return v.Items }).(ClusterRoleBindingTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ClusterRoleBindingListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterRoleBindingList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata.
func (o ClusterRoleBindingListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ClusterRoleBindingList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ClusterRoleBindingListArrayOutput struct{ *pulumi.OutputState }

func (ClusterRoleBindingListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterRoleBindingList)(nil)).Elem()
}

func (o ClusterRoleBindingListArrayOutput) ToClusterRoleBindingListArrayOutput() ClusterRoleBindingListArrayOutput {
	return o
}

func (o ClusterRoleBindingListArrayOutput) ToClusterRoleBindingListArrayOutputWithContext(ctx context.Context) ClusterRoleBindingListArrayOutput {
	return o
}

func (o ClusterRoleBindingListArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ClusterRoleBindingList] {
	return pulumix.Output[[]*ClusterRoleBindingList]{
		OutputState: o.OutputState,
	}
}

func (o ClusterRoleBindingListArrayOutput) Index(i pulumi.IntInput) ClusterRoleBindingListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterRoleBindingList {
		return vs[0].([]*ClusterRoleBindingList)[vs[1].(int)]
	}).(ClusterRoleBindingListOutput)
}

type ClusterRoleBindingListMapOutput struct{ *pulumi.OutputState }

func (ClusterRoleBindingListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterRoleBindingList)(nil)).Elem()
}

func (o ClusterRoleBindingListMapOutput) ToClusterRoleBindingListMapOutput() ClusterRoleBindingListMapOutput {
	return o
}

func (o ClusterRoleBindingListMapOutput) ToClusterRoleBindingListMapOutputWithContext(ctx context.Context) ClusterRoleBindingListMapOutput {
	return o
}

func (o ClusterRoleBindingListMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ClusterRoleBindingList] {
	return pulumix.Output[map[string]*ClusterRoleBindingList]{
		OutputState: o.OutputState,
	}
}

func (o ClusterRoleBindingListMapOutput) MapIndex(k pulumi.StringInput) ClusterRoleBindingListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterRoleBindingList {
		return vs[0].(map[string]*ClusterRoleBindingList)[vs[1].(string)]
	}).(ClusterRoleBindingListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoleBindingListInput)(nil)).Elem(), &ClusterRoleBindingList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoleBindingListArrayInput)(nil)).Elem(), ClusterRoleBindingListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoleBindingListMapInput)(nil)).Elem(), ClusterRoleBindingListMap{})
	pulumi.RegisterOutputType(ClusterRoleBindingListOutput{})
	pulumi.RegisterOutputType(ClusterRoleBindingListArrayOutput{})
	pulumi.RegisterOutputType(ClusterRoleBindingListMapOutput{})
}
