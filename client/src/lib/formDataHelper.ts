
import { OnboardingRequestValues } from "@/types/user";


export const createOnboardingFormData = (data: OnboardingRequestValues): FormData => {
  const formData = new FormData();

  const jsonData = JSON.parse(JSON.stringify(data));

  if (data.profileImage) {
    formData.append("profileImage", data.profileImage[0]);
  }
  delete jsonData.profileImage;

  data.services.forEach((service, index) => {
    if (service.images && service.images.length > 0) {
      formData.append(`service.${index}.image`, service.images[0].file);
    }
    if (jsonData.services[index]) {
      delete jsonData.services[index].images;
    }
  });

  data.projects?.forEach((project, projectIndex) => {
    if (project.images && project.images.length > 0) {
      project.images.forEach((image) => {
        formData.append(`project.${projectIndex}.images[]`, image.file);
      });
    }
    if (jsonData.projects[projectIndex]) {
      delete jsonData.projects[projectIndex].images;
    }
  });


  formData.append("body", JSON.stringify(jsonData));

  return formData;
};
