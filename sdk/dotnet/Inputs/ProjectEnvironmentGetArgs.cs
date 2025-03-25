// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Launchdarkly.Inputs
{

    public sealed class ProjectEnvironmentGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey")]
        private Input<string>? _apiKey;

        /// <summary>
        /// The environment's SDK key.
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("approvalSettings")]
        private InputList<Inputs.ProjectEnvironmentApprovalSettingGetArgs>? _approvalSettings;
        public InputList<Inputs.ProjectEnvironmentApprovalSettingGetArgs> ApprovalSettings
        {
            get => _approvalSettings ?? (_approvalSettings = new InputList<Inputs.ProjectEnvironmentApprovalSettingGetArgs>());
            set => _approvalSettings = value;
        }

        [Input("clientSideId")]
        private Input<string>? _clientSideId;

        /// <summary>
        /// The environment's client-side ID.
        /// </summary>
        public Input<string>? ClientSideId
        {
            get => _clientSideId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSideId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The color swatch as an RGB hex value with no leading `#`. For example: `000000`
        /// </summary>
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        /// <summary>
        /// Set to `true` if this environment requires confirmation for flag and segment changes. This field will default to `false` when not set.
        /// </summary>
        [Input("confirmChanges")]
        public Input<bool>? ConfirmChanges { get; set; }

        /// <summary>
        /// Denotes whether the environment is critical.
        /// </summary>
        [Input("critical")]
        public Input<bool>? Critical { get; set; }

        /// <summary>
        /// Set to `true` to enable data export for every flag created in this environment after you configure this argument. This field will default to `false` when not set. To learn more, read [Data Export](https://docs.launchdarkly.com/home/data-export).
        /// </summary>
        [Input("defaultTrackEvents")]
        public Input<bool>? DefaultTrackEvents { get; set; }

        /// <summary>
        /// The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
        /// </summary>
        [Input("defaultTtl")]
        public Input<int>? DefaultTtl { get; set; }

        /// <summary>
        /// The project-unique key for the environment. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("mobileKey")]
        private Input<string>? _mobileKey;

        /// <summary>
        /// The environment's mobile key.
        /// </summary>
        public Input<string>? MobileKey
        {
            get => _mobileKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _mobileKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Set to `true` if this environment requires comments for flag and segment changes. This field will default to `false` when not set.
        /// </summary>
        [Input("requireComments")]
        public Input<bool>? RequireComments { get; set; }

        /// <summary>
        /// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user. This field will default to `false` when not set.
        /// </summary>
        [Input("secureMode")]
        public Input<bool>? SecureMode { get; set; }

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

        public ProjectEnvironmentGetArgs()
        {
        }
        public static new ProjectEnvironmentGetArgs Empty => new ProjectEnvironmentGetArgs();
    }
}
