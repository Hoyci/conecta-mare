import { env } from "@/config/env";
import { toCamelCase } from "@/lib/utils";
import { ProfessionalUsersResponse } from "@/types/user";

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
