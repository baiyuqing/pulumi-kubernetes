// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.resource.v1alpha2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * ResourceClaimConsumerReference contains enough information to let you locate the consumer of a ResourceClaim. The user must be a resource in the same namespace as the ResourceClaim.
 * 
 */
public final class ResourceClaimConsumerReferenceArgs extends com.pulumi.resources.ResourceArgs {

    public static final ResourceClaimConsumerReferenceArgs Empty = new ResourceClaimConsumerReferenceArgs();

    /**
     * APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
     * 
     */
    @Import(name="apiGroup")
    private @Nullable Output<String> apiGroup;

    /**
     * @return APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
     * 
     */
    public Optional<Output<String>> apiGroup() {
        return Optional.ofNullable(this.apiGroup);
    }

    /**
     * Name is the name of resource being referenced.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name is the name of resource being referenced.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Resource is the type of resource being referenced, for example &#34;pods&#34;.
     * 
     */
    @Import(name="resource", required=true)
    private Output<String> resource;

    /**
     * @return Resource is the type of resource being referenced, for example &#34;pods&#34;.
     * 
     */
    public Output<String> resource() {
        return this.resource;
    }

    /**
     * UID identifies exactly one incarnation of the resource.
     * 
     */
    @Import(name="uid", required=true)
    private Output<String> uid;

    /**
     * @return UID identifies exactly one incarnation of the resource.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }

    private ResourceClaimConsumerReferenceArgs() {}

    private ResourceClaimConsumerReferenceArgs(ResourceClaimConsumerReferenceArgs $) {
        this.apiGroup = $.apiGroup;
        this.name = $.name;
        this.resource = $.resource;
        this.uid = $.uid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResourceClaimConsumerReferenceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResourceClaimConsumerReferenceArgs $;

        public Builder() {
            $ = new ResourceClaimConsumerReferenceArgs();
        }

        public Builder(ResourceClaimConsumerReferenceArgs defaults) {
            $ = new ResourceClaimConsumerReferenceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiGroup APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
         * 
         * @return builder
         * 
         */
        public Builder apiGroup(@Nullable Output<String> apiGroup) {
            $.apiGroup = apiGroup;
            return this;
        }

        /**
         * @param apiGroup APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
         * 
         * @return builder
         * 
         */
        public Builder apiGroup(String apiGroup) {
            return apiGroup(Output.of(apiGroup));
        }

        /**
         * @param name Name is the name of resource being referenced.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name is the name of resource being referenced.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param resource Resource is the type of resource being referenced, for example &#34;pods&#34;.
         * 
         * @return builder
         * 
         */
        public Builder resource(Output<String> resource) {
            $.resource = resource;
            return this;
        }

        /**
         * @param resource Resource is the type of resource being referenced, for example &#34;pods&#34;.
         * 
         * @return builder
         * 
         */
        public Builder resource(String resource) {
            return resource(Output.of(resource));
        }

        /**
         * @param uid UID identifies exactly one incarnation of the resource.
         * 
         * @return builder
         * 
         */
        public Builder uid(Output<String> uid) {
            $.uid = uid;
            return this;
        }

        /**
         * @param uid UID identifies exactly one incarnation of the resource.
         * 
         * @return builder
         * 
         */
        public Builder uid(String uid) {
            return uid(Output.of(uid));
        }

        public ResourceClaimConsumerReferenceArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.resource = Objects.requireNonNull($.resource, "expected parameter 'resource' to be non-null");
            $.uid = Objects.requireNonNull($.uid, "expected parameter 'uid' to be non-null");
            return $;
        }
    }

}
