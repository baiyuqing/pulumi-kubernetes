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
    /// IngressLoadBalancerIngress represents the status of a load-balancer ingress point.
    /// </summary>
    public class IngressLoadBalancerIngressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// hostname is set for load-balancer ingress points that are DNS based.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// ip is set for load-balancer ingress points that are IP based.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("ports")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressPortStatusArgs>? _ports;

        /// <summary>
        /// ports provides information about the ports exposed by this LoadBalancer.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressPortStatusArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Pulumi.Kubernetes.Types.Inputs.Networking.V1.IngressPortStatusArgs>());
            set => _ports = value;
        }

        public IngressLoadBalancerIngressArgs()
        {
        }
        public static new IngressLoadBalancerIngressArgs Empty => new IngressLoadBalancerIngressArgs();
    }
}
