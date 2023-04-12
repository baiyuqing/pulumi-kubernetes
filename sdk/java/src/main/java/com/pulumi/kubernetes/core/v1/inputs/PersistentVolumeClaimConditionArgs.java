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
 * PersistentVolumeClaimCondition contains details about state of pvc
 * 
 */
public final class PersistentVolumeClaimConditionArgs extends com.pulumi.resources.ResourceArgs {

    public static final PersistentVolumeClaimConditionArgs Empty = new PersistentVolumeClaimConditionArgs();

    /**
     * lastProbeTime is the time we probed the condition.
     * 
     */
    @Import(name="lastProbeTime")
    private @Nullable Output<String> lastProbeTime;

    /**
     * @return lastProbeTime is the time we probed the condition.
     * 
     */
    public Optional<Output<String>> lastProbeTime() {
        return Optional.ofNullable(this.lastProbeTime);
    }

    /**
     * lastTransitionTime is the time the condition transitioned from one status to another.
     * 
     */
    @Import(name="lastTransitionTime")
    private @Nullable Output<String> lastTransitionTime;

    /**
     * @return lastTransitionTime is the time the condition transitioned from one status to another.
     * 
     */
    public Optional<Output<String>> lastTransitionTime() {
        return Optional.ofNullable(this.lastTransitionTime);
    }

    /**
     * message is the human-readable message indicating details about last transition.
     * 
     */
    @Import(name="message")
    private @Nullable Output<String> message;

    /**
     * @return message is the human-readable message indicating details about last transition.
     * 
     */
    public Optional<Output<String>> message() {
        return Optional.ofNullable(this.message);
    }

    /**
     * reason is a unique, this should be a short, machine understandable string that gives the reason for condition&#39;s last transition. If it reports &#34;ResizeStarted&#34; that means the underlying persistent volume is being resized.
     * 
     */
    @Import(name="reason")
    private @Nullable Output<String> reason;

    /**
     * @return reason is a unique, this should be a short, machine understandable string that gives the reason for condition&#39;s last transition. If it reports &#34;ResizeStarted&#34; that means the underlying persistent volume is being resized.
     * 
     */
    public Optional<Output<String>> reason() {
        return Optional.ofNullable(this.reason);
    }

    @Import(name="status", required=true)
    private Output<String> status;

    public Output<String> status() {
        return this.status;
    }

    @Import(name="type", required=true)
    private Output<String> type;

    public Output<String> type() {
        return this.type;
    }

    private PersistentVolumeClaimConditionArgs() {}

    private PersistentVolumeClaimConditionArgs(PersistentVolumeClaimConditionArgs $) {
        this.lastProbeTime = $.lastProbeTime;
        this.lastTransitionTime = $.lastTransitionTime;
        this.message = $.message;
        this.reason = $.reason;
        this.status = $.status;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PersistentVolumeClaimConditionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PersistentVolumeClaimConditionArgs $;

        public Builder() {
            $ = new PersistentVolumeClaimConditionArgs();
        }

        public Builder(PersistentVolumeClaimConditionArgs defaults) {
            $ = new PersistentVolumeClaimConditionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param lastProbeTime lastProbeTime is the time we probed the condition.
         * 
         * @return builder
         * 
         */
        public Builder lastProbeTime(@Nullable Output<String> lastProbeTime) {
            $.lastProbeTime = lastProbeTime;
            return this;
        }

        /**
         * @param lastProbeTime lastProbeTime is the time we probed the condition.
         * 
         * @return builder
         * 
         */
        public Builder lastProbeTime(String lastProbeTime) {
            return lastProbeTime(Output.of(lastProbeTime));
        }

        /**
         * @param lastTransitionTime lastTransitionTime is the time the condition transitioned from one status to another.
         * 
         * @return builder
         * 
         */
        public Builder lastTransitionTime(@Nullable Output<String> lastTransitionTime) {
            $.lastTransitionTime = lastTransitionTime;
            return this;
        }

        /**
         * @param lastTransitionTime lastTransitionTime is the time the condition transitioned from one status to another.
         * 
         * @return builder
         * 
         */
        public Builder lastTransitionTime(String lastTransitionTime) {
            return lastTransitionTime(Output.of(lastTransitionTime));
        }

        /**
         * @param message message is the human-readable message indicating details about last transition.
         * 
         * @return builder
         * 
         */
        public Builder message(@Nullable Output<String> message) {
            $.message = message;
            return this;
        }

        /**
         * @param message message is the human-readable message indicating details about last transition.
         * 
         * @return builder
         * 
         */
        public Builder message(String message) {
            return message(Output.of(message));
        }

        /**
         * @param reason reason is a unique, this should be a short, machine understandable string that gives the reason for condition&#39;s last transition. If it reports &#34;ResizeStarted&#34; that means the underlying persistent volume is being resized.
         * 
         * @return builder
         * 
         */
        public Builder reason(@Nullable Output<String> reason) {
            $.reason = reason;
            return this;
        }

        /**
         * @param reason reason is a unique, this should be a short, machine understandable string that gives the reason for condition&#39;s last transition. If it reports &#34;ResizeStarted&#34; that means the underlying persistent volume is being resized.
         * 
         * @return builder
         * 
         */
        public Builder reason(String reason) {
            return reason(Output.of(reason));
        }

        public Builder status(Output<String> status) {
            $.status = status;
            return this;
        }

        public Builder status(String status) {
            return status(Output.of(status));
        }

        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public PersistentVolumeClaimConditionArgs build() {
            $.status = Objects.requireNonNull($.status, "expected parameter 'status' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
