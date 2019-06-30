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

package capi

const (
	manifests = `---
apiVersion: v1
kind: Namespace
metadata:
  labels:
    controller-tools.k8s.io: "1.0"
  name: cluster-api-system
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  name: clusters.cluster.k8s.io
spec:
  group: cluster.k8s.io
  names:
    kind: Cluster
    plural: clusters
    shortNames:
    - cl
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: / [Cluster] Cluster is the Schema for the clusters API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        metadata:
          properties:
            annotations:
              additionalProperties:
                type: string
              description: 'Annotations is an unstructured key value map stored with
                a resource that may be set by external tools to store and retrieve
                arbitrary metadata. They are not queryable and should be preserved
                when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations'
              type: object
            clusterName:
              description: The name of the cluster which the object belongs to. This
                is used to distinguish resources with same name and namespace in different
                clusters. This field is not set anywhere right now and apiserver is
                going to ignore it if set in create or update request.
              type: string
            creationTimestamp:
              description: "CreationTimestamp is a timestamp representing the server
                time when this object was created. It is not guaranteed to be set
                in happens-before order across separate operations. Clients may not
                set this value. It is represented in RFC3339 form and is in UTC. \n
                Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            deletionGracePeriodSeconds:
              description: Number of seconds allowed for this object to gracefully
                terminate before it will be removed from the system. Only set when
                deletionTimestamp is also set. May only be shortened. Read-only.
              format: int64
              type: integer
            deletionTimestamp:
              description: "DeletionTimestamp is RFC 3339 date and time at which this
                resource will be deleted. This field is set by the server when a graceful
                deletion is requested by the user, and is not directly settable by
                a client. The resource is expected to be deleted (no longer visible
                from resource lists, and not reachable by name) after the time in
                this field, once the finalizers list is empty. As long as the finalizers
                list contains items, deletion is blocked. Once the deletionTimestamp
                is set, this value may not be unset or be set further into the future,
                although it may be shortened or the resource may be deleted prior
                to this time. For example, a user may request that a pod is deleted
                in 30 seconds. The Kubelet will react by sending a graceful termination
                signal to the containers in the pod. After that 30 seconds, the Kubelet
                will send a hard termination signal (SIGKILL) to the container and
                after cleanup, remove the pod from the API. In the presence of network
                partitions, this object may still exist after this timestamp, until
                an administrator or automated process can determine the resource is
                fully terminated. If not set, graceful deletion of the object has
                not been requested. \n Populated by the system when a graceful deletion
                is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            finalizers:
              description: Must be empty before the object is deleted from the registry.
                Each entry is an identifier for the responsible component that will
                remove the entry from the list. If the deletionTimestamp of the object
                is non-nil, entries in this list can only be removed.
              items:
                type: string
              type: array
            generateName:
              description: "GenerateName is an optional prefix, used by the server,
                to generate a unique name ONLY IF the Name field has not been provided.
                If this field is used, the name returned to the client will be different
                than the name passed. This value will also be combined with a unique
                suffix. The provided value has the same validation rules as the Name
                field, and may be truncated by the length of the suffix required to
                make the value unique on the server. \n If this field is specified
                and the generated name exists, the server will NOT return a 409 -
                instead, it will either return 201 Created or 500 with Reason ServerTimeout
                indicating a unique name could not be found in the time allotted,
                and the client should retry (optionally after the time indicated in
                the Retry-After header). \n Applied only if Name is not specified.
                More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
              type: string
            generation:
              description: A sequence number representing a specific generation of
                the desired state. Populated by the system. Read-only.
              format: int64
              type: integer
            initializers:
              description: "An initializer is a controller which enforces some system
                invariant at object creation time. This field is a list of initializers
                that have not yet acted on this object. If nil or empty, this object
                has been completely initialized. Otherwise, the object is considered
                uninitialized and is hidden (in list/watch and get calls) from clients
                that haven't explicitly asked to observe uninitialized objects. \n
                When an object is created, the system will populate this list with
                the current set of initializers. Only privileged users may set or
                modify this list. Once it is empty, it may not be modified further
                by any user. \n DEPRECATED - initializers are an alpha field and will
                be removed in v1.15."
              properties:
                pending:
                  description: Pending is a list of initializers that must execute
                    in order before this object is visible. When the last pending
                    initializer is removed, and no failing result is set, the initializers
                    struct will be set to nil and the object is considered as initialized
                    and visible to all clients.
                  items:
                    properties:
                      name:
                        description: name of the process that is responsible for initializing
                          this object.
                        type: string
                    required:
                    - name
                    type: object
                  type: array
                result:
                  description: If result is set with the Failure field, the object
                    will be persisted to storage and then deleted, ensuring that other
                    clients can observe the deletion.
                  properties:
                    apiVersion:
                      description: 'APIVersion defines the versioned schema of this
                        representation of an object. Servers should convert recognized
                        schemas to the latest internal value, and may reject unrecognized
                        values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
                      type: string
                    code:
                      description: Suggested HTTP return code for this status, 0 if
                        not set.
                      format: int32
                      type: integer
                    details:
                      description: Extended data associated with the reason.  Each
                        reason may define its own extended details. This field is
                        optional and the data returned is not guaranteed to conform
                        to any schema except that defined by the reason type.
                      properties:
                        causes:
                          description: The Causes array includes more details associated
                            with the StatusReason failure. Not all StatusReasons may
                            provide detailed causes.
                          items:
                            properties:
                              field:
                                description: "The field of the resource that has caused
                                  this error, as named by its JSON serialization.
                                  May include dot and postfix notation for nested
                                  attributes. Arrays are zero-indexed.  Fields may
                                  appear more than once in an array of causes due
                                  to fields having multiple errors. Optional. \n Examples:
                                  \  \"name\" - the field \"name\" on the current
                                  resource   \"items[0].name\" - the field \"name\"
                                  on the first array entry in \"items\""
                                type: string
                              message:
                                description: A human-readable description of the cause
                                  of the error.  This field may be presented as-is
                                  to a reader.
                                type: string
                              reason:
                                description: A machine-readable description of the
                                  cause of the error. If this value is empty there
                                  is no information available.
                                type: string
                            type: object
                          type: array
                        group:
                          description: The group attribute of the resource associated
                            with the status StatusReason.
                          type: string
                        kind:
                          description: 'The kind attribute of the resource associated
                            with the status StatusReason. On some operations may differ
                            from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: The name attribute of the resource associated
                            with the status StatusReason (when there is a single name
                            which can be described).
                          type: string
                        retryAfterSeconds:
                          description: If specified, the time in seconds before the
                            operation should be retried. Some errors may indicate
                            the client must take an alternate action - for those errors
                            this field may indicate how long to wait before taking
                            the alternate action.
                          format: int32
                          type: integer
                        uid:
                          description: 'UID of the resource. (when there is a single
                            resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                          type: string
                      type: object
                    kind:
                      description: 'Kind is a string value representing the REST resource
                        this object represents. Servers may infer this from the endpoint
                        the client submits requests to. Cannot be updated. In CamelCase.
                        More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      type: string
                    message:
                      description: A human-readable description of the status of this
                        operation.
                      type: string
                    metadata:
                      description: 'Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      properties:
                        continue:
                          description: continue may be set if the user set a limit
                            on the number of items returned, and indicates that the
                            server has more data available. The value is opaque and
                            may be used to issue another request to the endpoint that
                            served this list to retrieve the next set of available
                            objects. Continuing a consistent list may not be possible
                            if the server configuration has changed or more than a
                            few minutes have passed. The resourceVersion field returned
                            when using this continue value will be identical to the
                            value in the first response, unless you have received
                            this token from an error message.
                          type: string
                        resourceVersion:
                          description: 'String that identifies the server''s internal
                            version of this object that can be used by clients to
                            determine when objects have changed. Value must be treated
                            as opaque by clients and passed unmodified back to the
                            server. Populated by the system. Read-only. More info:
                            https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        selfLink:
                          description: selfLink is a URL representing this object.
                            Populated by the system. Read-only.
                          type: string
                      type: object
                    reason:
                      description: A machine-readable description of why this operation
                        is in the "Failure" status. If this value is empty there is
                        no information available. A Reason clarifies an HTTP status
                        code but does not override it.
                      type: string
                    status:
                      description: 'Status of the operation. One of: "Success" or
                        "Failure". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status'
                      type: string
                  type: object
              required:
              - pending
              type: object
            labels:
              additionalProperties:
                type: string
              description: 'Map of string keys and values that can be used to organize
                and categorize (scope and select) objects. May match selectors of
                replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels'
              type: object
            managedFields:
              description: "ManagedFields maps workflow-id and version to the set
                of fields that are managed by that workflow. This is mostly for internal
                housekeeping, and users typically shouldn't need to set or understand
                this field. A workflow can be the user's name, a controller's name,
                or the name of a specific apply path like \"ci-cd\". The set of fields
                is always in the version that the workflow used when modifying the
                object. \n This field is alpha and can be changed or removed without
                notice."
              items:
                properties:
                  apiVersion:
                    description: APIVersion defines the version of this resource that
                      this field set applies to. The format is "group/version" just
                      like the top-level APIVersion field. It is necessary to track
                      the version of a field set because it cannot be automatically
                      converted.
                    type: string
                  fields:
                    additionalProperties: true
                    description: Fields identifies a set of fields.
                    type: object
                  manager:
                    description: Manager is an identifier of the workflow managing
                      these fields.
                    type: string
                  operation:
                    description: Operation is the type of operation which lead to
                      this ManagedFieldsEntry being created. The only valid values
                      for this field are 'Apply' and 'Update'.
                    type: string
                  time:
                    description: Time is timestamp of when these fields were set.
                      It should always be empty if Operation is 'Apply'
                    format: date-time
                    type: string
                type: object
              type: array
            name:
              description: 'Name must be unique within a namespace. Is required when
                creating resources, although some resources may allow a client to
                request the generation of an appropriate name automatically. Name
                is primarily intended for creation idempotence and configuration definition.
                Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
              type: string
            namespace:
              description: "Namespace defines the space within each name must be unique.
                An empty namespace is equivalent to the \"default\" namespace, but
                \"default\" is the canonical representation. Not all objects are required
                to be scoped to a namespace - the value of this field for those objects
                will be empty. \n Must be a DNS_LABEL. Cannot be updated. More info:
                http://kubernetes.io/docs/user-guide/namespaces"
              type: string
            ownerReferences:
              description: List of objects depended by this object. If ALL objects
                in the list have been deleted, this object will be garbage collected.
                If this object is managed by a controller, then an entry in this list
                will point to this controller, with the controller field set to true.
                There cannot be more than one managing controller.
              items:
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  blockOwnerDeletion:
                    description: If true, AND if the owner has the "foregroundDeletion"
                      finalizer, then the owner cannot be deleted from the key-value
                      store until this reference is removed. Defaults to false. To
                      set this field, a user needs "delete" permission of the owner,
                      otherwise 422 (Unprocessable Entity) will be returned.
                    type: boolean
                  controller:
                    description: If true, this reference points to the managing controller.
                    type: boolean
                  kind:
                    description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                    type: string
                  uid:
                    description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                    type: string
                required:
                - apiVersion
                - kind
                - name
                - uid
                type: object
              type: array
            resourceVersion:
              description: "An opaque value that represents the internal version of
                this object that can be used by clients to determine when objects
                have changed. May be used for optimistic concurrency, change detection,
                and the watch operation on a resource or set of resources. Clients
                must treat these values as opaque and passed unmodified back to the
                server. They may only be valid for a particular resource or set of
                resources. \n Populated by the system. Read-only. Value must be treated
                as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency"
              type: string
            selfLink:
              description: SelfLink is a URL representing this object. Populated by
                the system. Read-only.
              type: string
            uid:
              description: "UID is the unique in time and space value for this object.
                It is typically generated by the server on successful creation of
                a resource and is not allowed to change on PUT operations. \n Populated
                by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids"
              type: string
          type: object
        spec:
          properties:
            clusterNetwork:
              description: Cluster network configuration
              properties:
                pods:
                  description: The network ranges from which Pod networks are allocated.
                  properties:
                    cidrBlocks:
                      items:
                        type: string
                      type: array
                  required:
                  - cidrBlocks
                  type: object
                serviceDomain:
                  description: Domain name for services.
                  type: string
                services:
                  description: The network ranges from which service VIPs are allocated.
                  properties:
                    cidrBlocks:
                      items:
                        type: string
                      type: array
                  required:
                  - cidrBlocks
                  type: object
              required:
              - services
              - pods
              - serviceDomain
              type: object
            providerSpec:
              description: Provider-specific serialized configuration to use during
                cluster creation. It is recommended that providers maintain their
                own versioned API types that should be serialized/deserialized from
                this field.
              properties:
                value:
                  description: Value is an inlined, serialized representation of the
                    resource configuration. It is recommended that providers maintain
                    their own versioned API types that should be serialized/deserialized
                    from this field, akin to component config.
                  type: object
                valueFrom:
                  description: Source for the provider configuration. Cannot be used
                    if value is not empty.
                  properties:
                    machineClass:
                      description: The machine class from which the provider config
                        should be sourced.
                      properties:
                        apiVersion:
                          description: API version of the referent.
                          type: string
                        fieldPath:
                          description: 'If referring to a piece of an object instead
                            of an entire object, this string should contain a valid
                            JSON/Go field access statement, such as desiredState.manifest.containers[2].
                            For example, if the object reference is to a container
                            within a pod, this would take on a value like: "spec.containers{name}"
                            (where "name" refers to the name of the container that
                            triggered the event) or if no container name is specified
                            "spec.containers[2]" (container with index 2 in this pod).
                            This syntax is chosen only to have some well-defined way
                            of referencing a part of an object. TODO: this design
                            is not final and this field is subject to change in the
                            future.'
                          type: string
                        kind:
                          description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                          type: string
                        namespace:
                          description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                          type: string
                        provider:
                          description: Provider is the name of the cloud-provider
                            which MachineClass is intended for.
                          type: string
                        resourceVersion:
                          description: 'Specific resourceVersion to which this reference
                            is made, if any. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        uid:
                          description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                          type: string
                      type: object
                  type: object
              type: object
          required:
          - clusterNetwork
          type: object
        status:
          properties:
            apiEndpoints:
              description: APIEndpoint represents the endpoint to communicate with
                the IP.
              items:
                properties:
                  host:
                    description: The hostname on which the API server is serving.
                    type: string
                  port:
                    description: The port on which the API server is serving.
                    type: integer
                required:
                - host
                - port
                type: object
              type: array
            errorMessage:
              description: If set, indicates that there is a problem reconciling the
                state, and will be set to a descriptive error message.
              type: string
            errorReason:
              description: If set, indicates that there is a problem reconciling the
                state, and will be set to a token value suitable for programmatic
                interpretation.
              type: string
            providerStatus:
              description: Provider-specific status. It is recommended that providers
                maintain their own versioned API types that should be serialized/deserialized
                from this field.
              type: object
          type: object
      type: object
  versions:
  - name: v1alpha1
    served: true
    storage: true
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
  name: machineclasses.cluster.k8s.io
spec:
  group: cluster.k8s.io
  names:
    kind: MachineClass
    plural: machineclasses
    shortNames:
    - mc
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: / [MachineClass] MachineClass can be used to templatize and re-use
        provider configuration across multiple Machines / MachineSets / MachineDeployments.
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        metadata:
          properties:
            annotations:
              additionalProperties:
                type: string
              description: 'Annotations is an unstructured key value map stored with
                a resource that may be set by external tools to store and retrieve
                arbitrary metadata. They are not queryable and should be preserved
                when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations'
              type: object
            clusterName:
              description: The name of the cluster which the object belongs to. This
                is used to distinguish resources with same name and namespace in different
                clusters. This field is not set anywhere right now and apiserver is
                going to ignore it if set in create or update request.
              type: string
            creationTimestamp:
              description: "CreationTimestamp is a timestamp representing the server
                time when this object was created. It is not guaranteed to be set
                in happens-before order across separate operations. Clients may not
                set this value. It is represented in RFC3339 form and is in UTC. \n
                Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            deletionGracePeriodSeconds:
              description: Number of seconds allowed for this object to gracefully
                terminate before it will be removed from the system. Only set when
                deletionTimestamp is also set. May only be shortened. Read-only.
              format: int64
              type: integer
            deletionTimestamp:
              description: "DeletionTimestamp is RFC 3339 date and time at which this
                resource will be deleted. This field is set by the server when a graceful
                deletion is requested by the user, and is not directly settable by
                a client. The resource is expected to be deleted (no longer visible
                from resource lists, and not reachable by name) after the time in
                this field, once the finalizers list is empty. As long as the finalizers
                list contains items, deletion is blocked. Once the deletionTimestamp
                is set, this value may not be unset or be set further into the future,
                although it may be shortened or the resource may be deleted prior
                to this time. For example, a user may request that a pod is deleted
                in 30 seconds. The Kubelet will react by sending a graceful termination
                signal to the containers in the pod. After that 30 seconds, the Kubelet
                will send a hard termination signal (SIGKILL) to the container and
                after cleanup, remove the pod from the API. In the presence of network
                partitions, this object may still exist after this timestamp, until
                an administrator or automated process can determine the resource is
                fully terminated. If not set, graceful deletion of the object has
                not been requested. \n Populated by the system when a graceful deletion
                is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            finalizers:
              description: Must be empty before the object is deleted from the registry.
                Each entry is an identifier for the responsible component that will
                remove the entry from the list. If the deletionTimestamp of the object
                is non-nil, entries in this list can only be removed.
              items:
                type: string
              type: array
            generateName:
              description: "GenerateName is an optional prefix, used by the server,
                to generate a unique name ONLY IF the Name field has not been provided.
                If this field is used, the name returned to the client will be different
                than the name passed. This value will also be combined with a unique
                suffix. The provided value has the same validation rules as the Name
                field, and may be truncated by the length of the suffix required to
                make the value unique on the server. \n If this field is specified
                and the generated name exists, the server will NOT return a 409 -
                instead, it will either return 201 Created or 500 with Reason ServerTimeout
                indicating a unique name could not be found in the time allotted,
                and the client should retry (optionally after the time indicated in
                the Retry-After header). \n Applied only if Name is not specified.
                More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
              type: string
            generation:
              description: A sequence number representing a specific generation of
                the desired state. Populated by the system. Read-only.
              format: int64
              type: integer
            initializers:
              description: "An initializer is a controller which enforces some system
                invariant at object creation time. This field is a list of initializers
                that have not yet acted on this object. If nil or empty, this object
                has been completely initialized. Otherwise, the object is considered
                uninitialized and is hidden (in list/watch and get calls) from clients
                that haven't explicitly asked to observe uninitialized objects. \n
                When an object is created, the system will populate this list with
                the current set of initializers. Only privileged users may set or
                modify this list. Once it is empty, it may not be modified further
                by any user. \n DEPRECATED - initializers are an alpha field and will
                be removed in v1.15."
              properties:
                pending:
                  description: Pending is a list of initializers that must execute
                    in order before this object is visible. When the last pending
                    initializer is removed, and no failing result is set, the initializers
                    struct will be set to nil and the object is considered as initialized
                    and visible to all clients.
                  items:
                    properties:
                      name:
                        description: name of the process that is responsible for initializing
                          this object.
                        type: string
                    required:
                    - name
                    type: object
                  type: array
                result:
                  description: If result is set with the Failure field, the object
                    will be persisted to storage and then deleted, ensuring that other
                    clients can observe the deletion.
                  properties:
                    apiVersion:
                      description: 'APIVersion defines the versioned schema of this
                        representation of an object. Servers should convert recognized
                        schemas to the latest internal value, and may reject unrecognized
                        values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
                      type: string
                    code:
                      description: Suggested HTTP return code for this status, 0 if
                        not set.
                      format: int32
                      type: integer
                    details:
                      description: Extended data associated with the reason.  Each
                        reason may define its own extended details. This field is
                        optional and the data returned is not guaranteed to conform
                        to any schema except that defined by the reason type.
                      properties:
                        causes:
                          description: The Causes array includes more details associated
                            with the StatusReason failure. Not all StatusReasons may
                            provide detailed causes.
                          items:
                            properties:
                              field:
                                description: "The field of the resource that has caused
                                  this error, as named by its JSON serialization.
                                  May include dot and postfix notation for nested
                                  attributes. Arrays are zero-indexed.  Fields may
                                  appear more than once in an array of causes due
                                  to fields having multiple errors. Optional. \n Examples:
                                  \  \"name\" - the field \"name\" on the current
                                  resource   \"items[0].name\" - the field \"name\"
                                  on the first array entry in \"items\""
                                type: string
                              message:
                                description: A human-readable description of the cause
                                  of the error.  This field may be presented as-is
                                  to a reader.
                                type: string
                              reason:
                                description: A machine-readable description of the
                                  cause of the error. If this value is empty there
                                  is no information available.
                                type: string
                            type: object
                          type: array
                        group:
                          description: The group attribute of the resource associated
                            with the status StatusReason.
                          type: string
                        kind:
                          description: 'The kind attribute of the resource associated
                            with the status StatusReason. On some operations may differ
                            from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: The name attribute of the resource associated
                            with the status StatusReason (when there is a single name
                            which can be described).
                          type: string
                        retryAfterSeconds:
                          description: If specified, the time in seconds before the
                            operation should be retried. Some errors may indicate
                            the client must take an alternate action - for those errors
                            this field may indicate how long to wait before taking
                            the alternate action.
                          format: int32
                          type: integer
                        uid:
                          description: 'UID of the resource. (when there is a single
                            resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                          type: string
                      type: object
                    kind:
                      description: 'Kind is a string value representing the REST resource
                        this object represents. Servers may infer this from the endpoint
                        the client submits requests to. Cannot be updated. In CamelCase.
                        More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      type: string
                    message:
                      description: A human-readable description of the status of this
                        operation.
                      type: string
                    metadata:
                      description: 'Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      properties:
                        continue:
                          description: continue may be set if the user set a limit
                            on the number of items returned, and indicates that the
                            server has more data available. The value is opaque and
                            may be used to issue another request to the endpoint that
                            served this list to retrieve the next set of available
                            objects. Continuing a consistent list may not be possible
                            if the server configuration has changed or more than a
                            few minutes have passed. The resourceVersion field returned
                            when using this continue value will be identical to the
                            value in the first response, unless you have received
                            this token from an error message.
                          type: string
                        resourceVersion:
                          description: 'String that identifies the server''s internal
                            version of this object that can be used by clients to
                            determine when objects have changed. Value must be treated
                            as opaque by clients and passed unmodified back to the
                            server. Populated by the system. Read-only. More info:
                            https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        selfLink:
                          description: selfLink is a URL representing this object.
                            Populated by the system. Read-only.
                          type: string
                      type: object
                    reason:
                      description: A machine-readable description of why this operation
                        is in the "Failure" status. If this value is empty there is
                        no information available. A Reason clarifies an HTTP status
                        code but does not override it.
                      type: string
                    status:
                      description: 'Status of the operation. One of: "Success" or
                        "Failure". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status'
                      type: string
                  type: object
              required:
              - pending
              type: object
            labels:
              additionalProperties:
                type: string
              description: 'Map of string keys and values that can be used to organize
                and categorize (scope and select) objects. May match selectors of
                replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels'
              type: object
            managedFields:
              description: "ManagedFields maps workflow-id and version to the set
                of fields that are managed by that workflow. This is mostly for internal
                housekeeping, and users typically shouldn't need to set or understand
                this field. A workflow can be the user's name, a controller's name,
                or the name of a specific apply path like \"ci-cd\". The set of fields
                is always in the version that the workflow used when modifying the
                object. \n This field is alpha and can be changed or removed without
                notice."
              items:
                properties:
                  apiVersion:
                    description: APIVersion defines the version of this resource that
                      this field set applies to. The format is "group/version" just
                      like the top-level APIVersion field. It is necessary to track
                      the version of a field set because it cannot be automatically
                      converted.
                    type: string
                  fields:
                    additionalProperties: true
                    description: Fields identifies a set of fields.
                    type: object
                  manager:
                    description: Manager is an identifier of the workflow managing
                      these fields.
                    type: string
                  operation:
                    description: Operation is the type of operation which lead to
                      this ManagedFieldsEntry being created. The only valid values
                      for this field are 'Apply' and 'Update'.
                    type: string
                  time:
                    description: Time is timestamp of when these fields were set.
                      It should always be empty if Operation is 'Apply'
                    format: date-time
                    type: string
                type: object
              type: array
            name:
              description: 'Name must be unique within a namespace. Is required when
                creating resources, although some resources may allow a client to
                request the generation of an appropriate name automatically. Name
                is primarily intended for creation idempotence and configuration definition.
                Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
              type: string
            namespace:
              description: "Namespace defines the space within each name must be unique.
                An empty namespace is equivalent to the \"default\" namespace, but
                \"default\" is the canonical representation. Not all objects are required
                to be scoped to a namespace - the value of this field for those objects
                will be empty. \n Must be a DNS_LABEL. Cannot be updated. More info:
                http://kubernetes.io/docs/user-guide/namespaces"
              type: string
            ownerReferences:
              description: List of objects depended by this object. If ALL objects
                in the list have been deleted, this object will be garbage collected.
                If this object is managed by a controller, then an entry in this list
                will point to this controller, with the controller field set to true.
                There cannot be more than one managing controller.
              items:
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  blockOwnerDeletion:
                    description: If true, AND if the owner has the "foregroundDeletion"
                      finalizer, then the owner cannot be deleted from the key-value
                      store until this reference is removed. Defaults to false. To
                      set this field, a user needs "delete" permission of the owner,
                      otherwise 422 (Unprocessable Entity) will be returned.
                    type: boolean
                  controller:
                    description: If true, this reference points to the managing controller.
                    type: boolean
                  kind:
                    description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                    type: string
                  uid:
                    description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                    type: string
                required:
                - apiVersion
                - kind
                - name
                - uid
                type: object
              type: array
            resourceVersion:
              description: "An opaque value that represents the internal version of
                this object that can be used by clients to determine when objects
                have changed. May be used for optimistic concurrency, change detection,
                and the watch operation on a resource or set of resources. Clients
                must treat these values as opaque and passed unmodified back to the
                server. They may only be valid for a particular resource or set of
                resources. \n Populated by the system. Read-only. Value must be treated
                as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency"
              type: string
            selfLink:
              description: SelfLink is a URL representing this object. Populated by
                the system. Read-only.
              type: string
            uid:
              description: "UID is the unique in time and space value for this object.
                It is typically generated by the server on successful creation of
                a resource and is not allowed to change on PUT operations. \n Populated
                by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids"
              type: string
          type: object
        providerSpec:
          description: Provider-specific configuration to use during node creation.
          type: object
      required:
      - providerSpec
      type: object
  versions:
  - name: v1alpha1
    served: true
    storage: true
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
  name: machinedeployments.cluster.k8s.io
spec:
  group: cluster.k8s.io
  names:
    kind: MachineDeployment
    plural: machinedeployments
    shortNames:
    - md
  scope: Namespaced
  subresources:
    scale:
      labelSelectorPath: .status.labelSelector
      specReplicasPath: .spec.replicas
      statusReplicasPath: .status.replicas
    status: {}
  validation:
    openAPIV3Schema:
      description: / [MachineDeployment] MachineDeployment is the Schema for the machinedeployments
        API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        metadata:
          properties:
            annotations:
              additionalProperties:
                type: string
              description: 'Annotations is an unstructured key value map stored with
                a resource that may be set by external tools to store and retrieve
                arbitrary metadata. They are not queryable and should be preserved
                when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations'
              type: object
            clusterName:
              description: The name of the cluster which the object belongs to. This
                is used to distinguish resources with same name and namespace in different
                clusters. This field is not set anywhere right now and apiserver is
                going to ignore it if set in create or update request.
              type: string
            creationTimestamp:
              description: "CreationTimestamp is a timestamp representing the server
                time when this object was created. It is not guaranteed to be set
                in happens-before order across separate operations. Clients may not
                set this value. It is represented in RFC3339 form and is in UTC. \n
                Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            deletionGracePeriodSeconds:
              description: Number of seconds allowed for this object to gracefully
                terminate before it will be removed from the system. Only set when
                deletionTimestamp is also set. May only be shortened. Read-only.
              format: int64
              type: integer
            deletionTimestamp:
              description: "DeletionTimestamp is RFC 3339 date and time at which this
                resource will be deleted. This field is set by the server when a graceful
                deletion is requested by the user, and is not directly settable by
                a client. The resource is expected to be deleted (no longer visible
                from resource lists, and not reachable by name) after the time in
                this field, once the finalizers list is empty. As long as the finalizers
                list contains items, deletion is blocked. Once the deletionTimestamp
                is set, this value may not be unset or be set further into the future,
                although it may be shortened or the resource may be deleted prior
                to this time. For example, a user may request that a pod is deleted
                in 30 seconds. The Kubelet will react by sending a graceful termination
                signal to the containers in the pod. After that 30 seconds, the Kubelet
                will send a hard termination signal (SIGKILL) to the container and
                after cleanup, remove the pod from the API. In the presence of network
                partitions, this object may still exist after this timestamp, until
                an administrator or automated process can determine the resource is
                fully terminated. If not set, graceful deletion of the object has
                not been requested. \n Populated by the system when a graceful deletion
                is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            finalizers:
              description: Must be empty before the object is deleted from the registry.
                Each entry is an identifier for the responsible component that will
                remove the entry from the list. If the deletionTimestamp of the object
                is non-nil, entries in this list can only be removed.
              items:
                type: string
              type: array
            generateName:
              description: "GenerateName is an optional prefix, used by the server,
                to generate a unique name ONLY IF the Name field has not been provided.
                If this field is used, the name returned to the client will be different
                than the name passed. This value will also be combined with a unique
                suffix. The provided value has the same validation rules as the Name
                field, and may be truncated by the length of the suffix required to
                make the value unique on the server. \n If this field is specified
                and the generated name exists, the server will NOT return a 409 -
                instead, it will either return 201 Created or 500 with Reason ServerTimeout
                indicating a unique name could not be found in the time allotted,
                and the client should retry (optionally after the time indicated in
                the Retry-After header). \n Applied only if Name is not specified.
                More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
              type: string
            generation:
              description: A sequence number representing a specific generation of
                the desired state. Populated by the system. Read-only.
              format: int64
              type: integer
            initializers:
              description: "An initializer is a controller which enforces some system
                invariant at object creation time. This field is a list of initializers
                that have not yet acted on this object. If nil or empty, this object
                has been completely initialized. Otherwise, the object is considered
                uninitialized and is hidden (in list/watch and get calls) from clients
                that haven't explicitly asked to observe uninitialized objects. \n
                When an object is created, the system will populate this list with
                the current set of initializers. Only privileged users may set or
                modify this list. Once it is empty, it may not be modified further
                by any user. \n DEPRECATED - initializers are an alpha field and will
                be removed in v1.15."
              properties:
                pending:
                  description: Pending is a list of initializers that must execute
                    in order before this object is visible. When the last pending
                    initializer is removed, and no failing result is set, the initializers
                    struct will be set to nil and the object is considered as initialized
                    and visible to all clients.
                  items:
                    properties:
                      name:
                        description: name of the process that is responsible for initializing
                          this object.
                        type: string
                    required:
                    - name
                    type: object
                  type: array
                result:
                  description: If result is set with the Failure field, the object
                    will be persisted to storage and then deleted, ensuring that other
                    clients can observe the deletion.
                  properties:
                    apiVersion:
                      description: 'APIVersion defines the versioned schema of this
                        representation of an object. Servers should convert recognized
                        schemas to the latest internal value, and may reject unrecognized
                        values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
                      type: string
                    code:
                      description: Suggested HTTP return code for this status, 0 if
                        not set.
                      format: int32
                      type: integer
                    details:
                      description: Extended data associated with the reason.  Each
                        reason may define its own extended details. This field is
                        optional and the data returned is not guaranteed to conform
                        to any schema except that defined by the reason type.
                      properties:
                        causes:
                          description: The Causes array includes more details associated
                            with the StatusReason failure. Not all StatusReasons may
                            provide detailed causes.
                          items:
                            properties:
                              field:
                                description: "The field of the resource that has caused
                                  this error, as named by its JSON serialization.
                                  May include dot and postfix notation for nested
                                  attributes. Arrays are zero-indexed.  Fields may
                                  appear more than once in an array of causes due
                                  to fields having multiple errors. Optional. \n Examples:
                                  \  \"name\" - the field \"name\" on the current
                                  resource   \"items[0].name\" - the field \"name\"
                                  on the first array entry in \"items\""
                                type: string
                              message:
                                description: A human-readable description of the cause
                                  of the error.  This field may be presented as-is
                                  to a reader.
                                type: string
                              reason:
                                description: A machine-readable description of the
                                  cause of the error. If this value is empty there
                                  is no information available.
                                type: string
                            type: object
                          type: array
                        group:
                          description: The group attribute of the resource associated
                            with the status StatusReason.
                          type: string
                        kind:
                          description: 'The kind attribute of the resource associated
                            with the status StatusReason. On some operations may differ
                            from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: The name attribute of the resource associated
                            with the status StatusReason (when there is a single name
                            which can be described).
                          type: string
                        retryAfterSeconds:
                          description: If specified, the time in seconds before the
                            operation should be retried. Some errors may indicate
                            the client must take an alternate action - for those errors
                            this field may indicate how long to wait before taking
                            the alternate action.
                          format: int32
                          type: integer
                        uid:
                          description: 'UID of the resource. (when there is a single
                            resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                          type: string
                      type: object
                    kind:
                      description: 'Kind is a string value representing the REST resource
                        this object represents. Servers may infer this from the endpoint
                        the client submits requests to. Cannot be updated. In CamelCase.
                        More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      type: string
                    message:
                      description: A human-readable description of the status of this
                        operation.
                      type: string
                    metadata:
                      description: 'Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      properties:
                        continue:
                          description: continue may be set if the user set a limit
                            on the number of items returned, and indicates that the
                            server has more data available. The value is opaque and
                            may be used to issue another request to the endpoint that
                            served this list to retrieve the next set of available
                            objects. Continuing a consistent list may not be possible
                            if the server configuration has changed or more than a
                            few minutes have passed. The resourceVersion field returned
                            when using this continue value will be identical to the
                            value in the first response, unless you have received
                            this token from an error message.
                          type: string
                        resourceVersion:
                          description: 'String that identifies the server''s internal
                            version of this object that can be used by clients to
                            determine when objects have changed. Value must be treated
                            as opaque by clients and passed unmodified back to the
                            server. Populated by the system. Read-only. More info:
                            https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        selfLink:
                          description: selfLink is a URL representing this object.
                            Populated by the system. Read-only.
                          type: string
                      type: object
                    reason:
                      description: A machine-readable description of why this operation
                        is in the "Failure" status. If this value is empty there is
                        no information available. A Reason clarifies an HTTP status
                        code but does not override it.
                      type: string
                    status:
                      description: 'Status of the operation. One of: "Success" or
                        "Failure". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status'
                      type: string
                  type: object
              required:
              - pending
              type: object
            labels:
              additionalProperties:
                type: string
              description: 'Map of string keys and values that can be used to organize
                and categorize (scope and select) objects. May match selectors of
                replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels'
              type: object
            managedFields:
              description: "ManagedFields maps workflow-id and version to the set
                of fields that are managed by that workflow. This is mostly for internal
                housekeeping, and users typically shouldn't need to set or understand
                this field. A workflow can be the user's name, a controller's name,
                or the name of a specific apply path like \"ci-cd\". The set of fields
                is always in the version that the workflow used when modifying the
                object. \n This field is alpha and can be changed or removed without
                notice."
              items:
                properties:
                  apiVersion:
                    description: APIVersion defines the version of this resource that
                      this field set applies to. The format is "group/version" just
                      like the top-level APIVersion field. It is necessary to track
                      the version of a field set because it cannot be automatically
                      converted.
                    type: string
                  fields:
                    additionalProperties: true
                    description: Fields identifies a set of fields.
                    type: object
                  manager:
                    description: Manager is an identifier of the workflow managing
                      these fields.
                    type: string
                  operation:
                    description: Operation is the type of operation which lead to
                      this ManagedFieldsEntry being created. The only valid values
                      for this field are 'Apply' and 'Update'.
                    type: string
                  time:
                    description: Time is timestamp of when these fields were set.
                      It should always be empty if Operation is 'Apply'
                    format: date-time
                    type: string
                type: object
              type: array
            name:
              description: 'Name must be unique within a namespace. Is required when
                creating resources, although some resources may allow a client to
                request the generation of an appropriate name automatically. Name
                is primarily intended for creation idempotence and configuration definition.
                Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
              type: string
            namespace:
              description: "Namespace defines the space within each name must be unique.
                An empty namespace is equivalent to the \"default\" namespace, but
                \"default\" is the canonical representation. Not all objects are required
                to be scoped to a namespace - the value of this field for those objects
                will be empty. \n Must be a DNS_LABEL. Cannot be updated. More info:
                http://kubernetes.io/docs/user-guide/namespaces"
              type: string
            ownerReferences:
              description: List of objects depended by this object. If ALL objects
                in the list have been deleted, this object will be garbage collected.
                If this object is managed by a controller, then an entry in this list
                will point to this controller, with the controller field set to true.
                There cannot be more than one managing controller.
              items:
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  blockOwnerDeletion:
                    description: If true, AND if the owner has the "foregroundDeletion"
                      finalizer, then the owner cannot be deleted from the key-value
                      store until this reference is removed. Defaults to false. To
                      set this field, a user needs "delete" permission of the owner,
                      otherwise 422 (Unprocessable Entity) will be returned.
                    type: boolean
                  controller:
                    description: If true, this reference points to the managing controller.
                    type: boolean
                  kind:
                    description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                    type: string
                  uid:
                    description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                    type: string
                required:
                - apiVersion
                - kind
                - name
                - uid
                type: object
              type: array
            resourceVersion:
              description: "An opaque value that represents the internal version of
                this object that can be used by clients to determine when objects
                have changed. May be used for optimistic concurrency, change detection,
                and the watch operation on a resource or set of resources. Clients
                must treat these values as opaque and passed unmodified back to the
                server. They may only be valid for a particular resource or set of
                resources. \n Populated by the system. Read-only. Value must be treated
                as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency"
              type: string
            selfLink:
              description: SelfLink is a URL representing this object. Populated by
                the system. Read-only.
              type: string
            uid:
              description: "UID is the unique in time and space value for this object.
                It is typically generated by the server on successful creation of
                a resource and is not allowed to change on PUT operations. \n Populated
                by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids"
              type: string
          type: object
        spec:
          properties:
            minReadySeconds:
              description: Minimum number of seconds for which a newly created machine
                should be ready. Defaults to 0 (machine will be considered available
                as soon as it is ready)
              format: int32
              type: integer
            paused:
              description: Indicates that the deployment is paused.
              type: boolean
            progressDeadlineSeconds:
              description: The maximum time in seconds for a deployment to make progress
                before it is considered to be failed. The deployment controller will
                continue to process failed deployments and a condition with a ProgressDeadlineExceeded
                reason will be surfaced in the deployment status. Note that progress
                will not be estimated during the time a deployment is paused. Defaults
                to 600s.
              format: int32
              type: integer
            replicas:
              description: Number of desired machines. Defaults to 1. This is a pointer
                to distinguish between explicit zero and not specified.
              format: int32
              type: integer
            revisionHistoryLimit:
              description: The number of old MachineSets to retain to allow rollback.
                This is a pointer to distinguish between explicit zero and not specified.
                Defaults to 1.
              format: int32
              type: integer
            selector:
              description: Label selector for machines. Existing MachineSets whose
                machines are selected by this will be the ones affected by this deployment.
                It must match the machine template's labels.
              properties:
                matchExpressions:
                  description: matchExpressions is a list of label selector requirements.
                    The requirements are ANDed.
                  items:
                    properties:
                      key:
                        description: key is the label key that the selector applies
                          to.
                        type: string
                      operator:
                        description: operator represents a key's relationship to a
                          set of values. Valid operators are In, NotIn, Exists and
                          DoesNotExist.
                        type: string
                      values:
                        description: values is an array of string values. If the operator
                          is In or NotIn, the values array must be non-empty. If the
                          operator is Exists or DoesNotExist, the values array must
                          be empty. This array is replaced during a strategic merge
                          patch.
                        items:
                          type: string
                        type: array
                    required:
                    - key
                    - operator
                    type: object
                  type: array
                matchLabels:
                  additionalProperties:
                    type: string
                  description: matchLabels is a map of {key,value} pairs. A single
                    {key,value} in the matchLabels map is equivalent to an element
                    of matchExpressions, whose key field is "key", the operator is
                    "In", and the values array contains only "value". The requirements
                    are ANDed.
                  type: object
              type: object
            strategy:
              description: The deployment strategy to use to replace existing machines
                with new ones.
              properties:
                rollingUpdate:
                  description: Rolling update config params. Present only if MachineDeploymentStrategyType
                    = RollingUpdate.
                  properties:
                    maxSurge:
                      anyOf:
                      - type: string
                      - type: integer
                      description: 'The maximum number of machines that can be scheduled
                        above the desired number of machines. Value can be an absolute
                        number (ex: 5) or a percentage of desired machines (ex: 10%).
                        This can not be 0 if MaxUnavailable is 0. Absolute number
                        is calculated from percentage by rounding up. Defaults to
                        1. Example: when this is set to 30%, the new MachineSet can
                        be scaled up immediately when the rolling update starts, such
                        that the total number of old and new machines do not exceed
                        130% of desired machines. Once old machines have been killed,
                        new MachineSet can be scaled up further, ensuring that total
                        number of machines running at any time during the update is
                        at most 130% of desired machines.'
                    maxUnavailable:
                      anyOf:
                      - type: string
                      - type: integer
                      description: 'The maximum number of machines that can be unavailable
                        during the update. Value can be an absolute number (ex: 5)
                        or a percentage of desired machines (ex: 10%). Absolute number
                        is calculated from percentage by rounding down. This can not
                        be 0 if MaxSurge is 0. Defaults to 0. Example: when this is
                        set to 30%, the old MachineSet can be scaled down to 70% of
                        desired machines immediately when the rolling update starts.
                        Once new machines are ready, old MachineSet can be scaled
                        down further, followed by scaling up the new MachineSet, ensuring
                        that the total number of machines available at all times during
                        the update is at least 70% of desired machines.'
                  type: object
                type:
                  description: Type of deployment. Currently the only supported strategy
                    is "RollingUpdate". Default is RollingUpdate.
                  type: string
              type: object
            template:
              description: Template describes the machines that will be created.
              properties:
                metadata:
                  description: 'Standard object''s metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
                  properties:
                    annotations:
                      additionalProperties:
                        type: string
                      description: 'Annotations is an unstructured key value map stored
                        with a resource that may be set by external tools to store
                        and retrieve arbitrary metadata. They are not queryable and
                        should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations'
                      type: object
                    generateName:
                      description: "GenerateName is an optional prefix, used by the
                        server, to generate a unique name ONLY IF the Name field has
                        not been provided. If this field is used, the name returned
                        to the client will be different than the name passed. This
                        value will also be combined with a unique suffix. The provided
                        value has the same validation rules as the Name field, and
                        may be truncated by the length of the suffix required to make
                        the value unique on the server. \n If this field is specified
                        and the generated name exists, the server will NOT return
                        a 409 - instead, it will either return 201 Created or 500
                        with Reason ServerTimeout indicating a unique name could not
                        be found in the time allotted, and the client should retry
                        (optionally after the time indicated in the Retry-After header).
                        \n Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
                      type: string
                    labels:
                      additionalProperties:
                        type: string
                      description: 'Map of string keys and values that can be used
                        to organize and categorize (scope and select) objects. May
                        match selectors of replication controllers and services. More
                        info: http://kubernetes.io/docs/user-guide/labels'
                      type: object
                    name:
                      description: 'Name must be unique within a namespace. Is required
                        when creating resources, although some resources may allow
                        a client to request the generation of an appropriate name
                        automatically. Name is primarily intended for creation idempotence
                        and configuration definition. Cannot be updated. More info:
                        http://kubernetes.io/docs/user-guide/identifiers#names'
                      type: string
                    namespace:
                      description: "Namespace defines the space within each name must
                        be unique. An empty namespace is equivalent to the \"default\"
                        namespace, but \"default\" is the canonical representation.
                        Not all objects are required to be scoped to a namespace -
                        the value of this field for those objects will be empty. \n
                        Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces"
                      type: string
                    ownerReferences:
                      description: List of objects depended by this object. If ALL
                        objects in the list have been deleted, this object will be
                        garbage collected. If this object is managed by a controller,
                        then an entry in this list will point to this controller,
                        with the controller field set to true. There cannot be more
                        than one managing controller.
                      items:
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          blockOwnerDeletion:
                            description: If true, AND if the owner has the "foregroundDeletion"
                              finalizer, then the owner cannot be deleted from the
                              key-value store until this reference is removed. Defaults
                              to false. To set this field, a user needs "delete" permission
                              of the owner, otherwise 422 (Unprocessable Entity) will
                              be returned.
                            type: boolean
                          controller:
                            description: If true, this reference points to the managing
                              controller.
                            type: boolean
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                            type: string
                        required:
                        - apiVersion
                        - kind
                        - name
                        - uid
                        type: object
                      type: array
                  type: object
                spec:
                  description: 'Specification of the desired behavior of the machine.
                    More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status'
                  properties:
                    configSource:
                      description: ConfigSource is used to populate in the associated
                        Node for dynamic kubelet config. This field already exists
                        in Node, so any updates to it in the Machine spec will be
                        automatically copied to the linked NodeRef from the status.
                        The rest of dynamic kubelet config support should then work
                        as-is.
                      properties:
                        configMap:
                          description: ConfigMap is a reference to a Node's ConfigMap
                          properties:
                            kubeletConfigKey:
                              description: KubeletConfigKey declares which key of
                                the referenced ConfigMap corresponds to the KubeletConfiguration
                                structure This field is required in all cases.
                              type: string
                            name:
                              description: Name is the metadata.name of the referenced
                                ConfigMap. This field is required in all cases.
                              type: string
                            namespace:
                              description: Namespace is the metadata.namespace of
                                the referenced ConfigMap. This field is required in
                                all cases.
                              type: string
                            resourceVersion:
                              description: ResourceVersion is the metadata.ResourceVersion
                                of the referenced ConfigMap. This field is forbidden
                                in Node.Spec, and required in Node.Status.
                              type: string
                            uid:
                              description: UID is the metadata.UID of the referenced
                                ConfigMap. This field is forbidden in Node.Spec, and
                                required in Node.Status.
                              type: string
                          required:
                          - namespace
                          - name
                          - kubeletConfigKey
                          type: object
                      type: object
                    metadata:
                      description: ObjectMeta will autopopulate the Node created.
                        Use this to indicate what labels, annotations, name prefix,
                        etc., should be used when creating the Node.
                      properties:
                        annotations:
                          additionalProperties:
                            type: string
                          description: 'Annotations is an unstructured key value map
                            stored with a resource that may be set by external tools
                            to store and retrieve arbitrary metadata. They are not
                            queryable and should be preserved when modifying objects.
                            More info: http://kubernetes.io/docs/user-guide/annotations'
                          type: object
                        generateName:
                          description: "GenerateName is an optional prefix, used by
                            the server, to generate a unique name ONLY IF the Name
                            field has not been provided. If this field is used, the
                            name returned to the client will be different than the
                            name passed. This value will also be combined with a unique
                            suffix. The provided value has the same validation rules
                            as the Name field, and may be truncated by the length
                            of the suffix required to make the value unique on the
                            server. \n If this field is specified and the generated
                            name exists, the server will NOT return a 409 - instead,
                            it will either return 201 Created or 500 with Reason ServerTimeout
                            indicating a unique name could not be found in the time
                            allotted, and the client should retry (optionally after
                            the time indicated in the Retry-After header). \n Applied
                            only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
                          type: string
                        labels:
                          additionalProperties:
                            type: string
                          description: 'Map of string keys and values that can be
                            used to organize and categorize (scope and select) objects.
                            May match selectors of replication controllers and services.
                            More info: http://kubernetes.io/docs/user-guide/labels'
                          type: object
                        name:
                          description: 'Name must be unique within a namespace. Is
                            required when creating resources, although some resources
                            may allow a client to request the generation of an appropriate
                            name automatically. Name is primarily intended for creation
                            idempotence and configuration definition. Cannot be updated.
                            More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                          type: string
                        namespace:
                          description: "Namespace defines the space within each name
                            must be unique. An empty namespace is equivalent to the
                            \"default\" namespace, but \"default\" is the canonical
                            representation. Not all objects are required to be scoped
                            to a namespace - the value of this field for those objects
                            will be empty. \n Must be a DNS_LABEL. Cannot be updated.
                            More info: http://kubernetes.io/docs/user-guide/namespaces"
                          type: string
                        ownerReferences:
                          description: List of objects depended by this object. If
                            ALL objects in the list have been deleted, this object
                            will be garbage collected. If this object is managed by
                            a controller, then an entry in this list will point to
                            this controller, with the controller field set to true.
                            There cannot be more than one managing controller.
                          items:
                            properties:
                              apiVersion:
                                description: API version of the referent.
                                type: string
                              blockOwnerDeletion:
                                description: If true, AND if the owner has the "foregroundDeletion"
                                  finalizer, then the owner cannot be deleted from
                                  the key-value store until this reference is removed.
                                  Defaults to false. To set this field, a user needs
                                  "delete" permission of the owner, otherwise 422
                                  (Unprocessable Entity) will be returned.
                                type: boolean
                              controller:
                                description: If true, this reference points to the
                                  managing controller.
                                type: boolean
                              kind:
                                description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                                type: string
                              name:
                                description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                                type: string
                              uid:
                                description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                                type: string
                            required:
                            - apiVersion
                            - kind
                            - name
                            - uid
                            type: object
                          type: array
                      type: object
                    providerID:
                      description: ProviderID is the identification ID of the machine
                        provided by the provider. This field must match the provider
                        ID as seen on the node object corresponding to this machine.
                        This field is required by higher level consumers of cluster-api.
                        Example use case is cluster autoscaler with cluster-api as
                        provider. Clean-up logic in the autoscaler compares machines
                        to nodes to find out machines at provider which could not
                        get registered as Kubernetes nodes. With cluster-api as a
                        generic out-of-tree provider for autoscaler, this field is
                        required by autoscaler to be able to have a provider view
                        of the list of machines. Another list of nodes is queried
                        from the k8s apiserver and then a comparison is done to find
                        out unregistered machines and are marked for delete. This
                        field will be set by the actuators and consumed by higher
                        level entities like autoscaler that will be interfacing with
                        cluster-api as generic provider.
                      type: string
                    providerSpec:
                      description: ProviderSpec details Provider-specific configuration
                        to use during node creation.
                      properties:
                        value:
                          description: Value is an inlined, serialized representation
                            of the resource configuration. It is recommended that
                            providers maintain their own versioned API types that
                            should be serialized/deserialized from this field, akin
                            to component config.
                          type: object
                        valueFrom:
                          description: Source for the provider configuration. Cannot
                            be used if value is not empty.
                          properties:
                            machineClass:
                              description: The machine class from which the provider
                                config should be sourced.
                              properties:
                                apiVersion:
                                  description: API version of the referent.
                                  type: string
                                fieldPath:
                                  description: 'If referring to a piece of an object
                                    instead of an entire object, this string should
                                    contain a valid JSON/Go field access statement,
                                    such as desiredState.manifest.containers[2]. For
                                    example, if the object reference is to a container
                                    within a pod, this would take on a value like:
                                    "spec.containers{name}" (where "name" refers to
                                    the name of the container that triggered the event)
                                    or if no container name is specified "spec.containers[2]"
                                    (container with index 2 in this pod). This syntax
                                    is chosen only to have some well-defined way of
                                    referencing a part of an object. TODO: this design
                                    is not final and this field is subject to change
                                    in the future.'
                                  type: string
                                kind:
                                  description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                                  type: string
                                name:
                                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                  type: string
                                namespace:
                                  description: 'Namespace of the referent. More info:
                                    https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                  type: string
                                provider:
                                  description: Provider is the name of the cloud-provider
                                    which MachineClass is intended for.
                                  type: string
                                resourceVersion:
                                  description: 'Specific resourceVersion to which
                                    this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                                  type: string
                                uid:
                                  description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                                  type: string
                              type: object
                          type: object
                      type: object
                    taints:
                      description: The list of the taints to be applied to the corresponding
                        Node in additive manner. This list will not overwrite any
                        other taints added to the Node on an ongoing basis by other
                        entities. These taints should be actively reconciled e.g.
                        if you ask the machine controller to apply a taint and then
                        manually remove the taint the machine controller will put
                        it back) but not have the machine controller remove any taints
                      items:
                        properties:
                          effect:
                            description: Required. The effect of the taint on pods
                              that do not tolerate the taint. Valid effects are NoSchedule,
                              PreferNoSchedule and NoExecute.
                            type: string
                          key:
                            description: Required. The taint key to be applied to
                              a node.
                            type: string
                          timeAdded:
                            description: TimeAdded represents the time at which the
                              taint was added. It is only written for NoExecute taints.
                            format: date-time
                            type: string
                          value:
                            description: Required. The taint value corresponding to
                              the taint key.
                            type: string
                        required:
                        - key
                        - effect
                        type: object
                      type: array
                    versions:
                      description: Versions of key software to use. This field is
                        optional at cluster creation time, and omitting the field
                        indicates that the cluster installation tool should select
                        defaults for the user. These defaults may differ based on
                        the cluster installer, but the tool should populate the values
                        it uses when persisting Machine objects. A Machine spec missing
                        this field at runtime is invalid.
                      properties:
                        controlPlane:
                          description: ControlPlane is the semantic version of the
                            Kubernetes control plane to run. This should only be populated
                            when the machine is a control plane.
                          type: string
                        kubelet:
                          description: Kubelet is the semantic version of kubelet
                            to run
                          type: string
                      required:
                      - kubelet
                      type: object
                  required:
                  - providerSpec
                  type: object
              type: object
          required:
          - selector
          - template
          type: object
        status:
          properties:
            availableReplicas:
              description: Total number of available machines (ready for at least
                minReadySeconds) targeted by this deployment.
              format: int32
              type: integer
            observedGeneration:
              description: The generation observed by the deployment controller.
              format: int64
              type: integer
            readyReplicas:
              description: Total number of ready machines targeted by this deployment.
              format: int32
              type: integer
            replicas:
              description: Total number of non-terminated machines targeted by this
                deployment (their labels match the selector).
              format: int32
              type: integer
            unavailableReplicas:
              description: Total number of unavailable machines targeted by this deployment.
                This is the total number of machines that are still required for the
                deployment to have 100% available capacity. They may either be machines
                that are running but not yet available or machines that still have
                not been created.
              format: int32
              type: integer
            updatedReplicas:
              description: Total number of non-terminated machines targeted by this
                deployment that have the desired template spec.
              format: int32
              type: integer
          type: object
      type: object
  versions:
  - name: v1alpha1
    served: true
    storage: true
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
  name: machines.cluster.k8s.io
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.providerID
    description: Provider ID
    name: ProviderID
    type: string
  - JSONPath: .status.phase
    description: Machine status such as Terminating/Pending/Running/Failed etc
    name: Phase
    type: string
  - JSONPath: .status.nodeRef.name
    description: Node name associated with this machine
    name: NodeName
    priority: 1
    type: string
  group: cluster.k8s.io
  names:
    kind: Machine
    plural: machines
    shortNames:
    - ma
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: / [Machine] Machine is the Schema for the machines API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        metadata:
          properties:
            annotations:
              additionalProperties:
                type: string
              description: 'Annotations is an unstructured key value map stored with
                a resource that may be set by external tools to store and retrieve
                arbitrary metadata. They are not queryable and should be preserved
                when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations'
              type: object
            clusterName:
              description: The name of the cluster which the object belongs to. This
                is used to distinguish resources with same name and namespace in different
                clusters. This field is not set anywhere right now and apiserver is
                going to ignore it if set in create or update request.
              type: string
            creationTimestamp:
              description: "CreationTimestamp is a timestamp representing the server
                time when this object was created. It is not guaranteed to be set
                in happens-before order across separate operations. Clients may not
                set this value. It is represented in RFC3339 form and is in UTC. \n
                Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            deletionGracePeriodSeconds:
              description: Number of seconds allowed for this object to gracefully
                terminate before it will be removed from the system. Only set when
                deletionTimestamp is also set. May only be shortened. Read-only.
              format: int64
              type: integer
            deletionTimestamp:
              description: "DeletionTimestamp is RFC 3339 date and time at which this
                resource will be deleted. This field is set by the server when a graceful
                deletion is requested by the user, and is not directly settable by
                a client. The resource is expected to be deleted (no longer visible
                from resource lists, and not reachable by name) after the time in
                this field, once the finalizers list is empty. As long as the finalizers
                list contains items, deletion is blocked. Once the deletionTimestamp
                is set, this value may not be unset or be set further into the future,
                although it may be shortened or the resource may be deleted prior
                to this time. For example, a user may request that a pod is deleted
                in 30 seconds. The Kubelet will react by sending a graceful termination
                signal to the containers in the pod. After that 30 seconds, the Kubelet
                will send a hard termination signal (SIGKILL) to the container and
                after cleanup, remove the pod from the API. In the presence of network
                partitions, this object may still exist after this timestamp, until
                an administrator or automated process can determine the resource is
                fully terminated. If not set, graceful deletion of the object has
                not been requested. \n Populated by the system when a graceful deletion
                is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            finalizers:
              description: Must be empty before the object is deleted from the registry.
                Each entry is an identifier for the responsible component that will
                remove the entry from the list. If the deletionTimestamp of the object
                is non-nil, entries in this list can only be removed.
              items:
                type: string
              type: array
            generateName:
              description: "GenerateName is an optional prefix, used by the server,
                to generate a unique name ONLY IF the Name field has not been provided.
                If this field is used, the name returned to the client will be different
                than the name passed. This value will also be combined with a unique
                suffix. The provided value has the same validation rules as the Name
                field, and may be truncated by the length of the suffix required to
                make the value unique on the server. \n If this field is specified
                and the generated name exists, the server will NOT return a 409 -
                instead, it will either return 201 Created or 500 with Reason ServerTimeout
                indicating a unique name could not be found in the time allotted,
                and the client should retry (optionally after the time indicated in
                the Retry-After header). \n Applied only if Name is not specified.
                More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
              type: string
            generation:
              description: A sequence number representing a specific generation of
                the desired state. Populated by the system. Read-only.
              format: int64
              type: integer
            initializers:
              description: "An initializer is a controller which enforces some system
                invariant at object creation time. This field is a list of initializers
                that have not yet acted on this object. If nil or empty, this object
                has been completely initialized. Otherwise, the object is considered
                uninitialized and is hidden (in list/watch and get calls) from clients
                that haven't explicitly asked to observe uninitialized objects. \n
                When an object is created, the system will populate this list with
                the current set of initializers. Only privileged users may set or
                modify this list. Once it is empty, it may not be modified further
                by any user. \n DEPRECATED - initializers are an alpha field and will
                be removed in v1.15."
              properties:
                pending:
                  description: Pending is a list of initializers that must execute
                    in order before this object is visible. When the last pending
                    initializer is removed, and no failing result is set, the initializers
                    struct will be set to nil and the object is considered as initialized
                    and visible to all clients.
                  items:
                    properties:
                      name:
                        description: name of the process that is responsible for initializing
                          this object.
                        type: string
                    required:
                    - name
                    type: object
                  type: array
                result:
                  description: If result is set with the Failure field, the object
                    will be persisted to storage and then deleted, ensuring that other
                    clients can observe the deletion.
                  properties:
                    apiVersion:
                      description: 'APIVersion defines the versioned schema of this
                        representation of an object. Servers should convert recognized
                        schemas to the latest internal value, and may reject unrecognized
                        values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
                      type: string
                    code:
                      description: Suggested HTTP return code for this status, 0 if
                        not set.
                      format: int32
                      type: integer
                    details:
                      description: Extended data associated with the reason.  Each
                        reason may define its own extended details. This field is
                        optional and the data returned is not guaranteed to conform
                        to any schema except that defined by the reason type.
                      properties:
                        causes:
                          description: The Causes array includes more details associated
                            with the StatusReason failure. Not all StatusReasons may
                            provide detailed causes.
                          items:
                            properties:
                              field:
                                description: "The field of the resource that has caused
                                  this error, as named by its JSON serialization.
                                  May include dot and postfix notation for nested
                                  attributes. Arrays are zero-indexed.  Fields may
                                  appear more than once in an array of causes due
                                  to fields having multiple errors. Optional. \n Examples:
                                  \  \"name\" - the field \"name\" on the current
                                  resource   \"items[0].name\" - the field \"name\"
                                  on the first array entry in \"items\""
                                type: string
                              message:
                                description: A human-readable description of the cause
                                  of the error.  This field may be presented as-is
                                  to a reader.
                                type: string
                              reason:
                                description: A machine-readable description of the
                                  cause of the error. If this value is empty there
                                  is no information available.
                                type: string
                            type: object
                          type: array
                        group:
                          description: The group attribute of the resource associated
                            with the status StatusReason.
                          type: string
                        kind:
                          description: 'The kind attribute of the resource associated
                            with the status StatusReason. On some operations may differ
                            from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: The name attribute of the resource associated
                            with the status StatusReason (when there is a single name
                            which can be described).
                          type: string
                        retryAfterSeconds:
                          description: If specified, the time in seconds before the
                            operation should be retried. Some errors may indicate
                            the client must take an alternate action - for those errors
                            this field may indicate how long to wait before taking
                            the alternate action.
                          format: int32
                          type: integer
                        uid:
                          description: 'UID of the resource. (when there is a single
                            resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                          type: string
                      type: object
                    kind:
                      description: 'Kind is a string value representing the REST resource
                        this object represents. Servers may infer this from the endpoint
                        the client submits requests to. Cannot be updated. In CamelCase.
                        More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      type: string
                    message:
                      description: A human-readable description of the status of this
                        operation.
                      type: string
                    metadata:
                      description: 'Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      properties:
                        continue:
                          description: continue may be set if the user set a limit
                            on the number of items returned, and indicates that the
                            server has more data available. The value is opaque and
                            may be used to issue another request to the endpoint that
                            served this list to retrieve the next set of available
                            objects. Continuing a consistent list may not be possible
                            if the server configuration has changed or more than a
                            few minutes have passed. The resourceVersion field returned
                            when using this continue value will be identical to the
                            value in the first response, unless you have received
                            this token from an error message.
                          type: string
                        resourceVersion:
                          description: 'String that identifies the server''s internal
                            version of this object that can be used by clients to
                            determine when objects have changed. Value must be treated
                            as opaque by clients and passed unmodified back to the
                            server. Populated by the system. Read-only. More info:
                            https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        selfLink:
                          description: selfLink is a URL representing this object.
                            Populated by the system. Read-only.
                          type: string
                      type: object
                    reason:
                      description: A machine-readable description of why this operation
                        is in the "Failure" status. If this value is empty there is
                        no information available. A Reason clarifies an HTTP status
                        code but does not override it.
                      type: string
                    status:
                      description: 'Status of the operation. One of: "Success" or
                        "Failure". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status'
                      type: string
                  type: object
              required:
              - pending
              type: object
            labels:
              additionalProperties:
                type: string
              description: 'Map of string keys and values that can be used to organize
                and categorize (scope and select) objects. May match selectors of
                replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels'
              type: object
            managedFields:
              description: "ManagedFields maps workflow-id and version to the set
                of fields that are managed by that workflow. This is mostly for internal
                housekeeping, and users typically shouldn't need to set or understand
                this field. A workflow can be the user's name, a controller's name,
                or the name of a specific apply path like \"ci-cd\". The set of fields
                is always in the version that the workflow used when modifying the
                object. \n This field is alpha and can be changed or removed without
                notice."
              items:
                properties:
                  apiVersion:
                    description: APIVersion defines the version of this resource that
                      this field set applies to. The format is "group/version" just
                      like the top-level APIVersion field. It is necessary to track
                      the version of a field set because it cannot be automatically
                      converted.
                    type: string
                  fields:
                    additionalProperties: true
                    description: Fields identifies a set of fields.
                    type: object
                  manager:
                    description: Manager is an identifier of the workflow managing
                      these fields.
                    type: string
                  operation:
                    description: Operation is the type of operation which lead to
                      this ManagedFieldsEntry being created. The only valid values
                      for this field are 'Apply' and 'Update'.
                    type: string
                  time:
                    description: Time is timestamp of when these fields were set.
                      It should always be empty if Operation is 'Apply'
                    format: date-time
                    type: string
                type: object
              type: array
            name:
              description: 'Name must be unique within a namespace. Is required when
                creating resources, although some resources may allow a client to
                request the generation of an appropriate name automatically. Name
                is primarily intended for creation idempotence and configuration definition.
                Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
              type: string
            namespace:
              description: "Namespace defines the space within each name must be unique.
                An empty namespace is equivalent to the \"default\" namespace, but
                \"default\" is the canonical representation. Not all objects are required
                to be scoped to a namespace - the value of this field for those objects
                will be empty. \n Must be a DNS_LABEL. Cannot be updated. More info:
                http://kubernetes.io/docs/user-guide/namespaces"
              type: string
            ownerReferences:
              description: List of objects depended by this object. If ALL objects
                in the list have been deleted, this object will be garbage collected.
                If this object is managed by a controller, then an entry in this list
                will point to this controller, with the controller field set to true.
                There cannot be more than one managing controller.
              items:
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  blockOwnerDeletion:
                    description: If true, AND if the owner has the "foregroundDeletion"
                      finalizer, then the owner cannot be deleted from the key-value
                      store until this reference is removed. Defaults to false. To
                      set this field, a user needs "delete" permission of the owner,
                      otherwise 422 (Unprocessable Entity) will be returned.
                    type: boolean
                  controller:
                    description: If true, this reference points to the managing controller.
                    type: boolean
                  kind:
                    description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                    type: string
                  uid:
                    description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                    type: string
                required:
                - apiVersion
                - kind
                - name
                - uid
                type: object
              type: array
            resourceVersion:
              description: "An opaque value that represents the internal version of
                this object that can be used by clients to determine when objects
                have changed. May be used for optimistic concurrency, change detection,
                and the watch operation on a resource or set of resources. Clients
                must treat these values as opaque and passed unmodified back to the
                server. They may only be valid for a particular resource or set of
                resources. \n Populated by the system. Read-only. Value must be treated
                as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency"
              type: string
            selfLink:
              description: SelfLink is a URL representing this object. Populated by
                the system. Read-only.
              type: string
            uid:
              description: "UID is the unique in time and space value for this object.
                It is typically generated by the server on successful creation of
                a resource and is not allowed to change on PUT operations. \n Populated
                by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids"
              type: string
          type: object
        spec:
          properties:
            configSource:
              description: ConfigSource is used to populate in the associated Node
                for dynamic kubelet config. This field already exists in Node, so
                any updates to it in the Machine spec will be automatically copied
                to the linked NodeRef from the status. The rest of dynamic kubelet
                config support should then work as-is.
              properties:
                configMap:
                  description: ConfigMap is a reference to a Node's ConfigMap
                  properties:
                    kubeletConfigKey:
                      description: KubeletConfigKey declares which key of the referenced
                        ConfigMap corresponds to the KubeletConfiguration structure
                        This field is required in all cases.
                      type: string
                    name:
                      description: Name is the metadata.name of the referenced ConfigMap.
                        This field is required in all cases.
                      type: string
                    namespace:
                      description: Namespace is the metadata.namespace of the referenced
                        ConfigMap. This field is required in all cases.
                      type: string
                    resourceVersion:
                      description: ResourceVersion is the metadata.ResourceVersion
                        of the referenced ConfigMap. This field is forbidden in Node.Spec,
                        and required in Node.Status.
                      type: string
                    uid:
                      description: UID is the metadata.UID of the referenced ConfigMap.
                        This field is forbidden in Node.Spec, and required in Node.Status.
                      type: string
                  required:
                  - namespace
                  - name
                  - kubeletConfigKey
                  type: object
              type: object
            metadata:
              description: ObjectMeta will autopopulate the Node created. Use this
                to indicate what labels, annotations, name prefix, etc., should be
                used when creating the Node.
              properties:
                annotations:
                  additionalProperties:
                    type: string
                  description: 'Annotations is an unstructured key value map stored
                    with a resource that may be set by external tools to store and
                    retrieve arbitrary metadata. They are not queryable and should
                    be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations'
                  type: object
                generateName:
                  description: "GenerateName is an optional prefix, used by the server,
                    to generate a unique name ONLY IF the Name field has not been
                    provided. If this field is used, the name returned to the client
                    will be different than the name passed. This value will also be
                    combined with a unique suffix. The provided value has the same
                    validation rules as the Name field, and may be truncated by the
                    length of the suffix required to make the value unique on the
                    server. \n If this field is specified and the generated name exists,
                    the server will NOT return a 409 - instead, it will either return
                    201 Created or 500 with Reason ServerTimeout indicating a unique
                    name could not be found in the time allotted, and the client should
                    retry (optionally after the time indicated in the Retry-After
                    header). \n Applied only if Name is not specified. More info:
                    https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
                  type: string
                labels:
                  additionalProperties:
                    type: string
                  description: 'Map of string keys and values that can be used to
                    organize and categorize (scope and select) objects. May match
                    selectors of replication controllers and services. More info:
                    http://kubernetes.io/docs/user-guide/labels'
                  type: object
                name:
                  description: 'Name must be unique within a namespace. Is required
                    when creating resources, although some resources may allow a client
                    to request the generation of an appropriate name automatically.
                    Name is primarily intended for creation idempotence and configuration
                    definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                  type: string
                namespace:
                  description: "Namespace defines the space within each name must
                    be unique. An empty namespace is equivalent to the \"default\"
                    namespace, but \"default\" is the canonical representation. Not
                    all objects are required to be scoped to a namespace - the value
                    of this field for those objects will be empty. \n Must be a DNS_LABEL.
                    Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces"
                  type: string
                ownerReferences:
                  description: List of objects depended by this object. If ALL objects
                    in the list have been deleted, this object will be garbage collected.
                    If this object is managed by a controller, then an entry in this
                    list will point to this controller, with the controller field
                    set to true. There cannot be more than one managing controller.
                  items:
                    properties:
                      apiVersion:
                        description: API version of the referent.
                        type: string
                      blockOwnerDeletion:
                        description: If true, AND if the owner has the "foregroundDeletion"
                          finalizer, then the owner cannot be deleted from the key-value
                          store until this reference is removed. Defaults to false.
                          To set this field, a user needs "delete" permission of the
                          owner, otherwise 422 (Unprocessable Entity) will be returned.
                        type: boolean
                      controller:
                        description: If true, this reference points to the managing
                          controller.
                        type: boolean
                      kind:
                        description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                        type: string
                      name:
                        description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                        type: string
                      uid:
                        description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                        type: string
                    required:
                    - apiVersion
                    - kind
                    - name
                    - uid
                    type: object
                  type: array
              type: object
            providerID:
              description: ProviderID is the identification ID of the machine provided
                by the provider. This field must match the provider ID as seen on
                the node object corresponding to this machine. This field is required
                by higher level consumers of cluster-api. Example use case is cluster
                autoscaler with cluster-api as provider. Clean-up logic in the autoscaler
                compares machines to nodes to find out machines at provider which
                could not get registered as Kubernetes nodes. With cluster-api as
                a generic out-of-tree provider for autoscaler, this field is required
                by autoscaler to be able to have a provider view of the list of machines.
                Another list of nodes is queried from the k8s apiserver and then a
                comparison is done to find out unregistered machines and are marked
                for delete. This field will be set by the actuators and consumed by
                higher level entities like autoscaler that will be interfacing with
                cluster-api as generic provider.
              type: string
            providerSpec:
              description: ProviderSpec details Provider-specific configuration to
                use during node creation.
              properties:
                value:
                  description: Value is an inlined, serialized representation of the
                    resource configuration. It is recommended that providers maintain
                    their own versioned API types that should be serialized/deserialized
                    from this field, akin to component config.
                  type: object
                valueFrom:
                  description: Source for the provider configuration. Cannot be used
                    if value is not empty.
                  properties:
                    machineClass:
                      description: The machine class from which the provider config
                        should be sourced.
                      properties:
                        apiVersion:
                          description: API version of the referent.
                          type: string
                        fieldPath:
                          description: 'If referring to a piece of an object instead
                            of an entire object, this string should contain a valid
                            JSON/Go field access statement, such as desiredState.manifest.containers[2].
                            For example, if the object reference is to a container
                            within a pod, this would take on a value like: "spec.containers{name}"
                            (where "name" refers to the name of the container that
                            triggered the event) or if no container name is specified
                            "spec.containers[2]" (container with index 2 in this pod).
                            This syntax is chosen only to have some well-defined way
                            of referencing a part of an object. TODO: this design
                            is not final and this field is subject to change in the
                            future.'
                          type: string
                        kind:
                          description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                          type: string
                        namespace:
                          description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                          type: string
                        provider:
                          description: Provider is the name of the cloud-provider
                            which MachineClass is intended for.
                          type: string
                        resourceVersion:
                          description: 'Specific resourceVersion to which this reference
                            is made, if any. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        uid:
                          description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                          type: string
                      type: object
                  type: object
              type: object
            taints:
              description: The list of the taints to be applied to the corresponding
                Node in additive manner. This list will not overwrite any other taints
                added to the Node on an ongoing basis by other entities. These taints
                should be actively reconciled e.g. if you ask the machine controller
                to apply a taint and then manually remove the taint the machine controller
                will put it back) but not have the machine controller remove any taints
              items:
                properties:
                  effect:
                    description: Required. The effect of the taint on pods that do
                      not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule
                      and NoExecute.
                    type: string
                  key:
                    description: Required. The taint key to be applied to a node.
                    type: string
                  timeAdded:
                    description: TimeAdded represents the time at which the taint
                      was added. It is only written for NoExecute taints.
                    format: date-time
                    type: string
                  value:
                    description: Required. The taint value corresponding to the taint
                      key.
                    type: string
                required:
                - key
                - effect
                type: object
              type: array
            versions:
              description: Versions of key software to use. This field is optional
                at cluster creation time, and omitting the field indicates that the
                cluster installation tool should select defaults for the user. These
                defaults may differ based on the cluster installer, but the tool should
                populate the values it uses when persisting Machine objects. A Machine
                spec missing this field at runtime is invalid.
              properties:
                controlPlane:
                  description: ControlPlane is the semantic version of the Kubernetes
                    control plane to run. This should only be populated when the machine
                    is a control plane.
                  type: string
                kubelet:
                  description: Kubelet is the semantic version of kubelet to run
                  type: string
              required:
              - kubelet
              type: object
          required:
          - providerSpec
          type: object
        status:
          properties:
            addresses:
              description: Addresses is a list of addresses assigned to the machine.
                Queried from cloud provider, if available.
              items:
                properties:
                  address:
                    description: The node address.
                    type: string
                  type:
                    description: Node address type, one of Hostname, ExternalIP or
                      InternalIP.
                    type: string
                required:
                - type
                - address
                type: object
              type: array
            conditions:
              description: 'Conditions lists the conditions synced from the node conditions
                of the corresponding node-object. Machine-controller is responsible
                for keeping conditions up-to-date. MachineSet controller will be taking
                these conditions as a signal to decide if machine is healthy or needs
                to be replaced. Refer: https://kubernetes.io/docs/concepts/architecture/nodes/#condition'
              items:
                properties:
                  lastHeartbeatTime:
                    description: Last time we got an update on a given condition.
                    format: date-time
                    type: string
                  lastTransitionTime:
                    description: Last time the condition transit from one status to
                      another.
                    format: date-time
                    type: string
                  message:
                    description: Human readable message indicating details about last
                      transition.
                    type: string
                  reason:
                    description: (brief) reason for the condition's last transition.
                    type: string
                  status:
                    description: Status of the condition, one of True, False, Unknown.
                    type: string
                  type:
                    description: Type of node condition.
                    type: string
                required:
                - type
                - status
                type: object
              type: array
            errorMessage:
              description: "ErrorMessage will be set in the event that there is a
                terminal problem reconciling the Machine and will contain a more verbose
                string suitable for logging and human consumption. \n This field should
                not be set for transitive errors that a controller faces that are
                expected to be fixed automatically over time (like service outages),
                but instead indicate that something is fundamentally wrong with the
                Machine's spec or the configuration of the controller, and that manual
                intervention is required. Examples of terminal errors would be invalid
                combinations of settings in the spec, values that are unsupported
                by the controller, or the responsible controller itself being critically
                misconfigured. \n Any transient errors that occur during the reconciliation
                of Machines can be added as events to the Machine object and/or logged
                in the controller's output."
              type: string
            errorReason:
              description: "ErrorReason will be set in the event that there is a terminal
                problem reconciling the Machine and will contain a succinct value
                suitable for machine interpretation. \n This field should not be set
                for transitive errors that a controller faces that are expected to
                be fixed automatically over time (like service outages), but instead
                indicate that something is fundamentally wrong with the Machine's
                spec or the configuration of the controller, and that manual intervention
                is required. Examples of terminal errors would be invalid combinations
                of settings in the spec, values that are unsupported by the controller,
                or the responsible controller itself being critically misconfigured.
                \n Any transient errors that occur during the reconciliation of Machines
                can be added as events to the Machine object and/or logged in the
                controller's output."
              type: string
            lastOperation:
              description: LastOperation describes the last-operation performed by
                the machine-controller. This API should be useful as a history in
                terms of the latest operation performed on the specific machine. It
                should also convey the state of the latest-operation for example if
                it is still on-going, failed or completed successfully.
              properties:
                description:
                  description: Description is the human-readable description of the
                    last operation.
                  type: string
                lastUpdated:
                  description: LastUpdated is the timestamp at which LastOperation
                    API was last-updated.
                  format: date-time
                  type: string
                state:
                  description: State is the current status of the last performed operation.
                    E.g. Processing, Failed, Successful etc
                  type: string
                type:
                  description: Type is the type of operation which was last performed.
                    E.g. Create, Delete, Update etc
                  type: string
              type: object
            lastUpdated:
              description: LastUpdated identifies when this status was last observed.
              format: date-time
              type: string
            nodeRef:
              description: NodeRef will point to the corresponding Node if it exists.
              properties:
                apiVersion:
                  description: API version of the referent.
                  type: string
                fieldPath:
                  description: 'If referring to a piece of an object instead of an
                    entire object, this string should contain a valid JSON/Go field
                    access statement, such as desiredState.manifest.containers[2].
                    For example, if the object reference is to a container within
                    a pod, this would take on a value like: "spec.containers{name}"
                    (where "name" refers to the name of the container that triggered
                    the event) or if no container name is specified "spec.containers[2]"
                    (container with index 2 in this pod). This syntax is chosen only
                    to have some well-defined way of referencing a part of an object.
                    TODO: this design is not final and this field is subject to change
                    in the future.'
                  type: string
                kind:
                  description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                  type: string
                name:
                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                  type: string
                namespace:
                  description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                  type: string
                resourceVersion:
                  description: 'Specific resourceVersion to which this reference is
                    made, if any. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                  type: string
                uid:
                  description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                  type: string
              type: object
            phase:
              description: Phase represents the current phase of machine actuation.
                E.g. Pending, Running, Terminating, Failed etc.
              type: string
            providerStatus:
              description: ProviderStatus details a Provider-specific status. It is
                recommended that providers maintain their own versioned API types
                that should be serialized/deserialized from this field.
              type: object
            versions:
              description: "Versions specifies the current versions of software on
                the corresponding Node (if it exists). This is provided for a few
                reasons: \n 1) It is more convenient than checking the NodeRef, traversing
                it to    the Node, and finding the appropriate field in Node.Status.NodeInfo
                \   (which uses different field names and formatting). 2) It removes
                some of the dependency on the structure of the Node,    so that if
                the structure of Node.Status.NodeInfo changes, only    machine controllers
                need to be updated, rather than every client    of the Machines API.
                3) There is no other simple way to check the control plane    version.
                A client would have to connect directly to the apiserver    running
                on the target node in order to find out its version."
              properties:
                controlPlane:
                  description: ControlPlane is the semantic version of the Kubernetes
                    control plane to run. This should only be populated when the machine
                    is a control plane.
                  type: string
                kubelet:
                  description: Kubelet is the semantic version of kubelet to run
                  type: string
              required:
              - kubelet
              type: object
          type: object
      type: object
  versions:
  - name: v1alpha1
    served: true
    storage: true
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
  name: machinesets.cluster.k8s.io
spec:
  group: cluster.k8s.io
  names:
    kind: MachineSet
    plural: machinesets
    shortNames:
    - ms
  scope: Namespaced
  subresources:
    scale:
      labelSelectorPath: .status.labelSelector
      specReplicasPath: .spec.replicas
      statusReplicasPath: .status.replicas
    status: {}
  validation:
    openAPIV3Schema:
      description: / [MachineSet] MachineSet ensures that a specified number of machines
        replicas are running at any given time.
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        metadata:
          properties:
            annotations:
              additionalProperties:
                type: string
              description: 'Annotations is an unstructured key value map stored with
                a resource that may be set by external tools to store and retrieve
                arbitrary metadata. They are not queryable and should be preserved
                when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations'
              type: object
            clusterName:
              description: The name of the cluster which the object belongs to. This
                is used to distinguish resources with same name and namespace in different
                clusters. This field is not set anywhere right now and apiserver is
                going to ignore it if set in create or update request.
              type: string
            creationTimestamp:
              description: "CreationTimestamp is a timestamp representing the server
                time when this object was created. It is not guaranteed to be set
                in happens-before order across separate operations. Clients may not
                set this value. It is represented in RFC3339 form and is in UTC. \n
                Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            deletionGracePeriodSeconds:
              description: Number of seconds allowed for this object to gracefully
                terminate before it will be removed from the system. Only set when
                deletionTimestamp is also set. May only be shortened. Read-only.
              format: int64
              type: integer
            deletionTimestamp:
              description: "DeletionTimestamp is RFC 3339 date and time at which this
                resource will be deleted. This field is set by the server when a graceful
                deletion is requested by the user, and is not directly settable by
                a client. The resource is expected to be deleted (no longer visible
                from resource lists, and not reachable by name) after the time in
                this field, once the finalizers list is empty. As long as the finalizers
                list contains items, deletion is blocked. Once the deletionTimestamp
                is set, this value may not be unset or be set further into the future,
                although it may be shortened or the resource may be deleted prior
                to this time. For example, a user may request that a pod is deleted
                in 30 seconds. The Kubelet will react by sending a graceful termination
                signal to the containers in the pod. After that 30 seconds, the Kubelet
                will send a hard termination signal (SIGKILL) to the container and
                after cleanup, remove the pod from the API. In the presence of network
                partitions, this object may still exist after this timestamp, until
                an administrator or automated process can determine the resource is
                fully terminated. If not set, graceful deletion of the object has
                not been requested. \n Populated by the system when a graceful deletion
                is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
              format: date-time
              type: string
            finalizers:
              description: Must be empty before the object is deleted from the registry.
                Each entry is an identifier for the responsible component that will
                remove the entry from the list. If the deletionTimestamp of the object
                is non-nil, entries in this list can only be removed.
              items:
                type: string
              type: array
            generateName:
              description: "GenerateName is an optional prefix, used by the server,
                to generate a unique name ONLY IF the Name field has not been provided.
                If this field is used, the name returned to the client will be different
                than the name passed. This value will also be combined with a unique
                suffix. The provided value has the same validation rules as the Name
                field, and may be truncated by the length of the suffix required to
                make the value unique on the server. \n If this field is specified
                and the generated name exists, the server will NOT return a 409 -
                instead, it will either return 201 Created or 500 with Reason ServerTimeout
                indicating a unique name could not be found in the time allotted,
                and the client should retry (optionally after the time indicated in
                the Retry-After header). \n Applied only if Name is not specified.
                More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
              type: string
            generation:
              description: A sequence number representing a specific generation of
                the desired state. Populated by the system. Read-only.
              format: int64
              type: integer
            initializers:
              description: "An initializer is a controller which enforces some system
                invariant at object creation time. This field is a list of initializers
                that have not yet acted on this object. If nil or empty, this object
                has been completely initialized. Otherwise, the object is considered
                uninitialized and is hidden (in list/watch and get calls) from clients
                that haven't explicitly asked to observe uninitialized objects. \n
                When an object is created, the system will populate this list with
                the current set of initializers. Only privileged users may set or
                modify this list. Once it is empty, it may not be modified further
                by any user. \n DEPRECATED - initializers are an alpha field and will
                be removed in v1.15."
              properties:
                pending:
                  description: Pending is a list of initializers that must execute
                    in order before this object is visible. When the last pending
                    initializer is removed, and no failing result is set, the initializers
                    struct will be set to nil and the object is considered as initialized
                    and visible to all clients.
                  items:
                    properties:
                      name:
                        description: name of the process that is responsible for initializing
                          this object.
                        type: string
                    required:
                    - name
                    type: object
                  type: array
                result:
                  description: If result is set with the Failure field, the object
                    will be persisted to storage and then deleted, ensuring that other
                    clients can observe the deletion.
                  properties:
                    apiVersion:
                      description: 'APIVersion defines the versioned schema of this
                        representation of an object. Servers should convert recognized
                        schemas to the latest internal value, and may reject unrecognized
                        values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
                      type: string
                    code:
                      description: Suggested HTTP return code for this status, 0 if
                        not set.
                      format: int32
                      type: integer
                    details:
                      description: Extended data associated with the reason.  Each
                        reason may define its own extended details. This field is
                        optional and the data returned is not guaranteed to conform
                        to any schema except that defined by the reason type.
                      properties:
                        causes:
                          description: The Causes array includes more details associated
                            with the StatusReason failure. Not all StatusReasons may
                            provide detailed causes.
                          items:
                            properties:
                              field:
                                description: "The field of the resource that has caused
                                  this error, as named by its JSON serialization.
                                  May include dot and postfix notation for nested
                                  attributes. Arrays are zero-indexed.  Fields may
                                  appear more than once in an array of causes due
                                  to fields having multiple errors. Optional. \n Examples:
                                  \  \"name\" - the field \"name\" on the current
                                  resource   \"items[0].name\" - the field \"name\"
                                  on the first array entry in \"items\""
                                type: string
                              message:
                                description: A human-readable description of the cause
                                  of the error.  This field may be presented as-is
                                  to a reader.
                                type: string
                              reason:
                                description: A machine-readable description of the
                                  cause of the error. If this value is empty there
                                  is no information available.
                                type: string
                            type: object
                          type: array
                        group:
                          description: The group attribute of the resource associated
                            with the status StatusReason.
                          type: string
                        kind:
                          description: 'The kind attribute of the resource associated
                            with the status StatusReason. On some operations may differ
                            from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: The name attribute of the resource associated
                            with the status StatusReason (when there is a single name
                            which can be described).
                          type: string
                        retryAfterSeconds:
                          description: If specified, the time in seconds before the
                            operation should be retried. Some errors may indicate
                            the client must take an alternate action - for those errors
                            this field may indicate how long to wait before taking
                            the alternate action.
                          format: int32
                          type: integer
                        uid:
                          description: 'UID of the resource. (when there is a single
                            resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                          type: string
                      type: object
                    kind:
                      description: 'Kind is a string value representing the REST resource
                        this object represents. Servers may infer this from the endpoint
                        the client submits requests to. Cannot be updated. In CamelCase.
                        More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      type: string
                    message:
                      description: A human-readable description of the status of this
                        operation.
                      type: string
                    metadata:
                      description: 'Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                      properties:
                        continue:
                          description: continue may be set if the user set a limit
                            on the number of items returned, and indicates that the
                            server has more data available. The value is opaque and
                            may be used to issue another request to the endpoint that
                            served this list to retrieve the next set of available
                            objects. Continuing a consistent list may not be possible
                            if the server configuration has changed or more than a
                            few minutes have passed. The resourceVersion field returned
                            when using this continue value will be identical to the
                            value in the first response, unless you have received
                            this token from an error message.
                          type: string
                        resourceVersion:
                          description: 'String that identifies the server''s internal
                            version of this object that can be used by clients to
                            determine when objects have changed. Value must be treated
                            as opaque by clients and passed unmodified back to the
                            server. Populated by the system. Read-only. More info:
                            https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        selfLink:
                          description: selfLink is a URL representing this object.
                            Populated by the system. Read-only.
                          type: string
                      type: object
                    reason:
                      description: A machine-readable description of why this operation
                        is in the "Failure" status. If this value is empty there is
                        no information available. A Reason clarifies an HTTP status
                        code but does not override it.
                      type: string
                    status:
                      description: 'Status of the operation. One of: "Success" or
                        "Failure". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status'
                      type: string
                  type: object
              required:
              - pending
              type: object
            labels:
              additionalProperties:
                type: string
              description: 'Map of string keys and values that can be used to organize
                and categorize (scope and select) objects. May match selectors of
                replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels'
              type: object
            managedFields:
              description: "ManagedFields maps workflow-id and version to the set
                of fields that are managed by that workflow. This is mostly for internal
                housekeeping, and users typically shouldn't need to set or understand
                this field. A workflow can be the user's name, a controller's name,
                or the name of a specific apply path like \"ci-cd\". The set of fields
                is always in the version that the workflow used when modifying the
                object. \n This field is alpha and can be changed or removed without
                notice."
              items:
                properties:
                  apiVersion:
                    description: APIVersion defines the version of this resource that
                      this field set applies to. The format is "group/version" just
                      like the top-level APIVersion field. It is necessary to track
                      the version of a field set because it cannot be automatically
                      converted.
                    type: string
                  fields:
                    additionalProperties: true
                    description: Fields identifies a set of fields.
                    type: object
                  manager:
                    description: Manager is an identifier of the workflow managing
                      these fields.
                    type: string
                  operation:
                    description: Operation is the type of operation which lead to
                      this ManagedFieldsEntry being created. The only valid values
                      for this field are 'Apply' and 'Update'.
                    type: string
                  time:
                    description: Time is timestamp of when these fields were set.
                      It should always be empty if Operation is 'Apply'
                    format: date-time
                    type: string
                type: object
              type: array
            name:
              description: 'Name must be unique within a namespace. Is required when
                creating resources, although some resources may allow a client to
                request the generation of an appropriate name automatically. Name
                is primarily intended for creation idempotence and configuration definition.
                Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
              type: string
            namespace:
              description: "Namespace defines the space within each name must be unique.
                An empty namespace is equivalent to the \"default\" namespace, but
                \"default\" is the canonical representation. Not all objects are required
                to be scoped to a namespace - the value of this field for those objects
                will be empty. \n Must be a DNS_LABEL. Cannot be updated. More info:
                http://kubernetes.io/docs/user-guide/namespaces"
              type: string
            ownerReferences:
              description: List of objects depended by this object. If ALL objects
                in the list have been deleted, this object will be garbage collected.
                If this object is managed by a controller, then an entry in this list
                will point to this controller, with the controller field set to true.
                There cannot be more than one managing controller.
              items:
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  blockOwnerDeletion:
                    description: If true, AND if the owner has the "foregroundDeletion"
                      finalizer, then the owner cannot be deleted from the key-value
                      store until this reference is removed. Defaults to false. To
                      set this field, a user needs "delete" permission of the owner,
                      otherwise 422 (Unprocessable Entity) will be returned.
                    type: boolean
                  controller:
                    description: If true, this reference points to the managing controller.
                    type: boolean
                  kind:
                    description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                    type: string
                  uid:
                    description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                    type: string
                required:
                - apiVersion
                - kind
                - name
                - uid
                type: object
              type: array
            resourceVersion:
              description: "An opaque value that represents the internal version of
                this object that can be used by clients to determine when objects
                have changed. May be used for optimistic concurrency, change detection,
                and the watch operation on a resource or set of resources. Clients
                must treat these values as opaque and passed unmodified back to the
                server. They may only be valid for a particular resource or set of
                resources. \n Populated by the system. Read-only. Value must be treated
                as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency"
              type: string
            selfLink:
              description: SelfLink is a URL representing this object. Populated by
                the system. Read-only.
              type: string
            uid:
              description: "UID is the unique in time and space value for this object.
                It is typically generated by the server on successful creation of
                a resource and is not allowed to change on PUT operations. \n Populated
                by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids"
              type: string
          type: object
        spec:
          properties:
            deletePolicy:
              description: DeletePolicy defines the policy used to identify nodes
                to delete when downscaling. Defaults to "Random".  Valid values are
                "Random, "Newest", "Oldest"
              enum:
              - Random
              - Newest
              - Oldest
              type: string
            minReadySeconds:
              description: MinReadySeconds is the minimum number of seconds for which
                a newly created machine should be ready. Defaults to 0 (machine will
                be considered available as soon as it is ready)
              format: int32
              type: integer
            replicas:
              description: Replicas is the number of desired replicas. This is a pointer
                to distinguish between explicit zero and unspecified. Defaults to
                1.
              format: int32
              type: integer
            selector:
              description: 'Selector is a label query over machines that should match
                the replica count. Label keys and values that must match in order
                to be controlled by this MachineSet. It must match the machine template''s
                labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors'
              properties:
                matchExpressions:
                  description: matchExpressions is a list of label selector requirements.
                    The requirements are ANDed.
                  items:
                    properties:
                      key:
                        description: key is the label key that the selector applies
                          to.
                        type: string
                      operator:
                        description: operator represents a key's relationship to a
                          set of values. Valid operators are In, NotIn, Exists and
                          DoesNotExist.
                        type: string
                      values:
                        description: values is an array of string values. If the operator
                          is In or NotIn, the values array must be non-empty. If the
                          operator is Exists or DoesNotExist, the values array must
                          be empty. This array is replaced during a strategic merge
                          patch.
                        items:
                          type: string
                        type: array
                    required:
                    - key
                    - operator
                    type: object
                  type: array
                matchLabels:
                  additionalProperties:
                    type: string
                  description: matchLabels is a map of {key,value} pairs. A single
                    {key,value} in the matchLabels map is equivalent to an element
                    of matchExpressions, whose key field is "key", the operator is
                    "In", and the values array contains only "value". The requirements
                    are ANDed.
                  type: object
              type: object
            template:
              description: Template is the object that describes the machine that
                will be created if insufficient replicas are detected.
              properties:
                metadata:
                  description: 'Standard object''s metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
                  properties:
                    annotations:
                      additionalProperties:
                        type: string
                      description: 'Annotations is an unstructured key value map stored
                        with a resource that may be set by external tools to store
                        and retrieve arbitrary metadata. They are not queryable and
                        should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations'
                      type: object
                    generateName:
                      description: "GenerateName is an optional prefix, used by the
                        server, to generate a unique name ONLY IF the Name field has
                        not been provided. If this field is used, the name returned
                        to the client will be different than the name passed. This
                        value will also be combined with a unique suffix. The provided
                        value has the same validation rules as the Name field, and
                        may be truncated by the length of the suffix required to make
                        the value unique on the server. \n If this field is specified
                        and the generated name exists, the server will NOT return
                        a 409 - instead, it will either return 201 Created or 500
                        with Reason ServerTimeout indicating a unique name could not
                        be found in the time allotted, and the client should retry
                        (optionally after the time indicated in the Retry-After header).
                        \n Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
                      type: string
                    labels:
                      additionalProperties:
                        type: string
                      description: 'Map of string keys and values that can be used
                        to organize and categorize (scope and select) objects. May
                        match selectors of replication controllers and services. More
                        info: http://kubernetes.io/docs/user-guide/labels'
                      type: object
                    name:
                      description: 'Name must be unique within a namespace. Is required
                        when creating resources, although some resources may allow
                        a client to request the generation of an appropriate name
                        automatically. Name is primarily intended for creation idempotence
                        and configuration definition. Cannot be updated. More info:
                        http://kubernetes.io/docs/user-guide/identifiers#names'
                      type: string
                    namespace:
                      description: "Namespace defines the space within each name must
                        be unique. An empty namespace is equivalent to the \"default\"
                        namespace, but \"default\" is the canonical representation.
                        Not all objects are required to be scoped to a namespace -
                        the value of this field for those objects will be empty. \n
                        Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces"
                      type: string
                    ownerReferences:
                      description: List of objects depended by this object. If ALL
                        objects in the list have been deleted, this object will be
                        garbage collected. If this object is managed by a controller,
                        then an entry in this list will point to this controller,
                        with the controller field set to true. There cannot be more
                        than one managing controller.
                      items:
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          blockOwnerDeletion:
                            description: If true, AND if the owner has the "foregroundDeletion"
                              finalizer, then the owner cannot be deleted from the
                              key-value store until this reference is removed. Defaults
                              to false. To set this field, a user needs "delete" permission
                              of the owner, otherwise 422 (Unprocessable Entity) will
                              be returned.
                            type: boolean
                          controller:
                            description: If true, this reference points to the managing
                              controller.
                            type: boolean
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                            type: string
                        required:
                        - apiVersion
                        - kind
                        - name
                        - uid
                        type: object
                      type: array
                  type: object
                spec:
                  description: 'Specification of the desired behavior of the machine.
                    More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status'
                  properties:
                    configSource:
                      description: ConfigSource is used to populate in the associated
                        Node for dynamic kubelet config. This field already exists
                        in Node, so any updates to it in the Machine spec will be
                        automatically copied to the linked NodeRef from the status.
                        The rest of dynamic kubelet config support should then work
                        as-is.
                      properties:
                        configMap:
                          description: ConfigMap is a reference to a Node's ConfigMap
                          properties:
                            kubeletConfigKey:
                              description: KubeletConfigKey declares which key of
                                the referenced ConfigMap corresponds to the KubeletConfiguration
                                structure This field is required in all cases.
                              type: string
                            name:
                              description: Name is the metadata.name of the referenced
                                ConfigMap. This field is required in all cases.
                              type: string
                            namespace:
                              description: Namespace is the metadata.namespace of
                                the referenced ConfigMap. This field is required in
                                all cases.
                              type: string
                            resourceVersion:
                              description: ResourceVersion is the metadata.ResourceVersion
                                of the referenced ConfigMap. This field is forbidden
                                in Node.Spec, and required in Node.Status.
                              type: string
                            uid:
                              description: UID is the metadata.UID of the referenced
                                ConfigMap. This field is forbidden in Node.Spec, and
                                required in Node.Status.
                              type: string
                          required:
                          - namespace
                          - name
                          - kubeletConfigKey
                          type: object
                      type: object
                    metadata:
                      description: ObjectMeta will autopopulate the Node created.
                        Use this to indicate what labels, annotations, name prefix,
                        etc., should be used when creating the Node.
                      properties:
                        annotations:
                          additionalProperties:
                            type: string
                          description: 'Annotations is an unstructured key value map
                            stored with a resource that may be set by external tools
                            to store and retrieve arbitrary metadata. They are not
                            queryable and should be preserved when modifying objects.
                            More info: http://kubernetes.io/docs/user-guide/annotations'
                          type: object
                        generateName:
                          description: "GenerateName is an optional prefix, used by
                            the server, to generate a unique name ONLY IF the Name
                            field has not been provided. If this field is used, the
                            name returned to the client will be different than the
                            name passed. This value will also be combined with a unique
                            suffix. The provided value has the same validation rules
                            as the Name field, and may be truncated by the length
                            of the suffix required to make the value unique on the
                            server. \n If this field is specified and the generated
                            name exists, the server will NOT return a 409 - instead,
                            it will either return 201 Created or 500 with Reason ServerTimeout
                            indicating a unique name could not be found in the time
                            allotted, and the client should retry (optionally after
                            the time indicated in the Retry-After header). \n Applied
                            only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
                          type: string
                        labels:
                          additionalProperties:
                            type: string
                          description: 'Map of string keys and values that can be
                            used to organize and categorize (scope and select) objects.
                            May match selectors of replication controllers and services.
                            More info: http://kubernetes.io/docs/user-guide/labels'
                          type: object
                        name:
                          description: 'Name must be unique within a namespace. Is
                            required when creating resources, although some resources
                            may allow a client to request the generation of an appropriate
                            name automatically. Name is primarily intended for creation
                            idempotence and configuration definition. Cannot be updated.
                            More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                          type: string
                        namespace:
                          description: "Namespace defines the space within each name
                            must be unique. An empty namespace is equivalent to the
                            \"default\" namespace, but \"default\" is the canonical
                            representation. Not all objects are required to be scoped
                            to a namespace - the value of this field for those objects
                            will be empty. \n Must be a DNS_LABEL. Cannot be updated.
                            More info: http://kubernetes.io/docs/user-guide/namespaces"
                          type: string
                        ownerReferences:
                          description: List of objects depended by this object. If
                            ALL objects in the list have been deleted, this object
                            will be garbage collected. If this object is managed by
                            a controller, then an entry in this list will point to
                            this controller, with the controller field set to true.
                            There cannot be more than one managing controller.
                          items:
                            properties:
                              apiVersion:
                                description: API version of the referent.
                                type: string
                              blockOwnerDeletion:
                                description: If true, AND if the owner has the "foregroundDeletion"
                                  finalizer, then the owner cannot be deleted from
                                  the key-value store until this reference is removed.
                                  Defaults to false. To set this field, a user needs
                                  "delete" permission of the owner, otherwise 422
                                  (Unprocessable Entity) will be returned.
                                type: boolean
                              controller:
                                description: If true, this reference points to the
                                  managing controller.
                                type: boolean
                              kind:
                                description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                                type: string
                              name:
                                description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                                type: string
                              uid:
                                description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                                type: string
                            required:
                            - apiVersion
                            - kind
                            - name
                            - uid
                            type: object
                          type: array
                      type: object
                    providerID:
                      description: ProviderID is the identification ID of the machine
                        provided by the provider. This field must match the provider
                        ID as seen on the node object corresponding to this machine.
                        This field is required by higher level consumers of cluster-api.
                        Example use case is cluster autoscaler with cluster-api as
                        provider. Clean-up logic in the autoscaler compares machines
                        to nodes to find out machines at provider which could not
                        get registered as Kubernetes nodes. With cluster-api as a
                        generic out-of-tree provider for autoscaler, this field is
                        required by autoscaler to be able to have a provider view
                        of the list of machines. Another list of nodes is queried
                        from the k8s apiserver and then a comparison is done to find
                        out unregistered machines and are marked for delete. This
                        field will be set by the actuators and consumed by higher
                        level entities like autoscaler that will be interfacing with
                        cluster-api as generic provider.
                      type: string
                    providerSpec:
                      description: ProviderSpec details Provider-specific configuration
                        to use during node creation.
                      properties:
                        value:
                          description: Value is an inlined, serialized representation
                            of the resource configuration. It is recommended that
                            providers maintain their own versioned API types that
                            should be serialized/deserialized from this field, akin
                            to component config.
                          type: object
                        valueFrom:
                          description: Source for the provider configuration. Cannot
                            be used if value is not empty.
                          properties:
                            machineClass:
                              description: The machine class from which the provider
                                config should be sourced.
                              properties:
                                apiVersion:
                                  description: API version of the referent.
                                  type: string
                                fieldPath:
                                  description: 'If referring to a piece of an object
                                    instead of an entire object, this string should
                                    contain a valid JSON/Go field access statement,
                                    such as desiredState.manifest.containers[2]. For
                                    example, if the object reference is to a container
                                    within a pod, this would take on a value like:
                                    "spec.containers{name}" (where "name" refers to
                                    the name of the container that triggered the event)
                                    or if no container name is specified "spec.containers[2]"
                                    (container with index 2 in this pod). This syntax
                                    is chosen only to have some well-defined way of
                                    referencing a part of an object. TODO: this design
                                    is not final and this field is subject to change
                                    in the future.'
                                  type: string
                                kind:
                                  description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                                  type: string
                                name:
                                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                  type: string
                                namespace:
                                  description: 'Namespace of the referent. More info:
                                    https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                  type: string
                                provider:
                                  description: Provider is the name of the cloud-provider
                                    which MachineClass is intended for.
                                  type: string
                                resourceVersion:
                                  description: 'Specific resourceVersion to which
                                    this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                                  type: string
                                uid:
                                  description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                                  type: string
                              type: object
                          type: object
                      type: object
                    taints:
                      description: The list of the taints to be applied to the corresponding
                        Node in additive manner. This list will not overwrite any
                        other taints added to the Node on an ongoing basis by other
                        entities. These taints should be actively reconciled e.g.
                        if you ask the machine controller to apply a taint and then
                        manually remove the taint the machine controller will put
                        it back) but not have the machine controller remove any taints
                      items:
                        properties:
                          effect:
                            description: Required. The effect of the taint on pods
                              that do not tolerate the taint. Valid effects are NoSchedule,
                              PreferNoSchedule and NoExecute.
                            type: string
                          key:
                            description: Required. The taint key to be applied to
                              a node.
                            type: string
                          timeAdded:
                            description: TimeAdded represents the time at which the
                              taint was added. It is only written for NoExecute taints.
                            format: date-time
                            type: string
                          value:
                            description: Required. The taint value corresponding to
                              the taint key.
                            type: string
                        required:
                        - key
                        - effect
                        type: object
                      type: array
                    versions:
                      description: Versions of key software to use. This field is
                        optional at cluster creation time, and omitting the field
                        indicates that the cluster installation tool should select
                        defaults for the user. These defaults may differ based on
                        the cluster installer, but the tool should populate the values
                        it uses when persisting Machine objects. A Machine spec missing
                        this field at runtime is invalid.
                      properties:
                        controlPlane:
                          description: ControlPlane is the semantic version of the
                            Kubernetes control plane to run. This should only be populated
                            when the machine is a control plane.
                          type: string
                        kubelet:
                          description: Kubelet is the semantic version of kubelet
                            to run
                          type: string
                      required:
                      - kubelet
                      type: object
                  required:
                  - providerSpec
                  type: object
              type: object
          required:
          - selector
          type: object
        status:
          properties:
            availableReplicas:
              description: The number of available replicas (ready for at least minReadySeconds)
                for this MachineSet.
              format: int32
              type: integer
            errorMessage:
              type: string
            errorReason:
              description: "In the event that there is a terminal problem reconciling
                the replicas, both ErrorReason and ErrorMessage will be set. ErrorReason
                will be populated with a succinct value suitable for machine interpretation,
                while ErrorMessage will contain a more verbose string suitable for
                logging and human consumption. \n These fields should not be set for
                transitive errors that a controller faces that are expected to be
                fixed automatically over time (like service outages), but instead
                indicate that something is fundamentally wrong with the MachineTemplate's
                spec or the configuration of the machine controller, and that manual
                intervention is required. Examples of terminal errors would be invalid
                combinations of settings in the spec, values that are unsupported
                by the machine controller, or the responsible machine controller itself
                being critically misconfigured. \n Any transient errors that occur
                during the reconciliation of Machines can be added as events to the
                MachineSet object and/or logged in the controller's output."
              type: string
            fullyLabeledReplicas:
              description: The number of replicas that have labels matching the labels
                of the machine template of the MachineSet.
              format: int32
              type: integer
            observedGeneration:
              description: ObservedGeneration reflects the generation of the most
                recently observed MachineSet.
              format: int64
              type: integer
            readyReplicas:
              description: The number of ready replicas for this MachineSet. A machine
                is considered ready when the node has been created and is "Ready".
              format: int32
              type: integer
            replicas:
              description: Replicas is the most recently observed number of replicas.
              format: int32
              type: integer
          required:
          - replicas
          type: object
      type: object
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: cluster-api-manager-role
rules:
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - get
  - list
  - watch
  - create
  - patch
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - nodes
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
  - machines
  - machines/status
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
  - machinedeployments
  - machinedeployments/status
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
  - machinesets
  - machinesets/status
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  creationTimestamp: null
  name: cluster-api-manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-api-manager-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: cluster-api-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    control-plane: controller-manager
    controller-tools.k8s.io: "1.0"
  name: cluster-api-controller-manager-service
  namespace: cluster-api-system
spec:
  ports:
  - port: 443
  selector:
    control-plane: controller-manager
    controller-tools.k8s.io: "1.0"
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    control-plane: controller-manager
    controller-tools.k8s.io: "1.0"
  name: cluster-api-controller-manager
  namespace: cluster-api-system
spec:
  selector:
    matchLabels:
      control-plane: controller-manager
      controller-tools.k8s.io: "1.0"
  serviceName: cluster-api-controller-manager-service
  template:
    metadata:
      labels:
        control-plane: controller-manager
        controller-tools.k8s.io: "1.0"
    spec:
      containers:
      - command:
        - /manager
        image: gcr.io/k8s-cluster-api/cluster-api-controller:0.1.0
        name: manager
        resources:
          limits:
            cpu: 100m
            memory: 30Mi
          requests:
            cpu: 100m
            memory: 20Mi
      terminationGracePeriodSeconds: 10
      tolerations:
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoExecute
        key: node.alpha.kubernetes.io/notReady
        operator: Exists
      - effect: NoExecute
        key: node.alpha.kubernetes.io/unreachable
        operator: Exists
`
)
