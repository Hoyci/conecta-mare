import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Button } from "@/components/ui/button";
import {
  Star,
  MapPin,
  Loader2,
  X,
  ChevronLeft,
  ChevronRight,
  MessageCircle,
} from "lucide-react";
import { ProfessionalUserResponse, Project, Service } from "@/types/user";
import { useQuery } from "@tanstack/react-query";
import { getProfessionalByID } from "@/services/user-service";
import { format } from "date-fns";
import { useEffect, useState } from "react";
import { createWhatsAppMessage, formatCentsToBRL } from "@/lib/utils";
import { getAnalytics } from "@/lib/analytics";

interface ProfessionalModalProps {
  userID: string;
  isOpen: boolean;
  onClose: () => void;
}

const ProfessionalModal = ({
  userID,
  isOpen,
  onClose,
}: ProfessionalModalProps) => {
  const [selectedService, setSelectedService] = useState<Project | null>(null);
  const [isImageModalOpen, setImageModalOpen] = useState(false);
  const [currentImageIndex, setCurrentImageIndex] = useState(0);

  const {
    data: professional,
    isLoading,
    isSuccess,
  } = useQuery<ProfessionalUserResponse>({
    queryKey: ["professionalUser", userID],
    queryFn: () => getProfessionalByID(userID),
    enabled: isOpen && !!userID,
  });

  useEffect(() => {
    if (professional) {
      const analytics = getAnalytics();
      analytics.track("profile_visited", {
        professional_id: userID,
        category: professional.category,
        subcategory: professional.subcategory,
        location: professional.location,
      });
    }
  }, [isSuccess, professional]);


  const [activeTab, setActiveTab] = useState("about");

  useEffect(() => {
    if (activeTab === "services") {
      const analytics = getAnalytics();
      analytics.track("viewed_services_tab", {
        user_id: userID,
        professional_id: userID,
      });
    }
  }, [activeTab, userID]);

  const handleContactService = (service: Service) => {
    const analytics = getAnalytics();
    analytics.track("service_interest", {
      user_id: userID,
      professional_id: userID,
      service_name: service.name,
      service_id: service.id,
    });


    window.open(createWhatsAppMessage(service.name, professional.phone), "_blank");
  };

  const goToPreviousImage = () => {
    if (!selectedService) return;
    setCurrentImageIndex((prevIndex: number) =>
      prevIndex === 0 ? selectedService.images.length - 1 : prevIndex - 1
    );
  };

  const goToNextImage = () => {
    if (!selectedService) return;
    setCurrentImageIndex((prevIndex: number) =>
      prevIndex === selectedService.images.length - 1 ? 0 : prevIndex + 1
    );
  };

  const openImageModal = (project: Project) => {
    setSelectedService(project);
    setCurrentImageIndex(0);
    setImageModalOpen(true);
  };

  if (isLoading || !professional) {
    return (
      <Dialog open={isOpen} onOpenChange={onClose}>
        <DialogContent
          className="sm:max-w-[800px] max-h-[90vh] overflow-y-auto"
          onInteractOutside={(e) => e.preventDefault()}
        >
          <DialogHeader>
            <DialogTitle>Perfil do Profissional</DialogTitle>
            <DialogDescription>
              Conheça mais sobre o serviço oferecido
            </DialogDescription>
          </DialogHeader>

          <div className="flex justify-center items-center py-10">
            <Loader2 className="w-6 h-6 mr-2 animate-spin text-conecta-blue" />
            <span className="text-gray-600">Carregando informações...</span>
          </div>
        </DialogContent>
      </Dialog>
    );
  }

  return (
    <>
      <Dialog open={isOpen} onOpenChange={onClose}>
        <DialogContent
          className="sm:max-w-[800px] max-h-[90vh] overflow-y-auto"
          onInteractOutside={(e) => e.preventDefault()}
        >
          <DialogHeader>
            <DialogTitle>Perfil do Profissional</DialogTitle>
            <DialogDescription>
              Conheça mais sobre o serviço oferecido
            </DialogDescription>
          </DialogHeader>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 py-4">
            {/* Coluna do perfil */}
            <div className="md:col-span-1">
              <div className="rounded-lg overflow-hidden mb-4">
                <img
                  src={professional.profileImage}
                  alt={professional.fullName}
                  className="w-full h-auto object-cover"
                />
              </div>
              <div className="space-y-4">
                <div className="bg-gray-50 p-4 rounded-md">
                  <h3 className="text-lg font-semibold mb-2">
                    {professional.fullName}
                  </h3>
                  <p className="text-conecta-blue font-medium">
                    {professional.subcategory.name}
                  </p>

                  <div className="flex items-center mt-2">
                    <Star
                      size={18}
                      className="text-yellow-400 fill-yellow-400"
                    />
                    <span className="text-sm ml-1 font-medium">
                      {professional.rating.toFixed(1)}
                    </span>
                    <span className="text-sm text-gray-500 ml-1">
                      (100 avaliações)
                    </span>
                  </div>

                  <div className="flex items-center text-gray-600 text-sm mt-2">
                    <MapPin size={16} className="mr-1" />
                    <span className="capitalize">
                      {professional.location.communityName}
                    </span>
                  </div>
                </div>
              </div>
            </div>

            {/* Coluna com as tabs */}
            <div className="md:col-span-2">
              <Tabs value={activeTab} onValueChange={setActiveTab}>
                <TabsList className="w-full">
                  <TabsTrigger value="about" className="flex-1">
                    Sobre
                  </TabsTrigger>
                  <TabsTrigger value="services" className="flex-1">
                    Serviços
                  </TabsTrigger>
                  <TabsTrigger value="portfolio" className="flex-1">
                    Portfólio
                  </TabsTrigger>
                  <TabsTrigger value="reviews" className="flex-1">
                    Avaliações
                  </TabsTrigger>
                </TabsList>

                {/* Tab: Sobre */}
                <TabsContent value="about" className="mt-4">
                  <div className="space-y-4">
                    <div>
                      <h3 className="font-semibold mb-2 text-lg">Sobre</h3>
                      <p className="text-gray-700">
                        {professional.jobDescription}
                      </p>
                    </div>
                    <div>
                      <h3 className="font-semibold mb-4 text-lg">
                        Certificações
                      </h3>

                      <div className="flex flex-col gap-4">
                        {professional.certifications.map(
                          ({ id, courseName, institution, startDate, endDate }) => (
                            <div
                              key={id}
                              className="py-2 px-4 border border-gray-200 rounded-md shadow-sm bg-white"
                            >
                              <h4 className="font-semibold text-gray-900 mb-1">
                                {courseName}
                              </h4>
                              <p className="text-gray-700 mb-1">
                                {institution}
                              </p>
                              <p className="text-gray-500 text-sm">
                                {startDate
                                  ? format(new Date(startDate), "dd/MM/yyyy")
                                  : "—"}{" "}
                                —{" "}
                                {endDate
                                  ? format(new Date(endDate), "dd/MM/yyyy")
                                  : "Presente"}
                              </p>
                            </div>
                          )
                        )}
                      </div>
                    </div>
                  </div>
                </TabsContent>

                {/* Tab: Serviços */}
                <TabsContent value="services" className="mt-4">
                  <div>
                    <h3 className="font-semibold mb-2 text-lg">
                      Serviços ofertados
                    </h3>
                  </div>
                  <div className="flex flex-col gap-4">
                    {professional.services.map((service) => (
                      <div
                        key={service.id}
                        className="py-3 px-4 border border-gray-200 rounded-md shadow-sm bg-white flex justify-between items-center"
                      >
                        <div>
                          <h4 className="font-semibold text-gray-900 mb-1">
                            {service.name}
                          </h4>
                          <p className="text-gray-700 mb-1">{service.description}</p>
                          <p className="text-gray-700 mb-1">
                            {service.price === 0
                              ? "Gratuito"
                              : `${formatCentsToBRL(service.price)}`}
                          </p>
                        </div>
                        <Button
                          size="sm"
                          className="bg-conecta-green hover:bg-conecta-green-dark text-white"
                          onClick={() => handleContactService(service)}
                        >
                          <MessageCircle size={16} className="mr-1" />
                          Contatar
                        </Button>
                      </div>
                    ))}
                  </div>
                </TabsContent>

                {/* Tab: Portfólio */}
                <TabsContent value="portfolio">
                  <div>
                    <h3 className="font-semibold mb-2 text-lg">
                      Projetos realizados
                    </h3>
                    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-2 gap-4">
                      {professional.projects.map((project: Project) => (
                        <div
                          key={project.name}
                          onClick={() => openImageModal(project)}
                          className="cursor-pointer bg-gray-50 rounded-md overflow-hidden shadow hover:shadow-md transition"
                        >
                          <img
                            src={project.images?.[0]?.url}
                            alt={project.name}
                            className="w-full h-40 object-cover"
                          />
                          <div className="p-3">
                            <h4 className="font-semibold text-sm">
                              {project.name}
                            </h4>
                            <p className="text-sm font-thin">
                              {project.description}
                            </p>
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>
                </TabsContent>

                {/* Tab: Avaliações */}
                <TabsContent value="reviews" className="mt-4">
                  <div className="space-y-4">
                    <div className="bg-gray-50 p-4 rounded-md">
                      <h3 className="font-semibold mb-2">
                        Classificação geral
                      </h3>
                      <div className="flex items-center">
                        <div className="flex">
                          {[...Array(5)].map((_, i) => (
                            <Star
                              key={i}
                              size={24}
                              className={
                                i < Math.floor(professional.rating)
                                  ? "text-yellow-400 fill-yellow-400"
                                  : "text-gray-300"
                              }
                            />
                          ))}
                        </div>
                        <span className="ml-2 text-2xl font-semibold">
                          {professional.rating.toFixed(1)}
                        </span>
                        <span className="ml-2 text-gray-500">
                          (100 avaliações)
                        </span>
                      </div>
                    </div>
                    <div>
                      <h3 className="font-semibold mb-2">
                        Comentários recentes
                      </h3>
                    </div>
                  </div>
                </TabsContent>
              </Tabs>
            </div>
          </div>
        </DialogContent>
      </Dialog>

      {/* Modal para imagens do portfólio */}
      <Dialog open={isImageModalOpen} onOpenChange={setImageModalOpen}>
        <DialogContent className="bg-black/90 p-0 max-w-4xl max-h-[90vh] flex items-center justify-center border-none">
          <button
            onClick={() => setImageModalOpen(false)}
            className="absolute top-4 right-4 z-20 text-white bg-black/50 rounded-full p-2 hover:bg-black/80 transition"
          >
            <X size={24} />
          </button>

          <img
            src={selectedService?.images[currentImageIndex]?.url}
            alt={`${selectedService?.name} - Imagem ${currentImageIndex + 1}`}
            className="w-full h-full max-h-[80vh] object-contain"
          />

          {selectedService?.images.length > 1 && (
            <>
              <button
                onClick={goToPreviousImage}
                className="absolute left-4 top-1/2 z-20 transform -translate-y-1/2 bg-black/50 text-white rounded-full p-2 hover:bg-black/80 transition"
              >
                <ChevronLeft size={24} />
              </button>

              <button
                onClick={goToNextImage}
                className="absolute right-4 top-1/2 z-20 transform -translate-y-1/2 bg-black/50 text-white rounded-full p-2 hover:bg-black/80 transition"
              >
                <ChevronRight size={24} />
              </button>

              <div className="absolute bottom-4 left-1/2 z-20 transform -translate-x-1/2 bg-black/50 text-white px-3 py-1 rounded-full text-sm">
                {currentImageIndex + 1} / {selectedService.images.length}
              </div>
            </>
          )}
        </DialogContent>
      </Dialog>
    </>
  );
};

export default ProfessionalModal;

