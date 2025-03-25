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
    public sealed class GetRelayProxyConfigurationPolicyResult
    {
        /// <summary>
        /// The list of action specifiers defining the actions to which the statement applies.
        /// Either `actions` or `not_actions` must be specified. For a list of available actions read [Actions reference](https://docs.launchdarkly.com/home/account-security/custom-roles/actions#actions-reference).
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        /// <summary>
        /// Either `allow` or `deny`. This argument defines whether the statement allows or denies access to the named resources and actions.
        /// </summary>
        public readonly string Effect;
        /// <summary>
        /// The list of action specifiers defining the actions to which the statement does not apply.
        /// </summary>
        public readonly ImmutableArray<string> NotActions;
        /// <summary>
        /// The list of resource specifiers defining the resources to which the statement does not apply.
        /// </summary>
        public readonly ImmutableArray<string> NotResources;
        /// <summary>
        /// The list of resource specifiers defining the resources to which the statement applies.
        /// </summary>
        public readonly ImmutableArray<string> Resources;

        [OutputConstructor]
        private GetRelayProxyConfigurationPolicyResult(
            ImmutableArray<string> actions,

            string effect,

            ImmutableArray<string> notActions,

            ImmutableArray<string> notResources,

            ImmutableArray<string> resources)
        {
            Actions = actions;
            Effect = effect;
            NotActions = notActions;
            NotResources = notResources;
            Resources = resources;
        }
    }
}
