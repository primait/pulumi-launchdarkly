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
    /// The provider type for the launchdarkly package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [LaunchdarklyResourceType("pulumi:providers:launchdarkly")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// The [personal access token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#personal-tokens) or
        /// [service token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens) used to
        /// authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_ACCESS_TOKEN` environment variable. You
        /// must provide either `access_token` or `oauth_token`.
        /// </summary>
        [Output("accessToken")]
        public Output<string?> AccessToken { get; private set; } = null!;

        /// <summary>
        /// The LaunchDarkly host address. If this argument is not specified, the default host address is
        /// `https://app.launchdarkly.com`
        /// </summary>
        [Output("apiHost")]
        public Output<string?> ApiHost { get; private set; } = null!;

        /// <summary>
        /// An OAuth V2 token you use to authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_OAUTH_TOKEN`
        /// environment variable. You must provide either `access_token` or `oauth_token`.
        /// </summary>
        [Output("oauthToken")]
        public Output<string?> OauthToken { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("launchdarkly", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [personal access token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#personal-tokens) or
        /// [service token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens) used to
        /// authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_ACCESS_TOKEN` environment variable. You
        /// must provide either `access_token` or `oauth_token`.
        /// </summary>
        [Input("accessToken")]
        public Input<string>? AccessToken { get; set; }

        /// <summary>
        /// The LaunchDarkly host address. If this argument is not specified, the default host address is
        /// `https://app.launchdarkly.com`
        /// </summary>
        [Input("apiHost")]
        public Input<string>? ApiHost { get; set; }

        /// <summary>
        /// The HTTP timeout (in seconds) when making API calls to LaunchDarkly.
        /// </summary>
        [Input("httpTimeout", json: true)]
        public Input<int>? HttpTimeout { get; set; }

        /// <summary>
        /// An OAuth V2 token you use to authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_OAUTH_TOKEN`
        /// environment variable. You must provide either `access_token` or `oauth_token`.
        /// </summary>
        [Input("oauthToken")]
        public Input<string>? OauthToken { get; set; }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
