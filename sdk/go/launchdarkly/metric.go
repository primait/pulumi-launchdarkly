// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package launchdarkly

import (
	"context"
	"reflect"

	"errors"
	"github.com/primait/pulumi-launchdarkly/sdk/go/launchdarkly/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a LaunchDarkly metric resource.
//
// This resource allows you to create and manage metrics within your LaunchDarkly organization.
//
// To learn more about metrics and experimentation, read [Experimentation Documentation](https://docs.launchdarkly.com/home/experimentation).
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
//			_, err := launchdarkly.NewMetric(ctx, "example", &launchdarkly.MetricArgs{
//				ProjectKey:  pulumi.Any(exampleLaunchdarklyProject.Key),
//				Key:         pulumi.String("example-metric"),
//				Name:        pulumi.String("Example Metric"),
//				Description: pulumi.String("Metric description."),
//				Kind:        pulumi.String("pageview"),
//				Tags: pulumi.StringArray{
//					pulumi.String("example"),
//				},
//				Urls: launchdarkly.MetricUrlArray{
//					&launchdarkly.MetricUrlArgs{
//						Kind:      pulumi.String("substring"),
//						Substring: pulumi.String("foo"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// LaunchDarkly metrics can be imported using the metric's ID in the form `project_key/metric_key`
//
// ```sh
// $ pulumi import launchdarkly:index/metric:Metric example example-project/example-metric-key
// ```
type Metric struct {
	pulumi.CustomResourceState

	// The method for analyzing metric events. Available choices are `mean` and `percentile`.
	AnalysisType pulumi.StringPtrOutput `pulumi:"analysisType"`
	// The description of the metric's purpose.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The event key for your metric (if custom metric)
	EventKey pulumi.StringPtrOutput `pulumi:"eventKey"`
	// Include units that did not send any events and set their value to 0.
	IncludeUnitsWithoutEvents pulumi.BoolOutput `pulumi:"includeUnitsWithoutEvents"`
	// Ignored. All metrics are considered active.
	//
	// Deprecated: No longer in use. This field will be removed in a future major release of the LaunchDarkly provider.
	IsActive pulumi.BoolOutput `pulumi:"isActive"`
	// Whether a `custom` metric is a numeric metric or not.
	IsNumeric pulumi.BoolPtrOutput `pulumi:"isNumeric"`
	// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key pulumi.StringOutput `pulumi:"key"`
	// The metric type. Available choices are `click`, `custom`, and `pageview`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Kind         pulumi.StringOutput `pulumi:"kind"`
	MaintainerId pulumi.StringOutput `pulumi:"maintainerId"`
	// The human-friendly name for the metric.
	Name pulumi.StringOutput `pulumi:"name"`
	// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when analysisType is percentile.
	PercentileValue pulumi.IntPtrOutput `pulumi:"percentileValue"`
	// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey pulumi.StringOutput `pulumi:"projectKey"`
	// A set of one or more context kinds that this metric can measure events from. Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
	RandomizationUnits pulumi.StringArrayOutput `pulumi:"randomizationUnits"`
	// The CSS selector for your metric (if click metric)
	Selector pulumi.StringPtrOutput `pulumi:"selector"`
	// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
	SuccessCriteria pulumi.StringOutput `pulumi:"successCriteria"`
	// Tags associated with your resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// (Required for kind `custom`) The unit for numeric `custom` metrics.
	Unit pulumi.StringPtrOutput `pulumi:"unit"`
	// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
	UnitAggregationType pulumi.StringPtrOutput `pulumi:"unitAggregationType"`
	// List of nested `url` blocks describing URLs that you want to associate with the metric.
	Urls MetricUrlArrayOutput `pulumi:"urls"`
	// Version of the metric
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewMetric registers a new resource with the given unique name, arguments, and options.
func NewMetric(ctx *pulumi.Context,
	name string, args *MetricArgs, opts ...pulumi.ResourceOption) (*Metric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ProjectKey == nil {
		return nil, errors.New("invalid value for required argument 'ProjectKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Metric
	err := ctx.RegisterResource("launchdarkly:index/metric:Metric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetric gets an existing Metric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricState, opts ...pulumi.ResourceOption) (*Metric, error) {
	var resource Metric
	err := ctx.ReadResource("launchdarkly:index/metric:Metric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Metric resources.
type metricState struct {
	// The method for analyzing metric events. Available choices are `mean` and `percentile`.
	AnalysisType *string `pulumi:"analysisType"`
	// The description of the metric's purpose.
	Description *string `pulumi:"description"`
	// The event key for your metric (if custom metric)
	EventKey *string `pulumi:"eventKey"`
	// Include units that did not send any events and set their value to 0.
	IncludeUnitsWithoutEvents *bool `pulumi:"includeUnitsWithoutEvents"`
	// Ignored. All metrics are considered active.
	//
	// Deprecated: No longer in use. This field will be removed in a future major release of the LaunchDarkly provider.
	IsActive *bool `pulumi:"isActive"`
	// Whether a `custom` metric is a numeric metric or not.
	IsNumeric *bool `pulumi:"isNumeric"`
	// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key *string `pulumi:"key"`
	// The metric type. Available choices are `click`, `custom`, and `pageview`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Kind         *string `pulumi:"kind"`
	MaintainerId *string `pulumi:"maintainerId"`
	// The human-friendly name for the metric.
	Name *string `pulumi:"name"`
	// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when analysisType is percentile.
	PercentileValue *int `pulumi:"percentileValue"`
	// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey *string `pulumi:"projectKey"`
	// A set of one or more context kinds that this metric can measure events from. Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
	RandomizationUnits []string `pulumi:"randomizationUnits"`
	// The CSS selector for your metric (if click metric)
	Selector *string `pulumi:"selector"`
	// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
	SuccessCriteria *string `pulumi:"successCriteria"`
	// Tags associated with your resource.
	Tags []string `pulumi:"tags"`
	// (Required for kind `custom`) The unit for numeric `custom` metrics.
	Unit *string `pulumi:"unit"`
	// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
	UnitAggregationType *string `pulumi:"unitAggregationType"`
	// List of nested `url` blocks describing URLs that you want to associate with the metric.
	Urls []MetricUrl `pulumi:"urls"`
	// Version of the metric
	Version *int `pulumi:"version"`
}

type MetricState struct {
	// The method for analyzing metric events. Available choices are `mean` and `percentile`.
	AnalysisType pulumi.StringPtrInput
	// The description of the metric's purpose.
	Description pulumi.StringPtrInput
	// The event key for your metric (if custom metric)
	EventKey pulumi.StringPtrInput
	// Include units that did not send any events and set their value to 0.
	IncludeUnitsWithoutEvents pulumi.BoolPtrInput
	// Ignored. All metrics are considered active.
	//
	// Deprecated: No longer in use. This field will be removed in a future major release of the LaunchDarkly provider.
	IsActive pulumi.BoolPtrInput
	// Whether a `custom` metric is a numeric metric or not.
	IsNumeric pulumi.BoolPtrInput
	// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key pulumi.StringPtrInput
	// The metric type. Available choices are `click`, `custom`, and `pageview`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Kind         pulumi.StringPtrInput
	MaintainerId pulumi.StringPtrInput
	// The human-friendly name for the metric.
	Name pulumi.StringPtrInput
	// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when analysisType is percentile.
	PercentileValue pulumi.IntPtrInput
	// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey pulumi.StringPtrInput
	// A set of one or more context kinds that this metric can measure events from. Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
	RandomizationUnits pulumi.StringArrayInput
	// The CSS selector for your metric (if click metric)
	Selector pulumi.StringPtrInput
	// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
	SuccessCriteria pulumi.StringPtrInput
	// Tags associated with your resource.
	Tags pulumi.StringArrayInput
	// (Required for kind `custom`) The unit for numeric `custom` metrics.
	Unit pulumi.StringPtrInput
	// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
	UnitAggregationType pulumi.StringPtrInput
	// List of nested `url` blocks describing URLs that you want to associate with the metric.
	Urls MetricUrlArrayInput
	// Version of the metric
	Version pulumi.IntPtrInput
}

func (MetricState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricState)(nil)).Elem()
}

type metricArgs struct {
	// The method for analyzing metric events. Available choices are `mean` and `percentile`.
	AnalysisType *string `pulumi:"analysisType"`
	// The description of the metric's purpose.
	Description *string `pulumi:"description"`
	// The event key for your metric (if custom metric)
	EventKey *string `pulumi:"eventKey"`
	// Include units that did not send any events and set their value to 0.
	IncludeUnitsWithoutEvents *bool `pulumi:"includeUnitsWithoutEvents"`
	// Ignored. All metrics are considered active.
	//
	// Deprecated: No longer in use. This field will be removed in a future major release of the LaunchDarkly provider.
	IsActive *bool `pulumi:"isActive"`
	// Whether a `custom` metric is a numeric metric or not.
	IsNumeric *bool `pulumi:"isNumeric"`
	// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key string `pulumi:"key"`
	// The metric type. Available choices are `click`, `custom`, and `pageview`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Kind         string  `pulumi:"kind"`
	MaintainerId *string `pulumi:"maintainerId"`
	// The human-friendly name for the metric.
	Name *string `pulumi:"name"`
	// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when analysisType is percentile.
	PercentileValue *int `pulumi:"percentileValue"`
	// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey string `pulumi:"projectKey"`
	// A set of one or more context kinds that this metric can measure events from. Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
	RandomizationUnits []string `pulumi:"randomizationUnits"`
	// The CSS selector for your metric (if click metric)
	Selector *string `pulumi:"selector"`
	// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
	SuccessCriteria *string `pulumi:"successCriteria"`
	// Tags associated with your resource.
	Tags []string `pulumi:"tags"`
	// (Required for kind `custom`) The unit for numeric `custom` metrics.
	Unit *string `pulumi:"unit"`
	// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
	UnitAggregationType *string `pulumi:"unitAggregationType"`
	// List of nested `url` blocks describing URLs that you want to associate with the metric.
	Urls []MetricUrl `pulumi:"urls"`
}

// The set of arguments for constructing a Metric resource.
type MetricArgs struct {
	// The method for analyzing metric events. Available choices are `mean` and `percentile`.
	AnalysisType pulumi.StringPtrInput
	// The description of the metric's purpose.
	Description pulumi.StringPtrInput
	// The event key for your metric (if custom metric)
	EventKey pulumi.StringPtrInput
	// Include units that did not send any events and set their value to 0.
	IncludeUnitsWithoutEvents pulumi.BoolPtrInput
	// Ignored. All metrics are considered active.
	//
	// Deprecated: No longer in use. This field will be removed in a future major release of the LaunchDarkly provider.
	IsActive pulumi.BoolPtrInput
	// Whether a `custom` metric is a numeric metric or not.
	IsNumeric pulumi.BoolPtrInput
	// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key pulumi.StringInput
	// The metric type. Available choices are `click`, `custom`, and `pageview`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Kind         pulumi.StringInput
	MaintainerId pulumi.StringPtrInput
	// The human-friendly name for the metric.
	Name pulumi.StringPtrInput
	// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when analysisType is percentile.
	PercentileValue pulumi.IntPtrInput
	// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey pulumi.StringInput
	// A set of one or more context kinds that this metric can measure events from. Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
	RandomizationUnits pulumi.StringArrayInput
	// The CSS selector for your metric (if click metric)
	Selector pulumi.StringPtrInput
	// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
	SuccessCriteria pulumi.StringPtrInput
	// Tags associated with your resource.
	Tags pulumi.StringArrayInput
	// (Required for kind `custom`) The unit for numeric `custom` metrics.
	Unit pulumi.StringPtrInput
	// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
	UnitAggregationType pulumi.StringPtrInput
	// List of nested `url` blocks describing URLs that you want to associate with the metric.
	Urls MetricUrlArrayInput
}

func (MetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricArgs)(nil)).Elem()
}

type MetricInput interface {
	pulumi.Input

	ToMetricOutput() MetricOutput
	ToMetricOutputWithContext(ctx context.Context) MetricOutput
}

func (*Metric) ElementType() reflect.Type {
	return reflect.TypeOf((**Metric)(nil)).Elem()
}

func (i *Metric) ToMetricOutput() MetricOutput {
	return i.ToMetricOutputWithContext(context.Background())
}

func (i *Metric) ToMetricOutputWithContext(ctx context.Context) MetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricOutput)
}

// MetricArrayInput is an input type that accepts MetricArray and MetricArrayOutput values.
// You can construct a concrete instance of `MetricArrayInput` via:
//
//	MetricArray{ MetricArgs{...} }
type MetricArrayInput interface {
	pulumi.Input

	ToMetricArrayOutput() MetricArrayOutput
	ToMetricArrayOutputWithContext(context.Context) MetricArrayOutput
}

type MetricArray []MetricInput

func (MetricArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Metric)(nil)).Elem()
}

func (i MetricArray) ToMetricArrayOutput() MetricArrayOutput {
	return i.ToMetricArrayOutputWithContext(context.Background())
}

func (i MetricArray) ToMetricArrayOutputWithContext(ctx context.Context) MetricArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricArrayOutput)
}

// MetricMapInput is an input type that accepts MetricMap and MetricMapOutput values.
// You can construct a concrete instance of `MetricMapInput` via:
//
//	MetricMap{ "key": MetricArgs{...} }
type MetricMapInput interface {
	pulumi.Input

	ToMetricMapOutput() MetricMapOutput
	ToMetricMapOutputWithContext(context.Context) MetricMapOutput
}

type MetricMap map[string]MetricInput

func (MetricMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Metric)(nil)).Elem()
}

func (i MetricMap) ToMetricMapOutput() MetricMapOutput {
	return i.ToMetricMapOutputWithContext(context.Background())
}

func (i MetricMap) ToMetricMapOutputWithContext(ctx context.Context) MetricMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricMapOutput)
}

type MetricOutput struct{ *pulumi.OutputState }

func (MetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Metric)(nil)).Elem()
}

func (o MetricOutput) ToMetricOutput() MetricOutput {
	return o
}

func (o MetricOutput) ToMetricOutputWithContext(ctx context.Context) MetricOutput {
	return o
}

// The method for analyzing metric events. Available choices are `mean` and `percentile`.
func (o MetricOutput) AnalysisType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringPtrOutput { return v.AnalysisType }).(pulumi.StringPtrOutput)
}

// The description of the metric's purpose.
func (o MetricOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The event key for your metric (if custom metric)
func (o MetricOutput) EventKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringPtrOutput { return v.EventKey }).(pulumi.StringPtrOutput)
}

// Include units that did not send any events and set their value to 0.
func (o MetricOutput) IncludeUnitsWithoutEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *Metric) pulumi.BoolOutput { return v.IncludeUnitsWithoutEvents }).(pulumi.BoolOutput)
}

// Ignored. All metrics are considered active.
//
// Deprecated: No longer in use. This field will be removed in a future major release of the LaunchDarkly provider.
func (o MetricOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v *Metric) pulumi.BoolOutput { return v.IsActive }).(pulumi.BoolOutput)
}

// Whether a `custom` metric is a numeric metric or not.
func (o MetricOutput) IsNumeric() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Metric) pulumi.BoolPtrOutput { return v.IsNumeric }).(pulumi.BoolPtrOutput)
}

// The unique key that references the metric. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
func (o MetricOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The metric type. Available choices are `click`, `custom`, and `pageview`. A change in this field will force the destruction of the existing resource and the creation of a new one.
func (o MetricOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o MetricOutput) MaintainerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringOutput { return v.MaintainerId }).(pulumi.StringOutput)
}

// The human-friendly name for the metric.
func (o MetricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when analysisType is percentile.
func (o MetricOutput) PercentileValue() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Metric) pulumi.IntPtrOutput { return v.PercentileValue }).(pulumi.IntPtrOutput)
}

// The metrics's project key. A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
func (o MetricOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// A set of one or more context kinds that this metric can measure events from. Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
func (o MetricOutput) RandomizationUnits() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringArrayOutput { return v.RandomizationUnits }).(pulumi.StringArrayOutput)
}

// The CSS selector for your metric (if click metric)
func (o MetricOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringPtrOutput { return v.Selector }).(pulumi.StringPtrOutput)
}

// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
func (o MetricOutput) SuccessCriteria() pulumi.StringOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringOutput { return v.SuccessCriteria }).(pulumi.StringOutput)
}

// Tags associated with your resource.
func (o MetricOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// (Required for kind `custom`) The unit for numeric `custom` metrics.
func (o MetricOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringPtrOutput { return v.Unit }).(pulumi.StringPtrOutput)
}

// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
func (o MetricOutput) UnitAggregationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metric) pulumi.StringPtrOutput { return v.UnitAggregationType }).(pulumi.StringPtrOutput)
}

// List of nested `url` blocks describing URLs that you want to associate with the metric.
func (o MetricOutput) Urls() MetricUrlArrayOutput {
	return o.ApplyT(func(v *Metric) MetricUrlArrayOutput { return v.Urls }).(MetricUrlArrayOutput)
}

// Version of the metric
func (o MetricOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *Metric) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type MetricArrayOutput struct{ *pulumi.OutputState }

func (MetricArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Metric)(nil)).Elem()
}

func (o MetricArrayOutput) ToMetricArrayOutput() MetricArrayOutput {
	return o
}

func (o MetricArrayOutput) ToMetricArrayOutputWithContext(ctx context.Context) MetricArrayOutput {
	return o
}

func (o MetricArrayOutput) Index(i pulumi.IntInput) MetricOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Metric {
		return vs[0].([]*Metric)[vs[1].(int)]
	}).(MetricOutput)
}

type MetricMapOutput struct{ *pulumi.OutputState }

func (MetricMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Metric)(nil)).Elem()
}

func (o MetricMapOutput) ToMetricMapOutput() MetricMapOutput {
	return o
}

func (o MetricMapOutput) ToMetricMapOutputWithContext(ctx context.Context) MetricMapOutput {
	return o
}

func (o MetricMapOutput) MapIndex(k pulumi.StringInput) MetricOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Metric {
		return vs[0].(map[string]*Metric)[vs[1].(string)]
	}).(MetricOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetricInput)(nil)).Elem(), &Metric{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricArrayInput)(nil)).Elem(), MetricArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricMapInput)(nil)).Elem(), MetricMap{})
	pulumi.RegisterOutputType(MetricOutput{})
	pulumi.RegisterOutputType(MetricArrayOutput{})
	pulumi.RegisterOutputType(MetricMapOutput{})
}
