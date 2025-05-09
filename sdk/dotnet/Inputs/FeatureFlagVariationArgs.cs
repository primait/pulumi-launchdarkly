// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Launchdarkly.Inputs
{

    public sealed class FeatureFlagVariationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The feature flag's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The human-readable name of the feature flag.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The variation value. The value's type must correspond to the `variation_type` argument. For example: `variation_type = "boolean"` accepts only `true` or `false`. The `number` variation type accepts both floats and ints, but please note that any trailing zeroes on floats will be trimmed (i.e. `1.1` and `1.100` will both be converted to `1.1`).
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public FeatureFlagVariationArgs()
        {
        }
        public static new FeatureFlagVariationArgs Empty => new FeatureFlagVariationArgs();
    }
}
