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
  Clock,
  Calendar,
  Phone,
  Mail,
  Shield,
  Award,
  User,
} from "lucide-react";
import { Professional } from "@/types/professional";
import { useToast } from "@/components/ui/use-toast";
import { ProfessionalUser } from "@/types/user";

interface ProfessionalModalProps {
  professional: ProfessionalUser;
  isOpen: boolean;
  onClose: () => void;
}

const ProfessionalModal = ({
  professional,
  isOpen,
  onClose,
}: ProfessionalModalProps) => {
  const { toast } = useToast();

  const handleContact = () => {
    toast({
      title: "Solicitação enviada!",
      description: `Sua mensagem foi enviada para ${professional.fullName}. Aguarde o contato!`,
      variant: "default",
    });
  };

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[800px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Perfil do Profissional</DialogTitle>
          <DialogDescription>
            Conheça mais sobre o serviço oferecido
          </DialogDescription>
        </DialogHeader>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 py-4">
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
                  {professional.subcategoryName}
                </p>

                {/* <div className="flex items-center mt-2"> */}
                {/*   <Star size={18} className="text-yellow-400 fill-yellow-400" /> */}
                {/*   <span className="text-sm ml-1 font-medium"> */}
                {/*     {professional.rating.toFixed(1)} */}
                {/*   </span> */}
                {/*   <span className="text-sm text-gray-500 ml-1"> */}
                {/*     ({professional.reviewsCount} avaliações) */}
                {/*   </span> */}
                {/* </div> */}

                {/* <div className="flex items-center text-gray-600 text-sm mt-2"> */}
                {/*   <MapPin size={16} className="mr-1" /> */}
                {/*   <span>{professional.location}</span> */}
                {/* </div> */}

                {/* <div className="flex items-center text-gray-600 text-sm mt-1"> */}
                {/*   <Clock size={16} className="mr-1" /> */}
                {/*   <span>Tempo de resposta: {professional.responseTimes}</span> */}
                {/* </div> */}

                {/* <div className="mt-4 flex flex-wrap gap-2"> */}
                {/*   {professional.specialties.map((specialty, index) => ( */}
                {/*     <span */}
                {/*       key={index} */}
                {/*       className="bg-gray-100 text-gray-600 text-xs px-2 py-1 rounded-md" */}
                {/*     > */}
                {/*       {specialty} */}
                {/*     </span> */}
                {/*   ))} */}
                {/* </div> */}
              </div>

              <div className="flex gap-2 mt-4">
                <Button
                  className="w-full bg-conecta-green hover:bg-conecta-green-dark text-white"
                  onClick={handleContact}
                >
                  Contatar
                </Button>
              </div>
            </div>
          </div>

          <div className="md:col-span-2">
            <Tabs defaultValue="about">
              <TabsList className="w-full">
                <TabsTrigger value="about" className="flex-1">
                  Sobre
                </TabsTrigger>
                <TabsTrigger value="services" className="flex-1">
                  Serviços
                </TabsTrigger>
                <TabsTrigger value="reviews" className="flex-1">
                  Avaliações
                </TabsTrigger>
                <TabsTrigger value="portfolio" className="flex-1">
                  Portfólio
                </TabsTrigger>
              </TabsList>

              {/* <TabsContent value="about" className="mt-4"> */}
              {/*   <div className="space-y-4"> */}
              {/*     <div> */}
              {/*       <h3 className="font-semibold mb-2">Sobre</h3> */}
              {/*       <p className="text-gray-600">{professional.description}</p> */}
              {/*     </div> */}
              {/**/}
              {/*     <div> */}
              {/*       <h3 className="font-semibold mb-2">Experiência</h3> */}
              {/*       <p className="text-gray-600">{professional.experience}</p> */}
              {/*     </div> */}
              {/**/}
              {/*     <div> */}
              {/*       <h3 className="font-semibold mb-2">Credenciais</h3> */}
              {/*       <div className="space-y-2"> */}
              {/*         {professional.credentials.map((credential, index) => ( */}
              {/*           <div key={index} className="flex items-start"> */}
              {/*             <Shield */}
              {/*               size={18} */}
              {/*               className="text-conecta-blue mt-1 mr-2" */}
              {/*             /> */}
              {/*             <span className="text-gray-600">{credential}</span> */}
              {/*           </div> */}
              {/*         ))} */}
              {/*       </div> */}
              {/*     </div> */}
              {/**/}
              {/*     <div> */}
              {/*       <h3 className="font-semibold mb-2"> */}
              {/*         Informações de Contato */}
              {/*       </h3> */}
              {/*       <div className="space-y-2"> */}
              {/*         <div className="flex items-center"> */}
              {/*           <Phone size={18} className="text-conecta-blue mr-2" /> */}
              {/*           <span className="text-gray-600"> */}
              {/*             Disponível após contratação */}
              {/*           </span> */}
              {/*         </div> */}
              {/*         <div className="flex items-center"> */}
              {/*           <Mail size={18} className="text-conecta-blue mr-2" /> */}
              {/*           <span className="text-gray-600"> */}
              {/*             Disponível após contratação */}
              {/*           </span> */}
              {/*         </div> */}
              {/*         <div className="flex items-center"> */}
              {/*           <Calendar */}
              {/*             size={18} */}
              {/*             className="text-conecta-blue mr-2" */}
              {/*           /> */}
              {/*           <span className="text-gray-600"> */}
              {/*             Disponibilidade: {professional.availability} */}
              {/*           </span> */}
              {/*         </div> */}
              {/*       </div> */}
              {/*     </div> */}
              {/*   </div> */}
              {/* </TabsContent> */}
              {/**/}
              {/* <TabsContent value="services" className="mt-4"> */}
              {/*   <div className="space-y-4"> */}
              {/*     <div> */}
              {/*       <h3 className="font-semibold mb-2">Serviços oferecidos</h3> */}
              {/*       <ul className="space-y-3"> */}
              {/*         {professional.services.map((service, index) => ( */}
              {/*           <li key={index} className="bg-gray-50 p-3 rounded-md"> */}
              {/*             <div className="flex justify-between"> */}
              {/*               <span className="font-medium">{service.name}</span> */}
              {/*               <span className="text-conecta-blue font-semibold"> */}
              {/*                 {typeof service.price === "string" */}
              {/*                   ? service.price */}
              {/*                   : `R$ ${service.price.toFixed(2)}`} */}
              {/*               </span> */}
              {/*             </div> */}
              {/*             <p className="text-sm text-gray-600 mt-1"> */}
              {/*               {service.description} */}
              {/*             </p> */}
              {/*           </li> */}
              {/*         ))} */}
              {/*       </ul> */}
              {/*     </div> */}
              {/**/}
              {/*     <div> */}
              {/*       <h3 className="font-semibold mb-2">Áreas de atendimento</h3> */}
              {/*       <p className="text-gray-600">{professional.serviceAreas}</p> */}
              {/*     </div> */}
              {/*   </div> */}
              {/* </TabsContent> */}
              {/**/}
              {/* <TabsContent value="reviews" className="mt-4"> */}
              {/*   <div className="space-y-4"> */}
              {/*     <div className="bg-gray-50 p-4 rounded-md"> */}
              {/*       <h3 className="font-semibold mb-2">Classificação geral</h3> */}
              {/*       <div className="flex items-center"> */}
              {/*         <div className="flex"> */}
              {/*           {[...Array(5)].map((_, i) => ( */}
              {/*             <Star */}
              {/*               key={i} */}
              {/*               size={24} */}
              {/*               className={ */}
              {/*                 i < Math.floor(professional.rating) */}
              {/*                   ? "text-yellow-400 fill-yellow-400" */}
              {/*                   : "text-gray-300" */}
              {/*               } */}
              {/*             /> */}
              {/*           ))} */}
              {/*         </div> */}
              {/*         <span className="ml-2 text-2xl font-semibold"> */}
              {/*           {professional.rating.toFixed(1)} */}
              {/*         </span> */}
              {/*         <span className="ml-2 text-gray-500"> */}
              {/*           ({professional.reviewsCount} avaliações) */}
              {/*         </span> */}
              {/*       </div> */}
              {/*     </div> */}
              {/**/}
              {/*     <div> */}
              {/*       <h3 className="font-semibold mb-2">Comentários recentes</h3> */}
              {/*       <div className="space-y-4"> */}
              {/*         {professional.reviews.map((review, index) => ( */}
              {/*           <div */}
              {/*             key={index} */}
              {/*             className="border-b pb-4 last:border-b-0" */}
              {/*           > */}
              {/*             <div className="flex justify-between"> */}
              {/*               <div className="flex items-center"> */}
              {/*                 <User size={18} className="text-gray-400 mr-2" /> */}
              {/*                 <span className="font-medium">{review.user}</span> */}
              {/*               </div> */}
              {/*               <div className="flex"> */}
              {/*                 {[...Array(5)].map((_, i) => ( */}
              {/*                   <Star */}
              {/*                     key={i} */}
              {/*                     size={16} */}
              {/*                     className={ */}
              {/*                       i < review.rating */}
              {/*                         ? "text-yellow-400 fill-yellow-400" */}
              {/*                         : "text-gray-300" */}
              {/*                     } */}
              {/*                   /> */}
              {/*                 ))} */}
              {/*               </div> */}
              {/*             </div> */}
              {/*             <p className="text-gray-600 mt-2">{review.comment}</p> */}
              {/*             <p className="text-gray-400 text-sm mt-1"> */}
              {/*               {review.date} */}
              {/*             </p> */}
              {/*           </div> */}
              {/*         ))} */}
              {/*       </div> */}
              {/*     </div> */}
              {/*   </div> */}
              {/* </TabsContent> */}
              {/**/}
              {/* <TabsContent value="portfolio" className="mt-4"> */}
              {/*   <div className="space-y-4"> */}
              {/*     <h3 className="font-semibold mb-2">Trabalhos anteriores</h3> */}
              {/**/}
              {/*     <div className="grid grid-cols-2 sm:grid-cols-3 gap-3"> */}
              {/*       {professional.portfolioImages?.map((image, index) => ( */}
              {/*         <div key={index} className="rounded-md overflow-hidden"> */}
              {/*           <img */}
              {/*             src={ */}
              {/*               image.url || */}
              {/*               "https://via.placeholder.com/300x200?text=Portfolio" */}
              {/*             } */}
              {/*             alt={image.title} */}
              {/*             className="w-full h-40 object-cover" */}
              {/*           /> */}
              {/*         </div> */}
              {/*       ))} */}
              {/*     </div> */}
              {/**/}
              {/*     <div> */}
              {/*       <h3 className="font-semibold mb-2"> */}
              {/*         Certificações e Prêmios */}
              {/*       </h3> */}
              {/*       <div className="space-y-2"> */}
              {/*         {professional.awards?.map((award, index) => ( */}
              {/*           <div key={index} className="flex items-start"> */}
              {/*             <Award */}
              {/*               size={18} */}
              {/*               className="text-conecta-green mt-1 mr-2" */}
              {/*             /> */}
              {/*             <div> */}
              {/*               <p className="font-medium">{award.title}</p> */}
              {/*               <p className="text-sm text-gray-600"> */}
              {/*                 {award.description} */}
              {/*               </p> */}
              {/*             </div> */}
              {/*           </div> */}
              {/*         ))} */}
              {/*       </div> */}
              {/*     </div> */}
              {/*   </div> */}
              {/* </TabsContent> */}
            </Tabs>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
};

export default ProfessionalModal;
