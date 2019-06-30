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

package kubernetes

import (
	"fmt"
	"io"
	"os/exec"

	v1 "k8s.io/api/core/v1"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicclient "k8s.io/client-go/dynamic"
	kubernetesclient "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clusterapiclient "sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset"
)

func Apply(kubeConfigPath string, manifests string) error {
	cmd := exec.Command("kubectl", fmt.Sprintf("--kubeconfig=%s", kubeConfigPath), "apply", "-f=-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, manifests)
	}()

	return cmd.Run()
}

func GetKubernetesClient(kubeConfigPath string) (kubernetesclient.Interface, dynamicclient.Interface, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, nil, err
	}
	kubernetesClient, err := kubernetesclient.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}
	dynamicClient, err := dynamicclient.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}
	return kubernetesClient, dynamicClient, nil
}

func GetClusterAPIClient(kubeConfigPath string) (clusterapiclient.Interface, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}
	return clusterapiclient.NewForConfig(config)
}

func CreateComponents(kubernetesClient kubernetesclient.Interface, dynamicClient dynamicclient.Interface, components []runtime.Object) error {
	kubernetesClient.CoreV1().Namespaces().Create(&v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "kube-ship",
		},
	})
	for _, component := range components {
		gvk := component.GetObjectKind().GroupVersionKind()
		gvr, _ := apimeta.UnsafeGuessKindToResource(gvk)
		unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(component)
		if err != nil {
			return err
		}
		unstructuredFinalObj := unstructured.Unstructured{unstructuredObj}
		namespace := unstructuredFinalObj.GetNamespace()
		if len(namespace) == 0 {
			namespace = "kube-ship"
		}
		if _, err := dynamicClient.Resource(gvr).Namespace(namespace).Create(&unstructuredFinalObj, metav1.CreateOptions{}); err != nil {
			return err
		}
	}
	return nil
}

func WaitForComponents(clusterAPIClient clusterapiclient.Interface, components []runtime.Object) error {

	return nil
}
