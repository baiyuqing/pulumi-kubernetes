// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Resource.V1Alpha2
{

    /// <summary>
    /// ResourceClaimSchedulingStatus contains information about one particular ResourceClaim with "WaitForFirstConsumer" allocation mode.
    /// </summary>
    [OutputType]
    public sealed class ResourceClaimSchedulingStatus
    {
        /// <summary>
        /// Name matches the pod.spec.resourceClaims[*].Name field.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// UnsuitableNodes lists nodes that the ResourceClaim cannot be allocated for.
        /// 
        /// The size of this field is limited to 128, the same as for PodSchedulingSpec.PotentialNodes. This may get increased in the future, but not reduced.
        /// </summary>
        public readonly ImmutableArray<string> UnsuitableNodes;

        [OutputConstructor]
        private ResourceClaimSchedulingStatus(
            string name,

            ImmutableArray<string> unsuitableNodes)
        {
            Name = name;
            UnsuitableNodes = unsuitableNodes;
        }
    }
}
