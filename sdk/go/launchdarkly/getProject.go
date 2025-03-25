// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package launchdarkly

import (
	"context"
	"reflect"

	"github.com/primait/pulumi-launchdarkly/sdk/go/launchdarkly/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a LaunchDarkly project data source.
//
// This data source allows you to retrieve project information from your LaunchDarkly organization.
//
// > **Note:** LaunchDarkly data sources do not provide access to the project's environments. If you wish to import environment configurations as data sources you must use the [`Environment` data source](https://www.terraform.io/docs/providers/launchdarkly/d/environment.html).
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
//			_, err := launchdarkly.LookupProject(ctx, &launchdarkly.LookupProjectArgs{
//				Key: "example-project",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("launchdarkly:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// The project's unique key.
	Key string `pulumi:"key"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	// A map describing which client-side SDKs can use new flags by default. Please migrate to `defaultClientSideAvailability` to maintain future compatibility.
	//
	// Deprecated: 'client_side_availability' is now deprecated. Please migrate to 'default_client_side_availability' to maintain future compatibility.
	ClientSideAvailabilities []GetProjectClientSideAvailability `pulumi:"clientSideAvailabilities"`
	// A block describing which client-side SDKs can use new flags by default.
	DefaultClientSideAvailabilities []GetProjectDefaultClientSideAvailability `pulumi:"defaultClientSideAvailabilities"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The project's unique key.
	Key string `pulumi:"key"`
	// The project's name.
	Name string `pulumi:"name"`
	// Tags associated with your resource.
	Tags []string `pulumi:"tags"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectResultOutput, error) {
			args := v.(LookupProjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("launchdarkly:index/getProject:getProject", args, LookupProjectResultOutput{}, options).(LookupProjectResultOutput), nil
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// The project's unique key.
	Key pulumi.StringInput `pulumi:"key"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// A map describing which client-side SDKs can use new flags by default. Please migrate to `defaultClientSideAvailability` to maintain future compatibility.
//
// Deprecated: 'client_side_availability' is now deprecated. Please migrate to 'default_client_side_availability' to maintain future compatibility.
func (o LookupProjectResultOutput) ClientSideAvailabilities() GetProjectClientSideAvailabilityArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []GetProjectClientSideAvailability { return v.ClientSideAvailabilities }).(GetProjectClientSideAvailabilityArrayOutput)
}

// A block describing which client-side SDKs can use new flags by default.
func (o LookupProjectResultOutput) DefaultClientSideAvailabilities() GetProjectDefaultClientSideAvailabilityArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []GetProjectDefaultClientSideAvailability {
		return v.DefaultClientSideAvailabilities
	}).(GetProjectDefaultClientSideAvailabilityArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// The project's unique key.
func (o LookupProjectResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Key }).(pulumi.StringOutput)
}

// The project's name.
func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// Tags associated with your resource.
func (o LookupProjectResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
