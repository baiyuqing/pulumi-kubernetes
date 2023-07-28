// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/internal"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 RoleBinding, and will no longer be served in v1.20.
type RoleBindingPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefPatchPtrOutput `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectPatchArrayOutput `pulumi:"subjects"`
}

// NewRoleBindingPatch registers a new resource with the given unique name, arguments, and options.
func NewRoleBindingPatch(ctx *pulumi.Context,
	name string, args *RoleBindingPatchArgs, opts ...pulumi.ResourceOption) (*RoleBindingPatch, error) {
	if args == nil {
		args = &RoleBindingPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("RoleBinding")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1:RoleBindingPatch"),
		},
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBindingPatch"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RoleBindingPatch
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBindingPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleBindingPatch gets an existing RoleBindingPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleBindingPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleBindingPatchState, opts ...pulumi.ResourceOption) (*RoleBindingPatch, error) {
	var resource RoleBindingPatch
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBindingPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleBindingPatch resources.
type roleBindingPatchState struct {
}

type RoleBindingPatchState struct {
}

func (RoleBindingPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleBindingPatchState)(nil)).Elem()
}

type roleBindingPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef *RoleRefPatch `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects []SubjectPatch `pulumi:"subjects"`
}

// The set of arguments for constructing a RoleBindingPatch resource.
type RoleBindingPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPatchPtrInput
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefPatchPtrInput
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectPatchArrayInput
}

func (RoleBindingPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleBindingPatchArgs)(nil)).Elem()
}

type RoleBindingPatchInput interface {
	pulumi.Input

	ToRoleBindingPatchOutput() RoleBindingPatchOutput
	ToRoleBindingPatchOutputWithContext(ctx context.Context) RoleBindingPatchOutput
}

func (*RoleBindingPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleBindingPatch)(nil)).Elem()
}

func (i *RoleBindingPatch) ToRoleBindingPatchOutput() RoleBindingPatchOutput {
	return i.ToRoleBindingPatchOutputWithContext(context.Background())
}

func (i *RoleBindingPatch) ToRoleBindingPatchOutputWithContext(ctx context.Context) RoleBindingPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingPatchOutput)
}

// RoleBindingPatchArrayInput is an input type that accepts RoleBindingPatchArray and RoleBindingPatchArrayOutput values.
// You can construct a concrete instance of `RoleBindingPatchArrayInput` via:
//
//	RoleBindingPatchArray{ RoleBindingPatchArgs{...} }
type RoleBindingPatchArrayInput interface {
	pulumi.Input

	ToRoleBindingPatchArrayOutput() RoleBindingPatchArrayOutput
	ToRoleBindingPatchArrayOutputWithContext(context.Context) RoleBindingPatchArrayOutput
}

type RoleBindingPatchArray []RoleBindingPatchInput

func (RoleBindingPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleBindingPatch)(nil)).Elem()
}

func (i RoleBindingPatchArray) ToRoleBindingPatchArrayOutput() RoleBindingPatchArrayOutput {
	return i.ToRoleBindingPatchArrayOutputWithContext(context.Background())
}

func (i RoleBindingPatchArray) ToRoleBindingPatchArrayOutputWithContext(ctx context.Context) RoleBindingPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingPatchArrayOutput)
}

// RoleBindingPatchMapInput is an input type that accepts RoleBindingPatchMap and RoleBindingPatchMapOutput values.
// You can construct a concrete instance of `RoleBindingPatchMapInput` via:
//
//	RoleBindingPatchMap{ "key": RoleBindingPatchArgs{...} }
type RoleBindingPatchMapInput interface {
	pulumi.Input

	ToRoleBindingPatchMapOutput() RoleBindingPatchMapOutput
	ToRoleBindingPatchMapOutputWithContext(context.Context) RoleBindingPatchMapOutput
}

type RoleBindingPatchMap map[string]RoleBindingPatchInput

func (RoleBindingPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleBindingPatch)(nil)).Elem()
}

func (i RoleBindingPatchMap) ToRoleBindingPatchMapOutput() RoleBindingPatchMapOutput {
	return i.ToRoleBindingPatchMapOutputWithContext(context.Background())
}

func (i RoleBindingPatchMap) ToRoleBindingPatchMapOutputWithContext(ctx context.Context) RoleBindingPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingPatchMapOutput)
}

type RoleBindingPatchOutput struct{ *pulumi.OutputState }

func (RoleBindingPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleBindingPatch)(nil)).Elem()
}

func (o RoleBindingPatchOutput) ToRoleBindingPatchOutput() RoleBindingPatchOutput {
	return o
}

func (o RoleBindingPatchOutput) ToRoleBindingPatchOutputWithContext(ctx context.Context) RoleBindingPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o RoleBindingPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleBindingPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o RoleBindingPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleBindingPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata.
func (o RoleBindingPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *RoleBindingPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
func (o RoleBindingPatchOutput) RoleRef() RoleRefPatchPtrOutput {
	return o.ApplyT(func(v *RoleBindingPatch) RoleRefPatchPtrOutput { return v.RoleRef }).(RoleRefPatchPtrOutput)
}

// Subjects holds references to the objects the role applies to.
func (o RoleBindingPatchOutput) Subjects() SubjectPatchArrayOutput {
	return o.ApplyT(func(v *RoleBindingPatch) SubjectPatchArrayOutput { return v.Subjects }).(SubjectPatchArrayOutput)
}

type RoleBindingPatchArrayOutput struct{ *pulumi.OutputState }

func (RoleBindingPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleBindingPatch)(nil)).Elem()
}

func (o RoleBindingPatchArrayOutput) ToRoleBindingPatchArrayOutput() RoleBindingPatchArrayOutput {
	return o
}

func (o RoleBindingPatchArrayOutput) ToRoleBindingPatchArrayOutputWithContext(ctx context.Context) RoleBindingPatchArrayOutput {
	return o
}

func (o RoleBindingPatchArrayOutput) Index(i pulumi.IntInput) RoleBindingPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RoleBindingPatch {
		return vs[0].([]*RoleBindingPatch)[vs[1].(int)]
	}).(RoleBindingPatchOutput)
}

type RoleBindingPatchMapOutput struct{ *pulumi.OutputState }

func (RoleBindingPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleBindingPatch)(nil)).Elem()
}

func (o RoleBindingPatchMapOutput) ToRoleBindingPatchMapOutput() RoleBindingPatchMapOutput {
	return o
}

func (o RoleBindingPatchMapOutput) ToRoleBindingPatchMapOutputWithContext(ctx context.Context) RoleBindingPatchMapOutput {
	return o
}

func (o RoleBindingPatchMapOutput) MapIndex(k pulumi.StringInput) RoleBindingPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RoleBindingPatch {
		return vs[0].(map[string]*RoleBindingPatch)[vs[1].(string)]
	}).(RoleBindingPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleBindingPatchInput)(nil)).Elem(), &RoleBindingPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleBindingPatchArrayInput)(nil)).Elem(), RoleBindingPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleBindingPatchMapInput)(nil)).Elem(), RoleBindingPatchMap{})
	pulumi.RegisterOutputType(RoleBindingPatchOutput{})
	pulumi.RegisterOutputType(RoleBindingPatchArrayOutput{})
	pulumi.RegisterOutputType(RoleBindingPatchMapOutput{})
}
