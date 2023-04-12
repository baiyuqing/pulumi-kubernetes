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
    /// IngressStatus describe the current state of the Ingress.
    /// </summary>
    [OutputType]
    public sealed class IngressStatus
    {
        /// <summary>
        /// loadBalancer contains the current status of the load-balancer.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Networking.V1.IngressLoadBalancerStatus LoadBalancer;

        [OutputConstructor]
        private IngressStatus(Pulumi.Kubernetes.Types.Outputs.Networking.V1.IngressLoadBalancerStatus loadBalancer)
        {
            LoadBalancer = loadBalancer;
        }
    }
}
