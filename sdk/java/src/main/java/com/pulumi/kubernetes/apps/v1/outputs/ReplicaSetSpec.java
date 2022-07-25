// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.apps.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.core.v1.outputs.PodTemplateSpec;
import com.pulumi.kubernetes.meta.v1.outputs.LabelSelector;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ReplicaSetSpec {
    /**
     * @return Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
     * 
     */
    private @Nullable Integer minReadySeconds;
    /**
     * @return Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
     * 
     */
    private @Nullable Integer replicas;
    /**
     * @return Selector is a label query over pods that should match the replica count. Label keys and values that must match in order to be controlled by this replica set. It must match the pod template&#39;s labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
     * 
     */
    private LabelSelector selector;
    /**
     * @return Template is the object that describes the pod that will be created if insufficient replicas are detected. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
     * 
     */
    private @Nullable PodTemplateSpec template;

    private ReplicaSetSpec() {}
    /**
     * @return Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
     * 
     */
    public Optional<Integer> minReadySeconds() {
        return Optional.ofNullable(this.minReadySeconds);
    }
    /**
     * @return Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
     * 
     */
    public Optional<Integer> replicas() {
        return Optional.ofNullable(this.replicas);
    }
    /**
     * @return Selector is a label query over pods that should match the replica count. Label keys and values that must match in order to be controlled by this replica set. It must match the pod template&#39;s labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
     * 
     */
    public LabelSelector selector() {
        return this.selector;
    }
    /**
     * @return Template is the object that describes the pod that will be created if insufficient replicas are detected. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
     * 
     */
    public Optional<PodTemplateSpec> template() {
        return Optional.ofNullable(this.template);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ReplicaSetSpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer minReadySeconds;
        private @Nullable Integer replicas;
        private LabelSelector selector;
        private @Nullable PodTemplateSpec template;
        public Builder() {}
        public Builder(ReplicaSetSpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.minReadySeconds = defaults.minReadySeconds;
    	      this.replicas = defaults.replicas;
    	      this.selector = defaults.selector;
    	      this.template = defaults.template;
        }

        @CustomType.Setter
        public Builder minReadySeconds(@Nullable Integer minReadySeconds) {
            this.minReadySeconds = minReadySeconds;
            return this;
        }
        @CustomType.Setter
        public Builder replicas(@Nullable Integer replicas) {
            this.replicas = replicas;
            return this;
        }
        @CustomType.Setter
        public Builder selector(LabelSelector selector) {
            this.selector = Objects.requireNonNull(selector);
            return this;
        }
        @CustomType.Setter
        public Builder template(@Nullable PodTemplateSpec template) {
            this.template = template;
            return this;
        }
        public ReplicaSetSpec build() {
            final var o = new ReplicaSetSpec();
            o.minReadySeconds = minReadySeconds;
            o.replicas = replicas;
            o.selector = selector;
            o.template = template;
            return o;
        }
    }
}