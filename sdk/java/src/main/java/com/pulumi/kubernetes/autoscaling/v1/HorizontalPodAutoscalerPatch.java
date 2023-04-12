// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.autoscaling.v1;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.autoscaling.v1.HorizontalPodAutoscalerPatchArgs;
import com.pulumi.kubernetes.autoscaling.v1.outputs.HorizontalPodAutoscalerSpecPatch;
import com.pulumi.kubernetes.autoscaling.v1.outputs.HorizontalPodAutoscalerStatusPatch;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMetaPatch;
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
 * configuration of a horizontal pod autoscaler.
 * 
 */
@ResourceType(type="kubernetes:autoscaling/v1:HorizontalPodAutoscalerPatch")
public class HorizontalPodAutoscalerPatch extends com.pulumi.resources.CustomResource {
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
     * Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Export(name="metadata", refs={ObjectMetaPatch.class}, tree="[0]")
    private Output</* @Nullable */ ObjectMetaPatch> metadata;

    /**
     * @return Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Output<Optional<ObjectMetaPatch>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * spec defines the behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
     * 
     */
    @Export(name="spec", refs={HorizontalPodAutoscalerSpecPatch.class}, tree="[0]")
    private Output</* @Nullable */ HorizontalPodAutoscalerSpecPatch> spec;

    /**
     * @return spec defines the behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
     * 
     */
    public Output<Optional<HorizontalPodAutoscalerSpecPatch>> spec() {
        return Codegen.optional(this.spec);
    }
    /**
     * status is the current information about the autoscaler.
     * 
     */
    @Export(name="status", refs={HorizontalPodAutoscalerStatusPatch.class}, tree="[0]")
    private Output</* @Nullable */ HorizontalPodAutoscalerStatusPatch> status;

    /**
     * @return status is the current information about the autoscaler.
     * 
     */
    public Output<Optional<HorizontalPodAutoscalerStatusPatch>> status() {
        return Codegen.optional(this.status);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HorizontalPodAutoscalerPatch(String name) {
        this(name, HorizontalPodAutoscalerPatchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HorizontalPodAutoscalerPatch(String name, @Nullable HorizontalPodAutoscalerPatchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HorizontalPodAutoscalerPatch(String name, @Nullable HorizontalPodAutoscalerPatchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:autoscaling/v1:HorizontalPodAutoscalerPatch", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private HorizontalPodAutoscalerPatch(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:autoscaling/v1:HorizontalPodAutoscalerPatch", name, null, makeResourceOptions(options, id));
    }

    private static HorizontalPodAutoscalerPatchArgs makeArgs(@Nullable HorizontalPodAutoscalerPatchArgs args) {
        var builder = args == null ? HorizontalPodAutoscalerPatchArgs.builder() : HorizontalPodAutoscalerPatchArgs.builder(args);
        return builder
            .apiVersion("autoscaling/v1")
            .kind("HorizontalPodAutoscaler")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("kubernetes:autoscaling/v2:HorizontalPodAutoscalerPatch").build()),
                Output.of(Alias.builder().type("kubernetes:autoscaling/v2beta1:HorizontalPodAutoscalerPatch").build()),
                Output.of(Alias.builder().type("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscalerPatch").build())
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
    public static HorizontalPodAutoscalerPatch get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HorizontalPodAutoscalerPatch(name, id, options);
    }
}
