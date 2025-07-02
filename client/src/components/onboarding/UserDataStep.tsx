import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { Plus, Upload } from "lucide-react";
import { MAX_CERTIFICATIONS, ProfessionalProfile } from "@/types/user";
import { useFieldArray, useFormContext } from "react-hook-form";
import { cn } from "@/lib/utils";
import { useEffect, useState } from "react";

export const UserDataStep = () => {
  const {
    register,
    formState: { errors },
    watch,
  } = useFormContext<ProfessionalProfile>();

  const { fields, append, remove } = useFieldArray<ProfessionalProfile>({
    name: "certifications",
  });

  const profileImage = watch("profileImage") as FileList | null;
  const [imageUrl, setImageUrl] = useState<string | null>(null);

  useEffect(() => {
    if (profileImage && profileImage.length > 0) {
      const file = profileImage[0];
      const url = URL.createObjectURL(file);
      setImageUrl(url);

      return () => URL.revokeObjectURL(url);
    }
    setImageUrl(null);
  }, [profileImage]);

  console.log(fields);

  return (
    <div className="space-y-6">
      <Card>
        <CardContent className="p-6">
          <div className="flex items-center gap-6">
            <div className="flex flex-col items-center gap-3">
              <Avatar className="w-24 h-24">
                {imageUrl ? (
                  <AvatarImage src={imageUrl} alt="Profile" />
                ) : (
                  <AvatarFallback className="bg-conecta-gray text-gray-600 text-2xl">
                    <Upload className="w-8 h-8" />
                  </AvatarFallback>
                )}
              </Avatar>
              <label htmlFor="profile-image">
                <Button variant="outline" className="cursor-pointer" asChild>
                  <span>
                    <Upload className="w-4 h-4 mr-2" />
                    Carregar Foto
                  </span>
                </Button>
                <input
                  id="profile-image"
                  type="file"
                  accept="image/*"
                  className="hidden"
                  {...register("profileImage")}
                />
              </label>
            </div>
            <div className="flex-1">
              <h3 className="font-semibold text-lg mb-2">Foto de Perfil</h3>
              <p className="text-gray-600 text-sm">
                Adicione uma foto profissional que represente bem você e seu
                trabalho. Formatos aceitos: JPG, PNG (máx. 5MB)
              </p>
            </div>
          </div>
        </CardContent>
      </Card>
      <div className="space-y-2">
        <Label htmlFor="description">Descrição sobre você *</Label>
        <Textarea
          id="description"
          placeholder="Padeiro da padaria mais famosa da maré..."
          {...register("jobDescription")}
          className={cn(
            "min-h-[100px] resize-none",
            errors.jobDescription && "border-red-500",
          )}
        />
        {errors.jobDescription && (
          <p className="text-red-500 text-xs mt-1">
            {errors.jobDescription.message}
          </p>
        )}
      </div>
      <div className="space-y-2">
        <Label htmlFor="phone">Telefone de contato *</Label>
        <Input
          id="phone"
          placeholder="(21) 99999-9999"
          {...register("phone")}
          className={errors.phone && "border-red-500"}
        />
        {errors.phone && (
          <p className="text-red-500 text-xs mt-1">{errors.phone.message}</p>
        )}
      </div>
      <div className="space-y-2">
        <Label htmlFor="socialMedia.instagram">Instagram</Label>
        <Input
          id="socialMedia.instagram"
          placeholder="@seuinstagram"
          {...register("socialLinks.instagram")}
        />
        {errors.socialLinks?.instagram && (
          <p className="text-red-500 text-xs mt-1">
            {errors.socialLinks.instagram.message}
          </p>
        )}
      </div>
      <div className="space-y-2">
        <Label htmlFor="socialMedia.linkedin">LinkedIn</Label>
        <Input
          id="socialMedia.linkedin"
          placeholder="https://www.linkedin.com/in/john-doe"
          {...register("socialLinks.linkedin")}
        />
        {errors.socialLinks?.linkedin && (
          <p className="text-red-500 text-xs mt-1">
            {errors.socialLinks.linkedin.message}
          </p>
        )}
      </div>
      <div className="space-y-4">
        <div className="flex justify-between items-center">
          <Label>Certificações</Label>
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
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor={`certifications.${index}.courseName`}>
                Nome do Curso *
              </Label>
              <Input
                id={`certifications.${index}.courseName`}
                placeholder="Nome do curso/certificação"
                {...register(`certifications.${index}.courseName`)}
              />
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
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor={`certifications.${index}.endDate`}>
                Data de Fim (opcional)
              </Label>
              <Input
                id={`certifications.${index}.endDate`}
                type="date"
                {...register(`certifications.${index}.endDate`)}
              />
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
