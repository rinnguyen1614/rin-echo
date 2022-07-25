import { API_ENDPOINT, KEY_PERMISSIONS } from "../../constants";
import { Permission } from "../../types/models/Permission";
import { fetchJsonFromConfig } from "../../utils/fetch";

export const getPermissionsStore: () => Promise<Permission[]> = () => {
  try {
    const permissions = JSON.parse(
      localStorage.getItem(KEY_PERMISSIONS) ?? ""
    ) as Permission[];
    return Promise.resolve(permissions);
  } catch (err) {
    return Promise.reject(err);
  }
};

export const getPermissions: () => Promise<any> = async () => {
  const response = await fetchJsonFromConfig()(
    `${API_ENDPOINT}/account/permissions`,
    {
      method: "GET",
    }
  );
  return Promise.resolve(response.json.data);
};
