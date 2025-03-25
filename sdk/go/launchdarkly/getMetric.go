// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package launchdarkly

import (
	"context"
	"reflect"

	"github.com/primait/pulumi-launchdarkly/sdk/go/launchdarkly/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a LaunchDarkly metric data source.
//
// This data source allows you to retrieve metric information from your LaunchDarkly organization.
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
//			_, err := launchdarkly.LookupMetric(ctx, &launchdarkly.LookupMetricArgs{
//				Key:        "example-metric",
//				ProjectKey: "example-project",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupMetric(ctx *pulumi.Context, args *LookupMetricArgs, opts ...pulumi.InvokeOption) (*LookupMetricResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMetricResult
	err := ctx.Invoke("launchdarkly:index/getMetric:getMetric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMetric.
type LookupMetricArgs struct {
	// Ignored. All metrics are considered active.
	//
	// Deprecated: No longer in use. This field will be removed in a future major release of the LaunchDarkly provider.
	IsActive *bool `pulumi:"isActive"`
	// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key string `pulumi:"key"`
	// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey string `pulumi:"projectKey"`
}

// A collection of values returned by getMetric.
type LookupMetricResult struct {
	// The method for analyzing metric events. Available choices are `mean` and `percentile`.
	AnalysisType string `pulumi:"analysisType"`
	// The description of the metric's purpose.
	Description string `pulumi:"description"`
	// The event key for your metric (if custom metric)
	EventKey string `pulumi:"eventKey"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Include units that did not send any events and set their value to 0.
	IncludeUnitsWithoutEvents bool `pulumi:"includeUnitsWithoutEvents"`
	// Ignored. All metrics are considered active.
	//
	// Deprecated: No longer in use. This field will be removed in a future major release of the LaunchDarkly provider.
	IsActive bool `pulumi:"isActive"`
	// Whether a `custom` metric is a numeric metric or not.
	IsNumeric bool `pulumi:"isNumeric"`
	// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key string `pulumi:"key"`
	// The metric type. Available choices are `click`, `custom`, and `pageview`.
	Kind         string `pulumi:"kind"`
	MaintainerId string `pulumi:"maintainerId"`
	// The human-friendly name for the metric.
	Name string `pulumi:"name"`
	// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when analysisType is percentile.
	PercentileValue int `pulumi:"percentileValue"`
	// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey string `pulumi:"projectKey"`
	// A set of one or more context kinds that this metric can measure events from. Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
	RandomizationUnits []string `pulumi:"randomizationUnits"`
	// The CSS selector for your metric (if click metric)
	Selector string `pulumi:"selector"`
	// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
	SuccessCriteria string `pulumi:"successCriteria"`
	// Tags associated with your resource.
	Tags []string `pulumi:"tags"`
	// (Required for kind `custom`) The unit for numeric `custom` metrics.
	Unit string `pulumi:"unit"`
	// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
	UnitAggregationType string `pulumi:"unitAggregationType"`
	// List of nested `url` blocks describing URLs that you want to associate with the metric.
	Urls []GetMetricUrl `pulumi:"urls"`
	// Version of the metric
	Version int `pulumi:"version"`
}

func LookupMetricOutput(ctx *pulumi.Context, args LookupMetricOutputArgs, opts ...pulumi.InvokeOption) LookupMetricResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMetricResultOutput, error) {
			args := v.(LookupMetricArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("launchdarkly:index/getMetric:getMetric", args, LookupMetricResultOutput{}, options).(LookupMetricResultOutput), nil
		}).(LookupMetricResultOutput)
}

// A collection of arguments for invoking getMetric.
type LookupMetricOutputArgs struct {
	// Ignored. All metrics are considered active.
	//
	// Deprecated: No longer in use. This field will be removed in a future major release of the LaunchDarkly provider.
	IsActive pulumi.BoolPtrInput `pulumi:"isActive"`
	// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key pulumi.StringInput `pulumi:"key"`
	// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey pulumi.StringInput `pulumi:"projectKey"`
}

func (LookupMetricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricArgs)(nil)).Elem()
}

// A collection of values returned by getMetric.
type LookupMetricResultOutput struct{ *pulumi.OutputState }

func (LookupMetricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricResult)(nil)).Elem()
}

func (o LookupMetricResultOutput) ToLookupMetricResultOutput() LookupMetricResultOutput {
	return o
}

func (o LookupMetricResultOutput) ToLookupMetricResultOutputWithContext(ctx context.Context) LookupMetricResultOutput {
	return o
}

// The method for analyzing metric events. Available choices are `mean` and `percentile`.
func (o LookupMetricResultOutput) AnalysisType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.AnalysisType }).(pulumi.StringOutput)
}

// The description of the metric's purpose.
func (o LookupMetricResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.Description }).(pulumi.StringOutput)
}

// The event key for your metric (if custom metric)
func (o LookupMetricResultOutput) EventKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.EventKey }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMetricResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.Id }).(pulumi.StringOutput)
}

// Include units that did not send any events and set their value to 0.
func (o LookupMetricResultOutput) IncludeUnitsWithoutEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMetricResult) bool { return v.IncludeUnitsWithoutEvents }).(pulumi.BoolOutput)
}

// Ignored. All metrics are considered active.
//
// Deprecated: No longer in use. This field will be removed in a future major release of the LaunchDarkly provider.
func (o LookupMetricResultOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMetricResult) bool { return v.IsActive }).(pulumi.BoolOutput)
}

// Whether a `custom` metric is a numeric metric or not.
func (o LookupMetricResultOutput) IsNumeric() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMetricResult) bool { return v.IsNumeric }).(pulumi.BoolOutput)
}

// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one.
func (o LookupMetricResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.Key }).(pulumi.StringOutput)
}

// The metric type. Available choices are `click`, `custom`, and `pageview`.
func (o LookupMetricResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupMetricResultOutput) MaintainerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.MaintainerId }).(pulumi.StringOutput)
}

// The human-friendly name for the metric.
func (o LookupMetricResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.Name }).(pulumi.StringOutput)
}

// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when analysisType is percentile.
func (o LookupMetricResultOutput) PercentileValue() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetricResult) int { return v.PercentileValue }).(pulumi.IntOutput)
}

// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
func (o LookupMetricResultOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.ProjectKey }).(pulumi.StringOutput)
}

// A set of one or more context kinds that this metric can measure events from. Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
func (o LookupMetricResultOutput) RandomizationUnits() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetricResult) []string { return v.RandomizationUnits }).(pulumi.StringArrayOutput)
}

// The CSS selector for your metric (if click metric)
func (o LookupMetricResultOutput) Selector() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.Selector }).(pulumi.StringOutput)
}

// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
func (o LookupMetricResultOutput) SuccessCriteria() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.SuccessCriteria }).(pulumi.StringOutput)
}

// Tags associated with your resource.
func (o LookupMetricResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetricResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// (Required for kind `custom`) The unit for numeric `custom` metrics.
func (o LookupMetricResultOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.Unit }).(pulumi.StringOutput)
}

// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
func (o LookupMetricResultOutput) UnitAggregationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricResult) string { return v.UnitAggregationType }).(pulumi.StringOutput)
}

// List of nested `url` blocks describing URLs that you want to associate with the metric.
func (o LookupMetricResultOutput) Urls() GetMetricUrlArrayOutput {
	return o.ApplyT(func(v LookupMetricResult) []GetMetricUrl { return v.Urls }).(GetMetricUrlArrayOutput)
}

// Version of the metric
func (o LookupMetricResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetricResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetricResultOutput{})
}
