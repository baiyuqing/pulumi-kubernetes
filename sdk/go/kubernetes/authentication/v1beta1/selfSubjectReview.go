// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SelfSubjectReview contains the user information that the kube-apiserver has about the user making this request. When using impersonation, users will receive the user info of the user being impersonated.  If impersonation or request header authentication is used, any extra keys will have their case ignored and returned as lowercase.
type SelfSubjectReview struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Status is filled in by the server with the user attributes.
	Status SelfSubjectReviewStatusPtrOutput `pulumi:"status"`
}

// NewSelfSubjectReview registers a new resource with the given unique name, arguments, and options.
func NewSelfSubjectReview(ctx *pulumi.Context,
	name string, args *SelfSubjectReviewArgs, opts ...pulumi.ResourceOption) (*SelfSubjectReview, error) {
	if args == nil {
		args = &SelfSubjectReviewArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("authentication.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("SelfSubjectReview")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authentication.k8s.io/v1alpha1:SelfSubjectReview"),
		},
	})
	opts = append(opts, aliases)
	var resource SelfSubjectReview
	err := ctx.RegisterResource("kubernetes:authentication.k8s.io/v1beta1:SelfSubjectReview", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSelfSubjectReview gets an existing SelfSubjectReview resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSelfSubjectReview(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SelfSubjectReviewState, opts ...pulumi.ResourceOption) (*SelfSubjectReview, error) {
	var resource SelfSubjectReview
	err := ctx.ReadResource("kubernetes:authentication.k8s.io/v1beta1:SelfSubjectReview", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SelfSubjectReview resources.
type selfSubjectReviewState struct {
}

type SelfSubjectReviewState struct {
}

func (SelfSubjectReviewState) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectReviewState)(nil)).Elem()
}

type selfSubjectReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a SelfSubjectReview resource.
type SelfSubjectReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
}

func (SelfSubjectReviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectReviewArgs)(nil)).Elem()
}

type SelfSubjectReviewInput interface {
	pulumi.Input

	ToSelfSubjectReviewOutput() SelfSubjectReviewOutput
	ToSelfSubjectReviewOutputWithContext(ctx context.Context) SelfSubjectReviewOutput
}

func (*SelfSubjectReview) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSubjectReview)(nil)).Elem()
}

func (i *SelfSubjectReview) ToSelfSubjectReviewOutput() SelfSubjectReviewOutput {
	return i.ToSelfSubjectReviewOutputWithContext(context.Background())
}

func (i *SelfSubjectReview) ToSelfSubjectReviewOutputWithContext(ctx context.Context) SelfSubjectReviewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectReviewOutput)
}

// SelfSubjectReviewArrayInput is an input type that accepts SelfSubjectReviewArray and SelfSubjectReviewArrayOutput values.
// You can construct a concrete instance of `SelfSubjectReviewArrayInput` via:
//
//	SelfSubjectReviewArray{ SelfSubjectReviewArgs{...} }
type SelfSubjectReviewArrayInput interface {
	pulumi.Input

	ToSelfSubjectReviewArrayOutput() SelfSubjectReviewArrayOutput
	ToSelfSubjectReviewArrayOutputWithContext(context.Context) SelfSubjectReviewArrayOutput
}

type SelfSubjectReviewArray []SelfSubjectReviewInput

func (SelfSubjectReviewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SelfSubjectReview)(nil)).Elem()
}

func (i SelfSubjectReviewArray) ToSelfSubjectReviewArrayOutput() SelfSubjectReviewArrayOutput {
	return i.ToSelfSubjectReviewArrayOutputWithContext(context.Background())
}

func (i SelfSubjectReviewArray) ToSelfSubjectReviewArrayOutputWithContext(ctx context.Context) SelfSubjectReviewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectReviewArrayOutput)
}

// SelfSubjectReviewMapInput is an input type that accepts SelfSubjectReviewMap and SelfSubjectReviewMapOutput values.
// You can construct a concrete instance of `SelfSubjectReviewMapInput` via:
//
//	SelfSubjectReviewMap{ "key": SelfSubjectReviewArgs{...} }
type SelfSubjectReviewMapInput interface {
	pulumi.Input

	ToSelfSubjectReviewMapOutput() SelfSubjectReviewMapOutput
	ToSelfSubjectReviewMapOutputWithContext(context.Context) SelfSubjectReviewMapOutput
}

type SelfSubjectReviewMap map[string]SelfSubjectReviewInput

func (SelfSubjectReviewMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SelfSubjectReview)(nil)).Elem()
}

func (i SelfSubjectReviewMap) ToSelfSubjectReviewMapOutput() SelfSubjectReviewMapOutput {
	return i.ToSelfSubjectReviewMapOutputWithContext(context.Background())
}

func (i SelfSubjectReviewMap) ToSelfSubjectReviewMapOutputWithContext(ctx context.Context) SelfSubjectReviewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectReviewMapOutput)
}

type SelfSubjectReviewOutput struct{ *pulumi.OutputState }

func (SelfSubjectReviewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSubjectReview)(nil)).Elem()
}

func (o SelfSubjectReviewOutput) ToSelfSubjectReviewOutput() SelfSubjectReviewOutput {
	return o
}

func (o SelfSubjectReviewOutput) ToSelfSubjectReviewOutputWithContext(ctx context.Context) SelfSubjectReviewOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o SelfSubjectReviewOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSubjectReview) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o SelfSubjectReviewOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSubjectReview) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o SelfSubjectReviewOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *SelfSubjectReview) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Status is filled in by the server with the user attributes.
func (o SelfSubjectReviewOutput) Status() SelfSubjectReviewStatusPtrOutput {
	return o.ApplyT(func(v *SelfSubjectReview) SelfSubjectReviewStatusPtrOutput { return v.Status }).(SelfSubjectReviewStatusPtrOutput)
}

type SelfSubjectReviewArrayOutput struct{ *pulumi.OutputState }

func (SelfSubjectReviewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SelfSubjectReview)(nil)).Elem()
}

func (o SelfSubjectReviewArrayOutput) ToSelfSubjectReviewArrayOutput() SelfSubjectReviewArrayOutput {
	return o
}

func (o SelfSubjectReviewArrayOutput) ToSelfSubjectReviewArrayOutputWithContext(ctx context.Context) SelfSubjectReviewArrayOutput {
	return o
}

func (o SelfSubjectReviewArrayOutput) Index(i pulumi.IntInput) SelfSubjectReviewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SelfSubjectReview {
		return vs[0].([]*SelfSubjectReview)[vs[1].(int)]
	}).(SelfSubjectReviewOutput)
}

type SelfSubjectReviewMapOutput struct{ *pulumi.OutputState }

func (SelfSubjectReviewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SelfSubjectReview)(nil)).Elem()
}

func (o SelfSubjectReviewMapOutput) ToSelfSubjectReviewMapOutput() SelfSubjectReviewMapOutput {
	return o
}

func (o SelfSubjectReviewMapOutput) ToSelfSubjectReviewMapOutputWithContext(ctx context.Context) SelfSubjectReviewMapOutput {
	return o
}

func (o SelfSubjectReviewMapOutput) MapIndex(k pulumi.StringInput) SelfSubjectReviewOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SelfSubjectReview {
		return vs[0].(map[string]*SelfSubjectReview)[vs[1].(string)]
	}).(SelfSubjectReviewOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectReviewInput)(nil)).Elem(), &SelfSubjectReview{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectReviewArrayInput)(nil)).Elem(), SelfSubjectReviewArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectReviewMapInput)(nil)).Elem(), SelfSubjectReviewMap{})
	pulumi.RegisterOutputType(SelfSubjectReviewOutput{})
	pulumi.RegisterOutputType(SelfSubjectReviewArrayOutput{})
	pulumi.RegisterOutputType(SelfSubjectReviewMapOutput{})
}
