// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.
 * 
 */
public final class EmptyDirVolumeSourcePatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final EmptyDirVolumeSourcePatchArgs Empty = new EmptyDirVolumeSourcePatchArgs();

    /**
     * medium represents what type of storage medium should back this directory. The default is &#34;&#34; which means to use the node&#39;s default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
     * 
     */
    @Import(name="medium")
    private @Nullable Output<String> medium;

    /**
     * @return medium represents what type of storage medium should back this directory. The default is &#34;&#34; which means to use the node&#39;s default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
     * 
     */
    public Optional<Output<String>> medium() {
        return Optional.ofNullable(this.medium);
    }

    /**
     * sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
     * 
     */
    @Import(name="sizeLimit")
    private @Nullable Output<String> sizeLimit;

    /**
     * @return sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
     * 
     */
    public Optional<Output<String>> sizeLimit() {
        return Optional.ofNullable(this.sizeLimit);
    }

    private EmptyDirVolumeSourcePatchArgs() {}

    private EmptyDirVolumeSourcePatchArgs(EmptyDirVolumeSourcePatchArgs $) {
        this.medium = $.medium;
        this.sizeLimit = $.sizeLimit;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EmptyDirVolumeSourcePatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EmptyDirVolumeSourcePatchArgs $;

        public Builder() {
            $ = new EmptyDirVolumeSourcePatchArgs();
        }

        public Builder(EmptyDirVolumeSourcePatchArgs defaults) {
            $ = new EmptyDirVolumeSourcePatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param medium medium represents what type of storage medium should back this directory. The default is &#34;&#34; which means to use the node&#39;s default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
         * 
         * @return builder
         * 
         */
        public Builder medium(@Nullable Output<String> medium) {
            $.medium = medium;
            return this;
        }

        /**
         * @param medium medium represents what type of storage medium should back this directory. The default is &#34;&#34; which means to use the node&#39;s default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
         * 
         * @return builder
         * 
         */
        public Builder medium(String medium) {
            return medium(Output.of(medium));
        }

        /**
         * @param sizeLimit sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
         * 
         * @return builder
         * 
         */
        public Builder sizeLimit(@Nullable Output<String> sizeLimit) {
            $.sizeLimit = sizeLimit;
            return this;
        }

        /**
         * @param sizeLimit sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
         * 
         * @return builder
         * 
         */
        public Builder sizeLimit(String sizeLimit) {
            return sizeLimit(Output.of(sizeLimit));
        }

        public EmptyDirVolumeSourcePatchArgs build() {
            return $;
        }
    }

}
