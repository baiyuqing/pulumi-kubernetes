// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.apiextensions.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * ValidationRule describes a validation rule written in the CEL expression language.
 * 
 */
public final class ValidationRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final ValidationRuleArgs Empty = new ValidationRuleArgs();

    /**
     * Message represents the message displayed when validation fails. The message is required if the Rule contains line breaks. The message must not contain line breaks. If unset, the message is &#34;failed rule: {Rule}&#34;. e.g. &#34;must be a URL with the host matching spec.host&#34;
     * 
     */
    @Import(name="message")
    private @Nullable Output<String> message;

    /**
     * @return Message represents the message displayed when validation fails. The message is required if the Rule contains line breaks. The message must not contain line breaks. If unset, the message is &#34;failed rule: {Rule}&#34;. e.g. &#34;must be a URL with the host matching spec.host&#34;
     * 
     */
    public Optional<Output<String>> message() {
        return Optional.ofNullable(this.message);
    }

    /**
     * MessageExpression declares a CEL expression that evaluates to the validation failure message that is returned when this rule fails. Since messageExpression is used as a failure message, it must evaluate to a string. If both message and messageExpression are present on a rule, then messageExpression will be used if validation fails. If messageExpression results in a runtime error, the runtime error is logged, and the validation failure message is produced as if the messageExpression field were unset. If messageExpression evaluates to an empty string, a string with only spaces, or a string that contains line breaks, then the validation failure message will also be produced as if the messageExpression field were unset, and the fact that messageExpression produced an empty string/string with only spaces/string with line breaks will be logged. messageExpression has access to all the same variables as the rule; the only difference is the return type. Example: &#34;x must be less than max (&#34;+string(self.max)+&#34;)&#34;
     * 
     */
    @Import(name="messageExpression")
    private @Nullable Output<String> messageExpression;

    /**
     * @return MessageExpression declares a CEL expression that evaluates to the validation failure message that is returned when this rule fails. Since messageExpression is used as a failure message, it must evaluate to a string. If both message and messageExpression are present on a rule, then messageExpression will be used if validation fails. If messageExpression results in a runtime error, the runtime error is logged, and the validation failure message is produced as if the messageExpression field were unset. If messageExpression evaluates to an empty string, a string with only spaces, or a string that contains line breaks, then the validation failure message will also be produced as if the messageExpression field were unset, and the fact that messageExpression produced an empty string/string with only spaces/string with line breaks will be logged. messageExpression has access to all the same variables as the rule; the only difference is the return type. Example: &#34;x must be less than max (&#34;+string(self.max)+&#34;)&#34;
     * 
     */
    public Optional<Output<String>> messageExpression() {
        return Optional.ofNullable(this.messageExpression);
    }

    /**
     * Rule represents the expression which will be evaluated by CEL. ref: https://github.com/google/cel-spec The Rule is scoped to the location of the x-kubernetes-validations extension in the schema. The `self` variable in the CEL expression is bound to the scoped value. Example: - Rule scoped to the root of a resource with a status subresource: {&#34;rule&#34;: &#34;self.status.actual &lt;= self.spec.maxDesired&#34;}
     * 
     * If the Rule is scoped to an object with properties, the accessible properties of the object are field selectable via `self.field` and field presence can be checked via `has(self.field)`. Null valued fields are treated as absent fields in CEL expressions. If the Rule is scoped to an object with additionalProperties (i.e. a map) the value of the map are accessible via `self[mapKey]`, map containment can be checked via `mapKey in self` and all entries of the map are accessible via CEL macros and functions such as `self.all(...)`. If the Rule is scoped to an array, the elements of the array are accessible via `self[i]` and also by macros and functions. If the Rule is scoped to a scalar, `self` is bound to the scalar value. Examples: - Rule scoped to a map of objects: {&#34;rule&#34;: &#34;self.components[&#39;Widget&#39;].priority &lt; 10&#34;} - Rule scoped to a list of integers: {&#34;rule&#34;: &#34;self.values.all(value, value &gt;= 0 &amp;&amp; value &lt; 100)&#34;} - Rule scoped to a string value: {&#34;rule&#34;: &#34;self.startsWith(&#39;kube&#39;)&#34;}
     * 
     * The `apiVersion`, `kind`, `metadata.name` and `metadata.generateName` are always accessible from the root of the object and from any x-kubernetes-embedded-resource annotated objects. No other metadata properties are accessible.
     * 
     * Unknown data preserved in custom resources via x-kubernetes-preserve-unknown-fields is not accessible in CEL expressions. This includes: - Unknown field values that are preserved by object schemas with x-kubernetes-preserve-unknown-fields. - Object properties where the property schema is of an &#34;unknown type&#34;. An &#34;unknown type&#34; is recursively defined as:
     *   - A schema with no type and x-kubernetes-preserve-unknown-fields set to true
     *   - An array where the items schema is of an &#34;unknown type&#34;
     *   - An object where the additionalProperties schema is of an &#34;unknown type&#34;
     * 
     * Only property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - &#39;__&#39; escapes to &#39;__underscores__&#39; - &#39;.&#39; escapes to &#39;__dot__&#39; - &#39;-&#39; escapes to &#39;__dash__&#39; - &#39;/&#39; escapes to &#39;__slash__&#39; - Property names that exactly match a CEL RESERVED keyword escape to &#39;__{keyword}__&#39;. The keywords are:
     * 	  &#34;true&#34;, &#34;false&#34;, &#34;null&#34;, &#34;in&#34;, &#34;as&#34;, &#34;break&#34;, &#34;const&#34;, &#34;continue&#34;, &#34;else&#34;, &#34;for&#34;, &#34;function&#34;, &#34;if&#34;,
     * 	  &#34;import&#34;, &#34;let&#34;, &#34;loop&#34;, &#34;package&#34;, &#34;namespace&#34;, &#34;return&#34;.
     * Examples:
     *   - Rule accessing a property named &#34;namespace&#34;: {&#34;rule&#34;: &#34;self.__namespace__ &gt; 0&#34;}
     *   - Rule accessing a property named &#34;x-prop&#34;: {&#34;rule&#34;: &#34;self.x__dash__prop &gt; 0&#34;}
     *   - Rule accessing a property named &#34;redact__d&#34;: {&#34;rule&#34;: &#34;self.redact__underscores__d &gt; 0&#34;}
     * 
     * Equality on arrays with x-kubernetes-list-type of &#39;set&#39; or &#39;map&#39; ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type:
     *   - &#39;set&#39;: `X + Y` performs a union where the array positions of all elements in `X` are preserved and
     *     non-intersecting elements in `Y` are appended, retaining their partial order.
     *   - &#39;map&#39;: `X + Y` performs a merge where the array positions of all keys in `X` are preserved but the values
     *     are overwritten by values in `Y` when the key sets of `X` and `Y` intersect. Elements in `Y` with
     *     non-intersecting keys are appended, retaining their partial order.
     * 
     */
    @Import(name="rule", required=true)
    private Output<String> rule;

    /**
     * @return Rule represents the expression which will be evaluated by CEL. ref: https://github.com/google/cel-spec The Rule is scoped to the location of the x-kubernetes-validations extension in the schema. The `self` variable in the CEL expression is bound to the scoped value. Example: - Rule scoped to the root of a resource with a status subresource: {&#34;rule&#34;: &#34;self.status.actual &lt;= self.spec.maxDesired&#34;}
     * 
     * If the Rule is scoped to an object with properties, the accessible properties of the object are field selectable via `self.field` and field presence can be checked via `has(self.field)`. Null valued fields are treated as absent fields in CEL expressions. If the Rule is scoped to an object with additionalProperties (i.e. a map) the value of the map are accessible via `self[mapKey]`, map containment can be checked via `mapKey in self` and all entries of the map are accessible via CEL macros and functions such as `self.all(...)`. If the Rule is scoped to an array, the elements of the array are accessible via `self[i]` and also by macros and functions. If the Rule is scoped to a scalar, `self` is bound to the scalar value. Examples: - Rule scoped to a map of objects: {&#34;rule&#34;: &#34;self.components[&#39;Widget&#39;].priority &lt; 10&#34;} - Rule scoped to a list of integers: {&#34;rule&#34;: &#34;self.values.all(value, value &gt;= 0 &amp;&amp; value &lt; 100)&#34;} - Rule scoped to a string value: {&#34;rule&#34;: &#34;self.startsWith(&#39;kube&#39;)&#34;}
     * 
     * The `apiVersion`, `kind`, `metadata.name` and `metadata.generateName` are always accessible from the root of the object and from any x-kubernetes-embedded-resource annotated objects. No other metadata properties are accessible.
     * 
     * Unknown data preserved in custom resources via x-kubernetes-preserve-unknown-fields is not accessible in CEL expressions. This includes: - Unknown field values that are preserved by object schemas with x-kubernetes-preserve-unknown-fields. - Object properties where the property schema is of an &#34;unknown type&#34;. An &#34;unknown type&#34; is recursively defined as:
     *   - A schema with no type and x-kubernetes-preserve-unknown-fields set to true
     *   - An array where the items schema is of an &#34;unknown type&#34;
     *   - An object where the additionalProperties schema is of an &#34;unknown type&#34;
     * 
     * Only property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - &#39;__&#39; escapes to &#39;__underscores__&#39; - &#39;.&#39; escapes to &#39;__dot__&#39; - &#39;-&#39; escapes to &#39;__dash__&#39; - &#39;/&#39; escapes to &#39;__slash__&#39; - Property names that exactly match a CEL RESERVED keyword escape to &#39;__{keyword}__&#39;. The keywords are:
     * 	  &#34;true&#34;, &#34;false&#34;, &#34;null&#34;, &#34;in&#34;, &#34;as&#34;, &#34;break&#34;, &#34;const&#34;, &#34;continue&#34;, &#34;else&#34;, &#34;for&#34;, &#34;function&#34;, &#34;if&#34;,
     * 	  &#34;import&#34;, &#34;let&#34;, &#34;loop&#34;, &#34;package&#34;, &#34;namespace&#34;, &#34;return&#34;.
     * Examples:
     *   - Rule accessing a property named &#34;namespace&#34;: {&#34;rule&#34;: &#34;self.__namespace__ &gt; 0&#34;}
     *   - Rule accessing a property named &#34;x-prop&#34;: {&#34;rule&#34;: &#34;self.x__dash__prop &gt; 0&#34;}
     *   - Rule accessing a property named &#34;redact__d&#34;: {&#34;rule&#34;: &#34;self.redact__underscores__d &gt; 0&#34;}
     * 
     * Equality on arrays with x-kubernetes-list-type of &#39;set&#39; or &#39;map&#39; ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type:
     *   - &#39;set&#39;: `X + Y` performs a union where the array positions of all elements in `X` are preserved and
     *     non-intersecting elements in `Y` are appended, retaining their partial order.
     *   - &#39;map&#39;: `X + Y` performs a merge where the array positions of all keys in `X` are preserved but the values
     *     are overwritten by values in `Y` when the key sets of `X` and `Y` intersect. Elements in `Y` with
     *     non-intersecting keys are appended, retaining their partial order.
     * 
     */
    public Output<String> rule() {
        return this.rule;
    }

    private ValidationRuleArgs() {}

    private ValidationRuleArgs(ValidationRuleArgs $) {
        this.message = $.message;
        this.messageExpression = $.messageExpression;
        this.rule = $.rule;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ValidationRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ValidationRuleArgs $;

        public Builder() {
            $ = new ValidationRuleArgs();
        }

        public Builder(ValidationRuleArgs defaults) {
            $ = new ValidationRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param message Message represents the message displayed when validation fails. The message is required if the Rule contains line breaks. The message must not contain line breaks. If unset, the message is &#34;failed rule: {Rule}&#34;. e.g. &#34;must be a URL with the host matching spec.host&#34;
         * 
         * @return builder
         * 
         */
        public Builder message(@Nullable Output<String> message) {
            $.message = message;
            return this;
        }

        /**
         * @param message Message represents the message displayed when validation fails. The message is required if the Rule contains line breaks. The message must not contain line breaks. If unset, the message is &#34;failed rule: {Rule}&#34;. e.g. &#34;must be a URL with the host matching spec.host&#34;
         * 
         * @return builder
         * 
         */
        public Builder message(String message) {
            return message(Output.of(message));
        }

        /**
         * @param messageExpression MessageExpression declares a CEL expression that evaluates to the validation failure message that is returned when this rule fails. Since messageExpression is used as a failure message, it must evaluate to a string. If both message and messageExpression are present on a rule, then messageExpression will be used if validation fails. If messageExpression results in a runtime error, the runtime error is logged, and the validation failure message is produced as if the messageExpression field were unset. If messageExpression evaluates to an empty string, a string with only spaces, or a string that contains line breaks, then the validation failure message will also be produced as if the messageExpression field were unset, and the fact that messageExpression produced an empty string/string with only spaces/string with line breaks will be logged. messageExpression has access to all the same variables as the rule; the only difference is the return type. Example: &#34;x must be less than max (&#34;+string(self.max)+&#34;)&#34;
         * 
         * @return builder
         * 
         */
        public Builder messageExpression(@Nullable Output<String> messageExpression) {
            $.messageExpression = messageExpression;
            return this;
        }

        /**
         * @param messageExpression MessageExpression declares a CEL expression that evaluates to the validation failure message that is returned when this rule fails. Since messageExpression is used as a failure message, it must evaluate to a string. If both message and messageExpression are present on a rule, then messageExpression will be used if validation fails. If messageExpression results in a runtime error, the runtime error is logged, and the validation failure message is produced as if the messageExpression field were unset. If messageExpression evaluates to an empty string, a string with only spaces, or a string that contains line breaks, then the validation failure message will also be produced as if the messageExpression field were unset, and the fact that messageExpression produced an empty string/string with only spaces/string with line breaks will be logged. messageExpression has access to all the same variables as the rule; the only difference is the return type. Example: &#34;x must be less than max (&#34;+string(self.max)+&#34;)&#34;
         * 
         * @return builder
         * 
         */
        public Builder messageExpression(String messageExpression) {
            return messageExpression(Output.of(messageExpression));
        }

        /**
         * @param rule Rule represents the expression which will be evaluated by CEL. ref: https://github.com/google/cel-spec The Rule is scoped to the location of the x-kubernetes-validations extension in the schema. The `self` variable in the CEL expression is bound to the scoped value. Example: - Rule scoped to the root of a resource with a status subresource: {&#34;rule&#34;: &#34;self.status.actual &lt;= self.spec.maxDesired&#34;}
         * 
         * If the Rule is scoped to an object with properties, the accessible properties of the object are field selectable via `self.field` and field presence can be checked via `has(self.field)`. Null valued fields are treated as absent fields in CEL expressions. If the Rule is scoped to an object with additionalProperties (i.e. a map) the value of the map are accessible via `self[mapKey]`, map containment can be checked via `mapKey in self` and all entries of the map are accessible via CEL macros and functions such as `self.all(...)`. If the Rule is scoped to an array, the elements of the array are accessible via `self[i]` and also by macros and functions. If the Rule is scoped to a scalar, `self` is bound to the scalar value. Examples: - Rule scoped to a map of objects: {&#34;rule&#34;: &#34;self.components[&#39;Widget&#39;].priority &lt; 10&#34;} - Rule scoped to a list of integers: {&#34;rule&#34;: &#34;self.values.all(value, value &gt;= 0 &amp;&amp; value &lt; 100)&#34;} - Rule scoped to a string value: {&#34;rule&#34;: &#34;self.startsWith(&#39;kube&#39;)&#34;}
         * 
         * The `apiVersion`, `kind`, `metadata.name` and `metadata.generateName` are always accessible from the root of the object and from any x-kubernetes-embedded-resource annotated objects. No other metadata properties are accessible.
         * 
         * Unknown data preserved in custom resources via x-kubernetes-preserve-unknown-fields is not accessible in CEL expressions. This includes: - Unknown field values that are preserved by object schemas with x-kubernetes-preserve-unknown-fields. - Object properties where the property schema is of an &#34;unknown type&#34;. An &#34;unknown type&#34; is recursively defined as:
         *   - A schema with no type and x-kubernetes-preserve-unknown-fields set to true
         *   - An array where the items schema is of an &#34;unknown type&#34;
         *   - An object where the additionalProperties schema is of an &#34;unknown type&#34;
         * 
         * Only property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - &#39;__&#39; escapes to &#39;__underscores__&#39; - &#39;.&#39; escapes to &#39;__dot__&#39; - &#39;-&#39; escapes to &#39;__dash__&#39; - &#39;/&#39; escapes to &#39;__slash__&#39; - Property names that exactly match a CEL RESERVED keyword escape to &#39;__{keyword}__&#39;. The keywords are:
         * 	  &#34;true&#34;, &#34;false&#34;, &#34;null&#34;, &#34;in&#34;, &#34;as&#34;, &#34;break&#34;, &#34;const&#34;, &#34;continue&#34;, &#34;else&#34;, &#34;for&#34;, &#34;function&#34;, &#34;if&#34;,
         * 	  &#34;import&#34;, &#34;let&#34;, &#34;loop&#34;, &#34;package&#34;, &#34;namespace&#34;, &#34;return&#34;.
         * Examples:
         *   - Rule accessing a property named &#34;namespace&#34;: {&#34;rule&#34;: &#34;self.__namespace__ &gt; 0&#34;}
         *   - Rule accessing a property named &#34;x-prop&#34;: {&#34;rule&#34;: &#34;self.x__dash__prop &gt; 0&#34;}
         *   - Rule accessing a property named &#34;redact__d&#34;: {&#34;rule&#34;: &#34;self.redact__underscores__d &gt; 0&#34;}
         * 
         * Equality on arrays with x-kubernetes-list-type of &#39;set&#39; or &#39;map&#39; ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type:
         *   - &#39;set&#39;: `X + Y` performs a union where the array positions of all elements in `X` are preserved and
         *     non-intersecting elements in `Y` are appended, retaining their partial order.
         *   - &#39;map&#39;: `X + Y` performs a merge where the array positions of all keys in `X` are preserved but the values
         *     are overwritten by values in `Y` when the key sets of `X` and `Y` intersect. Elements in `Y` with
         *     non-intersecting keys are appended, retaining their partial order.
         * 
         * @return builder
         * 
         */
        public Builder rule(Output<String> rule) {
            $.rule = rule;
            return this;
        }

        /**
         * @param rule Rule represents the expression which will be evaluated by CEL. ref: https://github.com/google/cel-spec The Rule is scoped to the location of the x-kubernetes-validations extension in the schema. The `self` variable in the CEL expression is bound to the scoped value. Example: - Rule scoped to the root of a resource with a status subresource: {&#34;rule&#34;: &#34;self.status.actual &lt;= self.spec.maxDesired&#34;}
         * 
         * If the Rule is scoped to an object with properties, the accessible properties of the object are field selectable via `self.field` and field presence can be checked via `has(self.field)`. Null valued fields are treated as absent fields in CEL expressions. If the Rule is scoped to an object with additionalProperties (i.e. a map) the value of the map are accessible via `self[mapKey]`, map containment can be checked via `mapKey in self` and all entries of the map are accessible via CEL macros and functions such as `self.all(...)`. If the Rule is scoped to an array, the elements of the array are accessible via `self[i]` and also by macros and functions. If the Rule is scoped to a scalar, `self` is bound to the scalar value. Examples: - Rule scoped to a map of objects: {&#34;rule&#34;: &#34;self.components[&#39;Widget&#39;].priority &lt; 10&#34;} - Rule scoped to a list of integers: {&#34;rule&#34;: &#34;self.values.all(value, value &gt;= 0 &amp;&amp; value &lt; 100)&#34;} - Rule scoped to a string value: {&#34;rule&#34;: &#34;self.startsWith(&#39;kube&#39;)&#34;}
         * 
         * The `apiVersion`, `kind`, `metadata.name` and `metadata.generateName` are always accessible from the root of the object and from any x-kubernetes-embedded-resource annotated objects. No other metadata properties are accessible.
         * 
         * Unknown data preserved in custom resources via x-kubernetes-preserve-unknown-fields is not accessible in CEL expressions. This includes: - Unknown field values that are preserved by object schemas with x-kubernetes-preserve-unknown-fields. - Object properties where the property schema is of an &#34;unknown type&#34;. An &#34;unknown type&#34; is recursively defined as:
         *   - A schema with no type and x-kubernetes-preserve-unknown-fields set to true
         *   - An array where the items schema is of an &#34;unknown type&#34;
         *   - An object where the additionalProperties schema is of an &#34;unknown type&#34;
         * 
         * Only property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - &#39;__&#39; escapes to &#39;__underscores__&#39; - &#39;.&#39; escapes to &#39;__dot__&#39; - &#39;-&#39; escapes to &#39;__dash__&#39; - &#39;/&#39; escapes to &#39;__slash__&#39; - Property names that exactly match a CEL RESERVED keyword escape to &#39;__{keyword}__&#39;. The keywords are:
         * 	  &#34;true&#34;, &#34;false&#34;, &#34;null&#34;, &#34;in&#34;, &#34;as&#34;, &#34;break&#34;, &#34;const&#34;, &#34;continue&#34;, &#34;else&#34;, &#34;for&#34;, &#34;function&#34;, &#34;if&#34;,
         * 	  &#34;import&#34;, &#34;let&#34;, &#34;loop&#34;, &#34;package&#34;, &#34;namespace&#34;, &#34;return&#34;.
         * Examples:
         *   - Rule accessing a property named &#34;namespace&#34;: {&#34;rule&#34;: &#34;self.__namespace__ &gt; 0&#34;}
         *   - Rule accessing a property named &#34;x-prop&#34;: {&#34;rule&#34;: &#34;self.x__dash__prop &gt; 0&#34;}
         *   - Rule accessing a property named &#34;redact__d&#34;: {&#34;rule&#34;: &#34;self.redact__underscores__d &gt; 0&#34;}
         * 
         * Equality on arrays with x-kubernetes-list-type of &#39;set&#39; or &#39;map&#39; ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type:
         *   - &#39;set&#39;: `X + Y` performs a union where the array positions of all elements in `X` are preserved and
         *     non-intersecting elements in `Y` are appended, retaining their partial order.
         *   - &#39;map&#39;: `X + Y` performs a merge where the array positions of all keys in `X` are preserved but the values
         *     are overwritten by values in `Y` when the key sets of `X` and `Y` intersect. Elements in `Y` with
         *     non-intersecting keys are appended, retaining their partial order.
         * 
         * @return builder
         * 
         */
        public Builder rule(String rule) {
            return rule(Output.of(rule));
        }

        public ValidationRuleArgs build() {
            $.rule = Objects.requireNonNull($.rule, "expected parameter 'rule' to be non-null");
            return $;
        }
    }

}
