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

package main

import (
	"github.com/kubeman/app/config"
	"github.com/kubeman/pkg/kubeman"
	"github.com/kubeman/pkg/terraform"

	"github.com/google/wire"
)

// wire set for kubeman.
var kubemanSet = wire.NewSet(
	provideKubeman,
)

// provideKubeman is a Wire provider function that returns an
// kubeman application that is configured from the environment.
func provideKubeman(terraform *terraform.Terraform, config config.Config) *kubeman.Kubeman {
	return &kubeman.Kubeman{
		Terraform:    terraform,
	}
}

