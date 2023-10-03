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
// EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.
type EndpointSlicePatch struct {
	pulumi.CustomResourceState

	// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
	AddressType pulumi.StringPtrOutput `pulumi:"addressType"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
	Endpoints EndpointPatchArrayOutput `pulumi:"endpoints"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
	Ports EndpointPortPatchArrayOutput `pulumi:"ports"`
}

// NewEndpointSlicePatch registers a new resource with the given unique name, arguments, and options.
func NewEndpointSlicePatch(ctx *pulumi.Context,
	name string, args *EndpointSlicePatchArgs, opts ...pulumi.ResourceOption) (*EndpointSlicePatch, error) {
	if args == nil {
		args = &EndpointSlicePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("discovery.k8s.io/v1")
	args.Kind = pulumi.StringPtr("EndpointSlice")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:discovery.k8s.io/v1beta1:EndpointSlicePatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EndpointSlicePatch
	err := ctx.RegisterResource("kubernetes:discovery.k8s.io/v1:EndpointSlicePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointSlicePatch gets an existing EndpointSlicePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointSlicePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointSlicePatchState, opts ...pulumi.ResourceOption) (*EndpointSlicePatch, error) {
	var resource EndpointSlicePatch
	err := ctx.ReadResource("kubernetes:discovery.k8s.io/v1:EndpointSlicePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointSlicePatch resources.
type endpointSlicePatchState struct {
}

type EndpointSlicePatchState struct {
}

func (EndpointSlicePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointSlicePatchState)(nil)).Elem()
}

type endpointSlicePatchArgs struct {
	// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
	AddressType *string `pulumi:"addressType"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
	Endpoints []EndpointPatch `pulumi:"endpoints"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
	Ports []EndpointPortPatch `pulumi:"ports"`
}

// The set of arguments for constructing a EndpointSlicePatch resource.
type EndpointSlicePatchArgs struct {
	// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
	AddressType pulumi.StringPtrInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
	Endpoints EndpointPatchArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPatchPtrInput
	// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
	Ports EndpointPortPatchArrayInput
}

func (EndpointSlicePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointSlicePatchArgs)(nil)).Elem()
}

type EndpointSlicePatchInput interface {
	pulumi.Input

	ToEndpointSlicePatchOutput() EndpointSlicePatchOutput
	ToEndpointSlicePatchOutputWithContext(ctx context.Context) EndpointSlicePatchOutput
}

func (*EndpointSlicePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointSlicePatch)(nil)).Elem()
}

func (i *EndpointSlicePatch) ToEndpointSlicePatchOutput() EndpointSlicePatchOutput {
	return i.ToEndpointSlicePatchOutputWithContext(context.Background())
}

func (i *EndpointSlicePatch) ToEndpointSlicePatchOutputWithContext(ctx context.Context) EndpointSlicePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointSlicePatchOutput)
}

func (i *EndpointSlicePatch) ToOutput(ctx context.Context) pulumix.Output[*EndpointSlicePatch] {
	return pulumix.Output[*EndpointSlicePatch]{
		OutputState: i.ToEndpointSlicePatchOutputWithContext(ctx).OutputState,
	}
}

// EndpointSlicePatchArrayInput is an input type that accepts EndpointSlicePatchArray and EndpointSlicePatchArrayOutput values.
// You can construct a concrete instance of `EndpointSlicePatchArrayInput` via:
//
//	EndpointSlicePatchArray{ EndpointSlicePatchArgs{...} }
type EndpointSlicePatchArrayInput interface {
	pulumi.Input

	ToEndpointSlicePatchArrayOutput() EndpointSlicePatchArrayOutput
	ToEndpointSlicePatchArrayOutputWithContext(context.Context) EndpointSlicePatchArrayOutput
}

type EndpointSlicePatchArray []EndpointSlicePatchInput

func (EndpointSlicePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointSlicePatch)(nil)).Elem()
}

func (i EndpointSlicePatchArray) ToEndpointSlicePatchArrayOutput() EndpointSlicePatchArrayOutput {
	return i.ToEndpointSlicePatchArrayOutputWithContext(context.Background())
}

func (i EndpointSlicePatchArray) ToEndpointSlicePatchArrayOutputWithContext(ctx context.Context) EndpointSlicePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointSlicePatchArrayOutput)
}

func (i EndpointSlicePatchArray) ToOutput(ctx context.Context) pulumix.Output[[]*EndpointSlicePatch] {
	return pulumix.Output[[]*EndpointSlicePatch]{
		OutputState: i.ToEndpointSlicePatchArrayOutputWithContext(ctx).OutputState,
	}
}

// EndpointSlicePatchMapInput is an input type that accepts EndpointSlicePatchMap and EndpointSlicePatchMapOutput values.
// You can construct a concrete instance of `EndpointSlicePatchMapInput` via:
//
//	EndpointSlicePatchMap{ "key": EndpointSlicePatchArgs{...} }
type EndpointSlicePatchMapInput interface {
	pulumi.Input

	ToEndpointSlicePatchMapOutput() EndpointSlicePatchMapOutput
	ToEndpointSlicePatchMapOutputWithContext(context.Context) EndpointSlicePatchMapOutput
}

type EndpointSlicePatchMap map[string]EndpointSlicePatchInput

func (EndpointSlicePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointSlicePatch)(nil)).Elem()
}

func (i EndpointSlicePatchMap) ToEndpointSlicePatchMapOutput() EndpointSlicePatchMapOutput {
	return i.ToEndpointSlicePatchMapOutputWithContext(context.Background())
}

func (i EndpointSlicePatchMap) ToEndpointSlicePatchMapOutputWithContext(ctx context.Context) EndpointSlicePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointSlicePatchMapOutput)
}

func (i EndpointSlicePatchMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*EndpointSlicePatch] {
	return pulumix.Output[map[string]*EndpointSlicePatch]{
		OutputState: i.ToEndpointSlicePatchMapOutputWithContext(ctx).OutputState,
	}
}

type EndpointSlicePatchOutput struct{ *pulumi.OutputState }

func (EndpointSlicePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointSlicePatch)(nil)).Elem()
}

func (o EndpointSlicePatchOutput) ToEndpointSlicePatchOutput() EndpointSlicePatchOutput {
	return o
}

func (o EndpointSlicePatchOutput) ToEndpointSlicePatchOutputWithContext(ctx context.Context) EndpointSlicePatchOutput {
	return o
}

func (o EndpointSlicePatchOutput) ToOutput(ctx context.Context) pulumix.Output[*EndpointSlicePatch] {
	return pulumix.Output[*EndpointSlicePatch]{
		OutputState: o.OutputState,
	}
}

// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
func (o EndpointSlicePatchOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointSlicePatch) pulumi.StringPtrOutput { return v.AddressType }).(pulumi.StringPtrOutput)
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o EndpointSlicePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointSlicePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
func (o EndpointSlicePatchOutput) Endpoints() EndpointPatchArrayOutput {
	return o.ApplyT(func(v *EndpointSlicePatch) EndpointPatchArrayOutput { return v.Endpoints }).(EndpointPatchArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o EndpointSlicePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointSlicePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata.
func (o EndpointSlicePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *EndpointSlicePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
func (o EndpointSlicePatchOutput) Ports() EndpointPortPatchArrayOutput {
	return o.ApplyT(func(v *EndpointSlicePatch) EndpointPortPatchArrayOutput { return v.Ports }).(EndpointPortPatchArrayOutput)
}

type EndpointSlicePatchArrayOutput struct{ *pulumi.OutputState }

func (EndpointSlicePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointSlicePatch)(nil)).Elem()
}

func (o EndpointSlicePatchArrayOutput) ToEndpointSlicePatchArrayOutput() EndpointSlicePatchArrayOutput {
	return o
}

func (o EndpointSlicePatchArrayOutput) ToEndpointSlicePatchArrayOutputWithContext(ctx context.Context) EndpointSlicePatchArrayOutput {
	return o
}

func (o EndpointSlicePatchArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*EndpointSlicePatch] {
	return pulumix.Output[[]*EndpointSlicePatch]{
		OutputState: o.OutputState,
	}
}

func (o EndpointSlicePatchArrayOutput) Index(i pulumi.IntInput) EndpointSlicePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointSlicePatch {
		return vs[0].([]*EndpointSlicePatch)[vs[1].(int)]
	}).(EndpointSlicePatchOutput)
}

type EndpointSlicePatchMapOutput struct{ *pulumi.OutputState }

func (EndpointSlicePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointSlicePatch)(nil)).Elem()
}

func (o EndpointSlicePatchMapOutput) ToEndpointSlicePatchMapOutput() EndpointSlicePatchMapOutput {
	return o
}

func (o EndpointSlicePatchMapOutput) ToEndpointSlicePatchMapOutputWithContext(ctx context.Context) EndpointSlicePatchMapOutput {
	return o
}

func (o EndpointSlicePatchMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*EndpointSlicePatch] {
	return pulumix.Output[map[string]*EndpointSlicePatch]{
		OutputState: o.OutputState,
	}
}

func (o EndpointSlicePatchMapOutput) MapIndex(k pulumi.StringInput) EndpointSlicePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointSlicePatch {
		return vs[0].(map[string]*EndpointSlicePatch)[vs[1].(string)]
	}).(EndpointSlicePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointSlicePatchInput)(nil)).Elem(), &EndpointSlicePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointSlicePatchArrayInput)(nil)).Elem(), EndpointSlicePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointSlicePatchMapInput)(nil)).Elem(), EndpointSlicePatchMap{})
	pulumi.RegisterOutputType(EndpointSlicePatchOutput{})
	pulumi.RegisterOutputType(EndpointSlicePatchArrayOutput{})
	pulumi.RegisterOutputType(EndpointSlicePatchMapOutput{})
}
