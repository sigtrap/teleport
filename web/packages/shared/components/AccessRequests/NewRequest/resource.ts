/**
 * Teleport
 * Copyright (C) 2024 Gravitational, Inc.
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

// eslint-disable-next-line no-restricted-imports -- FIXME
import { ResourceIdKind } from 'teleport/services/agents';
// eslint-disable-next-line no-restricted-imports -- FIXME
import { KubeResourceKind } from 'teleport/services/kube';

/** Available request kinds for resource-based and role-based access requests. */
export type RequestableResourceKind =
  | ResourceIdKind
  | 'role'
  | 'resource'
  | Exclude<KubeResourceKind, '*'>;

/**
 * Maps a resource ID (usually agent name) to resource description (usually the
 * same, but not necessarily).
 */
export type ResourceMap = {
  [K in Exclude<RequestableResourceKind, 'resource'>]: Record<string, string>;
};

export function getEmptyResourceState(): ResourceMap {
  return {
    node: {},
    db: {},
    app: {},
    kube_cluster: {},
    user_group: {},
    windows_desktop: {},
    role: {},
    saml_idp_service_provider: {},
    namespace: {},
    aws_ic_account_assignment: {},
    git_server: {},
  };
}
