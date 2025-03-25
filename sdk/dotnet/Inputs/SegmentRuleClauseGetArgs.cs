// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Launchdarkly.Inputs
{

    public sealed class SegmentRuleClauseGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The user attribute to operate on
        /// </summary>
        [Input("attribute", required: true)]
        public Input<string> Attribute { get; set; } = null!;

        /// <summary>
        /// The context kind associated with this rule clause. If omitted, defaults to `user`.
        /// </summary>
        [Input("contextKind")]
        public Input<string>? ContextKind { get; set; }

        /// <summary>
        /// Whether to negate the rule clause.
        /// </summary>
        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        /// <summary>
        /// The operator associated with the rule clause. Available options are `in`, `endsWith`, `startsWith`, `matches`, `contains`, `lessThan`, `lessThanOrEqual`, `greaterThanOrEqual`, `before`, `after`, `segmentMatch`, `semVerEqual`, `semVerLessThan`, and `semVerGreaterThan`. Read LaunchDarkly's [Operators](https://docs.launchdarkly.com/sdk/concepts/flag-evaluation-rules#operators) documentation for more information.
        /// </summary>
        [Input("op", required: true)]
        public Input<string> Op { get; set; } = null!;

        /// <summary>
        /// The type for each of the clause's values. Available types are `boolean`, `string`, and `number`. If omitted, `value_type` defaults to `string`.
        /// </summary>
        [Input("valueType")]
        public Input<string>? ValueType { get; set; }

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// The list of values associated with the rule clause.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public SegmentRuleClauseGetArgs()
        {
        }
        public static new SegmentRuleClauseGetArgs Empty => new SegmentRuleClauseGetArgs();
    }
}
