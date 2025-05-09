// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a LaunchDarkly audit log subscription resource.
 *
 * This resource allows you to create and manage LaunchDarkly audit log subscriptions.
 */
export class AuditLogSubscription extends pulumi.CustomResource {
    /**
     * Get an existing AuditLogSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuditLogSubscriptionState, opts?: pulumi.CustomResourceOptions): AuditLogSubscription {
        return new AuditLogSubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'launchdarkly:index/auditLogSubscription:AuditLogSubscription';

    /**
     * Returns true if the given object is an instance of AuditLogSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuditLogSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuditLogSubscription.__pulumiType;
    }

    public readonly config!: pulumi.Output<{[key: string]: string}>;
    /**
     * The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    public readonly integrationKey!: pulumi.Output<string>;
    /**
     * A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether or not you want your subscription enabled, i.e. to actively send events.
     */
    public readonly on!: pulumi.Output<boolean>;
    /**
     * A block representing the resources to which you wish to subscribe.
     */
    public readonly statements!: pulumi.Output<outputs.AuditLogSubscriptionStatement[]>;
    /**
     * Tags associated with your resource.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a AuditLogSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuditLogSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuditLogSubscriptionArgs | AuditLogSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuditLogSubscriptionState | undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["integrationKey"] = state ? state.integrationKey : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["on"] = state ? state.on : undefined;
            resourceInputs["statements"] = state ? state.statements : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AuditLogSubscriptionArgs | undefined;
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.integrationKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationKey'");
            }
            if ((!args || args.on === undefined) && !opts.urn) {
                throw new Error("Missing required property 'on'");
            }
            if ((!args || args.statements === undefined) && !opts.urn) {
                throw new Error("Missing required property 'statements'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["integrationKey"] = args ? args.integrationKey : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["on"] = args ? args.on : undefined;
            resourceInputs["statements"] = args ? args.statements : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuditLogSubscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuditLogSubscription resources.
 */
export interface AuditLogSubscriptionState {
    config?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    integrationKey?: pulumi.Input<string>;
    /**
     * A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether or not you want your subscription enabled, i.e. to actively send events.
     */
    on?: pulumi.Input<boolean>;
    /**
     * A block representing the resources to which you wish to subscribe.
     */
    statements?: pulumi.Input<pulumi.Input<inputs.AuditLogSubscriptionStatement>[]>;
    /**
     * Tags associated with your resource.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AuditLogSubscription resource.
 */
export interface AuditLogSubscriptionArgs {
    config: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    integrationKey: pulumi.Input<string>;
    /**
     * A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether or not you want your subscription enabled, i.e. to actively send events.
     */
    on: pulumi.Input<boolean>;
    /**
     * A block representing the resources to which you wish to subscribe.
     */
    statements: pulumi.Input<pulumi.Input<inputs.AuditLogSubscriptionStatement>[]>;
    /**
     * Tags associated with your resource.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
