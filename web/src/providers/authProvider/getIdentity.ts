import { Identity } from "@app/types";
import { API_ENDPOINT, FILE_ENDPOINT, KEY_IDENTITY } from "../../constants";
import { fetchJsonFromConfig } from "../../utils/fetch";

export const getIdentityStore: () => Promise<Identity> = () => {
  try {
    const identity = JSON.parse(
      localStorage.getItem(KEY_IDENTITY) ?? ""
    ) as Identity;
    return Promise.resolve(identity);
  } catch (err) {
    return Promise.reject(err);
  }
};

export const getIdentity: () => Promise<Identity> = async () => {
  const response = await fetchJsonFromConfig()(
    `${API_ENDPOINT}/account/profile`,
    {
      method: "GET",
    }
  );

  if (response.json.data?.avatar_path) {
    response.json.data.avatar_path =
      FILE_ENDPOINT + response.json.data.avatar_path;
  }
  return Promise.resolve(response.json.data);
};
