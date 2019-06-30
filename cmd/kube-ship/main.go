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

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	kindcluster "sigs.k8s.io/kind/pkg/cluster"
	kindapi "sigs.k8s.io/kind/pkg/cluster/config"
	"sigs.k8s.io/kind/pkg/cluster/config/encoding"
	kindapiv1alpha3 "sigs.k8s.io/kind/pkg/cluster/config/v1alpha3"
	kindclustercreate "sigs.k8s.io/kind/pkg/cluster/create"
	"sigs.k8s.io/kind/pkg/container/cri"

	"github.com/ereslibre/kube-ship/pkg/capi"
	"github.com/ereslibre/kube-ship/pkg/capiproviders"
	"github.com/ereslibre/kube-ship/pkg/kubernetes"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type cmdArgs struct {
	cfgFile string
}

var (
	args cmdArgs
)

func buildClusterConfig() kindapi.Cluster {
	versionedClusterCfg := kindapiv1alpha3.Cluster{
		Nodes: []kindapiv1alpha3.Node{
			{
				ExtraMounts: []cri.Mount{
					{
						ContainerPath: "/etc/resolv.conf",
						HostPath:      "/etc/resolv.conf",
						Readonly:      true,
					},
				},
			},
		},
	}
	clusterCfg := kindapi.Cluster{}
	encoding.Scheme.Convert(&versionedClusterCfg, &clusterCfg, nil)

	return clusterCfg
}

func createCluster(cfg *kindapi.Cluster) (*kindcluster.Context, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	kindContext := kindcluster.NewContext(uuid.String())
	return kindContext, kindContext.Create(cfg, kindclustercreate.Retain(true), kindclustercreate.WaitForReady(5*time.Minute), kindclustercreate.SetupKubernetes(true))
}

func infraConfig() (string, error) {
	rawConfig, err := ioutil.ReadFile(args.cfgFile)
	if err != nil {
		return "", errors.Wrap(err, "error reading infrastructure configuration file")
	}
	return string(rawConfig), nil
}

func doCreateCluster(cmd *cobra.Command, args []string) {
	infraCfg, err := infraConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating temporary Kind cluster")

	clusterCfg := buildClusterConfig()

	kindContext, err := createCluster(&clusterCfg)
	if err != nil {
		log.Fatal("could not create a cluster")
	}
	defer kindContext.Delete()

	kubeConfigPath := kindContext.KubeConfigPath()
	if err := capi.Deploy(kubeConfigPath); err != nil {
		log.Fatalf("could not deploy cluster-api: %v", err)
	}

	kubernetesClient, dynamicClient, err := kubernetes.GetKubernetesClient(kubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deploying components on top of the temporary Kind cluster")

	components, err := capiproviders.ListComponents(infraCfg)
	if err != nil {
		log.Fatalf("error listing components: %v", err)
	}
	if err := capi.Deploy(kubeConfigPath); err != nil {
		log.Fatalf("error deploying cluster-api: %v", err)
	}
	if err := capiproviders.Deploy(kubeConfigPath, components); err != nil {
		log.Fatalf("error deploying cluster-api provider: %v", err)
	}
	if err := kubernetes.CreateComponents(kubernetesClient, dynamicClient, components); err != nil {
		log.Fatalf("error creating components: %v", err)
	}
	if err := capi.WaitForMachines(components, kubeConfigPath); err != nil {
		log.Fatalf("error waiting for machines: %v", err)
	}

	fmt.Println("All ready! Cleaning up...")
}

func main() {
	var rootCmd = &cobra.Command{Use: os.Args[0]}
	rootCmd.PersistentFlags().StringVar(&args.cfgFile, "config", "", "infrastructure configuration file")
	rootCmd.MarkPersistentFlagRequired("config")
	rootCmd.Run = doCreateCluster
	rootCmd.Execute()
}
