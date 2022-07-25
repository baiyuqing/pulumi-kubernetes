// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.core.v1.inputs.NodeSelectorPatchArgs;
import com.pulumi.kubernetes.core.v1.inputs.PreferredSchedulingTermPatchArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * Node affinity is a group of node affinity scheduling rules.
 * 
 */
public final class NodeAffinityPatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final NodeAffinityPatchArgs Empty = new NodeAffinityPatchArgs();

    /**
     * The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding &#34;weight&#34; to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
     * 
     */
    @Import(name="preferredDuringSchedulingIgnoredDuringExecution")
    private @Nullable Output<List<PreferredSchedulingTermPatchArgs>> preferredDuringSchedulingIgnoredDuringExecution;

    /**
     * @return The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding &#34;weight&#34; to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
     * 
     */
    public Optional<Output<List<PreferredSchedulingTermPatchArgs>>> preferredDuringSchedulingIgnoredDuringExecution() {
        return Optional.ofNullable(this.preferredDuringSchedulingIgnoredDuringExecution);
    }

    /**
     * If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.
     * 
     */
    @Import(name="requiredDuringSchedulingIgnoredDuringExecution")
    private @Nullable Output<NodeSelectorPatchArgs> requiredDuringSchedulingIgnoredDuringExecution;

    /**
     * @return If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.
     * 
     */
    public Optional<Output<NodeSelectorPatchArgs>> requiredDuringSchedulingIgnoredDuringExecution() {
        return Optional.ofNullable(this.requiredDuringSchedulingIgnoredDuringExecution);
    }

    private NodeAffinityPatchArgs() {}

    private NodeAffinityPatchArgs(NodeAffinityPatchArgs $) {
        this.preferredDuringSchedulingIgnoredDuringExecution = $.preferredDuringSchedulingIgnoredDuringExecution;
        this.requiredDuringSchedulingIgnoredDuringExecution = $.requiredDuringSchedulingIgnoredDuringExecution;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NodeAffinityPatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NodeAffinityPatchArgs $;

        public Builder() {
            $ = new NodeAffinityPatchArgs();
        }

        public Builder(NodeAffinityPatchArgs defaults) {
            $ = new NodeAffinityPatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param preferredDuringSchedulingIgnoredDuringExecution The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding &#34;weight&#34; to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
         * 
         * @return builder
         * 
         */
        public Builder preferredDuringSchedulingIgnoredDuringExecution(@Nullable Output<List<PreferredSchedulingTermPatchArgs>> preferredDuringSchedulingIgnoredDuringExecution) {
            $.preferredDuringSchedulingIgnoredDuringExecution = preferredDuringSchedulingIgnoredDuringExecution;
            return this;
        }

        /**
         * @param preferredDuringSchedulingIgnoredDuringExecution The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding &#34;weight&#34; to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
         * 
         * @return builder
         * 
         */
        public Builder preferredDuringSchedulingIgnoredDuringExecution(List<PreferredSchedulingTermPatchArgs> preferredDuringSchedulingIgnoredDuringExecution) {
            return preferredDuringSchedulingIgnoredDuringExecution(Output.of(preferredDuringSchedulingIgnoredDuringExecution));
        }

        /**
         * @param preferredDuringSchedulingIgnoredDuringExecution The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding &#34;weight&#34; to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
         * 
         * @return builder
         * 
         */
        public Builder preferredDuringSchedulingIgnoredDuringExecution(PreferredSchedulingTermPatchArgs... preferredDuringSchedulingIgnoredDuringExecution) {
            return preferredDuringSchedulingIgnoredDuringExecution(List.of(preferredDuringSchedulingIgnoredDuringExecution));
        }

        /**
         * @param requiredDuringSchedulingIgnoredDuringExecution If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.
         * 
         * @return builder
         * 
         */
        public Builder requiredDuringSchedulingIgnoredDuringExecution(@Nullable Output<NodeSelectorPatchArgs> requiredDuringSchedulingIgnoredDuringExecution) {
            $.requiredDuringSchedulingIgnoredDuringExecution = requiredDuringSchedulingIgnoredDuringExecution;
            return this;
        }

        /**
         * @param requiredDuringSchedulingIgnoredDuringExecution If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.
         * 
         * @return builder
         * 
         */
        public Builder requiredDuringSchedulingIgnoredDuringExecution(NodeSelectorPatchArgs requiredDuringSchedulingIgnoredDuringExecution) {
            return requiredDuringSchedulingIgnoredDuringExecution(Output.of(requiredDuringSchedulingIgnoredDuringExecution));
        }

        public NodeAffinityPatchArgs build() {
            return $;
        }
    }

}