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
 * const example = new launchdarkly.RelayProxyConfiguration("example", {
 *     name: "example-config",
 *     policies: [{
 *         actions: ["*"],
 *         effect: "allow",
 *         resources: ["proj/*:env/*"],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Relay Proxy configurations can be imported using the configuration's unique 24 character ID, e.g.
 *
 * ```sh
 * $ pulumi import launchdarkly:index/relayProxyConfiguration:RelayProxyConfiguration example 51d440e30c9ff61457c710f6
 * ```
 *
 * The unique relay proxy ID can be found in the relay proxy edit page URL, which you can locate by clicking the three dot menu on your relay proxy item in the UI and selecting 'Edit configuration':
 *
 * https://app.launchdarkly.com/settings/relay/THIS_IS_YOUR_RELAY_PROXY_ID/edit
 */
export class RelayProxyConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing RelayProxyConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RelayProxyConfigurationState, opts?: pulumi.CustomResourceOptions): RelayProxyConfiguration {
        return new RelayProxyConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'launchdarkly:index/relayProxyConfiguration:RelayProxyConfiguration';

    /**
     * Returns true if the given object is an instance of RelayProxyConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RelayProxyConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RelayProxyConfiguration.__pulumiType;
    }

    /**
     * The last 4 characters of the Relay Proxy configuration's unique key.
     */
    public /*out*/ readonly displayKey!: pulumi.Output<string>;
    /**
     * The Relay Proxy configuration's unique key. Because the `fullKey` is only exposed upon creation, it will not be available if the resource is imported.
     */
    public /*out*/ readonly fullKey!: pulumi.Output<string>;
    /**
     * The human-readable name for your Relay Proxy configuration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Relay Proxy configuration's rule policy block. This determines what content the Relay Proxy receives. To learn more, read [Understanding policies](https://docs.launchdarkly.com/home/members/role-policies#understanding-policies).
     */
    public readonly policies!: pulumi.Output<outputs.RelayProxyConfigurationPolicy[]>;

    /**
     * Create a RelayProxyConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RelayProxyConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RelayProxyConfigurationArgs | RelayProxyConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RelayProxyConfigurationState | undefined;
            resourceInputs["displayKey"] = state ? state.displayKey : undefined;
            resourceInputs["fullKey"] = state ? state.fullKey : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
        } else {
            const args = argsOrState as RelayProxyConfigurationArgs | undefined;
            if ((!args || args.policies === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policies'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["displayKey"] = undefined /*out*/;
            resourceInputs["fullKey"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["fullKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(RelayProxyConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RelayProxyConfiguration resources.
 */
export interface RelayProxyConfigurationState {
    /**
     * The last 4 characters of the Relay Proxy configuration's unique key.
     */
    displayKey?: pulumi.Input<string>;
    /**
     * The Relay Proxy configuration's unique key. Because the `fullKey` is only exposed upon creation, it will not be available if the resource is imported.
     */
    fullKey?: pulumi.Input<string>;
    /**
     * The human-readable name for your Relay Proxy configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * The Relay Proxy configuration's rule policy block. This determines what content the Relay Proxy receives. To learn more, read [Understanding policies](https://docs.launchdarkly.com/home/members/role-policies#understanding-policies).
     */
    policies?: pulumi.Input<pulumi.Input<inputs.RelayProxyConfigurationPolicy>[]>;
}

/**
 * The set of arguments for constructing a RelayProxyConfiguration resource.
 */
export interface RelayProxyConfigurationArgs {
    /**
     * The human-readable name for your Relay Proxy configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * The Relay Proxy configuration's rule policy block. This determines what content the Relay Proxy receives. To learn more, read [Understanding policies](https://docs.launchdarkly.com/home/members/role-policies#understanding-policies).
     */
    policies: pulumi.Input<pulumi.Input<inputs.RelayProxyConfigurationPolicy>[]>;
}
