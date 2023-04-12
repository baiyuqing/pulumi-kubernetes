# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ... import core as _core
from ... import meta as _meta

__all__ = [
    'ClusterCIDRSpecPatchArgs',
    'ClusterCIDRSpecArgs',
    'ClusterCIDRArgs',
    'IPAddressSpecPatchArgs',
    'IPAddressSpecArgs',
    'IPAddressArgs',
    'ParentReferencePatchArgs',
    'ParentReferenceArgs',
]

@pulumi.input_type
class ClusterCIDRSpecPatchArgs:
    def __init__(__self__, *,
                 ipv4: Optional[pulumi.Input[str]] = None,
                 ipv6: Optional[pulumi.Input[str]] = None,
                 node_selector: Optional[pulumi.Input['_core.v1.NodeSelectorPatchArgs']] = None,
                 per_node_host_bits: Optional[pulumi.Input[int]] = None):
        """
        ClusterCIDRSpec defines the desired state of ClusterCIDR.
        :param pulumi.Input[str] ipv4: ipv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        :param pulumi.Input[str] ipv6: ipv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        :param pulumi.Input['_core.v1.NodeSelectorPatchArgs'] node_selector: nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.
        :param pulumi.Input[int] per_node_host_bits: perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
        """
        if ipv4 is not None:
            pulumi.set(__self__, "ipv4", ipv4)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)
        if node_selector is not None:
            pulumi.set(__self__, "node_selector", node_selector)
        if per_node_host_bits is not None:
            pulumi.set(__self__, "per_node_host_bits", per_node_host_bits)

    @property
    @pulumi.getter
    def ipv4(self) -> Optional[pulumi.Input[str]]:
        """
        ipv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        """
        return pulumi.get(self, "ipv4")

    @ipv4.setter
    def ipv4(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4", value)

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[pulumi.Input[str]]:
        """
        ipv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        """
        return pulumi.get(self, "ipv6")

    @ipv6.setter
    def ipv6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6", value)

    @property
    @pulumi.getter(name="nodeSelector")
    def node_selector(self) -> Optional[pulumi.Input['_core.v1.NodeSelectorPatchArgs']]:
        """
        nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.
        """
        return pulumi.get(self, "node_selector")

    @node_selector.setter
    def node_selector(self, value: Optional[pulumi.Input['_core.v1.NodeSelectorPatchArgs']]):
        pulumi.set(self, "node_selector", value)

    @property
    @pulumi.getter(name="perNodeHostBits")
    def per_node_host_bits(self) -> Optional[pulumi.Input[int]]:
        """
        perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
        """
        return pulumi.get(self, "per_node_host_bits")

    @per_node_host_bits.setter
    def per_node_host_bits(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "per_node_host_bits", value)


@pulumi.input_type
class ClusterCIDRSpecArgs:
    def __init__(__self__, *,
                 per_node_host_bits: pulumi.Input[int],
                 ipv4: Optional[pulumi.Input[str]] = None,
                 ipv6: Optional[pulumi.Input[str]] = None,
                 node_selector: Optional[pulumi.Input['_core.v1.NodeSelectorArgs']] = None):
        """
        ClusterCIDRSpec defines the desired state of ClusterCIDR.
        :param pulumi.Input[int] per_node_host_bits: perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
        :param pulumi.Input[str] ipv4: ipv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        :param pulumi.Input[str] ipv6: ipv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        :param pulumi.Input['_core.v1.NodeSelectorArgs'] node_selector: nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.
        """
        pulumi.set(__self__, "per_node_host_bits", per_node_host_bits)
        if ipv4 is not None:
            pulumi.set(__self__, "ipv4", ipv4)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)
        if node_selector is not None:
            pulumi.set(__self__, "node_selector", node_selector)

    @property
    @pulumi.getter(name="perNodeHostBits")
    def per_node_host_bits(self) -> pulumi.Input[int]:
        """
        perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
        """
        return pulumi.get(self, "per_node_host_bits")

    @per_node_host_bits.setter
    def per_node_host_bits(self, value: pulumi.Input[int]):
        pulumi.set(self, "per_node_host_bits", value)

    @property
    @pulumi.getter
    def ipv4(self) -> Optional[pulumi.Input[str]]:
        """
        ipv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        """
        return pulumi.get(self, "ipv4")

    @ipv4.setter
    def ipv4(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4", value)

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[pulumi.Input[str]]:
        """
        ipv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        """
        return pulumi.get(self, "ipv6")

    @ipv6.setter
    def ipv6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6", value)

    @property
    @pulumi.getter(name="nodeSelector")
    def node_selector(self) -> Optional[pulumi.Input['_core.v1.NodeSelectorArgs']]:
        """
        nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.
        """
        return pulumi.get(self, "node_selector")

    @node_selector.setter
    def node_selector(self, value: Optional[pulumi.Input['_core.v1.NodeSelectorArgs']]):
        pulumi.set(self, "node_selector", value)


@pulumi.input_type
class ClusterCIDRArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 spec: Optional[pulumi.Input['ClusterCIDRSpecArgs']] = None):
        """
        ClusterCIDR represents a single configuration for per-Node Pod CIDR allocations when the MultiCIDRRangeAllocator is enabled (see the config for kube-controller-manager).  A cluster may have any number of ClusterCIDR resources, all of which will be considered when allocating a CIDR for a Node.  A ClusterCIDR is eligible to be used for a given Node when the node selector matches the node in question and has free CIDRs to allocate.  In case of multiple matching ClusterCIDR resources, the allocator will attempt to break ties using internal heuristics, but any ClusterCIDR whose node selector matches the Node may be used.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['ClusterCIDRSpecArgs'] spec: spec is the desired state of the ClusterCIDR. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'networking.k8s.io/v1alpha1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'ClusterCIDR')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['ClusterCIDRSpecArgs']]:
        """
        spec is the desired state of the ClusterCIDR. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['ClusterCIDRSpecArgs']]):
        pulumi.set(self, "spec", value)


@pulumi.input_type
class IPAddressSpecPatchArgs:
    def __init__(__self__, *,
                 parent_ref: Optional[pulumi.Input['ParentReferencePatchArgs']] = None):
        """
        IPAddressSpec describe the attributes in an IP Address.
        :param pulumi.Input['ParentReferencePatchArgs'] parent_ref: ParentRef references the resource that an IPAddress is attached to. An IPAddress must reference a parent object.
        """
        if parent_ref is not None:
            pulumi.set(__self__, "parent_ref", parent_ref)

    @property
    @pulumi.getter(name="parentRef")
    def parent_ref(self) -> Optional[pulumi.Input['ParentReferencePatchArgs']]:
        """
        ParentRef references the resource that an IPAddress is attached to. An IPAddress must reference a parent object.
        """
        return pulumi.get(self, "parent_ref")

    @parent_ref.setter
    def parent_ref(self, value: Optional[pulumi.Input['ParentReferencePatchArgs']]):
        pulumi.set(self, "parent_ref", value)


@pulumi.input_type
class IPAddressSpecArgs:
    def __init__(__self__, *,
                 parent_ref: Optional[pulumi.Input['ParentReferenceArgs']] = None):
        """
        IPAddressSpec describe the attributes in an IP Address.
        :param pulumi.Input['ParentReferenceArgs'] parent_ref: ParentRef references the resource that an IPAddress is attached to. An IPAddress must reference a parent object.
        """
        if parent_ref is not None:
            pulumi.set(__self__, "parent_ref", parent_ref)

    @property
    @pulumi.getter(name="parentRef")
    def parent_ref(self) -> Optional[pulumi.Input['ParentReferenceArgs']]:
        """
        ParentRef references the resource that an IPAddress is attached to. An IPAddress must reference a parent object.
        """
        return pulumi.get(self, "parent_ref")

    @parent_ref.setter
    def parent_ref(self, value: Optional[pulumi.Input['ParentReferenceArgs']]):
        pulumi.set(self, "parent_ref", value)


@pulumi.input_type
class IPAddressArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 spec: Optional[pulumi.Input['IPAddressSpecArgs']] = None):
        """
        IPAddress represents a single IP of a single IP Family. The object is designed to be used by APIs that operate on IP addresses. The object is used by the Service core API for allocation of IP addresses. An IP address can be represented in different formats, to guarantee the uniqueness of the IP, the name of the object is the IP address in canonical format, four decimal digits separated by dots suppressing leading zeros for IPv4 and the representation defined by RFC 5952 for IPv6. Valid: 192.168.1.5 or 2001:db8::1 or 2001:db8:aaaa:bbbb:cccc:dddd:eeee:1 Invalid: 10.01.2.3 or 2001:db8:0:0:0::1
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['IPAddressSpecArgs'] spec: spec is the desired state of the IPAddress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'networking.k8s.io/v1alpha1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'IPAddress')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['IPAddressSpecArgs']]:
        """
        spec is the desired state of the IPAddress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['IPAddressSpecArgs']]):
        pulumi.set(self, "spec", value)


@pulumi.input_type
class ParentReferencePatchArgs:
    def __init__(__self__, *,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 uid: Optional[pulumi.Input[str]] = None):
        """
        ParentReference describes a reference to a parent object.
        :param pulumi.Input[str] group: Group is the group of the object being referenced.
        :param pulumi.Input[str] name: Name is the name of the object being referenced.
        :param pulumi.Input[str] namespace: Namespace is the namespace of the object being referenced.
        :param pulumi.Input[str] resource: Resource is the resource of the object being referenced.
        :param pulumi.Input[str] uid: UID is the uid of the object being referenced.
        """
        if group is not None:
            pulumi.set(__self__, "group", group)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if resource is not None:
            pulumi.set(__self__, "resource", resource)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        Group is the group of the object being referenced.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name is the name of the object being referenced.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Namespace is the namespace of the object being referenced.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def resource(self) -> Optional[pulumi.Input[str]]:
        """
        Resource is the resource of the object being referenced.
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[str]]:
        """
        UID is the uid of the object being referenced.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uid", value)


@pulumi.input_type
class ParentReferenceArgs:
    def __init__(__self__, *,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 uid: Optional[pulumi.Input[str]] = None):
        """
        ParentReference describes a reference to a parent object.
        :param pulumi.Input[str] group: Group is the group of the object being referenced.
        :param pulumi.Input[str] name: Name is the name of the object being referenced.
        :param pulumi.Input[str] namespace: Namespace is the namespace of the object being referenced.
        :param pulumi.Input[str] resource: Resource is the resource of the object being referenced.
        :param pulumi.Input[str] uid: UID is the uid of the object being referenced.
        """
        if group is not None:
            pulumi.set(__self__, "group", group)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if resource is not None:
            pulumi.set(__self__, "resource", resource)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        Group is the group of the object being referenced.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name is the name of the object being referenced.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Namespace is the namespace of the object being referenced.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def resource(self) -> Optional[pulumi.Input[str]]:
        """
        Resource is the resource of the object being referenced.
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[str]]:
        """
        UID is the uid of the object being referenced.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uid", value)


