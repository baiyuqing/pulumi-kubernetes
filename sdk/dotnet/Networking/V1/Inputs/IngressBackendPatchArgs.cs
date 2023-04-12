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
    /// IngressBackend describes all endpoints for a given service and port.
    /// </summary>
    public class IngressBackendPatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with "Service".
        /// </summary>
        [Input("resource")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.TypedLocalObjectReferencePatchArgs>? Resource { get; set; }

        /// <summary>
        /// service references a service as a backend. This is a mutually exclusive setting with "Resource".
        /// </summary>
        [Input("service")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressServiceBackendPatchArgs>? Service { get; set; }

        public IngressBackendPatchArgs()
        {
        }
        public static new IngressBackendPatchArgs Empty => new IngressBackendPatchArgs();
    }
}
