import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { useToast } from "@/hooks/use-toast";
import { useState } from "react";
import {
  getSignedUser,
  submitOnboardingProfile,
} from "@/services/user-service";
import { OnboardingRequestSchema, OnboardingRequestValues } from "@/types/user";
import { getAnalytics } from "@/lib/analytics";

export const useOnboarding = () => {
  const navigate = useNavigate();
  const { toast } = useToast();
  const [currentStep, setCurrentStep] = useState(1);

  const methods = useForm<OnboardingRequestValues>({
    resolver: zodResolver(OnboardingRequestSchema),
    mode: "all",
    defaultValues: {
      jobDescription: "",
      profileImage: undefined,
      phone: "",
      socialLinks: {
        instagram: "",
        linkedin: "",
      },
      certifications: [],
      projects: [],
      services: [],
      hasOwnLocation: false,
      location: {
        street: "",
        number: "",
        complement: "",
        communityId: "",
      },
    },
  });

  const { data: userData = {} } = useQuery({
    queryKey: ["userData"],
    queryFn: getSignedUser,
  });

  const { mutate, isPending } = useMutation({
    mutationFn: submitOnboardingProfile,
    onSuccess: (_, { location, subcategoryID }) => {
      const analytics = getAnalytics();
      toast({
        title: "Perfil configurado com sucesso! ✅",
        description: "Você será redirecionado para o seu dashboard.",
      });
      navigate("/dashboard");
      analytics.track("onboarding_completed", {
        community_id: location.communityId,
        subcategory_id: subcategoryID,
      });
    },
    onError: (error) => {
      toast({
        variant: "destructive",
        title: "Ocorreu um erro",
        description: error.message,
      });
    },
  });

  return {
    userData,
    methods,
    currentStep,
    setCurrentStep,
    isPending,
    mutate,
  };
};
