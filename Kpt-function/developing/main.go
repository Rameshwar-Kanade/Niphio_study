// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"os"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)

var _ fn.Runner = &YourFunction{}

// TODO: Change to your functionConfig "Kind" name.
type YourFunction struct {
	FnConfigBool bool
	FnConfigInt  int
	FnConfigFoo  string
}

// Run is the main function logic.
// `items` is parsed from the STDIN "ResourceList.Items".
// `functionConfig` is from the STDIN "ResourceList.FunctionConfig". The value has been assigned to the r attributes
// `results` is the "ResourceList.Results" that you can write result info to.
func (r *YourFunction) Run(ctx *fn.Context, functionConfig *fn.KubeObject, items fn.KubeObjects, results *fn.Results) bool {
	for _, kubeObject := range items {
		if kubeObject.IsGVK("app", "v1", "Deployment") { // Corrected Group/Version/Kind
			kubeObject.SetAnnotation("k8s.v1.cni.cncf.io/networks", "kpt")
			kubeObject.SetAPIVersion("apps/v1")
			kubeObject.SetKind("Deployment")
			kubeObject.SetLabel("app", "my-app")

			arrMaps := map[string]int{
				"replicas": 3,
			}
			kubeObject.SetNestedField(arrMaps, "spec")
			kubeObject.SetNestedStringMap(map[string]string{"app": "my-app"}, "spec", "selector", "matchLabels")
			kubeObject.SetNestedStringMap(map[string]string{"app": "my-app", "name": "my-pod"}, "spec", "template", "metadata", "labels")
			kubeObject.SetNestedField([]map[string]interface{}{{"name": "my-container", "image": "nginx:latest"}}, "spec", "template", "spec", "containers")
			kubeObject.RemoveNestedField("data")
		}
	}
	return true // Indicate successful processing
}

func main() {
	runner := fn.WithContext(context.Background(), &YourFunction{})
	if err := fn.AsMain(runner); err != nil {
		os.Exit(1)
	}
}
