// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Discovery.V1
{

    /// <summary>
    /// EndpointPort represents a Port used by an EndpointSlice
    /// </summary>
    public class EndpointPortArgs : global::Pulumi.ResourceArgs
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
        [Input("appProtocol")]
        public Input<string>? AppProtocol { get; set; }

        /// <summary>
        /// name represents the name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// port represents the port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// protocol represents the IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public EndpointPortArgs()
        {
        }
        public static new EndpointPortArgs Empty => new EndpointPortArgs();
    }
}
