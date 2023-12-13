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
 * ModifyVolumeStatus represents the status object of ControllerModifyVolume operation
 * 
 */
public final class ModifyVolumeStatusArgs extends com.pulumi.resources.ResourceArgs {

    public static final ModifyVolumeStatusArgs Empty = new ModifyVolumeStatusArgs();

    /**
     * status is the status of the ControllerModifyVolume operation. It can be in any of following states:
     *  - Pending
     *    Pending indicates that the PersistentVolumeClaim cannot be modified due to unmet requirements, such as
     *    the specified VolumeAttributesClass not existing.
     *  - InProgress
     *    InProgress indicates that the volume is being modified.
     *  - Infeasible
     *      Infeasible indicates that the request has been rejected as invalid by the CSI driver. To
     *       resolve the error, a valid VolumeAttributesClass needs to be specified.
     *    Note: New statuses can be added in the future. Consumers should check for unknown statuses and fail appropriately.
     * 
     */
    @Import(name="status", required=true)
    private Output<String> status;

    /**
     * @return status is the status of the ControllerModifyVolume operation. It can be in any of following states:
     *  - Pending
     *    Pending indicates that the PersistentVolumeClaim cannot be modified due to unmet requirements, such as
     *    the specified VolumeAttributesClass not existing.
     *  - InProgress
     *    InProgress indicates that the volume is being modified.
     *  - Infeasible
     *      Infeasible indicates that the request has been rejected as invalid by the CSI driver. To
     *       resolve the error, a valid VolumeAttributesClass needs to be specified.
     *    Note: New statuses can be added in the future. Consumers should check for unknown statuses and fail appropriately.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     * targetVolumeAttributesClassName is the name of the VolumeAttributesClass the PVC currently being reconciled
     * 
     */
    @Import(name="targetVolumeAttributesClassName")
    private @Nullable Output<String> targetVolumeAttributesClassName;

    /**
     * @return targetVolumeAttributesClassName is the name of the VolumeAttributesClass the PVC currently being reconciled
     * 
     */
    public Optional<Output<String>> targetVolumeAttributesClassName() {
        return Optional.ofNullable(this.targetVolumeAttributesClassName);
    }

    private ModifyVolumeStatusArgs() {}

    private ModifyVolumeStatusArgs(ModifyVolumeStatusArgs $) {
        this.status = $.status;
        this.targetVolumeAttributesClassName = $.targetVolumeAttributesClassName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ModifyVolumeStatusArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ModifyVolumeStatusArgs $;

        public Builder() {
            $ = new ModifyVolumeStatusArgs();
        }

        public Builder(ModifyVolumeStatusArgs defaults) {
            $ = new ModifyVolumeStatusArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param status status is the status of the ControllerModifyVolume operation. It can be in any of following states:
         *  - Pending
         *    Pending indicates that the PersistentVolumeClaim cannot be modified due to unmet requirements, such as
         *    the specified VolumeAttributesClass not existing.
         *  - InProgress
         *    InProgress indicates that the volume is being modified.
         *  - Infeasible
         *      Infeasible indicates that the request has been rejected as invalid by the CSI driver. To
         *       resolve the error, a valid VolumeAttributesClass needs to be specified.
         *    Note: New statuses can be added in the future. Consumers should check for unknown statuses and fail appropriately.
         * 
         * @return builder
         * 
         */
        public Builder status(Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status status is the status of the ControllerModifyVolume operation. It can be in any of following states:
         *  - Pending
         *    Pending indicates that the PersistentVolumeClaim cannot be modified due to unmet requirements, such as
         *    the specified VolumeAttributesClass not existing.
         *  - InProgress
         *    InProgress indicates that the volume is being modified.
         *  - Infeasible
         *      Infeasible indicates that the request has been rejected as invalid by the CSI driver. To
         *       resolve the error, a valid VolumeAttributesClass needs to be specified.
         *    Note: New statuses can be added in the future. Consumers should check for unknown statuses and fail appropriately.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param targetVolumeAttributesClassName targetVolumeAttributesClassName is the name of the VolumeAttributesClass the PVC currently being reconciled
         * 
         * @return builder
         * 
         */
        public Builder targetVolumeAttributesClassName(@Nullable Output<String> targetVolumeAttributesClassName) {
            $.targetVolumeAttributesClassName = targetVolumeAttributesClassName;
            return this;
        }

        /**
         * @param targetVolumeAttributesClassName targetVolumeAttributesClassName is the name of the VolumeAttributesClass the PVC currently being reconciled
         * 
         * @return builder
         * 
         */
        public Builder targetVolumeAttributesClassName(String targetVolumeAttributesClassName) {
            return targetVolumeAttributesClassName(Output.of(targetVolumeAttributesClassName));
        }

        public ModifyVolumeStatusArgs build() {
            $.status = Objects.requireNonNull($.status, "expected parameter 'status' to be non-null");
            return $;
        }
    }

}
