// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package launchdarkly

import (
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	launchdarkly "github.com/launchdarkly/terraform-provider-launchdarkly/launchdarkly"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/primait/pulumi-launchdarkly/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "launchdarkly"
	// modules:
	mainMod = "index" // the launchdarkly module
)

//go:embed cmd/pulumi-resource-launchdarkly/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {

	frameworkProvider := pfbridge.ShimProvider(launchdarkly.NewPluginProvider(version.Version)())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		// Instantiate the Terraform provider
		//
		// The [pulumi-terraform-bridge](https://github.com/pulumi/pulumi-terraform-bridge) supports 3
		// types of Terraform providers:
		//
		// 1. Providers written with the terraform-plugin-sdk/v1:
		//
		//    If the provider you are bridging is written with the terraform-plugin-sdk/v1, then you
		//    will need to adapt the boilerplate:
		//
		//    - Change the import "shimv2" to "shimv1" and change the associated import to
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1".
		//
		//    You can then proceed as normal.
		//
		// 2. Providers written with terraform-plugin-sdk/v2:
		//
		//    This boilerplate is already geared towards providers written with the
		//    terraform-plugin-sdk/v2, since it is the most common provider framework used. No
		//    adaptions are needed.
		//
		// 3. Providers written with terraform-plugin-framework:
		//
		//    If the provider you are bridging is written with the terraform-plugin-framework, then
		//    you will need to adapt the boilerplate:
		//
		//    - Remove the `shimv2` import and add:
		//
		//      	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
		//
		//    - Replace `shimv2.NewProvider` with `pfbridge.ShimProvider`.
		//
		//    - In provider/cmd/pulumi-tfgen-launchdarkly/main.go, replace the
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen" import with
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen". Remove the `version.Version`
		//      argument to `tfgen.Main`.
		//
		//    - In provider/cmd/pulumi-resource-launchdarkly/main.go, replace the
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge" import with
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge". Replace the arguments to the
		//      `tfbridge.Main` so it looks like this:
		//
		//      	tfbridge.Main(context.Background(), "launchdarkly", launchdarkly.Provider(),
		//			tfbridge.ProviderMetadata{PulumiSchema: pulumiSchema})
		//
		//   Detailed instructions can be found at
		//   https://pulumi-developer-docs.readthedocs.io/projects/pulumi-terraform-bridge/en/latest/docs/guides/new-pf-provider.html
		//   After that, you can proceed as normal.
		//
		// This is where you give the bridge a handle to the upstream terraform provider. SDKv2
		// convention is to have a function at "launchdarkly/terraform-provider-launchdarkly/provider".New
		// which takes a version and produces a factory function. The provider you are bridging may
		// not do that. You will need to find the function (generally called in upstream's main.go)
		// that produces a:
		//
		// - *"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema".Provider (for SDKv2)
		// - *"github.com/hashicorp/terraform-plugin-sdk/v1/helper/schema".Provider (for SDKv1)
		// - "github.com/hashicorp/terraform-plugin-framework/provider".Provider (for plugin-framework)
		//
		//nolint:lll
		P: frameworkProvider,

		Name:    "launchdarkly",
		Version: version.Version,
		// DisplayName is a way to be able to change the casing of the provider name when being
		// displayed on the Pulumi registry
		DisplayName: "",
		// Change this to your personal name (or a company name) that you would like to be shown in
		// the Pulumi Registry if this package is published there.
		Publisher: "Pulumi",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an PNG logo (100x100) for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g. https://github.com/org/pulumi-provider-name/releases/download/v${VERSION}/
		PluginDownloadURL: "github://api.github.com/primait/pulumi-launchdarkly",
		Description:       "A Pulumi package for creating and managing launchdarkly cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"launchdarkly", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/primait/pulumi-launchdarkly",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this should
		// match the TF provider module's require directive, not any replace directives.
		GitHubOrg:        "launchdarkly",
		UpstreamRepoPath: "./upstream",
		MetadataInfo:     tfbridge.NewProviderMetadata(metadata),
		// PreConfigureCallbackWithLogger: func(ctx context.Context, host *pulumi_provider.HostClient, vars pulumi_resource.PropertyMap, config pulumi_tf_shim.ResourceConfig) error {
		// 	host.Log(ctx, diag.Warning, "urn:pulumi:production::stub",
		// 		fmt.Sprintf("PreConfigureCallbackWithLogger called with \nhost: \n%+v \nvars:\n %+v \nconfig:\n %+v \n", host, vars, config),
		// 	)
		// 	return nil
		// },
		Config: map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			"region": {
				Type: "launchdarkly:region/region:Region",
			},
			"access_token": {
				Secret: tfbridge.True(),
				Name:   "access_token",
			},
			"oauth_token": {
				Secret: tfbridge.True(),
				Name:   "oauth_token",
			},
			"api_host": {
				Name: "api_host",
			},
			"http_timeout": {
				Name: "http_timeout",
			},
		},
		// If extra types are needed for configuration, they can be added here.
		ExtraTypes: map[string]schema.ComplexTypeSpec{
			"launchdarkly:region/region:Region": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []schema.EnumValueSpec{
					{Name: "here", Value: "HERE"},
					{Name: "overThere", Value: "OVER_THERE"},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Enable modern PyProject support in the generated Python SDK.
			PyProject: struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			// Set where the SDK is going to be published to.
			ImportBasePath: path.Join(
				"github.com/primait/pulumi-launchdarkly/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			// Opt in to all available code generation features.
			GenerateResourceContainerTypes: true,
			GenerateExtraInputTypes:        true,
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Use a wildcard import so NuGet will prefer the latest possible version.
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// MustComputeTokens maps all resources and datasources from the upstream provider into Pulumi.
	//
	// tokens.SingleModule puts every upstream item into your provider's main module.
	//
	// You shouldn't need to override anything, but if you do, use the [tfbridge.ProviderInfo.Resources]
	// and [tfbridge.ProviderInfo.DataSources].
	prov.MustComputeTokens(tokens.SingleModule("launchdarkly_", mainMod,
		tokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	// Migrate Pulumi state written by pulumi-launchdarkly <= 0.0.x
	// (terraform-provider-launchdarkly v2, terraform-plugin-sdk/v2) so it can be
	// read by this v3 (plugin-framework) based provider. See the comment on
	// bumpStateSchemaVersion for the rationale.
	//
	// Resources whose old state already has the current Pulumi shape only need a
	// schema-version bump; the remaining hooks also migrate renamed fields.
	stateMigrations := map[string]tfbridge.PreStateUpgradeHook{
		"launchdarkly_access_token":             migrateAccessTokenState,
		"launchdarkly_custom_role":              migrateCustomRoleState,
		"launchdarkly_feature_flag":             migrateFeatureFlagState,
		"launchdarkly_feature_flag_environment": bumpStateSchemaVersion,
		"launchdarkly_metric":                   migrateMetricState,
		"launchdarkly_project":                  migrateProjectState,
		"launchdarkly_segment":                  bumpStateSchemaVersion,
	}
	for tfName, hook := range stateMigrations {
		if r, ok := prov.Resources[tfName]; ok {
			r.PreStateUpgradeHook = hook
		}
	}

	return prov
}

// bumpStateSchemaVersion upgrades Pulumi state written by pulumi-launchdarkly
// <= 0.0.x (terraform-provider-launchdarkly v2, terraform-plugin-sdk/v2) so it
// can be read by the v3 (plugin-framework) based provider without invoking the
// upstream Terraform state upgraders.
//
// The upstream v0/v1 upgraders operate on the raw terraform-plugin-sdk/v2 state
// shape, in which single-nested blocks (MaxItems: 1, e.g. fallthrough) are
// encoded as single-element lists. The Pulumi bridge has always flattened those
// MaxItemsOne blocks into objects, so the state it persisted is already in the
// object shape that the v3 plugin-framework schema expects. Feeding that
// object-shaped state to the v0 upgrader fails with errors such as:
//
//	AttributeName("fallthrough"): invalid JSON, expected "[", got "{"
//
// This is sufficient only where legacy Pulumi state already matches the
// current Pulumi schema. Resources with renamed or re-keyed fields use a
// resource-specific hook below before their schema version is advanced.
func bumpStateSchemaVersion(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
	if args.PriorStateSchemaVersion >= args.ResourceSchemaVersion {
		return args.PriorStateSchemaVersion, args.PriorState, nil
	}
	return args.ResourceSchemaVersion, args.PriorState, nil
}

// migrateFeatureFlagState upgrades Pulumi state written by pulumi-launchdarkly
// <= 0.0.x for launchdarkly_feature_flag.
//
// The old Pulumi schema represented client-side availability as the plural
// clientSideAvailabilities array and customProperties as an array of objects.
// v3 uses a single clientSideAvailability object and a map keyed by property
// key. includeInSnippet was replaced by client-side availability.
func migrateFeatureFlagState(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
	return migrateLegacyState(args, func(state resource.PropertyMap) {
		moveFirstArrayElement(state, "clientSideAvailabilities", "clientSideAvailability")
		if propertyIsUnset(state, "clientSideAvailability") {
			if includeInSnippet, ok := state["includeInSnippet"]; ok {
				state["clientSideAvailability"] = resource.NewObjectProperty(resource.PropertyMap{
					"usingEnvironmentId": includeInSnippet,
					"usingMobileKey":     resource.NewBoolProperty(false),
				})
			}
		}
		delete(state, "includeInSnippet")
		convertCustomPropertiesToMap(state)
	})
}

// migrateMetricState renames randomizationUnits and removes the obsolete
// isActive field from metric state.
func migrateMetricState(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
	return migrateLegacyState(args, func(state resource.PropertyMap) {
		moveProperty(state, "randomizationUnits", "analysisUnits")
		delete(state, "isActive")
	})
}

// migrateProjectState converts the legacy environment array to a map keyed by
// environment key and applies the v2 includeInSnippet compatibility mapping.
func migrateProjectState(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
	return migrateLegacyState(args, func(state resource.PropertyMap) {
		moveFirstArrayElement(state, "defaultClientSideAvailabilities", "defaultClientSideAvailability")
		if propertyIsUnset(state, "defaultClientSideAvailability") {
			if includeInSnippet, ok := state["includeInSnippet"]; ok {
				state["defaultClientSideAvailability"] = resource.NewObjectProperty(resource.PropertyMap{
					"usingEnvironmentId": includeInSnippet,
					"usingMobileKey":     resource.NewBoolProperty(true),
				})
			}
		}
		delete(state, "includeInSnippet")
		convertArrayToMap(state, "environments", "key", convertProjectEnvironment)
	})
}

// migrateAccessTokenState moves the removed policyStatements field to
// inlineRoles and discards the obsolete expire field.
func migrateAccessTokenState(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
	return migrateLegacyState(args, func(state resource.PropertyMap) {
		moveProperty(state, "policyStatements", "inlineRoles")
		delete(state, "expire")
	})
}

// migrateCustomRoleState preserves the legacy policies value under the v3
// policyStatements field.
func migrateCustomRoleState(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
	return migrateLegacyState(args, func(state resource.PropertyMap) {
		moveProperty(state, "policies", "policyStatements")
	})
}

func migrateLegacyState(
	args tfbridge.PreStateUpgradeHookArgs,
	transform func(resource.PropertyMap),
) (int64, resource.PropertyMap, error) {
	if args.PriorStateSchemaVersion >= args.ResourceSchemaVersion {
		return args.PriorStateSchemaVersion, args.PriorState, nil
	}

	state := args.PriorState
	transform(state)
	return args.ResourceSchemaVersion, state, nil
}

// moveProperty transfers a legacy field only when the v3 field is absent,
// null, or an empty list. A populated v3 field remains authoritative.
func moveProperty(state resource.PropertyMap, from, to resource.PropertyKey) {
	if !propertyIsUnset(state, to) {
		delete(state, from)
		return
	}
	if value, exists := state[from]; exists {
		state[to] = value
		delete(state, from)
	}
}

// flattenFirstArrayElement converts a v2 single nested block to its v3 object
// representation. Empty blocks are treated as unset.
func flattenFirstArrayElement(state resource.PropertyMap, key resource.PropertyKey) {
	value, exists := state[key]
	if !exists || !value.IsArray() {
		return
	}
	entries := value.ArrayValue()
	if len(entries) == 0 {
		delete(state, key)
		return
	}
	state[key] = entries[0]
}

// moveFirstArrayElement transfers a legacy singleton array to its v3 object
// field without overwriting a populated v3 value.
func moveFirstArrayElement(state resource.PropertyMap, from, to resource.PropertyKey) {
	if !propertyIsUnset(state, to) {
		delete(state, from)
		return
	}
	value, exists := state[from]
	if !exists {
		return
	}
	if value.IsArray() {
		entries := value.ArrayValue()
		if len(entries) > 0 {
			state[to] = entries[0]
		}
	} else {
		state[to] = value
	}
	delete(state, from)
}

func propertyIsUnset(state resource.PropertyMap, key resource.PropertyKey) bool {
	value, exists := state[key]
	return !exists || value.IsNull() || (value.IsArray() && len(value.ArrayValue()) == 0)
}

// convertArrayToMap re-keys a legacy array of objects into a v3 object/map
// using an element field as the map key. The optional converter handles
// element-level shape changes.
func convertArrayToMap(
	state resource.PropertyMap,
	key, mapKey resource.PropertyKey,
	convert func(resource.PropertyMap) resource.PropertyMap,
) {
	if cp, ok := state[key]; ok && cp.IsArray() {
		entries := cp.ArrayValue()
		properties := resource.PropertyMap{}
		for _, entry := range entries {
			if !entry.IsObject() {
				continue
			}
			obj := entry.ObjectValue()
			entryKey, ok := obj[mapKey]
			if !ok || !entryKey.IsString() || entryKey.StringValue() == "" {
				continue
			}
			if convert != nil {
				obj = convert(obj)
			}
			properties[resource.PropertyKey(entryKey.StringValue())] = resource.NewObjectProperty(obj)
		}
		if len(properties) == 0 {
			delete(state, key)
			return
		}
		state[key] = resource.NewObjectProperty(properties)
	}
}

// convertCustomPropertiesToMap converts feature flag customProperties from the
// v2 array form to the v3 map form keyed by property key.
func convertCustomPropertiesToMap(state resource.PropertyMap) {
	convertArrayToMap(state, "customProperties", "key", convertCustomProperty)
}

// convertCustomProperty accepts the upstream Terraform v2 spelling (value) as
// well as the Pulumi v2 spelling (values), which is already the v3 field name.
func convertCustomProperty(obj resource.PropertyMap) resource.PropertyMap {
	if values, ok := obj["value"]; ok {
		obj["values"] = values
		delete(obj, "value")
	}
	return obj
}

// convertProjectEnvironment flattens the nested v2 approvalSettings block
// while its containing environment is re-keyed into the v3 map.
func convertProjectEnvironment(obj resource.PropertyMap) resource.PropertyMap {
	flattenFirstArrayElement(obj, "approvalSettings")
	return obj
}
