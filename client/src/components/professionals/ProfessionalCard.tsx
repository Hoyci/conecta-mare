import { useState } from "react";
import { MapPin, Star } from "lucide-react";
import ProfessionalModal from "./ProfessionalModal";
import { ProfessionalUser } from "@/types/user";

interface ProfessionalCardProps {
  professional: ProfessionalUser;
}

const ProfessionalCard = ({ professional }: ProfessionalCardProps) => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  return (
    <>
      <div
        onClick={() => setIsModalOpen(true)}
        className="bg-white rounded-lg shadow-md overflow-hidden transition-transform hover:shadow-lg hover:-translate-y-1 cursor-pointer"
      >
        <div className="relative">
          <img
            src={professional.profileImage}
            alt={professional.fullName}
            className="w-full h-48 object-cover"
          />
          {
            <div className="absolute top-2 right-2 bg-conecta-blue text-white text-xs font-bold px-2 py-1 rounded-full">
              Verificado
            </div>
          }
        </div>
        <div className="p-4">
          <div className="flex items-center justify-between mb-2">
            <h3 className="text-lg font-semibold">{professional.fullName}</h3>
            <div className="flex items-center">
              <Star size={16} className="text-yellow-400 fill-yellow-400" />
              <span className="text-sm ml-1 font-medium">
                {professional.rating.toFixed(1)}
              </span>
            </div>
          </div>
          <p className="text-conecta-blue font-medium mb-2">
            {professional.subcategoryName}
          </p>
          <div className="flex items-center text-gray-500 text-sm mb-2">
            <MapPin size={14} className="mr-1" />
            <span className="capitalize">{professional.location}</span>
          </div>
          <p className="mt-3 text-sm text-gray-600 line-clamp-2">
            {professional.jobDescription}
          </p>
        </div>
      </div>

      <ProfessionalModal
        userID={professional.userId}
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
      />
    </>
  );
};

export default ProfessionalCard;
