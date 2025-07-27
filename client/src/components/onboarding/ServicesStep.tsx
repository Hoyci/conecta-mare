import { useFormContext, useFieldArray, Controller } from "react-hook-form";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { Textarea } from "../ui/textarea";
import { Button } from "../ui/button";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { Plus, Upload, X } from "lucide-react";
import { Card } from "../ui/card";
import { useCallback } from "react";
import {
  MAX_SERVICE_IMAGES,
  MAX_SERVICES,
  ProfessionalProfile,
  ServiceImage,
} from "@/types/user";
import { getCommunities } from "@/services/communities.service";
import { useQuery } from "@tanstack/react-query";

interface ServiceImageUploadProps {
  index: number;
}

interface ServiceImageGalleryProps {
  images: ServiceImage[];
  onRemoveImage: (index: number) => void;
}

interface ServiceImageItemProps {
  image: ServiceImage;
  onRemove: () => void;
}

const ServicesStep = () => {
  const { register, watch, control } = useFormContext<ProfessionalProfile>();
  const hasOwnLocation = watch("hasOwnLocation");

  const {
    data: { communities } = {},
    isLoading: isLoadingCommunities,
    isError,
  } = useQuery({
    queryKey: ["communities"],
    queryFn: getCommunities,
  });

  const {
    fields: services,
    append,
    remove,
  } = useFieldArray({
    name: "services",
    control,
  });

  return (
    <div className="space-y-6">
      <div className="space-y-2">
        <Label
          htmlFor="hasOwnLocation"
          className="flex items-center gap-3 cursor-pointer"
        >
          <input
            id="hasOwnLocation"
            type="checkbox"
            {...register("hasOwnLocation")}
            className="w-5 h-5 text-conecta-blue border-gray-300 rounded focus:ring-conecta-blue"
          />
          <span className="text-gray-700 text-sm md:text-base">
            Possuo <span className="font-bold">local próprio</span> para
            atendimento
          </span>
        </Label>
      </div>

      {hasOwnLocation && (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 border p-4 rounded-md">
          <div>
            <Label>Rua / Avenida</Label>
            <Input
              {...register("location.street")}
              placeholder="Avenida Bento Ribeiro Dantas"
            />
          </div>
          <div>
            <Label>Número</Label>
            <Input {...register("location.number")} placeholder="123" />
          </div>
          <div>
            <Label>Complemento</Label>
            <Input
              {...register("location.complement")}
              placeholder="Apartamento, bloco, etc."
            />
          </div>
          <div>
            <Label>Bairro</Label>
            <Controller
              name="location.communityId"
              control={control}
              render={({ field }) => (
                <Select
                  onValueChange={field.onChange}
                  value={field.value}
                  defaultValue={field.value}
                  disabled={isLoadingCommunities || isError}                >
                  <SelectTrigger>
                    <SelectValue
                      placeholder={
                        isLoadingCommunities
                          ? "Carregando bairros..."
                          : "Selecione um bairro"
                      }
                    />
                  </SelectTrigger>
                  <SelectContent>
                    {isError && (
                      <SelectItem value="" disabled>
                        Erro ao carregar bairros
                      </SelectItem>
                    )}
                    {communities?.map((community) => (
                      <SelectItem key={community.id} value={community.id}>
                        {community.name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              )}
            />
          </div>
        </div>
      )}

      <div className="flex justify-between items-center">
        <Label className="text-base text-gray-800">Serviços</Label>
        <Button
          type="button"
          disabled={services.length === MAX_SERVICES}
          onClick={() =>
            append({
              name: "",
              description: "",
              price: 0,
              ownLocationPrice: null,
              images: [],
            })
          }
          variant="outline"
          size="sm"
        >
          <Plus className="w-4 h-4 mr-2" /> Adicionar Serviço
        </Button>
      </div>

      {services.map((service, index) => (
        <Card
          key={service.id}
          className="p-6 border border-conecta-blue/20 space-y-4"
        >
          <div className="flex justify-between items-center">
            <h3 className="font-semibold text-lg">Serviço {index + 1}</h3>
            <Button
              type="button"
              onClick={() => remove(index)}
              variant="ghost"
              className="text-red-500"
            >
              <X className="w-4 h-4 mr-1" /> Remover
            </Button>
          </div>

          <div className="space-y-2">
            <Label>Nome do Serviço</Label>
            <Input
              {...register(`services.${index}.name`)}
              placeholder="Ex: Manicure completa"
            />
          </div>

          <div className="space-y-2">
            <Label>Descrição do Serviço</Label>
            <Textarea
              {...register(`services.${index}.description`)}
              placeholder="Descreva seu serviço..."
              className="min-h-[120px]"
            />
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {hasOwnLocation && (
              <div>
                <Label>Valor do serviço no local próprio</Label>
                <Input
                  type="number"
                  min="0"
                  {...register(`services.${index}.ownLocationPrice`)}
                  placeholder="R$"
                />
              </div>
            )}
            <div>
              <Label>Valor do serviço a domicílio</Label>
              <Input
                type="number"
                min="0"
                {...register(`services.${index}.price`)}
                placeholder="R$"
              />
            </div>
          </div>

          <div className="space-y-2">
            <ServiceImageUpload index={index} />
          </div>
        </Card>
      ))}
    </div>
  );
};

export default ServicesStep;

const ServiceImageUpload = ({ index }: ServiceImageUploadProps) => {
  const { watch, setValue } = useFormContext();
  const images = watch(`services.${index}.images`);

  const handleImageUpload = useCallback(
    (event: React.ChangeEvent<HTMLInputElement>) => {
      const files = Array.from(event.target.files || []);
      const availableSlots = MAX_SERVICE_IMAGES - images.length;

      if (availableSlots <= 0) return;
      const filesToAdd = files.slice(0, availableSlots);

      const newImages = filesToAdd.map((file) => ({
        file,
        url: URL.createObjectURL(file),
      }));

      setValue(`services.${index}.images`, [...images, ...newImages], {
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
      setValue(`services.${index}.images`, updated, {
        shouldValidate: true,
      });
    },
    [images, index, setValue],
  );

  return (
    <div className="space-y-4">
      <Label>Foto do Serviço</Label>
      <label htmlFor={`services.${index}.images`}>
        <div className="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center cursor-pointer hover:border-conecta-blue transition-colors">
          <Upload className="w-12 h-12 mx-auto text-gray-400 mb-4" />
          <p className="text-gray-600 mb-2">
            Clique para carregar foto do serviço
          </p>
          <p className="text-sm text-gray-500">
            PNG, JPG até 5MB cada (máximo 1 foto)
          </p>
        </div>
        <input
          id={`services.${index}.images`}
          type="file"
          accept="image/*"
          multiple
          className="hidden"
          onChange={handleImageUpload}
          disabled={images.length >= MAX_SERVICE_IMAGES}
        />
      </label>

      <ServiceImageGallery images={images} onRemoveImage={handleRemoveImage} />
    </div>
  );
};

const ServiceImageGallery = ({
  images,
  onRemoveImage,
}: ServiceImageGalleryProps) => {
  if (images.length === 0) return null;

  return (
    <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
      {images.map((image, imgIndex) => (
        <ServiceImageItem
          key={imgIndex}
          image={image}
          onRemove={() => onRemoveImage(imgIndex)}
        />
      ))}
    </div>
  );
};

const ServiceImageItem = ({ image, onRemove }: ServiceImageItemProps) => (
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
