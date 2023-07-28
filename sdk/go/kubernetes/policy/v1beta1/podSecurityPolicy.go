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

// PodSecurityPolicy governs the ability to make requests that affect the Security Context that will be applied to a pod and container.
type PodSecurityPolicy struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// spec defines the policy enforced.
	Spec PodSecurityPolicySpecOutput `pulumi:"spec"`
}

// NewPodSecurityPolicy registers a new resource with the given unique name, arguments, and options.
func NewPodSecurityPolicy(ctx *pulumi.Context,
	name string, args *PodSecurityPolicyArgs, opts ...pulumi.ResourceOption) (*PodSecurityPolicy, error) {
	if args == nil {
		args = &PodSecurityPolicyArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("policy/v1beta1")
	args.Kind = pulumi.StringPtr("PodSecurityPolicy")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:extensions/v1beta1:PodSecurityPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PodSecurityPolicy
	err := ctx.RegisterResource("kubernetes:policy/v1beta1:PodSecurityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodSecurityPolicy gets an existing PodSecurityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodSecurityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodSecurityPolicyState, opts ...pulumi.ResourceOption) (*PodSecurityPolicy, error) {
	var resource PodSecurityPolicy
	err := ctx.ReadResource("kubernetes:policy/v1beta1:PodSecurityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodSecurityPolicy resources.
type podSecurityPolicyState struct {
}

type PodSecurityPolicyState struct {
}

func (PodSecurityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*podSecurityPolicyState)(nil)).Elem()
}

type podSecurityPolicyArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec defines the policy enforced.
	Spec *PodSecurityPolicySpec `pulumi:"spec"`
}

// The set of arguments for constructing a PodSecurityPolicy resource.
type PodSecurityPolicyArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// spec defines the policy enforced.
	Spec PodSecurityPolicySpecPtrInput
}

func (PodSecurityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podSecurityPolicyArgs)(nil)).Elem()
}

type PodSecurityPolicyInput interface {
	pulumi.Input

	ToPodSecurityPolicyOutput() PodSecurityPolicyOutput
	ToPodSecurityPolicyOutputWithContext(ctx context.Context) PodSecurityPolicyOutput
}

func (*PodSecurityPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**PodSecurityPolicy)(nil)).Elem()
}

func (i *PodSecurityPolicy) ToPodSecurityPolicyOutput() PodSecurityPolicyOutput {
	return i.ToPodSecurityPolicyOutputWithContext(context.Background())
}

func (i *PodSecurityPolicy) ToPodSecurityPolicyOutputWithContext(ctx context.Context) PodSecurityPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodSecurityPolicyOutput)
}

// PodSecurityPolicyArrayInput is an input type that accepts PodSecurityPolicyArray and PodSecurityPolicyArrayOutput values.
// You can construct a concrete instance of `PodSecurityPolicyArrayInput` via:
//
//	PodSecurityPolicyArray{ PodSecurityPolicyArgs{...} }
type PodSecurityPolicyArrayInput interface {
	pulumi.Input

	ToPodSecurityPolicyArrayOutput() PodSecurityPolicyArrayOutput
	ToPodSecurityPolicyArrayOutputWithContext(context.Context) PodSecurityPolicyArrayOutput
}

type PodSecurityPolicyArray []PodSecurityPolicyInput

func (PodSecurityPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodSecurityPolicy)(nil)).Elem()
}

func (i PodSecurityPolicyArray) ToPodSecurityPolicyArrayOutput() PodSecurityPolicyArrayOutput {
	return i.ToPodSecurityPolicyArrayOutputWithContext(context.Background())
}

func (i PodSecurityPolicyArray) ToPodSecurityPolicyArrayOutputWithContext(ctx context.Context) PodSecurityPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodSecurityPolicyArrayOutput)
}

// PodSecurityPolicyMapInput is an input type that accepts PodSecurityPolicyMap and PodSecurityPolicyMapOutput values.
// You can construct a concrete instance of `PodSecurityPolicyMapInput` via:
//
//	PodSecurityPolicyMap{ "key": PodSecurityPolicyArgs{...} }
type PodSecurityPolicyMapInput interface {
	pulumi.Input

	ToPodSecurityPolicyMapOutput() PodSecurityPolicyMapOutput
	ToPodSecurityPolicyMapOutputWithContext(context.Context) PodSecurityPolicyMapOutput
}

type PodSecurityPolicyMap map[string]PodSecurityPolicyInput

func (PodSecurityPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodSecurityPolicy)(nil)).Elem()
}

func (i PodSecurityPolicyMap) ToPodSecurityPolicyMapOutput() PodSecurityPolicyMapOutput {
	return i.ToPodSecurityPolicyMapOutputWithContext(context.Background())
}

func (i PodSecurityPolicyMap) ToPodSecurityPolicyMapOutputWithContext(ctx context.Context) PodSecurityPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodSecurityPolicyMapOutput)
}

type PodSecurityPolicyOutput struct{ *pulumi.OutputState }

func (PodSecurityPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PodSecurityPolicy)(nil)).Elem()
}

func (o PodSecurityPolicyOutput) ToPodSecurityPolicyOutput() PodSecurityPolicyOutput {
	return o
}

func (o PodSecurityPolicyOutput) ToPodSecurityPolicyOutputWithContext(ctx context.Context) PodSecurityPolicyOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PodSecurityPolicyOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PodSecurityPolicy) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PodSecurityPolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PodSecurityPolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PodSecurityPolicyOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *PodSecurityPolicy) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// spec defines the policy enforced.
func (o PodSecurityPolicyOutput) Spec() PodSecurityPolicySpecOutput {
	return o.ApplyT(func(v *PodSecurityPolicy) PodSecurityPolicySpecOutput { return v.Spec }).(PodSecurityPolicySpecOutput)
}

type PodSecurityPolicyArrayOutput struct{ *pulumi.OutputState }

func (PodSecurityPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodSecurityPolicy)(nil)).Elem()
}

func (o PodSecurityPolicyArrayOutput) ToPodSecurityPolicyArrayOutput() PodSecurityPolicyArrayOutput {
	return o
}

func (o PodSecurityPolicyArrayOutput) ToPodSecurityPolicyArrayOutputWithContext(ctx context.Context) PodSecurityPolicyArrayOutput {
	return o
}

func (o PodSecurityPolicyArrayOutput) Index(i pulumi.IntInput) PodSecurityPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PodSecurityPolicy {
		return vs[0].([]*PodSecurityPolicy)[vs[1].(int)]
	}).(PodSecurityPolicyOutput)
}

type PodSecurityPolicyMapOutput struct{ *pulumi.OutputState }

func (PodSecurityPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodSecurityPolicy)(nil)).Elem()
}

func (o PodSecurityPolicyMapOutput) ToPodSecurityPolicyMapOutput() PodSecurityPolicyMapOutput {
	return o
}

func (o PodSecurityPolicyMapOutput) ToPodSecurityPolicyMapOutputWithContext(ctx context.Context) PodSecurityPolicyMapOutput {
	return o
}

func (o PodSecurityPolicyMapOutput) MapIndex(k pulumi.StringInput) PodSecurityPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PodSecurityPolicy {
		return vs[0].(map[string]*PodSecurityPolicy)[vs[1].(string)]
	}).(PodSecurityPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PodSecurityPolicyInput)(nil)).Elem(), &PodSecurityPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodSecurityPolicyArrayInput)(nil)).Elem(), PodSecurityPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodSecurityPolicyMapInput)(nil)).Elem(), PodSecurityPolicyMap{})
	pulumi.RegisterOutputType(PodSecurityPolicyOutput{})
	pulumi.RegisterOutputType(PodSecurityPolicyArrayOutput{})
	pulumi.RegisterOutputType(PodSecurityPolicyMapOutput{})
}
