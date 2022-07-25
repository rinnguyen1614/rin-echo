import { API_ENDPOINT, KEY_MENUS } from "../../constants";
import { Menu } from "../../types/models/Menu";
import { fetchJsonFromConfig } from "../../utils/fetch";

export const getMenusStore: () => Promise<Menu[]> = () => {
  try {
    const menus = JSON.parse(localStorage.getItem(KEY_MENUS) ?? "") as Menu[];
    return Promise.resolve(menus);
  } catch (err) {
    return Promise.reject(err);
  }
};

export const getMenus: () => Promise<any> = async () => {
  const response = await fetchJsonFromConfig()(
    `${API_ENDPOINT}/account/menus`,
    {
      method: "GET",
    }
  );
  return Promise.resolve(response.json.data);
};
