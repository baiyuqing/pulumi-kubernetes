// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.networking.v1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.meta.v1.outputs.ListMeta;
import com.pulumi.kubernetes.networking.v1.NetworkPolicyListArgs;
import com.pulumi.kubernetes.networking.v1.outputs.NetworkPolicy;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * NetworkPolicyList is a list of NetworkPolicy objects.
 * 
 */
@ResourceType(type="kubernetes:networking.k8s.io/v1:NetworkPolicyList")
public class NetworkPolicyList extends com.pulumi.resources.CustomResource {
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
     * items is a list of schema objects.
     * 
     */
    @Export(name="items", refs={List.class,NetworkPolicy.class}, tree="[0,1]")
    private Output<List<NetworkPolicy>> items;

    /**
     * @return items is a list of schema objects.
     * 
     */
    public Output<List<NetworkPolicy>> items() {
        return this.items;
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
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Export(name="metadata", refs={ListMeta.class}, tree="[0]")
    private Output</* @Nullable */ ListMeta> metadata;

    /**
     * @return Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Output<Optional<ListMeta>> metadata() {
        return Codegen.optional(this.metadata);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkPolicyList(String name) {
        this(name, NetworkPolicyListArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkPolicyList(String name, NetworkPolicyListArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkPolicyList(String name, NetworkPolicyListArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:networking.k8s.io/v1:NetworkPolicyList", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private NetworkPolicyList(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:networking.k8s.io/v1:NetworkPolicyList", name, null, makeResourceOptions(options, id));
    }

    private static NetworkPolicyListArgs makeArgs(NetworkPolicyListArgs args) {
        var builder = args == null ? NetworkPolicyListArgs.builder() : NetworkPolicyListArgs.builder(args);
        return builder
            .apiVersion("networking.k8s.io/v1")
            .kind("NetworkPolicyList")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static NetworkPolicyList get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkPolicyList(name, id, options);
    }
}
