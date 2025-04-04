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
//			_, err := launchdarkly.NewEnvironment(ctx, "staging", &launchdarkly.EnvironmentArgs{
//				Name:  pulumi.String("Staging"),
//				Key:   pulumi.String("staging"),
//				Color: pulumi.String("ff00ff"),
//				Tags: pulumi.StringArray{
//					pulumi.String("terraform"),
//					pulumi.String("staging"),
//				},
//				ProjectKey: pulumi.Any(example.Key),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = launchdarkly.NewEnvironment(ctx, "approvals_example", &launchdarkly.EnvironmentArgs{
//				Name:  pulumi.String("Approvals Example Environment"),
//				Key:   pulumi.String("approvals-example"),
//				Color: pulumi.String("ff00ff"),
//				Tags: pulumi.StringArray{
//					pulumi.String("terraform"),
//					pulumi.String("staging"),
//				},
//				ApprovalSettings: launchdarkly.EnvironmentApprovalSettingArray{
//					&launchdarkly.EnvironmentApprovalSettingArgs{
//						Required:                pulumi.Bool(true),
//						CanReviewOwnRequest:     pulumi.Bool(true),
//						MinNumApprovals:         pulumi.Int(2),
//						CanApplyDeclinedChanges: pulumi.Bool(true),
//					},
//				},
//				ProjectKey: pulumi.Any(example.Key),
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
// Import a LaunchDarkly environment using this format: `project_key/environment_key`.
//
// ```sh
// $ pulumi import launchdarkly:index/environment:Environment staging example-project/staging
// ```
type Environment struct {
	pulumi.CustomResourceState

	// The environment's SDK key.
	ApiKey           pulumi.StringOutput                   `pulumi:"apiKey"`
	ApprovalSettings EnvironmentApprovalSettingArrayOutput `pulumi:"approvalSettings"`
	// The environment's client-side ID.
	ClientSideId pulumi.StringOutput `pulumi:"clientSideId"`
	// The color swatch as an RGB hex value with no leading `#`. For example: `000000`
	Color pulumi.StringOutput `pulumi:"color"`
	// Set to `true` if this environment requires confirmation for flag and segment changes. This field will default to `false`
	// when not set.
	ConfirmChanges pulumi.BoolPtrOutput `pulumi:"confirmChanges"`
	// Denotes whether the environment is critical.
	Critical pulumi.BoolPtrOutput `pulumi:"critical"`
	// Set to `true` to enable data export for every flag created in this environment after you configure this argument. This
	// field will default to `false` when not set. To learn more, read [Data
	// Export](https://docs.launchdarkly.com/home/data-export).
	DefaultTrackEvents pulumi.BoolPtrOutput `pulumi:"defaultTrackEvents"`
	// The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using
	// the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL
	// settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
	DefaultTtl pulumi.IntPtrOutput `pulumi:"defaultTtl"`
	// The project-unique key for the environment. A change in this field will force the destruction of the existing resource
	// and the creation of a new one.
	Key pulumi.StringOutput `pulumi:"key"`
	// The environment's mobile key.
	MobileKey pulumi.StringOutput `pulumi:"mobileKey"`
	// The name of the environment.
	Name pulumi.StringOutput `pulumi:"name"`
	// The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the
	// creation of a new one.
	ProjectKey pulumi.StringOutput `pulumi:"projectKey"`
	// Set to `true` if this environment requires comments for flag and segment changes. This field will default to `false`
	// when not set.
	RequireComments pulumi.BoolPtrOutput `pulumi:"requireComments"`
	// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user. This field will default to
	// `false` when not set.
	SecureMode pulumi.BoolPtrOutput `pulumi:"secureMode"`
	// Tags associated with your resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Color == nil {
		return nil, errors.New("invalid value for required argument 'Color'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.ProjectKey == nil {
		return nil, errors.New("invalid value for required argument 'ProjectKey'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiKey",
		"clientSideId",
		"mobileKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Environment
	err := ctx.RegisterResource("launchdarkly:index/environment:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("launchdarkly:index/environment:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// The environment's SDK key.
	ApiKey           *string                      `pulumi:"apiKey"`
	ApprovalSettings []EnvironmentApprovalSetting `pulumi:"approvalSettings"`
	// The environment's client-side ID.
	ClientSideId *string `pulumi:"clientSideId"`
	// The color swatch as an RGB hex value with no leading `#`. For example: `000000`
	Color *string `pulumi:"color"`
	// Set to `true` if this environment requires confirmation for flag and segment changes. This field will default to `false`
	// when not set.
	ConfirmChanges *bool `pulumi:"confirmChanges"`
	// Denotes whether the environment is critical.
	Critical *bool `pulumi:"critical"`
	// Set to `true` to enable data export for every flag created in this environment after you configure this argument. This
	// field will default to `false` when not set. To learn more, read [Data
	// Export](https://docs.launchdarkly.com/home/data-export).
	DefaultTrackEvents *bool `pulumi:"defaultTrackEvents"`
	// The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using
	// the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL
	// settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
	DefaultTtl *int `pulumi:"defaultTtl"`
	// The project-unique key for the environment. A change in this field will force the destruction of the existing resource
	// and the creation of a new one.
	Key *string `pulumi:"key"`
	// The environment's mobile key.
	MobileKey *string `pulumi:"mobileKey"`
	// The name of the environment.
	Name *string `pulumi:"name"`
	// The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the
	// creation of a new one.
	ProjectKey *string `pulumi:"projectKey"`
	// Set to `true` if this environment requires comments for flag and segment changes. This field will default to `false`
	// when not set.
	RequireComments *bool `pulumi:"requireComments"`
	// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user. This field will default to
	// `false` when not set.
	SecureMode *bool `pulumi:"secureMode"`
	// Tags associated with your resource.
	Tags []string `pulumi:"tags"`
}

type EnvironmentState struct {
	// The environment's SDK key.
	ApiKey           pulumi.StringPtrInput
	ApprovalSettings EnvironmentApprovalSettingArrayInput
	// The environment's client-side ID.
	ClientSideId pulumi.StringPtrInput
	// The color swatch as an RGB hex value with no leading `#`. For example: `000000`
	Color pulumi.StringPtrInput
	// Set to `true` if this environment requires confirmation for flag and segment changes. This field will default to `false`
	// when not set.
	ConfirmChanges pulumi.BoolPtrInput
	// Denotes whether the environment is critical.
	Critical pulumi.BoolPtrInput
	// Set to `true` to enable data export for every flag created in this environment after you configure this argument. This
	// field will default to `false` when not set. To learn more, read [Data
	// Export](https://docs.launchdarkly.com/home/data-export).
	DefaultTrackEvents pulumi.BoolPtrInput
	// The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using
	// the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL
	// settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
	DefaultTtl pulumi.IntPtrInput
	// The project-unique key for the environment. A change in this field will force the destruction of the existing resource
	// and the creation of a new one.
	Key pulumi.StringPtrInput
	// The environment's mobile key.
	MobileKey pulumi.StringPtrInput
	// The name of the environment.
	Name pulumi.StringPtrInput
	// The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the
	// creation of a new one.
	ProjectKey pulumi.StringPtrInput
	// Set to `true` if this environment requires comments for flag and segment changes. This field will default to `false`
	// when not set.
	RequireComments pulumi.BoolPtrInput
	// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user. This field will default to
	// `false` when not set.
	SecureMode pulumi.BoolPtrInput
	// Tags associated with your resource.
	Tags pulumi.StringArrayInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	ApprovalSettings []EnvironmentApprovalSetting `pulumi:"approvalSettings"`
	// The color swatch as an RGB hex value with no leading `#`. For example: `000000`
	Color string `pulumi:"color"`
	// Set to `true` if this environment requires confirmation for flag and segment changes. This field will default to `false`
	// when not set.
	ConfirmChanges *bool `pulumi:"confirmChanges"`
	// Denotes whether the environment is critical.
	Critical *bool `pulumi:"critical"`
	// Set to `true` to enable data export for every flag created in this environment after you configure this argument. This
	// field will default to `false` when not set. To learn more, read [Data
	// Export](https://docs.launchdarkly.com/home/data-export).
	DefaultTrackEvents *bool `pulumi:"defaultTrackEvents"`
	// The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using
	// the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL
	// settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
	DefaultTtl *int `pulumi:"defaultTtl"`
	// The project-unique key for the environment. A change in this field will force the destruction of the existing resource
	// and the creation of a new one.
	Key string `pulumi:"key"`
	// The name of the environment.
	Name *string `pulumi:"name"`
	// The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the
	// creation of a new one.
	ProjectKey string `pulumi:"projectKey"`
	// Set to `true` if this environment requires comments for flag and segment changes. This field will default to `false`
	// when not set.
	RequireComments *bool `pulumi:"requireComments"`
	// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user. This field will default to
	// `false` when not set.
	SecureMode *bool `pulumi:"secureMode"`
	// Tags associated with your resource.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	ApprovalSettings EnvironmentApprovalSettingArrayInput
	// The color swatch as an RGB hex value with no leading `#`. For example: `000000`
	Color pulumi.StringInput
	// Set to `true` if this environment requires confirmation for flag and segment changes. This field will default to `false`
	// when not set.
	ConfirmChanges pulumi.BoolPtrInput
	// Denotes whether the environment is critical.
	Critical pulumi.BoolPtrInput
	// Set to `true` to enable data export for every flag created in this environment after you configure this argument. This
	// field will default to `false` when not set. To learn more, read [Data
	// Export](https://docs.launchdarkly.com/home/data-export).
	DefaultTrackEvents pulumi.BoolPtrInput
	// The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using
	// the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL
	// settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
	DefaultTtl pulumi.IntPtrInput
	// The project-unique key for the environment. A change in this field will force the destruction of the existing resource
	// and the creation of a new one.
	Key pulumi.StringInput
	// The name of the environment.
	Name pulumi.StringPtrInput
	// The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the
	// creation of a new one.
	ProjectKey pulumi.StringInput
	// Set to `true` if this environment requires comments for flag and segment changes. This field will default to `false`
	// when not set.
	RequireComments pulumi.BoolPtrInput
	// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user. This field will default to
	// `false` when not set.
	SecureMode pulumi.BoolPtrInput
	// Tags associated with your resource.
	Tags pulumi.StringArrayInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

// EnvironmentArrayInput is an input type that accepts EnvironmentArray and EnvironmentArrayOutput values.
// You can construct a concrete instance of `EnvironmentArrayInput` via:
//
//	EnvironmentArray{ EnvironmentArgs{...} }
type EnvironmentArrayInput interface {
	pulumi.Input

	ToEnvironmentArrayOutput() EnvironmentArrayOutput
	ToEnvironmentArrayOutputWithContext(context.Context) EnvironmentArrayOutput
}

type EnvironmentArray []EnvironmentInput

func (EnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Environment)(nil)).Elem()
}

func (i EnvironmentArray) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return i.ToEnvironmentArrayOutputWithContext(context.Background())
}

func (i EnvironmentArray) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentArrayOutput)
}

// EnvironmentMapInput is an input type that accepts EnvironmentMap and EnvironmentMapOutput values.
// You can construct a concrete instance of `EnvironmentMapInput` via:
//
//	EnvironmentMap{ "key": EnvironmentArgs{...} }
type EnvironmentMapInput interface {
	pulumi.Input

	ToEnvironmentMapOutput() EnvironmentMapOutput
	ToEnvironmentMapOutputWithContext(context.Context) EnvironmentMapOutput
}

type EnvironmentMap map[string]EnvironmentInput

func (EnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Environment)(nil)).Elem()
}

func (i EnvironmentMap) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return i.ToEnvironmentMapOutputWithContext(context.Background())
}

func (i EnvironmentMap) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentMapOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

// The environment's SDK key.
func (o EnvironmentOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.ApiKey }).(pulumi.StringOutput)
}

func (o EnvironmentOutput) ApprovalSettings() EnvironmentApprovalSettingArrayOutput {
	return o.ApplyT(func(v *Environment) EnvironmentApprovalSettingArrayOutput { return v.ApprovalSettings }).(EnvironmentApprovalSettingArrayOutput)
}

// The environment's client-side ID.
func (o EnvironmentOutput) ClientSideId() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.ClientSideId }).(pulumi.StringOutput)
}

// The color swatch as an RGB hex value with no leading `#`. For example: `000000`
func (o EnvironmentOutput) Color() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Color }).(pulumi.StringOutput)
}

// Set to `true` if this environment requires confirmation for flag and segment changes. This field will default to `false`
// when not set.
func (o EnvironmentOutput) ConfirmChanges() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.BoolPtrOutput { return v.ConfirmChanges }).(pulumi.BoolPtrOutput)
}

// Denotes whether the environment is critical.
func (o EnvironmentOutput) Critical() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.BoolPtrOutput { return v.Critical }).(pulumi.BoolPtrOutput)
}

// Set to `true` to enable data export for every flag created in this environment after you configure this argument. This
// field will default to `false` when not set. To learn more, read [Data
// Export](https://docs.launchdarkly.com/home/data-export).
func (o EnvironmentOutput) DefaultTrackEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.BoolPtrOutput { return v.DefaultTrackEvents }).(pulumi.BoolPtrOutput)
}

// The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using
// the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL
// settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
func (o EnvironmentOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.IntPtrOutput { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

// The project-unique key for the environment. A change in this field will force the destruction of the existing resource
// and the creation of a new one.
func (o EnvironmentOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The environment's mobile key.
func (o EnvironmentOutput) MobileKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.MobileKey }).(pulumi.StringOutput)
}

// The name of the environment.
func (o EnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the
// creation of a new one.
func (o EnvironmentOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// Set to `true` if this environment requires comments for flag and segment changes. This field will default to `false`
// when not set.
func (o EnvironmentOutput) RequireComments() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.BoolPtrOutput { return v.RequireComments }).(pulumi.BoolPtrOutput)
}

// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user. This field will default to
// `false` when not set.
func (o EnvironmentOutput) SecureMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.BoolPtrOutput { return v.SecureMode }).(pulumi.BoolPtrOutput)
}

// Tags associated with your resource.
func (o EnvironmentOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type EnvironmentArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Environment)(nil)).Elem()
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) Index(i pulumi.IntInput) EnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Environment {
		return vs[0].([]*Environment)[vs[1].(int)]
	}).(EnvironmentOutput)
}

type EnvironmentMapOutput struct{ *pulumi.OutputState }

func (EnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Environment)(nil)).Elem()
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) MapIndex(k pulumi.StringInput) EnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Environment {
		return vs[0].(map[string]*Environment)[vs[1].(string)]
	}).(EnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentInput)(nil)).Elem(), &Environment{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentArrayInput)(nil)).Elem(), EnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentMapInput)(nil)).Elem(), EnvironmentMap{})
	pulumi.RegisterOutputType(EnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentMapOutput{})
}
