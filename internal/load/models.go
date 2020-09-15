/*
Copyright © 2020 The Homeport Team

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package load

import (
	"time"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// KubeAccess contains Kubernetes cluster access objects in a single place
type KubeAccess struct {
	RestConfig *rest.Config
	Client     kubernetes.Interface
	DynClient  dynamic.Interface
}

// BuildRunResultSet is an aggregated result set based on multiple
// buildrun results
type BuildRunResultSet struct {
	NumberOfResults int

	Minimum BuildRunResult
	Maximum BuildRunResult
	Mean    BuildRunResult
	Median  BuildRunResult
}

// BuildRunResult contains the raw time results of a buildrun
type BuildRunResult struct {
	TotalBuildRunTime      time.Duration
	BuildRunRampUpDuration time.Duration
	TaskRunRampUpDuration  time.Duration
	PodRampUpDuration      time.Duration
	InternalProcessingTime time.Duration
}

// BuildRunSettings contains all required settings for a buildrun
type BuildRunSettings struct {
	// type settings
	ClusterBuildStrategy string
	BuildType            string

	// location and test resource naming settings
	Namespace string
	Prefix    string
	Name      string

	// build source settings
	SourceURL        string
	SourceContextDir string
	Dockerfile       string
	SourceSecret     string

	// build target settings
	TargetRegistryHostname  string
	TargetRegistryNamespace string
	TargetRegistrySecretRef string
}
