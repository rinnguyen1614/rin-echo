import { API_ENDPOINT } from "../../constants";
import { fetchJsonFromConfig } from "../../utils/fetch";

export const changePassword: (params: any) => Promise<any> = async ({
  current_password,
  new_password,
}) => {
  await fetchJsonFromConfig()(`${API_ENDPOINT}/account/password`, {
    method: "PUT",
    body: JSON.stringify({
      current_password,
      new_password,
    }),
  });

  return Promise.resolve();
};
