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

import 'whatwg-fetch';

import auth, { MfaChallengeScope } from 'teleport/services/auth/auth';
import websession from 'teleport/services/websession';

import { MfaChallengeResponse } from '../mfa';
import { storageService } from '../storageService';
import parseError, { ApiError, parseProxyVersion } from './parseError';

export const MFA_HEADER = 'Teleport-Mfa-Response';

type RequestOptions = {
  /**
   * Usually, an HTTP/404 with a "role not found" message means that the user
   * can't be authorized and needs to sign in again. In such case the API
   * service immediately signs the user out. Setting this flag to `true`
   * overrides this behavior and allows a "role not found" error to be
   * propagated up the call stack.
   */
  allowRoleNotFound?: boolean;

  /**
   * If set to `true`, the API service will not attempt to retry after an MFA
   * challenge.
   */
  skipAuthnRetry?: boolean;
};

const api = {
  get(
    url: string,
    abortSignal?: AbortSignal,
    mfaResponse?: MfaChallengeResponse,
    options?: RequestOptions
  ) {
    return api.fetchJsonWithMfaAuthnRetry(
      url,
      { signal: abortSignal },
      mfaResponse,
      options
    );
  },

  post(url, data?, abortSignal?, mfaResponse?: MfaChallengeResponse) {
    return api.fetchJsonWithMfaAuthnRetry(
      url,
      {
        body: JSON.stringify(data),
        method: 'POST',
        signal: abortSignal,
      },
      mfaResponse
    );
  },

  postFormData(url, formData, mfaResponse?: MfaChallengeResponse) {
    if (formData instanceof FormData) {
      return api.fetchJsonWithMfaAuthnRetry(
        url,
        {
          body: formData,
          method: 'POST',
          // Overrides the default header from `defaultRequestOptions`.
          headers: {
            Accept: 'application/json',
            // Let the browser infer the content-type for FormData types
            // to set the correct boundary:
            // 1) https://developer.mozilla.org/en-US/docs/Web/API/FormData/Using_FormData_Objects#sending_files_using_a_formdata_object
            // 2) https://stackoverflow.com/a/64653976
          },
        },
        mfaResponse
      );
    }

    throw new Error('data for body is not a type of FormData');
  },

  delete(url, data?, mfaResponse?: MfaChallengeResponse) {
    return api.fetchJsonWithMfaAuthnRetry(
      url,
      {
        body: JSON.stringify(data),
        method: 'DELETE',
      },
      mfaResponse
    );
  },

  deleteWithHeaders(
    url,
    headers?: Record<string, string>,
    signal?,
    mfaResponse?: MfaChallengeResponse
  ) {
    return api.fetchJsonWithMfaAuthnRetry(
      url,
      {
        method: 'DELETE',
        headers,
        signal,
      },
      mfaResponse
    );
  },

  // TODO (avatus) add abort signal to this
  put(url, data, mfaResponse?: MfaChallengeResponse) {
    return api.fetchJsonWithMfaAuthnRetry(
      url,
      {
        body: JSON.stringify(data),
        method: 'PUT',
      },
      mfaResponse
    );
  },

  putWithHeaders(
    url,
    data,
    headers?: Record<string, string>,
    mfaResponse?: MfaChallengeResponse
  ) {
    return api.fetchJsonWithMfaAuthnRetry(
      url,
      {
        body: JSON.stringify(data),
        method: 'PUT',
        headers,
      },
      mfaResponse
    );
  },

  /**
   * fetchJsonWithMfaAuthnRetry calls on `api.fetch` and
   * processes the response.
   *
   * It returns the JSON data if it is a valid JSON and
   * there were no response errors.
   *
   * The field "mfaResponse" accepts a pre-made response to an MFA challenge
   * for an admin action with allowReuse set to true.
   *
   * The cluster requires the user to re-authenticate before performing certain
   * admin actions (e.g., creating join tokens or deleting users) when
   * second_factor is set to webauthn.
   *
   * Generally, mfaResponse is not needed because this func will first attempt
   * a fetch without it and if the response had an error and it contained a MFA
   * authn required message, then this func will fetch a challenge and prompt
   * the user to re-authn. After successfully re-authenticating, a retry fetch
   * with the original URL is attempted.
   *
   * There are a few cases where providing a mfaResponse is required and it
   * starts by creating a MFA challenge with `allowReuse` set to true
   * (see services/auth/auth.ts > getMfaChallengeResponseForAdminAction).
   *
   * This allow users to require re-authenticating ONCE for the following:
   *   - In the web app, we can be calling multiple endpoints back to back,
   *     and some or all endpoints require re-authenticating. If a non reusable
   *     mfaResponse was provided, or mfaResponse wasn't provided then each
   *     fetch will ask the user to re-authenticate.
   *   - We can make a single fetch without a mfaResponse, re-authenticate
   *     successfully as required, but the retry attempt still fails with a
   *     vague "access denied" error. This is because there are some endpoints
   *     where it's in the backend that are calling multiple endpoints that
   *     require re-authenticating. Providing a reusable mfaResponse resolves
   *     this issue.
   *
   * The mfaResponse lasts for WebauthnChallengeTimeout defined in:
   * https://github.com/gravitational/teleport/blob/b8a65486844b4125ea1cf1f08ae17e2fc5a4db5a/lib/defaults/defaults.go#L598
   */
  async fetchJsonWithMfaAuthnRetry(
    url: string,
    customOptions: RequestInit,
    mfaResponse?: MfaChallengeResponse,
    options: RequestOptions = {}
  ): Promise<any> {
    try {
      const response = await api.fetch(url, customOptions, mfaResponse);
      return await api.getJsonFromFetchResponse(response, options);
    } catch (err) {
      // Retry with MFA if we get an admin action MFA error.
      if (
        !mfaResponse &&
        !options.skipAuthnRetry &&
        isAdminActionRequiresMfaError(err)
      ) {
        mfaResponse = await api.getAdminActionMfaResponse();
        const response = await api.fetch(url, customOptions, mfaResponse);
        return await api.getJsonFromFetchResponse(response, options);
      } else {
        throw err;
      }
    }
  },

  async getJsonFromFetchResponse(response: Response, options: RequestOptions) {
    const { allowRoleNotFound = false } = options;
    let json;
    try {
      json = await response.json();
    } catch (err) {
      // error reading JSON
      const message = response.ok
        ? err.message
        : `${response.status} - ${response.url}`;
      throw new ApiError({ message, response, opts: { cause: err } });
    }

    if (response.ok) {
      return json;
    }

    /** This error can occur in the edge case where a role in the user's certificate was deleted during their session. */
    const isRoleNotFoundErr =
      !allowRoleNotFound && isRoleNotFoundError(parseError(json));
    if (isRoleNotFoundErr) {
      websession.logoutWithoutSlo({
        /* Don't remember location after login, since they may no longer have access to the page they were on. */
        rememberLocation: false,
        /* Show "access changed" notice on login page. */
        withAccessChangedMessage: true,
      });
      return;
    }

    throw new ApiError({
      message: parseError(json),
      response,
      proxyVersion: parseProxyVersion(json),
      messages: json.messages,
    });
  },

  async getAdminActionMfaResponse() {
    const challenge = await auth.getMfaChallenge({
      scope: MfaChallengeScope.ADMIN_ACTION,
    });

    if (!challenge) {
      throw new Error(
        'This is an admin-level API request and requires MFA verification. Please try again with a registered MFA device. If you do not have an MFA device registered, you can add one in the account settings page.'
      );
    }

    return auth.getMfaChallengeResponse(challenge);
  },

  /**
   * fetch calls upon native fetch with options and headers set
   * to default (or provided) values.
   *
   * @param customOptions is an optional RequestInit object.
   * It can be provided to either override some fields defined in
   * `defaultRequestOptions` or define new fields not in
   * `defaultRequestOptions`.
   *
   * customOptions gets shallowly merged with `defaultRequestOptions` where
   * inner objects do not get merged if overrided.
   *
   * Example with an example customOption:
   * {
   *  body: formData,
   *  method: 'POST',
   *  headers: {
   *    Accept: 'application/json',
   *  }
   * }
   *
   * 'headers' is a field also defined in `defaultRequestOptions`, because of
   * shallow merging, the customOption.headers will get completely overrided.
   * After merge:
   *
   * {
   *  body: formData,
   *  method: 'POST',
   *  headers: {
   *    Accept: 'application/json',
   *  },
   *  credentials: 'same-origin',
   *  mode: 'same-origin',
   *  cache: 'no-store'
   * }
   *
   * If customOptions field is not provided, only fields defined in
   * `defaultRequestOptions` will be used.
   *
   * @param mfaResponse if defined (eg: `fetchJsonWithMfaAuthnRetry`)
   * will add a custom MFA header field that will hold the mfaResponse.
   *
   * @returns the native fetch's response object.
   */
  async fetch(
    url: string,
    customOptions: RequestInit = {},
    mfaResponse?: MfaChallengeResponse
  ): Promise<Response> {
    url = window.location.origin + url;
    const options = {
      ...defaultRequestOptions,
      ...customOptions,
    };

    options.headers = {
      ...options.headers,
      ...getAuthHeaders(),
    };

    if (mfaResponse) {
      options.headers[MFA_HEADER] = JSON.stringify({
        ...mfaResponse,
        // TODO(Joerger): DELETE IN v19.0.0.
        // We include webauthnAssertionResponse for backwards compatibility.
        webauthnAssertionResponse: mfaResponse.webauthn_response,
      });
    }

    // native call
    return await fetch(url, options);
  },
};

export const defaultRequestOptions: RequestInit = {
  credentials: 'same-origin',
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json; charset=utf-8',
  },
  mode: 'same-origin',
  cache: 'no-store',
};

export function getAuthHeaders() {
  const accessToken = getAccessToken();
  const csrfToken = getXCSRFToken();
  return {
    'X-CSRF-Token': csrfToken,
    Authorization: `Bearer ${accessToken}`,
  };
}

export function getNoCacheHeaders() {
  return {
    'cache-control': 'max-age=0',
    expires: '0',
    pragma: 'no-cache',
  };
}

export const getXCSRFToken = () => {
  const metaTag = document.querySelector(
    '[name=grv_csrf_token]'
  ) as HTMLMetaElement;
  return metaTag ? metaTag.content : '';
};

export function getAccessToken() {
  return storageService.getBearerToken()?.accessToken;
}

export function getHostName() {
  return location.hostname + (location.port ? ':' + location.port : '');
}

export function isAdminActionRequiresMfaError(err: Error) {
  return err.message.includes(
    'admin-level API request requires MFA verification'
  );
}

/** isRoleNotFoundError returns true if the error message is due to a role not being found. */
export function isRoleNotFoundError(errMessage: string): boolean {
  // This error message format should be kept in sync with the NotFound error message returned in lib/services/local/access.GetRole
  return /role \S+ is not found/.test(errMessage);
}

export default api;
