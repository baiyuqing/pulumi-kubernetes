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

// CSINode holds information about all CSI drivers installed on a node. CSI drivers do not need to create the CSINode object directly. As long as they use the node-driver-registrar sidecar container, the kubelet will automatically populate the CSINode object for the CSI driver as part of kubelet plugin registration. CSINode has the same name as a node. If the object is missing, it means either there are no CSI Drivers available on the node, or the Kubelet version is low enough that it doesn't create this object. CSINode has an OwnerReference that points to the corresponding node object.
type CSINode struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. metadata.name must be the Kubernetes node name.
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// spec is the specification of CSINode
	Spec CSINodeSpecOutput `pulumi:"spec"`
}

// NewCSINode registers a new resource with the given unique name, arguments, and options.
func NewCSINode(ctx *pulumi.Context,
	name string, args *CSINodeArgs, opts ...pulumi.ResourceOption) (*CSINode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CSINode")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1beta1:CSINode"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CSINode
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1:CSINode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSINode gets an existing CSINode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSINode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSINodeState, opts ...pulumi.ResourceOption) (*CSINode, error) {
	var resource CSINode
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1:CSINode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSINode resources.
type csinodeState struct {
}

type CSINodeState struct {
}

func (CSINodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*csinodeState)(nil)).Elem()
}

type csinodeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. metadata.name must be the Kubernetes node name.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec is the specification of CSINode
	Spec CSINodeSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CSINode resource.
type CSINodeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. metadata.name must be the Kubernetes node name.
	Metadata metav1.ObjectMetaPtrInput
	// spec is the specification of CSINode
	Spec CSINodeSpecInput
}

func (CSINodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csinodeArgs)(nil)).Elem()
}

type CSINodeInput interface {
	pulumi.Input

	ToCSINodeOutput() CSINodeOutput
	ToCSINodeOutputWithContext(ctx context.Context) CSINodeOutput
}

func (*CSINode) ElementType() reflect.Type {
	return reflect.TypeOf((**CSINode)(nil)).Elem()
}

func (i *CSINode) ToCSINodeOutput() CSINodeOutput {
	return i.ToCSINodeOutputWithContext(context.Background())
}

func (i *CSINode) ToCSINodeOutputWithContext(ctx context.Context) CSINodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodeOutput)
}

func (i *CSINode) ToOutput(ctx context.Context) pulumix.Output[*CSINode] {
	return pulumix.Output[*CSINode]{
		OutputState: i.ToCSINodeOutputWithContext(ctx).OutputState,
	}
}

// CSINodeArrayInput is an input type that accepts CSINodeArray and CSINodeArrayOutput values.
// You can construct a concrete instance of `CSINodeArrayInput` via:
//
//	CSINodeArray{ CSINodeArgs{...} }
type CSINodeArrayInput interface {
	pulumi.Input

	ToCSINodeArrayOutput() CSINodeArrayOutput
	ToCSINodeArrayOutputWithContext(context.Context) CSINodeArrayOutput
}

type CSINodeArray []CSINodeInput

func (CSINodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSINode)(nil)).Elem()
}

func (i CSINodeArray) ToCSINodeArrayOutput() CSINodeArrayOutput {
	return i.ToCSINodeArrayOutputWithContext(context.Background())
}

func (i CSINodeArray) ToCSINodeArrayOutputWithContext(ctx context.Context) CSINodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodeArrayOutput)
}

func (i CSINodeArray) ToOutput(ctx context.Context) pulumix.Output[[]*CSINode] {
	return pulumix.Output[[]*CSINode]{
		OutputState: i.ToCSINodeArrayOutputWithContext(ctx).OutputState,
	}
}

// CSINodeMapInput is an input type that accepts CSINodeMap and CSINodeMapOutput values.
// You can construct a concrete instance of `CSINodeMapInput` via:
//
//	CSINodeMap{ "key": CSINodeArgs{...} }
type CSINodeMapInput interface {
	pulumi.Input

	ToCSINodeMapOutput() CSINodeMapOutput
	ToCSINodeMapOutputWithContext(context.Context) CSINodeMapOutput
}

type CSINodeMap map[string]CSINodeInput

func (CSINodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSINode)(nil)).Elem()
}

func (i CSINodeMap) ToCSINodeMapOutput() CSINodeMapOutput {
	return i.ToCSINodeMapOutputWithContext(context.Background())
}

func (i CSINodeMap) ToCSINodeMapOutputWithContext(ctx context.Context) CSINodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodeMapOutput)
}

func (i CSINodeMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CSINode] {
	return pulumix.Output[map[string]*CSINode]{
		OutputState: i.ToCSINodeMapOutputWithContext(ctx).OutputState,
	}
}

type CSINodeOutput struct{ *pulumi.OutputState }

func (CSINodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CSINode)(nil)).Elem()
}

func (o CSINodeOutput) ToCSINodeOutput() CSINodeOutput {
	return o
}

func (o CSINodeOutput) ToCSINodeOutputWithContext(ctx context.Context) CSINodeOutput {
	return o
}

func (o CSINodeOutput) ToOutput(ctx context.Context) pulumix.Output[*CSINode] {
	return pulumix.Output[*CSINode]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CSINodeOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CSINode) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CSINodeOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CSINode) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. metadata.name must be the Kubernetes node name.
func (o CSINodeOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *CSINode) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// spec is the specification of CSINode
func (o CSINodeOutput) Spec() CSINodeSpecOutput {
	return o.ApplyT(func(v *CSINode) CSINodeSpecOutput { return v.Spec }).(CSINodeSpecOutput)
}

type CSINodeArrayOutput struct{ *pulumi.OutputState }

func (CSINodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSINode)(nil)).Elem()
}

func (o CSINodeArrayOutput) ToCSINodeArrayOutput() CSINodeArrayOutput {
	return o
}

func (o CSINodeArrayOutput) ToCSINodeArrayOutputWithContext(ctx context.Context) CSINodeArrayOutput {
	return o
}

func (o CSINodeArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CSINode] {
	return pulumix.Output[[]*CSINode]{
		OutputState: o.OutputState,
	}
}

func (o CSINodeArrayOutput) Index(i pulumi.IntInput) CSINodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CSINode {
		return vs[0].([]*CSINode)[vs[1].(int)]
	}).(CSINodeOutput)
}

type CSINodeMapOutput struct{ *pulumi.OutputState }

func (CSINodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSINode)(nil)).Elem()
}

func (o CSINodeMapOutput) ToCSINodeMapOutput() CSINodeMapOutput {
	return o
}

func (o CSINodeMapOutput) ToCSINodeMapOutputWithContext(ctx context.Context) CSINodeMapOutput {
	return o
}

func (o CSINodeMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CSINode] {
	return pulumix.Output[map[string]*CSINode]{
		OutputState: o.OutputState,
	}
}

func (o CSINodeMapOutput) MapIndex(k pulumi.StringInput) CSINodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CSINode {
		return vs[0].(map[string]*CSINode)[vs[1].(string)]
	}).(CSINodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CSINodeInput)(nil)).Elem(), &CSINode{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSINodeArrayInput)(nil)).Elem(), CSINodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSINodeMapInput)(nil)).Elem(), CSINodeMap{})
	pulumi.RegisterOutputType(CSINodeOutput{})
	pulumi.RegisterOutputType(CSINodeArrayOutput{})
	pulumi.RegisterOutputType(CSINodeMapOutput{})
}
