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
    /// TokenRequest contains parameters of a service account token.
    /// </summary>
    [OutputType]
    public sealed class TokenRequestPatch
    {
        /// <summary>
        /// audience is the intended audience of the token in "TokenRequestSpec". It will default to the audiences of kube apiserver.
        /// </summary>
        public readonly string Audience;
        /// <summary>
        /// expirationSeconds is the duration of validity of the token in "TokenRequestSpec". It has the same default value of "ExpirationSeconds" in "TokenRequestSpec".
        /// </summary>
        public readonly int ExpirationSeconds;

        [OutputConstructor]
        private TokenRequestPatch(
            string audience,

            int expirationSeconds)
        {
            Audience = audience;
            ExpirationSeconds = expirationSeconds;
        }
    }
}
