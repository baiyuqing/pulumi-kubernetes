// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Batch.V1
{

    /// <summary>
    /// CronJobSpec describes how the job execution will look like and when it will actually run.
    /// </summary>
    [OutputType]
    public sealed class CronJobSpec
    {
        /// <summary>
        /// Specifies how to treat concurrent executions of a Job. Valid values are:
        /// 
        /// - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
        /// </summary>
        public readonly string ConcurrencyPolicy;
        /// <summary>
        /// The number of failed finished jobs to retain. Value must be non-negative integer. Defaults to 1.
        /// </summary>
        public readonly int FailedJobsHistoryLimit;
        /// <summary>
        /// Specifies the job that will be created when executing a CronJob.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobTemplateSpec JobTemplate;
        /// <summary>
        /// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
        /// </summary>
        public readonly string Schedule;
        /// <summary>
        /// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
        /// </summary>
        public readonly int StartingDeadlineSeconds;
        /// <summary>
        /// The number of successful finished jobs to retain. Value must be non-negative integer. Defaults to 3.
        /// </summary>
        public readonly int SuccessfulJobsHistoryLimit;
        /// <summary>
        /// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
        /// </summary>
        public readonly bool Suspend;
        /// <summary>
        /// The time zone name for the given schedule, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones. If not specified, this will default to the time zone of the kube-controller-manager process. The set of valid time zone names and the time zone offset is loaded from the system-wide time zone database by the API server during CronJob validation and the controller manager during execution. If no system-wide time zone database can be found a bundled version of the database is used instead. If the time zone name becomes invalid during the lifetime of a CronJob or due to a change in host configuration, the controller will stop creating new new Jobs and will create a system event with the reason UnknownTimeZone. More information can be found in https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#time-zones
        /// </summary>
        public readonly string TimeZone;

        [OutputConstructor]
        private CronJobSpec(
            string concurrencyPolicy,

            int failedJobsHistoryLimit,

            Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobTemplateSpec jobTemplate,

            string schedule,

            int startingDeadlineSeconds,

            int successfulJobsHistoryLimit,

            bool suspend,

            string timeZone)
        {
            ConcurrencyPolicy = concurrencyPolicy;
            FailedJobsHistoryLimit = failedJobsHistoryLimit;
            JobTemplate = jobTemplate;
            Schedule = schedule;
            StartingDeadlineSeconds = startingDeadlineSeconds;
            SuccessfulJobsHistoryLimit = successfulJobsHistoryLimit;
            Suspend = suspend;
            TimeZone = timeZone;
        }
    }
}
