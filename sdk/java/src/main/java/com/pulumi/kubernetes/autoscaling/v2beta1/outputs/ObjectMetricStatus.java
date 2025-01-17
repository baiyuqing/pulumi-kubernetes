// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.autoscaling.v2beta1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.autoscaling.v2beta1.outputs.CrossVersionObjectReference;
import com.pulumi.kubernetes.meta.v1.outputs.LabelSelector;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ObjectMetricStatus {
    /**
     * @return averageValue is the current value of the average of the metric across all relevant pods (as a quantity)
     * 
     */
    private @Nullable String averageValue;
    /**
     * @return currentValue is the current value of the metric (as a quantity).
     * 
     */
    private String currentValue;
    /**
     * @return metricName is the name of the metric in question.
     * 
     */
    private String metricName;
    /**
     * @return selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the ObjectMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
     * 
     */
    private @Nullable LabelSelector selector;
    /**
     * @return target is the described Kubernetes object.
     * 
     */
    private CrossVersionObjectReference target;

    private ObjectMetricStatus() {}
    /**
     * @return averageValue is the current value of the average of the metric across all relevant pods (as a quantity)
     * 
     */
    public Optional<String> averageValue() {
        return Optional.ofNullable(this.averageValue);
    }
    /**
     * @return currentValue is the current value of the metric (as a quantity).
     * 
     */
    public String currentValue() {
        return this.currentValue;
    }
    /**
     * @return metricName is the name of the metric in question.
     * 
     */
    public String metricName() {
        return this.metricName;
    }
    /**
     * @return selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the ObjectMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
     * 
     */
    public Optional<LabelSelector> selector() {
        return Optional.ofNullable(this.selector);
    }
    /**
     * @return target is the described Kubernetes object.
     * 
     */
    public CrossVersionObjectReference target() {
        return this.target;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ObjectMetricStatus defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String averageValue;
        private String currentValue;
        private String metricName;
        private @Nullable LabelSelector selector;
        private CrossVersionObjectReference target;
        public Builder() {}
        public Builder(ObjectMetricStatus defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.averageValue = defaults.averageValue;
    	      this.currentValue = defaults.currentValue;
    	      this.metricName = defaults.metricName;
    	      this.selector = defaults.selector;
    	      this.target = defaults.target;
        }

        @CustomType.Setter
        public Builder averageValue(@Nullable String averageValue) {
            this.averageValue = averageValue;
            return this;
        }
        @CustomType.Setter
        public Builder currentValue(String currentValue) {
            this.currentValue = Objects.requireNonNull(currentValue);
            return this;
        }
        @CustomType.Setter
        public Builder metricName(String metricName) {
            this.metricName = Objects.requireNonNull(metricName);
            return this;
        }
        @CustomType.Setter
        public Builder selector(@Nullable LabelSelector selector) {
            this.selector = selector;
            return this;
        }
        @CustomType.Setter
        public Builder target(CrossVersionObjectReference target) {
            this.target = Objects.requireNonNull(target);
            return this;
        }
        public ObjectMetricStatus build() {
            final var o = new ObjectMetricStatus();
            o.averageValue = averageValue;
            o.currentValue = currentValue;
            o.metricName = metricName;
            o.selector = selector;
            o.target = target;
            return o;
        }
    }
}
