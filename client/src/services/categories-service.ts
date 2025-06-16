import { env } from "@/config/env";
import {
  CategoryWithSubsResponse,
  CategoryWithUserCountResponse,
} from "@/types/categories";

export const getCategoriesWithUserCount = async () => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/categories`;

  const res = await fetch(apiUrl, {
    method: "GET",
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => null);
    throw new Error(errorData);
  }

  const user = await res.json();
  return user as CategoryWithUserCountResponse;
};

export const getCategoriesWithSubs = async () => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/categories?include=subcategories`;

  const res = await fetch(apiUrl, {
    method: "GET",
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => null);
    throw new Error(errorData);
  }

  const user = await res.json();
  return user as CategoryWithSubsResponse;
};
