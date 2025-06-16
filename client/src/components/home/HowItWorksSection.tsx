
import { SearchIcon, Calendar, CheckCircle, Star } from 'lucide-react';

const steps = [
  {
    icon: SearchIcon,
    title: "Busque profissionais",
    description: "Encontre profissionais qualificados baseados na sua localização e necessidades específicas."
  },
  {
    icon: CheckCircle,
    title: "Compare e escolha",
    description: "Veja perfis, avaliações e portfólios para escolher o profissional ideal para seu serviço."
  },
  {
    icon: Calendar,
    title: "Agende o serviço",
    description: "Entre em contato e agende diretamente com o profissional em um horário que funcione para você."
  },
  {
    icon: Star,
    title: "Avalie o serviço",
    description: "Após a conclusão, deixe sua avaliação para ajudar outros usuários da plataforma."
  }
];

const HowItWorksSection = () => {
  return (
    <section className="py-16 lg:py-24">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-16">
          <h2 className="text-3xl lg:text-4xl font-bold text-gray-900 mb-4">
            Como funciona
          </h2>
          <p className="text-lg text-gray-600 max-w-2xl mx-auto">
            Encontrar e contratar profissionais qualificados nunca foi tão simples
          </p>
        </div>
        
        <div className="relative">
          {/* Connecting line */}
          <div className="hidden lg:block absolute left-0 right-0 top-1/2 h-0.5 bg-gray-200 -translate-y-1/2 z-0" />
          
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            {steps.map((step, index) => (
              <div key={index} className="relative z-10 flex flex-col items-center text-center">
                <div className="bg-white border-2 border-conecta-blue text-conecta-blue rounded-full w-14 h-14 flex items-center justify-center mb-6">
                  <step.icon size={28} />
                </div>
                <div className="bg-white p-4 rounded-lg">
                  <h3 className="font-bold text-xl mb-3">{step.title}</h3>
                  <p className="text-gray-600">{step.description}</p>
                </div>
                {index < steps.length - 1 && (
                  <div className="hidden lg:block absolute top-1/2 right-0 transform translate-x-1/2 -translate-y-16">
                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path d="M5 12H19M19 12L12 5M19 12L12 19" stroke="#0056b3" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
                    </svg>
                  </div>
                )}
              </div>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
};

export default HowItWorksSection;
