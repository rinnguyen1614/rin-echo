import { API_ENDPOINT, KEY_TOKEN } from "../../constants";
import { fetchJson } from "../../utils/fetch";

export const login: (params: any) => Promise<any> = async ({
  username,
  password,
}) => {
  const response = await fetchJson(`${API_ENDPOINT}/account/login`, {
    method: "POST",
    body: JSON.stringify({ username, password }),
  });

  var result = response.json.data;

  localStorage.setItem(KEY_TOKEN, result.access_token);

  return Promise.resolve();
};
