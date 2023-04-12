// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.networking.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.networking.v1.inputs.IngressPortStatusArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * IngressLoadBalancerIngress represents the status of a load-balancer ingress point.
 * 
 */
public final class IngressLoadBalancerIngressArgs extends com.pulumi.resources.ResourceArgs {

    public static final IngressLoadBalancerIngressArgs Empty = new IngressLoadBalancerIngressArgs();

    /**
     * hostname is set for load-balancer ingress points that are DNS based.
     * 
     */
    @Import(name="hostname")
    private @Nullable Output<String> hostname;

    /**
     * @return hostname is set for load-balancer ingress points that are DNS based.
     * 
     */
    public Optional<Output<String>> hostname() {
        return Optional.ofNullable(this.hostname);
    }

    /**
     * ip is set for load-balancer ingress points that are IP based.
     * 
     */
    @Import(name="ip")
    private @Nullable Output<String> ip;

    /**
     * @return ip is set for load-balancer ingress points that are IP based.
     * 
     */
    public Optional<Output<String>> ip() {
        return Optional.ofNullable(this.ip);
    }

    /**
     * ports provides information about the ports exposed by this LoadBalancer.
     * 
     */
    @Import(name="ports")
    private @Nullable Output<List<IngressPortStatusArgs>> ports;

    /**
     * @return ports provides information about the ports exposed by this LoadBalancer.
     * 
     */
    public Optional<Output<List<IngressPortStatusArgs>>> ports() {
        return Optional.ofNullable(this.ports);
    }

    private IngressLoadBalancerIngressArgs() {}

    private IngressLoadBalancerIngressArgs(IngressLoadBalancerIngressArgs $) {
        this.hostname = $.hostname;
        this.ip = $.ip;
        this.ports = $.ports;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IngressLoadBalancerIngressArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IngressLoadBalancerIngressArgs $;

        public Builder() {
            $ = new IngressLoadBalancerIngressArgs();
        }

        public Builder(IngressLoadBalancerIngressArgs defaults) {
            $ = new IngressLoadBalancerIngressArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param hostname hostname is set for load-balancer ingress points that are DNS based.
         * 
         * @return builder
         * 
         */
        public Builder hostname(@Nullable Output<String> hostname) {
            $.hostname = hostname;
            return this;
        }

        /**
         * @param hostname hostname is set for load-balancer ingress points that are DNS based.
         * 
         * @return builder
         * 
         */
        public Builder hostname(String hostname) {
            return hostname(Output.of(hostname));
        }

        /**
         * @param ip ip is set for load-balancer ingress points that are IP based.
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip ip is set for load-balancer ingress points that are IP based.
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param ports ports provides information about the ports exposed by this LoadBalancer.
         * 
         * @return builder
         * 
         */
        public Builder ports(@Nullable Output<List<IngressPortStatusArgs>> ports) {
            $.ports = ports;
            return this;
        }

        /**
         * @param ports ports provides information about the ports exposed by this LoadBalancer.
         * 
         * @return builder
         * 
         */
        public Builder ports(List<IngressPortStatusArgs> ports) {
            return ports(Output.of(ports));
        }

        /**
         * @param ports ports provides information about the ports exposed by this LoadBalancer.
         * 
         * @return builder
         * 
         */
        public Builder ports(IngressPortStatusArgs... ports) {
            return ports(List.of(ports));
        }

        public IngressLoadBalancerIngressArgs build() {
            return $;
        }
    }

}
