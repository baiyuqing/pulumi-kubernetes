// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.certificates.v1alpha1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterTrustBundleSpecPatch {
    /**
     * @return signerName indicates the associated signer, if any.
     * 
     * In order to create or update a ClusterTrustBundle that sets signerName, you must have the following cluster-scoped permission: group=certificates.k8s.io resource=signers resourceName=&lt;the signer name&gt; verb=attest.
     * 
     * If signerName is not empty, then the ClusterTrustBundle object must be named with the signer name as a prefix (translating slashes to colons). For example, for the signer name `example.com/foo`, valid ClusterTrustBundle object names include `example.com:foo:abc` and `example.com:foo:v1`.
     * 
     * If signerName is empty, then the ClusterTrustBundle object&#39;s name must not have such a prefix.
     * 
     * List/watch requests for ClusterTrustBundles can filter on this field using a `spec.signerName=NAME` field selector.
     * 
     */
    private @Nullable String signerName;
    /**
     * @return trustBundle contains the individual X.509 trust anchors for this bundle, as PEM bundle of PEM-wrapped, DER-formatted X.509 certificates.
     * 
     * The data must consist only of PEM certificate blocks that parse as valid X.509 certificates.  Each certificate must include a basic constraints extension with the CA bit set.  The API server will reject objects that contain duplicate certificates, or that use PEM block headers.
     * 
     * Users of ClusterTrustBundles, including Kubelet, are free to reorder and deduplicate certificate blocks in this file according to their own logic, as well as to drop PEM block headers and inter-block data.
     * 
     */
    private @Nullable String trustBundle;

    private ClusterTrustBundleSpecPatch() {}
    /**
     * @return signerName indicates the associated signer, if any.
     * 
     * In order to create or update a ClusterTrustBundle that sets signerName, you must have the following cluster-scoped permission: group=certificates.k8s.io resource=signers resourceName=&lt;the signer name&gt; verb=attest.
     * 
     * If signerName is not empty, then the ClusterTrustBundle object must be named with the signer name as a prefix (translating slashes to colons). For example, for the signer name `example.com/foo`, valid ClusterTrustBundle object names include `example.com:foo:abc` and `example.com:foo:v1`.
     * 
     * If signerName is empty, then the ClusterTrustBundle object&#39;s name must not have such a prefix.
     * 
     * List/watch requests for ClusterTrustBundles can filter on this field using a `spec.signerName=NAME` field selector.
     * 
     */
    public Optional<String> signerName() {
        return Optional.ofNullable(this.signerName);
    }
    /**
     * @return trustBundle contains the individual X.509 trust anchors for this bundle, as PEM bundle of PEM-wrapped, DER-formatted X.509 certificates.
     * 
     * The data must consist only of PEM certificate blocks that parse as valid X.509 certificates.  Each certificate must include a basic constraints extension with the CA bit set.  The API server will reject objects that contain duplicate certificates, or that use PEM block headers.
     * 
     * Users of ClusterTrustBundles, including Kubelet, are free to reorder and deduplicate certificate blocks in this file according to their own logic, as well as to drop PEM block headers and inter-block data.
     * 
     */
    public Optional<String> trustBundle() {
        return Optional.ofNullable(this.trustBundle);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterTrustBundleSpecPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String signerName;
        private @Nullable String trustBundle;
        public Builder() {}
        public Builder(ClusterTrustBundleSpecPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.signerName = defaults.signerName;
    	      this.trustBundle = defaults.trustBundle;
        }

        @CustomType.Setter
        public Builder signerName(@Nullable String signerName) {
            this.signerName = signerName;
            return this;
        }
        @CustomType.Setter
        public Builder trustBundle(@Nullable String trustBundle) {
            this.trustBundle = trustBundle;
            return this;
        }
        public ClusterTrustBundleSpecPatch build() {
            final var o = new ClusterTrustBundleSpecPatch();
            o.signerName = signerName;
            o.trustBundle = trustBundle;
            return o;
        }
    }
}
