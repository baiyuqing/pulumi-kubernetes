// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Autoscaling.V1
{

    /// <summary>
    /// CrossVersionObjectReference contains enough information to let you identify the referred resource.
    /// </summary>
    [OutputType]
    public sealed class CrossVersionObjectReference
    {
        /// <summary>
        /// apiVersion is the API version of the referent
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private CrossVersionObjectReference(
            string apiVersion,

            string kind,

            string name)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Name = name;
        }
    }
}
