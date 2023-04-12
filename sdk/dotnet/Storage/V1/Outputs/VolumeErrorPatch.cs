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
    /// VolumeError captures an error encountered during a volume operation.
    /// </summary>
    [OutputType]
    public sealed class VolumeErrorPatch
    {
        /// <summary>
        /// message represents the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// time represents the time the error was encountered.
        /// </summary>
        public readonly string Time;

        [OutputConstructor]
        private VolumeErrorPatch(
            string message,

            string time)
        {
            Message = message;
            Time = time;
        }
    }
}
