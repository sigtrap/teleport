/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package version

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/coreos/go-semver/semver"
	"github.com/gravitational/trace"
	"github.com/stretchr/testify/require"

	basichttp2 "github.com/gravitational/teleport/lib/automaticupgrades/basichttp"
	"github.com/gravitational/teleport/lib/automaticupgrades/constants"
)

const basicHTTPTestPath = "/v1/cloud-stable"

func Test_basicHTTPVersionClient_Get(t *testing.T) {
	mock := basichttp2.NewServerMock(basicHTTPTestPath + "/" + constants.VersionPath)
	t.Cleanup(mock.Srv.Close)
	serverURL, err := url.Parse(mock.Srv.URL)
	serverURL.Path = basicHTTPTestPath
	require.NoError(t, err)
	ctx := context.Background()

	tests := []struct {
		name       string
		statusCode int
		response   string
		expected   *semver.Version
		assertErr  require.ErrorAssertionFunc
	}{
		{
			name:       "all good",
			statusCode: http.StatusOK,
			response:   "12.0.3",
			expected:   semver.Must(EnsureSemver("12.0.3")),
			assertErr:  require.NoError,
		},
		{
			name:       "all good with newline",
			statusCode: http.StatusOK,
			response:   "12.0.3\n",
			expected:   semver.Must(EnsureSemver("12.0.3")),
			assertErr:  require.NoError,
		},
		{
			name:       "all good with leading v",
			statusCode: http.StatusOK,
			response:   "v12.0.3",
			expected:   semver.Must(EnsureSemver("12.0.3")),
			assertErr:  require.NoError,
		},
		{
			name:       "non-semver",
			statusCode: http.StatusOK,
			response:   "hello",
			expected:   nil,
			assertErr: func(t2 require.TestingT, err2 error, _ ...interface{}) {
				require.IsType(t2, &trace.BadParameterError{}, trace.Unwrap(err2))
			},
		},
		{
			name:       "empty",
			statusCode: http.StatusOK,
			response:   "",
			expected:   nil,
			assertErr: func(t2 require.TestingT, err2 error, _ ...interface{}) {
				require.IsType(t2, &trace.BadParameterError{}, trace.Unwrap(err2))
			},
		},
		{
			name:       "non-200 response",
			statusCode: http.StatusInternalServerError,
			response:   "ERROR - SOMETHING WENT WRONG",
			expected:   nil,
			assertErr:  require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &basicHTTPVersionClient{
				baseURL: serverURL,
				client:  &basichttp2.Client{Client: mock.Srv.Client()},
			}
			mock.SetResponse(t, tt.statusCode, tt.response)
			result, err := b.Get(ctx)
			tt.assertErr(t, err)
			require.Equal(t, tt.expected, result)
		})
	}
}
