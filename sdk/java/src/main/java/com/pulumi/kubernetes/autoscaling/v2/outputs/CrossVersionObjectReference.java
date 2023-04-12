// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.autoscaling.v2.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CrossVersionObjectReference {
    /**
     * @return apiVersion is the API version of the referent
     * 
     */
    private @Nullable String apiVersion;
    /**
     * @return kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    private String kind;
    /**
     * @return name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     * 
     */
    private String name;

    private CrossVersionObjectReference() {}
    /**
     * @return apiVersion is the API version of the referent
     * 
     */
    public Optional<String> apiVersion() {
        return Optional.ofNullable(this.apiVersion);
    }
    /**
     * @return kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public String kind() {
        return this.kind;
    }
    /**
     * @return name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CrossVersionObjectReference defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String apiVersion;
        private String kind;
        private String name;
        public Builder() {}
        public Builder(CrossVersionObjectReference defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiVersion = defaults.apiVersion;
    	      this.kind = defaults.kind;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder apiVersion(@Nullable String apiVersion) {
            this.apiVersion = apiVersion;
            return this;
        }
        @CustomType.Setter
        public Builder kind(String kind) {
            this.kind = Objects.requireNonNull(kind);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public CrossVersionObjectReference build() {
            final var o = new CrossVersionObjectReference();
            o.apiVersion = apiVersion;
            o.kind = kind;
            o.name = name;
            return o;
        }
    }
}
