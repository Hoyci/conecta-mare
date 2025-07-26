import { useState, useEffect } from "react";
import { useLocation } from "react-router-dom";
import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import ProfessionalCard from "@/components/professionals/ProfessionalCard";
import ProfessionalFilters from "@/components/professionals/ProfessionalFilters";
import { professionals, categories } from "@/data/professionalsData";
import { Professional } from "@/types/professional";
import { useQuery } from "@tanstack/react-query";
import { ProfessionalUser, ProfessionalUsersResponse } from "@/types/user";
import { getProfessionals } from "@/services/user-service";

const Professionals = () => {
  // const location = useLocation();
  // const [filteredProfessionals, setFilteredProfessionals] =
  //   useState<Professional[]>(professionals);
  // const [filters, setFilters] = useState<any>({
  //   searchTerm: "",
  //   categories: [],
  //   minRating: 0,
  //   verified: false,
  //   availability: [],
  //   distance: 50,
  // });
  // const [loading, setLoading] = useState(false);
  //
  // useEffect(() => {
  //   // Check for category param in URL
  //   const params = new URLSearchParams(location.search);
  //   const categoryParam = params.get("category");
  //
  //   if (categoryParam) {
  //     setFilters((prevFilters) => ({
  //       ...prevFilters,
  //       categories: [categoryParam],
  //     }));
  //   }
  // }, [location.search]);
  //
  // useEffect(() => {
  //   setLoading(true);
  //
  //   // Simulate API call with filters
  //   const timer = setTimeout(() => {
  //     const filtered = professionals.filter((professional) => {
  //       // Filter by search term
  //       if (
  //         filters.searchTerm &&
  //         !professional.name
  //           .toLowerCase()
  //           .includes(filters.searchTerm.toLowerCase()) &&
  //         !professional.profession
  //           .toLowerCase()
  //           .includes(filters.searchTerm.toLowerCase()) &&
  //         !professional.description
  //           .toLowerCase()
  //           .includes(filters.searchTerm.toLowerCase())
  //       ) {
  //         return false;
  //       }
  //
  //       // Filter by categories
  //       if (
  //         filters.categories.length > 0 &&
  //         !filters.categories.includes(professional.profession)
  //       ) {
  //         return false;
  //       }
  //
  //       // Filter by minimum rating
  //       if (professional.rating < filters.minRating) {
  //         return false;
  //       }
  //
  //       // Filter by verification
  //       if (filters.verified && !professional.isVerified) {
  //         return false;
  //       }
  //
  //       // Filter by availability (simplified)
  //       if (filters.availability.length > 0) {
  //         // This is a simplified check - in a real app, you'd have more detailed availability data
  //         if (!professional.availability.includes(filters.availability[0])) {
  //           return false;
  //         }
  //       }
  //
  //       return true;
  //     });
  //
  //     setFilteredProfessionals(filtered);
  //     setLoading(false);
  //   }, 500); // Simulate network delay
  //
  //   return () => clearTimeout(timer);
  // }, [filters]);
  //
  // const handleFilterChange = (newFilters: any) => {
  //   setFilters(newFilters);
  // };
  //
  const { data: { professionals } = [], isLoading } =
    useQuery<ProfessionalUsersResponse>({
      queryKey: ["professionalUsers"],
      queryFn: getProfessionals,
    });

  return (
    <div className="min-h-screen flex flex-col">
      <Navbar />
      <main className="flex-grow bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
          <div className="mb-10">
            <h1 className="text-3xl font-bold text-gray-900 mb-2">
              Encontre Profissionais
            </h1>
            <p className="text-lg text-gray-600">
              Conecte-se com os melhores profissionais da sua região
            </p>
          </div>

          <ProfessionalFilters
            onFilterChange={() => console.log("oi")}
            categories={categories}
          />

          {isLoading ? (
            <div className="flex justify-center items-center py-12">
              <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-conecta-blue"></div>
            </div>
          ) : professionals.length > 0 ? (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {professionals.map((professional: ProfessionalUser) => (
                <ProfessionalCard
                  key={professional.userId}
                  professional={professional}
                />
              ))}
            </div>
          ) : (
            <div className="text-center py-12">
              <h3 className="text-xl font-medium text-gray-900 mb-2">
                Nenhum profissional encontrado
              </h3>
              <p className="text-gray-600 mb-6">
                Tente ajustar seus filtros para encontrar mais resultados
              </p>
              <button
                onClick={
                  () => console.log("clear")
                  // setFilters({
                  //   searchTerm: "",
                  //   categories: [],
                  //   minRating: 0,
                  //   verified: false,
                  //   availability: [],
                  //   distance: 50,
                  // })
                }
                className="text-conecta-blue hover:underline font-medium"
              >
                Limpar todos os filtros
              </button>
            </div>
          )}
        </div>
      </main>
      <Footer />
    </div>
  );
};

export default Professionals;
