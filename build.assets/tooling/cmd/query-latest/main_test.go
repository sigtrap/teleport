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

package main

import (
	"context"
	"testing"

	"github.com/google/go-github/v41/github"
	"github.com/stretchr/testify/require"
)

func TestGetLatest(t *testing.T) {
	tests := []struct {
		desc     string
		spec     string
		releases map[string][]string
		wantErr  require.ErrorAssertionFunc
		latest   string
	}{
		{
			desc: "pass",
			spec: "v8",
			releases: map[string][]string{
				"teleport":         {"v8.1.9", "v8.1.10", "v8.0.11"},
				"teleport-private": {"v1.0.1"},
			},
			wantErr: require.NoError,
			latest:  "v8.1.10",
		},
		{
			desc: "fail-bad-spec",
			spec: "v9",
			releases: map[string][]string{
				"teleport":         {"v8.1.9", "v8.1.10", "v8.0.11"},
				"teleport-private": {"v1.0.1"},
			},
			wantErr: require.Error,
			latest:  "",
		},
		{
			desc: "pass-prerelease",
			spec: "v8",
			releases: map[string][]string{
				"teleport":         {"v8.1.10-rc.1", "v8.1.10", "v8.1.10-alpha.1"},
				"teleport-private": {"v1.0.1"},
			},
			wantErr: require.NoError,
			latest:  "v8.1.10",
		},
		{
			desc: "pass-major-minor",
			spec: "v8.1",
			releases: map[string][]string{
				"teleport":         {"v8.1.9", "v8.2.1", "v8.1.10", "v8.0.11"},
				"teleport-private": {"v1.0.1"},
			},
			wantErr: require.NoError,
			latest:  "v8.1.10",
		},
		{
			desc: "latest-private",
			spec: "v8",
			releases: map[string][]string{
				"teleport":         {"v8.1.9", "v8.2.1", "v8.1.10", "v8.0.11"},
				"teleport-private": {"v8.2.2"},
			},
			wantErr: require.NoError,
			latest:  "private-v8.2.2",
		},
		{
			desc: "latest-not-private",
			spec: "v8",
			releases: map[string][]string{
				"teleport":         {"v8.1.9", "v8.2.3", "v8.1.10", "v8.0.11"},
				"teleport-private": {"v8.2.2"},
			},
			wantErr: require.NoError,
			latest:  "v8.2.3",
		},
		{
			desc: "no-private-releases",
			spec: "v8",
			releases: map[string][]string{
				"teleport":         {"v8.1.9", "v8.2.3", "v8.1.10", "v8.0.11"},
				"teleport-private": {},
			},
			wantErr: require.Error,
			latest:  "",
		},
		{
			desc: "no-matching-private",
			spec: "v8",
			releases: map[string][]string{
				"teleport":         {"v8.1.9", "v8.2.3", "v8.1.10", "v8.0.11"},
				"teleport-private": {"v7.5.2"},
			},
			wantErr: require.NoError,
			latest:  "v8.2.3",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gh := &fakeGitHub{
				releases: test.releases,
			}
			latest, err := getLatest(context.Background(), test.spec, gh)
			test.wantErr(t, err)
			require.Equal(t, test.latest, latest)
		})
	}

}

type fakeGitHub struct {
	releases map[string][]string
}

func (f *fakeGitHub) ListReleases(ctx context.Context, organization, repository string) ([]github.RepositoryRelease, error) {
	ghReleases := make([]github.RepositoryRelease, 0)
	for _, r := range f.releases[repository] {
		tag := r
		ghReleases = append(ghReleases, github.RepositoryRelease{TagName: &tag})
	}
	return ghReleases, nil
}
