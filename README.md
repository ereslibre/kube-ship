kube-ship
============

`kube-ship` allows you to create Kubernetes clusters leveraging
`cluster-api`. It supports the following `cluster-api` providers as of today:

* [Proxmox](https://github.com/ereslibre/cluster-api-provider-proxmox)
* [OpenStack](https://github.com/kubernetes-sigs/cluster-api-provider-openstack)

`cluster-api` providers usually allow you to create infrastructure
using their own `clusterctl` command. `kube-ship` aims to unify this
setup, so you can use `kube-ship` to create the infrastructure using any
`cluster-api` provider.

You can think about `kube-ship` as the `clusterctl` for many `cluster-api`
provider implementations.

You can find some examples in the [`samples`](samples) directory. Each
file defines a complete way of deploying a cluster in the given
provider. All you should need is a `Cluster` definition, and a couple
of `MachineDeployment` objects, aside from provider-specific
authentication.

## Building and installing

Building:

```
# make build
```

Installing:

```
# make install
```

## Creating the infrastructure

### Proxmox

Several environment variables need to be set so they can be propagated
to the proxmox cluster-api provider manager:

* `PROXMOX_HOSTPORT`
  * Example: `a-proxmox.some-company.intra.net:8006`
* `PROXMOX_USERNAME`
  * Example: `root@pam`
* `PROXMOX_PASSWORD`
  * Example: `mypassword`
* `PROXMOX_HYPERVISOR_NAME`
  * Example: `proxmox`
* `PROXMOX_HYPERVISOR_SNIPPETS_STORAGE`
  * Example: `ci-snippets`
  * Important: make sure you create this volume using the Proxmox UI
    with `Snippets` content.
* `VM_TEMPLATE_ID`
  * Example: `9000`
  * Important: make sure you have created this template beforehand.

Note: [envvars are required at this time because the proxmox
cluster-api provider is not yet in its final shape](https://github.com/ereslibre/cluster-api-provider-proxmox#todo).

```
# PROXMOX_HOSTPORT=... PROXMOX_USERNAME=... (...) kube-ship --config samples/proxmox.yaml
```

### OpenStack

Before running the command, please check `samples/openstack.yaml` file
for all settings that have to be configured. After you have configured
everything required in this file, you can go ahead and create the infrastructure:

```
# kube-ship --config samples/openstack.yaml
```

## How it works

`kube-ship` will create a temporary [`kind`](`https://sigs.k8s.io/kind`)
cluster, it will then deploy `cluster-api` CRD's and controller inside
this temporary cluster.

`kube-ship` will read your infrastructure definition, then it will
decide what specific provider needs to be created based on this
definition, and will deploy it on top of the temporary cluster. This
definition will also be applied automatically on this temporary
cluster.

After this temporary cluster using `cluster-api` has created and
initialized the infrastructure remotely, this `kind` temporary cluster
will be dropped.

## TODO

- [ ] General cleanup and refactoring
- [ ] Support for AWS
- [ ] Support for Azure
- [ ] Support for GCE
- [ ] Allow for a "management cluster" (`--kubeconfig` param) instead of
      starting a temporary `kind` based cluster
- [ ] Check the best way to check when a deployment on the different
      providers have finished
  * Right now `kube-ship` waits for the `Phase` status on each target
    machine to be `Running`, but not all providers behave in the same way
- [ ] Download the target `kubeconfig` file
- [ ] Pivot `cluster-api` definitions and actuators to the target
      cluster if a "management cluster" is not used

## SUSE Hackweek

This project started as a [SUSE Hackweek 2019 project](https://hackweek.suse.com/).

## License

```
Copyright 2019 Rafael Fernández López <ereslibre@ereslibre.es>
Copyright 2019 SUSE LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
