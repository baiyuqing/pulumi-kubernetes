// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.resource.v1alpha2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.resource.v1alpha2.inputs.ResourceClaimParametersReferencePatchArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * ResourceClaimSpec defines how a resource is to be allocated.
 * 
 */
public final class ResourceClaimSpecPatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final ResourceClaimSpecPatchArgs Empty = new ResourceClaimSpecPatchArgs();

    /**
     * Allocation can start immediately or when a Pod wants to use the resource. &#34;WaitForFirstConsumer&#34; is the default.
     * 
     */
    @Import(name="allocationMode")
    private @Nullable Output<String> allocationMode;

    /**
     * @return Allocation can start immediately or when a Pod wants to use the resource. &#34;WaitForFirstConsumer&#34; is the default.
     * 
     */
    public Optional<Output<String>> allocationMode() {
        return Optional.ofNullable(this.allocationMode);
    }

    /**
     * ParametersRef references a separate object with arbitrary parameters that will be used by the driver when allocating a resource for the claim.
     * 
     * The object must be in the same namespace as the ResourceClaim.
     * 
     */
    @Import(name="parametersRef")
    private @Nullable Output<ResourceClaimParametersReferencePatchArgs> parametersRef;

    /**
     * @return ParametersRef references a separate object with arbitrary parameters that will be used by the driver when allocating a resource for the claim.
     * 
     * The object must be in the same namespace as the ResourceClaim.
     * 
     */
    public Optional<Output<ResourceClaimParametersReferencePatchArgs>> parametersRef() {
        return Optional.ofNullable(this.parametersRef);
    }

    /**
     * ResourceClassName references the driver and additional parameters via the name of a ResourceClass that was created as part of the driver deployment.
     * 
     */
    @Import(name="resourceClassName")
    private @Nullable Output<String> resourceClassName;

    /**
     * @return ResourceClassName references the driver and additional parameters via the name of a ResourceClass that was created as part of the driver deployment.
     * 
     */
    public Optional<Output<String>> resourceClassName() {
        return Optional.ofNullable(this.resourceClassName);
    }

    private ResourceClaimSpecPatchArgs() {}

    private ResourceClaimSpecPatchArgs(ResourceClaimSpecPatchArgs $) {
        this.allocationMode = $.allocationMode;
        this.parametersRef = $.parametersRef;
        this.resourceClassName = $.resourceClassName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResourceClaimSpecPatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResourceClaimSpecPatchArgs $;

        public Builder() {
            $ = new ResourceClaimSpecPatchArgs();
        }

        public Builder(ResourceClaimSpecPatchArgs defaults) {
            $ = new ResourceClaimSpecPatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allocationMode Allocation can start immediately or when a Pod wants to use the resource. &#34;WaitForFirstConsumer&#34; is the default.
         * 
         * @return builder
         * 
         */
        public Builder allocationMode(@Nullable Output<String> allocationMode) {
            $.allocationMode = allocationMode;
            return this;
        }

        /**
         * @param allocationMode Allocation can start immediately or when a Pod wants to use the resource. &#34;WaitForFirstConsumer&#34; is the default.
         * 
         * @return builder
         * 
         */
        public Builder allocationMode(String allocationMode) {
            return allocationMode(Output.of(allocationMode));
        }

        /**
         * @param parametersRef ParametersRef references a separate object with arbitrary parameters that will be used by the driver when allocating a resource for the claim.
         * 
         * The object must be in the same namespace as the ResourceClaim.
         * 
         * @return builder
         * 
         */
        public Builder parametersRef(@Nullable Output<ResourceClaimParametersReferencePatchArgs> parametersRef) {
            $.parametersRef = parametersRef;
            return this;
        }

        /**
         * @param parametersRef ParametersRef references a separate object with arbitrary parameters that will be used by the driver when allocating a resource for the claim.
         * 
         * The object must be in the same namespace as the ResourceClaim.
         * 
         * @return builder
         * 
         */
        public Builder parametersRef(ResourceClaimParametersReferencePatchArgs parametersRef) {
            return parametersRef(Output.of(parametersRef));
        }

        /**
         * @param resourceClassName ResourceClassName references the driver and additional parameters via the name of a ResourceClass that was created as part of the driver deployment.
         * 
         * @return builder
         * 
         */
        public Builder resourceClassName(@Nullable Output<String> resourceClassName) {
            $.resourceClassName = resourceClassName;
            return this;
        }

        /**
         * @param resourceClassName ResourceClassName references the driver and additional parameters via the name of a ResourceClass that was created as part of the driver deployment.
         * 
         * @return builder
         * 
         */
        public Builder resourceClassName(String resourceClassName) {
            return resourceClassName(Output.of(resourceClassName));
        }

        public ResourceClaimSpecPatchArgs build() {
            return $;
        }
    }

}
