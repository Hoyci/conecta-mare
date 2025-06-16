import { Shield, Clock, Award, Search } from "lucide-react";

const features = [
  {
    icon: Shield,
    title: "Profissionais Verificados",
    description:
      "Todos os profissionais passam por um processo de verificação de identidade e credenciais para garantir sua segurança.",
  },
  {
    icon: Clock,
    title: "Agende com Facilidade",
    description:
      "Marque serviços diretamente pelo aplicativo de acordo com a sua disponibilidade e a do profissional.",
  },
  {
    icon: Award,
    title: "Avaliações Reais",
    description:
      "Veja avaliações e comentários de outros clientes para escolher o profissional ideal para você.",
  },
  {
    icon: Search,
    title: "Busca Inteligente",
    description:
      "Encontre profissionais com base na sua localização, especialidade e disponibilidade em poucos cliques.",
  },
];

const FeaturesSection = () => {
  return (
    <section className="py-16 lg:py-24">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-16">
          <h2 className="text-3xl lg:text-4xl font-bold text-gray-900 mb-4">
            Por que escolher o ConectaMaré
          </h2>
          <p className="text-lg text-gray-600 max-w-2xl mx-auto">
            Nossa plataforma foi criada para facilitar a conexão entre clientes
            e profissionais, garantindo qualidade, segurança e praticidade.
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {features.map((feature, index) => (
            <div
              key={index}
              className="bg-white rounded-lg p-6 shadow-md hover:shadow-lg transition-shadow flex flex-col items-center text-center"
            >
              <div className="bg-conecta-blue/10 p-4 rounded-full mb-6">
                <feature.icon size={28} className="text-conecta-blue" />
              </div>
              <h3 className="font-bold text-xl mb-3">{feature.title}</h3>
              <p className="text-gray-600">{feature.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default FeaturesSection;
