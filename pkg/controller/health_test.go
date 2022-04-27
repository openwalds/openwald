/*
Copyright 2022 The OpenWald Authors.

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

package controller

import (
	"testing"

	"github.com/astaxie/beego"
)

func TestHealthController_HealthCheck(t *testing.T) {
	type fields struct {
		Controller beego.Controller
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "HealthCheck",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &HealthController{
				Controller: tt.fields.Controller,
			}
		})
	}
}
