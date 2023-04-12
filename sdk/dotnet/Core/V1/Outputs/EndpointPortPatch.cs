// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// EndpointPort is a tuple that describes a single port.
    /// </summary>
    [OutputType]
    public sealed class EndpointPortPatch
    {
        /// <summary>
        /// The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:
        /// 
        /// * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).
        /// 
        /// * Kubernetes-defined prefixed names:
        ///   * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540
        /// 
        /// * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
        /// </summary>
        public readonly string AppProtocol;
        /// <summary>
        /// The name of this port.  This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The port number of the endpoint.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private EndpointPortPatch(
            string appProtocol,

            string name,

            int port,

            string protocol)
        {
            AppProtocol = appProtocol;
            Name = name;
            Port = port;
            Protocol = protocol;
        }
    }
}
