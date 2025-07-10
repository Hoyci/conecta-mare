import { OnboardingRequestValues } from "@/types/user";
import { toSnakeCase } from "./utils";

export const createOnboardingFormData = (
  data: OnboardingRequestValues,
): FormData => {
  const formData = new FormData();

  const payload = JSON.parse(JSON.stringify(toSnakeCase(data)));

  if (payload.profile_image) {
    delete payload.profile_image;
  }

  if (payload.projects && payload.projects.length > 0) {
    payload.projects.forEach((project: any) => {
      if (project.images) {
        delete project.images;
      }
    });
  }

  if (payload.services && payload.services.length > 0) {
    payload.services.forEach((service: any) => {
      if (service.images) {
        delete service.images;
      }
    });
  }

  if ("has_own_location" in payload) {
    delete payload.has_own_location;
  }

  formData.append("body", JSON.stringify(payload));

  if (data.profileImage && data.profileImage.length > 0) {
    formData.append("profile_image", data.profileImage[0]);
  }

  if (data.projects && data.projects.length > 0) {
    data.projects.forEach((project, projectIndex) => {
      if (project.images && project.images.length > 0) {
        project.images.forEach((image) => {
          if (image.file instanceof File) {
            formData.append(`projects[${projectIndex}].images`, image.file);
          }
        });
      }
    });
  }

  if (data.services && data.services.length > 0) {
    data.services.forEach((service, serviceIndex) => {
      if (service.images && service.images.length > 0) {
        service.images.forEach((image) => {
          if (image.file instanceof File) {
            formData.append(`services[${serviceIndex}].images`, image.file);
          }
        });
      }
    });
  }

  console.log("FormData a ser enviado:", Array.from(formData.entries()));

  return formData;
};
