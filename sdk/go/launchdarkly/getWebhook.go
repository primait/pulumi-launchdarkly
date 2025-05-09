// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package launchdarkly

import (
	"context"
	"reflect"

	"github.com/primait/pulumi-launchdarkly/sdk/go/launchdarkly/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a LaunchDarkly webhook data source.
//
// This data source allows you to retrieve webhook information from your LaunchDarkly organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/primait/pulumi-launchdarkly/sdk/go/launchdarkly"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := launchdarkly.LookupWebhook(ctx, &launchdarkly.LookupWebhookArgs{
//				Id: "57c0af6099690907435299",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupWebhook(ctx *pulumi.Context, args *LookupWebhookArgs, opts ...pulumi.InvokeOption) (*LookupWebhookResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWebhookResult
	err := ctx.Invoke("launchdarkly:index/getWebhook:getWebhook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWebhook.
type LookupWebhookArgs struct {
	// The unique webhook ID.
	Id string `pulumi:"id"`
}

// A collection of values returned by getWebhook.
type LookupWebhookResult struct {
	// The unique webhook ID.
	Id string `pulumi:"id"`
	// The webhook's human-readable name.
	Name string `pulumi:"name"`
	// Whether the webhook is enabled.
	On bool `pulumi:"on"`
	// The secret used to sign the webhook.
	Secret string `pulumi:"secret"`
	// List of policy statement blocks used to filter webhook events. For more information on webhook policy filters read [Adding a policy filter](https://docs.launchdarkly.com/integrations/webhooks#adding-a-policy-filter).
	Statements []GetWebhookStatement `pulumi:"statements"`
	// Tags associated with your resource.
	Tags []string `pulumi:"tags"`
	// The URL of the remote webhook.
	Url string `pulumi:"url"`
}

func LookupWebhookOutput(ctx *pulumi.Context, args LookupWebhookOutputArgs, opts ...pulumi.InvokeOption) LookupWebhookResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWebhookResultOutput, error) {
			args := v.(LookupWebhookArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("launchdarkly:index/getWebhook:getWebhook", args, LookupWebhookResultOutput{}, options).(LookupWebhookResultOutput), nil
		}).(LookupWebhookResultOutput)
}

// A collection of arguments for invoking getWebhook.
type LookupWebhookOutputArgs struct {
	// The unique webhook ID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupWebhookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebhookArgs)(nil)).Elem()
}

// A collection of values returned by getWebhook.
type LookupWebhookResultOutput struct{ *pulumi.OutputState }

func (LookupWebhookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebhookResult)(nil)).Elem()
}

func (o LookupWebhookResultOutput) ToLookupWebhookResultOutput() LookupWebhookResultOutput {
	return o
}

func (o LookupWebhookResultOutput) ToLookupWebhookResultOutputWithContext(ctx context.Context) LookupWebhookResultOutput {
	return o
}

// The unique webhook ID.
func (o LookupWebhookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Id }).(pulumi.StringOutput)
}

// The webhook's human-readable name.
func (o LookupWebhookResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether the webhook is enabled.
func (o LookupWebhookResultOutput) On() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWebhookResult) bool { return v.On }).(pulumi.BoolOutput)
}

// The secret used to sign the webhook.
func (o LookupWebhookResultOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Secret }).(pulumi.StringOutput)
}

// List of policy statement blocks used to filter webhook events. For more information on webhook policy filters read [Adding a policy filter](https://docs.launchdarkly.com/integrations/webhooks#adding-a-policy-filter).
func (o LookupWebhookResultOutput) Statements() GetWebhookStatementArrayOutput {
	return o.ApplyT(func(v LookupWebhookResult) []GetWebhookStatement { return v.Statements }).(GetWebhookStatementArrayOutput)
}

// Tags associated with your resource.
func (o LookupWebhookResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebhookResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The URL of the remote webhook.
func (o LookupWebhookResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebhookResultOutput{})
}
