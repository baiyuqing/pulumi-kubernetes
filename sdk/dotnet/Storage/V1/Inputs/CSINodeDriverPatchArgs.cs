// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Storage.V1
{

    /// <summary>
    /// CSINodeDriver holds information about the specification of one CSI driver installed on a node
    /// </summary>
    public class CSINodeDriverPatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// allocatable represents the volume resources of a node that are available for scheduling. This field is beta.
        /// </summary>
        [Input("allocatable")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Storage.V1.VolumeNodeResourcesPatchArgs>? Allocatable { get; set; }

        /// <summary>
        /// name represents the name of the CSI driver that this object refers to. This MUST be the same name returned by the CSI GetPluginName() call for that driver.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// nodeID of the node from the driver point of view. This field enables Kubernetes to communicate with storage systems that do not share the same nomenclature for nodes. For example, Kubernetes may refer to a given node as "node1", but the storage system may refer to the same node as "nodeA". When Kubernetes issues a command to the storage system to attach a volume to a specific node, it can use this field to refer to the node name using the ID that the storage system will understand, e.g. "nodeA" instead of "node1". This field is required.
        /// </summary>
        [Input("nodeID")]
        public Input<string>? NodeID { get; set; }

        [Input("topologyKeys")]
        private InputList<string>? _topologyKeys;

        /// <summary>
        /// topologyKeys is the list of keys supported by the driver. When a driver is initialized on a cluster, it provides a set of topology keys that it understands (e.g. "company.com/zone", "company.com/region"). When a driver is initialized on a node, it provides the same topology keys along with values. Kubelet will expose these topology keys as labels on its own node object. When Kubernetes does topology aware provisioning, it can use this list to determine which labels it should retrieve from the node object and pass back to the driver. It is possible for different nodes to use different topology keys. This can be empty if driver does not support topology.
        /// </summary>
        public InputList<string> TopologyKeys
        {
            get => _topologyKeys ?? (_topologyKeys = new InputList<string>());
            set => _topologyKeys = value;
        }

        public CSINodeDriverPatchArgs()
        {
        }
        public static new CSINodeDriverPatchArgs Empty => new CSINodeDriverPatchArgs();
    }
}
