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
    /// Provides a LaunchDarkly metric resource.
    /// 
    /// This resource allows you to create and manage metrics within your LaunchDarkly organization.
    /// 
    /// To learn more about metrics and experimentation, read [Experimentation Documentation](https://docs.launchdarkly.com/home/experimentation).
    /// 
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
    ///     var example = new Launchdarkly.Metric("example", new()
    ///     {
    ///         ProjectKey = exampleLaunchdarklyProject.Key,
    ///         Key = "example-metric",
    ///         Name = "Example Metric",
    ///         Description = "Metric description.",
    ///         Kind = "pageview",
    ///         Tags = new[]
    ///         {
    ///             "example",
    ///         },
    ///         Urls = new[]
    ///         {
    ///             new Launchdarkly.Inputs.MetricUrlArgs
    ///             {
    ///                 Kind = "substring",
    ///                 Substring = "foo",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LaunchDarkly metrics can be imported using the metric's ID in the form `project_key/metric_key`
    /// 
    /// ```sh
    /// $ pulumi import launchdarkly:index/metric:Metric example example-project/example-metric-key
    /// ```
    /// </summary>
    [LaunchdarklyResourceType("launchdarkly:index/metric:Metric")]
    public partial class Metric : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The method for analyzing metric events. Available choices are `mean` and `percentile`.
        /// </summary>
        [Output("analysisType")]
        public Output<string?> AnalysisType { get; private set; } = null!;

        /// <summary>
        /// The description of the metric's purpose.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The event key for your metric (if custom metric)
        /// </summary>
        [Output("eventKey")]
        public Output<string?> EventKey { get; private set; } = null!;

        /// <summary>
        /// Include units that did not send any events and set their value to 0.
        /// </summary>
        [Output("includeUnitsWithoutEvents")]
        public Output<bool> IncludeUnitsWithoutEvents { get; private set; } = null!;

        /// <summary>
        /// Ignored. All metrics are considered active.
        /// </summary>
        [Output("isActive")]
        public Output<bool> IsActive { get; private set; } = null!;

        /// <summary>
        /// Whether a `custom` metric is a numeric metric or not.
        /// </summary>
        [Output("isNumeric")]
        public Output<bool?> IsNumeric { get; private set; } = null!;

        /// <summary>
        /// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The metric type. Available choices are `click`, `custom`, and `pageview`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        [Output("maintainerId")]
        public Output<string> MaintainerId { get; private set; } = null!;

        /// <summary>
        /// The human-friendly name for the metric.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when analysis_type is percentile.
        /// </summary>
        [Output("percentileValue")]
        public Output<int?> PercentileValue { get; private set; } = null!;

        /// <summary>
        /// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Output("projectKey")]
        public Output<string> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// A set of one or more context kinds that this metric can measure events from. Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
        /// </summary>
        [Output("randomizationUnits")]
        public Output<ImmutableArray<string>> RandomizationUnits { get; private set; } = null!;

        /// <summary>
        /// The CSS selector for your metric (if click metric)
        /// </summary>
        [Output("selector")]
        public Output<string?> Selector { get; private set; } = null!;

        /// <summary>
        /// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
        /// </summary>
        [Output("successCriteria")]
        public Output<string> SuccessCriteria { get; private set; } = null!;

        /// <summary>
        /// Tags associated with your resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// (Required for kind `custom`) The unit for numeric `custom` metrics.
        /// </summary>
        [Output("unit")]
        public Output<string?> Unit { get; private set; } = null!;

        /// <summary>
        /// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
        /// </summary>
        [Output("unitAggregationType")]
        public Output<string?> UnitAggregationType { get; private set; } = null!;

        /// <summary>
        /// List of nested `url` blocks describing URLs that you want to associate with the metric.
        /// </summary>
        [Output("urls")]
        public Output<ImmutableArray<Outputs.MetricUrl>> Urls { get; private set; } = null!;

        /// <summary>
        /// Version of the metric
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Metric resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Metric(string name, MetricArgs args, CustomResourceOptions? options = null)
            : base("launchdarkly:index/metric:Metric", name, args ?? new MetricArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Metric(string name, Input<string> id, MetricState? state = null, CustomResourceOptions? options = null)
            : base("launchdarkly:index/metric:Metric", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/primait/pulumi-launchdarkly",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Metric resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Metric Get(string name, Input<string> id, MetricState? state = null, CustomResourceOptions? options = null)
        {
            return new Metric(name, id, state, options);
        }
    }

    public sealed class MetricArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The method for analyzing metric events. Available choices are `mean` and `percentile`.
        /// </summary>
        [Input("analysisType")]
        public Input<string>? AnalysisType { get; set; }

        /// <summary>
        /// The description of the metric's purpose.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The event key for your metric (if custom metric)
        /// </summary>
        [Input("eventKey")]
        public Input<string>? EventKey { get; set; }

        /// <summary>
        /// Include units that did not send any events and set their value to 0.
        /// </summary>
        [Input("includeUnitsWithoutEvents")]
        public Input<bool>? IncludeUnitsWithoutEvents { get; set; }

        /// <summary>
        /// Ignored. All metrics are considered active.
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Whether a `custom` metric is a numeric metric or not.
        /// </summary>
        [Input("isNumeric")]
        public Input<bool>? IsNumeric { get; set; }

        /// <summary>
        /// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The metric type. Available choices are `click`, `custom`, and `pageview`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        [Input("maintainerId")]
        public Input<string>? MaintainerId { get; set; }

        /// <summary>
        /// The human-friendly name for the metric.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when analysis_type is percentile.
        /// </summary>
        [Input("percentileValue")]
        public Input<int>? PercentileValue { get; set; }

        /// <summary>
        /// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("projectKey", required: true)]
        public Input<string> ProjectKey { get; set; } = null!;

        [Input("randomizationUnits")]
        private InputList<string>? _randomizationUnits;

        /// <summary>
        /// A set of one or more context kinds that this metric can measure events from. Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
        /// </summary>
        public InputList<string> RandomizationUnits
        {
            get => _randomizationUnits ?? (_randomizationUnits = new InputList<string>());
            set => _randomizationUnits = value;
        }

        /// <summary>
        /// The CSS selector for your metric (if click metric)
        /// </summary>
        [Input("selector")]
        public Input<string>? Selector { get; set; }

        /// <summary>
        /// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
        /// </summary>
        [Input("successCriteria")]
        public Input<string>? SuccessCriteria { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags associated with your resource.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// (Required for kind `custom`) The unit for numeric `custom` metrics.
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        /// <summary>
        /// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
        /// </summary>
        [Input("unitAggregationType")]
        public Input<string>? UnitAggregationType { get; set; }

        [Input("urls")]
        private InputList<Inputs.MetricUrlArgs>? _urls;

        /// <summary>
        /// List of nested `url` blocks describing URLs that you want to associate with the metric.
        /// </summary>
        public InputList<Inputs.MetricUrlArgs> Urls
        {
            get => _urls ?? (_urls = new InputList<Inputs.MetricUrlArgs>());
            set => _urls = value;
        }

        public MetricArgs()
        {
        }
        public static new MetricArgs Empty => new MetricArgs();
    }

    public sealed class MetricState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The method for analyzing metric events. Available choices are `mean` and `percentile`.
        /// </summary>
        [Input("analysisType")]
        public Input<string>? AnalysisType { get; set; }

        /// <summary>
        /// The description of the metric's purpose.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The event key for your metric (if custom metric)
        /// </summary>
        [Input("eventKey")]
        public Input<string>? EventKey { get; set; }

        /// <summary>
        /// Include units that did not send any events and set their value to 0.
        /// </summary>
        [Input("includeUnitsWithoutEvents")]
        public Input<bool>? IncludeUnitsWithoutEvents { get; set; }

        /// <summary>
        /// Ignored. All metrics are considered active.
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Whether a `custom` metric is a numeric metric or not.
        /// </summary>
        [Input("isNumeric")]
        public Input<bool>? IsNumeric { get; set; }

        /// <summary>
        /// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The metric type. Available choices are `click`, `custom`, and `pageview`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("maintainerId")]
        public Input<string>? MaintainerId { get; set; }

        /// <summary>
        /// The human-friendly name for the metric.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when analysis_type is percentile.
        /// </summary>
        [Input("percentileValue")]
        public Input<int>? PercentileValue { get; set; }

        /// <summary>
        /// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("randomizationUnits")]
        private InputList<string>? _randomizationUnits;

        /// <summary>
        /// A set of one or more context kinds that this metric can measure events from. Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
        /// </summary>
        public InputList<string> RandomizationUnits
        {
            get => _randomizationUnits ?? (_randomizationUnits = new InputList<string>());
            set => _randomizationUnits = value;
        }

        /// <summary>
        /// The CSS selector for your metric (if click metric)
        /// </summary>
        [Input("selector")]
        public Input<string>? Selector { get; set; }

        /// <summary>
        /// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
        /// </summary>
        [Input("successCriteria")]
        public Input<string>? SuccessCriteria { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags associated with your resource.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// (Required for kind `custom`) The unit for numeric `custom` metrics.
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        /// <summary>
        /// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
        /// </summary>
        [Input("unitAggregationType")]
        public Input<string>? UnitAggregationType { get; set; }

        [Input("urls")]
        private InputList<Inputs.MetricUrlGetArgs>? _urls;

        /// <summary>
        /// List of nested `url` blocks describing URLs that you want to associate with the metric.
        /// </summary>
        public InputList<Inputs.MetricUrlGetArgs> Urls
        {
            get => _urls ?? (_urls = new InputList<Inputs.MetricUrlGetArgs>());
            set => _urls = value;
        }

        /// <summary>
        /// Version of the metric
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public MetricState()
        {
        }
        public static new MetricState Empty => new MetricState();
    }
}
