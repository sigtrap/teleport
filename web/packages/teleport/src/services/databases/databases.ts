/**
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

import cfg, { UrlResourcesParams } from 'teleport/config';
import { ResourcesResponse } from 'teleport/services/agents';
import api from 'teleport/services/api';

import { withGenericUnsupportedError } from '../version/unsupported';
import {
  makeDatabase,
  makeDatabaseServer,
  makeDatabaseService,
} from './makeDatabase';
import type {
  CreateDatabaseRequest,
  Database,
  DatabaseIamPolicyResponse,
  DatabaseServer,
  DatabaseServicesResponse,
  UpdateDatabaseRequest,
} from './types';

class DatabaseService {
  fetchDatabases(
    clusterId: string,
    params: UrlResourcesParams,
    signal?: AbortSignal
  ): Promise<ResourcesResponse<Database>> {
    return api
      .get(cfg.getDatabasesUrl(clusterId, params), signal)
      .then(json => {
        const items = json?.items || [];

        return {
          agents: items.map(makeDatabase),
          startKey: json?.startKey,
          totalCount: json?.totalCount,
        };
      });
  }

  fetchDatabase(clusterId: string, dbName: string): Promise<Database> {
    return api.get(cfg.getDatabaseUrl(clusterId, dbName)).then(makeDatabase);
  }

  fetchDatabaseIamPolicy(
    clusterId: string,
    dbName: string
  ): Promise<DatabaseIamPolicyResponse> {
    return api.get(cfg.getDatabaseIamPolicyUrl(clusterId, dbName));
  }

  fetchDatabaseServices(clusterId: string): Promise<DatabaseServicesResponse> {
    return api.get(cfg.getDatabaseServicesUrl(clusterId)).then(json => {
      const items = json?.items || [];

      return {
        services: items.map(makeDatabaseService),
        startKey: json?.startKey,
        totalCount: json?.totalCount,
      };
    });
  }

  updateDatabase(
    clusterId: string,
    req: UpdateDatabaseRequest
  ): Promise<Database> {
    return api
      .put(cfg.getDatabaseUrl(clusterId, req.name), req)
      .then(makeDatabase);
  }

  createDatabase(
    clusterId: string,
    req: CreateDatabaseRequest
  ): Promise<Database> {
    return api.post(cfg.getDatabasesUrl(clusterId), req).then(makeDatabase);
  }
}

export default DatabaseService;

export function fetchDatabaseServers({
  clusterId,
  params,
  signal,
}: {
  clusterId: string;
  params: UrlResourcesParams;
  signal?: AbortSignal;
}): Promise<ResourcesResponse<DatabaseServer>> {
  return (
    api
      .get(cfg.getDatabaseServerUrl(clusterId, params), signal)
      .then(json => {
        const items = json?.items || [];

        return {
          agents: items.map(makeDatabaseServer),
          startKey: json?.startKey,
        };
      })
      // TODO(kimlisa): DELETE IN v19.0
      .catch(err => withGenericUnsupportedError(err, 'v18.0.0'))
  );
}
