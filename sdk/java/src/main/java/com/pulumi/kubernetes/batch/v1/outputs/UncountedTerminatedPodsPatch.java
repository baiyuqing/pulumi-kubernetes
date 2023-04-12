// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.batch.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class UncountedTerminatedPodsPatch {
    /**
     * @return failed holds UIDs of failed Pods.
     * 
     */
    private @Nullable List<String> failed;
    /**
     * @return succeeded holds UIDs of succeeded Pods.
     * 
     */
    private @Nullable List<String> succeeded;

    private UncountedTerminatedPodsPatch() {}
    /**
     * @return failed holds UIDs of failed Pods.
     * 
     */
    public List<String> failed() {
        return this.failed == null ? List.of() : this.failed;
    }
    /**
     * @return succeeded holds UIDs of succeeded Pods.
     * 
     */
    public List<String> succeeded() {
        return this.succeeded == null ? List.of() : this.succeeded;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(UncountedTerminatedPodsPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> failed;
        private @Nullable List<String> succeeded;
        public Builder() {}
        public Builder(UncountedTerminatedPodsPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.failed = defaults.failed;
    	      this.succeeded = defaults.succeeded;
        }

        @CustomType.Setter
        public Builder failed(@Nullable List<String> failed) {
            this.failed = failed;
            return this;
        }
        public Builder failed(String... failed) {
            return failed(List.of(failed));
        }
        @CustomType.Setter
        public Builder succeeded(@Nullable List<String> succeeded) {
            this.succeeded = succeeded;
            return this;
        }
        public Builder succeeded(String... succeeded) {
            return succeeded(List.of(succeeded));
        }
        public UncountedTerminatedPodsPatch build() {
            final var o = new UncountedTerminatedPodsPatch();
            o.failed = failed;
            o.succeeded = succeeded;
            return o;
        }
    }
}
