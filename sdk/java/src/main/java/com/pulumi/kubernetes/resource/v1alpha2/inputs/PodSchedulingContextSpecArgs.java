// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.resource.v1alpha2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * PodSchedulingContextSpec describes where resources for the Pod are needed.
 * 
 */
public final class PodSchedulingContextSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final PodSchedulingContextSpecArgs Empty = new PodSchedulingContextSpecArgs();

    /**
     * PotentialNodes lists nodes where the Pod might be able to run.
     * 
     * The size of this field is limited to 128. This is large enough for many clusters. Larger clusters may need more attempts to find a node that suits all pending resources. This may get increased in the future, but not reduced.
     * 
     */
    @Import(name="potentialNodes")
    private @Nullable Output<List<String>> potentialNodes;

    /**
     * @return PotentialNodes lists nodes where the Pod might be able to run.
     * 
     * The size of this field is limited to 128. This is large enough for many clusters. Larger clusters may need more attempts to find a node that suits all pending resources. This may get increased in the future, but not reduced.
     * 
     */
    public Optional<Output<List<String>>> potentialNodes() {
        return Optional.ofNullable(this.potentialNodes);
    }

    /**
     * SelectedNode is the node for which allocation of ResourceClaims that are referenced by the Pod and that use &#34;WaitForFirstConsumer&#34; allocation is to be attempted.
     * 
     */
    @Import(name="selectedNode")
    private @Nullable Output<String> selectedNode;

    /**
     * @return SelectedNode is the node for which allocation of ResourceClaims that are referenced by the Pod and that use &#34;WaitForFirstConsumer&#34; allocation is to be attempted.
     * 
     */
    public Optional<Output<String>> selectedNode() {
        return Optional.ofNullable(this.selectedNode);
    }

    private PodSchedulingContextSpecArgs() {}

    private PodSchedulingContextSpecArgs(PodSchedulingContextSpecArgs $) {
        this.potentialNodes = $.potentialNodes;
        this.selectedNode = $.selectedNode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PodSchedulingContextSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PodSchedulingContextSpecArgs $;

        public Builder() {
            $ = new PodSchedulingContextSpecArgs();
        }

        public Builder(PodSchedulingContextSpecArgs defaults) {
            $ = new PodSchedulingContextSpecArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param potentialNodes PotentialNodes lists nodes where the Pod might be able to run.
         * 
         * The size of this field is limited to 128. This is large enough for many clusters. Larger clusters may need more attempts to find a node that suits all pending resources. This may get increased in the future, but not reduced.
         * 
         * @return builder
         * 
         */
        public Builder potentialNodes(@Nullable Output<List<String>> potentialNodes) {
            $.potentialNodes = potentialNodes;
            return this;
        }

        /**
         * @param potentialNodes PotentialNodes lists nodes where the Pod might be able to run.
         * 
         * The size of this field is limited to 128. This is large enough for many clusters. Larger clusters may need more attempts to find a node that suits all pending resources. This may get increased in the future, but not reduced.
         * 
         * @return builder
         * 
         */
        public Builder potentialNodes(List<String> potentialNodes) {
            return potentialNodes(Output.of(potentialNodes));
        }

        /**
         * @param potentialNodes PotentialNodes lists nodes where the Pod might be able to run.
         * 
         * The size of this field is limited to 128. This is large enough for many clusters. Larger clusters may need more attempts to find a node that suits all pending resources. This may get increased in the future, but not reduced.
         * 
         * @return builder
         * 
         */
        public Builder potentialNodes(String... potentialNodes) {
            return potentialNodes(List.of(potentialNodes));
        }

        /**
         * @param selectedNode SelectedNode is the node for which allocation of ResourceClaims that are referenced by the Pod and that use &#34;WaitForFirstConsumer&#34; allocation is to be attempted.
         * 
         * @return builder
         * 
         */
        public Builder selectedNode(@Nullable Output<String> selectedNode) {
            $.selectedNode = selectedNode;
            return this;
        }

        /**
         * @param selectedNode SelectedNode is the node for which allocation of ResourceClaims that are referenced by the Pod and that use &#34;WaitForFirstConsumer&#34; allocation is to be attempted.
         * 
         * @return builder
         * 
         */
        public Builder selectedNode(String selectedNode) {
            return selectedNode(Output.of(selectedNode));
        }

        public PodSchedulingContextSpecArgs build() {
            return $;
        }
    }

}
