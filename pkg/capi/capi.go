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

import (
	"fmt"
	"time"

	"github.com/ereslibre/kube-ship/pkg/kubernetes"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	capiapis "sigs.k8s.io/cluster-api/pkg/apis"
	capiv1alpha1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	clusterapiclient "sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset"
)

func Deploy(kubeConfigPath string) error {
	return kubernetes.Apply(kubeConfigPath, manifests)
}

func NewScheme() *runtime.Scheme {
	scheme := runtime.NewScheme()
	v1.AddToScheme(scheme)
	capiapis.AddToScheme(scheme)
	return scheme
}

func numExpectedMachines(components []runtime.Object) int {
	scheme := NewScheme()

	expectedMachines := 0
	for _, component := range components {
		switch component.GetObjectKind().GroupVersionKind().Kind {
		case "MachineDeployment":
			md := capiv1alpha1.MachineDeployment{}
			scheme.Convert(component, &md, nil)
			expectedMachines += int(*md.Spec.Replicas)
		}
	}

	return expectedMachines
}

func numCurrentMachines(clusterAPIClient clusterapiclient.Interface) int {
	currentMachines, _ := clusterAPIClient.ClusterV1alpha1().Machines("kube-ship").List(metav1.ListOptions{})
	return len(currentMachines.Items)
}

func numRunningMachines(clusterAPIClient clusterapiclient.Interface) int {
	runningMachines := 0
	currentMachines, _ := clusterAPIClient.ClusterV1alpha1().Machines("kube-ship").List(metav1.ListOptions{})
	for _, machine := range currentMachines.Items {
		if (machine.Status.Phase != nil) && *machine.Status.Phase == "Running" {
			runningMachines += 1
		}
	}
	return runningMachines
}

func WaitForMachines(components []runtime.Object, kubeConfigPath string) error {
	expectedMachines := numExpectedMachines(components)
	clusterAPIClient, err := kubernetes.GetClusterAPIClient(kubeConfigPath)
	if err != nil {
		return err
	}
	for {
		currentMachineNumber := numCurrentMachines(clusterAPIClient)
		currentRunningMachineNumber := numRunningMachines(clusterAPIClient)
		fmt.Printf("Expecting %d machines; %d present (%d running). Waiting...\n", expectedMachines, currentMachineNumber, currentRunningMachineNumber)
		if currentRunningMachineNumber == int(expectedMachines) {
			break
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}
