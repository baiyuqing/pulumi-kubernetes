// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Networking.V1
{

    /// <summary>
    /// NetworkPolicyStatus describes the current state of the NetworkPolicy.
    /// </summary>
    [OutputType]
    public sealed class NetworkPolicyStatusPatch
    {
        /// <summary>
        /// conditions holds an array of metav1.Condition that describe the state of the NetworkPolicy. Current service state
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ConditionPatch> Conditions;

        [OutputConstructor]
        private NetworkPolicyStatusPatch(ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ConditionPatch> conditions)
        {
            Conditions = conditions;
        }
    }
}
