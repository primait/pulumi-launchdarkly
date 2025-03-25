// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Launchdarkly.Outputs
{

    [OutputType]
    public sealed class GetFeatureFlagEnvironmentTargetResult
    {
        /// <summary>
        /// List of `user` strings to target.
        /// </summary>
        public readonly ImmutableArray<string> Values;
        /// <summary>
        /// The index of the variation to serve if a user target value is matched.
        /// </summary>
        public readonly int Variation;

        [OutputConstructor]
        private GetFeatureFlagEnvironmentTargetResult(
            ImmutableArray<string> values,

            int variation)
        {
            Values = values;
            Variation = variation;
        }
    }
}
