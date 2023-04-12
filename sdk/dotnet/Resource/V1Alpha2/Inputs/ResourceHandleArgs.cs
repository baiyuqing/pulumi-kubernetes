// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2
{

    /// <summary>
    /// ResourceHandle holds opaque resource data for processing by a specific kubelet plugin.
    /// </summary>
    public class ResourceHandleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data contains the opaque data associated with this ResourceHandle. It is set by the controller component of the resource driver whose name matches the DriverName set in the ResourceClaimStatus this ResourceHandle is embedded in. It is set at allocation time and is intended for processing by the kubelet plugin whose name matches the DriverName set in this ResourceHandle.
        /// 
        /// The maximum size of this field is 16KiB. This may get increased in the future, but not reduced.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// DriverName specifies the name of the resource driver whose kubelet plugin should be invoked to process this ResourceHandle's data once it lands on a node. This may differ from the DriverName set in ResourceClaimStatus this ResourceHandle is embedded in.
        /// </summary>
        [Input("driverName")]
        public Input<string>? DriverName { get; set; }

        public ResourceHandleArgs()
        {
        }
        public static new ResourceHandleArgs Empty => new ResourceHandleArgs();
    }
}
