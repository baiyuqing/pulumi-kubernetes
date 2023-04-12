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
    /// NetworkPolicyEgressRule describes a particular set of traffic that is allowed out of pods matched by a NetworkPolicySpec's podSelector. The traffic must match both ports and to. This type is beta-level in 1.8
    /// </summary>
    [OutputType]
    public sealed class NetworkPolicyEgressRulePatch
    {
        /// <summary>
        /// ports is a list of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Networking.V1.NetworkPolicyPortPatch> Ports;
        /// <summary>
        /// to is a list of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Networking.V1.NetworkPolicyPeerPatch> To;

        [OutputConstructor]
        private NetworkPolicyEgressRulePatch(
            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Networking.V1.NetworkPolicyPortPatch> ports,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Networking.V1.NetworkPolicyPeerPatch> to)
        {
            Ports = ports;
            To = to;
        }
    }
}
