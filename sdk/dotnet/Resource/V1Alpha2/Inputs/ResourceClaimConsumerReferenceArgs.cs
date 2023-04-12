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
    /// ResourceClaimConsumerReference contains enough information to let you locate the consumer of a ResourceClaim. The user must be a resource in the same namespace as the ResourceClaim.
    /// </summary>
    public class ResourceClaimConsumerReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
        /// </summary>
        [Input("apiGroup")]
        public Input<string>? ApiGroup { get; set; }

        /// <summary>
        /// Name is the name of resource being referenced.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Resource is the type of resource being referenced, for example "pods".
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// UID identifies exactly one incarnation of the resource.
        /// </summary>
        [Input("uid", required: true)]
        public Input<string> Uid { get; set; } = null!;

        public ResourceClaimConsumerReferenceArgs()
        {
        }
        public static new ResourceClaimConsumerReferenceArgs Empty => new ResourceClaimConsumerReferenceArgs();
    }
}
