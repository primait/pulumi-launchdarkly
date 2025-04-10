// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package launchdarkly

import (
	"context"
	"reflect"

	"github.com/primait/pulumi-launchdarkly/sdk/go/launchdarkly/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the launchdarkly package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// The [personal access token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#personal-tokens) or
	// [service token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens) used to
	// authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_ACCESS_TOKEN` environment variable. You
	// must provide either `accessToken` or `oauthToken`.
	AccessToken pulumi.StringPtrOutput `pulumi:"accessToken"`
	// The LaunchDarkly host address. If this argument is not specified, the default host address is
	// `https://app.launchdarkly.com`
	ApiHost pulumi.StringPtrOutput `pulumi:"apiHost"`
	// An OAuth V2 token you use to authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_OAUTH_TOKEN`
	// environment variable. You must provide either `accessToken` or `oauthToken`.
	OauthToken pulumi.StringPtrOutput `pulumi:"oauthToken"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.AccessToken != nil {
		args.AccessToken = pulumi.ToSecret(args.AccessToken).(pulumi.StringPtrInput)
	}
	if args.OauthToken != nil {
		args.OauthToken = pulumi.ToSecret(args.OauthToken).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessToken",
		"oauthToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:launchdarkly", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The [personal access token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#personal-tokens) or
	// [service token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens) used to
	// authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_ACCESS_TOKEN` environment variable. You
	// must provide either `accessToken` or `oauthToken`.
	AccessToken *string `pulumi:"accessToken"`
	// The LaunchDarkly host address. If this argument is not specified, the default host address is
	// `https://app.launchdarkly.com`
	ApiHost *string `pulumi:"apiHost"`
	// The HTTP timeout (in seconds) when making API calls to LaunchDarkly.
	HttpTimeout *int `pulumi:"httpTimeout"`
	// An OAuth V2 token you use to authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_OAUTH_TOKEN`
	// environment variable. You must provide either `accessToken` or `oauthToken`.
	OauthToken *string `pulumi:"oauthToken"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The [personal access token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#personal-tokens) or
	// [service token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens) used to
	// authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_ACCESS_TOKEN` environment variable. You
	// must provide either `accessToken` or `oauthToken`.
	AccessToken pulumi.StringPtrInput
	// The LaunchDarkly host address. If this argument is not specified, the default host address is
	// `https://app.launchdarkly.com`
	ApiHost pulumi.StringPtrInput
	// The HTTP timeout (in seconds) when making API calls to LaunchDarkly.
	HttpTimeout pulumi.IntPtrInput
	// An OAuth V2 token you use to authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_OAUTH_TOKEN`
	// environment variable. You must provide either `accessToken` or `oauthToken`.
	OauthToken pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// The [personal access token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#personal-tokens) or
// [service token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens) used to
// authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_ACCESS_TOKEN` environment variable. You
// must provide either `accessToken` or `oauthToken`.
func (o ProviderOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.AccessToken }).(pulumi.StringPtrOutput)
}

// The LaunchDarkly host address. If this argument is not specified, the default host address is
// `https://app.launchdarkly.com`
func (o ProviderOutput) ApiHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ApiHost }).(pulumi.StringPtrOutput)
}

// An OAuth V2 token you use to authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_OAUTH_TOKEN`
// environment variable. You must provide either `accessToken` or `oauthToken`.
func (o ProviderOutput) OauthToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OauthToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
