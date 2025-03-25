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
    /// Provides a LaunchDarkly project resource.
    /// 
    /// This resource allows you to create and manage projects within your LaunchDarkly organization.
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
    ///     var example = new Launchdarkly.Project("example", new()
    ///     {
    ///         Key = "example-project",
    ///         Name = "Example project",
    ///         Tags = new[]
    ///         {
    ///             "terraform",
    ///         },
    ///         Environments = new[]
    ///         {
    ///             new Launchdarkly.Inputs.ProjectEnvironmentArgs
    ///             {
    ///                 Key = "production",
    ///                 Name = "Production",
    ///                 Color = "EEEEEE",
    ///                 Tags = new[]
    ///                 {
    ///                     "terraform",
    ///                 },
    ///                 ApprovalSettings = new[]
    ///                 {
    ///                     new Launchdarkly.Inputs.ProjectEnvironmentApprovalSettingArgs
    ///                     {
    ///                         CanReviewOwnRequest = false,
    ///                         CanApplyDeclinedChanges = false,
    ///                         MinNumApprovals = 3,
    ///                         RequiredApprovalTags = new[]
    ///                         {
    ///                             "approvals_required",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             new Launchdarkly.Inputs.ProjectEnvironmentArgs
    ///             {
    ///                 Key = "staging",
    ///                 Name = "Staging",
    ///                 Color = "000000",
    ///                 Tags = new[]
    ///                 {
    ///                     "terraform",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LaunchDarkly projects can be imported using the project's key.
    /// 
    /// ```sh
    /// $ pulumi import launchdarkly:index/project:Project example example-project
    /// ```
    /// 
    /// **IMPORTANT:** Please note that, regardless of how many `environments` blocks you include on your import, _all_ of the project's environments will be saved to the Terraform state and will update with subsequent applies. This means that any environments not included in your import configuration will be torn down with any subsequent apply. If you wish to manage project properties with Terraform but not nested environments consider using Terraform's ignore changes lifecycle meta-argument; see below for example.
    /// 
    /// terraform
    /// 
    /// resource "launchdarkly_project" "example" {
    /// 
    ///   lifecycle {
    /// 
    ///     ignore_changes = [environments]
    /// 
    ///   }
    /// 
    ///   name = "testProject"
    /// 
    ///   key = "%s"
    /// 
    /// # environments not included on this configuration will not be affected by subsequent applies
    /// 
    /// }
    /// 
    /// **Managing environment resources with Terraform should always be done on the project unless the project is not also managed with Terraform.**
    /// </summary>
    [LaunchdarklyResourceType("launchdarkly:index/project:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A block describing which client-side SDKs can use new flags by default.
        /// </summary>
        [Output("defaultClientSideAvailabilities")]
        public Output<ImmutableArray<Outputs.ProjectDefaultClientSideAvailability>> DefaultClientSideAvailabilities { get; private set; } = null!;

        [Output("environments")]
        public Output<ImmutableArray<Outputs.ProjectEnvironment>> Environments { get; private set; } = null!;

        /// <summary>
        /// Whether feature flags created under the project should be available to client-side SDKs by default. Please migrate to
        /// `default_client_side_availability` to maintain future compatibility.
        /// </summary>
        [Output("includeInSnippet")]
        public Output<bool> IncludeInSnippet { get; private set; } = null!;

        /// <summary>
        /// The project's unique key. A change in this field will force the destruction of the existing resource and the creation of
        /// a new one.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The project's name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Tags associated with your resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs args, CustomResourceOptions? options = null)
            : base("launchdarkly:index/project:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
            : base("launchdarkly:index/project:Project", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new Project(name, id, state, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultClientSideAvailabilities")]
        private InputList<Inputs.ProjectDefaultClientSideAvailabilityArgs>? _defaultClientSideAvailabilities;

        /// <summary>
        /// A block describing which client-side SDKs can use new flags by default.
        /// </summary>
        public InputList<Inputs.ProjectDefaultClientSideAvailabilityArgs> DefaultClientSideAvailabilities
        {
            get => _defaultClientSideAvailabilities ?? (_defaultClientSideAvailabilities = new InputList<Inputs.ProjectDefaultClientSideAvailabilityArgs>());
            set => _defaultClientSideAvailabilities = value;
        }

        [Input("environments", required: true)]
        private InputList<Inputs.ProjectEnvironmentArgs>? _environments;
        public InputList<Inputs.ProjectEnvironmentArgs> Environments
        {
            get => _environments ?? (_environments = new InputList<Inputs.ProjectEnvironmentArgs>());
            set => _environments = value;
        }

        /// <summary>
        /// Whether feature flags created under the project should be available to client-side SDKs by default. Please migrate to
        /// `default_client_side_availability` to maintain future compatibility.
        /// </summary>
        [Input("includeInSnippet")]
        public Input<bool>? IncludeInSnippet { get; set; }

        /// <summary>
        /// The project's unique key. A change in this field will force the destruction of the existing resource and the creation of
        /// a new one.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The project's name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }

    public sealed class ProjectState : global::Pulumi.ResourceArgs
    {
        [Input("defaultClientSideAvailabilities")]
        private InputList<Inputs.ProjectDefaultClientSideAvailabilityGetArgs>? _defaultClientSideAvailabilities;

        /// <summary>
        /// A block describing which client-side SDKs can use new flags by default.
        /// </summary>
        public InputList<Inputs.ProjectDefaultClientSideAvailabilityGetArgs> DefaultClientSideAvailabilities
        {
            get => _defaultClientSideAvailabilities ?? (_defaultClientSideAvailabilities = new InputList<Inputs.ProjectDefaultClientSideAvailabilityGetArgs>());
            set => _defaultClientSideAvailabilities = value;
        }

        [Input("environments")]
        private InputList<Inputs.ProjectEnvironmentGetArgs>? _environments;
        public InputList<Inputs.ProjectEnvironmentGetArgs> Environments
        {
            get => _environments ?? (_environments = new InputList<Inputs.ProjectEnvironmentGetArgs>());
            set => _environments = value;
        }

        /// <summary>
        /// Whether feature flags created under the project should be available to client-side SDKs by default. Please migrate to
        /// `default_client_side_availability` to maintain future compatibility.
        /// </summary>
        [Input("includeInSnippet")]
        public Input<bool>? IncludeInSnippet { get; set; }

        /// <summary>
        /// The project's unique key. A change in this field will force the destruction of the existing resource and the creation of
        /// a new one.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The project's name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public ProjectState()
        {
        }
        public static new ProjectState Empty => new ProjectState();
    }
}
