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
//			_, err := launchdarkly.NewFlagTrigger(ctx, "example", &launchdarkly.FlagTriggerArgs{
//				ProjectKey:     pulumi.Any(exampleLaunchdarklyProject.Key),
//				EnvKey:         pulumi.String("test"),
//				FlagKey:        pulumi.Any(triggerFlag.Key),
//				IntegrationKey: pulumi.String("generic-trigger"),
//				Instructions: &launchdarkly.FlagTriggerInstructionsArgs{
//					Kind: pulumi.String("turnFlagOn"),
//				},
//				Enabled: pulumi.Bool(false),
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
// Import a LaunchDarkly flag trigger using this format: `project_key/environment_key/flag_key/trigger_id`.
//
// ```sh
// $ pulumi import launchdarkly:index/flagTrigger:FlagTrigger example example-project-key/example-env-key/example-flag-key/62581d4488def814b831abc3
// ```
//
// The unique trigger ID can be found in your saved trigger URL:
//
// https://app.launchdarkly.com/webhook/triggers/THIS_IS_YOUR_TRIGGER_ID/aff25a53-17d9-4112-a9b8-12718d1a2e79
//
// Please note that if you did not save this upon creation of the resource, you will have to reset it to get a new value, which can cause breaking changes.
type FlagTrigger struct {
	pulumi.CustomResourceState

	// Whether the trigger is currently active or not.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
	EnvKey pulumi.StringOutput `pulumi:"envKey"`
	// The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
	FlagKey pulumi.StringOutput `pulumi:"flagKey"`
	// Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "<flag_action>" }`.
	Instructions FlagTriggerInstructionsOutput `pulumi:"instructions"`
	// The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
	IntegrationKey pulumi.StringOutput `pulumi:"integrationKey"`
	MaintainerId   pulumi.StringOutput `pulumi:"maintainerId"`
	// The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey pulumi.StringOutput `pulumi:"projectKey"`
	// The unique URL used to invoke the trigger.
	TriggerUrl pulumi.StringOutput `pulumi:"triggerUrl"`
}

// NewFlagTrigger registers a new resource with the given unique name, arguments, and options.
func NewFlagTrigger(ctx *pulumi.Context,
	name string, args *FlagTriggerArgs, opts ...pulumi.ResourceOption) (*FlagTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.EnvKey == nil {
		return nil, errors.New("invalid value for required argument 'EnvKey'")
	}
	if args.FlagKey == nil {
		return nil, errors.New("invalid value for required argument 'FlagKey'")
	}
	if args.Instructions == nil {
		return nil, errors.New("invalid value for required argument 'Instructions'")
	}
	if args.IntegrationKey == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationKey'")
	}
	if args.ProjectKey == nil {
		return nil, errors.New("invalid value for required argument 'ProjectKey'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"triggerUrl",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FlagTrigger
	err := ctx.RegisterResource("launchdarkly:index/flagTrigger:FlagTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlagTrigger gets an existing FlagTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlagTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlagTriggerState, opts ...pulumi.ResourceOption) (*FlagTrigger, error) {
	var resource FlagTrigger
	err := ctx.ReadResource("launchdarkly:index/flagTrigger:FlagTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlagTrigger resources.
type flagTriggerState struct {
	// Whether the trigger is currently active or not.
	Enabled *bool `pulumi:"enabled"`
	// The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
	EnvKey *string `pulumi:"envKey"`
	// The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
	FlagKey *string `pulumi:"flagKey"`
	// Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "<flag_action>" }`.
	Instructions *FlagTriggerInstructions `pulumi:"instructions"`
	// The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
	IntegrationKey *string `pulumi:"integrationKey"`
	MaintainerId   *string `pulumi:"maintainerId"`
	// The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey *string `pulumi:"projectKey"`
	// The unique URL used to invoke the trigger.
	TriggerUrl *string `pulumi:"triggerUrl"`
}

type FlagTriggerState struct {
	// Whether the trigger is currently active or not.
	Enabled pulumi.BoolPtrInput
	// The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
	EnvKey pulumi.StringPtrInput
	// The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
	FlagKey pulumi.StringPtrInput
	// Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "<flag_action>" }`.
	Instructions FlagTriggerInstructionsPtrInput
	// The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
	IntegrationKey pulumi.StringPtrInput
	MaintainerId   pulumi.StringPtrInput
	// The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey pulumi.StringPtrInput
	// The unique URL used to invoke the trigger.
	TriggerUrl pulumi.StringPtrInput
}

func (FlagTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*flagTriggerState)(nil)).Elem()
}

type flagTriggerArgs struct {
	// Whether the trigger is currently active or not.
	Enabled bool `pulumi:"enabled"`
	// The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
	EnvKey string `pulumi:"envKey"`
	// The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
	FlagKey string `pulumi:"flagKey"`
	// Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "<flag_action>" }`.
	Instructions FlagTriggerInstructions `pulumi:"instructions"`
	// The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
	IntegrationKey string `pulumi:"integrationKey"`
	// The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey string `pulumi:"projectKey"`
}

// The set of arguments for constructing a FlagTrigger resource.
type FlagTriggerArgs struct {
	// Whether the trigger is currently active or not.
	Enabled pulumi.BoolInput
	// The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
	EnvKey pulumi.StringInput
	// The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
	FlagKey pulumi.StringInput
	// Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "<flag_action>" }`.
	Instructions FlagTriggerInstructionsInput
	// The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
	IntegrationKey pulumi.StringInput
	// The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey pulumi.StringInput
}

func (FlagTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flagTriggerArgs)(nil)).Elem()
}

type FlagTriggerInput interface {
	pulumi.Input

	ToFlagTriggerOutput() FlagTriggerOutput
	ToFlagTriggerOutputWithContext(ctx context.Context) FlagTriggerOutput
}

func (*FlagTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**FlagTrigger)(nil)).Elem()
}

func (i *FlagTrigger) ToFlagTriggerOutput() FlagTriggerOutput {
	return i.ToFlagTriggerOutputWithContext(context.Background())
}

func (i *FlagTrigger) ToFlagTriggerOutputWithContext(ctx context.Context) FlagTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlagTriggerOutput)
}

// FlagTriggerArrayInput is an input type that accepts FlagTriggerArray and FlagTriggerArrayOutput values.
// You can construct a concrete instance of `FlagTriggerArrayInput` via:
//
//	FlagTriggerArray{ FlagTriggerArgs{...} }
type FlagTriggerArrayInput interface {
	pulumi.Input

	ToFlagTriggerArrayOutput() FlagTriggerArrayOutput
	ToFlagTriggerArrayOutputWithContext(context.Context) FlagTriggerArrayOutput
}

type FlagTriggerArray []FlagTriggerInput

func (FlagTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlagTrigger)(nil)).Elem()
}

func (i FlagTriggerArray) ToFlagTriggerArrayOutput() FlagTriggerArrayOutput {
	return i.ToFlagTriggerArrayOutputWithContext(context.Background())
}

func (i FlagTriggerArray) ToFlagTriggerArrayOutputWithContext(ctx context.Context) FlagTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlagTriggerArrayOutput)
}

// FlagTriggerMapInput is an input type that accepts FlagTriggerMap and FlagTriggerMapOutput values.
// You can construct a concrete instance of `FlagTriggerMapInput` via:
//
//	FlagTriggerMap{ "key": FlagTriggerArgs{...} }
type FlagTriggerMapInput interface {
	pulumi.Input

	ToFlagTriggerMapOutput() FlagTriggerMapOutput
	ToFlagTriggerMapOutputWithContext(context.Context) FlagTriggerMapOutput
}

type FlagTriggerMap map[string]FlagTriggerInput

func (FlagTriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlagTrigger)(nil)).Elem()
}

func (i FlagTriggerMap) ToFlagTriggerMapOutput() FlagTriggerMapOutput {
	return i.ToFlagTriggerMapOutputWithContext(context.Background())
}

func (i FlagTriggerMap) ToFlagTriggerMapOutputWithContext(ctx context.Context) FlagTriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlagTriggerMapOutput)
}

type FlagTriggerOutput struct{ *pulumi.OutputState }

func (FlagTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlagTrigger)(nil)).Elem()
}

func (o FlagTriggerOutput) ToFlagTriggerOutput() FlagTriggerOutput {
	return o
}

func (o FlagTriggerOutput) ToFlagTriggerOutputWithContext(ctx context.Context) FlagTriggerOutput {
	return o
}

// Whether the trigger is currently active or not.
func (o FlagTriggerOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *FlagTrigger) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
func (o FlagTriggerOutput) EnvKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FlagTrigger) pulumi.StringOutput { return v.EnvKey }).(pulumi.StringOutput)
}

// The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
func (o FlagTriggerOutput) FlagKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FlagTrigger) pulumi.StringOutput { return v.FlagKey }).(pulumi.StringOutput)
}

// Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "<flag_action>" }`.
func (o FlagTriggerOutput) Instructions() FlagTriggerInstructionsOutput {
	return o.ApplyT(func(v *FlagTrigger) FlagTriggerInstructionsOutput { return v.Instructions }).(FlagTriggerInstructionsOutput)
}

// The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
func (o FlagTriggerOutput) IntegrationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FlagTrigger) pulumi.StringOutput { return v.IntegrationKey }).(pulumi.StringOutput)
}

func (o FlagTriggerOutput) MaintainerId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlagTrigger) pulumi.StringOutput { return v.MaintainerId }).(pulumi.StringOutput)
}

// The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
func (o FlagTriggerOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FlagTrigger) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// The unique URL used to invoke the trigger.
func (o FlagTriggerOutput) TriggerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *FlagTrigger) pulumi.StringOutput { return v.TriggerUrl }).(pulumi.StringOutput)
}

type FlagTriggerArrayOutput struct{ *pulumi.OutputState }

func (FlagTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlagTrigger)(nil)).Elem()
}

func (o FlagTriggerArrayOutput) ToFlagTriggerArrayOutput() FlagTriggerArrayOutput {
	return o
}

func (o FlagTriggerArrayOutput) ToFlagTriggerArrayOutputWithContext(ctx context.Context) FlagTriggerArrayOutput {
	return o
}

func (o FlagTriggerArrayOutput) Index(i pulumi.IntInput) FlagTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlagTrigger {
		return vs[0].([]*FlagTrigger)[vs[1].(int)]
	}).(FlagTriggerOutput)
}

type FlagTriggerMapOutput struct{ *pulumi.OutputState }

func (FlagTriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlagTrigger)(nil)).Elem()
}

func (o FlagTriggerMapOutput) ToFlagTriggerMapOutput() FlagTriggerMapOutput {
	return o
}

func (o FlagTriggerMapOutput) ToFlagTriggerMapOutputWithContext(ctx context.Context) FlagTriggerMapOutput {
	return o
}

func (o FlagTriggerMapOutput) MapIndex(k pulumi.StringInput) FlagTriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlagTrigger {
		return vs[0].(map[string]*FlagTrigger)[vs[1].(string)]
	}).(FlagTriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlagTriggerInput)(nil)).Elem(), &FlagTrigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlagTriggerArrayInput)(nil)).Elem(), FlagTriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlagTriggerMapInput)(nil)).Elem(), FlagTriggerMap{})
	pulumi.RegisterOutputType(FlagTriggerOutput{})
	pulumi.RegisterOutputType(FlagTriggerArrayOutput{})
	pulumi.RegisterOutputType(FlagTriggerMapOutput{})
}
