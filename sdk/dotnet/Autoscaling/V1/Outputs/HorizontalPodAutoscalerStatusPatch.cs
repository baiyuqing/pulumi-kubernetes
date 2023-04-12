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
    /// current status of a horizontal pod autoscaler
    /// </summary>
    [OutputType]
    public sealed class HorizontalPodAutoscalerStatusPatch
    {
        /// <summary>
        /// currentCPUUtilizationPercentage is the current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
        /// </summary>
        public readonly int CurrentCPUUtilizationPercentage;
        /// <summary>
        /// currentReplicas is the current number of replicas of pods managed by this autoscaler.
        /// </summary>
        public readonly int CurrentReplicas;
        /// <summary>
        /// desiredReplicas is the  desired number of replicas of pods managed by this autoscaler.
        /// </summary>
        public readonly int DesiredReplicas;
        /// <summary>
        /// lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
        /// </summary>
        public readonly string LastScaleTime;
        /// <summary>
        /// observedGeneration is the most recent generation observed by this autoscaler.
        /// </summary>
        public readonly int ObservedGeneration;

        [OutputConstructor]
        private HorizontalPodAutoscalerStatusPatch(
            int currentCPUUtilizationPercentage,

            int currentReplicas,

            int desiredReplicas,

            string lastScaleTime,

            int observedGeneration)
        {
            CurrentCPUUtilizationPercentage = currentCPUUtilizationPercentage;
            CurrentReplicas = currentReplicas;
            DesiredReplicas = desiredReplicas;
            LastScaleTime = lastScaleTime;
            ObservedGeneration = observedGeneration;
        }
    }
}
