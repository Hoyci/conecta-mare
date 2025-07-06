import { useFormContext, useFieldArray } from "react-hook-form";
import {
  MAX_PROJECT_IMAGES,
  ProfessionalProfile,
  ProjectImage,
} from "@/types/user";
import { Label } from "../ui/label";
import { Button } from "../ui/button";
import { Plus, Upload, X } from "lucide-react";

import { Card } from "../ui/card";
import { Input } from "../ui/input";
import { Textarea } from "../ui/textarea";
import { cn } from "@/lib/utils";
import { useCallback } from "react";

interface ProjectCardProps {
  index: number;
  onRemove: () => void;
}

interface ProjectImageUploadProps {
  index: number;
}

interface ProjectImageGalleryProps {
  images: ProjectImage[];
  onRemoveImage: (index: number) => void;
}

interface ProjectImageItemProps {
  image: ProjectImage;
  onRemove: () => void;
}

export const ProjectStep = () => {
  const { control } = useFormContext<ProfessionalProfile>();
  const {
    fields: projects,
    append,
    remove,
  } = useFieldArray({
    name: "projects",
    control,
  });

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <Label className="text-base text-gray-800">Projetos</Label>
        <Button
          type="button"
          disabled={projects.length === MAX_PROJECT_IMAGES}
          onClick={() => append({ name: "", description: "", images: [] })}
          variant="outline"
          size="sm"
        >
          <Plus className="w-4 h-4 mr-2" /> Adicionar Projeto
        </Button>
      </div>

      {projects.map((project, index) => (
        <ProjectCard
          key={project.id}
          index={index}
          onRemove={() => remove(index)}
        />
      ))}
    </div>
  );
};

export const ProjectCard = ({ index, onRemove }: ProjectCardProps) => {
  const { register, formState } = useFormContext<ProfessionalProfile>();
  const errors = formState.errors;

  const handleRemoveProject = useCallback(() => {
    onRemove();
  }, [onRemove]);

  return (
    <Card className="p-6 border border-conecta-blue/20 space-y-4">
      <div className="flex justify-between items-center">
        <h3 className="font-semibold text-lg">Projeto {index + 1}</h3>
        <Button
          type="button"
          onClick={handleRemoveProject}
          variant="ghost"
          className="text-red-500"
        >
          <X className="w-4 h-4 mr-1" /> Remover
        </Button>
      </div>

      <div className="space-y-2">
        <Label htmlFor={`projects.${index}.name`}>Nome do Projeto *</Label>
        <Input
          id={`projects.${index}.name`}
          placeholder="Ex: Design de Logo Profissional"
          {...register(`projects.${index}.name`)}
          className={errors.projects?.[index]?.name && "border-red-500"}
        />
        {errors.projects?.[index]?.name && (
          <p className="text-red-500 text-xs mt-1">
            {errors.projects[index]?.name?.message}
          </p>
        )}
      </div>

      <div className="space-y-2">
        <Label htmlFor={`projects.${index}.description`}>Descrição *</Label>
        <Textarea
          id={`projects.${index}.description`}
          placeholder="Descreva detalhadamente..."
          {...register(`projects.${index}.description`)}
          className={cn(
            "min-h-[120px]",
            errors.projects?.[index]?.description && "border-red-500",
          )}
        />
        {errors.projects?.[index]?.description && (
          <p className="text-red-500 text-xs mt-1">
            {errors.projects[index]?.description?.message}
          </p>
        )}
      </div>

      <ProjectImageUpload index={index} />
    </Card>
  );
};

export const ProjectImageUpload = ({ index }: ProjectImageUploadProps) => {
  const { setValue, watch } = useFormContext();
  const images: ProjectImage[] = watch(`projects.${index}.images`) || [];

  const handleImageUpload = useCallback(
    (event: React.ChangeEvent<HTMLInputElement>) => {
      const files = Array.from(event.target.files || []);
      const availableSlots = MAX_PROJECT_IMAGES - images.length;

      if (availableSlots <= 0) return;
      const filesToAdd = files.slice(0, availableSlots);

      const newImages = filesToAdd.map((file) => ({
        file,
        url: URL.createObjectURL(file),
      }));

      setValue(`projects.${index}.images`, [...images, ...newImages], {
        shouldValidate: true,
      });
    },
    [images, index, setValue],
  );

  const handleRemoveImage = useCallback(
    (imgIndex: number) => {
      const imageToRemove = images[imgIndex];
      if (imageToRemove.file && imageToRemove.url) {
        URL.revokeObjectURL(imageToRemove.url);
      }

      const updated = images.filter((_, i) => i !== imgIndex);
      setValue(`projects.${index}.images`, updated, {
        shouldValidate: true,
      });
    },
    [images, index, setValue],
  );

  return (
    <div className="space-y-4">
      <Label>Fotos do Projeto</Label>
      <label htmlFor={`projects.${index}.images`}>
        <div className="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center cursor-pointer hover:border-conecta-blue transition-colors">
          <Upload className="w-12 h-12 mx-auto text-gray-400 mb-4" />
          <p className="text-gray-600 mb-2">
            Clique para carregar fotos do projeto
          </p>
          <p className="text-sm text-gray-500">
            PNG, JPG até 5MB cada (máximo 5 fotos)
          </p>
        </div>
        <input
          id={`projects.${index}.images`}
          type="file"
          accept="image/*"
          multiple
          className="hidden"
          onChange={handleImageUpload}
          disabled={images.length >= MAX_PROJECT_IMAGES}
        />
      </label>

      <ProjectImageGallery images={images} onRemoveImage={handleRemoveImage} />
    </div>
  );
};

export const ProjectImageGallery = ({
  images,
  onRemoveImage,
}: ProjectImageGalleryProps) => {
  if (images.length === 0) return null;

  return (
    <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
      {images.map((image, imgIndex) => (
        <ProjectImageItem
          key={imgIndex}
          image={image}
          onRemove={() => onRemoveImage(imgIndex)}
        />
      ))}
    </div>
  );
};

export const ProjectImageItem = ({
  image,
  onRemove,
}: ProjectImageItemProps) => (
  <div className="relative group">
    <div className="aspect-square rounded-lg overflow-hidden bg-gray-100">
      <img
        src={image.url || (image.file ? URL.createObjectURL(image.file) : "")}
        alt="Imagem do projeto"
        className="w-full h-full object-cover"
      />
    </div>
    <button
      type="button"
      onClick={onRemove}
      className="absolute -top-2 -right-2 bg-red-500 text-white rounded-full p-1 opacity-0 group-hover:opacity-100 transition-opacity"
    >
      <X className="w-4 h-4" />
    </button>
  </div>
);
