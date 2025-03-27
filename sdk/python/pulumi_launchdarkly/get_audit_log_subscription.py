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

__all__ = [
    'GetAuditLogSubscriptionResult',
    'AwaitableGetAuditLogSubscriptionResult',
    'get_audit_log_subscription',
    'get_audit_log_subscription_output',
]

@pulumi.output_type
class GetAuditLogSubscriptionResult:
    """
    A collection of values returned by getAuditLogSubscription.
    """
    def __init__(__self__, config=None, id=None, integration_key=None, name=None, on=None, statements=None, tags=None):
        if config and not isinstance(config, dict):
            raise TypeError("Expected argument 'config' to be a dict")
        pulumi.set(__self__, "config", config)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if integration_key and not isinstance(integration_key, str):
            raise TypeError("Expected argument 'integration_key' to be a str")
        pulumi.set(__self__, "integration_key", integration_key)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if on and not isinstance(on, bool):
            raise TypeError("Expected argument 'on' to be a bool")
        pulumi.set(__self__, "on", on)
        if statements and not isinstance(statements, list):
            raise TypeError("Expected argument 'statements' to be a list")
        pulumi.set(__self__, "statements", statements)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def config(self) -> Mapping[str, str]:
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The audit log subscription ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="integrationKey")
    def integration_key(self) -> str:
        """
        The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`.
        """
        return pulumi.get(self, "integration_key")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def on(self) -> bool:
        """
        Whether or not you want your subscription enabled, i.e. to actively send events.
        """
        return pulumi.get(self, "on")

    @property
    @pulumi.getter
    def statements(self) -> Sequence['outputs.GetAuditLogSubscriptionStatementResult']:
        """
        A block representing the resources to which you wish to subscribe.
        """
        return pulumi.get(self, "statements")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        Tags associated with your resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetAuditLogSubscriptionResult(GetAuditLogSubscriptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuditLogSubscriptionResult(
            config=self.config,
            id=self.id,
            integration_key=self.integration_key,
            name=self.name,
            on=self.on,
            statements=self.statements,
            tags=self.tags)


def get_audit_log_subscription(id: Optional[str] = None,
                               integration_key: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuditLogSubscriptionResult:
    """
    Provides a LaunchDarkly audit log subscription data source.

    This data source allows you to retrieve information about LaunchDarkly audit log subscriptions.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_launchdarkly as launchdarkly

    test = launchdarkly.get_audit_log_subscription(id="5f0cd446a77cba0b4c5644a7",
        integration_key="msteams")
    ```


    :param str id: The audit log subscription ID.
    :param str integration_key: The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['integrationKey'] = integration_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('launchdarkly:index/getAuditLogSubscription:getAuditLogSubscription', __args__, opts=opts, typ=GetAuditLogSubscriptionResult).value

    return AwaitableGetAuditLogSubscriptionResult(
        config=pulumi.get(__ret__, 'config'),
        id=pulumi.get(__ret__, 'id'),
        integration_key=pulumi.get(__ret__, 'integration_key'),
        name=pulumi.get(__ret__, 'name'),
        on=pulumi.get(__ret__, 'on'),
        statements=pulumi.get(__ret__, 'statements'),
        tags=pulumi.get(__ret__, 'tags'))
def get_audit_log_subscription_output(id: Optional[pulumi.Input[str]] = None,
                                      integration_key: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAuditLogSubscriptionResult]:
    """
    Provides a LaunchDarkly audit log subscription data source.

    This data source allows you to retrieve information about LaunchDarkly audit log subscriptions.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_launchdarkly as launchdarkly

    test = launchdarkly.get_audit_log_subscription(id="5f0cd446a77cba0b4c5644a7",
        integration_key="msteams")
    ```


    :param str id: The audit log subscription ID.
    :param str integration_key: The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `pagerduty`, `signalfx`, `slack`, and `splunk`.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['integrationKey'] = integration_key
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('launchdarkly:index/getAuditLogSubscription:getAuditLogSubscription', __args__, opts=opts, typ=GetAuditLogSubscriptionResult)
    return __ret__.apply(lambda __response__: GetAuditLogSubscriptionResult(
        config=pulumi.get(__response__, 'config'),
        id=pulumi.get(__response__, 'id'),
        integration_key=pulumi.get(__response__, 'integration_key'),
        name=pulumi.get(__response__, 'name'),
        on=pulumi.get(__response__, 'on'),
        statements=pulumi.get(__response__, 'statements'),
        tags=pulumi.get(__response__, 'tags')))
