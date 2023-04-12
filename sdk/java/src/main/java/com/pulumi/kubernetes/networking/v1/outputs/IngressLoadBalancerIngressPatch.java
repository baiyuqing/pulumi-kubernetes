// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.networking.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.networking.v1.outputs.IngressPortStatusPatch;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class IngressLoadBalancerIngressPatch {
    /**
     * @return hostname is set for load-balancer ingress points that are DNS based.
     * 
     */
    private @Nullable String hostname;
    /**
     * @return ip is set for load-balancer ingress points that are IP based.
     * 
     */
    private @Nullable String ip;
    /**
     * @return ports provides information about the ports exposed by this LoadBalancer.
     * 
     */
    private @Nullable List<IngressPortStatusPatch> ports;

    private IngressLoadBalancerIngressPatch() {}
    /**
     * @return hostname is set for load-balancer ingress points that are DNS based.
     * 
     */
    public Optional<String> hostname() {
        return Optional.ofNullable(this.hostname);
    }
    /**
     * @return ip is set for load-balancer ingress points that are IP based.
     * 
     */
    public Optional<String> ip() {
        return Optional.ofNullable(this.ip);
    }
    /**
     * @return ports provides information about the ports exposed by this LoadBalancer.
     * 
     */
    public List<IngressPortStatusPatch> ports() {
        return this.ports == null ? List.of() : this.ports;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IngressLoadBalancerIngressPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String hostname;
        private @Nullable String ip;
        private @Nullable List<IngressPortStatusPatch> ports;
        public Builder() {}
        public Builder(IngressLoadBalancerIngressPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.hostname = defaults.hostname;
    	      this.ip = defaults.ip;
    	      this.ports = defaults.ports;
        }

        @CustomType.Setter
        public Builder hostname(@Nullable String hostname) {
            this.hostname = hostname;
            return this;
        }
        @CustomType.Setter
        public Builder ip(@Nullable String ip) {
            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder ports(@Nullable List<IngressPortStatusPatch> ports) {
            this.ports = ports;
            return this;
        }
        public Builder ports(IngressPortStatusPatch... ports) {
            return ports(List.of(ports));
        }
        public IngressLoadBalancerIngressPatch build() {
            final var o = new IngressLoadBalancerIngressPatch();
            o.hostname = hostname;
            o.ip = ip;
            o.ports = ports;
            return o;
        }
    }
}
