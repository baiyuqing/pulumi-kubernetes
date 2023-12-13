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
    /// SleepAction describes a "sleep" action.
    /// </summary>
    [OutputType]
    public sealed class SleepAction
    {
        /// <summary>
        /// Seconds is the number of seconds to sleep.
        /// </summary>
        public readonly int Seconds;

        [OutputConstructor]
        private SleepAction(int seconds)
        {
            Seconds = seconds;
        }
    }
}
