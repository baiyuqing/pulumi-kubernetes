// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.AdmissionRegistration.V1Alpha1
{

    /// <summary>
    /// ExpressionWarning is a warning information that targets a specific expression.
    /// </summary>
    [OutputType]
    public sealed class ExpressionWarningPatch
    {
        /// <summary>
        /// The path to the field that refers the expression. For example, the reference to the expression of the first item of validations is "spec.validations[0].expression"
        /// </summary>
        public readonly string FieldRef;
        /// <summary>
        /// The content of type checking information in a human-readable form. Each line of the warning contains the type that the expression is checked against, followed by the type check error from the compiler.
        /// </summary>
        public readonly string Warning;

        [OutputConstructor]
        private ExpressionWarningPatch(
            string fieldRef,

            string warning)
        {
            FieldRef = fieldRef;
            Warning = warning;
        }
    }
}
