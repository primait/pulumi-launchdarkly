// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Launchdarkly.Outputs
{

    [OutputType]
    public sealed class ProjectDefaultClientSideAvailability
    {
        public readonly bool UsingEnvironmentId;
        public readonly bool UsingMobileKey;

        [OutputConstructor]
        private ProjectDefaultClientSideAvailability(
            bool usingEnvironmentId,

            bool usingMobileKey)
        {
            UsingEnvironmentId = usingEnvironmentId;
            UsingMobileKey = usingMobileKey;
        }
    }
}
