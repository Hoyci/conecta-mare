import { Professional } from "@/types/professional";

export const professionals: Professional[] = [
  {
    id: "1",
    name: "Ricardo Oliveira",
    profession: "Eletricista",
    profileImage: "https://picsum.photos/300/300",
    rating: 4.9,
    reviewsCount: 47,
    location: "São Paulo, SP",
    responseTimes: "2 horas em média",
    isVerified: true,
    isPremium: true,
    description:
      "Eletricista com mais de 15 anos de experiência em instalações residenciais e comerciais. Especialista em manutenção preventiva e corretiva de redes elétricas.",
    specialties: ["Instalações", "Reparos", "Automação", "Iluminação"],
    experience:
      "15 anos de atuação em projetos residenciais e comerciais. Formado em Técnico em Elétrica pelo SENAI com diversas especializações em eficiência energética.",
    availability: "Segunda a Sábado, das 8h às 18h",
    serviceAreas: "São Paulo (Zona Sul, Zona Oeste e Centro)",
    credentials: [
      "Técnico em Elétrica - SENAI",
      "Certificado NR10 - Segurança em Instalações Elétricas",
      "Especialização em Eficiência Energética",
    ],
    portfolioImages: [
      {
        url: "https://via.placeholder.com/500x300?text=Projeto1",
        title: "Instalação elétrica em condomínio",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Projeto2",
        title: "Automação residencial",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Projeto3",
        title: "Iluminação de destaque",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Projeto4",
        title: "Instalação de painéis solares",
      },
    ],
    awards: [
      {
        title: "Técnico do Ano 2022",
        description: "Reconhecimento pelo sindicato dos eletricistas",
      },
      {
        title: "Certificação Ouro",
        description: "Melhor avaliação em atendimento ao cliente",
      },
    ],
    services: [
      {
        name: "Visita técnica",
        price: "Gratuito",
        description: "Avaliação inicial e orçamento sem compromisso",
      },
      {
        name: "Instalação elétrica completa",
        price: 1500,
        description:
          "Instalação de toda rede elétrica para residências até 150m²",
      },
      {
        name: "Troca de quadro de energia",
        price: 600,
        description: "Substituição e atualização de quadros antigos",
      },
      {
        name: "Reparo emergencial",
        price: 250,
        description: "Atendimento prioritário para problemas urgentes",
      },
    ],
    reviews: [
      {
        user: "Marcela Silva",
        rating: 5,
        comment:
          "Atendimento rápido e eficiente. Resolveu o problema da queda de energia imediatamente.",
        date: "15/04/2023",
      },
      {
        user: "Paulo Mendes",
        rating: 5,
        comment:
          "Excelente profissional! Fez a instalação elétrica da minha casa com muita qualidade e no prazo combinado.",
        date: "03/03/2023",
      },
      {
        user: "Ana Ribeiro",
        rating: 4,
        comment:
          "Trabalho bem feito, apenas um pequeno atraso no dia agendado, mas comunicou com antecedência.",
        date: "12/02/2023",
      },
    ],
  },
  {
    id: "2",
    name: "Juliana Martins",
    profession: "Designer de Interiores",
    profileImage: "https://picsum.photos/300/300",
    rating: 4.8,
    reviewsCount: 36,
    location: "Rio de Janeiro, RJ",
    responseTimes: "Mesmo dia",
    isVerified: true,
    isPremium: false,
    description:
      "Designer de interiores especializada em transformar espaços residenciais com foco em sustentabilidade e aproveitamento máximo de espaços.",
    specialties: [
      "Design Residencial",
      "Ambientes Compactos",
      "Design Sustentável",
      "Consultoria",
    ],
    experience:
      "Graduada em Design de Interiores pela UFRJ com 8 anos de experiência em projetos residenciais e comerciais. Especialização em Design Sustentável.",
    availability: "Segunda a Sexta, das 9h às 18h",
    serviceAreas: "Rio de Janeiro (Zona Sul, Zona Oeste e Zona Norte)",
    credentials: [
      "Bacharel em Design de Interiores - UFRJ",
      "Especialização em Design Sustentável - PUC-Rio",
      "Membro da Associação Brasileira de Designers de Interiores",
    ],
    portfolioImages: [
      {
        url: "https://via.placeholder.com/500x300?text=Interior1",
        title: "Apartamento compacto em Ipanema",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Interior2",
        title: "Reforma de cozinha sustentável",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Interior3",
        title: "Home office integrado",
      },
    ],
    awards: [
      {
        title: "Prêmio Casa Cor 2021",
        description: "Melhor ambiente sustentável",
      },
      { title: "Destaque em Design", description: "Revista Casa & Jardim" },
    ],
    services: [
      {
        name: "Consultoria inicial",
        price: 350,
        description: "Análise do ambiente e sugestões iniciais (2 horas)",
      },
      {
        name: "Projeto completo",
        price: "A partir de R$ 5.000",
        description:
          "Projeto detalhado com plantas, renderizações e lista de compras",
      },
      {
        name: "Acompanhamento de obra",
        price: 250,
        description: "Visita técnica para acompanhamento (por hora)",
      },
    ],
    reviews: [
      {
        user: "Roberto Almeida",
        rating: 5,
        comment:
          "A Juliana transformou meu apartamento de 45m² em um espaço incrível e funcional. Superou todas as expectativas!",
        date: "20/05/2023",
      },
      {
        user: "Fernanda Costa",
        rating: 5,
        comment:
          "Profissional extremamente talentosa e atenciosa. O projeto ficou perfeito e dentro do orçamento.",
        date: "11/04/2023",
      },
      {
        user: "Carlos Eduardo",
        rating: 4,
        comment: "Ótimas ideias e soluções criativas. Recomendo fortemente.",
        date: "28/03/2023",
      },
    ],
  },
  {
    id: "3",
    name: "André Santos",
    profession: "Encanador",
    profileImage: "https://picsum.photos/300/300",
    rating: 4.7,
    reviewsCount: 52,
    location: "Belo Horizonte, MG",
    responseTimes: "1 hora em média",
    isVerified: true,
    isPremium: true,
    description:
      "Especialista em reparos hidráulicos com atendimento rápido e eficiente. Soluciono vazamentos, entupimentos e instalação de equipamentos com garantia de serviço.",
    specialties: [
      "Reparos Emergenciais",
      "Detecção de Vazamentos",
      "Instalação de Equipamentos",
      "Manutenção Preventiva",
    ],
    experience:
      "Mais de 12 anos trabalhando com sistemas hidráulicos residenciais e comerciais. Técnico certificado em instalações hidráulicas.",
    availability: "Todos os dias, das 7h às 22h - Atendimento emergencial 24h",
    serviceAreas: "Toda região metropolitana de Belo Horizonte",
    credentials: [
      "Técnico em Instalações Hidráulicas - SENAI",
      "Certificação em Reparos Emergenciais",
      "Especialista em Economia de Água",
    ],
    portfolioImages: [
      {
        url: "https://via.placeholder.com/500x300?text=Hidraulica1",
        title: "Instalação de sistema de reuso de água",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Hidraulica2",
        title: "Reparo em tubulação complexa",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Hidraulica3",
        title: "Instalação de banheira com hidromassagem",
      },
    ],
    awards: [
      {
        title: "Selo Qualidade em Serviços",
        description: "Reconhecimento da Associação Comercial de MG",
      },
    ],
    services: [
      {
        name: "Visita técnica",
        price: 100,
        description:
          "Diagnóstico do problema (valor abatido em caso de contratação)",
      },
      {
        name: "Reparo de vazamentos",
        price: "A partir de R$ 180",
        description: "Identificação e correção de vazamentos simples",
      },
      {
        name: "Desentupimento",
        price: 220,
        description: "Desobstrução de pias, ralos, vasos sanitários",
      },
      {
        name: "Instalação de equipamentos",
        price: "A partir de R$ 250",
        description: "Instalação de torneiras, chuveiros, duchas, etc",
      },
    ],
    reviews: [
      {
        user: "Mariana Duarte",
        rating: 5,
        comment:
          "Atendimento rapidíssimo! Resolveu um vazamento complicado em menos de uma hora.",
        date: "10/06/2023",
      },
      {
        user: "João Victor",
        rating: 5,
        comment:
          "Excelente trabalho na instalação da minha banheira. Acabamento perfeito e limpeza total após o serviço.",
        date: "22/05/2023",
      },
      {
        user: "Luísa Mendes",
        rating: 4,
        comment:
          "Muito bom profissional, preço justo e trabalho bem feito. Recomendo.",
        date: "15/04/2023",
      },
    ],
  },
  {
    id: "4",
    name: "Patrícia Lima",
    profession: "Professora de Matemática",
    profileImage: "https://picsum.photos/300/300",
    rating: 4.9,
    reviewsCount: 64,
    location: "Curitiba, PR",
    responseTimes: "Mesmo dia",
    isVerified: true,
    isPremium: false,
    description:
      "Professora de Matemática com mais de 15 anos de experiência, especializada em ensino fundamental, médio e preparatório para vestibular/ENEM. Método pedagógico próprio focado em resultados rápidos.",
    specialties: [
      "Ensino Fundamental",
      "Ensino Médio",
      "Preparatório ENEM",
      "Cálculo para Universitários",
    ],
    experience:
      "Licenciatura em Matemática pela UFPR, Mestrado em Educação Matemática. 15 anos lecionando em escolas particulares e como professora particular.",
    availability: "Segunda a Sexta das 14h às 20h, Sábados das 9h às 14h",
    serviceAreas:
      "Atendimento online para todo Brasil. Presencial em Curitiba.",
    credentials: [
      "Licenciatura em Matemática - UFPR",
      "Mestrado em Educação Matemática - UFPR",
      "Especialização em Metodologias Ativas de Ensino",
    ],
    portfolioImages: [
      {
        url: "https://via.placeholder.com/500x300?text=Material1",
        title: "Material didático exclusivo",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Material2",
        title: "Jogos matemáticos para ensino fundamental",
      },
    ],
    awards: [
      {
        title: "Prêmio Educador Destaque 2020",
        description: "Secretaria de Educação de Curitiba",
      },
    ],
    services: [
      {
        name: "Aula particular individual",
        price: 120,
        description:
          "Aula personalizada de 1 hora com foco nas dificuldades do aluno",
      },
      {
        name: "Pacote mensal (8 aulas)",
        price: 800,
        description: "2 aulas semanais de 1 hora cada, material incluso",
      },
      {
        name: "Preparatório intensivo ENEM",
        price: 1500,
        description:
          "Pacote de 15 aulas focadas em resolver questões anteriores do ENEM",
      },
    ],
    reviews: [
      {
        user: "Renata Gomes",
        rating: 5,
        comment:
          "A professora Patrícia conseguiu fazer meu filho gostar de Matemática! As notas melhoraram muito após 2 meses de aula.",
        date: "12/06/2023",
      },
      {
        user: "Lucas Ferreira",
        rating: 5,
        comment:
          "Didática excelente! Consegui entender conceitos que nunca tinha compreendido na escola.",
        date: "05/05/2023",
      },
      {
        user: "Marisa Silveira",
        rating: 5,
        comment:
          "Graças às aulas da Patrícia, minha filha conseguiu nota máxima em matemática no ENEM. Método sensacional!",
        date: "10/03/2023",
      },
    ],
  },
  {
    id: "5",
    name: "Rafael Costa",
    profession: "Designer Gráfico",
    profileImage: "https://picsum.photos/300/300",
    rating: 4.7,
    reviewsCount: 28,
    location: "São Paulo, SP",
    responseTimes: "24 horas",
    isVerified: true,
    isPremium: true,
    description:
      "Designer gráfico especializado em identidade visual para pequenas e médias empresas. Desenvolvimento de logotipos, papelaria, materiais promocionais e presença digital.",
    specialties: [
      "Identidade Visual",
      "Logotipos",
      "Material Impresso",
      "Social Media",
    ],
    experience:
      "8 anos atuando como designer freelancer para agências e empresas de diversos segmentos. Formado em Design Gráfico pela FAAP.",
    availability: "Segunda a Sexta, horário comercial",
    serviceAreas: "Atendimento online para todo Brasil",
    credentials: [
      "Bacharel em Design Gráfico - FAAP",
      "Especialização em Design Digital - Belas Artes",
      "Certificado Adobe Certified Professional",
    ],
    portfolioImages: [
      {
        url: "https://via.placeholder.com/500x300?text=Logo1",
        title: "Identidade visual para restaurante",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Branding1",
        title: "Branding completo para startup",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Social1",
        title: "Kit de posts para redes sociais",
      },
    ],
    awards: [
      {
        title: "Prêmio ABEDESIGN 2021",
        description: "Categoria Identidade Visual",
      },
    ],
    services: [
      {
        name: "Consultoria de identidade visual",
        price: 300,
        description: "Análise da marca e recomendações estratégicas",
      },
      {
        name: "Criação de logotipo",
        price: 1200,
        description: "3 propostas com ajustes ilimitados na escolhida",
      },
      {
        name: "Identidade visual completa",
        price: 3500,
        description: "Logotipo, papelaria, manual da marca e aplicações",
      },
      {
        name: "Kit mensal para redes sociais",
        price: 800,
        description: "30 posts para Instagram, Facebook e LinkedIn",
      },
    ],
    reviews: [
      {
        user: "Marcos Pereira",
        rating: 5,
        comment:
          "Trabalho excepcional! O Rafael captou perfeitamente a essência da minha empresa no logotipo.",
        date: "15/06/2023",
      },
      {
        user: "Camila Rodrigues",
        rating: 4,
        comment:
          "Entrega dentro do prazo e ótima comunicação durante todo o processo. Recomendo!",
        date: "28/04/2023",
      },
      {
        user: "Empresa Sabor & Arte",
        rating: 5,
        comment:
          "Nossa identidade visual ficou incrível, muito além das expectativas. Valeu cada centavo!",
        date: "10/03/2023",
      },
    ],
  },
  {
    id: "6",
    name: "Amanda Souza",
    profession: "Fotógrafa",
    profileImage: "https://picsum.photos/300/300",
    rating: 4.8,
    reviewsCount: 42,
    location: "Porto Alegre, RS",
    responseTimes: "48 horas",
    isVerified: true,
    isPremium: false,
    description:
      "Fotógrafa especializada em retratos, eventos sociais e ensaios. Capturo momentos únicos com um olhar artístico que valoriza a naturalidade e emoção de cada ocasião.",
    specialties: [
      "Casamentos",
      "Ensaios",
      "Eventos Corporativos",
      "Fotografia de Produto",
    ],
    experience:
      "10 anos de experiência em fotografia profissional. Formada em Fotografia pelo Instituto Europeu de Design com especialização em fotografia digital.",
    availability: "Disponível para agendamento com 2 semanas de antecedência",
    serviceAreas: "Porto Alegre e região metropolitana",
    credentials: [
      "Graduação em Fotografia - IED",
      "Curso Avançado de Iluminação Fotográfica - SENAC",
      "Especialização em Fotografia de Casamentos - Wedding Brasil",
    ],
    portfolioImages: [
      {
        url: "https://via.placeholder.com/500x300?text=Wedding1",
        title: "Casamento na praia",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Portrait1",
        title: "Ensaio feminino",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Event1",
        title: "Evento corporativo",
      },
      {
        url: "https://via.placeholder.com/500x300?text=Product1",
        title: "Fotografia de produto",
      },
    ],
    awards: [
      {
        title: "Fotógrafa do Ano 2022",
        description: "Associação Gaúcha de Fotografia",
      },
      {
        title: "Top 10 Wedding Photographers",
        description: "Portal Casar é Fácil",
      },
    ],
    services: [
      {
        name: "Ensaio pessoal (1 hora)",
        price: 500,
        description: "Sessão de 1 hora com 30 fotos editadas",
      },
      {
        name: "Cobertura de evento (4 horas)",
        price: 1800,
        description: "Cobertura completa com entrega de 150 fotos editadas",
      },
      {
        name: "Casamento - pacote básico",
        price: 3500,
        description: "Cerimônia e recepção, 300 fotos editadas, álbum digital",
      },
      {
        name: "Fotografia de produto (10 itens)",
        price: 1200,
        description: "Fotos profissionais para e-commerce ou catálogo",
      },
    ],
    reviews: [
      {
        user: "Caroline e Pedro",
        rating: 5,
        comment:
          "A Amanda registrou nosso casamento de forma incrível! As fotos capturam perfeitamente cada emoção daquele dia especial.",
        date: "10/06/2023",
      },
      {
        user: "Empresa Tech Solutions",
        rating: 5,
        comment:
          "Contratamos para um evento corporativo e o resultado foi excelente. Profissionalismo impecável!",
        date: "22/05/2023",
      },
      {
        user: "Beatriz Lopes",
        rating: 4,
        comment:
          "Meu ensaio ficou lindo! Ótimo atendimento, apenas um pequeno atraso na entrega final das fotos.",
        date: "15/04/2023",
      },
    ],
  },
];

export const categories = [
  "Eletricistas",
  "Encanadores",
  "Professores",
  "Designers",
  "Fotógrafos",
  "Programadores",
  "Jardineiros",
  "Pedreiros",
  "Pintores",
  "Marceneiros",
  "Arquitetos",
  "Nutricionistas",
  "Personal Trainers",
  "Advogados",
  "Contadores",
];
