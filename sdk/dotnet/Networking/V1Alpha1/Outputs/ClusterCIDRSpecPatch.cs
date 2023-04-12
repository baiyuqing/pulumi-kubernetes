// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Networking.V1Alpha1
{

    /// <summary>
    /// ClusterCIDRSpec defines the desired state of ClusterCIDR.
    /// </summary>
    [OutputType]
    public sealed class ClusterCIDRSpecPatch
    {
        /// <summary>
        /// ipv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        /// </summary>
        public readonly string Ipv4;
        /// <summary>
        /// ipv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        /// </summary>
        public readonly string Ipv6;
        /// <summary>
        /// nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeSelectorPatch NodeSelector;
        /// <summary>
        /// perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
        /// </summary>
        public readonly int PerNodeHostBits;

        [OutputConstructor]
        private ClusterCIDRSpecPatch(
            string ipv4,

            string ipv6,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeSelectorPatch nodeSelector,

            int perNodeHostBits)
        {
            Ipv4 = ipv4;
            Ipv6 = ipv6;
            NodeSelector = nodeSelector;
            PerNodeHostBits = perNodeHostBits;
        }
    }
}
