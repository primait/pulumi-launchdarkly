# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .access_token import *
from .audit_log_subscription import *
from .custom_role import *
from .destination import *
from .environment import *
from .feature_flag import *
from .feature_flag_environment import *
from .flag_trigger import *
from .get_audit_log_subscription import *
from .get_environment import *
from .get_feature_flag import *
from .get_feature_flag_environment import *
from .get_flag_trigger import *
from .get_metric import *
from .get_project import *
from .get_relay_proxy_configuration import *
from .get_segment import *
from .get_team import *
from .get_team_member import *
from .get_team_members import *
from .get_webhook import *
from .metric import *
from .project import *
from .provider import *
from .relay_proxy_configuration import *
from .segment import *
from .team import *
from .team_member import *
from .team_role_mapping import *
from .webhook import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_launchdarkly.config as __config
    config = __config
    import pulumi_launchdarkly.region as __region
    region = __region
else:
    config = _utilities.lazy_import('pulumi_launchdarkly.config')
    region = _utilities.lazy_import('pulumi_launchdarkly.region')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "launchdarkly",
  "mod": "index/accessToken",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/accessToken:AccessToken": "AccessToken"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/auditLogSubscription",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/auditLogSubscription:AuditLogSubscription": "AuditLogSubscription"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/customRole",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/customRole:CustomRole": "CustomRole"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/destination",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/destination:Destination": "Destination"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/environment",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/environment:Environment": "Environment"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/featureFlag",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/featureFlag:FeatureFlag": "FeatureFlag"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/featureFlagEnvironment",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/featureFlagEnvironment:FeatureFlagEnvironment": "FeatureFlagEnvironment"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/flagTrigger",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/flagTrigger:FlagTrigger": "FlagTrigger"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/metric",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/metric:Metric": "Metric"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/project",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/project:Project": "Project"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/relayProxyConfiguration",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/relayProxyConfiguration:RelayProxyConfiguration": "RelayProxyConfiguration"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/segment",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/segment:Segment": "Segment"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/team",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/team:Team": "Team"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/teamMember",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/teamMember:TeamMember": "TeamMember"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/teamRoleMapping",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/teamRoleMapping:TeamRoleMapping": "TeamRoleMapping"
  }
 },
 {
  "pkg": "launchdarkly",
  "mod": "index/webhook",
  "fqn": "pulumi_launchdarkly",
  "classes": {
   "launchdarkly:index/webhook:Webhook": "Webhook"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "launchdarkly",
  "token": "pulumi:providers:launchdarkly",
  "fqn": "pulumi_launchdarkly",
  "class": "Provider"
 }
]
"""
)
