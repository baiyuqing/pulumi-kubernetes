// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.scheduling.v1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.meta.v1.inputs.ObjectMetaPatchArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PriorityClassPatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final PriorityClassPatchArgs Empty = new PriorityClassPatchArgs();

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    @Import(name="apiVersion")
    private @Nullable Output<String> apiVersion;

    /**
     * @return APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    public Optional<Output<String>> apiVersion() {
        return Optional.ofNullable(this.apiVersion);
    }

    /**
     * description is an arbitrary string that usually provides guidelines on when this priority class should be used.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return description is an arbitrary string that usually provides guidelines on when this priority class should be used.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
     * 
     */
    @Import(name="globalDefault")
    private @Nullable Output<Boolean> globalDefault;

    /**
     * @return globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
     * 
     */
    public Optional<Output<Boolean>> globalDefault() {
        return Optional.ofNullable(this.globalDefault);
    }

    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Import(name="kind")
    private @Nullable Output<String> kind;

    /**
     * @return Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Optional<Output<String>> kind() {
        return Optional.ofNullable(this.kind);
    }

    /**
     * Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<ObjectMetaPatchArgs> metadata;

    /**
     * @return Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Optional<Output<ObjectMetaPatchArgs>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
     * 
     */
    @Import(name="preemptionPolicy")
    private @Nullable Output<String> preemptionPolicy;

    /**
     * @return preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
     * 
     */
    public Optional<Output<String>> preemptionPolicy() {
        return Optional.ofNullable(this.preemptionPolicy);
    }

    /**
     * value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
     * 
     */
    @Import(name="value")
    private @Nullable Output<Integer> value;

    /**
     * @return value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
     * 
     */
    public Optional<Output<Integer>> value() {
        return Optional.ofNullable(this.value);
    }

    private PriorityClassPatchArgs() {}

    private PriorityClassPatchArgs(PriorityClassPatchArgs $) {
        this.apiVersion = $.apiVersion;
        this.description = $.description;
        this.globalDefault = $.globalDefault;
        this.kind = $.kind;
        this.metadata = $.metadata;
        this.preemptionPolicy = $.preemptionPolicy;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PriorityClassPatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PriorityClassPatchArgs $;

        public Builder() {
            $ = new PriorityClassPatchArgs();
        }

        public Builder(PriorityClassPatchArgs defaults) {
            $ = new PriorityClassPatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiVersion APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
         * 
         * @return builder
         * 
         */
        public Builder apiVersion(@Nullable Output<String> apiVersion) {
            $.apiVersion = apiVersion;
            return this;
        }

        /**
         * @param apiVersion APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
         * 
         * @return builder
         * 
         */
        public Builder apiVersion(String apiVersion) {
            return apiVersion(Output.of(apiVersion));
        }

        /**
         * @param description description is an arbitrary string that usually provides guidelines on when this priority class should be used.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description description is an arbitrary string that usually provides guidelines on when this priority class should be used.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param globalDefault globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
         * 
         * @return builder
         * 
         */
        public Builder globalDefault(@Nullable Output<Boolean> globalDefault) {
            $.globalDefault = globalDefault;
            return this;
        }

        /**
         * @param globalDefault globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
         * 
         * @return builder
         * 
         */
        public Builder globalDefault(Boolean globalDefault) {
            return globalDefault(Output.of(globalDefault));
        }

        /**
         * @param kind Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
         * 
         * @return builder
         * 
         */
        public Builder kind(@Nullable Output<String> kind) {
            $.kind = kind;
            return this;
        }

        /**
         * @param kind Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
         * 
         * @return builder
         * 
         */
        public Builder kind(String kind) {
            return kind(Output.of(kind));
        }

        /**
         * @param metadata Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<ObjectMetaPatchArgs> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
         * 
         * @return builder
         * 
         */
        public Builder metadata(ObjectMetaPatchArgs metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param preemptionPolicy preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
         * 
         * @return builder
         * 
         */
        public Builder preemptionPolicy(@Nullable Output<String> preemptionPolicy) {
            $.preemptionPolicy = preemptionPolicy;
            return this;
        }

        /**
         * @param preemptionPolicy preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
         * 
         * @return builder
         * 
         */
        public Builder preemptionPolicy(String preemptionPolicy) {
            return preemptionPolicy(Output.of(preemptionPolicy));
        }

        /**
         * @param value value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<Integer> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
         * 
         * @return builder
         * 
         */
        public Builder value(Integer value) {
            return value(Output.of(value));
        }

        public PriorityClassPatchArgs build() {
            $.apiVersion = Codegen.stringProp("apiVersion").output().arg($.apiVersion).getNullable();
            $.kind = Codegen.stringProp("kind").output().arg($.kind).getNullable();
            return $;
        }
    }

}
