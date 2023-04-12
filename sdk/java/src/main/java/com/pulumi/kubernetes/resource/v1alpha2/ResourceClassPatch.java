// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.resource.v1alpha2;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.core.v1.outputs.NodeSelectorPatch;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMetaPatch;
import com.pulumi.kubernetes.resource.v1alpha2.ResourceClassPatchArgs;
import com.pulumi.kubernetes.resource.v1alpha2.outputs.ResourceClassParametersReferencePatch;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Patch resources are used to modify existing Kubernetes resources by using
 * Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
 * one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
 * Conflicts will result in an error by default, but can be forced using the &#34;pulumi.com/patchForce&#34; annotation. See the
 * [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
 * additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
 * ResourceClass is used by administrators to influence how resources are allocated.
 * 
 * This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
 * 
 */
@ResourceType(type="kubernetes:resource.k8s.io/v1alpha2:ResourceClassPatch")
public class ResourceClassPatch extends com.pulumi.resources.CustomResource {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    @Export(name="apiVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> apiVersion;

    /**
     * @return APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    public Output<Optional<String>> apiVersion() {
        return Codegen.optional(this.apiVersion);
    }
    /**
     * DriverName defines the name of the dynamic resource driver that is used for allocation of a ResourceClaim that uses this class.
     * 
     * Resource drivers have a unique name in forward domain order (acme.example.com).
     * 
     */
    @Export(name="driverName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> driverName;

    /**
     * @return DriverName defines the name of the dynamic resource driver that is used for allocation of a ResourceClaim that uses this class.
     * 
     * Resource drivers have a unique name in forward domain order (acme.example.com).
     * 
     */
    public Output<Optional<String>> driverName() {
        return Codegen.optional(this.driverName);
    }
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Export(name="kind", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kind;

    /**
     * @return Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Output<Optional<String>> kind() {
        return Codegen.optional(this.kind);
    }
    /**
     * Standard object metadata
     * 
     */
    @Export(name="metadata", refs={ObjectMetaPatch.class}, tree="[0]")
    private Output</* @Nullable */ ObjectMetaPatch> metadata;

    /**
     * @return Standard object metadata
     * 
     */
    public Output<Optional<ObjectMetaPatch>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * ParametersRef references an arbitrary separate object that may hold parameters that will be used by the driver when allocating a resource that uses this class. A dynamic resource driver can distinguish between parameters stored here and and those stored in ResourceClaimSpec.
     * 
     */
    @Export(name="parametersRef", refs={ResourceClassParametersReferencePatch.class}, tree="[0]")
    private Output</* @Nullable */ ResourceClassParametersReferencePatch> parametersRef;

    /**
     * @return ParametersRef references an arbitrary separate object that may hold parameters that will be used by the driver when allocating a resource that uses this class. A dynamic resource driver can distinguish between parameters stored here and and those stored in ResourceClaimSpec.
     * 
     */
    public Output<Optional<ResourceClassParametersReferencePatch>> parametersRef() {
        return Codegen.optional(this.parametersRef);
    }
    /**
     * Only nodes matching the selector will be considered by the scheduler when trying to find a Node that fits a Pod when that Pod uses a ResourceClaim that has not been allocated yet.
     * 
     * Setting this field is optional. If null, all nodes are candidates.
     * 
     */
    @Export(name="suitableNodes", refs={NodeSelectorPatch.class}, tree="[0]")
    private Output</* @Nullable */ NodeSelectorPatch> suitableNodes;

    /**
     * @return Only nodes matching the selector will be considered by the scheduler when trying to find a Node that fits a Pod when that Pod uses a ResourceClaim that has not been allocated yet.
     * 
     * Setting this field is optional. If null, all nodes are candidates.
     * 
     */
    public Output<Optional<NodeSelectorPatch>> suitableNodes() {
        return Codegen.optional(this.suitableNodes);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ResourceClassPatch(String name) {
        this(name, ResourceClassPatchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ResourceClassPatch(String name, @Nullable ResourceClassPatchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ResourceClassPatch(String name, @Nullable ResourceClassPatchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:resource.k8s.io/v1alpha2:ResourceClassPatch", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private ResourceClassPatch(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:resource.k8s.io/v1alpha2:ResourceClassPatch", name, null, makeResourceOptions(options, id));
    }

    private static ResourceClassPatchArgs makeArgs(@Nullable ResourceClassPatchArgs args) {
        var builder = args == null ? ResourceClassPatchArgs.builder() : ResourceClassPatchArgs.builder(args);
        return builder
            .apiVersion("resource.k8s.io/v1alpha2")
            .kind("ResourceClass")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("kubernetes:resource.k8s.io/v1alpha1:ResourceClassPatch").build())
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ResourceClassPatch get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ResourceClassPatch(name, id, options);
    }
}
