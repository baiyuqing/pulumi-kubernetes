// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.resource.v1alpha2.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.resource.v1alpha2.outputs.AllocationResultPatch;
import com.pulumi.kubernetes.resource.v1alpha2.outputs.ResourceClaimConsumerReferencePatch;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ResourceClaimStatusPatch {
    /**
     * @return Allocation is set by the resource driver once a resource or set of resources has been allocated successfully. If this is not specified, the resources have not been allocated yet.
     * 
     */
    private @Nullable AllocationResultPatch allocation;
    /**
     * @return DeallocationRequested indicates that a ResourceClaim is to be deallocated.
     * 
     * The driver then must deallocate this claim and reset the field together with clearing the Allocation field.
     * 
     * While DeallocationRequested is set, no new consumers may be added to ReservedFor.
     * 
     */
    private @Nullable Boolean deallocationRequested;
    /**
     * @return DriverName is a copy of the driver name from the ResourceClass at the time when allocation started.
     * 
     */
    private @Nullable String driverName;
    /**
     * @return ReservedFor indicates which entities are currently allowed to use the claim. A Pod which references a ResourceClaim which is not reserved for that Pod will not be started.
     * 
     * There can be at most 32 such reservations. This may get increased in the future, but not reduced.
     * 
     */
    private @Nullable List<ResourceClaimConsumerReferencePatch> reservedFor;

    private ResourceClaimStatusPatch() {}
    /**
     * @return Allocation is set by the resource driver once a resource or set of resources has been allocated successfully. If this is not specified, the resources have not been allocated yet.
     * 
     */
    public Optional<AllocationResultPatch> allocation() {
        return Optional.ofNullable(this.allocation);
    }
    /**
     * @return DeallocationRequested indicates that a ResourceClaim is to be deallocated.
     * 
     * The driver then must deallocate this claim and reset the field together with clearing the Allocation field.
     * 
     * While DeallocationRequested is set, no new consumers may be added to ReservedFor.
     * 
     */
    public Optional<Boolean> deallocationRequested() {
        return Optional.ofNullable(this.deallocationRequested);
    }
    /**
     * @return DriverName is a copy of the driver name from the ResourceClass at the time when allocation started.
     * 
     */
    public Optional<String> driverName() {
        return Optional.ofNullable(this.driverName);
    }
    /**
     * @return ReservedFor indicates which entities are currently allowed to use the claim. A Pod which references a ResourceClaim which is not reserved for that Pod will not be started.
     * 
     * There can be at most 32 such reservations. This may get increased in the future, but not reduced.
     * 
     */
    public List<ResourceClaimConsumerReferencePatch> reservedFor() {
        return this.reservedFor == null ? List.of() : this.reservedFor;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ResourceClaimStatusPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable AllocationResultPatch allocation;
        private @Nullable Boolean deallocationRequested;
        private @Nullable String driverName;
        private @Nullable List<ResourceClaimConsumerReferencePatch> reservedFor;
        public Builder() {}
        public Builder(ResourceClaimStatusPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allocation = defaults.allocation;
    	      this.deallocationRequested = defaults.deallocationRequested;
    	      this.driverName = defaults.driverName;
    	      this.reservedFor = defaults.reservedFor;
        }

        @CustomType.Setter
        public Builder allocation(@Nullable AllocationResultPatch allocation) {
            this.allocation = allocation;
            return this;
        }
        @CustomType.Setter
        public Builder deallocationRequested(@Nullable Boolean deallocationRequested) {
            this.deallocationRequested = deallocationRequested;
            return this;
        }
        @CustomType.Setter
        public Builder driverName(@Nullable String driverName) {
            this.driverName = driverName;
            return this;
        }
        @CustomType.Setter
        public Builder reservedFor(@Nullable List<ResourceClaimConsumerReferencePatch> reservedFor) {
            this.reservedFor = reservedFor;
            return this;
        }
        public Builder reservedFor(ResourceClaimConsumerReferencePatch... reservedFor) {
            return reservedFor(List.of(reservedFor));
        }
        public ResourceClaimStatusPatch build() {
            final var o = new ResourceClaimStatusPatch();
            o.allocation = allocation;
            o.deallocationRequested = deallocationRequested;
            o.driverName = driverName;
            o.reservedFor = reservedFor;
            return o;
        }
    }
}
