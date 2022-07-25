import { Identity } from "@app/types";
import { fetchJson } from "utils/fetch";
import { API_ENDPOINT, KEY_IDENTITY } from "../../constants";

export const updateIdentityStore: (identity: Identity) => Promise<void> = (
  identity
) => {
  localStorage.setItem(KEY_IDENTITY, JSON.stringify(identity));

  return Promise.resolve();
};

export const updateIdentity: (identity: Identity) => Promise<Identity> = async (
  identity
) => {
  await fetchJson(`${API_ENDPOINT}/account/profile`, {
    method: "POST",
    body: JSON.stringify(identity),
  });

  return Promise.resolve(identity);
};
