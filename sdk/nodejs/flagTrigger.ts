// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as launchdarkly from "@pulumi/launchdarkly";
 *
 * const example = new launchdarkly.FlagTrigger("example", {
 *     projectKey: exampleLaunchdarklyProject.key,
 *     envKey: "test",
 *     flagKey: triggerFlag.key,
 *     integrationKey: "generic-trigger",
 *     instructions: {
 *         kind: "turnFlagOn",
 *     },
 *     enabled: false,
 * });
 * ```
 *
 * ## Import
 *
 * Import a LaunchDarkly flag trigger using this format: `project_key/environment_key/flag_key/trigger_id`.
 *
 * ```sh
 * $ pulumi import launchdarkly:index/flagTrigger:FlagTrigger example example-project-key/example-env-key/example-flag-key/62581d4488def814b831abc3
 * ```
 *
 * The unique trigger ID can be found in your saved trigger URL:
 *
 * https://app.launchdarkly.com/webhook/triggers/THIS_IS_YOUR_TRIGGER_ID/aff25a53-17d9-4112-a9b8-12718d1a2e79
 *
 * Please note that if you did not save this upon creation of the resource, you will have to reset it to get a new value, which can cause breaking changes.
 */
export class FlagTrigger extends pulumi.CustomResource {
    /**
     * Get an existing FlagTrigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlagTriggerState, opts?: pulumi.CustomResourceOptions): FlagTrigger {
        return new FlagTrigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'launchdarkly:index/flagTrigger:FlagTrigger';

    /**
     * Returns true if the given object is an instance of FlagTrigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlagTrigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlagTrigger.__pulumiType;
    }

    /**
     * Whether the trigger is currently active or not.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    public readonly envKey!: pulumi.Output<string>;
    /**
     * The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    public readonly flagKey!: pulumi.Output<string>;
    /**
     * Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "<flag_action>" }`.
     */
    public readonly instructions!: pulumi.Output<outputs.FlagTriggerInstructions>;
    /**
     * The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    public readonly integrationKey!: pulumi.Output<string>;
    public /*out*/ readonly maintainerId!: pulumi.Output<string>;
    /**
     * The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    public readonly projectKey!: pulumi.Output<string>;
    /**
     * The unique URL used to invoke the trigger.
     */
    public /*out*/ readonly triggerUrl!: pulumi.Output<string>;

    /**
     * Create a FlagTrigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlagTriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FlagTriggerArgs | FlagTriggerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FlagTriggerState | undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["envKey"] = state ? state.envKey : undefined;
            resourceInputs["flagKey"] = state ? state.flagKey : undefined;
            resourceInputs["instructions"] = state ? state.instructions : undefined;
            resourceInputs["integrationKey"] = state ? state.integrationKey : undefined;
            resourceInputs["maintainerId"] = state ? state.maintainerId : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["triggerUrl"] = state ? state.triggerUrl : undefined;
        } else {
            const args = argsOrState as FlagTriggerArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.envKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'envKey'");
            }
            if ((!args || args.flagKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flagKey'");
            }
            if ((!args || args.instructions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instructions'");
            }
            if ((!args || args.integrationKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationKey'");
            }
            if ((!args || args.projectKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectKey'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["envKey"] = args ? args.envKey : undefined;
            resourceInputs["flagKey"] = args ? args.flagKey : undefined;
            resourceInputs["instructions"] = args ? args.instructions : undefined;
            resourceInputs["integrationKey"] = args ? args.integrationKey : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["maintainerId"] = undefined /*out*/;
            resourceInputs["triggerUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["triggerUrl"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(FlagTrigger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FlagTrigger resources.
 */
export interface FlagTriggerState {
    /**
     * Whether the trigger is currently active or not.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    envKey?: pulumi.Input<string>;
    /**
     * The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    flagKey?: pulumi.Input<string>;
    /**
     * Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "<flag_action>" }`.
     */
    instructions?: pulumi.Input<inputs.FlagTriggerInstructions>;
    /**
     * The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    integrationKey?: pulumi.Input<string>;
    maintainerId?: pulumi.Input<string>;
    /**
     * The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * The unique URL used to invoke the trigger.
     */
    triggerUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FlagTrigger resource.
 */
export interface FlagTriggerArgs {
    /**
     * Whether the trigger is currently active or not.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The unique key of the environment the flag trigger will work in. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    envKey: pulumi.Input<string>;
    /**
     * The unique key of the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    flagKey: pulumi.Input<string>;
    /**
     * Instructions containing the action to perform when invoking the trigger. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`. This must be passed as the key-value pair `{ kind = "<flag_action>" }`.
     */
    instructions: pulumi.Input<inputs.FlagTriggerInstructions>;
    /**
     * The unique identifier of the integration you intend to set your trigger up with. Currently supported are `generic-trigger`, `datadog`, `dynatrace`, `dynatrace-cloud-automation`, `honeycomb`, `new-relic-apm`, and `signalfx`. `generic-trigger` should be used for integrations not explicitly supported. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    integrationKey: pulumi.Input<string>;
    /**
     * The unique key of the project encompassing the associated flag. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    projectKey: pulumi.Input<string>;
}
