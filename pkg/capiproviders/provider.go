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

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/ereslibre/kube-ship/pkg/capi"
	"github.com/ereslibre/kube-ship/pkg/kubernetes"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	errorsutil "k8s.io/apimachinery/pkg/util/errors"
	utilyaml "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/yaml"

	proxmoxapis "github.com/ereslibre/cluster-api-provider-proxmox/pkg/apis"
	openstackapis "sigs.k8s.io/cluster-api-provider-openstack/pkg/apis"
	clusterv1alpha1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

type apiVersionKind struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

type provider int

const (
	unknown provider = iota
	proxmox
	openstack
)

func (p provider) String() string {
	switch p {
	case proxmox:
		return "proxmox"
	case openstack:
		return "openstack"
	}
	return "unknown"
}

func splitYAMLDocuments(yamlBytes []byte) ([][]byte, error) {
	documents := [][]byte{}

	errs := []error{}
	buf := bytes.NewBuffer(yamlBytes)
	reader := utilyaml.NewYAMLReader(bufio.NewReader(buf))
	for {
		typeMetaInfo := runtime.TypeMeta{}
		// Read one YAML document at a time, until io.EOF is returned
		b, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if len(b) == 0 {
			break
		}
		// Deserialize the TypeMeta information of this byte slice
		if err := yaml.Unmarshal(b, &typeMetaInfo); err != nil {
			return nil, err
		}
		// Require TypeMeta information to be present
		if len(typeMetaInfo.APIVersion) == 0 || len(typeMetaInfo.Kind) == 0 {
			errs = append(errs, errors.New("invalid configuration: kind and apiVersion is mandatory information that needs to be specified in all YAML documents"))
			continue
		}

		documents = append(documents, b)
	}
	if err := errorsutil.NewAggregate(errs); err != nil {
		return nil, err
	}
	return documents, nil
}

func NewScheme() *runtime.Scheme {
	scheme := capi.NewScheme()
	proxmoxapis.AddToScheme(scheme)
	openstackapis.AddToScheme(scheme)
	return scheme
}

func ListComponents(definitions string) ([]runtime.Object, error) {
	scheme := NewScheme()
	codecFactory := serializer.NewCodecFactory(scheme)

	result := []runtime.Object{}

	documents, err := splitYAMLDocuments([]byte(definitions))
	if err != nil {
		return result, err
	}

	decode := codecFactory.UniversalDeserializer().Decode
	for _, rawDocument := range documents {
		obj, _, err := decode(rawDocument, nil, nil)
		if err != nil {
			return result, err
		}
		result = append(result, obj)
	}

	return result, nil
}

func providerForGroup(group string) provider {
	switch group {
	case "proxmoxproviderconfig.k8s.io":
		return proxmox
	case "openstackproviderconfig":
		return openstack
	}
	return unknown
}

func decideProvider(components []runtime.Object) (provider, error) {
	scheme := NewScheme()

	detectedProvider := unknown
	for _, component := range components {
		gvk := component.GetObjectKind().GroupVersionKind()
		if gvk.Group != "cluster.k8s.io" {
			continue
		}
		avk := apiVersionKind{}
		latestDetectedProvider := unknown
		switch gvk.Kind {
		case "Cluster":
			cluster := clusterv1alpha1.Cluster{}
			scheme.Convert(component, &cluster, nil)
			if err := json.Unmarshal(cluster.Spec.ProviderSpec.Value.Raw, &avk); err != nil {
				continue
			}
			gv, err := schema.ParseGroupVersion(avk.ApiVersion)
			if err != nil {
				continue
			}
			latestDetectedProvider = providerForGroup(gv.Group)
		case "MachineDeployment":
			md := clusterv1alpha1.MachineDeployment{}
			scheme.Convert(component, &md, nil)
			if err := json.Unmarshal(md.Spec.Template.Spec.ProviderSpec.Value.Raw, &avk); err != nil {
				continue
			}
			gv, err := schema.ParseGroupVersion(avk.ApiVersion)
			if err != nil {
				continue
			}
			latestDetectedProvider = providerForGroup(gv.Group)
		}
		if latestDetectedProvider == unknown {
			continue
		}
		if detectedProvider == unknown {
			detectedProvider = latestDetectedProvider
		}
		if latestDetectedProvider != detectedProvider {
			return unknown, errors.New("mixed providers detected. An infrastructure definition only supports creating resources for a single cluster-api provider")
		}
	}

	fmt.Printf("Detected provider: %s\n", detectedProvider.String())

	return detectedProvider, nil
}

func Deploy(kubeConfigPath string, components []runtime.Object) error {
	provider, err := decideProvider(components)
	if err != nil {
		return err
	}
	switch provider {
	case proxmox:
		return kubernetes.Apply(kubeConfigPath, proxmoxManifests)
	case openstack:
		return kubernetes.Apply(kubeConfigPath, openstackManifests)
	}
	return errors.New("it was not possible to detect a provider given the infrastructure definition")
}
