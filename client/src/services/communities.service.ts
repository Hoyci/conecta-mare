import { env } from "@/config/env";
import { GetCommunitiesResponse } from "@/types/communities"

export const getCommunities = async () => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/communities`;

  const res = await fetch(apiUrl, {
    method: "GET",
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => null);
    throw new Error(errorData);
  }

  const communties = await res.json();
  return communties as GetCommunitiesResponse;
};

