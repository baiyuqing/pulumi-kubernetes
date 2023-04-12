// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Coordination.V1
{

    /// <summary>
    /// LeaseSpec is a specification of a Lease.
    /// </summary>
    [OutputType]
    public sealed class LeaseSpecPatch
    {
        /// <summary>
        /// acquireTime is a time when the current lease was acquired.
        /// </summary>
        public readonly string AcquireTime;
        /// <summary>
        /// holderIdentity contains the identity of the holder of a current lease.
        /// </summary>
        public readonly string HolderIdentity;
        /// <summary>
        /// leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed renewTime.
        /// </summary>
        public readonly int LeaseDurationSeconds;
        /// <summary>
        /// leaseTransitions is the number of transitions of a lease between holders.
        /// </summary>
        public readonly int LeaseTransitions;
        /// <summary>
        /// renewTime is a time when the current holder of a lease has last updated the lease.
        /// </summary>
        public readonly string RenewTime;

        [OutputConstructor]
        private LeaseSpecPatch(
            string acquireTime,

            string holderIdentity,

            int leaseDurationSeconds,

            int leaseTransitions,

            string renewTime)
        {
            AcquireTime = acquireTime;
            HolderIdentity = holderIdentity;
            LeaseDurationSeconds = leaseDurationSeconds;
            LeaseTransitions = leaseTransitions;
            RenewTime = renewTime;
        }
    }
}
