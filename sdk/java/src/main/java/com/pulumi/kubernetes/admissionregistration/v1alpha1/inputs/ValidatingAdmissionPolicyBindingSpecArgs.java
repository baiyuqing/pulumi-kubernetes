// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.admissionregistration.v1alpha1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.admissionregistration.v1alpha1.inputs.MatchResourcesArgs;
import com.pulumi.kubernetes.admissionregistration.v1alpha1.inputs.ParamRefArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * ValidatingAdmissionPolicyBindingSpec is the specification of the ValidatingAdmissionPolicyBinding.
 * 
 */
public final class ValidatingAdmissionPolicyBindingSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final ValidatingAdmissionPolicyBindingSpecArgs Empty = new ValidatingAdmissionPolicyBindingSpecArgs();

    /**
     * MatchResources declares what resources match this binding and will be validated by it. Note that this is intersected with the policy&#39;s matchConstraints, so only requests that are matched by the policy can be selected by this. If this is unset, all resources matched by the policy are validated by this binding When resourceRules is unset, it does not constrain resource matching. If a resource is matched by the other fields of this object, it will be validated. Note that this is differs from ValidatingAdmissionPolicy matchConstraints, where resourceRules are required.
     * 
     */
    @Import(name="matchResources")
    private @Nullable Output<MatchResourcesArgs> matchResources;

    /**
     * @return MatchResources declares what resources match this binding and will be validated by it. Note that this is intersected with the policy&#39;s matchConstraints, so only requests that are matched by the policy can be selected by this. If this is unset, all resources matched by the policy are validated by this binding When resourceRules is unset, it does not constrain resource matching. If a resource is matched by the other fields of this object, it will be validated. Note that this is differs from ValidatingAdmissionPolicy matchConstraints, where resourceRules are required.
     * 
     */
    public Optional<Output<MatchResourcesArgs>> matchResources() {
        return Optional.ofNullable(this.matchResources);
    }

    /**
     * ParamRef specifies the parameter resource used to configure the admission control policy. It should point to a resource of the type specified in ParamKind of the bound ValidatingAdmissionPolicy. If the policy specifies a ParamKind and the resource referred to by ParamRef does not exist, this binding is considered mis-configured and the FailurePolicy of the ValidatingAdmissionPolicy applied.
     * 
     */
    @Import(name="paramRef")
    private @Nullable Output<ParamRefArgs> paramRef;

    /**
     * @return ParamRef specifies the parameter resource used to configure the admission control policy. It should point to a resource of the type specified in ParamKind of the bound ValidatingAdmissionPolicy. If the policy specifies a ParamKind and the resource referred to by ParamRef does not exist, this binding is considered mis-configured and the FailurePolicy of the ValidatingAdmissionPolicy applied.
     * 
     */
    public Optional<Output<ParamRefArgs>> paramRef() {
        return Optional.ofNullable(this.paramRef);
    }

    /**
     * PolicyName references a ValidatingAdmissionPolicy name which the ValidatingAdmissionPolicyBinding binds to. If the referenced resource does not exist, this binding is considered invalid and will be ignored Required.
     * 
     */
    @Import(name="policyName")
    private @Nullable Output<String> policyName;

    /**
     * @return PolicyName references a ValidatingAdmissionPolicy name which the ValidatingAdmissionPolicyBinding binds to. If the referenced resource does not exist, this binding is considered invalid and will be ignored Required.
     * 
     */
    public Optional<Output<String>> policyName() {
        return Optional.ofNullable(this.policyName);
    }

    /**
     * validationActions declares how Validations of the referenced ValidatingAdmissionPolicy are enforced. If a validation evaluates to false it is always enforced according to these actions.
     * 
     * Failures defined by the ValidatingAdmissionPolicy&#39;s FailurePolicy are enforced according to these actions only if the FailurePolicy is set to Fail, otherwise the failures are ignored. This includes compilation errors, runtime errors and misconfigurations of the policy.
     * 
     * validationActions is declared as a set of action values. Order does not matter. validationActions may not contain duplicates of the same action.
     * 
     * The supported actions values are:
     * 
     * &#34;Deny&#34; specifies that a validation failure results in a denied request.
     * 
     * &#34;Warn&#34; specifies that a validation failure is reported to the request client in HTTP Warning headers, with a warning code of 299. Warnings can be sent both for allowed or denied admission responses.
     * 
     * &#34;Audit&#34; specifies that a validation failure is included in the published audit event for the request. The audit event will contain a `validation.policy.admission.k8s.io/validation_failure` audit annotation with a value containing the details of the validation failures, formatted as a JSON list of objects, each with the following fields: - message: The validation failure message string - policy: The resource name of the ValidatingAdmissionPolicy - binding: The resource name of the ValidatingAdmissionPolicyBinding - expressionIndex: The index of the failed validations in the ValidatingAdmissionPolicy - validationActions: The enforcement actions enacted for the validation failure Example audit annotation: `&#34;validation.policy.admission.k8s.io/validation_failure&#34;: &#34;[{&#34;message&#34;: &#34;Invalid value&#34;, {&#34;policy&#34;: &#34;policy.example.com&#34;, {&#34;binding&#34;: &#34;policybinding.example.com&#34;, {&#34;expressionIndex&#34;: &#34;1&#34;, {&#34;validationActions&#34;: [&#34;Audit&#34;]}]&#34;`
     * 
     * Clients should expect to handle additional values by ignoring any values not recognized.
     * 
     * &#34;Deny&#34; and &#34;Warn&#34; may not be used together since this combination needlessly duplicates the validation failure both in the API response body and the HTTP warning headers.
     * 
     * Required.
     * 
     */
    @Import(name="validationActions")
    private @Nullable Output<List<String>> validationActions;

    /**
     * @return validationActions declares how Validations of the referenced ValidatingAdmissionPolicy are enforced. If a validation evaluates to false it is always enforced according to these actions.
     * 
     * Failures defined by the ValidatingAdmissionPolicy&#39;s FailurePolicy are enforced according to these actions only if the FailurePolicy is set to Fail, otherwise the failures are ignored. This includes compilation errors, runtime errors and misconfigurations of the policy.
     * 
     * validationActions is declared as a set of action values. Order does not matter. validationActions may not contain duplicates of the same action.
     * 
     * The supported actions values are:
     * 
     * &#34;Deny&#34; specifies that a validation failure results in a denied request.
     * 
     * &#34;Warn&#34; specifies that a validation failure is reported to the request client in HTTP Warning headers, with a warning code of 299. Warnings can be sent both for allowed or denied admission responses.
     * 
     * &#34;Audit&#34; specifies that a validation failure is included in the published audit event for the request. The audit event will contain a `validation.policy.admission.k8s.io/validation_failure` audit annotation with a value containing the details of the validation failures, formatted as a JSON list of objects, each with the following fields: - message: The validation failure message string - policy: The resource name of the ValidatingAdmissionPolicy - binding: The resource name of the ValidatingAdmissionPolicyBinding - expressionIndex: The index of the failed validations in the ValidatingAdmissionPolicy - validationActions: The enforcement actions enacted for the validation failure Example audit annotation: `&#34;validation.policy.admission.k8s.io/validation_failure&#34;: &#34;[{&#34;message&#34;: &#34;Invalid value&#34;, {&#34;policy&#34;: &#34;policy.example.com&#34;, {&#34;binding&#34;: &#34;policybinding.example.com&#34;, {&#34;expressionIndex&#34;: &#34;1&#34;, {&#34;validationActions&#34;: [&#34;Audit&#34;]}]&#34;`
     * 
     * Clients should expect to handle additional values by ignoring any values not recognized.
     * 
     * &#34;Deny&#34; and &#34;Warn&#34; may not be used together since this combination needlessly duplicates the validation failure both in the API response body and the HTTP warning headers.
     * 
     * Required.
     * 
     */
    public Optional<Output<List<String>>> validationActions() {
        return Optional.ofNullable(this.validationActions);
    }

    private ValidatingAdmissionPolicyBindingSpecArgs() {}

    private ValidatingAdmissionPolicyBindingSpecArgs(ValidatingAdmissionPolicyBindingSpecArgs $) {
        this.matchResources = $.matchResources;
        this.paramRef = $.paramRef;
        this.policyName = $.policyName;
        this.validationActions = $.validationActions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ValidatingAdmissionPolicyBindingSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ValidatingAdmissionPolicyBindingSpecArgs $;

        public Builder() {
            $ = new ValidatingAdmissionPolicyBindingSpecArgs();
        }

        public Builder(ValidatingAdmissionPolicyBindingSpecArgs defaults) {
            $ = new ValidatingAdmissionPolicyBindingSpecArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param matchResources MatchResources declares what resources match this binding and will be validated by it. Note that this is intersected with the policy&#39;s matchConstraints, so only requests that are matched by the policy can be selected by this. If this is unset, all resources matched by the policy are validated by this binding When resourceRules is unset, it does not constrain resource matching. If a resource is matched by the other fields of this object, it will be validated. Note that this is differs from ValidatingAdmissionPolicy matchConstraints, where resourceRules are required.
         * 
         * @return builder
         * 
         */
        public Builder matchResources(@Nullable Output<MatchResourcesArgs> matchResources) {
            $.matchResources = matchResources;
            return this;
        }

        /**
         * @param matchResources MatchResources declares what resources match this binding and will be validated by it. Note that this is intersected with the policy&#39;s matchConstraints, so only requests that are matched by the policy can be selected by this. If this is unset, all resources matched by the policy are validated by this binding When resourceRules is unset, it does not constrain resource matching. If a resource is matched by the other fields of this object, it will be validated. Note that this is differs from ValidatingAdmissionPolicy matchConstraints, where resourceRules are required.
         * 
         * @return builder
         * 
         */
        public Builder matchResources(MatchResourcesArgs matchResources) {
            return matchResources(Output.of(matchResources));
        }

        /**
         * @param paramRef ParamRef specifies the parameter resource used to configure the admission control policy. It should point to a resource of the type specified in ParamKind of the bound ValidatingAdmissionPolicy. If the policy specifies a ParamKind and the resource referred to by ParamRef does not exist, this binding is considered mis-configured and the FailurePolicy of the ValidatingAdmissionPolicy applied.
         * 
         * @return builder
         * 
         */
        public Builder paramRef(@Nullable Output<ParamRefArgs> paramRef) {
            $.paramRef = paramRef;
            return this;
        }

        /**
         * @param paramRef ParamRef specifies the parameter resource used to configure the admission control policy. It should point to a resource of the type specified in ParamKind of the bound ValidatingAdmissionPolicy. If the policy specifies a ParamKind and the resource referred to by ParamRef does not exist, this binding is considered mis-configured and the FailurePolicy of the ValidatingAdmissionPolicy applied.
         * 
         * @return builder
         * 
         */
        public Builder paramRef(ParamRefArgs paramRef) {
            return paramRef(Output.of(paramRef));
        }

        /**
         * @param policyName PolicyName references a ValidatingAdmissionPolicy name which the ValidatingAdmissionPolicyBinding binds to. If the referenced resource does not exist, this binding is considered invalid and will be ignored Required.
         * 
         * @return builder
         * 
         */
        public Builder policyName(@Nullable Output<String> policyName) {
            $.policyName = policyName;
            return this;
        }

        /**
         * @param policyName PolicyName references a ValidatingAdmissionPolicy name which the ValidatingAdmissionPolicyBinding binds to. If the referenced resource does not exist, this binding is considered invalid and will be ignored Required.
         * 
         * @return builder
         * 
         */
        public Builder policyName(String policyName) {
            return policyName(Output.of(policyName));
        }

        /**
         * @param validationActions validationActions declares how Validations of the referenced ValidatingAdmissionPolicy are enforced. If a validation evaluates to false it is always enforced according to these actions.
         * 
         * Failures defined by the ValidatingAdmissionPolicy&#39;s FailurePolicy are enforced according to these actions only if the FailurePolicy is set to Fail, otherwise the failures are ignored. This includes compilation errors, runtime errors and misconfigurations of the policy.
         * 
         * validationActions is declared as a set of action values. Order does not matter. validationActions may not contain duplicates of the same action.
         * 
         * The supported actions values are:
         * 
         * &#34;Deny&#34; specifies that a validation failure results in a denied request.
         * 
         * &#34;Warn&#34; specifies that a validation failure is reported to the request client in HTTP Warning headers, with a warning code of 299. Warnings can be sent both for allowed or denied admission responses.
         * 
         * &#34;Audit&#34; specifies that a validation failure is included in the published audit event for the request. The audit event will contain a `validation.policy.admission.k8s.io/validation_failure` audit annotation with a value containing the details of the validation failures, formatted as a JSON list of objects, each with the following fields: - message: The validation failure message string - policy: The resource name of the ValidatingAdmissionPolicy - binding: The resource name of the ValidatingAdmissionPolicyBinding - expressionIndex: The index of the failed validations in the ValidatingAdmissionPolicy - validationActions: The enforcement actions enacted for the validation failure Example audit annotation: `&#34;validation.policy.admission.k8s.io/validation_failure&#34;: &#34;[{&#34;message&#34;: &#34;Invalid value&#34;, {&#34;policy&#34;: &#34;policy.example.com&#34;, {&#34;binding&#34;: &#34;policybinding.example.com&#34;, {&#34;expressionIndex&#34;: &#34;1&#34;, {&#34;validationActions&#34;: [&#34;Audit&#34;]}]&#34;`
         * 
         * Clients should expect to handle additional values by ignoring any values not recognized.
         * 
         * &#34;Deny&#34; and &#34;Warn&#34; may not be used together since this combination needlessly duplicates the validation failure both in the API response body and the HTTP warning headers.
         * 
         * Required.
         * 
         * @return builder
         * 
         */
        public Builder validationActions(@Nullable Output<List<String>> validationActions) {
            $.validationActions = validationActions;
            return this;
        }

        /**
         * @param validationActions validationActions declares how Validations of the referenced ValidatingAdmissionPolicy are enforced. If a validation evaluates to false it is always enforced according to these actions.
         * 
         * Failures defined by the ValidatingAdmissionPolicy&#39;s FailurePolicy are enforced according to these actions only if the FailurePolicy is set to Fail, otherwise the failures are ignored. This includes compilation errors, runtime errors and misconfigurations of the policy.
         * 
         * validationActions is declared as a set of action values. Order does not matter. validationActions may not contain duplicates of the same action.
         * 
         * The supported actions values are:
         * 
         * &#34;Deny&#34; specifies that a validation failure results in a denied request.
         * 
         * &#34;Warn&#34; specifies that a validation failure is reported to the request client in HTTP Warning headers, with a warning code of 299. Warnings can be sent both for allowed or denied admission responses.
         * 
         * &#34;Audit&#34; specifies that a validation failure is included in the published audit event for the request. The audit event will contain a `validation.policy.admission.k8s.io/validation_failure` audit annotation with a value containing the details of the validation failures, formatted as a JSON list of objects, each with the following fields: - message: The validation failure message string - policy: The resource name of the ValidatingAdmissionPolicy - binding: The resource name of the ValidatingAdmissionPolicyBinding - expressionIndex: The index of the failed validations in the ValidatingAdmissionPolicy - validationActions: The enforcement actions enacted for the validation failure Example audit annotation: `&#34;validation.policy.admission.k8s.io/validation_failure&#34;: &#34;[{&#34;message&#34;: &#34;Invalid value&#34;, {&#34;policy&#34;: &#34;policy.example.com&#34;, {&#34;binding&#34;: &#34;policybinding.example.com&#34;, {&#34;expressionIndex&#34;: &#34;1&#34;, {&#34;validationActions&#34;: [&#34;Audit&#34;]}]&#34;`
         * 
         * Clients should expect to handle additional values by ignoring any values not recognized.
         * 
         * &#34;Deny&#34; and &#34;Warn&#34; may not be used together since this combination needlessly duplicates the validation failure both in the API response body and the HTTP warning headers.
         * 
         * Required.
         * 
         * @return builder
         * 
         */
        public Builder validationActions(List<String> validationActions) {
            return validationActions(Output.of(validationActions));
        }

        /**
         * @param validationActions validationActions declares how Validations of the referenced ValidatingAdmissionPolicy are enforced. If a validation evaluates to false it is always enforced according to these actions.
         * 
         * Failures defined by the ValidatingAdmissionPolicy&#39;s FailurePolicy are enforced according to these actions only if the FailurePolicy is set to Fail, otherwise the failures are ignored. This includes compilation errors, runtime errors and misconfigurations of the policy.
         * 
         * validationActions is declared as a set of action values. Order does not matter. validationActions may not contain duplicates of the same action.
         * 
         * The supported actions values are:
         * 
         * &#34;Deny&#34; specifies that a validation failure results in a denied request.
         * 
         * &#34;Warn&#34; specifies that a validation failure is reported to the request client in HTTP Warning headers, with a warning code of 299. Warnings can be sent both for allowed or denied admission responses.
         * 
         * &#34;Audit&#34; specifies that a validation failure is included in the published audit event for the request. The audit event will contain a `validation.policy.admission.k8s.io/validation_failure` audit annotation with a value containing the details of the validation failures, formatted as a JSON list of objects, each with the following fields: - message: The validation failure message string - policy: The resource name of the ValidatingAdmissionPolicy - binding: The resource name of the ValidatingAdmissionPolicyBinding - expressionIndex: The index of the failed validations in the ValidatingAdmissionPolicy - validationActions: The enforcement actions enacted for the validation failure Example audit annotation: `&#34;validation.policy.admission.k8s.io/validation_failure&#34;: &#34;[{&#34;message&#34;: &#34;Invalid value&#34;, {&#34;policy&#34;: &#34;policy.example.com&#34;, {&#34;binding&#34;: &#34;policybinding.example.com&#34;, {&#34;expressionIndex&#34;: &#34;1&#34;, {&#34;validationActions&#34;: [&#34;Audit&#34;]}]&#34;`
         * 
         * Clients should expect to handle additional values by ignoring any values not recognized.
         * 
         * &#34;Deny&#34; and &#34;Warn&#34; may not be used together since this combination needlessly duplicates the validation failure both in the API response body and the HTTP warning headers.
         * 
         * Required.
         * 
         * @return builder
         * 
         */
        public Builder validationActions(String... validationActions) {
            return validationActions(List.of(validationActions));
        }

        public ValidatingAdmissionPolicyBindingSpecArgs build() {
            return $;
        }
    }

}
