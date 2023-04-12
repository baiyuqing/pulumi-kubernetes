// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.apiextensions.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.apiextensions.v1.outputs.WebhookConversionPatch;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CustomResourceConversionPatch {
    /**
     * @return strategy specifies how custom resources are converted between versions. Allowed values are: - `&#34;None&#34;`: The converter only change the apiVersion and would not touch any other field in the custom resource. - `&#34;Webhook&#34;`: API Server will call to an external webhook to do the conversion. Additional information
     *   is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.
     * 
     */
    private @Nullable String strategy;
    /**
     * @return webhook describes how to call the conversion webhook. Required when `strategy` is set to `&#34;Webhook&#34;`.
     * 
     */
    private @Nullable WebhookConversionPatch webhook;

    private CustomResourceConversionPatch() {}
    /**
     * @return strategy specifies how custom resources are converted between versions. Allowed values are: - `&#34;None&#34;`: The converter only change the apiVersion and would not touch any other field in the custom resource. - `&#34;Webhook&#34;`: API Server will call to an external webhook to do the conversion. Additional information
     *   is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.
     * 
     */
    public Optional<String> strategy() {
        return Optional.ofNullable(this.strategy);
    }
    /**
     * @return webhook describes how to call the conversion webhook. Required when `strategy` is set to `&#34;Webhook&#34;`.
     * 
     */
    public Optional<WebhookConversionPatch> webhook() {
        return Optional.ofNullable(this.webhook);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CustomResourceConversionPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String strategy;
        private @Nullable WebhookConversionPatch webhook;
        public Builder() {}
        public Builder(CustomResourceConversionPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.strategy = defaults.strategy;
    	      this.webhook = defaults.webhook;
        }

        @CustomType.Setter
        public Builder strategy(@Nullable String strategy) {
            this.strategy = strategy;
            return this;
        }
        @CustomType.Setter
        public Builder webhook(@Nullable WebhookConversionPatch webhook) {
            this.webhook = webhook;
            return this;
        }
        public CustomResourceConversionPatch build() {
            final var o = new CustomResourceConversionPatch();
            o.strategy = strategy;
            o.webhook = webhook;
            return o;
        }
    }
}
