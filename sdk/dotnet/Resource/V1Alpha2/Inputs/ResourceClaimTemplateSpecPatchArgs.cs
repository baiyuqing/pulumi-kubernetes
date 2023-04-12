// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2
{

    /// <summary>
    /// ResourceClaimTemplateSpec contains the metadata and fields for a ResourceClaim.
    /// </summary>
    public class ResourceClaimTemplateSpecPatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ObjectMeta may contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaPatchArgs>? Metadata { get; set; }

        /// <summary>
        /// Spec for the ResourceClaim. The entire content is copied unchanged into the ResourceClaim that gets created from this template. The same fields as in a ResourceClaim are also valid here.
        /// </summary>
        [Input("spec")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceClaimSpecPatchArgs>? Spec { get; set; }

        public ResourceClaimTemplateSpecPatchArgs()
        {
        }
        public static new ResourceClaimTemplateSpecPatchArgs Empty => new ResourceClaimTemplateSpecPatchArgs();
    }
}
