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
    public sealed class GetFeatureFlagDefaultResult
    {
        /// <summary>
        /// The index of the variation the flag will default to in all new environments when off.
        /// </summary>
        public readonly int OffVariation;
        /// <summary>
        /// The index of the variation the flag will default to in all new environments when on.
        /// </summary>
        public readonly int OnVariation;

        [OutputConstructor]
        private GetFeatureFlagDefaultResult(
            int offVariation,

            int onVariation)
        {
            OffVariation = offVariation;
            OnVariation = onVariation;
        }
    }
}
