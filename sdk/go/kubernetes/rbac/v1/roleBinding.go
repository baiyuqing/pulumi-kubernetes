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

// RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace.
type RoleBinding struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefOutput `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectArrayOutput `pulumi:"subjects"`
}

// NewRoleBinding registers a new resource with the given unique name, arguments, and options.
func NewRoleBinding(ctx *pulumi.Context,
	name string, args *RoleBindingArgs, opts ...pulumi.ResourceOption) (*RoleBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleRef == nil {
		return nil, errors.New("invalid value for required argument 'RoleRef'")
	}
	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1")
	args.Kind = pulumi.StringPtr("RoleBinding")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBinding"),
		},
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBinding"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RoleBinding
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1:RoleBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleBinding gets an existing RoleBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleBindingState, opts ...pulumi.ResourceOption) (*RoleBinding, error) {
	var resource RoleBinding
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1:RoleBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleBinding resources.
type roleBindingState struct {
}

type RoleBindingState struct {
}

func (RoleBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleBindingState)(nil)).Elem()
}

type roleBindingArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRef `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects []Subject `pulumi:"subjects"`
}

// The set of arguments for constructing a RoleBinding resource.
type RoleBindingArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrInput
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefInput
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectArrayInput
}

func (RoleBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleBindingArgs)(nil)).Elem()
}

type RoleBindingInput interface {
	pulumi.Input

	ToRoleBindingOutput() RoleBindingOutput
	ToRoleBindingOutputWithContext(ctx context.Context) RoleBindingOutput
}

func (*RoleBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleBinding)(nil)).Elem()
}

func (i *RoleBinding) ToRoleBindingOutput() RoleBindingOutput {
	return i.ToRoleBindingOutputWithContext(context.Background())
}

func (i *RoleBinding) ToRoleBindingOutputWithContext(ctx context.Context) RoleBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingOutput)
}

// RoleBindingArrayInput is an input type that accepts RoleBindingArray and RoleBindingArrayOutput values.
// You can construct a concrete instance of `RoleBindingArrayInput` via:
//
//	RoleBindingArray{ RoleBindingArgs{...} }
type RoleBindingArrayInput interface {
	pulumi.Input

	ToRoleBindingArrayOutput() RoleBindingArrayOutput
	ToRoleBindingArrayOutputWithContext(context.Context) RoleBindingArrayOutput
}

type RoleBindingArray []RoleBindingInput

func (RoleBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleBinding)(nil)).Elem()
}

func (i RoleBindingArray) ToRoleBindingArrayOutput() RoleBindingArrayOutput {
	return i.ToRoleBindingArrayOutputWithContext(context.Background())
}

func (i RoleBindingArray) ToRoleBindingArrayOutputWithContext(ctx context.Context) RoleBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingArrayOutput)
}

// RoleBindingMapInput is an input type that accepts RoleBindingMap and RoleBindingMapOutput values.
// You can construct a concrete instance of `RoleBindingMapInput` via:
//
//	RoleBindingMap{ "key": RoleBindingArgs{...} }
type RoleBindingMapInput interface {
	pulumi.Input

	ToRoleBindingMapOutput() RoleBindingMapOutput
	ToRoleBindingMapOutputWithContext(context.Context) RoleBindingMapOutput
}

type RoleBindingMap map[string]RoleBindingInput

func (RoleBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleBinding)(nil)).Elem()
}

func (i RoleBindingMap) ToRoleBindingMapOutput() RoleBindingMapOutput {
	return i.ToRoleBindingMapOutputWithContext(context.Background())
}

func (i RoleBindingMap) ToRoleBindingMapOutputWithContext(ctx context.Context) RoleBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingMapOutput)
}

type RoleBindingOutput struct{ *pulumi.OutputState }

func (RoleBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleBinding)(nil)).Elem()
}

func (o RoleBindingOutput) ToRoleBindingOutput() RoleBindingOutput {
	return o
}

func (o RoleBindingOutput) ToRoleBindingOutputWithContext(ctx context.Context) RoleBindingOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o RoleBindingOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleBinding) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o RoleBindingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleBinding) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata.
func (o RoleBindingOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *RoleBinding) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
func (o RoleBindingOutput) RoleRef() RoleRefOutput {
	return o.ApplyT(func(v *RoleBinding) RoleRefOutput { return v.RoleRef }).(RoleRefOutput)
}

// Subjects holds references to the objects the role applies to.
func (o RoleBindingOutput) Subjects() SubjectArrayOutput {
	return o.ApplyT(func(v *RoleBinding) SubjectArrayOutput { return v.Subjects }).(SubjectArrayOutput)
}

type RoleBindingArrayOutput struct{ *pulumi.OutputState }

func (RoleBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleBinding)(nil)).Elem()
}

func (o RoleBindingArrayOutput) ToRoleBindingArrayOutput() RoleBindingArrayOutput {
	return o
}

func (o RoleBindingArrayOutput) ToRoleBindingArrayOutputWithContext(ctx context.Context) RoleBindingArrayOutput {
	return o
}

func (o RoleBindingArrayOutput) Index(i pulumi.IntInput) RoleBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RoleBinding {
		return vs[0].([]*RoleBinding)[vs[1].(int)]
	}).(RoleBindingOutput)
}

type RoleBindingMapOutput struct{ *pulumi.OutputState }

func (RoleBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleBinding)(nil)).Elem()
}

func (o RoleBindingMapOutput) ToRoleBindingMapOutput() RoleBindingMapOutput {
	return o
}

func (o RoleBindingMapOutput) ToRoleBindingMapOutputWithContext(ctx context.Context) RoleBindingMapOutput {
	return o
}

func (o RoleBindingMapOutput) MapIndex(k pulumi.StringInput) RoleBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RoleBinding {
		return vs[0].(map[string]*RoleBinding)[vs[1].(string)]
	}).(RoleBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleBindingInput)(nil)).Elem(), &RoleBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleBindingArrayInput)(nil)).Elem(), RoleBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleBindingMapInput)(nil)).Elem(), RoleBindingMap{})
	pulumi.RegisterOutputType(RoleBindingOutput{})
	pulumi.RegisterOutputType(RoleBindingArrayOutput{})
	pulumi.RegisterOutputType(RoleBindingMapOutput{})
}
