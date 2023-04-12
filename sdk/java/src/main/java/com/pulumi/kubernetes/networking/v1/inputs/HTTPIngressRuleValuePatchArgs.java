// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.networking.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.networking.v1.inputs.HTTPIngressPathPatchArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://&lt;host&gt;/&lt;path&gt;?&lt;searchpart&gt; -&gt; backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last &#39;/&#39; and before the first &#39;?&#39; or &#39;#&#39;.
 * 
 */
public final class HTTPIngressRuleValuePatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final HTTPIngressRuleValuePatchArgs Empty = new HTTPIngressRuleValuePatchArgs();

    /**
     * paths is a collection of paths that map requests to backends.
     * 
     */
    @Import(name="paths")
    private @Nullable Output<List<HTTPIngressPathPatchArgs>> paths;

    /**
     * @return paths is a collection of paths that map requests to backends.
     * 
     */
    public Optional<Output<List<HTTPIngressPathPatchArgs>>> paths() {
        return Optional.ofNullable(this.paths);
    }

    private HTTPIngressRuleValuePatchArgs() {}

    private HTTPIngressRuleValuePatchArgs(HTTPIngressRuleValuePatchArgs $) {
        this.paths = $.paths;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HTTPIngressRuleValuePatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HTTPIngressRuleValuePatchArgs $;

        public Builder() {
            $ = new HTTPIngressRuleValuePatchArgs();
        }

        public Builder(HTTPIngressRuleValuePatchArgs defaults) {
            $ = new HTTPIngressRuleValuePatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param paths paths is a collection of paths that map requests to backends.
         * 
         * @return builder
         * 
         */
        public Builder paths(@Nullable Output<List<HTTPIngressPathPatchArgs>> paths) {
            $.paths = paths;
            return this;
        }

        /**
         * @param paths paths is a collection of paths that map requests to backends.
         * 
         * @return builder
         * 
         */
        public Builder paths(List<HTTPIngressPathPatchArgs> paths) {
            return paths(Output.of(paths));
        }

        /**
         * @param paths paths is a collection of paths that map requests to backends.
         * 
         * @return builder
         * 
         */
        public Builder paths(HTTPIngressPathPatchArgs... paths) {
            return paths(List.of(paths));
        }

        public HTTPIngressRuleValuePatchArgs build() {
            return $;
        }
    }

}
