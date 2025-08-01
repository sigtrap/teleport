/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
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

package agent

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gravitational/trace"
	"gopkg.in/yaml.v3"

	"github.com/gravitational/teleport/lib/autoupdate"
)

const (
	// defaultSetting is used to represent the default value for updater config.
	defaultSetting = "default"
)

const (
	// updateConfigName specifies the name of the file inside versionsDirName containing configuration for the teleport update.
	updateConfigName = "update.yaml"

	// UpdateConfig metadata
	updateConfigVersion = "v1"
	updateConfigKind    = "update_config"
)

// UpdateConfig describes the update.yaml file schema.
// Pointer values should be replaced and not mutated, see Copy.
type UpdateConfig struct {
	// Version of the configuration file
	Version string `yaml:"version"`
	// Kind of configuration file (always "update_config")
	Kind string `yaml:"kind"`
	// Spec contains user-specified configuration.
	Spec UpdateSpec `yaml:"spec"`
	// Status contains state configuration.
	Status UpdateStatus `yaml:"status"`
}

// Copy an UpdateConfig. Pointers are not copied.
func (cfg *UpdateConfig) Copy() *UpdateConfig {
	if cfg == nil {
		return &UpdateConfig{}
	}
	return toPtr(*cfg)
}

// UpdateSpec describes the spec field in update.yaml.
// Pointer values should be replaced and not mutated, see Copy.
type UpdateSpec struct {
	// Proxy address
	Proxy string `yaml:"proxy"`
	// Path is the location the Teleport binaries are linked into.
	Path string `yaml:"path"`
	// Group specifies the update group identifier for the agent.
	Group string `yaml:"group,omitempty"`
	// BaseURL is CDN base URL used for the Teleport tgz download URL.
	BaseURL string `yaml:"base_url,omitempty"`
	// Enabled controls whether auto-updates are enabled.
	Enabled bool `yaml:"enabled"`
	// Pinned controls whether the active_version is pinned.
	Pinned bool `yaml:"pinned"`
	// SELinuxSSH controls whether an SELinux module will be installed to
	// constrain Teleport SSH.
	SELinuxSSH bool `yaml:"selinux_ssh,omitempty"`
}

// UpdateStatus describes the status field in update.yaml.
// Pointer values should be replaced and not mutated, see Copy.
type UpdateStatus struct {
	// IDFile is the path to a temporary file containing the updater ID.
	IDFile string `yaml:"id_file,omitempty"`
	// LastUpdate status, if attempted
	LastUpdate *LastUpdate `yaml:"last_update,omitempty"`
	// Active is the currently active revision of Teleport.
	Active Revision `yaml:"active"`
	// Backup is the last working revision of Teleport.
	Backup *Revision `yaml:"backup,omitempty"`
	// Skip is the skipped revision of Teleport.
	// Skipped revisions are not applied because they
	// are known to crash.
	Skip *Revision `yaml:"skip,omitempty"`
}

// LastUpdate describes the last attempted updated.
type LastUpdate struct {
	// Success or failure of the attempted update
	Success bool `yaml:"success"`
	// Time the update occurred
	Time time.Time `yaml:"time"`
	// Target revision for the update
	Target Revision `yaml:"target"`
}

// Revision is a version and edition of Teleport.
type Revision struct {
	// Version is the version of Teleport.
	Version string `yaml:"version" json:"version"`
	// Flags describe the edition of Teleport.
	Flags autoupdate.InstallFlags `yaml:"flags,flow,omitempty" json:"flags,omitempty"`
}

// NewRevision create a Revision.
// If version is not set, no flags are returned.
// This ensures that all Revisions without versions are zero-valued.
func NewRevision(version string, flags autoupdate.InstallFlags) Revision {
	if version != "" {
		return Revision{
			Version: version,
			Flags:   flags,
		}
	}
	return Revision{}
}

// NewRevisionFromDir translates a directory path containing Teleport into a Revision.
func NewRevisionFromDir(dir string) (Revision, error) {
	parts := strings.Split(dir, "_")
	var out Revision
	if len(parts) == 0 {
		return out, trace.Errorf("dir name empty")
	}
	out.Version = parts[0]
	if out.Version == "" {
		return out, trace.Errorf("version missing in dir %s", dir)
	}
	switch flags := parts[1:]; len(flags) {
	case 2:
		if flags[1] != autoupdate.FlagFIPS.DirFlag() {
			break
		}
		out.Flags |= autoupdate.FlagFIPS
		fallthrough
	case 1:
		if flags[0] != autoupdate.FlagEnterprise.DirFlag() {
			break
		}
		out.Flags |= autoupdate.FlagEnterprise
		fallthrough
	case 0:
		return out, nil
	}
	return out, trace.Errorf("invalid flag in %s", dir)
}

// Dir returns the directory path name of a Revision.
// These are unambiguous for semver and may be parsed with NewRevisionFromDir.
func (r Revision) Dir() string {
	// Do not change the order of these statements.
	// Otherwise, installed versions will no longer match update.yaml.
	var suffix string
	if r.Flags&(autoupdate.FlagEnterprise|autoupdate.FlagFIPS) != 0 {
		suffix += "_" + autoupdate.FlagEnterprise.DirFlag()
	}
	if r.Flags&autoupdate.FlagFIPS != 0 {
		suffix += "_" + autoupdate.FlagFIPS.DirFlag()
	}
	return r.Version + suffix
}

// String returns a human-readable description of a Teleport revision.
// These are semver-ambiguous and should not be parsed.
func (r Revision) String() string {
	if flags := r.Flags.Strings(); len(flags) > 0 {
		return fmt.Sprintf("%s+%s", r.Version, strings.Join(flags, "+"))
	}
	return r.Version
}

// readConfig reads UpdateConfig from a file.
func readConfig(path string) (*UpdateConfig, error) {
	f, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return &UpdateConfig{
			Version: updateConfigVersion,
			Kind:    updateConfigKind,
		}, nil
	}
	if err != nil {
		return nil, trace.Wrap(err, "failed to open")
	}
	defer f.Close()
	var cfg UpdateConfig
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, trace.Wrap(err, "failed to parse")
	}
	if k := cfg.Kind; k != updateConfigKind {
		return nil, trace.Errorf("invalid kind %s", k)
	}
	if v := cfg.Version; v != updateConfigVersion {
		return nil, trace.Errorf("invalid version %s", v)
	}
	return &cfg, nil
}

// writeConfig writes UpdateConfig to a file atomically, ensuring the file cannot be corrupted.
func writeConfig(filename string, cfg *UpdateConfig) error {
	return trace.Wrap(writeAtomicWithinDir(filename, configFileMode, func(w io.Writer) error {
		return trace.Wrap(yaml.NewEncoder(w).Encode(cfg))
	}))
}

func updateConfigSpec(spec *UpdateSpec, override OverrideConfig) error {
	if override.Proxy != "" {
		spec.Proxy = override.Proxy
	}
	if override.Path != "" {
		spec.Path = override.Path
	}
	spec.Group = overrideOptional(spec.Group, override.Group)
	spec.BaseURL = overrideOptional(spec.BaseURL, override.BaseURL)
	if spec.BaseURL != "" &&
		!strings.HasPrefix(strings.ToLower(spec.BaseURL), "https://") {
		return trace.Errorf("Teleport download base URL %s must use TLS (https://)", spec.BaseURL)
	}
	if override.Enabled {
		spec.Enabled = true
	}
	if override.Pinned {
		spec.Pinned = true
	}
	if override.SELinuxSSHChanged {
		spec.SELinuxSSH = override.SELinuxSSH
	}
	if spec.SELinuxSSH && runtime.GOOS != "linux" {
		return trace.BadParameter("SELinux is only supported on Linux")
	}
	return nil
}

func overrideOptional(orig, override string) string {
	switch override {
	case "":
		return orig
	case defaultSetting:
		return ""
	default:
		return override
	}
}

// Status of the agent auto-updates system.
type Status struct {
	UpdateSpec   `yaml:",inline"`
	UpdateStatus `yaml:",inline"`
	FindResp     `yaml:",inline"`
	// ID is the updater ID.
	ID string `yaml:"id,omitempty"`
}

// FindResp summarizes the auto-update status response from cluster.
type FindResp struct {
	// Target revision of Teleport to install
	Target Revision `yaml:"target"`
	// InWindow is true when the install should happen now.
	InWindow bool `yaml:"in_window"`
	// Jitter duration before an automated install
	Jitter time.Duration `yaml:"jitter"`
	// AGPL installations cannot use the official CDN.
	AGPL bool `yaml:"agpl,omitempty"`
}
