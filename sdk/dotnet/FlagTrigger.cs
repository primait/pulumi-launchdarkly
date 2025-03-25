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
    ///     var example = new Launchdarkly.FlagTrigger("example", new()
    ///     {
    ///         ProjectKey = exampleLaunchdarklyProject.Key,
    ///         EnvKey = "test",
    ///         FlagKey = triggerFlag.Key,
    ///         IntegrationKey = "generic-trigger",
    ///         Instructions = new Launchdarkly.Inputs.FlagTriggerInstructionsArgs
    ///         {
    ///             Kind = "turnFlagOn",
    ///         },
    ///         Enabled = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Import a LaunchDarkly flag trigger using this format: `project_key/environment_key/flag_key/trigger_id`.
    /// 
    /// ```sh
    /// $ pulumi import launchdarkly:index/flagTrigger:FlagTrigger example example-project-key/example-env-key/example-flag-key/62581d4488def814b831abc3
    /// ```
    /// 
    /// The unique trigger ID can be found in your saved trigger URL:
    /// 
    /// https://app.launchdarkly.com/webhook/triggers/THIS_IS_YOUR_TRIGGER_ID/aff25a53-17d9-4112-a9b8-12718d1a2e79
    /// 
    /// Please note that if you did not save this upon creation of the resource, you will have to reset it to get a new value, which can cause breaking changes.
    /// </summary>
    [LaunchdarklyResourceType("launchdarkly:index/flagTrigger:FlagTrigger")]
    public partial class FlagTrigger : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the trigger is currently active or not.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Output("envKey")]
        public Output<string> EnvKey { get; private set; } = null!;

        /// <summary>
        /// The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Output("flagKey")]
        public Output<string> FlagKey { get; private set; } = null!;

        /// <summary>
        /// Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "&lt;flag_action&gt;" }`.
        /// </summary>
        [Output("instructions")]
        public Output<Outputs.FlagTriggerInstructions> Instructions { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Output("integrationKey")]
        public Output<string> IntegrationKey { get; private set; } = null!;

        [Output("maintainerId")]
        public Output<string> MaintainerId { get; private set; } = null!;

        /// <summary>
        /// The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Output("projectKey")]
        public Output<string> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// The unique URL used to invoke the trigger.
        /// </summary>
        [Output("triggerUrl")]
        public Output<string> TriggerUrl { get; private set; } = null!;


        /// <summary>
        /// Create a FlagTrigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FlagTrigger(string name, FlagTriggerArgs args, CustomResourceOptions? options = null)
            : base("launchdarkly:index/flagTrigger:FlagTrigger", name, args ?? new FlagTriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FlagTrigger(string name, Input<string> id, FlagTriggerState? state = null, CustomResourceOptions? options = null)
            : base("launchdarkly:index/flagTrigger:FlagTrigger", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/primait/pulumi-launchdarkly/releases/",
                AdditionalSecretOutputs =
                {
                    "triggerUrl",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FlagTrigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FlagTrigger Get(string name, Input<string> id, FlagTriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new FlagTrigger(name, id, state, options);
        }
    }

    public sealed class FlagTriggerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the trigger is currently active or not.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("envKey", required: true)]
        public Input<string> EnvKey { get; set; } = null!;

        /// <summary>
        /// The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("flagKey", required: true)]
        public Input<string> FlagKey { get; set; } = null!;

        /// <summary>
        /// Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "&lt;flag_action&gt;" }`.
        /// </summary>
        [Input("instructions", required: true)]
        public Input<Inputs.FlagTriggerInstructionsArgs> Instructions { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("integrationKey", required: true)]
        public Input<string> IntegrationKey { get; set; } = null!;

        /// <summary>
        /// The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("projectKey", required: true)]
        public Input<string> ProjectKey { get; set; } = null!;

        public FlagTriggerArgs()
        {
        }
        public static new FlagTriggerArgs Empty => new FlagTriggerArgs();
    }

    public sealed class FlagTriggerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the trigger is currently active or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("envKey")]
        public Input<string>? EnvKey { get; set; }

        /// <summary>
        /// The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("flagKey")]
        public Input<string>? FlagKey { get; set; }

        /// <summary>
        /// Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "&lt;flag_action&gt;" }`.
        /// </summary>
        [Input("instructions")]
        public Input<Inputs.FlagTriggerInstructionsGetArgs>? Instructions { get; set; }

        /// <summary>
        /// The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("integrationKey")]
        public Input<string>? IntegrationKey { get; set; }

        [Input("maintainerId")]
        public Input<string>? MaintainerId { get; set; }

        /// <summary>
        /// The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("triggerUrl")]
        private Input<string>? _triggerUrl;

        /// <summary>
        /// The unique URL used to invoke the trigger.
        /// </summary>
        public Input<string>? TriggerUrl
        {
            get => _triggerUrl;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _triggerUrl = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public FlagTriggerState()
        {
        }
        public static new FlagTriggerState Empty => new FlagTriggerState();
    }
}
