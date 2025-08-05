import { useOnboarding } from "@/hooks/useOnboarding";
import { OnboardingLayout } from "@/components/onboarding/OnboardingLayout";
import { StepForm } from "@/components/onboarding/StepForm";
import { StepIndicator } from "@/components/onboarding/StepIndicator";
import { Button } from "@/components/ui/button";
import { CheckCircle, Loader2, LoaderCircle, ArrowLeft } from "lucide-react";

import { UserDataStep } from "@/components/onboarding/UserDataStep";
import { ProjectStep } from "@/components/onboarding/ProjectStep";
import ServicesStep from "@/components/onboarding/ServicesStep";
import CertificationsStep from "@/components/onboarding/CertificationsStep";
import { OnboardingRequestValues } from "@/types/user";

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

const Onboarding = () => {
  const { currentStep, setCurrentStep, methods, userData, isPending, mutate } =
    useOnboarding();

  const step = steps[currentStep - 1];
  const isLastStep = currentStep === steps.length;

  if (!userData?.id) return <LoaderCircle />;

  const handleNext = () => {
    setCurrentStep((prev: number) => Math.min(prev + 1, steps.length));
  };

  const handleBack = () => {
    setCurrentStep((prev: number) => Math.max(prev - 1, 1));
  };

  const onSubmit = (data: OnboardingRequestValues) => {
    const cleanedPhone = data.phone.replace(/\D/g, '');

    const formattedData = {
      ...data,
      phone: `55${cleanedPhone}`,
    };

    mutate(formattedData);
  };

  return (
    <OnboardingLayout>
      <div className="text-center mb-8">
        <h1 className="text-3xl font-bold text-gray-800 mb-3">
          Bem-vindo ao ConectaMaré! 🎉
        </h1>
        <p className="text-lg text-gray-600 max-w-2xl mx-auto">
          Vamos configurar seu perfil profissional em alguns passos simples.
        </p>
      </div>

      <StepIndicator currentStep={currentStep} steps={steps} />

      <StepForm step={step} methods={methods} onSubmit={onSubmit}>
        {step.component}
      </StepForm>

      <div className="flex justify-between mt-8">
        <Button
          variant="outline"
          onClick={handleBack}
          disabled={currentStep === 1}
        >
          <ArrowLeft className="w-4 h-4 mr-2" />
          Voltar
        </Button>

        {!isLastStep ? (
          <Button onClick={handleNext}>Próximo Passo</Button>
        ) : (
          <Button onClick={methods.handleSubmit(onSubmit)} disabled={isPending}>
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
    </OnboardingLayout>
  );
};

export default Onboarding;
