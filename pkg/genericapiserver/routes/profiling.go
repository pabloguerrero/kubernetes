/*
Copyright 2014 The Kubernetes Authors.

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

package routes

import (
	"github.com/emicklei/go-restful"

	"k8s.io/kubernetes/pkg/apiserver"
	"net/http/pprof"
)

// Profiling adds handlers for pprof under /debug/pprof.
type Profiling struct{}

// Install adds the Profiling webservice to the given mux.
func (d Profiling) Install(mux *apiserver.PathRecorderMux, c *restful.Container) {
	mux.BaseMux().HandleFunc("/debug/pprof/", pprof.Index)
	mux.BaseMux().HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.BaseMux().HandleFunc("/debug/pprof/symbol", pprof.Symbol)
}
