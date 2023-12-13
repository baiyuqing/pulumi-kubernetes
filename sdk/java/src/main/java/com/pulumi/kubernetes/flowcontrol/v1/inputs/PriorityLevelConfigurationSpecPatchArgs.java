// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.flowcontrol.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.flowcontrol.v1.inputs.ExemptPriorityLevelConfigurationPatchArgs;
import com.pulumi.kubernetes.flowcontrol.v1.inputs.LimitedPriorityLevelConfigurationPatchArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * PriorityLevelConfigurationSpec specifies the configuration of a priority level.
 * 
 */
public final class PriorityLevelConfigurationSpecPatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final PriorityLevelConfigurationSpecPatchArgs Empty = new PriorityLevelConfigurationSpecPatchArgs();

    /**
     * `exempt` specifies how requests are handled for an exempt priority level. This field MUST be empty if `type` is `&#34;Limited&#34;`. This field MAY be non-empty if `type` is `&#34;Exempt&#34;`. If empty and `type` is `&#34;Exempt&#34;` then the default values for `ExemptPriorityLevelConfiguration` apply.
     * 
     */
    @Import(name="exempt")
    private @Nullable Output<ExemptPriorityLevelConfigurationPatchArgs> exempt;

    /**
     * @return `exempt` specifies how requests are handled for an exempt priority level. This field MUST be empty if `type` is `&#34;Limited&#34;`. This field MAY be non-empty if `type` is `&#34;Exempt&#34;`. If empty and `type` is `&#34;Exempt&#34;` then the default values for `ExemptPriorityLevelConfiguration` apply.
     * 
     */
    public Optional<Output<ExemptPriorityLevelConfigurationPatchArgs>> exempt() {
        return Optional.ofNullable(this.exempt);
    }

    /**
     * `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `&#34;Limited&#34;`.
     * 
     */
    @Import(name="limited")
    private @Nullable Output<LimitedPriorityLevelConfigurationPatchArgs> limited;

    /**
     * @return `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `&#34;Limited&#34;`.
     * 
     */
    public Optional<Output<LimitedPriorityLevelConfigurationPatchArgs>> limited() {
        return Optional.ofNullable(this.limited);
    }

    /**
     * `type` indicates whether this priority level is subject to limitation on request execution.  A value of `&#34;Exempt&#34;` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `&#34;Limited&#34;` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server&#39;s limited capacity is made available exclusively to this priority level. Required.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return `type` indicates whether this priority level is subject to limitation on request execution.  A value of `&#34;Exempt&#34;` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `&#34;Limited&#34;` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server&#39;s limited capacity is made available exclusively to this priority level. Required.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private PriorityLevelConfigurationSpecPatchArgs() {}

    private PriorityLevelConfigurationSpecPatchArgs(PriorityLevelConfigurationSpecPatchArgs $) {
        this.exempt = $.exempt;
        this.limited = $.limited;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PriorityLevelConfigurationSpecPatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PriorityLevelConfigurationSpecPatchArgs $;

        public Builder() {
            $ = new PriorityLevelConfigurationSpecPatchArgs();
        }

        public Builder(PriorityLevelConfigurationSpecPatchArgs defaults) {
            $ = new PriorityLevelConfigurationSpecPatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param exempt `exempt` specifies how requests are handled for an exempt priority level. This field MUST be empty if `type` is `&#34;Limited&#34;`. This field MAY be non-empty if `type` is `&#34;Exempt&#34;`. If empty and `type` is `&#34;Exempt&#34;` then the default values for `ExemptPriorityLevelConfiguration` apply.
         * 
         * @return builder
         * 
         */
        public Builder exempt(@Nullable Output<ExemptPriorityLevelConfigurationPatchArgs> exempt) {
            $.exempt = exempt;
            return this;
        }

        /**
         * @param exempt `exempt` specifies how requests are handled for an exempt priority level. This field MUST be empty if `type` is `&#34;Limited&#34;`. This field MAY be non-empty if `type` is `&#34;Exempt&#34;`. If empty and `type` is `&#34;Exempt&#34;` then the default values for `ExemptPriorityLevelConfiguration` apply.
         * 
         * @return builder
         * 
         */
        public Builder exempt(ExemptPriorityLevelConfigurationPatchArgs exempt) {
            return exempt(Output.of(exempt));
        }

        /**
         * @param limited `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `&#34;Limited&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder limited(@Nullable Output<LimitedPriorityLevelConfigurationPatchArgs> limited) {
            $.limited = limited;
            return this;
        }

        /**
         * @param limited `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `&#34;Limited&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder limited(LimitedPriorityLevelConfigurationPatchArgs limited) {
            return limited(Output.of(limited));
        }

        /**
         * @param type `type` indicates whether this priority level is subject to limitation on request execution.  A value of `&#34;Exempt&#34;` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `&#34;Limited&#34;` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server&#39;s limited capacity is made available exclusively to this priority level. Required.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type `type` indicates whether this priority level is subject to limitation on request execution.  A value of `&#34;Exempt&#34;` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `&#34;Limited&#34;` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server&#39;s limited capacity is made available exclusively to this priority level. Required.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public PriorityLevelConfigurationSpecPatchArgs build() {
            return $;
        }
    }

}
