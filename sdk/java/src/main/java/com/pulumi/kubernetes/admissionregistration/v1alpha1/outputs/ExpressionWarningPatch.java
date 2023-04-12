// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.admissionregistration.v1alpha1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ExpressionWarningPatch {
    /**
     * @return The path to the field that refers the expression. For example, the reference to the expression of the first item of validations is &#34;spec.validations[0].expression&#34;
     * 
     */
    private @Nullable String fieldRef;
    /**
     * @return The content of type checking information in a human-readable form. Each line of the warning contains the type that the expression is checked against, followed by the type check error from the compiler.
     * 
     */
    private @Nullable String warning;

    private ExpressionWarningPatch() {}
    /**
     * @return The path to the field that refers the expression. For example, the reference to the expression of the first item of validations is &#34;spec.validations[0].expression&#34;
     * 
     */
    public Optional<String> fieldRef() {
        return Optional.ofNullable(this.fieldRef);
    }
    /**
     * @return The content of type checking information in a human-readable form. Each line of the warning contains the type that the expression is checked against, followed by the type check error from the compiler.
     * 
     */
    public Optional<String> warning() {
        return Optional.ofNullable(this.warning);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ExpressionWarningPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String fieldRef;
        private @Nullable String warning;
        public Builder() {}
        public Builder(ExpressionWarningPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fieldRef = defaults.fieldRef;
    	      this.warning = defaults.warning;
        }

        @CustomType.Setter
        public Builder fieldRef(@Nullable String fieldRef) {
            this.fieldRef = fieldRef;
            return this;
        }
        @CustomType.Setter
        public Builder warning(@Nullable String warning) {
            this.warning = warning;
            return this;
        }
        public ExpressionWarningPatch build() {
            final var o = new ExpressionWarningPatch();
            o.fieldRef = fieldRef;
            o.warning = warning;
            return o;
        }
    }
}
