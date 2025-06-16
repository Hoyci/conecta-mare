import { useEffect, useRef, useState } from "react";
import { User, X } from "lucide-react";
import { Button } from "./button";

type ProfilePhotoUploadProps = {
  value: File[];
  onChange: (files: File[]) => void;
};

function ProfilePhotoUpload({ value, onChange }: ProfilePhotoUploadProps) {
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [preview, setPreview] = useState<string | null>(null);

  const handleImageSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files ? Array.from(e.target.files) : [];
    onChange(files);
  };

  const handleCircleClick = () => {
    fileInputRef.current?.click();
  };

  const handleRemoveImage = (e: React.MouseEvent) => {
    e.stopPropagation();
    onChange(null);
    if (fileInputRef.current) {
      fileInputRef.current.value = "";
    }
  };

  useEffect(() => {
    if (value.length > 0) {
      setPreview(URL.createObjectURL(value[0]));
    } else {
      setPreview(null);
    }
  }, [value]);

  return (
    <div className="flex gap-4">
      <div className="relative">
        <div
          onClick={handleCircleClick}
          className={`
            size-16 rounded-full border-2 border-dashed border-gray-300 
            flex items-center justify-center cursor-pointer
            transition-all duration-200 hover:border-gray-400 hover:bg-gray-50
            ${preview ? "border-solid border-gray-200" : ""}
          `}
          style={{
            backgroundImage: preview ? `url(${preview})` : "none",
            backgroundSize: "cover",
            backgroundPosition: "center",
            backgroundRepeat: "no-repeat",
          }}
        >
          {!preview && (
            <div className="flex flex-col items-center gap-2 text-gray-400">
              <User size={24} />
            </div>
          )}

          {preview && (
            <div className="absolute inset-0 rounded-full bg-black bg-opacity-0 hover:bg-opacity-20 transition-all duration-200 flex items-center justify-center">
              <User
                size={16}
                className="text-white opacity-0 hover:opacity-100 transition-opacity duration-200"
              />
            </div>
          )}
        </div>

        {preview && (
          <Button
            onClick={handleRemoveImage}
            size="icon"
            variant="destructive"
            className="absolute -bottom-0 right-0 size-4 rounded-full shadow-lg hover:scale-110 transition-transform duration-200"
          >
            <X size={8} />
          </Button>
        )}

        <input
          ref={fileInputRef}
          type="file"
          accept="image/*"
          onChange={handleImageSelect}
          className="hidden"
          id="file-input"
        />
      </div>

      <div className="h-full text-sm text-muted-foreground">
        <label htmlFor="file-input" className="text-lg font-bold text-gray-700">
          Foto de Perfil
        </label>
        <p>Formatos aceitos: JPG, PNG</p>
        <p>Tamanho máximo: 5MB</p>
      </div>
    </div>
  );
}

export default ProfilePhotoUpload;
