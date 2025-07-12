import { refreshAccessToken } from "@/services/auth-service";
import { getCurrentAuthStore } from "@/store/auth-store-utils";

let refreshingPromise: Promise<string> | null = null;

export const authFetch = async (input: RequestInfo, init: RequestInit = {}) => {
  const store = getCurrentAuthStore();
  let accessToken = store.accessToken;

  const doFetch = async (token: string) => {
    const headers = new Headers(init.headers);
    headers.set("Authorization", `Bearer ${token}`);
    if (
      init.body &&
      !headers.has("Content-Type") &&
      !(init.body instanceof FormData)
    ) {
      headers.set("Content-Type", "application/json");
    }

    return fetch(input, {
      ...init,
      headers,
    });
  };

  let response = await doFetch(accessToken);

  if (response.status === 401) {
    if (!refreshingPromise) {
      refreshingPromise = (async () => {
        const data = await refreshAccessToken();
        store.updateAccessToken(data.access_token);
        return data.access_token;
      })();
      refreshingPromise.finally(() => {
        refreshingPromise = null;
      });
    }
    accessToken = await refreshingPromise;
    response = await doFetch(accessToken);
  }

  return response;
};
