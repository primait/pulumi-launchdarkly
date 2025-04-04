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
    public sealed class FlagTriggerInstructions
    {
        /// <summary>
        /// The action to perform when triggering. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`.
        /// </summary>
        public readonly string Kind;

        [OutputConstructor]
        private FlagTriggerInstructions(string kind)
        {
            Kind = kind;
        }
    }
}
