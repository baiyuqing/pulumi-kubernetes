// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.autoscaling.v2.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.autoscaling.v2.outputs.MetricValueStatus;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ContainerResourceMetricStatus {
    /**
     * @return container is the name of the container in the pods of the scaling target
     * 
     */
    private String container;
    /**
     * @return current contains the current value for the given metric
     * 
     */
    private MetricValueStatus current;
    /**
     * @return name is the name of the resource in question.
     * 
     */
    private String name;

    private ContainerResourceMetricStatus() {}
    /**
     * @return container is the name of the container in the pods of the scaling target
     * 
     */
    public String container() {
        return this.container;
    }
    /**
     * @return current contains the current value for the given metric
     * 
     */
    public MetricValueStatus current() {
        return this.current;
    }
    /**
     * @return name is the name of the resource in question.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerResourceMetricStatus defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String container;
        private MetricValueStatus current;
        private String name;
        public Builder() {}
        public Builder(ContainerResourceMetricStatus defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.container = defaults.container;
    	      this.current = defaults.current;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder container(String container) {
            this.container = Objects.requireNonNull(container);
            return this;
        }
        @CustomType.Setter
        public Builder current(MetricValueStatus current) {
            this.current = Objects.requireNonNull(current);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public ContainerResourceMetricStatus build() {
            final var o = new ContainerResourceMetricStatus();
            o.container = container;
            o.current = current;
            o.name = name;
            return o;
        }
    }
}
