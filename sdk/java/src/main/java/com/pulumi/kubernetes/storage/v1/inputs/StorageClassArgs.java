// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.storage.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.core.v1.inputs.TopologySelectorTermArgs;
import com.pulumi.kubernetes.meta.v1.inputs.ObjectMetaArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.
 * 
 * StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
 * 
 */
public final class StorageClassArgs extends com.pulumi.resources.ResourceArgs {

    public static final StorageClassArgs Empty = new StorageClassArgs();

    /**
     * allowVolumeExpansion shows whether the storage class allow volume expand.
     * 
     */
    @Import(name="allowVolumeExpansion")
    private @Nullable Output<Boolean> allowVolumeExpansion;

    /**
     * @return allowVolumeExpansion shows whether the storage class allow volume expand.
     * 
     */
    public Optional<Output<Boolean>> allowVolumeExpansion() {
        return Optional.ofNullable(this.allowVolumeExpansion);
    }

    /**
     * allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
     * 
     */
    @Import(name="allowedTopologies")
    private @Nullable Output<List<TopologySelectorTermArgs>> allowedTopologies;

    /**
     * @return allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
     * 
     */
    public Optional<Output<List<TopologySelectorTermArgs>>> allowedTopologies() {
        return Optional.ofNullable(this.allowedTopologies);
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    @Import(name="apiVersion")
    private @Nullable Output<String> apiVersion;

    /**
     * @return APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    public Optional<Output<String>> apiVersion() {
        return Optional.ofNullable(this.apiVersion);
    }

    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Import(name="kind")
    private @Nullable Output<String> kind;

    /**
     * @return Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Optional<Output<String>> kind() {
        return Optional.ofNullable(this.kind);
    }

    /**
     * Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<ObjectMetaArgs> metadata;

    /**
     * @return Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Optional<Output<ObjectMetaArgs>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. [&#34;ro&#34;, &#34;soft&#34;]. Not validated - mount of the PVs will simply fail if one is invalid.
     * 
     */
    @Import(name="mountOptions")
    private @Nullable Output<List<String>> mountOptions;

    /**
     * @return mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. [&#34;ro&#34;, &#34;soft&#34;]. Not validated - mount of the PVs will simply fail if one is invalid.
     * 
     */
    public Optional<Output<List<String>>> mountOptions() {
        return Optional.ofNullable(this.mountOptions);
    }

    /**
     * parameters holds the parameters for the provisioner that should create volumes of this storage class.
     * 
     */
    @Import(name="parameters")
    private @Nullable Output<Map<String,String>> parameters;

    /**
     * @return parameters holds the parameters for the provisioner that should create volumes of this storage class.
     * 
     */
    public Optional<Output<Map<String,String>>> parameters() {
        return Optional.ofNullable(this.parameters);
    }

    /**
     * provisioner indicates the type of the provisioner.
     * 
     */
    @Import(name="provisioner", required=true)
    private Output<String> provisioner;

    /**
     * @return provisioner indicates the type of the provisioner.
     * 
     */
    public Output<String> provisioner() {
        return this.provisioner;
    }

    /**
     * reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
     * 
     */
    @Import(name="reclaimPolicy")
    private @Nullable Output<String> reclaimPolicy;

    /**
     * @return reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
     * 
     */
    public Optional<Output<String>> reclaimPolicy() {
        return Optional.ofNullable(this.reclaimPolicy);
    }

    /**
     * volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
     * 
     */
    @Import(name="volumeBindingMode")
    private @Nullable Output<String> volumeBindingMode;

    /**
     * @return volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
     * 
     */
    public Optional<Output<String>> volumeBindingMode() {
        return Optional.ofNullable(this.volumeBindingMode);
    }

    private StorageClassArgs() {}

    private StorageClassArgs(StorageClassArgs $) {
        this.allowVolumeExpansion = $.allowVolumeExpansion;
        this.allowedTopologies = $.allowedTopologies;
        this.apiVersion = $.apiVersion;
        this.kind = $.kind;
        this.metadata = $.metadata;
        this.mountOptions = $.mountOptions;
        this.parameters = $.parameters;
        this.provisioner = $.provisioner;
        this.reclaimPolicy = $.reclaimPolicy;
        this.volumeBindingMode = $.volumeBindingMode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StorageClassArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StorageClassArgs $;

        public Builder() {
            $ = new StorageClassArgs();
        }

        public Builder(StorageClassArgs defaults) {
            $ = new StorageClassArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowVolumeExpansion allowVolumeExpansion shows whether the storage class allow volume expand.
         * 
         * @return builder
         * 
         */
        public Builder allowVolumeExpansion(@Nullable Output<Boolean> allowVolumeExpansion) {
            $.allowVolumeExpansion = allowVolumeExpansion;
            return this;
        }

        /**
         * @param allowVolumeExpansion allowVolumeExpansion shows whether the storage class allow volume expand.
         * 
         * @return builder
         * 
         */
        public Builder allowVolumeExpansion(Boolean allowVolumeExpansion) {
            return allowVolumeExpansion(Output.of(allowVolumeExpansion));
        }

        /**
         * @param allowedTopologies allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
         * 
         * @return builder
         * 
         */
        public Builder allowedTopologies(@Nullable Output<List<TopologySelectorTermArgs>> allowedTopologies) {
            $.allowedTopologies = allowedTopologies;
            return this;
        }

        /**
         * @param allowedTopologies allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
         * 
         * @return builder
         * 
         */
        public Builder allowedTopologies(List<TopologySelectorTermArgs> allowedTopologies) {
            return allowedTopologies(Output.of(allowedTopologies));
        }

        /**
         * @param allowedTopologies allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
         * 
         * @return builder
         * 
         */
        public Builder allowedTopologies(TopologySelectorTermArgs... allowedTopologies) {
            return allowedTopologies(List.of(allowedTopologies));
        }

        /**
         * @param apiVersion APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
         * 
         * @return builder
         * 
         */
        public Builder apiVersion(@Nullable Output<String> apiVersion) {
            $.apiVersion = apiVersion;
            return this;
        }

        /**
         * @param apiVersion APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
         * 
         * @return builder
         * 
         */
        public Builder apiVersion(String apiVersion) {
            return apiVersion(Output.of(apiVersion));
        }

        /**
         * @param kind Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
         * 
         * @return builder
         * 
         */
        public Builder kind(@Nullable Output<String> kind) {
            $.kind = kind;
            return this;
        }

        /**
         * @param kind Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
         * 
         * @return builder
         * 
         */
        public Builder kind(String kind) {
            return kind(Output.of(kind));
        }

        /**
         * @param metadata Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<ObjectMetaArgs> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
         * 
         * @return builder
         * 
         */
        public Builder metadata(ObjectMetaArgs metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param mountOptions mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. [&#34;ro&#34;, &#34;soft&#34;]. Not validated - mount of the PVs will simply fail if one is invalid.
         * 
         * @return builder
         * 
         */
        public Builder mountOptions(@Nullable Output<List<String>> mountOptions) {
            $.mountOptions = mountOptions;
            return this;
        }

        /**
         * @param mountOptions mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. [&#34;ro&#34;, &#34;soft&#34;]. Not validated - mount of the PVs will simply fail if one is invalid.
         * 
         * @return builder
         * 
         */
        public Builder mountOptions(List<String> mountOptions) {
            return mountOptions(Output.of(mountOptions));
        }

        /**
         * @param mountOptions mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. [&#34;ro&#34;, &#34;soft&#34;]. Not validated - mount of the PVs will simply fail if one is invalid.
         * 
         * @return builder
         * 
         */
        public Builder mountOptions(String... mountOptions) {
            return mountOptions(List.of(mountOptions));
        }

        /**
         * @param parameters parameters holds the parameters for the provisioner that should create volumes of this storage class.
         * 
         * @return builder
         * 
         */
        public Builder parameters(@Nullable Output<Map<String,String>> parameters) {
            $.parameters = parameters;
            return this;
        }

        /**
         * @param parameters parameters holds the parameters for the provisioner that should create volumes of this storage class.
         * 
         * @return builder
         * 
         */
        public Builder parameters(Map<String,String> parameters) {
            return parameters(Output.of(parameters));
        }

        /**
         * @param provisioner provisioner indicates the type of the provisioner.
         * 
         * @return builder
         * 
         */
        public Builder provisioner(Output<String> provisioner) {
            $.provisioner = provisioner;
            return this;
        }

        /**
         * @param provisioner provisioner indicates the type of the provisioner.
         * 
         * @return builder
         * 
         */
        public Builder provisioner(String provisioner) {
            return provisioner(Output.of(provisioner));
        }

        /**
         * @param reclaimPolicy reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
         * 
         * @return builder
         * 
         */
        public Builder reclaimPolicy(@Nullable Output<String> reclaimPolicy) {
            $.reclaimPolicy = reclaimPolicy;
            return this;
        }

        /**
         * @param reclaimPolicy reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
         * 
         * @return builder
         * 
         */
        public Builder reclaimPolicy(String reclaimPolicy) {
            return reclaimPolicy(Output.of(reclaimPolicy));
        }

        /**
         * @param volumeBindingMode volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
         * 
         * @return builder
         * 
         */
        public Builder volumeBindingMode(@Nullable Output<String> volumeBindingMode) {
            $.volumeBindingMode = volumeBindingMode;
            return this;
        }

        /**
         * @param volumeBindingMode volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
         * 
         * @return builder
         * 
         */
        public Builder volumeBindingMode(String volumeBindingMode) {
            return volumeBindingMode(Output.of(volumeBindingMode));
        }

        public StorageClassArgs build() {
            $.apiVersion = Codegen.stringProp("apiVersion").output().arg($.apiVersion).getNullable();
            $.kind = Codegen.stringProp("kind").output().arg($.kind).getNullable();
            $.provisioner = Objects.requireNonNull($.provisioner, "expected parameter 'provisioner' to be non-null");
            return $;
        }
    }

}
