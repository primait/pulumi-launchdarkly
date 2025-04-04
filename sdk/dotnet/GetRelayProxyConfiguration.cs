// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Launchdarkly
{
    public static class GetRelayProxyConfiguration
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
        ///     var example = Launchdarkly.GetRelayProxyConfiguration.Invoke(new()
        ///     {
        ///         Id = "123456789abcdef123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRelayProxyConfigurationResult> InvokeAsync(GetRelayProxyConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRelayProxyConfigurationResult>("launchdarkly:index/getRelayProxyConfiguration:getRelayProxyConfiguration", args ?? new GetRelayProxyConfigurationArgs(), options.WithDefaults());

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
        ///     var example = Launchdarkly.GetRelayProxyConfiguration.Invoke(new()
        ///     {
        ///         Id = "123456789abcdef123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRelayProxyConfigurationResult> Invoke(GetRelayProxyConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRelayProxyConfigurationResult>("launchdarkly:index/getRelayProxyConfiguration:getRelayProxyConfiguration", args ?? new GetRelayProxyConfigurationInvokeArgs(), options.WithDefaults());

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
        ///     var example = Launchdarkly.GetRelayProxyConfiguration.Invoke(new()
        ///     {
        ///         Id = "123456789abcdef123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRelayProxyConfigurationResult> Invoke(GetRelayProxyConfigurationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRelayProxyConfigurationResult>("launchdarkly:index/getRelayProxyConfiguration:getRelayProxyConfiguration", args ?? new GetRelayProxyConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRelayProxyConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Relay Proxy configuration's unique 24 character ID. The unique relay proxy ID can be found in the relay proxy edit page URL, which you can locate by clicking the three dot menu on your relay proxy item in the UI and selecting "Edit configuration":
        /// `
        /// https://app.launchdarkly.com/settings/relay/THIS_IS_YOUR_RELAY_PROXY_ID/edit
        /// `
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetRelayProxyConfigurationArgs()
        {
        }
        public static new GetRelayProxyConfigurationArgs Empty => new GetRelayProxyConfigurationArgs();
    }

    public sealed class GetRelayProxyConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Relay Proxy configuration's unique 24 character ID. The unique relay proxy ID can be found in the relay proxy edit page URL, which you can locate by clicking the three dot menu on your relay proxy item in the UI and selecting "Edit configuration":
        /// `
        /// https://app.launchdarkly.com/settings/relay/THIS_IS_YOUR_RELAY_PROXY_ID/edit
        /// `
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetRelayProxyConfigurationInvokeArgs()
        {
        }
        public static new GetRelayProxyConfigurationInvokeArgs Empty => new GetRelayProxyConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetRelayProxyConfigurationResult
    {
        /// <summary>
        /// The last 4 characters of the Relay Proxy configuration's unique key.
        /// </summary>
        public readonly string DisplayKey;
        /// <summary>
        /// The Relay Proxy configuration's unique 24 character ID. The unique relay proxy ID can be found in the relay proxy edit page URL, which you can locate by clicking the three dot menu on your relay proxy item in the UI and selecting "Edit configuration":
        /// `
        /// https://app.launchdarkly.com/settings/relay/THIS_IS_YOUR_RELAY_PROXY_ID/edit
        /// `
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The human-readable name for your Relay Proxy configuration.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Relay Proxy configuration's rule policy block. This determines what content the Relay Proxy receives. To learn more, read [Understanding policies](https://docs.launchdarkly.com/home/members/role-policies#understanding-policies).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRelayProxyConfigurationPolicyResult> Policies;

        [OutputConstructor]
        private GetRelayProxyConfigurationResult(
            string displayKey,

            string id,

            string name,

            ImmutableArray<Outputs.GetRelayProxyConfigurationPolicyResult> policies)
        {
            DisplayKey = displayKey;
            Id = id;
            Name = name;
            Policies = policies;
        }
    }
}
