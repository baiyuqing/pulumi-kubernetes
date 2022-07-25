// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.meta.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
 * 
 */
public final class LabelSelectorRequirementPatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final LabelSelectorRequirementPatchArgs Empty = new LabelSelectorRequirementPatchArgs();

    /**
     * key is the label key that the selector applies to.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return key is the label key that the selector applies to.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * operator represents a key&#39;s relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
     * 
     */
    @Import(name="operator")
    private @Nullable Output<String> operator;

    /**
     * @return operator represents a key&#39;s relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
     * 
     */
    public Optional<Output<String>> operator() {
        return Optional.ofNullable(this.operator);
    }

    /**
     * values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
     * 
     */
    @Import(name="values")
    private @Nullable Output<List<String>> values;

    /**
     * @return values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
     * 
     */
    public Optional<Output<List<String>>> values() {
        return Optional.ofNullable(this.values);
    }

    private LabelSelectorRequirementPatchArgs() {}

    private LabelSelectorRequirementPatchArgs(LabelSelectorRequirementPatchArgs $) {
        this.key = $.key;
        this.operator = $.operator;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LabelSelectorRequirementPatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LabelSelectorRequirementPatchArgs $;

        public Builder() {
            $ = new LabelSelectorRequirementPatchArgs();
        }

        public Builder(LabelSelectorRequirementPatchArgs defaults) {
            $ = new LabelSelectorRequirementPatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key key is the label key that the selector applies to.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key key is the label key that the selector applies to.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param operator operator represents a key&#39;s relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
         * 
         * @return builder
         * 
         */
        public Builder operator(@Nullable Output<String> operator) {
            $.operator = operator;
            return this;
        }

        /**
         * @param operator operator represents a key&#39;s relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
         * 
         * @return builder
         * 
         */
        public Builder operator(String operator) {
            return operator(Output.of(operator));
        }

        /**
         * @param values values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
         * 
         * @return builder
         * 
         */
        public Builder values(@Nullable Output<List<String>> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            return values(Output.of(values));
        }

        /**
         * @param values values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public LabelSelectorRequirementPatchArgs build() {
            return $;
        }
    }

}