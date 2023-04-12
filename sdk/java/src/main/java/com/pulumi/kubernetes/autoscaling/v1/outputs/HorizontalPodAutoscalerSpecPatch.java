// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.autoscaling.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.autoscaling.v1.outputs.CrossVersionObjectReferencePatch;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class HorizontalPodAutoscalerSpecPatch {
    /**
     * @return maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
     * 
     */
    private @Nullable Integer maxReplicas;
    /**
     * @return minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
     * 
     */
    private @Nullable Integer minReplicas;
    /**
     * @return reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
     * 
     */
    private @Nullable CrossVersionObjectReferencePatch scaleTargetRef;
    /**
     * @return targetCPUUtilizationPercentage is the target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
     * 
     */
    private @Nullable Integer targetCPUUtilizationPercentage;

    private HorizontalPodAutoscalerSpecPatch() {}
    /**
     * @return maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
     * 
     */
    public Optional<Integer> maxReplicas() {
        return Optional.ofNullable(this.maxReplicas);
    }
    /**
     * @return minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
     * 
     */
    public Optional<Integer> minReplicas() {
        return Optional.ofNullable(this.minReplicas);
    }
    /**
     * @return reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
     * 
     */
    public Optional<CrossVersionObjectReferencePatch> scaleTargetRef() {
        return Optional.ofNullable(this.scaleTargetRef);
    }
    /**
     * @return targetCPUUtilizationPercentage is the target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
     * 
     */
    public Optional<Integer> targetCPUUtilizationPercentage() {
        return Optional.ofNullable(this.targetCPUUtilizationPercentage);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(HorizontalPodAutoscalerSpecPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer maxReplicas;
        private @Nullable Integer minReplicas;
        private @Nullable CrossVersionObjectReferencePatch scaleTargetRef;
        private @Nullable Integer targetCPUUtilizationPercentage;
        public Builder() {}
        public Builder(HorizontalPodAutoscalerSpecPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.maxReplicas = defaults.maxReplicas;
    	      this.minReplicas = defaults.minReplicas;
    	      this.scaleTargetRef = defaults.scaleTargetRef;
    	      this.targetCPUUtilizationPercentage = defaults.targetCPUUtilizationPercentage;
        }

        @CustomType.Setter
        public Builder maxReplicas(@Nullable Integer maxReplicas) {
            this.maxReplicas = maxReplicas;
            return this;
        }
        @CustomType.Setter
        public Builder minReplicas(@Nullable Integer minReplicas) {
            this.minReplicas = minReplicas;
            return this;
        }
        @CustomType.Setter
        public Builder scaleTargetRef(@Nullable CrossVersionObjectReferencePatch scaleTargetRef) {
            this.scaleTargetRef = scaleTargetRef;
            return this;
        }
        @CustomType.Setter
        public Builder targetCPUUtilizationPercentage(@Nullable Integer targetCPUUtilizationPercentage) {
            this.targetCPUUtilizationPercentage = targetCPUUtilizationPercentage;
            return this;
        }
        public HorizontalPodAutoscalerSpecPatch build() {
            final var o = new HorizontalPodAutoscalerSpecPatch();
            o.maxReplicas = maxReplicas;
            o.minReplicas = minReplicas;
            o.scaleTargetRef = scaleTargetRef;
            o.targetCPUUtilizationPercentage = targetCPUUtilizationPercentage;
            return o;
        }
    }
}
