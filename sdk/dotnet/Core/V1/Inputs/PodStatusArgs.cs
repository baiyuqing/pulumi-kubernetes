// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    /// <summary>
    /// PodStatus represents information about the status of a pod. Status may trail the actual state of a system, especially if the node that hosts the pod cannot contact the control plane.
    /// </summary>
    public class PodStatusArgs : global::Pulumi.ResourceArgs
    {
        [Input("conditions")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.PodConditionArgs>? _conditions;

        /// <summary>
        /// Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.PodConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.PodConditionArgs>());
            set => _conditions = value;
        }

        [Input("containerStatuses")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerStatusArgs>? _containerStatuses;

        /// <summary>
        /// The list has one entry per container in the manifest. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerStatusArgs> ContainerStatuses
        {
            get => _containerStatuses ?? (_containerStatuses = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerStatusArgs>());
            set => _containerStatuses = value;
        }

        [Input("ephemeralContainerStatuses")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerStatusArgs>? _ephemeralContainerStatuses;

        /// <summary>
        /// Status for any ephemeral containers that have run in this pod.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerStatusArgs> EphemeralContainerStatuses
        {
            get => _ephemeralContainerStatuses ?? (_ephemeralContainerStatuses = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerStatusArgs>());
            set => _ephemeralContainerStatuses = value;
        }

        /// <summary>
        /// IP address of the host to which the pod is assigned. Empty if not yet scheduled.
        /// </summary>
        [Input("hostIP")]
        public Input<string>? HostIP { get; set; }

        [Input("initContainerStatuses")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerStatusArgs>? _initContainerStatuses;

        /// <summary>
        /// The list has one entry per init container in the manifest. The most recent successful init container will have ready = true, the most recently started container will have startTime set. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerStatusArgs> InitContainerStatuses
        {
            get => _initContainerStatuses ?? (_initContainerStatuses = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerStatusArgs>());
            set => _initContainerStatuses = value;
        }

        /// <summary>
        /// A human readable message indicating details about why the pod is in this condition.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to give the resources on this node to a higher priority pod that is created after preemption. As a result, this field may be different than PodSpec.nodeName when the pod is scheduled.
        /// </summary>
        [Input("nominatedNodeName")]
        public Input<string>? NominatedNodeName { get; set; }

        /// <summary>
        /// The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. The conditions array, the reason and message fields, and the individual container status arrays contain more detail about the pod's status. There are five possible phase values:
        /// 
        /// Pending: The pod has been accepted by the Kubernetes system, but one or more of the container images has not been created. This includes time before being scheduled as well as time spent downloading images over the network, which could take a while. Running: The pod has been bound to a node, and all of the containers have been created. At least one container is still running, or is in the process of starting or restarting. Succeeded: All containers in the pod have terminated in success, and will not be restarted. Failed: All containers in the pod have terminated, and at least one container has terminated in failure. The container either exited with non-zero status or was terminated by the system. Unknown: For some reason the state of the pod could not be obtained, typically due to an error in communicating with the host of the pod.
        /// 
        /// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase
        /// </summary>
        [Input("phase")]
        public Input<string>? Phase { get; set; }

        /// <summary>
        /// IP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated.
        /// </summary>
        [Input("podIP")]
        public Input<string>? PodIP { get; set; }

        [Input("podIPs")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.PodIPArgs>? _podIPs;

        /// <summary>
        /// podIPs holds the IP addresses allocated to the pod. If this field is specified, the 0th entry must match the podIP field. Pods may be allocated at most 1 value for each of IPv4 and IPv6. This list is empty if no IPs have been allocated yet.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.PodIPArgs> PodIPs
        {
            get => _podIPs ?? (_podIPs = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.PodIPArgs>());
            set => _podIPs = value;
        }

        /// <summary>
        /// The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-qos/#quality-of-service-classes
        /// </summary>
        [Input("qosClass")]
        public Input<string>? QosClass { get; set; }

        /// <summary>
        /// A brief CamelCase message indicating details about why the pod is in this state. e.g. 'Evicted'
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// Status of resources resize desired for pod's containers. It is empty if no resources resize is pending. Any changes to container resources will automatically set this to "Proposed"
        /// </summary>
        [Input("resize")]
        public Input<string>? Resize { get; set; }

        /// <summary>
        /// RFC 3339 date and time at which the object was acknowledged by the Kubelet. This is before the Kubelet pulled the container image(s) for the pod.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public PodStatusArgs()
        {
        }
        public static new PodStatusArgs Empty => new PodStatusArgs();
    }
}
