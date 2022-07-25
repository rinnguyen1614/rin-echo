import { GetListParams } from "react-admin";

export const getQueryParams = (params: GetListParams) => {
  let query: any = {};

  if (params.pagination) {
    const { page, perPage } = params.pagination;
    query["page_size"] = perPage;
    query["page"] = page;
  }

  if (params.filter) {
    query["filter"] = getQueryFilters(params.filter);
    query["select"] = getFields(params.filter);
  }

  if (params.sort) {
    const { field, order } = params.sort;
    query["sorts"] = field + ":" + order;
  }

  return query;
};

const getQueryFilters = (filter: any): string => {
  //clone and remove select key in filter
  filter = Object.assign({}, filter);
  delete filter["select"];

  // get query string from search field
  let qFilter: any = {};
  let hasQ = false;
  for (var k in filter) {
    if (k.startsWith("q=")) {
      var v = filter[k];
      if (v) {
        hasQ = true;
      }
      // remove key "q" from filter
      delete filter[k];
      k.substring(2, k.length)
        .split(",")
        // eslint-disable-next-line no-loop-func
        .forEach((f) => {
          qFilter[f] = v;
        });
      break;
    }
  }

  var queries = [];
  if (Object.keys(filter).length !== 0) {
    var otherQuery = getOperateFilter(filter, "and");
    queries.push(otherQuery);
  }

  if (hasQ) {
    var qQuery = getOperateFilter(qFilter, "or");
    queries.push(qQuery);
  }

  if (queries.length > 1) {
    return "(" + queries.join(") and (") + ")";
  }

  return queries.length ? queries[0] : "";
};

const getOperateFilter = (filter: any, joinOp: string): string => {
  let filters: string[] = [];
  for (var k in filter) {
    var v = filter[k];
    var ksplited = k.split(":");
    var field = ksplited[0];
    var operator = "=";

    if (ksplited.length > 1) {
      operator = ksplited[1];
    }

    if (typeof v === "string" || v instanceof String) {
      if (operator === "like") {
        v = '"%' + v + '%"';
      } else {
        v = '"' + v + '"';
      }
    }
    filters.push(field + " " + operator + " " + v);
  }

  return filters.join(" " + joinOp + " ");
};

const getFields = (filter: any): string => {
  return filter.select;
};
