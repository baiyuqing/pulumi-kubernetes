# *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from typing import Any, Callable, Optional, Sequence

import pulumi.runtime
from pulumi_kubernetes.yaml.yaml import _get_child_options, _get_invoke_options, _parse_yaml_document

from .. import _utilities, _tables

__all__ = ['Directory']


class Directory(pulumi.ComponentResource):
    resources: pulumi.Output[dict]
    """
    Kubernetes resources contained in this Directory.
    """

    def __init__(self,
                 name: str,
                 directory: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 transformations: Optional[Sequence[Callable[[Any, pulumi.ResourceOptions], None]]] = None,
                 resource_prefix: Optional[str] = None):
        """
        Directory is a component representing a collection of resources described by a kustomize directory
        (kustomization).

        ## Example Usage
        ### Local Kustomize Directory

        ```python
        from pulumi_kubernetes.kustomize import Directory

        hello_world = Directory(
            "hello-world-local",
            directory="./helloWorld",
        )
        ```
        ### Kustomize Directory from a Git Repo
        ```python
        from pulumi_kubernetes.kustomize import Directory

        hello_world = Directory(
            "hello-world-remote",
            directory="https://github.com/kubernetes-sigs/kustomize/tree/v3.3.1/examples/helloWorld",
        )
        ```
        ### Kustomize Directory with Transformations

        ```python
        from pulumi_kubernetes.helm.v3 import Chart, ChartOpts, FetchOpts

        # Make every service private to the cluster, i.e., turn all services into ClusterIP instead of LoadBalancer.
        def make_service_private(obj, opts):
            if obj["kind"] == "Service" and obj["apiVersion"] == "v1":
                try:
                    t = obj["spec"]["type"]
                    if t == "LoadBalancer":
                        obj["spec"]["type"] = "ClusterIP"
                except KeyError:
                    pass


        # Set a resource alias for a previous name.
        def alias(obj, opts):
            if obj["kind"] == "Deployment":
                opts.aliases = ["oldName"]


        # Omit a resource from the Chart by transforming the specified resource definition to an empty List.
        def omit_resource(obj, opts):
            if obj["kind"] == "Pod" and obj["metadata"]["name"] == "test":
                obj["apiVersion"] = "v1"
                obj["kind"] = "List"


        hello_world = Directory(
            "hello-world-remote",
            directory="https://github.com/kubernetes-sigs/kustomize/tree/v3.3.1/examples/helloWorld",
            transformations=[make_service_private, alias, omit_resource],
        )
        ```

        :param str name: A name for a resource.
        :param str directory: The directory containing the kustomization to apply. The value can be a local directory
               or a folder in a git repository.
               Example: ./helloWorld
               Example: https://github.com/kubernetes-sigs/kustomize/tree/master/examples/helloWorld
        :param Optional[pulumi.ResourceOptions] opts: A bag of optional settings that control a resource's behavior.
        :param Optional[Sequence[Callable[[Any, pulumi.ResourceOptions], None]]] transformations: A set of
               transformations to apply to Kubernetes resource definitions before registering with engine.
        :param Optional[str] resource_prefix: An optional prefix for the auto-generated resource names.
               Example: A resource created with resource_prefix="foo" would produce a resource named "foo-resourceName".
        """
        if not name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(name, str):
            raise TypeError('Expected resource name to be a string')
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if resource_prefix:
            name = f"{resource_prefix}-{name}"
        super(Directory, self).__init__("kubernetes:kustomize:Directory", name, __props__, opts)

        child_opts = _get_child_options(self, opts)
        invoke_opts = _get_invoke_options(child_opts)
        
        __ret__ = pulumi.runtime.invoke(
            'kubernetes:kustomize:directory', {'directory': directory}, invoke_opts)

        # Handle the cases when the provider is not fully configured:
        #   https://github.com/pulumi/pulumi/blob/v3.60.1/sdk/go/common/resource/plugin/provider_plugin.go#L1364-L1367
        result = (__ret__.value or {}).get('result', [])

        # Note: Unlike NodeJS, Python requires that we "pull" on our futures in order to get them scheduled for
        # execution. In order to do this, we leverage the engine's RegisterResourceOutputs to wait for the
        # resolution of all resources that this YAML document created.
        self.resources = _parse_yaml_document(result, child_opts, transformations, resource_prefix)
        self.register_outputs({"resources": self.resources})

    def translate_output_property(self, prop: str) -> str:
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

    def get_resource(self,
                     group_version_kind: str,
                     name: str,
                     namespace: Optional[str] = None) -> pulumi.Output[pulumi.CustomResource]:
        """
        get_resource returns a resource defined by a built-in Kubernetes group/version/kind and
        name. For example: `get_resource("apps/v1/Deployment", "nginx")`

        :param str group_version_kind: Group/Version/Kind of the resource, e.g., `apps/v1/Deployment`
        :param str name: Name of the resource to retrieve
        :param Optional[str] namespace: Optional namespace of the resource to retrieve
        """

        # `id` will either be `${name}` or `${namespace}/${name}`.
        id = pulumi.Output.from_input(name)
        if namespace is not None:
            id = pulumi.Output.concat(namespace, '/', name)

        resource_id = id.apply(lambda x: f'{group_version_kind}:{x}')
        return resource_id.apply(lambda x: self.resources[x])
