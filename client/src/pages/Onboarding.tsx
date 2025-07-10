import React from "react";
import { useForm, FormProvider, FieldErrors } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ArrowLeft, CheckCircle, Loader2 } from "lucide-react";
import { Link } from "react-router-dom";

import { OnboardingRequestSchema, OnboardingRequestValues } from "@/types/user";

import { UserDataStep } from "@/components/onboarding/UserDataStep";
import { ProjectStep } from "@/components/onboarding/ProjectStep";
import ServicesStep from "@/components/onboarding/ServicesStep";
import CertificationsStep from "@/components/onboarding/CertificationsStep";
import { useToast } from "@/hooks/use-toast";
import { submitOnboardingProfile } from "@/services/user-service";
import { useMutation } from "@tanstack/react-query";
import { watch } from "node:fs";

const steps = [
  {
    title: "Configure seu Perfil",
    subtitle: "Adicione suas informações pessoais e profissionais",
    indicatorTitle: "Dados Pessoais",
    indicatorSubtitle: "Seu perfil profissional",
    component: <UserDataStep />,
  },
  {
    title: "Adicione suas certificações",
    subtitle: "Descreva suas certificações para transmitir confiança",
    indicatorTitle: "Certificações",
    indicatorSubtitle: "Suas certificações",
    component: <CertificationsStep />,
  },
  {
    title: "Mostre seus Projetos",
    subtitle: "O que você já fez de melhor!",
    indicatorTitle: "Projetos",
    indicatorSubtitle: "O que você já fez de melhor!",
    component: <ProjectStep />,
  },
  {
    title: "Descreva seus Serviços",
    subtitle: "Conquiste clientes com suas habilidades",
    indicatorTitle: "Serviços",
    indicatorSubtitle: "Mostre o que você faz de melhor",
    component: <ServicesStep />,
  },
];

const fieldToStepMap: Record<string, number> = {
  profileImage: 1,
  jobDescription: 1,
  subcategoryID: 1,
  phone: 1,
  socialLinks: 1,
  certifications: 2,
  projects: 3,
  services: 4,
  hasOwnLocation: 4,
  location: 4,
};

const fieldOrder = [
  "profileImage",
  "jobDescription",
  "subcategoryID",
  "phone",
  "socialLinks",
  "certifications",
  "projects",
  "services",
  "hasOwnLocation",
  "location",
];

const Onboarding = () => {
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
        neighborhood: "",
      },
    },
  });

  const { mutate, isPending } = useMutation({
    mutationFn: submitOnboardingProfile,
    onSuccess: () => {
      toast({
        title: "Perfil configurado com sucesso! ✅",
        description: "Você será redirecionado para o seu dashboard.",
      });
      navigate("/dashboard");
    },
    onError: (error) => {
      toast({
        variant: "destructive",
        title: "Ocorreu um erro",
        description: error.message,
      });
    },
  });

  const watchedFields = methods.watch();

  const isStepValid = (step: number) => {
    switch (step) {
      case 1: {
        const { profileImage, jobDescription, subcategoryID, phone } =
          watchedFields;
        const isPhoneValid = phone && phone.replace(/\D/g, "").length >= 10;

        return profileImage && jobDescription && subcategoryID && isPhoneValid;
      }
      case 2: {
        const { certifications } = watchedFields;
        if (!certifications || certifications.length === 0) {
          return true;
        }

        return certifications.every(
          (cert) =>
            cert.institution.length > 0 &&
            cert.courseName.length > 0 &&
            cert.startDate,
        );
      }
      case 3: {
        const { projects } = watchedFields;
        if (!projects || projects.length === 0) {
          return true;
        }

        return projects.every(
          (project) =>
            project.name.length > 0 &&
            project.description.length > 0 &&
            project.images &&
            project.images.length >= 1,
        );
      }
      case 4: {
        const { services, hasOwnLocation, location } = watchedFields;

        if (!services || services.length === 0) {
          return false;
        }

        const firstService = services[0];
        const isServiceInfoFilled =
          firstService.name &&
          firstService.description &&
          firstService.price > 0;

        if (!isServiceInfoFilled) return false;

        if (hasOwnLocation) {
          const isLocationFilled =
            location?.street && location?.number && location?.neighborhood;

          const isOwnLocationPriceFilled =
            firstService.ownLocationPrice && firstService.ownLocationPrice > 0;

          return isLocationFilled && isOwnLocationPriceFilled;
        }

        return true;
      }
      default:
        return true;
    }
  };

  const onSubmit = (data: OnboardingRequestValues) => {
    console.log("Dados validados prontos para enviar:", data);
    mutate(data);
  };

  const onInvalid = (errors: FieldErrors<OnboardingRequestValues>) => {
    for (const field of fieldOrder) {
      if (errors[field as keyof OnboardingRequestValues]) {
        const errorStep = fieldToStepMap[field];
        if (errorStep) {
          setCurrentStep(errorStep);
          break;
        }
      }
    }
  };

  const handleNext = () => {
    if (isStepValid(currentStep) && currentStep < steps.length) {
      setCurrentStep((prev) => prev + 1);
    }
  };

  const handleBack = () => {
    if (currentStep > 1) {
      setCurrentStep((prev) => prev - 1);
    }
  };

  const step = steps[currentStep - 1];
  const totalSteps = steps.length;
  const isLastStep = currentStep === totalSteps;
  const isCurrentStepValid = isStepValid(currentStep);

  return (
    <div className="min-h-screen bg-gradient-to-br from-conecta-blue/5 to-conecta-green/5">
      <div className="bg-white shadow-sm border-b">
        <div className="max-w-4xl mx-auto px-4 py-4">
          <div className="flex items-center justify-between">
            <Link to="/" className="flex-shrink-0 flex items-center">
              <span className="text-conecta-blue text-xl font-bold">
                Conecta<span className="text-conecta-green">Maré</span>
              </span>
            </Link>
            <Button
              variant="ghost"
              onClick={() => navigate("/")}
              className="text-gray-500 hover:text-gray-700"
            >
              <ArrowLeft className="w-4 h-4 mr-2" />
              Voltar
            </Button>
          </div>
        </div>
      </div>

      <div className="max-w-6xl mx-auto px-4 py-8">
        <div className="text-center mb-8">
          <h1 className="text-3xl font-bold text-gray-800 mb-3">
            Bem-vindo ao ConectaMaré! 🎉
          </h1>
          <p className="text-lg text-gray-600 max-w-2xl mx-auto">
            Vamos configurar seu perfil profissional em alguns passos simples.
            Isso ajudará você a se conectar com clientes, mostrar seus serviços
            e projetos realizados.
          </p>
        </div>

        <StepIndicator currentStep={currentStep} steps={steps} />

        <Card className="border-none shadow-lg">
          <CardHeader className="text-center pb-6">
            <CardTitle className="text-2xl text-gray-800">
              {step.title}
            </CardTitle>
            <p className="text-gray-600">{step.subtitle}</p>
          </CardHeader>

          <CardContent className="px-8 pb-8">
            <FormProvider {...methods}>
              <form onSubmit={methods.handleSubmit(onSubmit)}>
                {step.component}
              </form>
            </FormProvider>
          </CardContent>
        </Card>

        <div className="flex justify-between mt-8">
          <Button
            variant="outline"
            onClick={handleBack}
            disabled={currentStep === 1}
            className="px-8 py-2 border-gray-300"
          >
            <ArrowLeft className="w-4 h-4 mr-2" />
            Voltar
          </Button>

          <div className="flex gap-3">
            {!isLastStep ? (
              <Button
                type="button"
                onClick={handleNext}
                disabled={!isCurrentStepValid}
                className="bg-conecta-blue hover:bg-conecta-blue-dark px-8 py-2 shadow-md"
              >
                Próximo Passo
              </Button>
            ) : (
              <Button
                type="button"
                onClick={methods.handleSubmit(onSubmit, onInvalid)}
                disabled={!isCurrentStepValid || isPending}
                className="bg-conecta-green hover:bg-conecta-green-dark px-8 py-2 shadow-md w-52"
              >
                {isPending ? (
                  <>
                    <Loader2 className="w-4 h-4 mr-2 animate-spin" />
                    Finalizando...
                  </>
                ) : (
                  <>
                    <CheckCircle className="w-4 h-4 mr-2" />
                    Finalizar Configuração
                  </>
                )}
              </Button>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

interface StepIndicatorProps {
  currentStep: number;
  steps: {
    indicatorTitle: string;
    indicatorSubtitle: string;
  }[];
}

const StepIndicator = ({ currentStep, steps }: StepIndicatorProps) => {
  return (
    <div className="flex items-center justify-center gap-8 mb-8">
      {steps.map((step, index) => {
        const stepNumber = index + 1;
        const isCompleted = currentStep > stepNumber;
        const isActive = currentStep >= stepNumber;

        return (
          <React.Fragment key={step.indicatorTitle}>
            <div className="flex items-center gap-3">
              <div
                className={`w-10 h-10 rounded-full flex items-center justify-center text-sm font-medium transition-all ${
                  isActive
                    ? "bg-conecta-blue text-white shadow-lg"
                    : "bg-gray-200 text-gray-600"
                }`}
              >
                {isCompleted ? <CheckCircle className="w-5 h-5" /> : stepNumber}
              </div>
              <div className="text-left">
                <p className="font-medium text-gray-800">
                  {step.indicatorTitle}
                </p>
                <p className="text-sm text-gray-500">
                  {step.indicatorSubtitle}
                </p>
              </div>
            </div>

            {stepNumber !== steps.length && (
              <div className="flex-1 h-px bg-gray-300 max-w-20"></div>
            )}
          </React.Fragment>
        );
      })}
    </div>
  );
};

export default Onboarding;
