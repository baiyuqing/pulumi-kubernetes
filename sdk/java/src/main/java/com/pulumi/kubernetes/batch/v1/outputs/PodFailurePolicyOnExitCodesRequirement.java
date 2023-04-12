// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.batch.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PodFailurePolicyOnExitCodesRequirement {
    /**
     * @return Restricts the check for exit codes to the container with the specified name. When null, the rule applies to all containers. When specified, it should match one the container or initContainer names in the pod template.
     * 
     */
    private @Nullable String containerName;
    /**
     * @return Represents the relationship between the container exit code(s) and the specified values. Containers completed with success (exit code 0) are excluded from the requirement check. Possible values are:
     * 
     * - In: the requirement is satisfied if at least one container exit code
     *   (might be multiple if there are multiple containers not restricted
     *   by the &#39;containerName&#39; field) is in the set of specified values.
     * - NotIn: the requirement is satisfied if at least one container exit code
     *   (might be multiple if there are multiple containers not restricted
     *   by the &#39;containerName&#39; field) is not in the set of specified values.
     *   Additional values are considered to be added in the future. Clients should react to an unknown operator by assuming the requirement is not satisfied.
     * 
     */
    private String operator;
    /**
     * @return Specifies the set of values. Each returned container exit code (might be multiple in case of multiple containers) is checked against this set of values with respect to the operator. The list of values must be ordered and must not contain duplicates. Value &#39;0&#39; cannot be used for the In operator. At least one element is required. At most 255 elements are allowed.
     * 
     */
    private List<Integer> values;

    private PodFailurePolicyOnExitCodesRequirement() {}
    /**
     * @return Restricts the check for exit codes to the container with the specified name. When null, the rule applies to all containers. When specified, it should match one the container or initContainer names in the pod template.
     * 
     */
    public Optional<String> containerName() {
        return Optional.ofNullable(this.containerName);
    }
    /**
     * @return Represents the relationship between the container exit code(s) and the specified values. Containers completed with success (exit code 0) are excluded from the requirement check. Possible values are:
     * 
     * - In: the requirement is satisfied if at least one container exit code
     *   (might be multiple if there are multiple containers not restricted
     *   by the &#39;containerName&#39; field) is in the set of specified values.
     * - NotIn: the requirement is satisfied if at least one container exit code
     *   (might be multiple if there are multiple containers not restricted
     *   by the &#39;containerName&#39; field) is not in the set of specified values.
     *   Additional values are considered to be added in the future. Clients should react to an unknown operator by assuming the requirement is not satisfied.
     * 
     */
    public String operator() {
        return this.operator;
    }
    /**
     * @return Specifies the set of values. Each returned container exit code (might be multiple in case of multiple containers) is checked against this set of values with respect to the operator. The list of values must be ordered and must not contain duplicates. Value &#39;0&#39; cannot be used for the In operator. At least one element is required. At most 255 elements are allowed.
     * 
     */
    public List<Integer> values() {
        return this.values;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PodFailurePolicyOnExitCodesRequirement defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String containerName;
        private String operator;
        private List<Integer> values;
        public Builder() {}
        public Builder(PodFailurePolicyOnExitCodesRequirement defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.containerName = defaults.containerName;
    	      this.operator = defaults.operator;
    	      this.values = defaults.values;
        }

        @CustomType.Setter
        public Builder containerName(@Nullable String containerName) {
            this.containerName = containerName;
            return this;
        }
        @CustomType.Setter
        public Builder operator(String operator) {
            this.operator = Objects.requireNonNull(operator);
            return this;
        }
        @CustomType.Setter
        public Builder values(List<Integer> values) {
            this.values = Objects.requireNonNull(values);
            return this;
        }
        public Builder values(Integer... values) {
            return values(List.of(values));
        }
        public PodFailurePolicyOnExitCodesRequirement build() {
            final var o = new PodFailurePolicyOnExitCodesRequirement();
            o.containerName = containerName;
            o.operator = operator;
            o.values = values;
            return o;
        }
    }
}
