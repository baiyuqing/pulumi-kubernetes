// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.policy.v1.outputs;

import com.pulumi.core.Either;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.meta.v1.outputs.LabelSelector;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PodDisruptionBudgetSpec {
    /**
     * @return An eviction is allowed if at most &#34;maxUnavailable&#34; pods selected by &#34;selector&#34; are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with &#34;minAvailable&#34;.
     * 
     */
    private @Nullable Either<Integer,String> maxUnavailable;
    /**
     * @return An eviction is allowed if at least &#34;minAvailable&#34; pods selected by &#34;selector&#34; will still be available after the eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by specifying &#34;100%&#34;.
     * 
     */
    private @Nullable Either<Integer,String> minAvailable;
    /**
     * @return Label query over pods whose evictions are managed by the disruption budget. A null selector will match no pods, while an empty ({}) selector will select all pods within the namespace.
     * 
     */
    private @Nullable LabelSelector selector;
    /**
     * @return UnhealthyPodEvictionPolicy defines the criteria for when unhealthy pods should be considered for eviction. Current implementation considers healthy pods, as pods that have status.conditions item with type=&#34;Ready&#34;,status=&#34;True&#34;.
     * 
     * Valid policies are IfHealthyBudget and AlwaysAllow. If no policy is specified, the default behavior will be used, which corresponds to the IfHealthyBudget policy.
     * 
     * IfHealthyBudget policy means that running pods (status.phase=&#34;Running&#34;), but not yet healthy can be evicted only if the guarded application is not disrupted (status.currentHealthy is at least equal to status.desiredHealthy). Healthy pods will be subject to the PDB for eviction.
     * 
     * AlwaysAllow policy means that all running pods (status.phase=&#34;Running&#34;), but not yet healthy are considered disrupted and can be evicted regardless of whether the criteria in a PDB is met. This means perspective running pods of a disrupted application might not get a chance to become healthy. Healthy pods will be subject to the PDB for eviction.
     * 
     * Additional policies may be added in the future. Clients making eviction decisions should disallow eviction of unhealthy pods if they encounter an unrecognized policy in this field.
     * 
     * This field is beta-level. The eviction API uses this field when the feature gate PDBUnhealthyPodEvictionPolicy is enabled (enabled by default).
     * 
     */
    private @Nullable String unhealthyPodEvictionPolicy;

    private PodDisruptionBudgetSpec() {}
    /**
     * @return An eviction is allowed if at most &#34;maxUnavailable&#34; pods selected by &#34;selector&#34; are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with &#34;minAvailable&#34;.
     * 
     */
    public Optional<Either<Integer,String>> maxUnavailable() {
        return Optional.ofNullable(this.maxUnavailable);
    }
    /**
     * @return An eviction is allowed if at least &#34;minAvailable&#34; pods selected by &#34;selector&#34; will still be available after the eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by specifying &#34;100%&#34;.
     * 
     */
    public Optional<Either<Integer,String>> minAvailable() {
        return Optional.ofNullable(this.minAvailable);
    }
    /**
     * @return Label query over pods whose evictions are managed by the disruption budget. A null selector will match no pods, while an empty ({}) selector will select all pods within the namespace.
     * 
     */
    public Optional<LabelSelector> selector() {
        return Optional.ofNullable(this.selector);
    }
    /**
     * @return UnhealthyPodEvictionPolicy defines the criteria for when unhealthy pods should be considered for eviction. Current implementation considers healthy pods, as pods that have status.conditions item with type=&#34;Ready&#34;,status=&#34;True&#34;.
     * 
     * Valid policies are IfHealthyBudget and AlwaysAllow. If no policy is specified, the default behavior will be used, which corresponds to the IfHealthyBudget policy.
     * 
     * IfHealthyBudget policy means that running pods (status.phase=&#34;Running&#34;), but not yet healthy can be evicted only if the guarded application is not disrupted (status.currentHealthy is at least equal to status.desiredHealthy). Healthy pods will be subject to the PDB for eviction.
     * 
     * AlwaysAllow policy means that all running pods (status.phase=&#34;Running&#34;), but not yet healthy are considered disrupted and can be evicted regardless of whether the criteria in a PDB is met. This means perspective running pods of a disrupted application might not get a chance to become healthy. Healthy pods will be subject to the PDB for eviction.
     * 
     * Additional policies may be added in the future. Clients making eviction decisions should disallow eviction of unhealthy pods if they encounter an unrecognized policy in this field.
     * 
     * This field is beta-level. The eviction API uses this field when the feature gate PDBUnhealthyPodEvictionPolicy is enabled (enabled by default).
     * 
     */
    public Optional<String> unhealthyPodEvictionPolicy() {
        return Optional.ofNullable(this.unhealthyPodEvictionPolicy);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PodDisruptionBudgetSpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Either<Integer,String> maxUnavailable;
        private @Nullable Either<Integer,String> minAvailable;
        private @Nullable LabelSelector selector;
        private @Nullable String unhealthyPodEvictionPolicy;
        public Builder() {}
        public Builder(PodDisruptionBudgetSpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.maxUnavailable = defaults.maxUnavailable;
    	      this.minAvailable = defaults.minAvailable;
    	      this.selector = defaults.selector;
    	      this.unhealthyPodEvictionPolicy = defaults.unhealthyPodEvictionPolicy;
        }

        @CustomType.Setter
        public Builder maxUnavailable(@Nullable Either<Integer,String> maxUnavailable) {
            this.maxUnavailable = maxUnavailable;
            return this;
        }
        @CustomType.Setter
        public Builder minAvailable(@Nullable Either<Integer,String> minAvailable) {
            this.minAvailable = minAvailable;
            return this;
        }
        @CustomType.Setter
        public Builder selector(@Nullable LabelSelector selector) {
            this.selector = selector;
            return this;
        }
        @CustomType.Setter
        public Builder unhealthyPodEvictionPolicy(@Nullable String unhealthyPodEvictionPolicy) {
            this.unhealthyPodEvictionPolicy = unhealthyPodEvictionPolicy;
            return this;
        }
        public PodDisruptionBudgetSpec build() {
            final var o = new PodDisruptionBudgetSpec();
            o.maxUnavailable = maxUnavailable;
            o.minAvailable = minAvailable;
            o.selector = selector;
            o.unhealthyPodEvictionPolicy = unhealthyPodEvictionPolicy;
            return o;
        }
    }
}
