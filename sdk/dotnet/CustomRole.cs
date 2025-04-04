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
    /// Provides a LaunchDarkly custom role resource.
    /// 
    /// &gt; **Note:** Custom roles are available to customers on an Enterprise LaunchDarkly plan. To learn more, [read about our pricing](https://launchdarkly.com/pricing/). To upgrade your plan, [contact LaunchDarkly Sales](https://launchdarkly.com/contact-sales/).
    /// 
    /// This resource allows you to create and manage custom roles within your LaunchDarkly organization.
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
    ///     var example = new Launchdarkly.CustomRole("example", new()
    ///     {
    ///         Key = "example-role-key-1",
    ///         Name = "example role",
    ///         Description = "This is an example role",
    ///         PolicyStatements = new[]
    ///         {
    ///             new Launchdarkly.Inputs.CustomRolePolicyStatementArgs
    ///             {
    ///                 Effect = "allow",
    ///                 Resources = new[]
    ///                 {
    ///                     "proj/*:env/production:flag/*",
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "*",
    ///                 },
    ///             },
    ///             new Launchdarkly.Inputs.CustomRolePolicyStatementArgs
    ///             {
    ///                 Effect = "allow",
    ///                 Resources = new[]
    ///                 {
    ///                     "proj/*:env/production",
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "*",
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
    /// ```sh
    /// $ pulumi import launchdarkly:index/customRole:CustomRole example example-role-key-1
    /// ```
    /// </summary>
    [LaunchdarklyResourceType("launchdarkly:index/customRole:CustomRole")]
    public partial class CustomRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The base permission level - either reader or no_access. Defaults to reader.
        /// </summary>
        [Output("basePermissions")]
        public Output<string?> BasePermissions { get; private set; } = null!;

        /// <summary>
        /// Description of the custom role.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A unique key that will be used to reference the custom role in your code. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// A name for the custom role. This must be unique within your organization.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("policies")]
        public Output<ImmutableArray<Outputs.CustomRolePolicy>> Policies { get; private set; } = null!;

        /// <summary>
        /// An array of the policy statements that define the permissions for the custom role. This field accepts [role attributes](https://docs.launchdarkly.com/home/getting-started/vocabulary#role-attribute). To use role attributes, use the syntax `$${roleAttribute/&lt;YOUR_ROLE_ATTRIBUTE&gt;}` in lieu of your usual resource keys.
        /// </summary>
        [Output("policyStatements")]
        public Output<ImmutableArray<Outputs.CustomRolePolicyStatement>> PolicyStatements { get; private set; } = null!;


        /// <summary>
        /// Create a CustomRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomRole(string name, CustomRoleArgs args, CustomResourceOptions? options = null)
            : base("launchdarkly:index/customRole:CustomRole", name, args ?? new CustomRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomRole(string name, Input<string> id, CustomRoleState? state = null, CustomResourceOptions? options = null)
            : base("launchdarkly:index/customRole:CustomRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomRole Get(string name, Input<string> id, CustomRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomRole(name, id, state, options);
        }
    }

    public sealed class CustomRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base permission level - either reader or no_access. Defaults to reader.
        /// </summary>
        [Input("basePermissions")]
        public Input<string>? BasePermissions { get; set; }

        /// <summary>
        /// Description of the custom role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique key that will be used to reference the custom role in your code. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// A name for the custom role. This must be unique within your organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policies")]
        private InputList<Inputs.CustomRolePolicyArgs>? _policies;
        [Obsolete(@"'policy' is now deprecated. Please migrate to 'policy_statements' to maintain future compatability.")]
        public InputList<Inputs.CustomRolePolicyArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.CustomRolePolicyArgs>());
            set => _policies = value;
        }

        [Input("policyStatements")]
        private InputList<Inputs.CustomRolePolicyStatementArgs>? _policyStatements;

        /// <summary>
        /// An array of the policy statements that define the permissions for the custom role. This field accepts [role attributes](https://docs.launchdarkly.com/home/getting-started/vocabulary#role-attribute). To use role attributes, use the syntax `$${roleAttribute/&lt;YOUR_ROLE_ATTRIBUTE&gt;}` in lieu of your usual resource keys.
        /// </summary>
        public InputList<Inputs.CustomRolePolicyStatementArgs> PolicyStatements
        {
            get => _policyStatements ?? (_policyStatements = new InputList<Inputs.CustomRolePolicyStatementArgs>());
            set => _policyStatements = value;
        }

        public CustomRoleArgs()
        {
        }
        public static new CustomRoleArgs Empty => new CustomRoleArgs();
    }

    public sealed class CustomRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base permission level - either reader or no_access. Defaults to reader.
        /// </summary>
        [Input("basePermissions")]
        public Input<string>? BasePermissions { get; set; }

        /// <summary>
        /// Description of the custom role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique key that will be used to reference the custom role in your code. A change in this field will force the destruction of the existing resource and the creation of a new one.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// A name for the custom role. This must be unique within your organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policies")]
        private InputList<Inputs.CustomRolePolicyGetArgs>? _policies;
        [Obsolete(@"'policy' is now deprecated. Please migrate to 'policy_statements' to maintain future compatability.")]
        public InputList<Inputs.CustomRolePolicyGetArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.CustomRolePolicyGetArgs>());
            set => _policies = value;
        }

        [Input("policyStatements")]
        private InputList<Inputs.CustomRolePolicyStatementGetArgs>? _policyStatements;

        /// <summary>
        /// An array of the policy statements that define the permissions for the custom role. This field accepts [role attributes](https://docs.launchdarkly.com/home/getting-started/vocabulary#role-attribute). To use role attributes, use the syntax `$${roleAttribute/&lt;YOUR_ROLE_ATTRIBUTE&gt;}` in lieu of your usual resource keys.
        /// </summary>
        public InputList<Inputs.CustomRolePolicyStatementGetArgs> PolicyStatements
        {
            get => _policyStatements ?? (_policyStatements = new InputList<Inputs.CustomRolePolicyStatementGetArgs>());
            set => _policyStatements = value;
        }

        public CustomRoleState()
        {
        }
        public static new CustomRoleState Empty => new CustomRoleState();
    }
}
