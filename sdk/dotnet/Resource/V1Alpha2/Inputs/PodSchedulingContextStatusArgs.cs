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
    /// PodSchedulingContextStatus describes where resources for the Pod can be allocated.
    /// </summary>
    public class PodSchedulingContextStatusArgs : global::Pulumi.ResourceArgs
    {
        [Input("resourceClaims")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceClaimSchedulingStatusArgs>? _resourceClaims;

        /// <summary>
        /// ResourceClaims describes resource availability for each pod.spec.resourceClaim entry where the corresponding ResourceClaim uses "WaitForFirstConsumer" allocation mode.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceClaimSchedulingStatusArgs> ResourceClaims
        {
            get => _resourceClaims ?? (_resourceClaims = new InputList<Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceClaimSchedulingStatusArgs>());
            set => _resourceClaims = value;
        }

        public PodSchedulingContextStatusArgs()
        {
        }
        public static new PodSchedulingContextStatusArgs Empty => new PodSchedulingContextStatusArgs();
    }
}
