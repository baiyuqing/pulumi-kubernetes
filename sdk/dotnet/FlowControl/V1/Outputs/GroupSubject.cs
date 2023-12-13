// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.FlowControl.V1
{

    /// <summary>
    /// GroupSubject holds detailed information for group-kind subject.
    /// </summary>
    [OutputType]
    public sealed class GroupSubject
    {
        /// <summary>
        /// name is the user group that matches, or "*" to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GroupSubject(string name)
        {
            Name = name;
        }
    }
}
