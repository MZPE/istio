// Copyright 2018 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha3

import (
	xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"

	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pilot/pkg/networking/plugin"
)

type ConfigGeneratorImpl struct {
	// List of plugins that modify code generated by this config generator
	Plugins []plugin.Plugin
}

func NewConfigGenerator(plugins []plugin.Plugin) *ConfigGeneratorImpl {
	// TODO: Stick indices and other stuff here
	return &ConfigGeneratorImpl{
		Plugins: plugins,
	}
}

// TODO
func (_ *ConfigGeneratorImpl) BuildRoutes(env model.Environment, node model.Proxy, routeName string) ([]*xdsapi.RouteConfiguration, error) {
	return nil, nil
}
