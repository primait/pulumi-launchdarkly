// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a LaunchDarkly flag trigger data source.
 *
 * > **Note:** Flag triggers are available to customers on an Enterprise LaunchDarkly plan. To learn more, [read about our pricing](https://launchdarkly.com/pricing/). To upgrade your plan, [contact LaunchDarkly Sales](https://launchdarkly.com/contact-sales/).
 *
 * This data source allows you to retrieve information about flag triggers from your LaunchDarkly organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as launchdarkly from "@pulumi/launchdarkly";
 *
 * const example = launchdarkly.getFlagTrigger({
 *     id: "61d490757f7821150815518f",
 *     flagKey: "example-flag",
 *     projectKey: "the-big-project",
 *     envKey: "production",
 * });
 * ```
 */
export function getFlagTrigger(args: GetFlagTriggerArgs, opts?: pulumi.InvokeOptions): Promise<GetFlagTriggerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("launchdarkly:index/getFlagTrigger:getFlagTrigger", {
        "envKey": args.envKey,
        "flagKey": args.flagKey,
        "id": args.id,
        "projectKey": args.projectKey,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlagTrigger.
 */
export interface GetFlagTriggerArgs {
    /**
     * The unique key of the environment the flag trigger will work in.
     */
    envKey: string;
    /**
     * The unique key of the associated flag.
     */
    flagKey: string;
    id: string;
    projectKey: string;
}

/**
 * A collection of values returned by getFlagTrigger.
 */
export interface GetFlagTriggerResult {
    readonly enabled: boolean;
    /**
     * The unique key of the environment the flag trigger will work in.
     */
    readonly envKey: string;
    /**
     * The unique key of the associated flag.
     */
    readonly flagKey: string;
    readonly id: string;
    readonly instructions: outputs.GetFlagTriggerInstruction[];
    readonly integrationKey: string;
    readonly maintainerId: string;
    readonly projectKey: string;
    readonly triggerUrl: string;
}
/**
 * Provides a LaunchDarkly flag trigger data source.
 *
 * > **Note:** Flag triggers are available to customers on an Enterprise LaunchDarkly plan. To learn more, [read about our pricing](https://launchdarkly.com/pricing/). To upgrade your plan, [contact LaunchDarkly Sales](https://launchdarkly.com/contact-sales/).
 *
 * This data source allows you to retrieve information about flag triggers from your LaunchDarkly organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as launchdarkly from "@pulumi/launchdarkly";
 *
 * const example = launchdarkly.getFlagTrigger({
 *     id: "61d490757f7821150815518f",
 *     flagKey: "example-flag",
 *     projectKey: "the-big-project",
 *     envKey: "production",
 * });
 * ```
 */
export function getFlagTriggerOutput(args: GetFlagTriggerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFlagTriggerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("launchdarkly:index/getFlagTrigger:getFlagTrigger", {
        "envKey": args.envKey,
        "flagKey": args.flagKey,
        "id": args.id,
        "projectKey": args.projectKey,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlagTrigger.
 */
export interface GetFlagTriggerOutputArgs {
    /**
     * The unique key of the environment the flag trigger will work in.
     */
    envKey: pulumi.Input<string>;
    /**
     * The unique key of the associated flag.
     */
    flagKey: pulumi.Input<string>;
    id: pulumi.Input<string>;
    projectKey: pulumi.Input<string>;
}
