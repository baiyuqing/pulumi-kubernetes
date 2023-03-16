// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../../types/input";
import * as outputs from "../../../types/output";
import * as enums from "../../../types/enums";
import * as utilities from "../../../utilities";

/**
 * MatchResources decides whether to run the admission control policy on an object based on whether it meets the match criteria. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)
 */
export interface MatchResources {
    /**
     * ExcludeResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy should not care about. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)
     */
    excludeResourceRules?: pulumi.Input<pulumi.Input<inputs.admissionregistration.v1alpha1.NamedRuleWithOperations>[]>;
    /**
     * matchPolicy defines how the "MatchResources" list is used to match incoming requests. Allowed values are "Exact" or "Equivalent".
     *
     * - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the ValidatingAdmissionPolicy.
     *
     * - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the ValidatingAdmissionPolicy.
     *
     * Defaults to "Equivalent"
     */
    matchPolicy?: pulumi.Input<string>;
    /**
     * NamespaceSelector decides whether to run the admission control policy on an object based on whether the namespace for that object matches the selector. If the object itself is a namespace, the matching is performed on object.metadata.labels. If the object is another cluster scoped resource, it never skips the policy.
     *
     * For example, to run the webhook on any objects whose namespace is not associated with "runlevel" of "0" or "1";  you will set the selector as follows: "namespaceSelector": {
     *   "matchExpressions": [
     *     {
     *       "key": "runlevel",
     *       "operator": "NotIn",
     *       "values": [
     *         "0",
     *         "1"
     *       ]
     *     }
     *   ]
     * }
     *
     * If instead you want to only run the policy on any objects whose namespace is associated with the "environment" of "prod" or "staging"; you will set the selector as follows: "namespaceSelector": {
     *   "matchExpressions": [
     *     {
     *       "key": "environment",
     *       "operator": "In",
     *       "values": [
     *         "prod",
     *         "staging"
     *       ]
     *     }
     *   ]
     * }
     *
     * See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more examples of label selectors.
     *
     * Default to the empty LabelSelector, which matches everything.
     */
    namespaceSelector?: pulumi.Input<inputs.meta.v1.LabelSelector>;
    /**
     * ObjectSelector decides whether to run the validation based on if the object has matching labels. objectSelector is evaluated against both the oldObject and newObject that would be sent to the cel validation, and is considered to match if either object matches the selector. A null object (oldObject in the case of create, or newObject in the case of delete) or an object that cannot have labels (like a DeploymentRollback or a PodProxyOptions object) is not considered to match. Use the object selector only if the webhook is opt-in, because end users may skip the admission webhook by setting the labels. Default to the empty LabelSelector, which matches everything.
     */
    objectSelector?: pulumi.Input<inputs.meta.v1.LabelSelector>;
    /**
     * ResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy matches. The policy cares about an operation if it matches _any_ Rule.
     */
    resourceRules?: pulumi.Input<pulumi.Input<inputs.admissionregistration.v1alpha1.NamedRuleWithOperations>[]>;
}

/**
 * MatchResources decides whether to run the admission control policy on an object based on whether it meets the match criteria. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)
 */
export interface MatchResourcesPatch {
    /**
     * ExcludeResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy should not care about. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)
     */
    excludeResourceRules?: pulumi.Input<pulumi.Input<inputs.admissionregistration.v1alpha1.NamedRuleWithOperationsPatch>[]>;
    /**
     * matchPolicy defines how the "MatchResources" list is used to match incoming requests. Allowed values are "Exact" or "Equivalent".
     *
     * - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the ValidatingAdmissionPolicy.
     *
     * - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the ValidatingAdmissionPolicy.
     *
     * Defaults to "Equivalent"
     */
    matchPolicy?: pulumi.Input<string>;
    /**
     * NamespaceSelector decides whether to run the admission control policy on an object based on whether the namespace for that object matches the selector. If the object itself is a namespace, the matching is performed on object.metadata.labels. If the object is another cluster scoped resource, it never skips the policy.
     *
     * For example, to run the webhook on any objects whose namespace is not associated with "runlevel" of "0" or "1";  you will set the selector as follows: "namespaceSelector": {
     *   "matchExpressions": [
     *     {
     *       "key": "runlevel",
     *       "operator": "NotIn",
     *       "values": [
     *         "0",
     *         "1"
     *       ]
     *     }
     *   ]
     * }
     *
     * If instead you want to only run the policy on any objects whose namespace is associated with the "environment" of "prod" or "staging"; you will set the selector as follows: "namespaceSelector": {
     *   "matchExpressions": [
     *     {
     *       "key": "environment",
     *       "operator": "In",
     *       "values": [
     *         "prod",
     *         "staging"
     *       ]
     *     }
     *   ]
     * }
     *
     * See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more examples of label selectors.
     *
     * Default to the empty LabelSelector, which matches everything.
     */
    namespaceSelector?: pulumi.Input<inputs.meta.v1.LabelSelectorPatch>;
    /**
     * ObjectSelector decides whether to run the validation based on if the object has matching labels. objectSelector is evaluated against both the oldObject and newObject that would be sent to the cel validation, and is considered to match if either object matches the selector. A null object (oldObject in the case of create, or newObject in the case of delete) or an object that cannot have labels (like a DeploymentRollback or a PodProxyOptions object) is not considered to match. Use the object selector only if the webhook is opt-in, because end users may skip the admission webhook by setting the labels. Default to the empty LabelSelector, which matches everything.
     */
    objectSelector?: pulumi.Input<inputs.meta.v1.LabelSelectorPatch>;
    /**
     * ResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy matches. The policy cares about an operation if it matches _any_ Rule.
     */
    resourceRules?: pulumi.Input<pulumi.Input<inputs.admissionregistration.v1alpha1.NamedRuleWithOperationsPatch>[]>;
}

/**
 * NamedRuleWithOperations is a tuple of Operations and Resources with ResourceNames.
 */
export interface NamedRuleWithOperations {
    /**
     * APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.
     */
    apiGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.
     */
    apiVersions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.
     */
    operations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
     */
    resourceNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resources is a list of resources this rule applies to.
     *
     * For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*&#47;scale' means all scale subresources. '*&#47;*' means all resources and their subresources.
     *
     * If wildcard is present, the validation rule will ensure resources do not overlap with each other.
     *
     * Depending on the enclosing object, subresources might not be allowed. Required.
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * scope specifies the scope of this rule. Valid values are "Cluster", "Namespaced", and "*" "Cluster" means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. "Namespaced" means that only namespaced resources will match this rule. "*" means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is "*".
     */
    scope?: pulumi.Input<string>;
}

/**
 * NamedRuleWithOperations is a tuple of Operations and Resources with ResourceNames.
 */
export interface NamedRuleWithOperationsPatch {
    /**
     * APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.
     */
    apiGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.
     */
    apiVersions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.
     */
    operations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
     */
    resourceNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resources is a list of resources this rule applies to.
     *
     * For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*&#47;scale' means all scale subresources. '*&#47;*' means all resources and their subresources.
     *
     * If wildcard is present, the validation rule will ensure resources do not overlap with each other.
     *
     * Depending on the enclosing object, subresources might not be allowed. Required.
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * scope specifies the scope of this rule. Valid values are "Cluster", "Namespaced", and "*" "Cluster" means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. "Namespaced" means that only namespaced resources will match this rule. "*" means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is "*".
     */
    scope?: pulumi.Input<string>;
}

/**
 * ParamKind is a tuple of Group Kind and Version.
 */
export interface ParamKind {
    /**
     * APIVersion is the API group version the resources belong to. In format of "group/version". Required.
     */
    apiVersion?: pulumi.Input<string>;
    /**
     * Kind is the API kind the resources belong to. Required.
     */
    kind?: pulumi.Input<string>;
}

/**
 * ParamKind is a tuple of Group Kind and Version.
 */
export interface ParamKindPatch {
    /**
     * APIVersion is the API group version the resources belong to. In format of "group/version". Required.
     */
    apiVersion?: pulumi.Input<string>;
    /**
     * Kind is the API kind the resources belong to. Required.
     */
    kind?: pulumi.Input<string>;
}

/**
 * ParamRef references a parameter resource
 */
export interface ParamRef {
    /**
     * Name of the resource being referenced.
     */
    name?: pulumi.Input<string>;
    /**
     * Namespace of the referenced resource. Should be empty for the cluster-scoped resources
     */
    namespace?: pulumi.Input<string>;
}

/**
 * ParamRef references a parameter resource
 */
export interface ParamRefPatch {
    /**
     * Name of the resource being referenced.
     */
    name?: pulumi.Input<string>;
    /**
     * Namespace of the referenced resource. Should be empty for the cluster-scoped resources
     */
    namespace?: pulumi.Input<string>;
}

/**
 * ValidatingAdmissionPolicy describes the definition of an admission validation policy that accepts or rejects an object without changing it.
 */
export interface ValidatingAdmissionPolicy {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: pulumi.Input<"admissionregistration.k8s.io/v1alpha1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: pulumi.Input<"ValidatingAdmissionPolicy">;
    /**
     * Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
     */
    metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * Specification of the desired behavior of the ValidatingAdmissionPolicy.
     */
    spec?: pulumi.Input<inputs.admissionregistration.v1alpha1.ValidatingAdmissionPolicySpec>;
}

/**
 * ValidatingAdmissionPolicyBinding binds the ValidatingAdmissionPolicy with paramerized resources. ValidatingAdmissionPolicyBinding and parameter CRDs together define how cluster administrators configure policies for clusters.
 */
export interface ValidatingAdmissionPolicyBinding {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: pulumi.Input<"admissionregistration.k8s.io/v1alpha1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: pulumi.Input<"ValidatingAdmissionPolicyBinding">;
    /**
     * Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
     */
    metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * Specification of the desired behavior of the ValidatingAdmissionPolicyBinding.
     */
    spec?: pulumi.Input<inputs.admissionregistration.v1alpha1.ValidatingAdmissionPolicyBindingSpec>;
}

/**
 * ValidatingAdmissionPolicyBindingSpec is the specification of the ValidatingAdmissionPolicyBinding.
 */
export interface ValidatingAdmissionPolicyBindingSpec {
    /**
     * MatchResources declares what resources match this binding and will be validated by it. Note that this is intersected with the policy's matchConstraints, so only requests that are matched by the policy can be selected by this. If this is unset, all resources matched by the policy are validated by this binding When resourceRules is unset, it does not constrain resource matching. If a resource is matched by the other fields of this object, it will be validated. Note that this is differs from ValidatingAdmissionPolicy matchConstraints, where resourceRules are required.
     */
    matchResources?: pulumi.Input<inputs.admissionregistration.v1alpha1.MatchResources>;
    /**
     * ParamRef specifies the parameter resource used to configure the admission control policy. It should point to a resource of the type specified in ParamKind of the bound ValidatingAdmissionPolicy. If the policy specifies a ParamKind and the resource referred to by ParamRef does not exist, this binding is considered mis-configured and the FailurePolicy of the ValidatingAdmissionPolicy applied.
     */
    paramRef?: pulumi.Input<inputs.admissionregistration.v1alpha1.ParamRef>;
    /**
     * PolicyName references a ValidatingAdmissionPolicy name which the ValidatingAdmissionPolicyBinding binds to. If the referenced resource does not exist, this binding is considered invalid and will be ignored Required.
     */
    policyName?: pulumi.Input<string>;
}

/**
 * ValidatingAdmissionPolicyBindingSpec is the specification of the ValidatingAdmissionPolicyBinding.
 */
export interface ValidatingAdmissionPolicyBindingSpecPatch {
    /**
     * MatchResources declares what resources match this binding and will be validated by it. Note that this is intersected with the policy's matchConstraints, so only requests that are matched by the policy can be selected by this. If this is unset, all resources matched by the policy are validated by this binding When resourceRules is unset, it does not constrain resource matching. If a resource is matched by the other fields of this object, it will be validated. Note that this is differs from ValidatingAdmissionPolicy matchConstraints, where resourceRules are required.
     */
    matchResources?: pulumi.Input<inputs.admissionregistration.v1alpha1.MatchResourcesPatch>;
    /**
     * ParamRef specifies the parameter resource used to configure the admission control policy. It should point to a resource of the type specified in ParamKind of the bound ValidatingAdmissionPolicy. If the policy specifies a ParamKind and the resource referred to by ParamRef does not exist, this binding is considered mis-configured and the FailurePolicy of the ValidatingAdmissionPolicy applied.
     */
    paramRef?: pulumi.Input<inputs.admissionregistration.v1alpha1.ParamRefPatch>;
    /**
     * PolicyName references a ValidatingAdmissionPolicy name which the ValidatingAdmissionPolicyBinding binds to. If the referenced resource does not exist, this binding is considered invalid and will be ignored Required.
     */
    policyName?: pulumi.Input<string>;
}

/**
 * ValidatingAdmissionPolicySpec is the specification of the desired behavior of the AdmissionPolicy.
 */
export interface ValidatingAdmissionPolicySpec {
    /**
     * FailurePolicy defines how to handle failures for the admission policy. Failures can occur from invalid or mis-configured policy definitions or bindings. A policy is invalid if spec.paramKind refers to a non-existent Kind. A binding is invalid if spec.paramRef.name refers to a non-existent resource. Allowed values are Ignore or Fail. Defaults to Fail.
     */
    failurePolicy?: pulumi.Input<string>;
    /**
     * MatchConstraints specifies what resources this policy is designed to validate. The AdmissionPolicy cares about a request if it matches _all_ Constraints. However, in order to prevent clusters from being put into an unstable state that cannot be recovered from via the API ValidatingAdmissionPolicy cannot match ValidatingAdmissionPolicy and ValidatingAdmissionPolicyBinding. Required.
     */
    matchConstraints?: pulumi.Input<inputs.admissionregistration.v1alpha1.MatchResources>;
    /**
     * ParamKind specifies the kind of resources used to parameterize this policy. If absent, there are no parameters for this policy and the param CEL variable will not be provided to validation expressions. If ParamKind refers to a non-existent kind, this policy definition is mis-configured and the FailurePolicy is applied. If paramKind is specified but paramRef is unset in ValidatingAdmissionPolicyBinding, the params variable will be null.
     */
    paramKind?: pulumi.Input<inputs.admissionregistration.v1alpha1.ParamKind>;
    /**
     * Validations contain CEL expressions which is used to apply the validation. A minimum of one validation is required for a policy definition. Required.
     */
    validations: pulumi.Input<pulumi.Input<inputs.admissionregistration.v1alpha1.Validation>[]>;
}

/**
 * ValidatingAdmissionPolicySpec is the specification of the desired behavior of the AdmissionPolicy.
 */
export interface ValidatingAdmissionPolicySpecPatch {
    /**
     * FailurePolicy defines how to handle failures for the admission policy. Failures can occur from invalid or mis-configured policy definitions or bindings. A policy is invalid if spec.paramKind refers to a non-existent Kind. A binding is invalid if spec.paramRef.name refers to a non-existent resource. Allowed values are Ignore or Fail. Defaults to Fail.
     */
    failurePolicy?: pulumi.Input<string>;
    /**
     * MatchConstraints specifies what resources this policy is designed to validate. The AdmissionPolicy cares about a request if it matches _all_ Constraints. However, in order to prevent clusters from being put into an unstable state that cannot be recovered from via the API ValidatingAdmissionPolicy cannot match ValidatingAdmissionPolicy and ValidatingAdmissionPolicyBinding. Required.
     */
    matchConstraints?: pulumi.Input<inputs.admissionregistration.v1alpha1.MatchResourcesPatch>;
    /**
     * ParamKind specifies the kind of resources used to parameterize this policy. If absent, there are no parameters for this policy and the param CEL variable will not be provided to validation expressions. If ParamKind refers to a non-existent kind, this policy definition is mis-configured and the FailurePolicy is applied. If paramKind is specified but paramRef is unset in ValidatingAdmissionPolicyBinding, the params variable will be null.
     */
    paramKind?: pulumi.Input<inputs.admissionregistration.v1alpha1.ParamKindPatch>;
    /**
     * Validations contain CEL expressions which is used to apply the validation. A minimum of one validation is required for a policy definition. Required.
     */
    validations?: pulumi.Input<pulumi.Input<inputs.admissionregistration.v1alpha1.ValidationPatch>[]>;
}

/**
 * Validation specifies the CEL expression which is used to apply the validation.
 */
export interface Validation {
    /**
     * Expression represents the expression which will be evaluated by CEL. ref: https://github.com/google/cel-spec CEL expressions have access to the contents of the Admission request/response, organized into CEL variables as well as some other useful variables:
     *
     * 'object' - The object from the incoming request. The value is null for DELETE requests. 'oldObject' - The existing object. The value is null for CREATE requests. 'request' - Attributes of the admission request([ref](/pkg/apis/admission/types.go#AdmissionRequest)). 'params' - Parameter resource referred to by the policy binding being evaluated. Only populated if the policy has a ParamKind.
     *
     * The `apiVersion`, `kind`, `metadata.name` and `metadata.generateName` are always accessible from the root of the object. No other metadata properties are accessible.
     *
     * Only property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - '__' escapes to '__underscores__' - '.' escapes to '__dot__' - '-' escapes to '__dash__' - '/' escapes to '__slash__' - Property names that exactly match a CEL RESERVED keyword escape to '__{keyword}__'. The keywords are:
     * 	  "true", "false", "null", "in", "as", "break", "const", "continue", "else", "for", "function", "if",
     * 	  "import", "let", "loop", "package", "namespace", "return".
     * Examples:
     *   - Expression accessing a property named "namespace": {"Expression": "object.__namespace__ > 0"}
     *   - Expression accessing a property named "x-prop": {"Expression": "object.x__dash__prop > 0"}
     *   - Expression accessing a property named "redact__d": {"Expression": "object.redact__underscores__d > 0"}
     *
     * Equality on arrays with list type of 'set' or 'map' ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type:
     *   - 'set': `X + Y` performs a union where the array positions of all elements in `X` are preserved and
     *     non-intersecting elements in `Y` are appended, retaining their partial order.
     *   - 'map': `X + Y` performs a merge where the array positions of all keys in `X` are preserved but the values
     *     are overwritten by values in `Y` when the key sets of `X` and `Y` intersect. Elements in `Y` with
     *     non-intersecting keys are appended, retaining their partial order.
     * Required.
     */
    expression: pulumi.Input<string>;
    /**
     * Message represents the message displayed when validation fails. The message is required if the Expression contains line breaks. The message must not contain line breaks. If unset, the message is "failed rule: {Rule}". e.g. "must be a URL with the host matching spec.host" If the Expression contains line breaks. Message is required. The message must not contain line breaks. If unset, the message is "failed Expression: {Expression}".
     */
    message?: pulumi.Input<string>;
    /**
     * Reason represents a machine-readable description of why this validation failed. If this is the first validation in the list to fail, this reason, as well as the corresponding HTTP response code, are used in the HTTP response to the client. The currently supported reasons are: "Unauthorized", "Forbidden", "Invalid", "RequestEntityTooLarge". If not set, StatusReasonInvalid is used in the response to the client.
     */
    reason?: pulumi.Input<string>;
}

/**
 * Validation specifies the CEL expression which is used to apply the validation.
 */
export interface ValidationPatch {
    /**
     * Expression represents the expression which will be evaluated by CEL. ref: https://github.com/google/cel-spec CEL expressions have access to the contents of the Admission request/response, organized into CEL variables as well as some other useful variables:
     *
     * 'object' - The object from the incoming request. The value is null for DELETE requests. 'oldObject' - The existing object. The value is null for CREATE requests. 'request' - Attributes of the admission request([ref](/pkg/apis/admission/types.go#AdmissionRequest)). 'params' - Parameter resource referred to by the policy binding being evaluated. Only populated if the policy has a ParamKind.
     *
     * The `apiVersion`, `kind`, `metadata.name` and `metadata.generateName` are always accessible from the root of the object. No other metadata properties are accessible.
     *
     * Only property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - '__' escapes to '__underscores__' - '.' escapes to '__dot__' - '-' escapes to '__dash__' - '/' escapes to '__slash__' - Property names that exactly match a CEL RESERVED keyword escape to '__{keyword}__'. The keywords are:
     * 	  "true", "false", "null", "in", "as", "break", "const", "continue", "else", "for", "function", "if",
     * 	  "import", "let", "loop", "package", "namespace", "return".
     * Examples:
     *   - Expression accessing a property named "namespace": {"Expression": "object.__namespace__ > 0"}
     *   - Expression accessing a property named "x-prop": {"Expression": "object.x__dash__prop > 0"}
     *   - Expression accessing a property named "redact__d": {"Expression": "object.redact__underscores__d > 0"}
     *
     * Equality on arrays with list type of 'set' or 'map' ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type:
     *   - 'set': `X + Y` performs a union where the array positions of all elements in `X` are preserved and
     *     non-intersecting elements in `Y` are appended, retaining their partial order.
     *   - 'map': `X + Y` performs a merge where the array positions of all keys in `X` are preserved but the values
     *     are overwritten by values in `Y` when the key sets of `X` and `Y` intersect. Elements in `Y` with
     *     non-intersecting keys are appended, retaining their partial order.
     * Required.
     */
    expression?: pulumi.Input<string>;
    /**
     * Message represents the message displayed when validation fails. The message is required if the Expression contains line breaks. The message must not contain line breaks. If unset, the message is "failed rule: {Rule}". e.g. "must be a URL with the host matching spec.host" If the Expression contains line breaks. Message is required. The message must not contain line breaks. If unset, the message is "failed Expression: {Expression}".
     */
    message?: pulumi.Input<string>;
    /**
     * Reason represents a machine-readable description of why this validation failed. If this is the first validation in the list to fail, this reason, as well as the corresponding HTTP response code, are used in the HTTP response to the client. The currently supported reasons are: "Unauthorized", "Forbidden", "Invalid", "RequestEntityTooLarge". If not set, StatusReasonInvalid is used in the response to the client.
     */
    reason?: pulumi.Input<string>;
}