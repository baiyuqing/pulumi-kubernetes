// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// ClusterTrustBundleProjection describes how to select a set of ClusterTrustBundle objects and project their contents into the pod filesystem.
    /// </summary>
    [OutputType]
    public sealed class ClusterTrustBundleProjectionPatch
    {
        /// <summary>
        /// Select all ClusterTrustBundles that match this label selector.  Only has effect if signerName is set.  Mutually-exclusive with name.  If unset, interpreted as "match nothing".  If set but empty, interpreted as "match everything".
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelectorPatch LabelSelector;
        /// <summary>
        /// Select a single ClusterTrustBundle by object name.  Mutually-exclusive with signerName and labelSelector.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// If true, don't block pod startup if the referenced ClusterTrustBundle(s) aren't available.  If using name, then the named ClusterTrustBundle is allowed not to exist.  If using signerName, then the combination of signerName and labelSelector is allowed to match zero ClusterTrustBundles.
        /// </summary>
        public readonly bool Optional;
        /// <summary>
        /// Relative path from the volume root to write the bundle.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Select all ClusterTrustBundles that match this signer name. Mutually-exclusive with name.  The contents of all selected ClusterTrustBundles will be unified and deduplicated.
        /// </summary>
        public readonly string SignerName;

        [OutputConstructor]
        private ClusterTrustBundleProjectionPatch(
            Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelectorPatch labelSelector,

            string name,

            bool optional,

            string path,

            string signerName)
        {
            LabelSelector = labelSelector;
            Name = name;
            Optional = optional;
            Path = path;
            SignerName = signerName;
        }
    }
}
