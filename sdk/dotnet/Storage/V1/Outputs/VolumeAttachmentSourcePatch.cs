// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Storage.V1
{

    /// <summary>
    /// VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
    /// </summary>
    [OutputType]
    public sealed class VolumeAttachmentSourcePatch
    {
        /// <summary>
        /// inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is beta-level and is only honored by servers that enabled the CSIMigration feature.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.PersistentVolumeSpecPatch InlineVolumeSpec;
        /// <summary>
        /// persistentVolumeName represents the name of the persistent volume to attach.
        /// </summary>
        public readonly string PersistentVolumeName;

        [OutputConstructor]
        private VolumeAttachmentSourcePatch(
            Pulumi.Kubernetes.Types.Outputs.Core.V1.PersistentVolumeSpecPatch inlineVolumeSpec,

            string persistentVolumeName)
        {
            InlineVolumeSpec = inlineVolumeSpec;
            PersistentVolumeName = persistentVolumeName;
        }
    }
}
