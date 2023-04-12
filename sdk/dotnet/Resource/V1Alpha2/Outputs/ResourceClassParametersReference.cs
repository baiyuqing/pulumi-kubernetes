// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Resource.V1Alpha2
{

    /// <summary>
    /// ResourceClassParametersReference contains enough information to let you locate the parameters for a ResourceClass.
    /// </summary>
    [OutputType]
    public sealed class ResourceClassParametersReference
    {
        /// <summary>
        /// APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
        /// </summary>
        public readonly string ApiGroup;
        /// <summary>
        /// Kind is the type of resource being referenced. This is the same value as in the parameter object's metadata.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name is the name of resource being referenced.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Namespace that contains the referenced resource. Must be empty for cluster-scoped resources and non-empty for namespaced resources.
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private ResourceClassParametersReference(
            string apiGroup,

            string kind,

            string name,

            string @namespace)
        {
            ApiGroup = apiGroup;
            Kind = kind;
            Name = name;
            Namespace = @namespace;
        }
    }
}
