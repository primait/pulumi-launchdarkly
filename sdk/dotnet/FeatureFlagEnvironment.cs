// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Launchdarkly
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Launchdarkly = Pulumi.Launchdarkly;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // This example shows the use of prerequisites, targets, context targets, rules, and fallthrough for a feature flag environment
    ///     var numberFfEnv = new Launchdarkly.FeatureFlagEnvironment("number_ff_env", new()
    ///     {
    ///         FlagId = number.Id,
    ///         EnvKey = staging.Key,
    ///         On = true,
    ///         Prerequisites = new[]
    ///         {
    ///             new Launchdarkly.Inputs.FeatureFlagEnvironmentPrerequisiteArgs
    ///             {
    ///                 FlagKey = basic.Key,
    ///                 Variation = 0,
    ///             },
    ///         },
    ///         Targets = new[]
    ///         {
    ///             new Launchdarkly.Inputs.FeatureFlagEnvironmentTargetArgs
    ///             {
    ///                 Values = new[]
    ///                 {
    ///                     "user0",
    ///                 },
    ///                 Variation = 0,
    ///             },
    ///             new Launchdarkly.Inputs.FeatureFlagEnvironmentTargetArgs
    ///             {
    ///                 Values = new[]
    ///                 {
    ///                     "user1",
    ///                     "user2",
    ///                 },
    ///                 Variation = 1,
    ///             },
    ///         },
    ///         ContextTargets = new[]
    ///         {
    ///             new Launchdarkly.Inputs.FeatureFlagEnvironmentContextTargetArgs
    ///             {
    ///                 Values = new[]
    ///                 {
    ///                     "accountX",
    ///                 },
    ///                 Variation = 1,
    ///                 ContextKind = "account",
    ///             },
    ///         },
    ///         Rules = new[]
    ///         {
    ///             new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleArgs
    ///             {
    ///                 Description = "example targeting rule with two clauses",
    ///                 Clauses = new[]
    ///                 {
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "country",
    ///                         Op = "startsWith",
    ///                         Values = new[]
    ///                         {
    ///                             "aus",
    ///                             "de",
    ///                             "united",
    ///                         },
    ///                         Negate = false,
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "segmentMatch",
    ///                         Op = "segmentMatch",
    ///                         Values = new[]
    ///                         {
    ///                             example.Key,
    ///                         },
    ///                         Negate = false,
    ///                     },
    ///                 },
    ///                 Variation = 0,
    ///             },
    ///         },
    ///         Fallthrough = new Launchdarkly.Inputs.FeatureFlagEnvironmentFallthroughArgs
    ///         {
    ///             RolloutWeights = new[]
    ///             {
    ///                 60000,
    ///                 40000,
    ///                 0,
    ///             },
    ///             ContextKind = "account",
    ///             BucketBy = "accountId",
    ///         },
    ///         OffVariation = 2,
    ///     });
    /// 
    ///     // This example shows the minimum configuration required to create a feature flag environment
    ///     var basicFlagEnvironment = new Launchdarkly.FeatureFlagEnvironment("basic_flag_environment", new()
    ///     {
    ///         FlagId = basicFlag.Id,
    ///         EnvKey = "development",
    ///         On = true,
    ///         Fallthrough = new Launchdarkly.Inputs.FeatureFlagEnvironmentFallthroughArgs
    ///         {
    ///             Variation = 1,
    ///         },
    ///         OffVariation = 0,
    ///     });
    /// 
    ///     // This example shows a feature flag environment with a targeting rule that uses every clause operator
    ///     var bigFlagEnvironment = new Launchdarkly.FeatureFlagEnvironment("big_flag_environment", new()
    ///     {
    ///         FlagId = bigFlag.Id,
    ///         EnvKey = "development",
    ///         On = true,
    ///         Rules = new[]
    ///         {
    ///             new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleArgs
    ///             {
    ///                 Description = "Example targeting rule with every clause operator",
    ///                 Clauses = new[]
    ///                 {
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "username",
    ///                         Op = "in",
    ///                         Values = new[]
    ///                         {
    ///                             "henrietta powell",
    ///                             "wally waterbear",
    ///                         },
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "username",
    ///                         Op = "endsWith",
    ///                         Values = new[]
    ///                         {
    ///                             "powell",
    ///                             "waterbear",
    ///                         },
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "username",
    ///                         Op = "startsWith",
    ///                         Values = new[]
    ///                         {
    ///                             "henrietta",
    ///                             "wally",
    ///                         },
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "username",
    ///                         Op = "matches",
    ///                         Values = new[]
    ///                         {
    ///                             "henr*",
    ///                         },
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "username",
    ///                         Op = "contains",
    ///                         Values = new[]
    ///                         {
    ///                             "water",
    ///                         },
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "pageVisits",
    ///                         Op = "lessThan",
    ///                         Values = new[]
    ///                         {
    ///                             "100",
    ///                         },
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "pageVisits",
    ///                         Op = "lessThanOrEqual",
    ///                         Values = new[]
    ///                         {
    ///                             "100",
    ///                         },
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "pageVisits",
    ///                         Op = "greaterThan",
    ///                         Values = new[]
    ///                         {
    ///                             "100",
    ///                         },
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "pageVisits",
    ///                         Op = "greaterThanOrEqual",
    ///                         Values = new[]
    ///                         {
    ///                             "100",
    ///                         },
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "creationDate",
    ///                         Op = "before",
    ///                         Values = new[]
    ///                         {
    ///                             "2024-05-03T15:57:30Z",
    ///                         },
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "creationDate",
    ///                         Op = "after",
    ///                         Values = new[]
    ///                         {
    ///                             "2024-05-03T15:57:30Z",
    ///                         },
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "version",
    ///                         Op = "semVerEqual",
    ///                         Values = new[]
    ///                         {
    ///                             "1.0.0",
    ///                             "1.0.1",
    ///                         },
    ///                         ContextKind = "application",
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "version",
    ///                         Op = "semVerLessThan",
    ///                         Values = new[]
    ///                         {
    ///                             "1.0.0",
    ///                         },
    ///                         ContextKind = "application",
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "version",
    ///                         Op = "semVerGreaterThan",
    ///                         Values = new[]
    ///                         {
    ///                             "1.0.0",
    ///                         },
    ///                         ContextKind = "application",
    ///                     },
    ///                     new Launchdarkly.Inputs.FeatureFlagEnvironmentRuleClauseArgs
    ///                     {
    ///                         Attribute = "context",
    ///                         Op = "segmentMatch",
    ///                         Values = new[]
    ///                         {
    ///                             "test-segment",
    ///                         },
    ///                     },
    ///                 },
    ///                 RolloutWeights = new[]
    ///                 {
    ///                     40000,
    ///                     60000,
    ///                 },
    ///                 BucketBy = "country",
    ///                 ContextKind = "account",
    ///             },
    ///         },
    ///         Fallthrough = new Launchdarkly.Inputs.FeatureFlagEnvironmentFallthroughArgs
    ///         {
    ///             Variation = 1,
    ///         },
    ///         OffVariation = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LaunchDarkly feature flag environments can be imported using the resource's ID in the form `project_key/env_key/flag_key`
    /// 
    /// ```sh
    /// $ pulumi import launchdarkly:index/featureFlagEnvironment:FeatureFlagEnvironment example example-project/example-env/example-flag-key
    /// ```
    /// </summary>
    [LaunchdarklyResourceType("launchdarkly:index/featureFlagEnvironment:FeatureFlagEnvironment")]
    public partial class FeatureFlagEnvironment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The set of nested blocks describing the individual targets for non-user context kinds for each variation.
        /// </summary>
        [Output("contextTargets")]
        public Output<ImmutableArray<Outputs.FeatureFlagEnvironmentContextTarget>> ContextTargets { get; private set; } = null!;

        /// <summary>
        /// The environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Output("envKey")]
        public Output<string> EnvKey { get; private set; } = null!;

        /// <summary>
        /// Nested block describing the default variation to serve if no `prerequisites`, `target`, or `rules` apply.
        /// </summary>
        [Output("fallthrough")]
        public Output<Outputs.FeatureFlagEnvironmentFallthrough> Fallthrough { get; private set; } = null!;

        /// <summary>
        /// The feature flag's unique `id` in the format `project_key/flag_key`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Output("flagId")]
        public Output<string> FlagId { get; private set; } = null!;

        /// <summary>
        /// The index of the variation to serve if targeting is disabled.
        /// </summary>
        [Output("offVariation")]
        public Output<int> OffVariation { get; private set; } = null!;

        /// <summary>
        /// Whether targeting is enabled. Defaults to `false` if not set.
        /// </summary>
        [Output("on")]
        public Output<bool?> On { get; private set; } = null!;

        /// <summary>
        /// List of nested blocks describing prerequisite feature flags rules.
        /// </summary>
        [Output("prerequisites")]
        public Output<ImmutableArray<Outputs.FeatureFlagEnvironmentPrerequisite>> Prerequisites { get; private set; } = null!;

        /// <summary>
        /// List of logical targeting rules.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.FeatureFlagEnvironmentRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// Set of nested blocks describing the individual user targets for each variation.
        /// </summary>
        [Output("targets")]
        public Output<ImmutableArray<Outputs.FeatureFlagEnvironmentTarget>> Targets { get; private set; } = null!;

        /// <summary>
        /// Whether to send event data back to LaunchDarkly. Defaults to `false` if not set.
        /// </summary>
        [Output("trackEvents")]
        public Output<bool?> TrackEvents { get; private set; } = null!;


        /// <summary>
        /// Create a FeatureFlagEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FeatureFlagEnvironment(string name, FeatureFlagEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("launchdarkly:index/featureFlagEnvironment:FeatureFlagEnvironment", name, args ?? new FeatureFlagEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FeatureFlagEnvironment(string name, Input<string> id, FeatureFlagEnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("launchdarkly:index/featureFlagEnvironment:FeatureFlagEnvironment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/primait/pulumi-launchdarkly/releases/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FeatureFlagEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FeatureFlagEnvironment Get(string name, Input<string> id, FeatureFlagEnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new FeatureFlagEnvironment(name, id, state, options);
        }
    }

    public sealed class FeatureFlagEnvironmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("contextTargets")]
        private InputList<Inputs.FeatureFlagEnvironmentContextTargetArgs>? _contextTargets;

        /// <summary>
        /// The set of nested blocks describing the individual targets for non-user context kinds for each variation.
        /// </summary>
        public InputList<Inputs.FeatureFlagEnvironmentContextTargetArgs> ContextTargets
        {
            get => _contextTargets ?? (_contextTargets = new InputList<Inputs.FeatureFlagEnvironmentContextTargetArgs>());
            set => _contextTargets = value;
        }

        /// <summary>
        /// The environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("envKey", required: true)]
        public Input<string> EnvKey { get; set; } = null!;

        /// <summary>
        /// Nested block describing the default variation to serve if no `prerequisites`, `target`, or `rules` apply.
        /// </summary>
        [Input("fallthrough", required: true)]
        public Input<Inputs.FeatureFlagEnvironmentFallthroughArgs> Fallthrough { get; set; } = null!;

        /// <summary>
        /// The feature flag's unique `id` in the format `project_key/flag_key`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("flagId", required: true)]
        public Input<string> FlagId { get; set; } = null!;

        /// <summary>
        /// The index of the variation to serve if targeting is disabled.
        /// </summary>
        [Input("offVariation", required: true)]
        public Input<int> OffVariation { get; set; } = null!;

        /// <summary>
        /// Whether targeting is enabled. Defaults to `false` if not set.
        /// </summary>
        [Input("on")]
        public Input<bool>? On { get; set; }

        [Input("prerequisites")]
        private InputList<Inputs.FeatureFlagEnvironmentPrerequisiteArgs>? _prerequisites;

        /// <summary>
        /// List of nested blocks describing prerequisite feature flags rules.
        /// </summary>
        public InputList<Inputs.FeatureFlagEnvironmentPrerequisiteArgs> Prerequisites
        {
            get => _prerequisites ?? (_prerequisites = new InputList<Inputs.FeatureFlagEnvironmentPrerequisiteArgs>());
            set => _prerequisites = value;
        }

        [Input("rules")]
        private InputList<Inputs.FeatureFlagEnvironmentRuleArgs>? _rules;

        /// <summary>
        /// List of logical targeting rules.
        /// </summary>
        public InputList<Inputs.FeatureFlagEnvironmentRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FeatureFlagEnvironmentRuleArgs>());
            set => _rules = value;
        }

        [Input("targets")]
        private InputList<Inputs.FeatureFlagEnvironmentTargetArgs>? _targets;

        /// <summary>
        /// Set of nested blocks describing the individual user targets for each variation.
        /// </summary>
        public InputList<Inputs.FeatureFlagEnvironmentTargetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.FeatureFlagEnvironmentTargetArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// Whether to send event data back to LaunchDarkly. Defaults to `false` if not set.
        /// </summary>
        [Input("trackEvents")]
        public Input<bool>? TrackEvents { get; set; }

        public FeatureFlagEnvironmentArgs()
        {
        }
        public static new FeatureFlagEnvironmentArgs Empty => new FeatureFlagEnvironmentArgs();
    }

    public sealed class FeatureFlagEnvironmentState : global::Pulumi.ResourceArgs
    {
        [Input("contextTargets")]
        private InputList<Inputs.FeatureFlagEnvironmentContextTargetGetArgs>? _contextTargets;

        /// <summary>
        /// The set of nested blocks describing the individual targets for non-user context kinds for each variation.
        /// </summary>
        public InputList<Inputs.FeatureFlagEnvironmentContextTargetGetArgs> ContextTargets
        {
            get => _contextTargets ?? (_contextTargets = new InputList<Inputs.FeatureFlagEnvironmentContextTargetGetArgs>());
            set => _contextTargets = value;
        }

        /// <summary>
        /// The environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("envKey")]
        public Input<string>? EnvKey { get; set; }

        /// <summary>
        /// Nested block describing the default variation to serve if no `prerequisites`, `target`, or `rules` apply.
        /// </summary>
        [Input("fallthrough")]
        public Input<Inputs.FeatureFlagEnvironmentFallthroughGetArgs>? Fallthrough { get; set; }

        /// <summary>
        /// The feature flag's unique `id` in the format `project_key/flag_key`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("flagId")]
        public Input<string>? FlagId { get; set; }

        /// <summary>
        /// The index of the variation to serve if targeting is disabled.
        /// </summary>
        [Input("offVariation")]
        public Input<int>? OffVariation { get; set; }

        /// <summary>
        /// Whether targeting is enabled. Defaults to `false` if not set.
        /// </summary>
        [Input("on")]
        public Input<bool>? On { get; set; }

        [Input("prerequisites")]
        private InputList<Inputs.FeatureFlagEnvironmentPrerequisiteGetArgs>? _prerequisites;

        /// <summary>
        /// List of nested blocks describing prerequisite feature flags rules.
        /// </summary>
        public InputList<Inputs.FeatureFlagEnvironmentPrerequisiteGetArgs> Prerequisites
        {
            get => _prerequisites ?? (_prerequisites = new InputList<Inputs.FeatureFlagEnvironmentPrerequisiteGetArgs>());
            set => _prerequisites = value;
        }

        [Input("rules")]
        private InputList<Inputs.FeatureFlagEnvironmentRuleGetArgs>? _rules;

        /// <summary>
        /// List of logical targeting rules.
        /// </summary>
        public InputList<Inputs.FeatureFlagEnvironmentRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FeatureFlagEnvironmentRuleGetArgs>());
            set => _rules = value;
        }

        [Input("targets")]
        private InputList<Inputs.FeatureFlagEnvironmentTargetGetArgs>? _targets;

        /// <summary>
        /// Set of nested blocks describing the individual user targets for each variation.
        /// </summary>
        public InputList<Inputs.FeatureFlagEnvironmentTargetGetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.FeatureFlagEnvironmentTargetGetArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// Whether to send event data back to LaunchDarkly. Defaults to `false` if not set.
        /// </summary>
        [Input("trackEvents")]
        public Input<bool>? TrackEvents { get; set; }

        public FeatureFlagEnvironmentState()
        {
        }
        public static new FeatureFlagEnvironmentState Empty => new FeatureFlagEnvironmentState();
    }
}
