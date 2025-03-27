# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AuditLogSubscriptionArgs', 'AuditLogSubscription']

@pulumi.input_type
class AuditLogSubscriptionArgs:
    def __init__(__self__, *,
                 config: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 integration_key: pulumi.Input[str],
                 on: pulumi.Input[bool],
                 statements: pulumi.Input[Sequence[pulumi.Input['AuditLogSubscriptionStatementArgs']]],
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a AuditLogSubscription resource.
        :param pulumi.Input[str] integration_key: The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        :param pulumi.Input[bool] on: Whether or not you want your subscription enabled, i.e. to actively send events.
        :param pulumi.Input[Sequence[pulumi.Input['AuditLogSubscriptionStatementArgs']]] statements: A block representing the resources to which you wish to subscribe.
        :param pulumi.Input[str] name: A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags associated with your resource.
        """
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "integration_key", integration_key)
        pulumi.set(__self__, "on", on)
        pulumi.set(__self__, "statements", statements)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="integrationKey")
    def integration_key(self) -> pulumi.Input[str]:
        """
        The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        """
        return pulumi.get(self, "integration_key")

    @integration_key.setter
    def integration_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "integration_key", value)

    @property
    @pulumi.getter
    def on(self) -> pulumi.Input[bool]:
        """
        Whether or not you want your subscription enabled, i.e. to actively send events.
        """
        return pulumi.get(self, "on")

    @on.setter
    def on(self, value: pulumi.Input[bool]):
        pulumi.set(self, "on", value)

    @property
    @pulumi.getter
    def statements(self) -> pulumi.Input[Sequence[pulumi.Input['AuditLogSubscriptionStatementArgs']]]:
        """
        A block representing the resources to which you wish to subscribe.
        """
        return pulumi.get(self, "statements")

    @statements.setter
    def statements(self, value: pulumi.Input[Sequence[pulumi.Input['AuditLogSubscriptionStatementArgs']]]):
        pulumi.set(self, "statements", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Tags associated with your resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _AuditLogSubscriptionState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 integration_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 on: Optional[pulumi.Input[bool]] = None,
                 statements: Optional[pulumi.Input[Sequence[pulumi.Input['AuditLogSubscriptionStatementArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering AuditLogSubscription resources.
        :param pulumi.Input[str] integration_key: The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        :param pulumi.Input[str] name: A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
        :param pulumi.Input[bool] on: Whether or not you want your subscription enabled, i.e. to actively send events.
        :param pulumi.Input[Sequence[pulumi.Input['AuditLogSubscriptionStatementArgs']]] statements: A block representing the resources to which you wish to subscribe.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags associated with your resource.
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if integration_key is not None:
            pulumi.set(__self__, "integration_key", integration_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if on is not None:
            pulumi.set(__self__, "on", on)
        if statements is not None:
            pulumi.set(__self__, "statements", statements)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="integrationKey")
    def integration_key(self) -> Optional[pulumi.Input[str]]:
        """
        The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        """
        return pulumi.get(self, "integration_key")

    @integration_key.setter
    def integration_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def on(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not you want your subscription enabled, i.e. to actively send events.
        """
        return pulumi.get(self, "on")

    @on.setter
    def on(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "on", value)

    @property
    @pulumi.getter
    def statements(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AuditLogSubscriptionStatementArgs']]]]:
        """
        A block representing the resources to which you wish to subscribe.
        """
        return pulumi.get(self, "statements")

    @statements.setter
    def statements(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AuditLogSubscriptionStatementArgs']]]]):
        pulumi.set(self, "statements", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Tags associated with your resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class AuditLogSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 integration_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 on: Optional[pulumi.Input[bool]] = None,
                 statements: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AuditLogSubscriptionStatementArgs', 'AuditLogSubscriptionStatementArgsDict']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a LaunchDarkly audit log subscription resource.

        This resource allows you to create and manage LaunchDarkly audit log subscriptions.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] integration_key: The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        :param pulumi.Input[str] name: A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
        :param pulumi.Input[bool] on: Whether or not you want your subscription enabled, i.e. to actively send events.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AuditLogSubscriptionStatementArgs', 'AuditLogSubscriptionStatementArgsDict']]]] statements: A block representing the resources to which you wish to subscribe.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags associated with your resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuditLogSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a LaunchDarkly audit log subscription resource.

        This resource allows you to create and manage LaunchDarkly audit log subscriptions.

        :param str resource_name: The name of the resource.
        :param AuditLogSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuditLogSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 integration_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 on: Optional[pulumi.Input[bool]] = None,
                 statements: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AuditLogSubscriptionStatementArgs', 'AuditLogSubscriptionStatementArgsDict']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuditLogSubscriptionArgs.__new__(AuditLogSubscriptionArgs)

            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            if integration_key is None and not opts.urn:
                raise TypeError("Missing required property 'integration_key'")
            __props__.__dict__["integration_key"] = integration_key
            __props__.__dict__["name"] = name
            if on is None and not opts.urn:
                raise TypeError("Missing required property 'on'")
            __props__.__dict__["on"] = on
            if statements is None and not opts.urn:
                raise TypeError("Missing required property 'statements'")
            __props__.__dict__["statements"] = statements
            __props__.__dict__["tags"] = tags
        super(AuditLogSubscription, __self__).__init__(
            'launchdarkly:index/auditLogSubscription:AuditLogSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            integration_key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            on: Optional[pulumi.Input[bool]] = None,
            statements: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AuditLogSubscriptionStatementArgs', 'AuditLogSubscriptionStatementArgsDict']]]]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'AuditLogSubscription':
        """
        Get an existing AuditLogSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] integration_key: The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        :param pulumi.Input[str] name: A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
        :param pulumi.Input[bool] on: Whether or not you want your subscription enabled, i.e. to actively send events.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AuditLogSubscriptionStatementArgs', 'AuditLogSubscriptionStatementArgsDict']]]] statements: A block representing the resources to which you wish to subscribe.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags associated with your resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuditLogSubscriptionState.__new__(_AuditLogSubscriptionState)

        __props__.__dict__["config"] = config
        __props__.__dict__["integration_key"] = integration_key
        __props__.__dict__["name"] = name
        __props__.__dict__["on"] = on
        __props__.__dict__["statements"] = statements
        __props__.__dict__["tags"] = tags
        return AuditLogSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="integrationKey")
    def integration_key(self) -> pulumi.Output[str]:
        """
        The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
        """
        return pulumi.get(self, "integration_key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def on(self) -> pulumi.Output[bool]:
        """
        Whether or not you want your subscription enabled, i.e. to actively send events.
        """
        return pulumi.get(self, "on")

    @property
    @pulumi.getter
    def statements(self) -> pulumi.Output[Sequence['outputs.AuditLogSubscriptionStatement']]:
        """
        A block representing the resources to which you wish to subscribe.
        """
        return pulumi.get(self, "statements")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Tags associated with your resource.
        """
        return pulumi.get(self, "tags")

