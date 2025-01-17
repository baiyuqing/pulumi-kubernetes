// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { MutatingWebhookConfigurationArgs } from "./mutatingWebhookConfiguration";
export type MutatingWebhookConfiguration = import("./mutatingWebhookConfiguration").MutatingWebhookConfiguration;
export const MutatingWebhookConfiguration: typeof import("./mutatingWebhookConfiguration").MutatingWebhookConfiguration = null as any;
utilities.lazyLoad(exports, ["MutatingWebhookConfiguration"], () => require("./mutatingWebhookConfiguration"));

export { MutatingWebhookConfigurationListArgs } from "./mutatingWebhookConfigurationList";
export type MutatingWebhookConfigurationList = import("./mutatingWebhookConfigurationList").MutatingWebhookConfigurationList;
export const MutatingWebhookConfigurationList: typeof import("./mutatingWebhookConfigurationList").MutatingWebhookConfigurationList = null as any;
utilities.lazyLoad(exports, ["MutatingWebhookConfigurationList"], () => require("./mutatingWebhookConfigurationList"));

export { MutatingWebhookConfigurationPatchArgs } from "./mutatingWebhookConfigurationPatch";
export type MutatingWebhookConfigurationPatch = import("./mutatingWebhookConfigurationPatch").MutatingWebhookConfigurationPatch;
export const MutatingWebhookConfigurationPatch: typeof import("./mutatingWebhookConfigurationPatch").MutatingWebhookConfigurationPatch = null as any;
utilities.lazyLoad(exports, ["MutatingWebhookConfigurationPatch"], () => require("./mutatingWebhookConfigurationPatch"));

export { ValidatingAdmissionPolicyArgs } from "./validatingAdmissionPolicy";
export type ValidatingAdmissionPolicy = import("./validatingAdmissionPolicy").ValidatingAdmissionPolicy;
export const ValidatingAdmissionPolicy: typeof import("./validatingAdmissionPolicy").ValidatingAdmissionPolicy = null as any;
utilities.lazyLoad(exports, ["ValidatingAdmissionPolicy"], () => require("./validatingAdmissionPolicy"));

export { ValidatingAdmissionPolicyBindingArgs } from "./validatingAdmissionPolicyBinding";
export type ValidatingAdmissionPolicyBinding = import("./validatingAdmissionPolicyBinding").ValidatingAdmissionPolicyBinding;
export const ValidatingAdmissionPolicyBinding: typeof import("./validatingAdmissionPolicyBinding").ValidatingAdmissionPolicyBinding = null as any;
utilities.lazyLoad(exports, ["ValidatingAdmissionPolicyBinding"], () => require("./validatingAdmissionPolicyBinding"));

export { ValidatingAdmissionPolicyBindingListArgs } from "./validatingAdmissionPolicyBindingList";
export type ValidatingAdmissionPolicyBindingList = import("./validatingAdmissionPolicyBindingList").ValidatingAdmissionPolicyBindingList;
export const ValidatingAdmissionPolicyBindingList: typeof import("./validatingAdmissionPolicyBindingList").ValidatingAdmissionPolicyBindingList = null as any;
utilities.lazyLoad(exports, ["ValidatingAdmissionPolicyBindingList"], () => require("./validatingAdmissionPolicyBindingList"));

export { ValidatingAdmissionPolicyBindingPatchArgs } from "./validatingAdmissionPolicyBindingPatch";
export type ValidatingAdmissionPolicyBindingPatch = import("./validatingAdmissionPolicyBindingPatch").ValidatingAdmissionPolicyBindingPatch;
export const ValidatingAdmissionPolicyBindingPatch: typeof import("./validatingAdmissionPolicyBindingPatch").ValidatingAdmissionPolicyBindingPatch = null as any;
utilities.lazyLoad(exports, ["ValidatingAdmissionPolicyBindingPatch"], () => require("./validatingAdmissionPolicyBindingPatch"));

export { ValidatingAdmissionPolicyListArgs } from "./validatingAdmissionPolicyList";
export type ValidatingAdmissionPolicyList = import("./validatingAdmissionPolicyList").ValidatingAdmissionPolicyList;
export const ValidatingAdmissionPolicyList: typeof import("./validatingAdmissionPolicyList").ValidatingAdmissionPolicyList = null as any;
utilities.lazyLoad(exports, ["ValidatingAdmissionPolicyList"], () => require("./validatingAdmissionPolicyList"));

export { ValidatingAdmissionPolicyPatchArgs } from "./validatingAdmissionPolicyPatch";
export type ValidatingAdmissionPolicyPatch = import("./validatingAdmissionPolicyPatch").ValidatingAdmissionPolicyPatch;
export const ValidatingAdmissionPolicyPatch: typeof import("./validatingAdmissionPolicyPatch").ValidatingAdmissionPolicyPatch = null as any;
utilities.lazyLoad(exports, ["ValidatingAdmissionPolicyPatch"], () => require("./validatingAdmissionPolicyPatch"));

export { ValidatingWebhookConfigurationArgs } from "./validatingWebhookConfiguration";
export type ValidatingWebhookConfiguration = import("./validatingWebhookConfiguration").ValidatingWebhookConfiguration;
export const ValidatingWebhookConfiguration: typeof import("./validatingWebhookConfiguration").ValidatingWebhookConfiguration = null as any;
utilities.lazyLoad(exports, ["ValidatingWebhookConfiguration"], () => require("./validatingWebhookConfiguration"));

export { ValidatingWebhookConfigurationListArgs } from "./validatingWebhookConfigurationList";
export type ValidatingWebhookConfigurationList = import("./validatingWebhookConfigurationList").ValidatingWebhookConfigurationList;
export const ValidatingWebhookConfigurationList: typeof import("./validatingWebhookConfigurationList").ValidatingWebhookConfigurationList = null as any;
utilities.lazyLoad(exports, ["ValidatingWebhookConfigurationList"], () => require("./validatingWebhookConfigurationList"));

export { ValidatingWebhookConfigurationPatchArgs } from "./validatingWebhookConfigurationPatch";
export type ValidatingWebhookConfigurationPatch = import("./validatingWebhookConfigurationPatch").ValidatingWebhookConfigurationPatch;
export const ValidatingWebhookConfigurationPatch: typeof import("./validatingWebhookConfigurationPatch").ValidatingWebhookConfigurationPatch = null as any;
utilities.lazyLoad(exports, ["ValidatingWebhookConfigurationPatch"], () => require("./validatingWebhookConfigurationPatch"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfiguration":
                return new MutatingWebhookConfiguration(name, <any>undefined, { urn })
            case "kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfigurationList":
                return new MutatingWebhookConfigurationList(name, <any>undefined, { urn })
            case "kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfigurationPatch":
                return new MutatingWebhookConfigurationPatch(name, <any>undefined, { urn })
            case "kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicy":
                return new ValidatingAdmissionPolicy(name, <any>undefined, { urn })
            case "kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicyBinding":
                return new ValidatingAdmissionPolicyBinding(name, <any>undefined, { urn })
            case "kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicyBindingList":
                return new ValidatingAdmissionPolicyBindingList(name, <any>undefined, { urn })
            case "kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicyBindingPatch":
                return new ValidatingAdmissionPolicyBindingPatch(name, <any>undefined, { urn })
            case "kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicyList":
                return new ValidatingAdmissionPolicyList(name, <any>undefined, { urn })
            case "kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicyPatch":
                return new ValidatingAdmissionPolicyPatch(name, <any>undefined, { urn })
            case "kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingWebhookConfiguration":
                return new ValidatingWebhookConfiguration(name, <any>undefined, { urn })
            case "kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingWebhookConfigurationList":
                return new ValidatingWebhookConfigurationList(name, <any>undefined, { urn })
            case "kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingWebhookConfigurationPatch":
                return new ValidatingWebhookConfigurationPatch(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes", "admissionregistration.k8s.io/v1beta1", _module)
