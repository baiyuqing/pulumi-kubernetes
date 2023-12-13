// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.flowcontrol.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FlowDistinguisherMethodPatch {
    /**
     * @return `type` is the type of flow distinguisher method The supported types are &#34;ByUser&#34; and &#34;ByNamespace&#34;. Required.
     * 
     */
    private @Nullable String type;

    private FlowDistinguisherMethodPatch() {}
    /**
     * @return `type` is the type of flow distinguisher method The supported types are &#34;ByUser&#34; and &#34;ByNamespace&#34;. Required.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FlowDistinguisherMethodPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String type;
        public Builder() {}
        public Builder(FlowDistinguisherMethodPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        public FlowDistinguisherMethodPatch build() {
            final var o = new FlowDistinguisherMethodPatch();
            o.type = type;
            return o;
        }
    }
}
