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
    'GetSegmentResult',
    'AwaitableGetSegmentResult',
    'get_segment',
    'get_segment_output',
]

@pulumi.output_type
class GetSegmentResult:
    """
    A collection of values returned by getSegment.
    """
    def __init__(__self__, creation_date=None, description=None, env_key=None, excluded_contexts=None, excludeds=None, id=None, included_contexts=None, includeds=None, key=None, name=None, project_key=None, rules=None, tags=None, unbounded=None, unbounded_context_kind=None):
        if creation_date and not isinstance(creation_date, int):
            raise TypeError("Expected argument 'creation_date' to be a int")
        pulumi.set(__self__, "creation_date", creation_date)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if env_key and not isinstance(env_key, str):
            raise TypeError("Expected argument 'env_key' to be a str")
        pulumi.set(__self__, "env_key", env_key)
        if excluded_contexts and not isinstance(excluded_contexts, list):
            raise TypeError("Expected argument 'excluded_contexts' to be a list")
        pulumi.set(__self__, "excluded_contexts", excluded_contexts)
        if excludeds and not isinstance(excludeds, list):
            raise TypeError("Expected argument 'excludeds' to be a list")
        pulumi.set(__self__, "excludeds", excludeds)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if included_contexts and not isinstance(included_contexts, list):
            raise TypeError("Expected argument 'included_contexts' to be a list")
        pulumi.set(__self__, "included_contexts", included_contexts)
        if includeds and not isinstance(includeds, list):
            raise TypeError("Expected argument 'includeds' to be a list")
        pulumi.set(__self__, "includeds", includeds)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_key and not isinstance(project_key, str):
            raise TypeError("Expected argument 'project_key' to be a str")
        pulumi.set(__self__, "project_key", project_key)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if unbounded and not isinstance(unbounded, bool):
            raise TypeError("Expected argument 'unbounded' to be a bool")
        pulumi.set(__self__, "unbounded", unbounded)
        if unbounded_context_kind and not isinstance(unbounded_context_kind, str):
            raise TypeError("Expected argument 'unbounded_context_kind' to be a str")
        pulumi.set(__self__, "unbounded_context_kind", unbounded_context_kind)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> int:
        """
        The segment's creation date represented as a UNIX epoch timestamp.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the segment's purpose.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="envKey")
    def env_key(self) -> str:
        """
        The segment's environment key.
        """
        return pulumi.get(self, "env_key")

    @property
    @pulumi.getter(name="excludedContexts")
    def excluded_contexts(self) -> Sequence['outputs.GetSegmentExcludedContextResult']:
        """
        List of non-user target objects excluded from the segment. This attribute is not valid when `unbounded` is set to `true`.
        """
        return pulumi.get(self, "excluded_contexts")

    @property
    @pulumi.getter
    def excludeds(self) -> Sequence[str]:
        """
        List of user keys excluded from the segment. To target on other context kinds, use the excluded_contexts block attribute. This attribute is not valid when `unbounded` is set to `true`.
        """
        return pulumi.get(self, "excludeds")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includedContexts")
    def included_contexts(self) -> Sequence['outputs.GetSegmentIncludedContextResult']:
        """
        List of non-user target objects included in the segment. This attribute is not valid when `unbounded` is set to `true`.
        """
        return pulumi.get(self, "included_contexts")

    @property
    @pulumi.getter
    def includeds(self) -> Sequence[str]:
        """
        List of user keys included in the segment. To target on other context kinds, use the included_contexts block attribute. This attribute is not valid when `unbounded` is set to `true`.
        """
        return pulumi.get(self, "includeds")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The unique key that references the segment.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The human-friendly name for the segment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> str:
        """
        The segment's project key.
        """
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetSegmentRuleResult']:
        """
        List of nested custom rule blocks to apply to the segment. This attribute is not valid when `unbounded` is set to `true`.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        Tags associated with your resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def unbounded(self) -> bool:
        return pulumi.get(self, "unbounded")

    @property
    @pulumi.getter(name="unboundedContextKind")
    def unbounded_context_kind(self) -> str:
        """
        For Big Segments, the targeted context kind. If this attribute is not specified it will default to `user`.
        """
        return pulumi.get(self, "unbounded_context_kind")


class AwaitableGetSegmentResult(GetSegmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSegmentResult(
            creation_date=self.creation_date,
            description=self.description,
            env_key=self.env_key,
            excluded_contexts=self.excluded_contexts,
            excludeds=self.excludeds,
            id=self.id,
            included_contexts=self.included_contexts,
            includeds=self.includeds,
            key=self.key,
            name=self.name,
            project_key=self.project_key,
            rules=self.rules,
            tags=self.tags,
            unbounded=self.unbounded,
            unbounded_context_kind=self.unbounded_context_kind)


def get_segment(env_key: Optional[str] = None,
                key: Optional[str] = None,
                project_key: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSegmentResult:
    """
    Provides a LaunchDarkly segment data source.

    This data source allows you to retrieve segment information from your LaunchDarkly organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_launchdarkly as launchdarkly

    example = launchdarkly.get_segment(key="example-segment",
        project_key="example-project",
        env_key="example-env")
    ```


    :param str env_key: The segment's environment key.
    :param str key: The unique key that references the segment.
    :param str project_key: The segment's project key.
    """
    __args__ = dict()
    __args__['envKey'] = env_key
    __args__['key'] = key
    __args__['projectKey'] = project_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('launchdarkly:index/getSegment:getSegment', __args__, opts=opts, typ=GetSegmentResult).value

    return AwaitableGetSegmentResult(
        creation_date=pulumi.get(__ret__, 'creation_date'),
        description=pulumi.get(__ret__, 'description'),
        env_key=pulumi.get(__ret__, 'env_key'),
        excluded_contexts=pulumi.get(__ret__, 'excluded_contexts'),
        excludeds=pulumi.get(__ret__, 'excludeds'),
        id=pulumi.get(__ret__, 'id'),
        included_contexts=pulumi.get(__ret__, 'included_contexts'),
        includeds=pulumi.get(__ret__, 'includeds'),
        key=pulumi.get(__ret__, 'key'),
        name=pulumi.get(__ret__, 'name'),
        project_key=pulumi.get(__ret__, 'project_key'),
        rules=pulumi.get(__ret__, 'rules'),
        tags=pulumi.get(__ret__, 'tags'),
        unbounded=pulumi.get(__ret__, 'unbounded'),
        unbounded_context_kind=pulumi.get(__ret__, 'unbounded_context_kind'))
def get_segment_output(env_key: Optional[pulumi.Input[str]] = None,
                       key: Optional[pulumi.Input[str]] = None,
                       project_key: Optional[pulumi.Input[str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSegmentResult]:
    """
    Provides a LaunchDarkly segment data source.

    This data source allows you to retrieve segment information from your LaunchDarkly organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_launchdarkly as launchdarkly

    example = launchdarkly.get_segment(key="example-segment",
        project_key="example-project",
        env_key="example-env")
    ```


    :param str env_key: The segment's environment key.
    :param str key: The unique key that references the segment.
    :param str project_key: The segment's project key.
    """
    __args__ = dict()
    __args__['envKey'] = env_key
    __args__['key'] = key
    __args__['projectKey'] = project_key
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('launchdarkly:index/getSegment:getSegment', __args__, opts=opts, typ=GetSegmentResult)
    return __ret__.apply(lambda __response__: GetSegmentResult(
        creation_date=pulumi.get(__response__, 'creation_date'),
        description=pulumi.get(__response__, 'description'),
        env_key=pulumi.get(__response__, 'env_key'),
        excluded_contexts=pulumi.get(__response__, 'excluded_contexts'),
        excludeds=pulumi.get(__response__, 'excludeds'),
        id=pulumi.get(__response__, 'id'),
        included_contexts=pulumi.get(__response__, 'included_contexts'),
        includeds=pulumi.get(__response__, 'includeds'),
        key=pulumi.get(__response__, 'key'),
        name=pulumi.get(__response__, 'name'),
        project_key=pulumi.get(__response__, 'project_key'),
        rules=pulumi.get(__response__, 'rules'),
        tags=pulumi.get(__response__, 'tags'),
        unbounded=pulumi.get(__response__, 'unbounded'),
        unbounded_context_kind=pulumi.get(__response__, 'unbounded_context_kind')))
