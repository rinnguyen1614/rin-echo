import { stringify } from "query-string";
import { DataProvider } from "ra-core";
import { API_ENDPOINT } from "../../constants";
import { fetchJson, fetchJsonFromConfig } from "../../utils/fetch";
import { getQueryParams } from "../../utils/query";

const RestProvider = (
  apiUrl: string,
  httpClient = fetchJson,
  countHeader: string = "Content-Range"
): DataProvider => ({
  getList: (resource, params) => {
    const query = getQueryParams(params);
    const url = `${apiUrl}/${resource}?${stringify(query)}`;
    const options =
      countHeader === "Content-Range"
        ? {
            // Chrome doesn't return `Content-Range` header if no `Range` is provided in the request.
            // headers: new Headers({
            //   Range: `${resource}=${rangeStart}-${rangeEnd}`,
            // }),
          }
        : {};

    return httpClient(url, options).then(({ headers, json }) => {
      return {
        data: json.data.records ?? [],
        total: json.data.total,
      };
    });
  },

  getOne: (resource, params) =>
    httpClient(`${apiUrl}/${resource}/${params.id}`).then(({ json }) => ({
      data: json.data,
    })),

  getMany: (resource, params) => {
    const query = {
      filter: JSON.stringify({ id: params.ids }),
    };
    const url = `${apiUrl}/${resource}?${stringify(query)}`;
    return httpClient(url).then(({ json }) => ({
      data: json.data.records ?? [],
    }));
  },

  getManyReference: (resource, params) => {
    const { page, perPage } = params.pagination;
    const { field, order } = params.sort;

    const rangeStart = (page - 1) * perPage;
    const rangeEnd = page * perPage - 1;

    const query = {
      sorts: JSON.stringify(field + ":" + order),
      range: JSON.stringify([(page - 1) * perPage, page * perPage - 1]),
      filters: JSON.stringify({
        ...params.filter,
        [params.target]: params.id,
      }),
    };
    const url = `${apiUrl}/${resource}?${stringify(query)}`;
    const options =
      countHeader === "Content-Range"
        ? {
            // Chrome doesn't return `Content-Range` header if no `Range` is provided in the request.
            headers: new Headers({
              Range: `${resource}=${rangeStart}-${rangeEnd}`,
            }),
          }
        : {};

    return httpClient(url, options).then(({ headers, json }) => {
      if (!headers.has(countHeader)) {
        throw new Error(
          `The ${countHeader} header is missing in the HTTP Response. The simple REST data provider expects responses for lists of resources to contain this header with the total number of results to build the pagination. If you are using CORS, did you declare ${countHeader} in the Access-Control-Expose-Headers header?`
        );
      }
      return {
        data: json,
        total:
          countHeader === "Content-Range"
            ? parseInt(headers.get("content-range")?.split("/").pop() ?? "", 10)
            : parseInt(headers.get(countHeader.toLowerCase()) ?? ""),
      };
    });
  },

  update: (resource, params) =>
    httpClient(`${apiUrl}/${resource}/${params.id}`, {
      method: "PUT",
      body: JSON.stringify(params.data),
    }).then(({ json }) => ({ data: json.data })),

  // simple-rest doesn't handle provide an updateMany route, so we fallback to calling update n times instead
  updateMany: (resource, params) =>
    Promise.all(
      params.ids.map((id) =>
        httpClient(`${apiUrl}/${resource}/${id}`, {
          method: "PUT",
          body: JSON.stringify(params.data),
        })
      )
    ).then((responses) => ({ data: responses.map(({ json }) => json.id) })),

  create: (resource, params) =>
    httpClient(`${apiUrl}/${resource}`, {
      method: "POST",
      body: JSON.stringify(params.data),
    }).then(({ json }) => ({
      data: { ...params.data, id: json.data.id },
    })),

  delete: (resource, params) =>
    httpClient(`${apiUrl}/${resource}/${params.id}`, {
      method: "DELETE",
      headers: new Headers({
        "Content-Type": "text/plain",
      }),
    }).then(({ json }) => ({ data: json })),

  // simple-rest doesn't handle filters on DELETE route, so we fallback to calling DELETE n times instead
  deleteMany: (resource, params) =>
    Promise.all(
      params.ids.map((id) =>
        httpClient(`${apiUrl}/${resource}/${id}`, {
          method: "DELETE",
          headers: new Headers({
            "Content-Type": "text/plain",
          }),
        })
      )
    ).then((responses) => ({
      data: responses.map(({ json }) => json.id),
    })),
});

const defaultProvider = RestProvider(API_ENDPOINT, fetchJsonFromConfig());
const dataProviders = {
  admin: defaultProvider,
};

const getProvider = (resource: string) => {
  const splited = resource.split("/");
  if (splited.length > 1 && dataProviders[splited[0]]) {
    return [dataProviders[splited[0]], resource.slice(splited[0].length + 1)];
  }

  return [defaultProvider, resource];
};

const delayedDataProvider = new Proxy(defaultProvider, {
  get: (target, name, self) =>
    name === "then" // as we await for the dataProvider, JS calls then on it. We must trap that call or else the dataProvider will be called with the then method
      ? self
      : (resource: string, params: any) => {
          var [provider, resourceSplited] = getProvider(resource);
          return new Promise((resolve) =>
            setTimeout(
              () => resolve(provider[name as string](resourceSplited, params)),
              500
            )
          );
        },
});

export default delayedDataProvider;
