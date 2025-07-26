import { CheckCircle } from "lucide-react";

interface StepIndicatorProps {
  currentStep: number;
  steps: {
    indicatorTitle: string;
    indicatorSubtitle: string;
  }[];
}

export const StepIndicator = ({ currentStep, steps }: StepIndicatorProps) => (
  <div className="flex items-center justify-center gap-8 mb-8">
    {steps.map((step, index) => {
      const stepNumber = index + 1;
      const isCompleted = currentStep > stepNumber;
      const isActive = currentStep >= stepNumber;

      return (
        <div key={step.indicatorTitle} className="flex items-center gap-3">
          <div
            className={`w-10 h-10 rounded-full flex items-center justify-center text-sm font-medium ${
              isActive
                ? "bg-conecta-blue text-white shadow-lg"
                : "bg-gray-200 text-gray-600"
            }`}
          >
            {isCompleted ? <CheckCircle className="w-5 h-5" /> : stepNumber}
          </div>
          <div>
            <p className="font-medium text-gray-800">{step.indicatorTitle}</p>
            <p className="text-sm text-gray-500">{step.indicatorSubtitle}</p>
          </div>
          {stepNumber !== steps.length && (
            <div className="flex-1 h-px bg-gray-300 max-w-20" />
          )}
        </div>
      );
    })}
  </div>
);
