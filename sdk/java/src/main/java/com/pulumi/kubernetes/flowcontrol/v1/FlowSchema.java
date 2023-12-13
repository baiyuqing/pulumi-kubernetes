// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.flowcontrol.v1;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.flowcontrol.v1.FlowSchemaArgs;
import com.pulumi.kubernetes.flowcontrol.v1.outputs.FlowSchemaSpec;
import com.pulumi.kubernetes.flowcontrol.v1.outputs.FlowSchemaStatus;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMeta;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a &#34;flow distinguisher&#34;.
 * 
 */
@ResourceType(type="kubernetes:flowcontrol.apiserver.k8s.io/v1:FlowSchema")
public class FlowSchema extends com.pulumi.resources.CustomResource {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    @Export(name="apiVersion", refs={String.class}, tree="[0]")
    private Output<String> apiVersion;

    /**
     * @return APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    public Output<String> apiVersion() {
        return this.apiVersion;
    }
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Export(name="kind", refs={String.class}, tree="[0]")
    private Output<String> kind;

    /**
     * @return Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Output<String> kind() {
        return this.kind;
    }
    /**
     * `metadata` is the standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Export(name="metadata", refs={ObjectMeta.class}, tree="[0]")
    private Output<ObjectMeta> metadata;

    /**
     * @return `metadata` is the standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Output<ObjectMeta> metadata() {
        return this.metadata;
    }
    /**
     * `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    @Export(name="spec", refs={FlowSchemaSpec.class}, tree="[0]")
    private Output<FlowSchemaSpec> spec;

    /**
     * @return `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    public Output<FlowSchemaSpec> spec() {
        return this.spec;
    }
    /**
     * `status` is the current status of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    @Export(name="status", refs={FlowSchemaStatus.class}, tree="[0]")
    private Output</* @Nullable */ FlowSchemaStatus> status;

    /**
     * @return `status` is the current status of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    public Output<Optional<FlowSchemaStatus>> status() {
        return Codegen.optional(this.status);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FlowSchema(String name) {
        this(name, FlowSchemaArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FlowSchema(String name, @Nullable FlowSchemaArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FlowSchema(String name, @Nullable FlowSchemaArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:flowcontrol.apiserver.k8s.io/v1:FlowSchema", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private FlowSchema(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:flowcontrol.apiserver.k8s.io/v1:FlowSchema", name, null, makeResourceOptions(options, id));
    }

    private static FlowSchemaArgs makeArgs(@Nullable FlowSchemaArgs args) {
        var builder = args == null ? FlowSchemaArgs.builder() : FlowSchemaArgs.builder(args);
        return builder
            .apiVersion("flowcontrol.apiserver.k8s.io/v1")
            .kind("FlowSchema")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchema").build()),
                Output.of(Alias.builder().type("kubernetes:flowcontrol.apiserver.k8s.io/v1beta1:FlowSchema").build()),
                Output.of(Alias.builder().type("kubernetes:flowcontrol.apiserver.k8s.io/v1beta2:FlowSchema").build()),
                Output.of(Alias.builder().type("kubernetes:flowcontrol.apiserver.k8s.io/v1beta3:FlowSchema").build())
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
    public static FlowSchema get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FlowSchema(name, id, options);
    }
}
