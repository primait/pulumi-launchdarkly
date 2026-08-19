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
	// Every resource prima-pulumi creates whose upstream schema version is > 0
	// carries Terraform state upgraders that cannot consume the object-shaped
	// state persisted by the Pulumi bridge, so each needs a PreStateUpgradeHook.
	// launchdarkly_feature_flag additionally needs custom_properties converted
	// from the old list-of-blocks shape to the new map shape.
	if r, ok := prov.Resources["launchdarkly_feature_flag"]; ok {
		r.PreStateUpgradeHook = migrateFeatureFlagState
	}
	for _, tfName := range []string{
		"launchdarkly_feature_flag_environment",
		"launchdarkly_segment",
		"launchdarkly_metric",
	} {
		if r, ok := prov.Resources[tfName]; ok {
			r.PreStateUpgradeHook = bumpStateSchemaVersion
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
// Because the persisted Pulumi state already matches the current schema shape,
// we record it at the current schema version and skip the upgraders. Fields
// with a genuine type change (e.g. custom_properties on launchdarkly_feature_flag,
// which became a map) fail earlier during encoding and are handled by a
// dedicated hook instead.
func bumpStateSchemaVersion(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
	if args.PriorStateSchemaVersion >= args.ResourceSchemaVersion {
		return args.PriorStateSchemaVersion, args.PriorState, nil
	}
	return args.ResourceSchemaVersion, args.PriorState, nil
}

// migrateFeatureFlagState upgrades Pulumi state written by pulumi-launchdarkly
// <= 0.0.x for launchdarkly_feature_flag.
//
// In addition to the version bump performed by bumpStateSchemaVersion, this
// resource has one field whose Pulumi representation genuinely changed shape:
// custom_properties went from a set/list of {key, name, value} blocks to a map
// keyed by the property key, with elements {name, values, key}. Old state
// stores the list shape, which the bridge cannot encode against the new map
// schema and fails with:
//
//	objectEncoder failed on property "custom_properties": Expected an Object PropertyValue, found []
//
// We convert custom_properties to the new map shape (the old plural
// client_side_availabilities key is simply dropped and reconciled on the next
// diff) and then record the current schema version to skip the upstream
// upgraders, exactly as bumpStateSchemaVersion does.
func migrateFeatureFlagState(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
	// Only migrate state written before the current schema version.
	if args.PriorStateSchemaVersion >= args.ResourceSchemaVersion {
		return args.PriorStateSchemaVersion, args.PriorState, nil
	}

	state := args.PriorState
	if cp, ok := state["customProperties"]; ok && cp.IsArray() {
		entries := cp.ArrayValue()
		if len(entries) == 0 {
			// Empty in old state means "unset"; the new map schema treats an
			// absent value the same way, avoiding a spurious empty-map diff.
			delete(state, "customProperties")
		} else {
			state["customProperties"] = convertCustomPropertiesToMap(entries)
		}
	}

	return args.ResourceSchemaVersion, state, nil
}

// convertCustomPropertiesToMap converts the old list-of-blocks representation of
// custom_properties into the new map representation keyed by the property key.
func convertCustomPropertiesToMap(entries []resource.PropertyValue) resource.PropertyValue {
	properties := resource.PropertyMap{}
	for _, entry := range entries {
		if !entry.IsObject() {
			continue
		}
		obj := entry.ObjectValue()

		key := ""
		if k, ok := obj["key"]; ok && k.IsString() {
			key = k.StringValue()
		}

		element := resource.PropertyMap{}
		if name, ok := obj["name"]; ok {
			element["name"] = name
		}
		// The old element field "value" (list of strings) is named "values" in
		// the new schema.
		if values, ok := obj["value"]; ok {
			element["values"] = values
		} else if values, ok := obj["values"]; ok {
			element["values"] = values
		}
		if key != "" {
			element["key"] = resource.NewStringProperty(key)
		}

		properties[resource.PropertyKey(key)] = resource.NewObjectProperty(element)
	}
	return resource.NewObjectProperty(properties)
}
