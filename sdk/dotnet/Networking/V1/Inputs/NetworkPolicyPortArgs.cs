// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Networking.V1
{

    /// <summary>
    /// NetworkPolicyPort describes a port to allow traffic on
    /// </summary>
    public class NetworkPolicyPortArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// endPort indicates that the range of ports from port to endPort if set, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port.
        /// </summary>
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        /// <summary>
        /// port represents the port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.
        /// </summary>
        [Input("port")]
        public InputUnion<int, string>? Port { get; set; }

        /// <summary>
        /// protocol represents the protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public NetworkPolicyPortArgs()
        {
        }
        public static new NetworkPolicyPortArgs Empty => new NetworkPolicyPortArgs();
    }
}
