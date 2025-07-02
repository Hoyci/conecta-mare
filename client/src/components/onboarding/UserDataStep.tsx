import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { Upload, User } from "lucide-react";
import { MAX_JOB_DESCRIPTION_CHARS, ProfessionalProfile } from "@/types/user";
import { useFormContext } from "react-hook-form";
import { cn } from "@/lib/utils";
import { useEffect, useState } from "react";
import { MaskedInput } from "../ui/masked-input";
import { useQuery } from "@tanstack/react-query";
import { getCategoriesWithSubs } from "@/services/categories-service";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { CategoryWithSubsResponse } from "@/types/categories";

export const UserDataStep = () => {
  const {
    register,
    setValue,
    formState: { errors },
    watch,
  } = useFormContext<ProfessionalProfile>();

  const { data, isLoading } = useQuery<CategoryWithSubsResponse>({
    queryKey: ["categoriesWithSubs"],
    queryFn: getCategoriesWithSubs,
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

  const [jobDescription, phone, instagram, subcategoryID] = watch([
    "jobDescription",
    "phone",
    "socialLinks.instagram",
    "subcategoryID",
  ]);

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
                    <User className="w-8 h-8" />
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
              <h3 className="font-semibold text-lg mb-2 text-gray-800">
                Foto de Perfil
              </h3>
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
          maxLength={MAX_JOB_DESCRIPTION_CHARS}
          placeholder="Padeiro da padaria mais famosa da maré..."
          {...register("jobDescription")}
          className={cn(
            "min-h-[100px] resize-none",
            errors.jobDescription && "border-red-500",
          )}
        />
        <div className="text-xs text-gray-500 text-right">
          {jobDescription.length}/{MAX_JOB_DESCRIPTION_CHARS} caracteres
        </div>
        {errors.jobDescription && (
          <p className="text-red-500 text-xs mt-1">
            {errors.jobDescription.message}
          </p>
        )}
      </div>
      <div className="space-y-2">
        <Label htmlFor="subcategory">Categoria *</Label>

        <Select
          value={subcategoryID || ""}
          onValueChange={(value) => setValue("subcategoryID", value)}
        >
          <SelectTrigger
            id="subcategory"
            aria-invalid={!!errors.subcategoryID}
            disabled={isLoading}
          >
            <SelectValue
              placeholder={
                isLoading ? "Carregando..." : "Selecione uma categoria"
              }
            />
          </SelectTrigger>

          <SelectContent>
            {isLoading ? (
              <SelectItem value="loading" disabled>
                Carregando...
              </SelectItem>
            ) : (
              data?.categories.map((category) => (
                <SelectGroup key={category.id}>
                  <SelectLabel>{`${category.icon} ${category.name}`}</SelectLabel>
                  {category.subcategories.map((subcat) => (
                    <SelectItem key={subcat.id} value={subcat.id}>
                      {subcat.name}
                    </SelectItem>
                  ))}
                </SelectGroup>
              ))
            )}
          </SelectContent>
        </Select>

        {errors.subcategoryID && (
          <p className="text-red-500 text-xs mt-1">
            {errors.subcategoryID.message}
          </p>
        )}
      </div>
      <div className="space-y-2">
        <Label htmlFor="phone">Telefone de contato *</Label>
        <MaskedInput
          id="phone"
          mask="(99) 99999-9999"
          value={phone}
          onChange={(e) => setValue("phone", e.target.value)}
          onBlur={(e) => setValue("phone", e.target.value)}
          placeholder="(21) 98765-4321"
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
          value={instagram}
          onChange={(e) => setValue("socialLinks.instagram", e.target.value)}
          onFocus={(e) => {
            if (e.target.value === "") {
              setValue("socialLinks.instagram", "@");
            }
          }}
          onBlur={(e) => {
            if (e.target.value === "@") {
              setValue("socialLinks.instagram", "");
            }
          }}
          placeholder="@seuinstagram"
          className={errors.socialLinks?.instagram && "border-red-500"}
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
    </div>
  );
};
