// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.admissionregistration.v1alpha1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class MatchConditionPatch {
    /**
     * @return Expression represents the expression which will be evaluated by CEL. Must evaluate to bool. CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables:
     * 
     * &#39;object&#39; - The object from the incoming request. The value is null for DELETE requests. &#39;oldObject&#39; - The existing object. The value is null for CREATE requests. &#39;request&#39; - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest). &#39;authorizer&#39; - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request.
     *   See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz
     * &#39;authorizer.requestResource&#39; - A CEL ResourceCheck constructed from the &#39;authorizer&#39; and configured with the
     *   request resource.
     * Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/
     * 
     * Required.
     * 
     */
    private @Nullable String expression;
    /**
     * @return Name is an identifier for this match condition, used for strategic merging of MatchConditions, as well as providing an identifier for logging purposes. A good name should be descriptive of the associated expression. Name must be a qualified name consisting of alphanumeric characters, &#39;-&#39;, &#39;_&#39; or &#39;.&#39;, and must start and end with an alphanumeric character (e.g. &#39;MyName&#39;,  or &#39;my.name&#39;,  or &#39;123-abc&#39;, regex used for validation is &#39;([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]&#39;) with an optional DNS subdomain prefix and &#39;/&#39; (e.g. &#39;example.com/MyName&#39;)
     * 
     * Required.
     * 
     */
    private @Nullable String name;

    private MatchConditionPatch() {}
    /**
     * @return Expression represents the expression which will be evaluated by CEL. Must evaluate to bool. CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables:
     * 
     * &#39;object&#39; - The object from the incoming request. The value is null for DELETE requests. &#39;oldObject&#39; - The existing object. The value is null for CREATE requests. &#39;request&#39; - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest). &#39;authorizer&#39; - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request.
     *   See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz
     * &#39;authorizer.requestResource&#39; - A CEL ResourceCheck constructed from the &#39;authorizer&#39; and configured with the
     *   request resource.
     * Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/
     * 
     * Required.
     * 
     */
    public Optional<String> expression() {
        return Optional.ofNullable(this.expression);
    }
    /**
     * @return Name is an identifier for this match condition, used for strategic merging of MatchConditions, as well as providing an identifier for logging purposes. A good name should be descriptive of the associated expression. Name must be a qualified name consisting of alphanumeric characters, &#39;-&#39;, &#39;_&#39; or &#39;.&#39;, and must start and end with an alphanumeric character (e.g. &#39;MyName&#39;,  or &#39;my.name&#39;,  or &#39;123-abc&#39;, regex used for validation is &#39;([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]&#39;) with an optional DNS subdomain prefix and &#39;/&#39; (e.g. &#39;example.com/MyName&#39;)
     * 
     * Required.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MatchConditionPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String expression;
        private @Nullable String name;
        public Builder() {}
        public Builder(MatchConditionPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.expression = defaults.expression;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder expression(@Nullable String expression) {
            this.expression = expression;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public MatchConditionPatch build() {
            final var o = new MatchConditionPatch();
            o.expression = expression;
            o.name = name;
            return o;
        }
    }
}
