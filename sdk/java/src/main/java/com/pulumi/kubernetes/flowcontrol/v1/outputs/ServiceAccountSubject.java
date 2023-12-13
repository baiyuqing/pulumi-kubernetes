// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.flowcontrol.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServiceAccountSubject {
    /**
     * @return `name` is the name of matching ServiceAccount objects, or &#34;*&#34; to match regardless of name. Required.
     * 
     */
    private String name;
    /**
     * @return `namespace` is the namespace of matching ServiceAccount objects. Required.
     * 
     */
    private String namespace;

    private ServiceAccountSubject() {}
    /**
     * @return `name` is the name of matching ServiceAccount objects, or &#34;*&#34; to match regardless of name. Required.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return `namespace` is the namespace of matching ServiceAccount objects. Required.
     * 
     */
    public String namespace() {
        return this.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceAccountSubject defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String namespace;
        public Builder() {}
        public Builder(ServiceAccountSubject defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.namespace = defaults.namespace;
        }

        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder namespace(String namespace) {
            this.namespace = Objects.requireNonNull(namespace);
            return this;
        }
        public ServiceAccountSubject build() {
            final var o = new ServiceAccountSubject();
            o.name = name;
            o.namespace = namespace;
            return o;
        }
    }
}
