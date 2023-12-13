// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Storage.V1Alpha1
{
    /// <summary>
    /// VolumeAttributesClassList is a collection of VolumeAttributesClass objects.
    /// </summary>
    [KubernetesResourceType("kubernetes:storage.k8s.io/v1alpha1:VolumeAttributesClassList")]
    public partial class VolumeAttributesClassList : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// items is the list of VolumeAttributesClass objects.
        /// </summary>
        [Output("items")]
        public Output<ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Storage.V1Alpha1.VolumeAttributesClass>> Items { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ListMeta> Metadata { get; private set; } = null!;


        /// <summary>
        /// Create a VolumeAttributesClassList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VolumeAttributesClassList(string name, Pulumi.Kubernetes.Types.Inputs.Storage.V1Alpha1.VolumeAttributesClassListArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:storage.k8s.io/v1alpha1:VolumeAttributesClassList", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal VolumeAttributesClassList(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:storage.k8s.io/v1alpha1:VolumeAttributesClassList", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private VolumeAttributesClassList(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:storage.k8s.io/v1alpha1:VolumeAttributesClassList", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Storage.V1Alpha1.VolumeAttributesClassListArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Storage.V1Alpha1.VolumeAttributesClassListArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Storage.V1Alpha1.VolumeAttributesClassListArgs();
            args.ApiVersion = "storage.k8s.io/v1alpha1";
            args.Kind = "VolumeAttributesClassList";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VolumeAttributesClassList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VolumeAttributesClassList Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VolumeAttributesClassList(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Storage.V1Alpha1
{

    public class VolumeAttributesClassListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        [Input("items", required: true)]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Storage.V1Alpha1.VolumeAttributesClassArgs>? _items;

        /// <summary>
        /// items is the list of VolumeAttributesClass objects.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Storage.V1Alpha1.VolumeAttributesClassArgs> Items
        {
            get => _items ?? (_items = new InputList<Pulumi.Kubernetes.Types.Inputs.Storage.V1Alpha1.VolumeAttributesClassArgs>());
            set => _items = value;
        }

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ListMetaArgs>? Metadata { get; set; }

        public VolumeAttributesClassListArgs()
        {
        }
        public static new VolumeAttributesClassListArgs Empty => new VolumeAttributesClassListArgs();
    }
}
