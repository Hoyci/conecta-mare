import { getCategoriesWithUserCount } from "@/services/categories-service";
import {
  CategoryWithUserCount,
  CategoryWithUserCountResponse,
} from "@/types/categories";
import { useQuery } from "@tanstack/react-query";
import { Link } from "react-router-dom";

const CategorySection = () => {
  const { data: { categories } = {}, isLoading } =
    useQuery<CategoryWithUserCountResponse>({
      queryKey: ["subcategories"],
      queryFn: getCategoriesWithUserCount,
    });

  return (
    <section className="py-16 bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-12">
          <h2 className="text-3xl lg:text-4xl font-bold text-gray-900 mb-4">
            Explore por Categorias
          </h2>
          <p className="text-lg text-gray-600 max-w-2xl mx-auto">
            Encontre profissionais especializados em diversas áreas para atender
            suas necessidades
          </p>
        </div>

        <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-4 md:gap-6">
          {isLoading
            ? Array.from({ length: 8 }).map((_, index) => (
                <div
                  key={index}
                  className="h-[136px] w-full bg-gray-200 rounded-lg relative overflow-hidden"
                >
                  <div className="absolute inset-0 -translate-x-full bg-gradient-to-r from-gray-200 via-white to-gray-200 animate-shimmer" />
                </div>
              ))
            : categories?.map((category: CategoryWithUserCount) => (
                <Link
                  key={category.id}
                  to={`/professionals?category=${category.id}`}
                  className="bg-white rounded-lg shadow-sm hover:shadow-md transition-shadow p-4 flex flex-col items-center text-center"
                >
                  <span className="text-4xl mb-3">{category.icon}</span>
                  <h3 className="font-medium text-lg mb-1">{category.name}</h3>
                  <p className="text-sm text-gray-500">
                    {category.user_count} profissionais
                  </p>
                </Link>
              ))}
        </div>

        <div className="text-center mt-10">
          <Link
            to="/professionals"
            className="inline-block text-conecta-blue font-medium hover:text-conecta-blue-dark hover:underline"
          >
            Clique aqui para pesquisar →
          </Link>
        </div>
      </div>
    </section>
  );
};

export default CategorySection;
