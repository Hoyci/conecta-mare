import { env } from "@/config/env";
import { authFetch } from "@/lib/auth-fetch";
import { createOnboardingFormData } from "@/lib/formDataHelper";
import { toCamelCase } from "@/lib/utils";
import { GetUserResponse, OnboardingRequestValues, ProfessionalProfile } from "@/types/user";

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

  const camelizedData: { professionals: ProfessionalProfile[] } = {
    professionals: rawData.professionals.map(toCamelCase),
  };

  return camelizedData;
};

export const getProfessionalByID = async (userID: string) => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/users/professionals/${userID}`;

  const response = await fetch(apiUrl, {
    method: "GET",
  });

  if (!response.ok) {
    const errorData = await response
      .json()
      .catch(() => ({ message: "Erro desconhecido ao requisitar dados do profissional" }));
    throw new Error(errorData.message || "Falha ao requisitar dados do profissional.");
  }

  const rawData = await response.json();

  return toCamelCase(rawData.data)
};

export const submitOnboardingProfile = async (
  data: OnboardingRequestValues,
) => {
  const formData = createOnboardingFormData(data);

  const apiUrl = `${env.data.VITE_API_URL}/api/v1/onboarding`;

  const response = await authFetch(apiUrl, {
    method: "POST",
    body: formData,
  });

  if (!response.ok) {
    const errorData = await response
      .json()
      .catch(() => ({ message: "Erro desconhecido ao processar a resposta." }));
    throw new Error(errorData.message || "Falha ao atualizar o perfil.");
  }

  return response.json();
};

export const getSignedUser = async () => {
  const apiUrl = `${env.data.VITE_API_URL}/api/v1/users/ `

  const response = await authFetch(apiUrl, {
    method: "GET",
  })

  if (!response.ok) {
    const errorData = await response
      .json()
      .catch(() => ({ message: "Erro desconhecido ao processar a resposta." }));
    throw new Error(errorData.message || "Falha ao requisitar dados do perfil.");
  }

  const rawData = await response.json();

  const camelizedData: GetUserResponse = toCamelCase(rawData);

  return camelizedData;
}
