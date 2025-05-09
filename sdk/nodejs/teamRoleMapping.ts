// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as launchdarkly from "@pulumi/launchdarkly";
 *
 * const platformTeam = new launchdarkly.TeamRoleMapping("platform_team", {
 *     teamKey: "platform_team",
 *     customRoleKeys: [
 *         "platform",
 *         "nomad-administrators",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * A LaunchDarkly team/role mapping can be imported using the team key:
 *
 * ```sh
 * $ pulumi import launchdarkly:index/teamRoleMapping:TeamRoleMapping platform_team platform_team
 * ```
 */
export class TeamRoleMapping extends pulumi.CustomResource {
    /**
     * Get an existing TeamRoleMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TeamRoleMappingState, opts?: pulumi.CustomResourceOptions): TeamRoleMapping {
        return new TeamRoleMapping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'launchdarkly:index/teamRoleMapping:TeamRoleMapping';

    /**
     * Returns true if the given object is an instance of TeamRoleMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TeamRoleMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TeamRoleMapping.__pulumiType;
    }

    /**
     * List of custom role keys the team will access. The referenced custom roles must already exist in LaunchDarkly. If they don't, the provider may behave unexpectedly.
     */
    public readonly customRoleKeys!: pulumi.Output<string[]>;
    /**
     * The team key.
     */
    public readonly teamKey!: pulumi.Output<string>;

    /**
     * Create a TeamRoleMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TeamRoleMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TeamRoleMappingArgs | TeamRoleMappingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TeamRoleMappingState | undefined;
            resourceInputs["customRoleKeys"] = state ? state.customRoleKeys : undefined;
            resourceInputs["teamKey"] = state ? state.teamKey : undefined;
        } else {
            const args = argsOrState as TeamRoleMappingArgs | undefined;
            if ((!args || args.customRoleKeys === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customRoleKeys'");
            }
            if ((!args || args.teamKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamKey'");
            }
            resourceInputs["customRoleKeys"] = args ? args.customRoleKeys : undefined;
            resourceInputs["teamKey"] = args ? args.teamKey : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TeamRoleMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TeamRoleMapping resources.
 */
export interface TeamRoleMappingState {
    /**
     * List of custom role keys the team will access. The referenced custom roles must already exist in LaunchDarkly. If they don't, the provider may behave unexpectedly.
     */
    customRoleKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The team key.
     */
    teamKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TeamRoleMapping resource.
 */
export interface TeamRoleMappingArgs {
    /**
     * List of custom role keys the team will access. The referenced custom roles must already exist in LaunchDarkly. If they don't, the provider may behave unexpectedly.
     */
    customRoleKeys: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The team key.
     */
    teamKey: pulumi.Input<string>;
}
