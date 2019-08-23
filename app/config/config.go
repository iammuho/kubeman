// Copyright 2019 Muhammet Arslan <https://github.com/geass>
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

package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)


type (
	// Config provides the system configuration.
	Config struct {
		Logging     Logging
		Terraform     Terraform
		Kubernetes     Kubernetes
	}

	// Logging provides the logging configuration.
	Logging struct {
		Debug  bool `envconfig:"KUBEMAN_LOGS_DEBUG"`
		Trace  bool `envconfig:"KUBEMAN_LOGS_TRACE"`
		Color  bool `envconfig:"KUBEMAN_LOGS_COLOR"`
		Pretty bool `envconfig:"KUBEMAN_LOGS_PRETTY"`
		Text   bool `envconfig:"KUBEMAN_LOGS_TEXT"`
	}

	// Kubeman provides the application configuration.
	Terraform struct {
		Provider  string `envconfig:"KUBEMAN_TERRAFORM_PROVIDER"`
	}

	// kubernetes provides the kubernetes configuration.
	Kubernetes struct {
		MasterCount  int `envconfig:"KUBEMAN_KUBERNETES_MASTER_COUNT"`
		WorkerCount  bool `envconfig:"KUBEMAN_KUBERNETES_WORKER_COUNT"`
	}
)

// Environ returns the settings from the environment.
func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	return cfg, err
}

// String returns the configuration in string format.
func (c *Config) String() string {
	out, _ := yaml.Marshal(c)
	return string(out)
}
