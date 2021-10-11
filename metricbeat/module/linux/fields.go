// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package linux

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "linux", asset.ModuleFieldsPri, AssetLinux); err != nil {
		panic(err)
	}
}

// AssetLinux returns asset data.
// This is the base64 encoded zlib format compressed contents of module/linux.
func AssetLinux() string {
	return "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"
}
