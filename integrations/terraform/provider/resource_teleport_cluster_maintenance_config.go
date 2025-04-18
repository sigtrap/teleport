// Code generated by _gen/main.go DO NOT EDIT
/*
Copyright 2015-2024 Gravitational, Inc.

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

package provider

import (
	"context"
	"fmt"
	 "math"

	apitypes "github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/integrations/lib/backoff"
	"github.com/gravitational/trace"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jonboulle/clockwork"

	"github.com/gravitational/teleport/integrations/terraform/tfschema"
)

// resourceTeleportClusterMaintenanceConfigType is the resource metadata type
type resourceTeleportClusterMaintenanceConfigType struct{}

// resourceTeleportClusterMaintenanceConfig is the resource
type resourceTeleportClusterMaintenanceConfig struct {
	p Provider
}

// GetSchema returns the resource schema
func (r resourceTeleportClusterMaintenanceConfigType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfschema.GenSchemaClusterMaintenanceConfigV1(ctx)
}

// NewResource creates the empty resource
func (r resourceTeleportClusterMaintenanceConfigType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceTeleportClusterMaintenanceConfig{
		p: *(p.(*Provider)),
	}, nil
}

// Create creates the ClusterMaintenanceConfig
func (r resourceTeleportClusterMaintenanceConfig) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.IsConfigured(resp.Diagnostics) {
		return
	}

	var plan types.Object
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterMaintenanceConfig := &apitypes.ClusterMaintenanceConfigV1{}
	diags = tfschema.CopyClusterMaintenanceConfigV1FromTerraform(ctx, plan, clusterMaintenanceConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := clusterMaintenanceConfig.CheckAndSetDefaults()
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error setting ClusterMaintenanceConfig defaults", trace.Wrap(err), "cluster_maintenance_config"))
		return
	}

	

	clusterMaintenanceConfigBefore, err := r.p.Client.GetClusterMaintenanceConfig(ctx)
	if err != nil && !trace.IsNotFound(err) {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
		return
	}

	if clusterMaintenanceConfigBefore == nil {
		clusterMaintenanceConfigBefore = &apitypes.ClusterMaintenanceConfigV1{}
	}
	clusterMaintenanceConfig = clusterMaintenanceConfig.WithNonce(math.MaxUint64).(*apitypes.ClusterMaintenanceConfigV1)

	err = r.p.Client.UpdateClusterMaintenanceConfig(ctx, clusterMaintenanceConfig)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error creating ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
		return
	}

	var clusterMaintenanceConfigI apitypes.ClusterMaintenanceConfig

	tries := 0
	backoff := backoff.NewDecorr(r.p.RetryConfig.Base, r.p.RetryConfig.Cap, clockwork.NewRealClock())
	for {
		tries = tries + 1
		clusterMaintenanceConfigI, err = r.p.Client.GetClusterMaintenanceConfig(ctx)
		if err != nil {
			resp.Diagnostics.Append(diagFromWrappedErr("Error reading ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
			return
		}

		previousMetadata := clusterMaintenanceConfigBefore.GetMetadata()
		currentMetadata := clusterMaintenanceConfigI.GetMetadata()
		if previousMetadata.GetRevision() != currentMetadata.GetRevision() || true {
			break
		}
		if bErr := backoff.Do(ctx); bErr != nil {
			resp.Diagnostics.Append(diagFromWrappedErr("Error reading ClusterMaintenanceConfig", trace.Wrap(bErr), "cluster_maintenance_config"))
			return
		}
		if tries >= r.p.RetryConfig.MaxTries {
			diagMessage := fmt.Sprintf("Error reading ClusterMaintenanceConfig (tried %d times) - state outdated, please import resource", tries)
			resp.Diagnostics.AddError(diagMessage, "cluster_maintenance_config")
			return
		}
	}
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
		return
	}

	clusterMaintenanceConfig, ok := clusterMaintenanceConfigI.(*apitypes.ClusterMaintenanceConfigV1)
	if !ok {
		resp.Diagnostics.Append(
			diagFromWrappedErr("Error reading ClusterMaintenanceConfig", trace.Errorf("Can not convert %T to ClusterMaintenanceConfigV1", clusterMaintenanceConfigI), "cluster_maintenance_config"),
		)
		return
	}

	diags = tfschema.CopyClusterMaintenanceConfigV1ToTerraform(ctx, clusterMaintenanceConfig, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Attrs["id"] = types.String{Value: "cluster_maintenance_config"}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read reads teleport ClusterMaintenanceConfig
func (r resourceTeleportClusterMaintenanceConfig) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state types.Object
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterMaintenanceConfigI, err := r.p.Client.GetClusterMaintenanceConfig(ctx)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
		return
	}

	
	clusterMaintenanceConfig := clusterMaintenanceConfigI.(*apitypes.ClusterMaintenanceConfigV1)
	diags = tfschema.CopyClusterMaintenanceConfigV1ToTerraform(ctx, clusterMaintenanceConfig, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates teleport ClusterMaintenanceConfig
func (r resourceTeleportClusterMaintenanceConfig) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	if !r.p.IsConfigured(resp.Diagnostics) {
		return
	}

	var plan types.Object
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterMaintenanceConfig := &apitypes.ClusterMaintenanceConfigV1{}
	diags = tfschema.CopyClusterMaintenanceConfigV1FromTerraform(ctx, plan, clusterMaintenanceConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := clusterMaintenanceConfig.CheckAndSetDefaults()
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error updating ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
		return
	}

	clusterMaintenanceConfigBefore, err := r.p.Client.GetClusterMaintenanceConfig(ctx)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
		return
	}
	clusterMaintenanceConfig = clusterMaintenanceConfig.WithNonce(math.MaxUint64).(*apitypes.ClusterMaintenanceConfigV1)

	err = r.p.Client.UpdateClusterMaintenanceConfig(ctx, clusterMaintenanceConfig)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error updating ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
		return
	}
	var clusterMaintenanceConfigI apitypes.ClusterMaintenanceConfig

	tries := 0
	backoff := backoff.NewDecorr(r.p.RetryConfig.Base, r.p.RetryConfig.Cap, clockwork.NewRealClock())
	for {
		tries = tries + 1
		clusterMaintenanceConfigI, err = r.p.Client.GetClusterMaintenanceConfig(ctx)
		if err != nil {
			resp.Diagnostics.Append(diagFromWrappedErr("Error reading ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
			return
		}
		if clusterMaintenanceConfigBefore.GetMetadata().Revision != clusterMaintenanceConfigI.GetMetadata().Revision || true {
			break
		}
		if bErr := backoff.Do(ctx); bErr != nil {
			resp.Diagnostics.Append(diagFromWrappedErr("Error reading ClusterMaintenanceConfig", trace.Wrap(bErr), "cluster_maintenance_config"))
			return
		}
		if tries >= r.p.RetryConfig.MaxTries {
			diagMessage := fmt.Sprintf("Error reading ClusterMaintenanceConfig (tried %d times) - state outdated, please import resource", tries)
			resp.Diagnostics.AddError(diagMessage, "cluster_maintenance_config")
			return
		}
	}
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
		return
	}

	
	clusterMaintenanceConfig = clusterMaintenanceConfigI.(*apitypes.ClusterMaintenanceConfigV1)
	diags = tfschema.CopyClusterMaintenanceConfigV1ToTerraform(ctx, clusterMaintenanceConfig, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes Teleport ClusterMaintenanceConfig
func (r resourceTeleportClusterMaintenanceConfig) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	err := r.p.Client.DeleteClusterMaintenanceConfig(ctx)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error deleting ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
		return
	}

	resp.State.RemoveResource(ctx)
}

// ImportState imports ClusterMaintenanceConfig state
func (r resourceTeleportClusterMaintenanceConfig) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	clusterMaintenanceConfigI, err := r.p.Client.GetClusterMaintenanceConfig(ctx)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error updating ClusterMaintenanceConfig", trace.Wrap(err), "cluster_maintenance_config"))
		return
	}
	clusterMaintenanceConfig := clusterMaintenanceConfigI.(*apitypes.ClusterMaintenanceConfigV1)

	var state types.Object

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = tfschema.CopyClusterMaintenanceConfigV1ToTerraform(ctx, clusterMaintenanceConfig, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	id := clusterMaintenanceConfig.GetName()

	state.Attrs["id"] = types.String{Value: id}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
