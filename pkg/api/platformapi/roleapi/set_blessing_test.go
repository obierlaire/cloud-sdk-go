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

package roleapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestSetBlessings(t *testing.T) {
	type args struct {
		params SetBlessingsParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "fails on parameter validation",
			args: args{},
			err: multierror.NewPrefixed("invalid role set blessing params",
				errors.New("api reference is required for the operation"),
				errors.New("blessing definitions not specified and is required for this operation"),
				errors.New("id not specified and is required for this operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "fails updating the role",
			args: args{params: SetBlessingsParams{
				Region: "us-east-1",
				API: api.NewMock(mock.New500Response(mock.NewStringBody(
					`{"error": "failed updating role"}`,
				))),
				Blessings: &models.Blessings{},
				ID:        "one",
			}},
			err: errors.New(`{"error": "failed updating role"}`),
		},
		{
			name: "succeeds",
			args: args{params: SetBlessingsParams{
				Region: "us-east-1",
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "PUT",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/blueprinter/roles/one/blessings",
						Body:   mock.NewStringBody(`{"runner_ids_to_blessing":null}` + "\n"),
					},
					mock.NewStringBody(""),
				)),
				Blessings: &models.Blessings{},
				ID:        "one",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetBlessings(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
		})
	}
}
