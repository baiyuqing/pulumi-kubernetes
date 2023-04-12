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
    /// IngressLoadBalancerStatus represents the status of a load-balancer.
    /// </summary>
    public class IngressLoadBalancerStatusArgs : global::Pulumi.ResourceArgs
    {
        [Input("ingress")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressLoadBalancerIngressArgs>? _ingress;

        /// <summary>
        /// ingress is a list containing ingress points for the load-balancer.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressLoadBalancerIngressArgs> Ingress
        {
            get => _ingress ?? (_ingress = new InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressLoadBalancerIngressArgs>());
            set => _ingress = value;
        }

        public IngressLoadBalancerStatusArgs()
        {
        }
        public static new IngressLoadBalancerStatusArgs Empty => new IngressLoadBalancerStatusArgs();
    }
}
