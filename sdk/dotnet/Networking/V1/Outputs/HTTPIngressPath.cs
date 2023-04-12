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
    /// HTTPIngressPath associates a path with a backend. Incoming urls matching the path are forwarded to the backend.
    /// </summary>
    [OutputType]
    public sealed class HTTPIngressPath
    {
        /// <summary>
        /// backend defines the referenced service endpoint to which the traffic will be forwarded to.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Networking.V1.IngressBackend Backend;
        /// <summary>
        /// path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value "Exact" or "Prefix".
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is
        ///   done on a path element by element basis. A path element refers is the
        ///   list of labels in the path split by the '/' separator. A request is a
        ///   match for path p if every p is an element-wise prefix of p of the
        ///   request path. Note that if the last element of the path is a substring
        ///   of the last element in request path, it is not a match (e.g. /foo/bar
        ///   matches /foo/bar/baz, but does not match /foo/barbaz).
        /// * ImplementationSpecific: Interpretation of the Path matching is up to
        ///   the IngressClass. Implementations can treat this as a separate PathType
        ///   or treat it identically to Prefix or Exact path types.
        /// Implementations are required to support all path types.
        /// </summary>
        public readonly string PathType;

        [OutputConstructor]
        private HTTPIngressPath(
            Pulumi.Kubernetes.Types.Outputs.Networking.V1.IngressBackend backend,

            string path,

            string pathType)
        {
            Backend = backend;
            Path = path;
            PathType = pathType;
        }
    }
}
