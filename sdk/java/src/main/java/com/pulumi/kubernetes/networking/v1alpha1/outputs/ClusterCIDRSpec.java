// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.networking.v1alpha1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.core.v1.outputs.NodeSelector;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterCIDRSpec {
    /**
     * @return ipv4 defines an IPv4 IP block in CIDR notation(e.g. &#34;10.0.0.0/8&#34;). At least one of ipv4 and ipv6 must be specified. This field is immutable.
     * 
     */
    private @Nullable String ipv4;
    /**
     * @return ipv6 defines an IPv6 IP block in CIDR notation(e.g. &#34;2001:db8::/64&#34;). At least one of ipv4 and ipv6 must be specified. This field is immutable.
     * 
     */
    private @Nullable String ipv6;
    /**
     * @return nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.
     * 
     */
    private @Nullable NodeSelector nodeSelector;
    /**
     * @return perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
     * 
     */
    private Integer perNodeHostBits;

    private ClusterCIDRSpec() {}
    /**
     * @return ipv4 defines an IPv4 IP block in CIDR notation(e.g. &#34;10.0.0.0/8&#34;). At least one of ipv4 and ipv6 must be specified. This field is immutable.
     * 
     */
    public Optional<String> ipv4() {
        return Optional.ofNullable(this.ipv4);
    }
    /**
     * @return ipv6 defines an IPv6 IP block in CIDR notation(e.g. &#34;2001:db8::/64&#34;). At least one of ipv4 and ipv6 must be specified. This field is immutable.
     * 
     */
    public Optional<String> ipv6() {
        return Optional.ofNullable(this.ipv6);
    }
    /**
     * @return nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.
     * 
     */
    public Optional<NodeSelector> nodeSelector() {
        return Optional.ofNullable(this.nodeSelector);
    }
    /**
     * @return perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
     * 
     */
    public Integer perNodeHostBits() {
        return this.perNodeHostBits;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterCIDRSpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String ipv4;
        private @Nullable String ipv6;
        private @Nullable NodeSelector nodeSelector;
        private Integer perNodeHostBits;
        public Builder() {}
        public Builder(ClusterCIDRSpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipv4 = defaults.ipv4;
    	      this.ipv6 = defaults.ipv6;
    	      this.nodeSelector = defaults.nodeSelector;
    	      this.perNodeHostBits = defaults.perNodeHostBits;
        }

        @CustomType.Setter
        public Builder ipv4(@Nullable String ipv4) {
            this.ipv4 = ipv4;
            return this;
        }
        @CustomType.Setter
        public Builder ipv6(@Nullable String ipv6) {
            this.ipv6 = ipv6;
            return this;
        }
        @CustomType.Setter
        public Builder nodeSelector(@Nullable NodeSelector nodeSelector) {
            this.nodeSelector = nodeSelector;
            return this;
        }
        @CustomType.Setter
        public Builder perNodeHostBits(Integer perNodeHostBits) {
            this.perNodeHostBits = Objects.requireNonNull(perNodeHostBits);
            return this;
        }
        public ClusterCIDRSpec build() {
            final var o = new ClusterCIDRSpec();
            o.ipv4 = ipv4;
            o.ipv6 = ipv6;
            o.nodeSelector = nodeSelector;
            o.perNodeHostBits = perNodeHostBits;
            return o;
        }
    }
}
