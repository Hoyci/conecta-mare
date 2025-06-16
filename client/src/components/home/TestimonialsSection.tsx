import { useState } from "react";
import { Star } from "lucide-react";

const testimonials = [
  {
    name: "Carlos Silva",
    role: "Cliente",
    image: "https://via.placeholder.com/100x100?text=C",
    rating: 5,
    content:
      "Encontrei um eletricista excelente através do ConectaPro. Resolveu meu problema rapidamente e com um preço justo. Excelente plataforma para quem precisa de serviços de qualidade!",
  },
  {
    name: "Ana Martins",
    role: "Cliente",
    image: "https://via.placeholder.com/100x100?text=A",
    rating: 5,
    content:
      "Como mãe solteira, sempre tive receio de contratar serviços para reparos em casa. Com o ConectaPro, me senti segura ao encontrar profissionais verificados e bem avaliados. Super recomendo!",
  },
  {
    name: "Roberto Almeida",
    role: "Profissional (Encanador)",
    image: "https://via.placeholder.com/100x100?text=R",
    rating: 5,
    content:
      "Desde que me cadastrei no ConectaPro, minha agenda está sempre cheia. A plataforma me trouxe clientes sérios e me ajudou a expandir meu negócio. Valeu muito a pena!",
  },
  {
    name: "Juliana Costa",
    role: "Cliente",
    image: "https://via.placeholder.com/100x100?text=J",
    rating: 4,
    content:
      "Precisei urgentemente de um professor particular para meu filho e em menos de uma hora já estava em contato com vários profissionais qualificados. O agendamento foi super prático!",
  },
  {
    name: "Felipe Gomes",
    role: "Profissional (Designer)",
    image: "https://via.placeholder.com/100x100?text=F",
    rating: 5,
    content:
      "Como freelancer, o ConectaPro se tornou minha principal fonte de clientes. A interface é intuitiva tanto para profissionais quanto para clientes. Nota 10!",
  },
];

const TestimonialsSection = () => {
  const [activeIndex, setActiveIndex] = useState(0);

  const showPrevious = () => {
    setActiveIndex((prevIndex) =>
      prevIndex === 0 ? testimonials.length - 1 : prevIndex - 1,
    );
  };

  const showNext = () => {
    setActiveIndex((prevIndex) =>
      prevIndex === testimonials.length - 1 ? 0 : prevIndex + 1,
    );
  };

  return (
    <section className="py-16 lg:py-24 bg-conecta-blue text-white">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-12">
          <h2 className="text-3xl lg:text-4xl font-bold mb-4">
            O que nossos usuários dizem
          </h2>
          <p className="text-lg opacity-90 max-w-2xl mx-auto">
            Histórias de sucesso de clientes e profissionais que utilizam o
            ConectaMaré
          </p>
        </div>

        <div className="relative max-w-4xl mx-auto">
          {/* Carousel controls */}
          <button
            onClick={showPrevious}
            className="absolute left-0 top-1/2 -translate-y-1/2 -translate-x-4 lg:-translate-x-10 bg-white/10 hover:bg-white/20 text-white rounded-full p-2 z-10"
            aria-label="Depoimento anterior"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={2}
              stroke="currentColor"
              className="w-6 h-6"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M15.75 19.5L8.25 12l7.5-7.5"
              />
            </svg>
          </button>

          <div className="overflow-hidden">
            <div
              className="transition-transform duration-500 ease-in-out flex"
              style={{ transform: `translateX(-${activeIndex * 100}%)` }}
            >
              {testimonials.map((testimonial, idx) => (
                <div key={idx} className="min-w-full px-4">
                  <div className="bg-white/10 backdrop-blur-sm rounded-xl p-6 md:p-8 text-center">
                    <div className="flex justify-center mb-4">
                      <img
                        src={testimonial.image}
                        alt={testimonial.name}
                        className="w-20 h-20 rounded-full object-cover"
                      />
                    </div>
                    <div className="flex justify-center mb-4">
                      {[...Array(5)].map((_, i) => (
                        <Star
                          key={i}
                          size={20}
                          className={
                            i < testimonial.rating
                              ? "fill-yellow-400 text-yellow-400"
                              : "text-gray-400"
                          }
                        />
                      ))}
                    </div>
                    <blockquote className="text-lg md:text-xl italic mb-6">
                      "{testimonial.content}"
                    </blockquote>
                    <div>
                      <div className="font-semibold text-lg">
                        {testimonial.name}
                      </div>
                      <div className="opacity-80">{testimonial.role}</div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>

          <button
            onClick={showNext}
            className="absolute right-0 top-1/2 -translate-y-1/2 translate-x-4 lg:translate-x-10 bg-white/10 hover:bg-white/20 text-white rounded-full p-2 z-10"
            aria-label="Próximo depoimento"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={2}
              stroke="currentColor"
              className="w-6 h-6"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M8.25 4.5l7.5 7.5-7.5 7.5"
              />
            </svg>
          </button>
        </div>

        {/* Indicators */}
        <div className="flex justify-center space-x-2 mt-8">
          {testimonials.map((_, idx) => (
            <button
              key={idx}
              onClick={() => setActiveIndex(idx)}
              className={`w-3 h-3 rounded-full transition-colors ${
                activeIndex === idx ? "bg-white" : "bg-white/30"
              }`}
              aria-label={`Ver depoimento ${idx + 1}`}
            />
          ))}
        </div>
      </div>
    </section>
  );
};

export default TestimonialsSection;
