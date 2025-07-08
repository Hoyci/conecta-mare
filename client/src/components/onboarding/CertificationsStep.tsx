import { MAX_CERTIFICATIONS, ProfessionalProfile } from "@/types/user";
import { useFieldArray, useFormContext } from "react-hook-form";
import { Label } from "../ui/label";
import { Button } from "../ui/button";
import { Plus } from "lucide-react";
import { Card } from "../ui/card";
import { Input } from "../ui/input";
import { cn } from "@/lib/utils";

const CertificationsStep = () => {
  const { register, formState: { errors } } = useFormContext<ProfessionalProfile>();

  const { fields, append, remove } = useFieldArray<ProfessionalProfile>({
    name: "certifications",
  });

  return (
    <div className="space-y-6">
      <div className="space-y-4">
        <div className="flex justify-between items-center">
          <Label className="text-base text-gray-800">Certificações</Label>
          <Button
            type="button"
            disabled={fields.length >= MAX_CERTIFICATIONS}
            onClick={() =>
              append({
                institution: "",
                courseName: "",
                startDate: null,
                endDate: null,
              })
            }
            variant="outline"
            size="sm"
          >
            <Plus className="w-4 h-4 mr-2" />
            Adicionar Certificação
          </Button>
        </div>
      </div>
      {fields.map((field, index) => (
        <Card
          key={field.id}
          className="border-2 border-conecta-blue/20 space-y-4 p-4"
        >
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor={`certifications.${index}.institution`}>
                Instituição *
              </Label>
              <Input
                id={`certifications.${index}.institution`}
                placeholder="Nome da instituição"
                {...register(`certifications.${index}.institution`)}
                className={cn(errors.certifications?.[index]?.institution && "border-red-500")}
              />
              {errors.certifications?.[index]?.institution && (
                <p className="text-red-500 text-xs mt-1">
                  {errors.certifications?.[index]?.institution?.message}
                </p>
              )}
            </div>
            <div className="space-y-2">
              <Label htmlFor={`certifications.${index}.courseName`}>
                Nome do Curso *
              </Label>
              <Input
                id={`certifications.${index}.courseName`}
                placeholder="Nome do curso/certificação"
                {...register(`certifications.${index}.courseName`)}
                className={cn(errors.certifications?.[index]?.courseName && "border-red-500")}
              />
              {errors.certifications?.[index]?.courseName && (
                <p className="text-red-500 text-xs mt-1">
                  {errors.certifications?.[index]?.courseName?.message}
                </p>
              )}
            </div>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor={`certifications.${index}.startDate`}>
                Data de Início *
              </Label>
              <Input
                id={`certifications.${index}.startDate`}
                type="date"
                {...register(`certifications.${index}.startDate`)}
                className={cn(errors.certifications?.[index]?.startDate && "border-red-500")}
              />
              {errors.certifications?.[index]?.startDate && (
                <p className="text-red-500 text-xs mt-1">
                  {errors.certifications?.[index]?.startDate?.message}
                </p>
              )}
            </div>
            <div className="space-y-2">
              <Label htmlFor={`certifications.${index}.endDate`}>
                Data de Fim (opcional)
              </Label>
              <Input
                id={`certifications.${index}.endDate`}
                type="date"
                {...register(`certifications.${index}.endDate`)}
                className={cn(errors.certifications?.[index]?.endDate && "border-red-500")}
              />
              {errors.certifications?.[index]?.endDate && (
                <p className="text-red-500 text-xs mt-1">
                  {errors.certifications?.[index]?.endDate?.message}
                </p>
              )}
            </div>
          </div>
          <div className="flex justify-end pt-2">
            <Button
              type="button"
              variant="ghost"
              className="text-red-500"
              onClick={() => remove(index)}
            >
              Remover
            </Button>
          </div>
        </Card>
      ))}
    </div>
  );
};

export default CertificationsStep;
