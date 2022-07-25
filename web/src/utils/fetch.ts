import { fetchUtils, HttpError } from "react-admin";
import { KEY_TOKEN, TOKEN_TYPE } from "../constants";

export const fetchJsonFromConfig = () => {
  switch (TOKEN_TYPE) {
    case "JWT":
      return fetchJsonWithJWTToken;
    case "AUTH":
      return fetchJsonWithAuthToken;
    default:
      return fetchJson;
  }
};

export const createHeadersFromOptions = (options: fetchUtils.Options) =>
  fetchUtils.createHeadersFromOptions(options);

export const fetchJson = (url: string, options: fetchUtils.Options = {}) => {
  const requestHeaders = createHeadersFromOptions(options);

  return fetch(url, { ...options, headers: requestHeaders })
    .then((response) =>
      response.text().then((text) => ({
        status: response.status,
        statusText: response.statusText,
        headers: response.headers,
        body: text,
      }))
    )
    .then(({ status, statusText, headers, body }) => {
      let json;
      try {
        json = JSON.parse(body);
      } catch (e) {
        // not json, no big deal
      }
      if (status < 200 || status >= 300) {
        const msg =
          json &&
          (json.error ? json.error.id || json.error.message : json.message);
        return Promise.reject(new HttpError(msg || statusText, status, json));
      }
      return Promise.resolve({ status, headers, body, json });
    });
};

export const createOptionsJWTToken = (): fetchUtils.Options => {
  const token = localStorage.getItem(KEY_TOKEN);
  return {
    user: {
      authenticated: true,
      token: "Bearer " + token,
    },
  };
};

export const fetchJsonWithJWTToken = (
  url: string,
  options: fetchUtils.Options = {}
) => {
  return fetchJson(url, { ...options, ...createOptionsJWTToken() });
};

export const createOptionsAuthToken = (): fetchUtils.Options => {
  const token = localStorage.getItem(KEY_TOKEN);
  return {
    user: {
      authenticated: true,
      token: "Token " + token,
    },
  };
};

export const fetchJsonWithAuthToken = (
  url: string,
  options: fetchUtils.Options = {}
) => {
  return fetchJson(url, { ...options, ...createOptionsAuthToken() });
};
