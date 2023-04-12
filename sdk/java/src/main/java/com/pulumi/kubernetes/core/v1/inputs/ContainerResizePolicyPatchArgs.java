// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * ContainerResizePolicy represents resource resize policy for the container.
 * 
 */
public final class ContainerResizePolicyPatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContainerResizePolicyPatchArgs Empty = new ContainerResizePolicyPatchArgs();

    /**
     * Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.
     * 
     */
    @Import(name="resourceName")
    private @Nullable Output<String> resourceName;

    /**
     * @return Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.
     * 
     */
    public Optional<Output<String>> resourceName() {
        return Optional.ofNullable(this.resourceName);
    }

    /**
     * Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.
     * 
     */
    @Import(name="restartPolicy")
    private @Nullable Output<String> restartPolicy;

    /**
     * @return Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.
     * 
     */
    public Optional<Output<String>> restartPolicy() {
        return Optional.ofNullable(this.restartPolicy);
    }

    private ContainerResizePolicyPatchArgs() {}

    private ContainerResizePolicyPatchArgs(ContainerResizePolicyPatchArgs $) {
        this.resourceName = $.resourceName;
        this.restartPolicy = $.restartPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerResizePolicyPatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerResizePolicyPatchArgs $;

        public Builder() {
            $ = new ContainerResizePolicyPatchArgs();
        }

        public Builder(ContainerResizePolicyPatchArgs defaults) {
            $ = new ContainerResizePolicyPatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param resourceName Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.
         * 
         * @return builder
         * 
         */
        public Builder resourceName(@Nullable Output<String> resourceName) {
            $.resourceName = resourceName;
            return this;
        }

        /**
         * @param resourceName Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.
         * 
         * @return builder
         * 
         */
        public Builder resourceName(String resourceName) {
            return resourceName(Output.of(resourceName));
        }

        /**
         * @param restartPolicy Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.
         * 
         * @return builder
         * 
         */
        public Builder restartPolicy(@Nullable Output<String> restartPolicy) {
            $.restartPolicy = restartPolicy;
            return this;
        }

        /**
         * @param restartPolicy Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.
         * 
         * @return builder
         * 
         */
        public Builder restartPolicy(String restartPolicy) {
            return restartPolicy(Output.of(restartPolicy));
        }

        public ContainerResizePolicyPatchArgs build() {
            return $;
        }
    }

}
