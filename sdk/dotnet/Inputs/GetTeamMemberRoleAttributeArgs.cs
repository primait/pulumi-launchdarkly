// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Launchdarkly.Inputs
{

    public sealed class GetTeamMemberRoleAttributeInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key / name of your role attribute. In the example `$${roleAttribute/testAttribute}`, the key is `testAttribute`.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// A list of values for your role attribute. For example, if your policy statement defines the resource `"proj/$${roleAttribute/testAttribute}"`, the values would be the keys of the projects you wanted to assign access to.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public GetTeamMemberRoleAttributeInputArgs()
        {
        }
        public static new GetTeamMemberRoleAttributeInputArgs Empty => new GetTeamMemberRoleAttributeInputArgs();
    }
}
