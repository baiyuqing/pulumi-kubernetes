# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ... import meta as _meta

__all__ = [
    'ClusterTrustBundleSpecPatchArgs',
    'ClusterTrustBundleSpecArgs',
    'ClusterTrustBundleArgs',
]

@pulumi.input_type
class ClusterTrustBundleSpecPatchArgs:
    def __init__(__self__, *,
                 signer_name: Optional[pulumi.Input[str]] = None,
                 trust_bundle: Optional[pulumi.Input[str]] = None):
        """
        ClusterTrustBundleSpec contains the signer and trust anchors.
        :param pulumi.Input[str] signer_name: signerName indicates the associated signer, if any.
               
               In order to create or update a ClusterTrustBundle that sets signerName, you must have the following cluster-scoped permission: group=certificates.k8s.io resource=signers resourceName=<the signer name> verb=attest.
               
               If signerName is not empty, then the ClusterTrustBundle object must be named with the signer name as a prefix (translating slashes to colons). For example, for the signer name `example.com/foo`, valid ClusterTrustBundle object names include `example.com:foo:abc` and `example.com:foo:v1`.
               
               If signerName is empty, then the ClusterTrustBundle object's name must not have such a prefix.
               
               List/watch requests for ClusterTrustBundles can filter on this field using a `spec.signerName=NAME` field selector.
        :param pulumi.Input[str] trust_bundle: trustBundle contains the individual X.509 trust anchors for this bundle, as PEM bundle of PEM-wrapped, DER-formatted X.509 certificates.
               
               The data must consist only of PEM certificate blocks that parse as valid X.509 certificates.  Each certificate must include a basic constraints extension with the CA bit set.  The API server will reject objects that contain duplicate certificates, or that use PEM block headers.
               
               Users of ClusterTrustBundles, including Kubelet, are free to reorder and deduplicate certificate blocks in this file according to their own logic, as well as to drop PEM block headers and inter-block data.
        """
        if signer_name is not None:
            pulumi.set(__self__, "signer_name", signer_name)
        if trust_bundle is not None:
            pulumi.set(__self__, "trust_bundle", trust_bundle)

    @property
    @pulumi.getter(name="signerName")
    def signer_name(self) -> Optional[pulumi.Input[str]]:
        """
        signerName indicates the associated signer, if any.

        In order to create or update a ClusterTrustBundle that sets signerName, you must have the following cluster-scoped permission: group=certificates.k8s.io resource=signers resourceName=<the signer name> verb=attest.

        If signerName is not empty, then the ClusterTrustBundle object must be named with the signer name as a prefix (translating slashes to colons). For example, for the signer name `example.com/foo`, valid ClusterTrustBundle object names include `example.com:foo:abc` and `example.com:foo:v1`.

        If signerName is empty, then the ClusterTrustBundle object's name must not have such a prefix.

        List/watch requests for ClusterTrustBundles can filter on this field using a `spec.signerName=NAME` field selector.
        """
        return pulumi.get(self, "signer_name")

    @signer_name.setter
    def signer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signer_name", value)

    @property
    @pulumi.getter(name="trustBundle")
    def trust_bundle(self) -> Optional[pulumi.Input[str]]:
        """
        trustBundle contains the individual X.509 trust anchors for this bundle, as PEM bundle of PEM-wrapped, DER-formatted X.509 certificates.

        The data must consist only of PEM certificate blocks that parse as valid X.509 certificates.  Each certificate must include a basic constraints extension with the CA bit set.  The API server will reject objects that contain duplicate certificates, or that use PEM block headers.

        Users of ClusterTrustBundles, including Kubelet, are free to reorder and deduplicate certificate blocks in this file according to their own logic, as well as to drop PEM block headers and inter-block data.
        """
        return pulumi.get(self, "trust_bundle")

    @trust_bundle.setter
    def trust_bundle(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trust_bundle", value)


@pulumi.input_type
class ClusterTrustBundleSpecArgs:
    def __init__(__self__, *,
                 trust_bundle: pulumi.Input[str],
                 signer_name: Optional[pulumi.Input[str]] = None):
        """
        ClusterTrustBundleSpec contains the signer and trust anchors.
        :param pulumi.Input[str] trust_bundle: trustBundle contains the individual X.509 trust anchors for this bundle, as PEM bundle of PEM-wrapped, DER-formatted X.509 certificates.
               
               The data must consist only of PEM certificate blocks that parse as valid X.509 certificates.  Each certificate must include a basic constraints extension with the CA bit set.  The API server will reject objects that contain duplicate certificates, or that use PEM block headers.
               
               Users of ClusterTrustBundles, including Kubelet, are free to reorder and deduplicate certificate blocks in this file according to their own logic, as well as to drop PEM block headers and inter-block data.
        :param pulumi.Input[str] signer_name: signerName indicates the associated signer, if any.
               
               In order to create or update a ClusterTrustBundle that sets signerName, you must have the following cluster-scoped permission: group=certificates.k8s.io resource=signers resourceName=<the signer name> verb=attest.
               
               If signerName is not empty, then the ClusterTrustBundle object must be named with the signer name as a prefix (translating slashes to colons). For example, for the signer name `example.com/foo`, valid ClusterTrustBundle object names include `example.com:foo:abc` and `example.com:foo:v1`.
               
               If signerName is empty, then the ClusterTrustBundle object's name must not have such a prefix.
               
               List/watch requests for ClusterTrustBundles can filter on this field using a `spec.signerName=NAME` field selector.
        """
        pulumi.set(__self__, "trust_bundle", trust_bundle)
        if signer_name is not None:
            pulumi.set(__self__, "signer_name", signer_name)

    @property
    @pulumi.getter(name="trustBundle")
    def trust_bundle(self) -> pulumi.Input[str]:
        """
        trustBundle contains the individual X.509 trust anchors for this bundle, as PEM bundle of PEM-wrapped, DER-formatted X.509 certificates.

        The data must consist only of PEM certificate blocks that parse as valid X.509 certificates.  Each certificate must include a basic constraints extension with the CA bit set.  The API server will reject objects that contain duplicate certificates, or that use PEM block headers.

        Users of ClusterTrustBundles, including Kubelet, are free to reorder and deduplicate certificate blocks in this file according to their own logic, as well as to drop PEM block headers and inter-block data.
        """
        return pulumi.get(self, "trust_bundle")

    @trust_bundle.setter
    def trust_bundle(self, value: pulumi.Input[str]):
        pulumi.set(self, "trust_bundle", value)

    @property
    @pulumi.getter(name="signerName")
    def signer_name(self) -> Optional[pulumi.Input[str]]:
        """
        signerName indicates the associated signer, if any.

        In order to create or update a ClusterTrustBundle that sets signerName, you must have the following cluster-scoped permission: group=certificates.k8s.io resource=signers resourceName=<the signer name> verb=attest.

        If signerName is not empty, then the ClusterTrustBundle object must be named with the signer name as a prefix (translating slashes to colons). For example, for the signer name `example.com/foo`, valid ClusterTrustBundle object names include `example.com:foo:abc` and `example.com:foo:v1`.

        If signerName is empty, then the ClusterTrustBundle object's name must not have such a prefix.

        List/watch requests for ClusterTrustBundles can filter on this field using a `spec.signerName=NAME` field selector.
        """
        return pulumi.get(self, "signer_name")

    @signer_name.setter
    def signer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signer_name", value)


@pulumi.input_type
class ClusterTrustBundleArgs:
    def __init__(__self__, *,
                 spec: pulumi.Input['ClusterTrustBundleSpecArgs'],
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None):
        """
        ClusterTrustBundle is a cluster-scoped container for X.509 trust anchors (root certificates).

        ClusterTrustBundle objects are considered to be readable by any authenticated user in the cluster, because they can be mounted by pods using the `clusterTrustBundle` projection.  All service accounts have read access to ClusterTrustBundles by default.  Users who only have namespace-level access to a cluster can read ClusterTrustBundles by impersonating a serviceaccount that they have access to.

        It can be optionally associated with a particular assigner, in which case it contains one valid set of trust anchors for that signer. Signers may have multiple associated ClusterTrustBundles; each is an independent set of trust anchors for that signer. Admission control is used to enforce that only users with permissions on the signer can create or modify the corresponding bundle.
        :param pulumi.Input['ClusterTrustBundleSpecArgs'] spec: spec contains the signer (if any) and trust anchors.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: metadata contains the object metadata.
        """
        pulumi.set(__self__, "spec", spec)
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'certificates.k8s.io/v1alpha1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'ClusterTrustBundle')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Input['ClusterTrustBundleSpecArgs']:
        """
        spec contains the signer (if any) and trust anchors.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: pulumi.Input['ClusterTrustBundleSpecArgs']):
        pulumi.set(self, "spec", value)

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
        metadata contains the object metadata.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)


