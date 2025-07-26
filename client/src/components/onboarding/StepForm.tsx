import { FormProvider } from "react-hook-form";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export const StepForm = ({ step, methods, onSubmit, children }: any) => (
  <Card className="border-none shadow-lg">
    <CardHeader className="text-center pb-6">
      <CardTitle className="text-2xl text-gray-800">{step.title}</CardTitle>
      <p className="text-gray-600">{step.subtitle}</p>
    </CardHeader>

    <CardContent className="px-8 pb-8">
      <FormProvider {...methods}>
        <form onSubmit={methods.handleSubmit(onSubmit)}>{children}</form>
      </FormProvider>
    </CardContent>
  </Card>
);
