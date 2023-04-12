// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.networking.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.networking.v1.outputs.IngressBackendPatch;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class HTTPIngressPathPatch {
    /**
     * @return backend defines the referenced service endpoint to which the traffic will be forwarded to.
     * 
     */
    private @Nullable IngressBackendPatch backend;
    /**
     * @return path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional &#34;path&#34; part of a URL as defined by RFC 3986. Paths must begin with a &#39;/&#39; and must be present when using PathType with value &#34;Exact&#34; or &#34;Prefix&#34;.
     * 
     */
    private @Nullable String path;
    /**
     * @return pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by &#39;/&#39;. Matching is
     *   done on a path element by element basis. A path element refers is the
     *   list of labels in the path split by the &#39;/&#39; separator. A request is a
     *   match for path p if every p is an element-wise prefix of p of the
     *   request path. Note that if the last element of the path is a substring
     *   of the last element in request path, it is not a match (e.g. /foo/bar
     *   matches /foo/bar/baz, but does not match /foo/barbaz).
     * * ImplementationSpecific: Interpretation of the Path matching is up to
     *   the IngressClass. Implementations can treat this as a separate PathType
     *   or treat it identically to Prefix or Exact path types.
     *   Implementations are required to support all path types.
     * 
     */
    private @Nullable String pathType;

    private HTTPIngressPathPatch() {}
    /**
     * @return backend defines the referenced service endpoint to which the traffic will be forwarded to.
     * 
     */
    public Optional<IngressBackendPatch> backend() {
        return Optional.ofNullable(this.backend);
    }
    /**
     * @return path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional &#34;path&#34; part of a URL as defined by RFC 3986. Paths must begin with a &#39;/&#39; and must be present when using PathType with value &#34;Exact&#34; or &#34;Prefix&#34;.
     * 
     */
    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }
    /**
     * @return pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by &#39;/&#39;. Matching is
     *   done on a path element by element basis. A path element refers is the
     *   list of labels in the path split by the &#39;/&#39; separator. A request is a
     *   match for path p if every p is an element-wise prefix of p of the
     *   request path. Note that if the last element of the path is a substring
     *   of the last element in request path, it is not a match (e.g. /foo/bar
     *   matches /foo/bar/baz, but does not match /foo/barbaz).
     * * ImplementationSpecific: Interpretation of the Path matching is up to
     *   the IngressClass. Implementations can treat this as a separate PathType
     *   or treat it identically to Prefix or Exact path types.
     *   Implementations are required to support all path types.
     * 
     */
    public Optional<String> pathType() {
        return Optional.ofNullable(this.pathType);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(HTTPIngressPathPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable IngressBackendPatch backend;
        private @Nullable String path;
        private @Nullable String pathType;
        public Builder() {}
        public Builder(HTTPIngressPathPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backend = defaults.backend;
    	      this.path = defaults.path;
    	      this.pathType = defaults.pathType;
        }

        @CustomType.Setter
        public Builder backend(@Nullable IngressBackendPatch backend) {
            this.backend = backend;
            return this;
        }
        @CustomType.Setter
        public Builder path(@Nullable String path) {
            this.path = path;
            return this;
        }
        @CustomType.Setter
        public Builder pathType(@Nullable String pathType) {
            this.pathType = pathType;
            return this;
        }
        public HTTPIngressPathPatch build() {
            final var o = new HTTPIngressPathPatch();
            o.backend = backend;
            o.path = path;
            o.pathType = pathType;
            return o;
        }
    }
}
