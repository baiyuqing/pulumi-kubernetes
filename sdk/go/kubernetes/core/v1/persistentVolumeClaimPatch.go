// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

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
// PersistentVolumeClaim is a user's request for and claim to a persistent volume
type PersistentVolumeClaimPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec PersistentVolumeClaimSpecPatchPtrOutput `pulumi:"spec"`
	// status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Status PersistentVolumeClaimStatusPatchPtrOutput `pulumi:"status"`
}

// NewPersistentVolumeClaimPatch registers a new resource with the given unique name, arguments, and options.
func NewPersistentVolumeClaimPatch(ctx *pulumi.Context,
	name string, args *PersistentVolumeClaimPatchArgs, opts ...pulumi.ResourceOption) (*PersistentVolumeClaimPatch, error) {
	if args == nil {
		args = &PersistentVolumeClaimPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("PersistentVolumeClaim")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PersistentVolumeClaimPatch
	err := ctx.RegisterResource("kubernetes:core/v1:PersistentVolumeClaimPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistentVolumeClaimPatch gets an existing PersistentVolumeClaimPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistentVolumeClaimPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersistentVolumeClaimPatchState, opts ...pulumi.ResourceOption) (*PersistentVolumeClaimPatch, error) {
	var resource PersistentVolumeClaimPatch
	err := ctx.ReadResource("kubernetes:core/v1:PersistentVolumeClaimPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersistentVolumeClaimPatch resources.
type persistentVolumeClaimPatchState struct {
}

type PersistentVolumeClaimPatchState struct {
}

func (PersistentVolumeClaimPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*persistentVolumeClaimPatchState)(nil)).Elem()
}

type persistentVolumeClaimPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec *PersistentVolumeClaimSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a PersistentVolumeClaimPatch resource.
type PersistentVolumeClaimPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec PersistentVolumeClaimSpecPatchPtrInput
}

func (PersistentVolumeClaimPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*persistentVolumeClaimPatchArgs)(nil)).Elem()
}

type PersistentVolumeClaimPatchInput interface {
	pulumi.Input

	ToPersistentVolumeClaimPatchOutput() PersistentVolumeClaimPatchOutput
	ToPersistentVolumeClaimPatchOutputWithContext(ctx context.Context) PersistentVolumeClaimPatchOutput
}

func (*PersistentVolumeClaimPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentVolumeClaimPatch)(nil)).Elem()
}

func (i *PersistentVolumeClaimPatch) ToPersistentVolumeClaimPatchOutput() PersistentVolumeClaimPatchOutput {
	return i.ToPersistentVolumeClaimPatchOutputWithContext(context.Background())
}

func (i *PersistentVolumeClaimPatch) ToPersistentVolumeClaimPatchOutputWithContext(ctx context.Context) PersistentVolumeClaimPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimPatchOutput)
}

// PersistentVolumeClaimPatchArrayInput is an input type that accepts PersistentVolumeClaimPatchArray and PersistentVolumeClaimPatchArrayOutput values.
// You can construct a concrete instance of `PersistentVolumeClaimPatchArrayInput` via:
//
//	PersistentVolumeClaimPatchArray{ PersistentVolumeClaimPatchArgs{...} }
type PersistentVolumeClaimPatchArrayInput interface {
	pulumi.Input

	ToPersistentVolumeClaimPatchArrayOutput() PersistentVolumeClaimPatchArrayOutput
	ToPersistentVolumeClaimPatchArrayOutputWithContext(context.Context) PersistentVolumeClaimPatchArrayOutput
}

type PersistentVolumeClaimPatchArray []PersistentVolumeClaimPatchInput

func (PersistentVolumeClaimPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistentVolumeClaimPatch)(nil)).Elem()
}

func (i PersistentVolumeClaimPatchArray) ToPersistentVolumeClaimPatchArrayOutput() PersistentVolumeClaimPatchArrayOutput {
	return i.ToPersistentVolumeClaimPatchArrayOutputWithContext(context.Background())
}

func (i PersistentVolumeClaimPatchArray) ToPersistentVolumeClaimPatchArrayOutputWithContext(ctx context.Context) PersistentVolumeClaimPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimPatchArrayOutput)
}

// PersistentVolumeClaimPatchMapInput is an input type that accepts PersistentVolumeClaimPatchMap and PersistentVolumeClaimPatchMapOutput values.
// You can construct a concrete instance of `PersistentVolumeClaimPatchMapInput` via:
//
//	PersistentVolumeClaimPatchMap{ "key": PersistentVolumeClaimPatchArgs{...} }
type PersistentVolumeClaimPatchMapInput interface {
	pulumi.Input

	ToPersistentVolumeClaimPatchMapOutput() PersistentVolumeClaimPatchMapOutput
	ToPersistentVolumeClaimPatchMapOutputWithContext(context.Context) PersistentVolumeClaimPatchMapOutput
}

type PersistentVolumeClaimPatchMap map[string]PersistentVolumeClaimPatchInput

func (PersistentVolumeClaimPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistentVolumeClaimPatch)(nil)).Elem()
}

func (i PersistentVolumeClaimPatchMap) ToPersistentVolumeClaimPatchMapOutput() PersistentVolumeClaimPatchMapOutput {
	return i.ToPersistentVolumeClaimPatchMapOutputWithContext(context.Background())
}

func (i PersistentVolumeClaimPatchMap) ToPersistentVolumeClaimPatchMapOutputWithContext(ctx context.Context) PersistentVolumeClaimPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimPatchMapOutput)
}

type PersistentVolumeClaimPatchOutput struct{ *pulumi.OutputState }

func (PersistentVolumeClaimPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentVolumeClaimPatch)(nil)).Elem()
}

func (o PersistentVolumeClaimPatchOutput) ToPersistentVolumeClaimPatchOutput() PersistentVolumeClaimPatchOutput {
	return o
}

func (o PersistentVolumeClaimPatchOutput) ToPersistentVolumeClaimPatchOutputWithContext(ctx context.Context) PersistentVolumeClaimPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PersistentVolumeClaimPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersistentVolumeClaimPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PersistentVolumeClaimPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersistentVolumeClaimPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PersistentVolumeClaimPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *PersistentVolumeClaimPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
func (o PersistentVolumeClaimPatchOutput) Spec() PersistentVolumeClaimSpecPatchPtrOutput {
	return o.ApplyT(func(v *PersistentVolumeClaimPatch) PersistentVolumeClaimSpecPatchPtrOutput { return v.Spec }).(PersistentVolumeClaimSpecPatchPtrOutput)
}

// status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
func (o PersistentVolumeClaimPatchOutput) Status() PersistentVolumeClaimStatusPatchPtrOutput {
	return o.ApplyT(func(v *PersistentVolumeClaimPatch) PersistentVolumeClaimStatusPatchPtrOutput { return v.Status }).(PersistentVolumeClaimStatusPatchPtrOutput)
}

type PersistentVolumeClaimPatchArrayOutput struct{ *pulumi.OutputState }

func (PersistentVolumeClaimPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistentVolumeClaimPatch)(nil)).Elem()
}

func (o PersistentVolumeClaimPatchArrayOutput) ToPersistentVolumeClaimPatchArrayOutput() PersistentVolumeClaimPatchArrayOutput {
	return o
}

func (o PersistentVolumeClaimPatchArrayOutput) ToPersistentVolumeClaimPatchArrayOutputWithContext(ctx context.Context) PersistentVolumeClaimPatchArrayOutput {
	return o
}

func (o PersistentVolumeClaimPatchArrayOutput) Index(i pulumi.IntInput) PersistentVolumeClaimPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PersistentVolumeClaimPatch {
		return vs[0].([]*PersistentVolumeClaimPatch)[vs[1].(int)]
	}).(PersistentVolumeClaimPatchOutput)
}

type PersistentVolumeClaimPatchMapOutput struct{ *pulumi.OutputState }

func (PersistentVolumeClaimPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistentVolumeClaimPatch)(nil)).Elem()
}

func (o PersistentVolumeClaimPatchMapOutput) ToPersistentVolumeClaimPatchMapOutput() PersistentVolumeClaimPatchMapOutput {
	return o
}

func (o PersistentVolumeClaimPatchMapOutput) ToPersistentVolumeClaimPatchMapOutputWithContext(ctx context.Context) PersistentVolumeClaimPatchMapOutput {
	return o
}

func (o PersistentVolumeClaimPatchMapOutput) MapIndex(k pulumi.StringInput) PersistentVolumeClaimPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PersistentVolumeClaimPatch {
		return vs[0].(map[string]*PersistentVolumeClaimPatch)[vs[1].(string)]
	}).(PersistentVolumeClaimPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PersistentVolumeClaimPatchInput)(nil)).Elem(), &PersistentVolumeClaimPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistentVolumeClaimPatchArrayInput)(nil)).Elem(), PersistentVolumeClaimPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistentVolumeClaimPatchMapInput)(nil)).Elem(), PersistentVolumeClaimPatchMap{})
	pulumi.RegisterOutputType(PersistentVolumeClaimPatchOutput{})
	pulumi.RegisterOutputType(PersistentVolumeClaimPatchArrayOutput{})
	pulumi.RegisterOutputType(PersistentVolumeClaimPatchMapOutput{})
}
