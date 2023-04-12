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
    /// IngressSpec describes the Ingress the user wishes to exist.
    /// </summary>
    public class IngressSpecPatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// defaultBackend is the backend that should handle requests that don't match any rule. If Rules are not specified, DefaultBackend must be specified. If DefaultBackend is not set, the handling of requests that do not match any of the rules will be up to the Ingress controller.
        /// </summary>
        [Input("defaultBackend")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressBackendPatchArgs>? DefaultBackend { get; set; }

        /// <summary>
        /// ingressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -&gt; IngressClass -&gt; Ingress resource). Although the `kubernetes.io/ingress.class` annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.
        /// </summary>
        [Input("ingressClassName")]
        public Input<string>? IngressClassName { get; set; }

        [Input("rules")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressRulePatchArgs>? _rules;

        /// <summary>
        /// rules is a list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressRulePatchArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressRulePatchArgs>());
            set => _rules = value;
        }

        [Input("tls")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressTLSPatchArgs>? _tls;

        /// <summary>
        /// tls represents the TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressTLSPatchArgs> Tls
        {
            get => _tls ?? (_tls = new InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressTLSPatchArgs>());
            set => _tls = value;
        }

        public IngressSpecPatchArgs()
        {
        }
        public static new IngressSpecPatchArgs Empty => new IngressSpecPatchArgs();
    }
}
