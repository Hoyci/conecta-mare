import { env } from "@/config/env";
import { authFetch } from "@/lib/auth-fetch";
import { toCamelCase } from "@/lib/utils";
import { UserProfileViewsResponse } from "@/types/metrics";

export const getUserProfileViews = async () => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/metrics/user-profile-views`;

  const res = await authFetch(apiUrl, {
    method: "GET",
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => null);
    throw new Error(errorData);
  }

  return toCamelCase(await res.json()) as UserProfileViewsResponse;
};
