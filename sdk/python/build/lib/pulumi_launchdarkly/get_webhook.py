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
    'GetWebhookResult',
    'AwaitableGetWebhookResult',
    'get_webhook',
    'get_webhook_output',
]

@pulumi.output_type
class GetWebhookResult:
    """
    A collection of values returned by getWebhook.
    """
    def __init__(__self__, id=None, name=None, on=None, secret=None, statements=None, tags=None, url=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if on and not isinstance(on, bool):
            raise TypeError("Expected argument 'on' to be a bool")
        pulumi.set(__self__, "on", on)
        if secret and not isinstance(secret, str):
            raise TypeError("Expected argument 'secret' to be a str")
        pulumi.set(__self__, "secret", secret)
        if statements and not isinstance(statements, list):
            raise TypeError("Expected argument 'statements' to be a list")
        pulumi.set(__self__, "statements", statements)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The unique webhook ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The webhook's human-readable name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def on(self) -> bool:
        """
        Whether the webhook is enabled.
        """
        return pulumi.get(self, "on")

    @property
    @pulumi.getter
    def secret(self) -> str:
        """
        The secret used to sign the webhook.
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def statements(self) -> Sequence['outputs.GetWebhookStatementResult']:
        """
        List of policy statement blocks used to filter webhook events. For more information on webhook policy filters read [Adding a policy filter](https://docs.launchdarkly.com/integrations/webhooks#adding-a-policy-filter).
        """
        return pulumi.get(self, "statements")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        Tags associated with your resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The URL of the remote webhook.
        """
        return pulumi.get(self, "url")


class AwaitableGetWebhookResult(GetWebhookResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebhookResult(
            id=self.id,
            name=self.name,
            on=self.on,
            secret=self.secret,
            statements=self.statements,
            tags=self.tags,
            url=self.url)


def get_webhook(id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebhookResult:
    """
    Provides a LaunchDarkly webhook data source.

    This data source allows you to retrieve webhook information from your LaunchDarkly organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_launchdarkly as launchdarkly

    example = launchdarkly.get_webhook(id="57c0af6099690907435299")
    ```


    :param str id: The unique webhook ID.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('launchdarkly:index/getWebhook:getWebhook', __args__, opts=opts, typ=GetWebhookResult).value

    return AwaitableGetWebhookResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        on=pulumi.get(__ret__, 'on'),
        secret=pulumi.get(__ret__, 'secret'),
        statements=pulumi.get(__ret__, 'statements'),
        tags=pulumi.get(__ret__, 'tags'),
        url=pulumi.get(__ret__, 'url'))
def get_webhook_output(id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWebhookResult]:
    """
    Provides a LaunchDarkly webhook data source.

    This data source allows you to retrieve webhook information from your LaunchDarkly organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_launchdarkly as launchdarkly

    example = launchdarkly.get_webhook(id="57c0af6099690907435299")
    ```


    :param str id: The unique webhook ID.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('launchdarkly:index/getWebhook:getWebhook', __args__, opts=opts, typ=GetWebhookResult)
    return __ret__.apply(lambda __response__: GetWebhookResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        on=pulumi.get(__response__, 'on'),
        secret=pulumi.get(__response__, 'secret'),
        statements=pulumi.get(__response__, 'statements'),
        tags=pulumi.get(__response__, 'tags'),
        url=pulumi.get(__response__, 'url')))
