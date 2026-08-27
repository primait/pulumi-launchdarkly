package launchdarkly

import (
	"testing"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func TestMigrateFeatureFlagState(t *testing.T) {
	t.Parallel()
	state := resource.PropertyMap{
		"includeInSnippet": resource.NewBoolProperty(true),
		"customProperties": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":    resource.NewStringProperty("owner"),
				"name":   resource.NewStringProperty("Owner"),
				"values": resource.NewArrayProperty([]resource.PropertyValue{resource.NewStringProperty("platform")}),
			}),
		}),
	}

	_, migrated, err := migrateFeatureFlagState(legacyStateArgs(state))
	if err != nil {
		t.Fatal(err)
	}
	if _, exists := migrated["includeInSnippet"]; exists {
		t.Fatal("includeInSnippet was not removed")
	}
	availability := migrated["clientSideAvailability"].ObjectValue()
	if !availability["usingEnvironmentId"].BoolValue() || availability["usingMobileKey"].BoolValue() {
		t.Fatalf("unexpected clientSideAvailability: %#v", availability)
	}
	properties := migrated["customProperties"].ObjectValue()
	owner := properties["owner"].ObjectValue()
	if owner["key"].StringValue() != "owner" || owner["values"].ArrayValue()[0].StringValue() != "platform" {
		t.Fatalf("unexpected customProperties: %#v", properties)
	}
}

func TestMigrateMetricState(t *testing.T) {
	t.Parallel()
	state := resource.PropertyMap{
		"isActive":           resource.NewBoolProperty(true),
		"randomizationUnits": resource.NewArrayProperty([]resource.PropertyValue{resource.NewStringProperty("user")}),
	}

	_, migrated, err := migrateMetricState(legacyStateArgs(state))
	if err != nil {
		t.Fatal(err)
	}
	if _, exists := migrated["isActive"]; exists {
		t.Fatal("isActive was not removed")
	}
	if migrated["analysisUnits"].ArrayValue()[0].StringValue() != "user" {
		t.Fatalf("unexpected analysisUnits: %#v", migrated["analysisUnits"])
	}
}

func TestMigrateProjectState(t *testing.T) {
	t.Parallel()
	state := resource.PropertyMap{
		"includeInSnippet": resource.NewBoolProperty(false),
		"environments": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key": resource.NewStringProperty("production"),
				"approvalSettings": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewObjectProperty(resource.PropertyMap{
						"required": resource.NewBoolProperty(true),
					}),
				}),
			}),
		}),
	}

	_, migrated, err := migrateProjectState(legacyStateArgs(state))
	if err != nil {
		t.Fatal(err)
	}
	availability := migrated["defaultClientSideAvailability"].ObjectValue()
	if availability["usingEnvironmentId"].BoolValue() || !availability["usingMobileKey"].BoolValue() {
		t.Fatalf("unexpected defaultClientSideAvailability: %#v", availability)
	}
	environments := migrated["environments"].ObjectValue()
	production := environments["production"].ObjectValue()
	if !production["approvalSettings"].ObjectValue()["required"].BoolValue() {
		t.Fatalf("unexpected environments: %#v", environments)
	}
}

func TestMigrateAccessTokenState(t *testing.T) {
	t.Parallel()
	tokenState := resource.PropertyMap{
		"expire": resource.NewNumberProperty(123),
		"policyStatements": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{"effect": resource.NewStringProperty("allow")}),
		}),
	}
	_, migratedToken, err := migrateAccessTokenState(legacyStateArgs(tokenState))
	if err != nil {
		t.Fatal(err)
	}
	if _, exists := migratedToken["expire"]; exists {
		t.Fatal("expire was not removed")
	}
	if migratedToken["inlineRoles"].ArrayValue()[0].ObjectValue()["effect"].StringValue() != "allow" {
		t.Fatalf("unexpected inlineRoles: %#v", migratedToken["inlineRoles"])
	}
}

func TestMigrateCustomRoleState(t *testing.T) {
	t.Parallel()
	customRoleState := resource.PropertyMap{
		"policies": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{"effect": resource.NewStringProperty("allow")}),
		}),
	}
	_, migratedCustomRole, err := migrateCustomRoleState(legacyStateArgs(customRoleState))
	if err != nil {
		t.Fatal(err)
	}
	if migratedCustomRole["policyStatements"].ArrayValue()[0].ObjectValue()["effect"].StringValue() != "allow" {
		t.Fatalf("unexpected policyStatements: %#v", migratedCustomRole["policyStatements"])
	}
}

func TestStateMigrationDoesNotChangeCurrentState(t *testing.T) {
	t.Parallel()
	state := resource.PropertyMap{"analysisUnits": resource.NewArrayProperty(nil)}
	args := tfbridge.PreStateUpgradeHookArgs{
		PriorState:              state,
		PriorStateSchemaVersion: 2,
		ResourceSchemaVersion:   2,
	}

	version, migrated, err := migrateMetricState(args)
	if err != nil {
		t.Fatal(err)
	}
	if version != 2 || len(migrated) != 1 {
		t.Fatalf("current state changed: version=%d state=%#v", version, migrated)
	}
}

func legacyStateArgs(state resource.PropertyMap) tfbridge.PreStateUpgradeHookArgs {
	return tfbridge.PreStateUpgradeHookArgs{
		PriorState:              state,
		PriorStateSchemaVersion: 0,
		ResourceSchemaVersion:   1,
	}
}
