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
    'GetFlagTriggerResult',
    'AwaitableGetFlagTriggerResult',
    'get_flag_trigger',
    'get_flag_trigger_output',
]

@pulumi.output_type
class GetFlagTriggerResult:
    """
    A collection of values returned by getFlagTrigger.
    """
    def __init__(__self__, enabled=None, env_key=None, flag_key=None, id=None, instructions=None, integration_key=None, maintainer_id=None, project_key=None, trigger_url=None):
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if env_key and not isinstance(env_key, str):
            raise TypeError("Expected argument 'env_key' to be a str")
        pulumi.set(__self__, "env_key", env_key)
        if flag_key and not isinstance(flag_key, str):
            raise TypeError("Expected argument 'flag_key' to be a str")
        pulumi.set(__self__, "flag_key", flag_key)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instructions and not isinstance(instructions, list):
            raise TypeError("Expected argument 'instructions' to be a list")
        pulumi.set(__self__, "instructions", instructions)
        if integration_key and not isinstance(integration_key, str):
            raise TypeError("Expected argument 'integration_key' to be a str")
        pulumi.set(__self__, "integration_key", integration_key)
        if maintainer_id and not isinstance(maintainer_id, str):
            raise TypeError("Expected argument 'maintainer_id' to be a str")
        pulumi.set(__self__, "maintainer_id", maintainer_id)
        if project_key and not isinstance(project_key, str):
            raise TypeError("Expected argument 'project_key' to be a str")
        pulumi.set(__self__, "project_key", project_key)
        if trigger_url and not isinstance(trigger_url, str):
            raise TypeError("Expected argument 'trigger_url' to be a str")
        pulumi.set(__self__, "trigger_url", trigger_url)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="envKey")
    def env_key(self) -> str:
        """
        The unique key of the environment the flag trigger will work in.
        """
        return pulumi.get(self, "env_key")

    @property
    @pulumi.getter(name="flagKey")
    def flag_key(self) -> str:
        """
        The unique key of the associated flag.
        """
        return pulumi.get(self, "flag_key")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def instructions(self) -> Sequence['outputs.GetFlagTriggerInstructionResult']:
        return pulumi.get(self, "instructions")

    @property
    @pulumi.getter(name="integrationKey")
    def integration_key(self) -> str:
        return pulumi.get(self, "integration_key")

    @property
    @pulumi.getter(name="maintainerId")
    def maintainer_id(self) -> str:
        return pulumi.get(self, "maintainer_id")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> str:
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter(name="triggerUrl")
    def trigger_url(self) -> str:
        return pulumi.get(self, "trigger_url")


class AwaitableGetFlagTriggerResult(GetFlagTriggerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlagTriggerResult(
            enabled=self.enabled,
            env_key=self.env_key,
            flag_key=self.flag_key,
            id=self.id,
            instructions=self.instructions,
            integration_key=self.integration_key,
            maintainer_id=self.maintainer_id,
            project_key=self.project_key,
            trigger_url=self.trigger_url)


def get_flag_trigger(env_key: Optional[str] = None,
                     flag_key: Optional[str] = None,
                     id: Optional[str] = None,
                     project_key: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlagTriggerResult:
    """
    Provides a LaunchDarkly flag trigger data source.

    > **Note:** Flag triggers are available to customers on an Enterprise LaunchDarkly plan. To learn more, [read about our pricing](https://launchdarkly.com/pricing/). To upgrade your plan, [contact LaunchDarkly Sales](https://launchdarkly.com/contact-sales/).

    This data source allows you to retrieve information about flag triggers from your LaunchDarkly organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_launchdarkly as launchdarkly

    example = launchdarkly.get_flag_trigger(id="61d490757f7821150815518f",
        flag_key="example-flag",
        project_key="the-big-project",
        env_key="production")
    ```


    :param str env_key: The unique key of the environment the flag trigger will work in.
    :param str flag_key: The unique key of the associated flag.
    """
    __args__ = dict()
    __args__['envKey'] = env_key
    __args__['flagKey'] = flag_key
    __args__['id'] = id
    __args__['projectKey'] = project_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('launchdarkly:index/getFlagTrigger:getFlagTrigger', __args__, opts=opts, typ=GetFlagTriggerResult).value

    return AwaitableGetFlagTriggerResult(
        enabled=pulumi.get(__ret__, 'enabled'),
        env_key=pulumi.get(__ret__, 'env_key'),
        flag_key=pulumi.get(__ret__, 'flag_key'),
        id=pulumi.get(__ret__, 'id'),
        instructions=pulumi.get(__ret__, 'instructions'),
        integration_key=pulumi.get(__ret__, 'integration_key'),
        maintainer_id=pulumi.get(__ret__, 'maintainer_id'),
        project_key=pulumi.get(__ret__, 'project_key'),
        trigger_url=pulumi.get(__ret__, 'trigger_url'))
def get_flag_trigger_output(env_key: Optional[pulumi.Input[str]] = None,
                            flag_key: Optional[pulumi.Input[str]] = None,
                            id: Optional[pulumi.Input[str]] = None,
                            project_key: Optional[pulumi.Input[str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFlagTriggerResult]:
    """
    Provides a LaunchDarkly flag trigger data source.

    > **Note:** Flag triggers are available to customers on an Enterprise LaunchDarkly plan. To learn more, [read about our pricing](https://launchdarkly.com/pricing/). To upgrade your plan, [contact LaunchDarkly Sales](https://launchdarkly.com/contact-sales/).

    This data source allows you to retrieve information about flag triggers from your LaunchDarkly organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_launchdarkly as launchdarkly

    example = launchdarkly.get_flag_trigger(id="61d490757f7821150815518f",
        flag_key="example-flag",
        project_key="the-big-project",
        env_key="production")
    ```


    :param str env_key: The unique key of the environment the flag trigger will work in.
    :param str flag_key: The unique key of the associated flag.
    """
    __args__ = dict()
    __args__['envKey'] = env_key
    __args__['flagKey'] = flag_key
    __args__['id'] = id
    __args__['projectKey'] = project_key
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('launchdarkly:index/getFlagTrigger:getFlagTrigger', __args__, opts=opts, typ=GetFlagTriggerResult)
    return __ret__.apply(lambda __response__: GetFlagTriggerResult(
        enabled=pulumi.get(__response__, 'enabled'),
        env_key=pulumi.get(__response__, 'env_key'),
        flag_key=pulumi.get(__response__, 'flag_key'),
        id=pulumi.get(__response__, 'id'),
        instructions=pulumi.get(__response__, 'instructions'),
        integration_key=pulumi.get(__response__, 'integration_key'),
        maintainer_id=pulumi.get(__response__, 'maintainer_id'),
        project_key=pulumi.get(__response__, 'project_key'),
        trigger_url=pulumi.get(__response__, 'trigger_url')))
