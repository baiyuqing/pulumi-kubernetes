// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/internal"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.
//
// StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
type StorageClass struct {
	pulumi.CustomResourceState

	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion pulumi.BoolOutput `pulumi:"allowVolumeExpansion"`
	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies corev1.TopologySelectorTermArrayOutput `pulumi:"allowedTopologies"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions pulumi.StringArrayOutput `pulumi:"mountOptions"`
	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Provisioner indicates the type of the provisioner.
	Provisioner pulumi.StringOutput `pulumi:"provisioner"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy pulumi.StringOutput `pulumi:"reclaimPolicy"`
	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode pulumi.StringOutput `pulumi:"volumeBindingMode"`
}

// NewStorageClass registers a new resource with the given unique name, arguments, and options.
func NewStorageClass(ctx *pulumi.Context,
	name string, args *StorageClassArgs, opts ...pulumi.ResourceOption) (*StorageClass, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Provisioner == nil {
		return nil, errors.New("invalid value for required argument 'Provisioner'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("StorageClass")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1:StorageClass"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StorageClass
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1beta1:StorageClass", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageClass gets an existing StorageClass resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageClass(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageClassState, opts ...pulumi.ResourceOption) (*StorageClass, error) {
	var resource StorageClass
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1beta1:StorageClass", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageClass resources.
type storageClassState struct {
}

type StorageClassState struct {
}

func (StorageClassState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageClassState)(nil)).Elem()
}

type storageClassArgs struct {
	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion *bool `pulumi:"allowVolumeExpansion"`
	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies []corev1.TopologySelectorTerm `pulumi:"allowedTopologies"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions []string `pulumi:"mountOptions"`
	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters map[string]string `pulumi:"parameters"`
	// Provisioner indicates the type of the provisioner.
	Provisioner string `pulumi:"provisioner"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy *string `pulumi:"reclaimPolicy"`
	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode *string `pulumi:"volumeBindingMode"`
}

// The set of arguments for constructing a StorageClass resource.
type StorageClassArgs struct {
	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion pulumi.BoolPtrInput
	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies corev1.TopologySelectorTermArrayInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions pulumi.StringArrayInput
	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters pulumi.StringMapInput
	// Provisioner indicates the type of the provisioner.
	Provisioner pulumi.StringInput
	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy pulumi.StringPtrInput
	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode pulumi.StringPtrInput
}

func (StorageClassArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageClassArgs)(nil)).Elem()
}

type StorageClassInput interface {
	pulumi.Input

	ToStorageClassOutput() StorageClassOutput
	ToStorageClassOutputWithContext(ctx context.Context) StorageClassOutput
}

func (*StorageClass) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageClass)(nil)).Elem()
}

func (i *StorageClass) ToStorageClassOutput() StorageClassOutput {
	return i.ToStorageClassOutputWithContext(context.Background())
}

func (i *StorageClass) ToStorageClassOutputWithContext(ctx context.Context) StorageClassOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageClassOutput)
}

// StorageClassArrayInput is an input type that accepts StorageClassArray and StorageClassArrayOutput values.
// You can construct a concrete instance of `StorageClassArrayInput` via:
//
//	StorageClassArray{ StorageClassArgs{...} }
type StorageClassArrayInput interface {
	pulumi.Input

	ToStorageClassArrayOutput() StorageClassArrayOutput
	ToStorageClassArrayOutputWithContext(context.Context) StorageClassArrayOutput
}

type StorageClassArray []StorageClassInput

func (StorageClassArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageClass)(nil)).Elem()
}

func (i StorageClassArray) ToStorageClassArrayOutput() StorageClassArrayOutput {
	return i.ToStorageClassArrayOutputWithContext(context.Background())
}

func (i StorageClassArray) ToStorageClassArrayOutputWithContext(ctx context.Context) StorageClassArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageClassArrayOutput)
}

// StorageClassMapInput is an input type that accepts StorageClassMap and StorageClassMapOutput values.
// You can construct a concrete instance of `StorageClassMapInput` via:
//
//	StorageClassMap{ "key": StorageClassArgs{...} }
type StorageClassMapInput interface {
	pulumi.Input

	ToStorageClassMapOutput() StorageClassMapOutput
	ToStorageClassMapOutputWithContext(context.Context) StorageClassMapOutput
}

type StorageClassMap map[string]StorageClassInput

func (StorageClassMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageClass)(nil)).Elem()
}

func (i StorageClassMap) ToStorageClassMapOutput() StorageClassMapOutput {
	return i.ToStorageClassMapOutputWithContext(context.Background())
}

func (i StorageClassMap) ToStorageClassMapOutputWithContext(ctx context.Context) StorageClassMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageClassMapOutput)
}

type StorageClassOutput struct{ *pulumi.OutputState }

func (StorageClassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageClass)(nil)).Elem()
}

func (o StorageClassOutput) ToStorageClassOutput() StorageClassOutput {
	return o
}

func (o StorageClassOutput) ToStorageClassOutputWithContext(ctx context.Context) StorageClassOutput {
	return o
}

// AllowVolumeExpansion shows whether the storage class allow volume expand
func (o StorageClassOutput) AllowVolumeExpansion() pulumi.BoolOutput {
	return o.ApplyT(func(v *StorageClass) pulumi.BoolOutput { return v.AllowVolumeExpansion }).(pulumi.BoolOutput)
}

// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
func (o StorageClassOutput) AllowedTopologies() corev1.TopologySelectorTermArrayOutput {
	return o.ApplyT(func(v *StorageClass) corev1.TopologySelectorTermArrayOutput { return v.AllowedTopologies }).(corev1.TopologySelectorTermArrayOutput)
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o StorageClassOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageClass) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o StorageClassOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageClass) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o StorageClassOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *StorageClass) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
func (o StorageClassOutput) MountOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StorageClass) pulumi.StringArrayOutput { return v.MountOptions }).(pulumi.StringArrayOutput)
}

// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
func (o StorageClassOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageClass) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

// Provisioner indicates the type of the provisioner.
func (o StorageClassOutput) Provisioner() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageClass) pulumi.StringOutput { return v.Provisioner }).(pulumi.StringOutput)
}

// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
func (o StorageClassOutput) ReclaimPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageClass) pulumi.StringOutput { return v.ReclaimPolicy }).(pulumi.StringOutput)
}

// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
func (o StorageClassOutput) VolumeBindingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageClass) pulumi.StringOutput { return v.VolumeBindingMode }).(pulumi.StringOutput)
}

type StorageClassArrayOutput struct{ *pulumi.OutputState }

func (StorageClassArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageClass)(nil)).Elem()
}

func (o StorageClassArrayOutput) ToStorageClassArrayOutput() StorageClassArrayOutput {
	return o
}

func (o StorageClassArrayOutput) ToStorageClassArrayOutputWithContext(ctx context.Context) StorageClassArrayOutput {
	return o
}

func (o StorageClassArrayOutput) Index(i pulumi.IntInput) StorageClassOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StorageClass {
		return vs[0].([]*StorageClass)[vs[1].(int)]
	}).(StorageClassOutput)
}

type StorageClassMapOutput struct{ *pulumi.OutputState }

func (StorageClassMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageClass)(nil)).Elem()
}

func (o StorageClassMapOutput) ToStorageClassMapOutput() StorageClassMapOutput {
	return o
}

func (o StorageClassMapOutput) ToStorageClassMapOutputWithContext(ctx context.Context) StorageClassMapOutput {
	return o
}

func (o StorageClassMapOutput) MapIndex(k pulumi.StringInput) StorageClassOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StorageClass {
		return vs[0].(map[string]*StorageClass)[vs[1].(string)]
	}).(StorageClassOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageClassInput)(nil)).Elem(), &StorageClass{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageClassArrayInput)(nil)).Elem(), StorageClassArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageClassMapInput)(nil)).Elem(), StorageClassMap{})
	pulumi.RegisterOutputType(StorageClassOutput{})
	pulumi.RegisterOutputType(StorageClassArrayOutput{})
	pulumi.RegisterOutputType(StorageClassMapOutput{})
}
