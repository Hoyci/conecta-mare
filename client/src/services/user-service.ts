import { env } from "@/config/env";
import { toCamelCase } from "@/lib/utils";
import {
  ProfessionalUserResponse,
  ProfessionalUsersResponse,
} from "@/types/user";

export const getProfessionals = async () => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/users/professionals`;

  const res = await fetch(apiUrl, {
    method: "GET",
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => null);
    throw new Error(errorData);
  }

  const rawData = await res.json();

  const camelizedData: ProfessionalUsersResponse = {
    professionals: rawData.professionals.map(toCamelCase),
  };

  return camelizedData;
};

export const getProfessionalByID = async (userID: string) => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/users/professionals/${userID}`;

  const res = await fetch(apiUrl, {
    method: "GET",
  });

  if (!res.ok) {
    const erroData = await res.json().catch((error: Error) => error);
    throw new Error(erroData);
  }

  const rawData = await res.json();

  const camelizedDada: ProfessionalUserResponse = {
    data: toCamelCase(rawData.data),
  };

  return camelizedDada;
};
