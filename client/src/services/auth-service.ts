import { env } from "@/config/env";
import { RenewAccessTokenResponse, Session } from "@/types/auth";
import { authFetch } from "@/lib/auth-fetch";
import { SignUpValues, User } from "@/types/user";
import { toSnakeCase } from "@/lib/utils";

export const loginUser = async (email: string, password: string) => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/users/login`;

  const res = await fetch(apiUrl, {
    method: "POST",
    body: JSON.stringify({ email, password }),
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => null);
    throw new Error(errorData);
  }

  const session = await res.json();
  return session as Session;
};

export const signUpUser = async (payload: SignUpValues) => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/users/register`;

  const res = await fetch(apiUrl, {
    method: "POST",
    body: JSON.stringify(toSnakeCase(payload)),
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => null);
    throw new Error(errorData);
  }

  const user = await res.json();
  return user as User;
};

export const logoutUser = async () => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/users/logout`;

  const res = await authFetch(apiUrl, {
    method: "PATCH",
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => null);
    throw new Error(errorData);
  }

  return res;
};

export const refreshAccessToken = async () => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/auth/refresh`;

  const res = await fetch(apiUrl, {
    method: "POST",
    credentials: "include",
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => null);
    throw new Error(errorData);
  }

  return (await res.json()) as RenewAccessTokenResponse;
};
