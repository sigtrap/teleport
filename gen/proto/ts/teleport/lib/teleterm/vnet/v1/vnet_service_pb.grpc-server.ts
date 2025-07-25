/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter eslint_disable,add_pb_suffix,server_grpc1,ts_nocheck
// @generated from protobuf file "teleport/lib/teleterm/vnet/v1/vnet_service.proto" (package "teleport.lib.teleterm.vnet.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
//
// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
import { AutoConfigureSSHResponse } from "./vnet_service_pb";
import { AutoConfigureSSHRequest } from "./vnet_service_pb";
import { RunDiagnosticsResponse } from "./vnet_service_pb";
import { RunDiagnosticsRequest } from "./vnet_service_pb";
import { GetBackgroundItemStatusResponse } from "./vnet_service_pb";
import { GetBackgroundItemStatusRequest } from "./vnet_service_pb";
import { GetServiceInfoResponse } from "./vnet_service_pb";
import { GetServiceInfoRequest } from "./vnet_service_pb";
import { StopResponse } from "./vnet_service_pb";
import { StopRequest } from "./vnet_service_pb";
import { StartResponse } from "./vnet_service_pb";
import { StartRequest } from "./vnet_service_pb";
import type * as grpc from "@grpc/grpc-js";
/**
 * VnetService provides methods to manage a VNet instance.
 *
 * @generated from protobuf service teleport.lib.teleterm.vnet.v1.VnetService
 */
export interface IVnetService extends grpc.UntypedServiceImplementation {
    /**
     * Start starts VNet.
     *
     * @generated from protobuf rpc: Start(teleport.lib.teleterm.vnet.v1.StartRequest) returns (teleport.lib.teleterm.vnet.v1.StartResponse);
     */
    start: grpc.handleUnaryCall<StartRequest, StartResponse>;
    /**
     * Stop stops VNet.
     *
     * @generated from protobuf rpc: Stop(teleport.lib.teleterm.vnet.v1.StopRequest) returns (teleport.lib.teleterm.vnet.v1.StopResponse);
     */
    stop: grpc.handleUnaryCall<StopRequest, StopResponse>;
    /**
     * GetServiceInfo returns info about the running VNet service.
     *
     * @generated from protobuf rpc: GetServiceInfo(teleport.lib.teleterm.vnet.v1.GetServiceInfoRequest) returns (teleport.lib.teleterm.vnet.v1.GetServiceInfoResponse);
     */
    getServiceInfo: grpc.handleUnaryCall<GetServiceInfoRequest, GetServiceInfoResponse>;
    /**
     * GetBackgroundItemStatus returns the status of the background item responsible for launching
     * VNet daemon. macOS only. tsh must be compiled with the vnetdaemon build tag.
     *
     * @generated from protobuf rpc: GetBackgroundItemStatus(teleport.lib.teleterm.vnet.v1.GetBackgroundItemStatusRequest) returns (teleport.lib.teleterm.vnet.v1.GetBackgroundItemStatusResponse);
     */
    getBackgroundItemStatus: grpc.handleUnaryCall<GetBackgroundItemStatusRequest, GetBackgroundItemStatusResponse>;
    /**
     * RunDiagnostics runs a set of heuristics to determine if VNet actually works on the device, that
     * is receives network traffic and DNS queries. RunDiagnostics requires VNet to be started.
     *
     * @generated from protobuf rpc: RunDiagnostics(teleport.lib.teleterm.vnet.v1.RunDiagnosticsRequest) returns (teleport.lib.teleterm.vnet.v1.RunDiagnosticsResponse);
     */
    runDiagnostics: grpc.handleUnaryCall<RunDiagnosticsRequest, RunDiagnosticsResponse>;
    /**
     * AutoConfigureSSH automatically configures OpenSSH-compatible clients for
     * connections to Teleport SSH hosts.
     *
     * @generated from protobuf rpc: AutoConfigureSSH(teleport.lib.teleterm.vnet.v1.AutoConfigureSSHRequest) returns (teleport.lib.teleterm.vnet.v1.AutoConfigureSSHResponse);
     */
    autoConfigureSSH: grpc.handleUnaryCall<AutoConfigureSSHRequest, AutoConfigureSSHResponse>;
}
/**
 * @grpc/grpc-js definition for the protobuf service teleport.lib.teleterm.vnet.v1.VnetService.
 *
 * Usage: Implement the interface IVnetService and add to a grpc server.
 *
 * ```typescript
 * const server = new grpc.Server();
 * const service: IVnetService = ...
 * server.addService(vnetServiceDefinition, service);
 * ```
 */
export const vnetServiceDefinition: grpc.ServiceDefinition<IVnetService> = {
    start: {
        path: "/teleport.lib.teleterm.vnet.v1.VnetService/Start",
        originalName: "Start",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => StartResponse.fromBinary(bytes),
        requestDeserialize: bytes => StartRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(StartResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(StartRequest.toBinary(value))
    },
    stop: {
        path: "/teleport.lib.teleterm.vnet.v1.VnetService/Stop",
        originalName: "Stop",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => StopResponse.fromBinary(bytes),
        requestDeserialize: bytes => StopRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(StopResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(StopRequest.toBinary(value))
    },
    getServiceInfo: {
        path: "/teleport.lib.teleterm.vnet.v1.VnetService/GetServiceInfo",
        originalName: "GetServiceInfo",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => GetServiceInfoResponse.fromBinary(bytes),
        requestDeserialize: bytes => GetServiceInfoRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(GetServiceInfoResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(GetServiceInfoRequest.toBinary(value))
    },
    getBackgroundItemStatus: {
        path: "/teleport.lib.teleterm.vnet.v1.VnetService/GetBackgroundItemStatus",
        originalName: "GetBackgroundItemStatus",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => GetBackgroundItemStatusResponse.fromBinary(bytes),
        requestDeserialize: bytes => GetBackgroundItemStatusRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(GetBackgroundItemStatusResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(GetBackgroundItemStatusRequest.toBinary(value))
    },
    runDiagnostics: {
        path: "/teleport.lib.teleterm.vnet.v1.VnetService/RunDiagnostics",
        originalName: "RunDiagnostics",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => RunDiagnosticsResponse.fromBinary(bytes),
        requestDeserialize: bytes => RunDiagnosticsRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(RunDiagnosticsResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(RunDiagnosticsRequest.toBinary(value))
    },
    autoConfigureSSH: {
        path: "/teleport.lib.teleterm.vnet.v1.VnetService/AutoConfigureSSH",
        originalName: "AutoConfigureSSH",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => AutoConfigureSSHResponse.fromBinary(bytes),
        requestDeserialize: bytes => AutoConfigureSSHRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(AutoConfigureSSHResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(AutoConfigureSSHRequest.toBinary(value))
    }
};
