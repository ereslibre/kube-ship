/*
 * Copyright 2019 Rafael Fernández López <ereslibre@ereslibre.es>
 * Copyright 2019 SUSE LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package capiproviders

const (
	openstackManifests = `---
apiVersion: v1
kind: Namespace
metadata:
  labels:
    controller-tools.k8s.io: "1.0"
  name: openstack-provider-system
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  labels:
    controller-tools.k8s.io: "1.0"
  name: openstackclusterproviderspecs.openstackproviderconfig.k8s.io
spec:
  group: openstackproviderconfig.k8s.io
  names:
    kind: OpenstackClusterProviderSpec
    plural: openstackclusterproviderspecs
  scope: Namespaced
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        cloudName:
          description: The name of the cloud to use from the clouds secret
          type: string
        cloudsSecret:
          description: The name of the secret containing the openstack credentials
          type: object
        disableServerTags:
          description: 'Default: True. In case of server tag errors, set to False'
          type: boolean
        dnsNameservers:
          description: DNSNameservers is the list of nameservers for OpenStack Subnet
            being created.
          items:
            type: string
          type: array
        externalNetworkId:
          description: ExternalNetworkID is the ID of an external OpenStack Network.
            This is necessary to get public internet to the VMs.
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        managedSecurityGroups:
          description: ManagedSecurityGroups defines that kubernetes manages the OpenStack
            security groups for now, that means that we'll create two security groups,
            one allowing SSH and API access from everywhere, and another one that
            allows all traffic to/from machines belonging to that group. In the future,
            we could make this more flexible.
          type: boolean
        metadata:
          type: object
        nodeCidr:
          description: NodeCIDR is the OpenStack Subnet to be created. Cluster actuator
            will create a network, a subnet with NodeCIDR, and a router connected
            to this subnet. If you leave this empty, no network will be created.
          type: string
        tags:
          description: Tags for all resources in cluster
          items:
            type: string
          type: array
      required:
      - cloudsSecret
      - cloudName
      - managedSecurityGroups
  version: v1alpha1
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  labels:
    controller-tools.k8s.io: "1.0"
  name: openstackclusterproviderstatuses.openstackproviderconfig.k8s.io
spec:
  group: openstackproviderconfig.k8s.io
  names:
    kind: OpenstackClusterProviderStatus
    plural: openstackclusterproviderstatuses
  scope: Namespaced
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        controlPlaneSecurityGroup:
          description: 'ControlPlaneSecurityGroups contains all the information about
            the OpenStack Security Group that needs to be applied to control plane
            nodes. TODO: Maybe instead of two properties, we add a property to the
            group?'
          properties:
            id:
              type: string
            name:
              type: string
            rules:
              items:
                properties:
                  direction:
                    type: string
                  etherType:
                    type: string
                  name:
                    type: string
                  portRangeMax:
                    format: int64
                    type: integer
                  portRangeMin:
                    format: int64
                    type: integer
                  protocol:
                    type: string
                  remoteGroupID:
                    type: string
                  remoteIPPrefix:
                    type: string
                  securityGroupID:
                    type: string
                required:
                - name
                - direction
                - etherType
                - securityGroupID
                - portRangeMin
                - portRangeMax
                - protocol
                - remoteGroupID
                - remoteIPPrefix
                type: object
              type: array
          required:
          - name
          - id
          - rules
          type: object
        globalSecurityGroup:
          description: GlobalSecurityGroup contains all the information about the
            OpenStack Security Group that needs to be applied to all nodes, both control
            plane and worker nodes.
          properties:
            id:
              type: string
            name:
              type: string
            rules:
              items:
                properties:
                  direction:
                    type: string
                  etherType:
                    type: string
                  name:
                    type: string
                  portRangeMax:
                    format: int64
                    type: integer
                  portRangeMin:
                    format: int64
                    type: integer
                  protocol:
                    type: string
                  remoteGroupID:
                    type: string
                  remoteIPPrefix:
                    type: string
                  securityGroupID:
                    type: string
                required:
                - name
                - direction
                - etherType
                - securityGroupID
                - portRangeMin
                - portRangeMax
                - protocol
                - remoteGroupID
                - remoteIPPrefix
                type: object
              type: array
          required:
          - name
          - id
          - rules
          type: object
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        network:
          description: Network contains all information about the created OpenStack
            Network. It includes Subnets and Router.
          properties:
            id:
              type: string
            name:
              type: string
            router:
              properties:
                id:
                  type: string
                name:
                  type: string
              required:
              - name
              - id
              type: object
            subnet:
              properties:
                cidr:
                  type: string
                id:
                  type: string
                name:
                  type: string
              required:
              - name
              - id
              - cidr
              type: object
          required:
          - name
          - id
          type: object
  version: v1alpha1
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  labels:
    controller-tools.k8s.io: "1.0"
  name: openstackproviderspecs.openstackproviderconfig.k8s.io
spec:
  group: openstackproviderconfig.k8s.io
  names:
    kind: OpenstackProviderSpec
    plural: openstackproviderspecs
  scope: Namespaced
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        availabilityZone:
          description: The availability zone from which to launch the server.
          type: string
        cloudName:
          description: The name of the cloud to use from the clouds secret
          type: string
        cloudsSecret:
          description: The name of the secret containing the openstack credentials
          type: object
        configDrive:
          description: Config Drive support
          type: boolean
        flavor:
          description: The flavor reference for the flavor for your server instance.
          type: string
        floatingIP:
          description: The floatingIP which will be associated to the machine, only
            used for master. The floatingIP should have been created and haven't been
            associated.
          type: string
        image:
          description: The name of the image to use for your server instance. If the
            RootVolume is specified, this will be ignored and use rootVolume directly.
          type: string
        keyName:
          description: The ssh key to inject in the instance
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        networks:
          description: A networks object. Required parameter when there are multiple
            networks defined for the tenant. When you do not specify the networks
            parameter, the server attaches to the only network created for the current
            tenant.
          items:
            properties:
              filter:
                description: Filters for optional network query
                properties:
                  adminStateUp:
                    type: boolean
                  description:
                    type: string
                  id:
                    type: string
                  limit:
                    format: int64
                    type: integer
                  marker:
                    type: string
                  name:
                    type: string
                  notTags:
                    type: string
                  notTagsAny:
                    type: string
                  projectId:
                    type: string
                  shared:
                    type: boolean
                  sortDir:
                    type: string
                  sortKey:
                    type: string
                  status:
                    type: string
                  tags:
                    type: string
                  tagsAny:
                    type: string
                  tenantId:
                    type: string
                type: object
              fixedIp:
                description: A fixed IPv4 address for the NIC.
                type: string
              subnets:
                description: Subnet within a network to use
                items:
                  properties:
                    filter:
                      description: Filters for optional network query
                      properties:
                        cidr:
                          type: string
                        description:
                          type: string
                        enableDhcp:
                          type: boolean
                        id:
                          type: string
                        ipVersion:
                          format: int64
                          type: integer
                        limit:
                          format: int64
                          type: integer
                        marker:
                          type: string
                        name:
                          type: string
                        networkId:
                          type: string
                        notTags:
                          type: string
                        notTagsAny:
                          type: string
                        projectId:
                          type: string
                        sortDir:
                          type: string
                        sortKey:
                          type: string
                        subnetpoolId:
                          type: string
                        tags:
                          type: string
                        tagsAny:
                          type: string
                        tenantId:
                          type: string
                      type: object
                    uuid:
                      description: The UUID of the network. Required if you omit the
                        port attribute.
                      type: string
                  type: object
                type: array
              uuid:
                description: The UUID of the network. Required if you omit the port
                  attribute.
                type: string
            type: object
          type: array
        rootVolume:
          description: The volume metadata to boot from
          properties:
            deviceType:
              type: string
            diskSize:
              format: int64
              type: integer
            sourceType:
              type: string
            sourceUUID:
              type: string
          required:
          - deviceType
          type: object
        securityGroups:
          description: The names of the security groups to assign to the instance
          items:
            properties:
              filter:
                description: Filters used to query security groups in openstack
                properties:
                  description:
                    type: string
                  id:
                    type: string
                  limit:
                    format: int64
                    type: integer
                  marker:
                    type: string
                  name:
                    type: string
                  notTags:
                    type: string
                  notTagsAny:
                    type: string
                  projectId:
                    type: string
                  sortDir:
                    type: string
                  sortKey:
                    type: string
                  tags:
                    type: string
                  tagsAny:
                    type: string
                  tenantId:
                    type: string
                type: object
              name:
                description: Security Group name
                type: string
              uuid:
                description: Security Group UID
                type: string
            type: object
          type: array
        serverMetadata:
          description: Metadata mapping. Allows you to create a map of key value pairs
            to add to the server instance.
          type: object
        sshUserName:
          description: The machine ssh username
          type: string
        tags:
          description: Machine tags Requires Nova api 2.52 minimum!
          items:
            type: string
          type: array
        trunk:
          description: Whether the server instance is created on a trunk port or not.
          type: boolean
        userDataSecret:
          description: The name of the secret containing the user data (startup script
            in most cases)
          type: object
      required:
      - cloudsSecret
      - cloudName
      - flavor
      - image
  version: v1alpha1
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: openstack-provider-manager-secrets
  namespace: kube-system
rules:
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - create
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: openstack-provider-manager-role
rules:
- apiGroups:
  - openstackproviderconfig.k8s.io
  resources:
  - openstackmachineproviderconfigs
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - cluster.k8s.io
  resources:
  - clusters
  - clusters/status
  - machinedeployments
  - machinesets
  - machines
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - nodes
  - events
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: openstack-provider-manager-secrets-binding
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: openstack-provider-manager-secrets
subjects:
- kind: ServiceAccount
  name: default
  namespace: openstack-provider-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  creationTimestamp: null
  name: openstack-provider-manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: openstack-provider-manager-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: openstack-provider-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    control-plane: controller-manager
    controller-tools.k8s.io: "1.0"
  name: openstack-provider-controller-manager-service
  namespace: openstack-provider-system
spec:
  ports:
  - port: 443
  selector:
    control-plane: controller-manager
    controller-tools.k8s.io: "1.0"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: controller-manager
    controller-tools.k8s.io: "1.0"
  name: clusterapi-controllers
  namespace: openstack-provider-system
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: controller-manager
      controller-tools.k8s.io: "1.0"
  template:
    metadata:
      labels:
        control-plane: controller-manager
        controller-tools.k8s.io: "1.0"
    spec:
      containers:
      - env:
        - name: USER
          value: root
        - name: OS_CLOUD
          valueFrom:
            secretKeyRef:
              key: OS_CLOUD
              name: cloud-selector
        image: k8scloudprovider/openstack-cluster-api-controller:latest
        name: openstack-machine-controller
        resources:
          limits:
            cpu: 100m
            memory: 30Mi
          requests:
            cpu: 100m
            memory: 20Mi
        volumeMounts:
        - mountPath: /etc/kubernetes
          name: config
        - mountPath: /etc/openstack
          name: cloud-config
        - mountPath: /usr/bin/kubeadm
          name: kubeadm
      tolerations:
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoExecute
        key: node.kubernetes.io/not-ready
        operator: Exists
      - effect: NoExecute
        key: node.kubernetes.io/unreachable
        operator: Exists
      volumes:
      - hostPath:
          path: /etc/kubernetes
        name: config
      - name: cloud-config
        secret:
          secretName: cloud-config
      - hostPath:
          path: /usr/bin/kubeadm
        name: kubeadm
`
)
