import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { CheckCircle2 } from "lucide-react";

const CtaSection = () => {
  return (
    <section className="py-16 lg:py-24">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="bg-gradient-to-br from-conecta-blue to-conecta-blue-dark rounded-2xl overflow-hidden">
          <div className="md:grid md:grid-cols-2">
            <div className="p-8 md:p-12 flex flex-col justify-center">
              <h2 className="text-3xl md:text-4xl font-bold text-white mb-6">
                É um profissional qualificado?
              </h2>
              <p className="text-white/90 text-lg mb-8 max-w-md">
                Junte-se à nossa comunidade de profissionais e expanda seu
                negócio. O ConectaMaré conecta você a novos clientes todos os
                dias.
              </p>

              <div className="space-y-3 mb-8">
                <div className="flex items-center">
                  <CheckCircle2 className="text-conecta-green mr-3 h-6 w-6 flex-shrink-0" />
                  <span className="text-white">Cadastro gratuito e fácil</span>
                </div>
                <div className="flex items-center">
                  <CheckCircle2 className="text-conecta-green mr-3 h-6 w-6 flex-shrink-0" />
                  <span className="text-white">
                    Conecte-se com clientes próximos a você
                  </span>
                </div>
                <div className="flex items-center">
                  <CheckCircle2 className="text-conecta-green mr-3 h-6 w-6 flex-shrink-0" />
                  <span className="text-white">
                    Gerencie sua agenda de forma eficiente
                  </span>
                </div>
                <div className="flex items-center">
                  <CheckCircle2 className="text-conecta-green mr-3 h-6 w-6 flex-shrink-0" />
                  <span className="text-white">
                    Destaque suas habilidades e experiência
                  </span>
                </div>
              </div>

              <div className="space-x-4">
                <Link to="/signup">
                  <Button className="bg-conecta-green hover:bg-conecta-green-dark text-white px-8 py-6">
                    Cadastre-se como Profissional
                  </Button>
                </Link>
                <Link to="/forpros">
                  <Button
                    variant="outline"
                    className="border-white text-conecta-blue hover:bg-white hover:text-conecta-blue"
                  >
                    Saiba mais
                  </Button>
                </Link>
              </div>
            </div>

            <div className="hidden md:block relative">
              <div className="absolute inset-0 bg-black/10" />
              <img
                src="https://via.placeholder.com/800x600?text=Profissionais"
                alt="Profissionais ConectaMaré"
                className="h-full w-full object-cover"
              />
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default CtaSection;
