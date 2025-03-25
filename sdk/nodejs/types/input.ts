// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

export interface AccessTokenInlineRole {
    /**
     * The list of action specifiers defining the actions to which the statement applies.
     * Either `actions` or `notActions` must be specified. For a list of available actions read [Actions reference](https://docs.launchdarkly.com/home/account-security/custom-roles/actions#actions-reference).
     */
    actions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either `allow` or `deny`. This argument defines whether the statement allows or denies access to the named resources and actions.
     */
    effect: pulumi.Input<string>;
    /**
     * The list of action specifiers defining the actions to which the statement does not apply.
     */
    notActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement does not apply.
     */
    notResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement applies.
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface AccessTokenPolicyStatement {
    /**
     * The list of action specifiers defining the actions to which the statement applies.
     * Either `actions` or `notActions` must be specified. For a list of available actions read [Actions reference](https://docs.launchdarkly.com/home/account-security/custom-roles/actions#actions-reference).
     */
    actions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either `allow` or `deny`. This argument defines whether the statement allows or denies access to the named resources and actions.
     */
    effect: pulumi.Input<string>;
    /**
     * The list of action specifiers defining the actions to which the statement does not apply.
     */
    notActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement does not apply.
     */
    notResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement applies.
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface AuditLogSubscriptionStatement {
    /**
     * The list of action specifiers defining the actions to which the statement applies.
     * Either `actions` or `notActions` must be specified. For a list of available actions read [Actions reference](https://docs.launchdarkly.com/home/account-security/custom-roles/actions#actions-reference).
     */
    actions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either `allow` or `deny`. This argument defines whether the statement allows or denies access to the named resources and actions.
     */
    effect: pulumi.Input<string>;
    /**
     * The list of action specifiers defining the actions to which the statement does not apply.
     */
    notActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement does not apply.
     */
    notResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement applies.
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface CustomRolePolicy {
    actions: pulumi.Input<pulumi.Input<string>[]>;
    effect: pulumi.Input<string>;
    resources: pulumi.Input<pulumi.Input<string>[]>;
}

export interface CustomRolePolicyStatement {
    /**
     * The list of action specifiers defining the actions to which the statement applies.
     * Either `actions` or `notActions` must be specified. For a list of available actions read [Actions reference](https://docs.launchdarkly.com/home/account-security/custom-roles/actions#actions-reference).
     */
    actions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either `allow` or `deny`. This argument defines whether the statement allows or denies access to the named resources and actions.
     */
    effect: pulumi.Input<string>;
    /**
     * The list of action specifiers defining the actions to which the statement does not apply.
     */
    notActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement does not apply.
     */
    notResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement applies.
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface EnvironmentApprovalSetting {
    /**
     * Automatically apply changes that have been approved by all reviewers. This field is only applicable for approval service kinds other than `launchdarkly`.
     */
    autoApplyApprovedChanges?: pulumi.Input<boolean>;
    /**
     * Set to `true` if changes can be applied as long as the `minNumApprovals` is met, regardless of whether any reviewers have declined a request. Defaults to `true`.
     */
    canApplyDeclinedChanges?: pulumi.Input<boolean>;
    /**
     * Set to `true` if requesters can approve or decline their own request. They may always comment. Defaults to `false`.
     */
    canReviewOwnRequest?: pulumi.Input<boolean>;
    /**
     * The number of approvals required before an approval request can be applied. This number must be between 1 and 5. Defaults to 1.
     */
    minNumApprovals?: pulumi.Input<number>;
    /**
     * Set to `true` for changes to flags in this environment to require approval. You may only set `required` to true if `requiredApprovalTags` is not set and vice versa. Defaults to `false`.
     */
    required?: pulumi.Input<boolean>;
    /**
     * An array of tags used to specify which flags with those tags require approval. You may only set `requiredApprovalTags` if `required` is set to `false` and vice versa.
     */
    requiredApprovalTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configuration for the service associated with this approval. This is specific to each approval service. For a `serviceKind` of `servicenow`, the following fields apply:
     *
     * 	 - `template` (String) The sysId of the Standard Change Request Template in ServiceNow that LaunchDarkly will use when creating the change request.
     * 	 - `detailColumn` (String) The name of the ServiceNow Change Request column LaunchDarkly uses to populate detailed approval request information. This is most commonly "justification".
     */
    serviceConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The kind of service associated with this approval. This determines which platform is used for requesting approval. Valid values are `servicenow`, `launchdarkly`. If you use a value other than `launchdarkly`, you must have already configured the integration in the LaunchDarkly UI or your apply will fail.
     */
    serviceKind?: pulumi.Input<string>;
}

export interface FeatureFlagClientSideAvailability {
    /**
     * Whether this flag is available to SDKs using the client-side ID.
     */
    usingEnvironmentId?: pulumi.Input<boolean>;
    /**
     * Whether this flag is available to SDKs using a mobile key.
     */
    usingMobileKey?: pulumi.Input<boolean>;
}

export interface FeatureFlagCustomProperty {
    /**
     * The unique custom property key.
     */
    key: pulumi.Input<string>;
    /**
     * The name of the custom property.
     */
    name: pulumi.Input<string>;
    /**
     * The list of custom property value strings.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface FeatureFlagDefaults {
    /**
     * The index of the variation the flag will default to in all new environments when off.
     */
    offVariation: pulumi.Input<number>;
    /**
     * The index of the variation the flag will default to in all new environments when on.
     */
    onVariation: pulumi.Input<number>;
}

export interface FeatureFlagEnvironmentContextTarget {
    /**
     * The context kind on which the flag should target in this environment. User (`user`) targets should be specified as `targets` attribute blocks.
     */
    contextKind: pulumi.Input<string>;
    /**
     * List of `user` strings to target.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The index of the variation to serve if a user target value is matched.
     */
    variation: pulumi.Input<number>;
}

export interface FeatureFlagEnvironmentFallthrough {
    /**
     * Group percentage rollout by a custom attribute. This argument is only valid if rolloutWeights is also specified.
     */
    bucketBy?: pulumi.Input<string>;
    /**
     * The context kind associated with the specified rollout. This argument is only valid if rolloutWeights is also specified. If omitted, defaults to `user`.
     */
    contextKind?: pulumi.Input<string>;
    /**
     * List of integer percentage rollout weights (in thousandths of a percent) to apply to each variation if the rule clauses evaluates to `true`. The sum of the `rolloutWeights` must equal 100000 and the number of rollout weights specified in the array must match the number of flag variations. You must specify either `variation` or `rolloutWeights`.
     */
    rolloutWeights?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The default integer variation index to serve if no `prerequisites`, `target`, or `rules` apply. You must specify either `variation` or `rolloutWeights`.
     */
    variation?: pulumi.Input<number>;
}

export interface FeatureFlagEnvironmentPrerequisite {
    /**
     * The prerequisite feature flag's `key`.
     */
    flagKey: pulumi.Input<string>;
    /**
     * The index of the prerequisite feature flag's variation to target.
     */
    variation: pulumi.Input<number>;
}

export interface FeatureFlagEnvironmentRule {
    /**
     * Group percentage rollout by a custom attribute. This argument is only valid if `rolloutWeights` is also specified.
     */
    bucketBy?: pulumi.Input<string>;
    /**
     * List of nested blocks specifying the logical clauses to evaluate
     */
    clauses?: pulumi.Input<pulumi.Input<inputs.FeatureFlagEnvironmentRuleClause>[]>;
    /**
     * The context kind associated with the specified rollout. This argument is only valid if `rolloutWeights` is also specified. Defaults to `user` if omitted.
     */
    contextKind?: pulumi.Input<string>;
    /**
     * A human-readable description of the targeting rule.
     */
    description?: pulumi.Input<string>;
    /**
     * List of integer percentage rollout weights (in thousandths of a percent) to apply to each variation if the rule clauses evaluates to `true`. The sum of the `rolloutWeights` must equal 100000 and the number of rollout weights specified in the array must match the number of flag variations. You must specify either `variation` or `rolloutWeights`.
     */
    rolloutWeights?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The integer variation index to serve if the rule clauses evaluate to `true`. You must specify either `variation` or `rolloutWeights`
     */
    variation?: pulumi.Input<number>;
}

export interface FeatureFlagEnvironmentRuleClause {
    /**
     * The user attribute to operate on
     */
    attribute: pulumi.Input<string>;
    /**
     * The context kind associated with this rule clause. If omitted, defaults to `user`.
     */
    contextKind?: pulumi.Input<string>;
    /**
     * Whether to negate the rule clause.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * The operator associated with the rule clause. Available options are `in`, `endsWith`, `startsWith`, `matches`, `contains`, `lessThan`, `lessThanOrEqual`, `greaterThanOrEqual`, `before`, `after`, `segmentMatch`, `semVerEqual`, `semVerLessThan`, and `semVerGreaterThan`. Read LaunchDarkly's [Operators](https://docs.launchdarkly.com/sdk/concepts/flag-evaluation-rules#operators) documentation for more information.
     */
    op: pulumi.Input<string>;
    /**
     * The type for each of the clause's values. Available types are `boolean`, `string`, and `number`. If omitted, `valueType` defaults to `string`.
     */
    valueType?: pulumi.Input<string>;
    /**
     * The list of values associated with the rule clause.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface FeatureFlagEnvironmentTarget {
    /**
     * List of `user` strings to target.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The index of the variation to serve if a user target value is matched.
     */
    variation: pulumi.Input<number>;
}

export interface FeatureFlagVariation {
    /**
     * The feature flag's description.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name of the feature flag.
     */
    name?: pulumi.Input<string>;
    /**
     * The variation value. The value's type must correspond to the `variationType` argument. For example: `variationType = "boolean"` accepts only `true` or `false`. The `number` variation type accepts both floats and ints, but please note that any trailing zeroes on floats will be trimmed (i.e. `1.1` and `1.100` will both be converted to `1.1`).
     */
    value: pulumi.Input<string>;
}

export interface FlagTriggerInstructions {
    /**
     * The action to perform when triggering. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`.
     */
    kind: pulumi.Input<string>;
}

export interface GetTeamMemberRoleAttribute {
    /**
     * The key / name of your role attribute. In the example `$${roleAttribute/testAttribute}`, the key is `testAttribute`.
     */
    key: string;
    /**
     * A list of values for your role attribute. For example, if your policy statement defines the resource `"proj/$${roleAttribute/testAttribute}"`, the values would be the keys of the projects you wanted to assign access to.
     */
    values: string[];
}

export interface GetTeamMemberRoleAttributeArgs {
    /**
     * The key / name of your role attribute. In the example `$${roleAttribute/testAttribute}`, the key is `testAttribute`.
     */
    key: pulumi.Input<string>;
    /**
     * A list of values for your role attribute. For example, if your policy statement defines the resource `"proj/$${roleAttribute/testAttribute}"`, the values would be the keys of the projects you wanted to assign access to.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetTeamRoleAttribute {
    /**
     * The key / name of your role attribute. In the example `$${roleAttribute/testAttribute}`, the key is `testAttribute`.
     */
    key: string;
    /**
     * A list of values for your role attribute. For example, if your policy statement defines the resource `"proj/$${roleAttribute/testAttribute}"`, the values would be the keys of the projects you wanted to assign access to.
     */
    values: string[];
}

export interface GetTeamRoleAttributeArgs {
    /**
     * The key / name of your role attribute. In the example `$${roleAttribute/testAttribute}`, the key is `testAttribute`.
     */
    key: pulumi.Input<string>;
    /**
     * A list of values for your role attribute. For example, if your policy statement defines the resource `"proj/$${roleAttribute/testAttribute}"`, the values would be the keys of the projects you wanted to assign access to.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface MetricUrl {
    /**
     * The URL type. Available choices are `exact`, `canonical`, `substring` and `regex`.
     */
    kind: pulumi.Input<string>;
    /**
     * (Required for kind `regex`) The regex pattern to match by.
     */
    pattern?: pulumi.Input<string>;
    /**
     * (Required for kind `substring`) The URL substring to match by.
     */
    substring?: pulumi.Input<string>;
    /**
     * (Required for kind `exact` and `canonical`) The exact or canonical URL.
     */
    url?: pulumi.Input<string>;
}

export interface ProjectDefaultClientSideAvailability {
    usingEnvironmentId: pulumi.Input<boolean>;
    usingMobileKey: pulumi.Input<boolean>;
}

export interface ProjectEnvironment {
    /**
     * The environment's SDK key.
     */
    apiKey?: pulumi.Input<string>;
    approvalSettings?: pulumi.Input<pulumi.Input<inputs.ProjectEnvironmentApprovalSetting>[]>;
    /**
     * The environment's client-side ID.
     */
    clientSideId?: pulumi.Input<string>;
    /**
     * The color swatch as an RGB hex value with no leading `#`. For example: `000000`
     */
    color: pulumi.Input<string>;
    /**
     * Set to `true` if this environment requires confirmation for flag and segment changes. This field will default to `false` when not set.
     */
    confirmChanges?: pulumi.Input<boolean>;
    /**
     * Denotes whether the environment is critical.
     */
    critical?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable data export for every flag created in this environment after you configure this argument. This field will default to `false` when not set. To learn more, read [Data Export](https://docs.launchdarkly.com/home/data-export).
     */
    defaultTrackEvents?: pulumi.Input<boolean>;
    /**
     * The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
     */
    defaultTtl?: pulumi.Input<number>;
    /**
     * The project-unique key for the environment. A change in this field will force the destruction of the existing resource and the creation of a new one.
     */
    key: pulumi.Input<string>;
    /**
     * The environment's mobile key.
     */
    mobileKey?: pulumi.Input<string>;
    /**
     * The name of the environment.
     */
    name: pulumi.Input<string>;
    /**
     * Set to `true` if this environment requires comments for flag and segment changes. This field will default to `false` when not set.
     */
    requireComments?: pulumi.Input<boolean>;
    /**
     * Set to `true` to ensure a user of the client-side SDK cannot impersonate another user. This field will default to `false` when not set.
     */
    secureMode?: pulumi.Input<boolean>;
    /**
     * Tags associated with your resource.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ProjectEnvironmentApprovalSetting {
    /**
     * Automatically apply changes that have been approved by all reviewers. This field is only applicable for approval service kinds other than `launchdarkly`.
     */
    autoApplyApprovedChanges?: pulumi.Input<boolean>;
    /**
     * Set to `true` if changes can be applied as long as the `minNumApprovals` is met, regardless of whether any reviewers have declined a request. Defaults to `true`.
     */
    canApplyDeclinedChanges?: pulumi.Input<boolean>;
    /**
     * Set to `true` if requesters can approve or decline their own request. They may always comment. Defaults to `false`.
     */
    canReviewOwnRequest?: pulumi.Input<boolean>;
    /**
     * The number of approvals required before an approval request can be applied. This number must be between 1 and 5. Defaults to 1.
     */
    minNumApprovals?: pulumi.Input<number>;
    /**
     * Set to `true` for changes to flags in this environment to require approval. You may only set `required` to true if `requiredApprovalTags` is not set and vice versa. Defaults to `false`.
     */
    required?: pulumi.Input<boolean>;
    /**
     * An array of tags used to specify which flags with those tags require approval. You may only set `requiredApprovalTags` if `required` is set to `false` and vice versa.
     */
    requiredApprovalTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configuration for the service associated with this approval. This is specific to each approval service. For a `serviceKind` of `servicenow`, the following fields apply:
     *
     * 	 - `template` (String) The sysId of the Standard Change Request Template in ServiceNow that LaunchDarkly will use when creating the change request.
     * 	 - `detailColumn` (String) The name of the ServiceNow Change Request column LaunchDarkly uses to populate detailed approval request information. This is most commonly "justification".
     */
    serviceConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The kind of service associated with this approval. This determines which platform is used for requesting approval. Valid values are `servicenow`, `launchdarkly`. If you use a value other than `launchdarkly`, you must have already configured the integration in the LaunchDarkly UI or your apply will fail.
     */
    serviceKind?: pulumi.Input<string>;
}

export interface RelayProxyConfigurationPolicy {
    /**
     * The list of action specifiers defining the actions to which the statement applies.
     * Either `actions` or `notActions` must be specified. For a list of available actions read [Actions reference](https://docs.launchdarkly.com/home/account-security/custom-roles/actions#actions-reference).
     */
    actions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either `allow` or `deny`. This argument defines whether the statement allows or denies access to the named resources and actions.
     */
    effect: pulumi.Input<string>;
    /**
     * The list of action specifiers defining the actions to which the statement does not apply.
     */
    notActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement does not apply.
     */
    notResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement applies.
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface SegmentExcludedContext {
    /**
     * The context kind associated with this segment target. To target on user contexts, use the included and excluded attributes.
     */
    contextKind: pulumi.Input<string>;
    /**
     * List of target object keys included in or excluded from the segment.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface SegmentIncludedContext {
    /**
     * The context kind associated with this segment target. To target on user contexts, use the included and excluded attributes.
     */
    contextKind: pulumi.Input<string>;
    /**
     * List of target object keys included in or excluded from the segment.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface SegmentRule {
    /**
     * The attribute by which to group contexts together.
     */
    bucketBy?: pulumi.Input<string>;
    /**
     * List of nested blocks specifying the logical clauses to evaluate
     */
    clauses?: pulumi.Input<pulumi.Input<inputs.SegmentRuleClause>[]>;
    /**
     * The context kind associated with this segment rule. This argument is only valid if `weight` is also specified. If omitted, defaults to `user`.
     */
    rolloutContextKind?: pulumi.Input<string>;
    /**
     * The integer weight of the rule (between 1 and 100000).
     */
    weight?: pulumi.Input<number>;
}

export interface SegmentRuleClause {
    /**
     * The user attribute to operate on
     */
    attribute: pulumi.Input<string>;
    /**
     * The context kind associated with this rule clause. If omitted, defaults to `user`.
     */
    contextKind?: pulumi.Input<string>;
    /**
     * Whether to negate the rule clause.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * The operator associated with the rule clause. Available options are `in`, `endsWith`, `startsWith`, `matches`, `contains`, `lessThan`, `lessThanOrEqual`, `greaterThanOrEqual`, `before`, `after`, `segmentMatch`, `semVerEqual`, `semVerLessThan`, and `semVerGreaterThan`. Read LaunchDarkly's [Operators](https://docs.launchdarkly.com/sdk/concepts/flag-evaluation-rules#operators) documentation for more information.
     */
    op: pulumi.Input<string>;
    /**
     * The type for each of the clause's values. Available types are `boolean`, `string`, and `number`. If omitted, `valueType` defaults to `string`.
     */
    valueType?: pulumi.Input<string>;
    /**
     * The list of values associated with the rule clause.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface TeamMemberRoleAttribute {
    /**
     * The key / name of your role attribute. In the example `$${roleAttribute/testAttribute}`, the key is `testAttribute`.
     */
    key: pulumi.Input<string>;
    /**
     * A list of values for your role attribute. For example, if your policy statement defines the resource `"proj/$${roleAttribute/testAttribute}"`, the values would be the keys of the projects you wanted to assign access to.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface TeamRoleAttribute {
    /**
     * The key / name of your role attribute. In the example `$${roleAttribute/testAttribute}`, the key is `testAttribute`.
     */
    key: pulumi.Input<string>;
    /**
     * A list of values for your role attribute. For example, if your policy statement defines the resource `"proj/$${roleAttribute/testAttribute}"`, the values would be the keys of the projects you wanted to assign access to.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface WebhookStatement {
    /**
     * The list of action specifiers defining the actions to which the statement applies.
     * Either `actions` or `notActions` must be specified. For a list of available actions read [Actions reference](https://docs.launchdarkly.com/home/account-security/custom-roles/actions#actions-reference).
     */
    actions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either `allow` or `deny`. This argument defines whether the statement allows or denies access to the named resources and actions.
     */
    effect: pulumi.Input<string>;
    /**
     * The list of action specifiers defining the actions to which the statement does not apply.
     */
    notActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement does not apply.
     */
    notResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of resource specifiers defining the resources to which the statement applies.
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
}
