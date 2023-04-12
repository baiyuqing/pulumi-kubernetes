// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.admissionregistration.v1alpha1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class AuditAnnotation {
    /**
     * @return key specifies the audit annotation key. The audit annotation keys of a ValidatingAdmissionPolicy must be unique. The key must be a qualified name ([A-Za-z0-9][-A-Za-z0-9_.]*) no more than 63 bytes in length.
     * 
     * The key is combined with the resource name of the ValidatingAdmissionPolicy to construct an audit annotation key: &#34;{ValidatingAdmissionPolicy name}/{key}&#34;.
     * 
     * If an admission webhook uses the same resource name as this ValidatingAdmissionPolicy and the same audit annotation key, the annotation key will be identical. In this case, the first annotation written with the key will be included in the audit event and all subsequent annotations with the same key will be discarded.
     * 
     * Required.
     * 
     */
    private String key;
    /**
     * @return valueExpression represents the expression which is evaluated by CEL to produce an audit annotation value. The expression must evaluate to either a string or null value. If the expression evaluates to a string, the audit annotation is included with the string value. If the expression evaluates to null or empty string the audit annotation will be omitted. The valueExpression may be no longer than 5kb in length. If the result of the valueExpression is more than 10kb in length, it will be truncated to 10kb.
     * 
     * If multiple ValidatingAdmissionPolicyBinding resources match an API request, then the valueExpression will be evaluated for each binding. All unique values produced by the valueExpressions will be joined together in a comma-separated list.
     * 
     * Required.
     * 
     */
    private String valueExpression;

    private AuditAnnotation() {}
    /**
     * @return key specifies the audit annotation key. The audit annotation keys of a ValidatingAdmissionPolicy must be unique. The key must be a qualified name ([A-Za-z0-9][-A-Za-z0-9_.]*) no more than 63 bytes in length.
     * 
     * The key is combined with the resource name of the ValidatingAdmissionPolicy to construct an audit annotation key: &#34;{ValidatingAdmissionPolicy name}/{key}&#34;.
     * 
     * If an admission webhook uses the same resource name as this ValidatingAdmissionPolicy and the same audit annotation key, the annotation key will be identical. In this case, the first annotation written with the key will be included in the audit event and all subsequent annotations with the same key will be discarded.
     * 
     * Required.
     * 
     */
    public String key() {
        return this.key;
    }
    /**
     * @return valueExpression represents the expression which is evaluated by CEL to produce an audit annotation value. The expression must evaluate to either a string or null value. If the expression evaluates to a string, the audit annotation is included with the string value. If the expression evaluates to null or empty string the audit annotation will be omitted. The valueExpression may be no longer than 5kb in length. If the result of the valueExpression is more than 10kb in length, it will be truncated to 10kb.
     * 
     * If multiple ValidatingAdmissionPolicyBinding resources match an API request, then the valueExpression will be evaluated for each binding. All unique values produced by the valueExpressions will be joined together in a comma-separated list.
     * 
     * Required.
     * 
     */
    public String valueExpression() {
        return this.valueExpression;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AuditAnnotation defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String key;
        private String valueExpression;
        public Builder() {}
        public Builder(AuditAnnotation defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.key = defaults.key;
    	      this.valueExpression = defaults.valueExpression;
        }

        @CustomType.Setter
        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
            return this;
        }
        @CustomType.Setter
        public Builder valueExpression(String valueExpression) {
            this.valueExpression = Objects.requireNonNull(valueExpression);
            return this;
        }
        public AuditAnnotation build() {
            final var o = new AuditAnnotation();
            o.key = key;
            o.valueExpression = valueExpression;
            return o;
        }
    }
}
