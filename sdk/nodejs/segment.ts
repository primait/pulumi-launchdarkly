// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a LaunchDarkly segment resource.
 *
 * This resource allows you to create and manage segments within your LaunchDarkly organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as launchdarkly from "@pulumi/launchdarkly";
 *
 * // This example shows the use of tags, targets, context targets, and rules for a segment
 * const example = new launchdarkly.Segment("example", {
 *     key: "example-segment-key",
 *     projectKey: exampleLaunchdarklyProject.key,
 *     envKey: exampleLaunchdarklyEnvironment.key,
 *     name: "example segment",
 *     description: "This segment is managed by Terraform",
 *     tags: [
 *         "segment-tag-1",
 *         "segment-tag-2",
 *     ],
 *     includeds: [
 *         "user1",
 *         "user2",
 *     ],
 *     excludeds: [
 *         "user3",
 *         "user4",
 *     ],
 *     includedContexts: [{
 *         values: [
 *             "account1",
 *             "account2",
 *         ],
 *         contextKind: "account",
 *     }],
 *     rules: [{
 *         clauses: [{
 *             attribute: "country",
 *             op: "startsWith",
 *             values: [
 *                 "en",
 *                 "de",
 *                 "un",
 *             ],
 *             negate: false,
 *             contextKind: "location-data",
 *         }],
 *     }],
 * });
 * // This example shows a segment configured to have an unbounded number of individual targets
 * const big_example = new launchdarkly.Segment("big-example", {
 *     key: "example-big-segment-key",
 *     projectKey: exampleLaunchdarklyProject.key,
 *     envKey: exampleLaunchdarklyEnvironment.key,
 *     name: "example big segment",
 *     description: "This big segment is managed by Terraform",
 *     tags: [
 *         "segment-tag-1",
 *         "segment-tag-2",
 *     ],
 *     unbounded: true,
 *     unboundedContextKind: "user",
 * });
 * // This example shows a segment with a targeting rule that uses all clause operators
 * const segmentWithAllClauseOperators = new launchdarkly.Segment("segment_with_all_clause_operators", {
 *     name: "Segment with all clause operators",
 *     key: "segment-operators",
 *     projectKey: "projectx",
 *     envKey: "development",
 *     rules: [{
 *         clauses: [
 *             {
 *                 attribute: "username",
 *                 op: "in",
 *                 values: [
 *                     "henrietta powell",
 *                     "wally waterbear",
 *                 ],
 *             },
 *             {
 *                 attribute: "username",
 *                 op: "endsWith",
 *                 values: [
 *                     "powell",
 *                     "waterbear",
 *                 ],
 *             },
 *             {
 *                 attribute: "username",
 *                 op: "startsWith",
 *                 values: [
 *                     "henrietta",
 *                     "wally",
 *                 ],
 *             },
 *             {
 *                 attribute: "username",
 *                 op: "matches",
 *                 values: ["henr*"],
 *             },
 *             {
 *                 attribute: "username",
 *                 op: "contains",
 *                 values: ["water"],
 *             },
 *             {
 *                 attribute: "pageVisits",
 *                 op: "lessThan",
 *                 values: ["100"],
 *             },
 *             {
 *                 attribute: "pageVisits",
 *                 op: "lessThanOrEqual",
 *                 values: ["100"],
 *             },
 *             {
 *                 attribute: "pageVisits",
 *                 op: "greaterThan",
 *                 values: ["100"],
 *             },
 *             {
 *                 attribute: "pageVisits",
 *                 op: "greaterThanOrEqual",
 *                 values: ["100"],
 *             },
 *             {
 *                 attribute: "creationDate",
 *                 op: "before",
 *                 values: ["2024-05-03T15:57:30Z"],
 *             },
 *             {
 *                 attribute: "creationDate",
 *                 op: "after",
 *                 values: ["2024-05-03T15:57:30Z"],
 *             },
 *             {
 *                 attribute: "version",
 *                 op: "semVerEqual",
 *                 values: [
 *                     "1.0.0",
 *                     "1.0.1",
 *                 ],
 *                 contextKind: "application",
 *             },
 *             {
 *                 attribute: "version",
 *                 op: "semVerLessThan",
 *                 values: ["1.0.0"],
 *                 contextKind: "application",
 *             },
 *             {
 *                 attribute: "version",
 *                 op: "semVerGreaterThan",
 *                 values: ["1.0.0"],
 *                 contextKind: "application",
 *             },
 *             {
 *                 attribute: "context",
 *                 op: "segmentMatch",
 *                 values: ["test-segment"],
 *             },
 *         ],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * #LaunchDarkly segments can be imported using the segment's ID in the form `project_key/env_key/segment_key`
 *
 * ```sh
 * $ pulumi import launchdarkly:index/segment:Segment example example-project/example-environment/example-segment-key
 * ```
 */
export class Segment extends pulumi.CustomResource {
    /**
     * Get an existing Segment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SegmentState, opts?: pulumi.CustomResourceOptions): Segment {
        return new Segment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'launchdarkly:index/segment:Segment';

    /**
     * Returns true if the given object is an instance of Segment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Segment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Segment.__pulumiType;
    }

    /**
     * The segment's creation date represented as a UNIX epoch timestamp.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<number>;
    /**
     * The description of the segment's purpose.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The segment's environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    public readonly envKey!: pulumi.Output<string>;
    /**
     * List of non-user target objects excluded from the segment. This attribute is not valid when `unbounded` is set to `true`.
     */
    public readonly excludedContexts!: pulumi.Output<outputs.SegmentExcludedContext[] | undefined>;
    /**
     * List of user keys excluded from the segment. To target on other context kinds, use the excludedContexts block attribute. This attribute is not valid when `unbounded` is set to `true`.
     */
    public readonly excludeds!: pulumi.Output<string[] | undefined>;
    /**
     * List of non-user target objects included in the segment. This attribute is not valid when `unbounded` is set to `true`.
     */
    public readonly includedContexts!: pulumi.Output<outputs.SegmentIncludedContext[] | undefined>;
    /**
     * List of user keys included in the segment. To target on other context kinds, use the includedContexts block attribute. This attribute is not valid when `unbounded` is set to `true`.
     */
    public readonly includeds!: pulumi.Output<string[] | undefined>;
    /**
     * The unique key that references the segment. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The human-friendly name for the segment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The segment's project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    public readonly projectKey!: pulumi.Output<string>;
    /**
     * List of nested custom rule blocks to apply to the segment. This attribute is not valid when `unbounded` is set to `true`.
     */
    public readonly rules!: pulumi.Output<outputs.SegmentRule[] | undefined>;
    /**
     * Tags associated with your resource.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly unbounded!: pulumi.Output<boolean | undefined>;
    /**
     * For Big Segments, the targeted context kind. If this attribute is not specified it will default to `user`. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    public readonly unboundedContextKind!: pulumi.Output<string>;

    /**
     * Create a Segment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SegmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SegmentArgs | SegmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SegmentState | undefined;
            resourceInputs["creationDate"] = state ? state.creationDate : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["envKey"] = state ? state.envKey : undefined;
            resourceInputs["excludedContexts"] = state ? state.excludedContexts : undefined;
            resourceInputs["excludeds"] = state ? state.excludeds : undefined;
            resourceInputs["includedContexts"] = state ? state.includedContexts : undefined;
            resourceInputs["includeds"] = state ? state.includeds : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["unbounded"] = state ? state.unbounded : undefined;
            resourceInputs["unboundedContextKind"] = state ? state.unboundedContextKind : undefined;
        } else {
            const args = argsOrState as SegmentArgs | undefined;
            if ((!args || args.envKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'envKey'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.projectKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectKey'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["envKey"] = args ? args.envKey : undefined;
            resourceInputs["excludedContexts"] = args ? args.excludedContexts : undefined;
            resourceInputs["excludeds"] = args ? args.excludeds : undefined;
            resourceInputs["includedContexts"] = args ? args.includedContexts : undefined;
            resourceInputs["includeds"] = args ? args.includeds : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["unbounded"] = args ? args.unbounded : undefined;
            resourceInputs["unboundedContextKind"] = args ? args.unboundedContextKind : undefined;
            resourceInputs["creationDate"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Segment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Segment resources.
 */
export interface SegmentState {
    /**
     * The segment's creation date represented as a UNIX epoch timestamp.
     */
    creationDate?: pulumi.Input<number>;
    /**
     * The description of the segment's purpose.
     */
    description?: pulumi.Input<string>;
    /**
     * The segment's environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    envKey?: pulumi.Input<string>;
    /**
     * List of non-user target objects excluded from the segment. This attribute is not valid when `unbounded` is set to `true`.
     */
    excludedContexts?: pulumi.Input<pulumi.Input<inputs.SegmentExcludedContext>[]>;
    /**
     * List of user keys excluded from the segment. To target on other context kinds, use the excludedContexts block attribute. This attribute is not valid when `unbounded` is set to `true`.
     */
    excludeds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of non-user target objects included in the segment. This attribute is not valid when `unbounded` is set to `true`.
     */
    includedContexts?: pulumi.Input<pulumi.Input<inputs.SegmentIncludedContext>[]>;
    /**
     * List of user keys included in the segment. To target on other context kinds, use the includedContexts block attribute. This attribute is not valid when `unbounded` is set to `true`.
     */
    includeds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique key that references the segment. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    key?: pulumi.Input<string>;
    /**
     * The human-friendly name for the segment.
     */
    name?: pulumi.Input<string>;
    /**
     * The segment's project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * List of nested custom rule blocks to apply to the segment. This attribute is not valid when `unbounded` is set to `true`.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.SegmentRule>[]>;
    /**
     * Tags associated with your resource.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    unbounded?: pulumi.Input<boolean>;
    /**
     * For Big Segments, the targeted context kind. If this attribute is not specified it will default to `user`. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    unboundedContextKind?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Segment resource.
 */
export interface SegmentArgs {
    /**
     * The description of the segment's purpose.
     */
    description?: pulumi.Input<string>;
    /**
     * The segment's environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    envKey: pulumi.Input<string>;
    /**
     * List of non-user target objects excluded from the segment. This attribute is not valid when `unbounded` is set to `true`.
     */
    excludedContexts?: pulumi.Input<pulumi.Input<inputs.SegmentExcludedContext>[]>;
    /**
     * List of user keys excluded from the segment. To target on other context kinds, use the excludedContexts block attribute. This attribute is not valid when `unbounded` is set to `true`.
     */
    excludeds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of non-user target objects included in the segment. This attribute is not valid when `unbounded` is set to `true`.
     */
    includedContexts?: pulumi.Input<pulumi.Input<inputs.SegmentIncludedContext>[]>;
    /**
     * List of user keys included in the segment. To target on other context kinds, use the includedContexts block attribute. This attribute is not valid when `unbounded` is set to `true`.
     */
    includeds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique key that references the segment. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    key: pulumi.Input<string>;
    /**
     * The human-friendly name for the segment.
     */
    name?: pulumi.Input<string>;
    /**
     * The segment's project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    projectKey: pulumi.Input<string>;
    /**
     * List of nested custom rule blocks to apply to the segment. This attribute is not valid when `unbounded` is set to `true`.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.SegmentRule>[]>;
    /**
     * Tags associated with your resource.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    unbounded?: pulumi.Input<boolean>;
    /**
     * For Big Segments, the targeted context kind. If this attribute is not specified it will default to `user`. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    unboundedContextKind?: pulumi.Input<string>;
}
