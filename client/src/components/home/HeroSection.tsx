import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";

const HeroSection = () => {
  return (
    <div className="relative bg-gradient-to-r from-conecta-blue to-conecta-blue-dark text-white">
      <div className="absolute inset-0 bg-black opacity-10" />
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative">
        <div className="pt-16 pb-20 md:pt-24 md:pb-28 lg:pt-32 lg:pb-36">
          <div className="text-center max-w-3xl mx-auto">
            <h1 className="text-4xl md:text-5xl lg:text-6xl font-bold mb-6">
              Conectamos você aos melhores profissionais do Complexo da Maré
            </h1>
            <p className="text-lg md:text-xl opacity-90 mb-10">
              Encontre profissionais qualificados e confiáveis para qualquer
              serviço que você precise, de forma rápida e segura.
            </p>
            <div className="flex flex-col sm:flex-row justify-center gap-4">
              <Link to="/professionals">
                <Button className="bg-conecta-green hover:bg-conecta-green-dark text-white px-8 py-6 text-lg">
                  Buscar Profissionais
                </Button>
              </Link>
              <Link to="/signup?role=professional">
                <Button
                  variant="outline"
                  className="border-white text-conecta-blue hover:bg-white px-8 py-6 text-lg"
                >
                  Oferecer Serviços
                </Button>
              </Link>
            </div>
          </div>
        </div>
      </div>

      <div className="absolute bottom-0 left-0 right-0">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 1440 100"
          className="w-full h-auto"
        >
          <path
            fill="#ffffff"
            fillOpacity="1"
            d="M0,64L60,58.7C120,53,240,43,360,42.7C480,43,600,53,720,58.7C840,64,960,64,1080,56C1200,48,1320,32,1380,24L1440,16L1440,100L1380,100C1320,100,1200,100,1080,100C960,100,840,100,720,100C600,100,480,100,360,100C240,100,120,100,60,100L0,100Z"
          />
        </svg>
      </div>
    </div>
  );
};

export default HeroSection;
