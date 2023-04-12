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
    /// IngressBackend describes all endpoints for a given service and port.
    /// </summary>
    [OutputType]
    public sealed class IngressBackendPatch
    {
        /// <summary>
        /// resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with "Service".
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.TypedLocalObjectReferencePatch Resource;
        /// <summary>
        /// service references a service as a backend. This is a mutually exclusive setting with "Resource".
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Networking.V1.IngressServiceBackendPatch Service;

        [OutputConstructor]
        private IngressBackendPatch(
            Pulumi.Kubernetes.Types.Outputs.Core.V1.TypedLocalObjectReferencePatch resource,

            Pulumi.Kubernetes.Types.Outputs.Networking.V1.IngressServiceBackendPatch service)
        {
            Resource = resource;
            Service = service;
        }
    }
}
