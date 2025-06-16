import { useState } from "react";
import { Search, Filter } from "lucide-react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { Slider } from "@/components/ui/slider";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
  SheetFooter,
  SheetClose,
} from "@/components/ui/sheet";

interface ProfessionalFiltersProps {
  onFilterChange: (filters: any) => void;
  categories: string[];
}

const ProfessionalFilters = ({
  onFilterChange,
  categories,
}: ProfessionalFiltersProps) => {
  const [searchTerm, setSearchTerm] = useState("");
  const [selectedCategories, setSelectedCategories] = useState<string[]>([]);
  const [minRating, setMinRating] = useState(0);
  const [verified, setVerified] = useState(false);
  const [availability, setAvailability] = useState<string[]>([]);
  const [distance, setDistance] = useState(50);

  const handleSearch = () => {
    onFilterChange({
      searchTerm,
      categories: selectedCategories,
      minRating,
      verified,
      availability,
      distance,
    });
  };

  const handleCategoryToggle = (category: string) => {
    setSelectedCategories((prev) =>
      prev.includes(category)
        ? prev.filter((c) => c !== category)
        : [...prev, category],
    );
  };

  const handleAvailabilityToggle = (value: string) => {
    setAvailability((prev) =>
      prev.includes(value) ? prev.filter((a) => a !== value) : [...prev, value],
    );
  };

  const handleReset = () => {
    setSearchTerm("");
    setSelectedCategories([]);
    setMinRating(0);
    setVerified(false);
    setAvailability([]);
    setDistance(50);

    onFilterChange({
      searchTerm: "",
      categories: [],
      minRating: 0,
      verified: false,
      availability: [],
      distance: 50,
    });
  };

  return (
    <div className="bg-white rounded-lg shadow p-4 mb-6">
      <div className="flex flex-col md:flex-row gap-4">
        <div className="relative flex-grow">
          <Search
            className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
            size={20}
          />
          <Input
            placeholder="Buscar profissionais, serviços..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="pl-10 py-6"
          />
        </div>

        <Sheet>
          <SheetTrigger asChild>
            <Button variant="outline" className="flex items-center gap-2">
              <Filter size={18} />
              <span>Filtros</span>
            </Button>
          </SheetTrigger>
          <SheetContent className="w-full sm:max-w-md">
            <SheetHeader>
              <SheetTitle>Filtros</SheetTitle>
              <SheetDescription>
                Refine sua busca por profissionais
              </SheetDescription>
            </SheetHeader>
            <div className="py-4 space-y-6">
              <div>
                <h4 className="text-sm font-medium mb-3">Categorias</h4>
                <div className="space-y-2">
                  {categories.map((category) => (
                    <div key={category} className="flex items-center space-x-2">
                      <Checkbox
                        id={`category-${category}`}
                        checked={selectedCategories.includes(category)}
                        onCheckedChange={() => handleCategoryToggle(category)}
                      />
                      <label
                        htmlFor={`category-${category}`}
                        className="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                      >
                        {category}
                      </label>
                    </div>
                  ))}
                </div>
              </div>

              <div>
                <div className="flex items-center justify-between mb-3">
                  <h4 className="text-sm font-medium">Avaliação mínima</h4>
                  <span className="text-sm font-medium">
                    {minRating} estrelas
                  </span>
                </div>
                <Slider
                  min={0}
                  max={5}
                  step={0.5}
                  value={[minRating]}
                  onValueChange={(values) => setMinRating(values[0])}
                  className="my-4"
                />
              </div>

              <div>
                <h4 className="text-sm font-medium mb-3">Disponibilidade</h4>
                <div className="space-y-2">
                  {[
                    "Segunda a Sexta",
                    "Fins de semana",
                    "Horário comercial",
                    "Noites",
                    "24/7",
                  ].map((day) => (
                    <div key={day} className="flex items-center space-x-2">
                      <Checkbox
                        id={`day-${day}`}
                        checked={availability.includes(day)}
                        onCheckedChange={() => handleAvailabilityToggle(day)}
                      />
                      <label
                        htmlFor={`day-${day}`}
                        className="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                      >
                        {day}
                      </label>
                    </div>
                  ))}
                </div>
              </div>

              <div>
                <div className="flex items-center justify-between mb-3">
                  <h4 className="text-sm font-medium">Distância máxima</h4>
                  <span className="text-sm font-medium">{distance} km</span>
                </div>
                <Slider
                  min={1}
                  max={100}
                  step={1}
                  value={[distance]}
                  onValueChange={(values) => setDistance(values[0])}
                  className="my-4"
                />
              </div>

              <div className="flex items-center space-x-2">
                <Checkbox
                  id="verified"
                  checked={verified}
                  onCheckedChange={() => setVerified(!verified)}
                />
                <label
                  htmlFor="verified"
                  className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                >
                  Mostrar apenas profissionais verificados
                </label>
              </div>
            </div>
            <SheetFooter className="flex-row gap-3 sm:justify-end mt-4">
              <Button variant="outline" onClick={handleReset}>
                Limpar filtros
              </Button>
              <SheetClose asChild>
                <Button onClick={handleSearch}>Aplicar filtros</Button>
              </SheetClose>
            </SheetFooter>
          </SheetContent>
        </Sheet>

        <Button
          onClick={handleSearch}
          className="bg-conecta-blue hover:bg-conecta-blue-dark"
        >
          Buscar
        </Button>
      </div>

      {selectedCategories.length > 0 && (
        <div className="flex flex-wrap gap-2 mt-4">
          <span className="text-sm text-gray-500">Filtros:</span>
          {selectedCategories.map((category) => (
            <span
              key={category}
              onClick={() => handleCategoryToggle(category)}
              className="bg-gray-100 text-gray-600 text-xs px-2 py-1 rounded-full flex items-center cursor-pointer hover:bg-gray-200"
            >
              {category} ×
            </span>
          ))}
          {(minRating > 0 || verified) && (
            <>
              {minRating > 0 && (
                <span className="bg-gray-100 text-gray-600 text-xs px-2 py-1 rounded-full flex items-center">
                  ≥ {minRating} estrelas
                </span>
              )}
              {verified && (
                <span className="bg-gray-100 text-gray-600 text-xs px-2 py-1 rounded-full flex items-center">
                  Verificado
                </span>
              )}
            </>
          )}
          {selectedCategories.length > 0 && (
            <button
              onClick={handleReset}
              className="text-xs text-conecta-blue hover:underline ml-2"
            >
              Limpar todos
            </button>
          )}
        </div>
      )}
    </div>
  );
};

export default ProfessionalFilters;
