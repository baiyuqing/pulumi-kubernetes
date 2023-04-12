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
    /// ResourceClaimParametersReference contains enough information to let you locate the parameters for a ResourceClaim. The object must be in the same namespace as the ResourceClaim.
    /// </summary>
    public class ResourceClaimParametersReferencePatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
        /// </summary>
        [Input("apiGroup")]
        public Input<string>? ApiGroup { get; set; }

        /// <summary>
        /// Kind is the type of resource being referenced. This is the same value as in the parameter object's metadata, for example "ConfigMap".
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name is the name of resource being referenced.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ResourceClaimParametersReferencePatchArgs()
        {
        }
        public static new ResourceClaimParametersReferencePatchArgs Empty => new ResourceClaimParametersReferencePatchArgs();
    }
}
