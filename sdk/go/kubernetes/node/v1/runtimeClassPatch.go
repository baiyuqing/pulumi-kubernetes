// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is used to determine which container runtime is used to run all containers in a pod. RuntimeClasses are manually defined by a user or cluster provisioner, and referenced in the PodSpec. The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.  For more details, see https://kubernetes.io/docs/concepts/containers/runtime-class/
type RuntimeClassPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The Handler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
	Handler pulumi.StringPtrOutput `pulumi:"handler"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see
	//  https://kubernetes.io/docs/concepts/scheduling-eviction/pod-overhead/
	Overhead OverheadPatchPtrOutput `pulumi:"overhead"`
	// scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
	Scheduling SchedulingPatchPtrOutput `pulumi:"scheduling"`
}

// NewRuntimeClassPatch registers a new resource with the given unique name, arguments, and options.
func NewRuntimeClassPatch(ctx *pulumi.Context,
	name string, args *RuntimeClassPatchArgs, opts ...pulumi.ResourceOption) (*RuntimeClassPatch, error) {
	if args == nil {
		args = &RuntimeClassPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("node.k8s.io/v1")
	args.Kind = pulumi.StringPtr("RuntimeClass")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:node.k8s.io/v1alpha1:RuntimeClassPatch"),
		},
		{
			Type: pulumi.String("kubernetes:node.k8s.io/v1beta1:RuntimeClassPatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RuntimeClassPatch
	err := ctx.RegisterResource("kubernetes:node.k8s.io/v1:RuntimeClassPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntimeClassPatch gets an existing RuntimeClassPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuntimeClassPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuntimeClassPatchState, opts ...pulumi.ResourceOption) (*RuntimeClassPatch, error) {
	var resource RuntimeClassPatch
	err := ctx.ReadResource("kubernetes:node.k8s.io/v1:RuntimeClassPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuntimeClassPatch resources.
type runtimeClassPatchState struct {
}

type RuntimeClassPatchState struct {
}

func (RuntimeClassPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeClassPatchState)(nil)).Elem()
}

type runtimeClassPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The Handler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
	Handler *string `pulumi:"handler"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see
	//  https://kubernetes.io/docs/concepts/scheduling-eviction/pod-overhead/
	Overhead *OverheadPatch `pulumi:"overhead"`
	// scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
	Scheduling *SchedulingPatch `pulumi:"scheduling"`
}

// The set of arguments for constructing a RuntimeClassPatch resource.
type RuntimeClassPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The Handler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
	Handler pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see
	//  https://kubernetes.io/docs/concepts/scheduling-eviction/pod-overhead/
	Overhead OverheadPatchPtrInput
	// scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
	Scheduling SchedulingPatchPtrInput
}

func (RuntimeClassPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeClassPatchArgs)(nil)).Elem()
}

type RuntimeClassPatchInput interface {
	pulumi.Input

	ToRuntimeClassPatchOutput() RuntimeClassPatchOutput
	ToRuntimeClassPatchOutputWithContext(ctx context.Context) RuntimeClassPatchOutput
}

func (*RuntimeClassPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeClassPatch)(nil)).Elem()
}

func (i *RuntimeClassPatch) ToRuntimeClassPatchOutput() RuntimeClassPatchOutput {
	return i.ToRuntimeClassPatchOutputWithContext(context.Background())
}

func (i *RuntimeClassPatch) ToRuntimeClassPatchOutputWithContext(ctx context.Context) RuntimeClassPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassPatchOutput)
}

func (i *RuntimeClassPatch) ToOutput(ctx context.Context) pulumix.Output[*RuntimeClassPatch] {
	return pulumix.Output[*RuntimeClassPatch]{
		OutputState: i.ToRuntimeClassPatchOutputWithContext(ctx).OutputState,
	}
}

// RuntimeClassPatchArrayInput is an input type that accepts RuntimeClassPatchArray and RuntimeClassPatchArrayOutput values.
// You can construct a concrete instance of `RuntimeClassPatchArrayInput` via:
//
//	RuntimeClassPatchArray{ RuntimeClassPatchArgs{...} }
type RuntimeClassPatchArrayInput interface {
	pulumi.Input

	ToRuntimeClassPatchArrayOutput() RuntimeClassPatchArrayOutput
	ToRuntimeClassPatchArrayOutputWithContext(context.Context) RuntimeClassPatchArrayOutput
}

type RuntimeClassPatchArray []RuntimeClassPatchInput

func (RuntimeClassPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuntimeClassPatch)(nil)).Elem()
}

func (i RuntimeClassPatchArray) ToRuntimeClassPatchArrayOutput() RuntimeClassPatchArrayOutput {
	return i.ToRuntimeClassPatchArrayOutputWithContext(context.Background())
}

func (i RuntimeClassPatchArray) ToRuntimeClassPatchArrayOutputWithContext(ctx context.Context) RuntimeClassPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassPatchArrayOutput)
}

func (i RuntimeClassPatchArray) ToOutput(ctx context.Context) pulumix.Output[[]*RuntimeClassPatch] {
	return pulumix.Output[[]*RuntimeClassPatch]{
		OutputState: i.ToRuntimeClassPatchArrayOutputWithContext(ctx).OutputState,
	}
}

// RuntimeClassPatchMapInput is an input type that accepts RuntimeClassPatchMap and RuntimeClassPatchMapOutput values.
// You can construct a concrete instance of `RuntimeClassPatchMapInput` via:
//
//	RuntimeClassPatchMap{ "key": RuntimeClassPatchArgs{...} }
type RuntimeClassPatchMapInput interface {
	pulumi.Input

	ToRuntimeClassPatchMapOutput() RuntimeClassPatchMapOutput
	ToRuntimeClassPatchMapOutputWithContext(context.Context) RuntimeClassPatchMapOutput
}

type RuntimeClassPatchMap map[string]RuntimeClassPatchInput

func (RuntimeClassPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuntimeClassPatch)(nil)).Elem()
}

func (i RuntimeClassPatchMap) ToRuntimeClassPatchMapOutput() RuntimeClassPatchMapOutput {
	return i.ToRuntimeClassPatchMapOutputWithContext(context.Background())
}

func (i RuntimeClassPatchMap) ToRuntimeClassPatchMapOutputWithContext(ctx context.Context) RuntimeClassPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassPatchMapOutput)
}

func (i RuntimeClassPatchMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RuntimeClassPatch] {
	return pulumix.Output[map[string]*RuntimeClassPatch]{
		OutputState: i.ToRuntimeClassPatchMapOutputWithContext(ctx).OutputState,
	}
}

type RuntimeClassPatchOutput struct{ *pulumi.OutputState }

func (RuntimeClassPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeClassPatch)(nil)).Elem()
}

func (o RuntimeClassPatchOutput) ToRuntimeClassPatchOutput() RuntimeClassPatchOutput {
	return o
}

func (o RuntimeClassPatchOutput) ToRuntimeClassPatchOutputWithContext(ctx context.Context) RuntimeClassPatchOutput {
	return o
}

func (o RuntimeClassPatchOutput) ToOutput(ctx context.Context) pulumix.Output[*RuntimeClassPatch] {
	return pulumix.Output[*RuntimeClassPatch]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o RuntimeClassPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuntimeClassPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The Handler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
func (o RuntimeClassPatchOutput) Handler() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuntimeClassPatch) pulumi.StringPtrOutput { return v.Handler }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o RuntimeClassPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuntimeClassPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o RuntimeClassPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *RuntimeClassPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see
//
//	https://kubernetes.io/docs/concepts/scheduling-eviction/pod-overhead/
func (o RuntimeClassPatchOutput) Overhead() OverheadPatchPtrOutput {
	return o.ApplyT(func(v *RuntimeClassPatch) OverheadPatchPtrOutput { return v.Overhead }).(OverheadPatchPtrOutput)
}

// scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
func (o RuntimeClassPatchOutput) Scheduling() SchedulingPatchPtrOutput {
	return o.ApplyT(func(v *RuntimeClassPatch) SchedulingPatchPtrOutput { return v.Scheduling }).(SchedulingPatchPtrOutput)
}

type RuntimeClassPatchArrayOutput struct{ *pulumi.OutputState }

func (RuntimeClassPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuntimeClassPatch)(nil)).Elem()
}

func (o RuntimeClassPatchArrayOutput) ToRuntimeClassPatchArrayOutput() RuntimeClassPatchArrayOutput {
	return o
}

func (o RuntimeClassPatchArrayOutput) ToRuntimeClassPatchArrayOutputWithContext(ctx context.Context) RuntimeClassPatchArrayOutput {
	return o
}

func (o RuntimeClassPatchArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RuntimeClassPatch] {
	return pulumix.Output[[]*RuntimeClassPatch]{
		OutputState: o.OutputState,
	}
}

func (o RuntimeClassPatchArrayOutput) Index(i pulumi.IntInput) RuntimeClassPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RuntimeClassPatch {
		return vs[0].([]*RuntimeClassPatch)[vs[1].(int)]
	}).(RuntimeClassPatchOutput)
}

type RuntimeClassPatchMapOutput struct{ *pulumi.OutputState }

func (RuntimeClassPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuntimeClassPatch)(nil)).Elem()
}

func (o RuntimeClassPatchMapOutput) ToRuntimeClassPatchMapOutput() RuntimeClassPatchMapOutput {
	return o
}

func (o RuntimeClassPatchMapOutput) ToRuntimeClassPatchMapOutputWithContext(ctx context.Context) RuntimeClassPatchMapOutput {
	return o
}

func (o RuntimeClassPatchMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RuntimeClassPatch] {
	return pulumix.Output[map[string]*RuntimeClassPatch]{
		OutputState: o.OutputState,
	}
}

func (o RuntimeClassPatchMapOutput) MapIndex(k pulumi.StringInput) RuntimeClassPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RuntimeClassPatch {
		return vs[0].(map[string]*RuntimeClassPatch)[vs[1].(string)]
	}).(RuntimeClassPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeClassPatchInput)(nil)).Elem(), &RuntimeClassPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeClassPatchArrayInput)(nil)).Elem(), RuntimeClassPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeClassPatchMapInput)(nil)).Elem(), RuntimeClassPatchMap{})
	pulumi.RegisterOutputType(RuntimeClassPatchOutput{})
	pulumi.RegisterOutputType(RuntimeClassPatchArrayOutput{})
	pulumi.RegisterOutputType(RuntimeClassPatchMapOutput{})
}
