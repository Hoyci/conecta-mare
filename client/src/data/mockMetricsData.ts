
import { ClientFeedback, Notification, ProfileMetrics, ServiceMetric } from "@/types/metrics";

// Dados de amostra para visualizações de perfil
const generateViewsTrends = () => {
  const daily = Array.from({ length: 7 }, (_, i) => ({
    date: `2025-05-${9 + i}`,
    value: Math.floor(Math.random() * 40) + 10,
  }));

  const weekly = Array.from({ length: 4 }, (_, i) => ({
    date: `Semana ${i + 1}`,
    value: Math.floor(Math.random() * 200) + 50,
  }));

  const monthly = Array.from({ length: 6 }, (_, i) => ({
    date: `${i + 1}/2025`,
    value: Math.floor(Math.random() * 800) + 200,
  }));

  return { daily, weekly, monthly };
};

// Dados de amostra para taxa de conversão
const generateConversionTrends = () => {
  const daily = Array.from({ length: 7 }, (_, i) => ({
    date: `2025-05-${9 + i}`,
    value: Math.floor(Math.random() * 30) + 10,
  }));

  const weekly = Array.from({ length: 4 }, (_, i) => ({
    date: `Semana ${i + 1}`,
    value: Math.floor(Math.random() * 25) + 15,
  }));

  const monthly = Array.from({ length: 6 }, (_, i) => ({
    date: `${i + 1}/2025`,
    value: Math.floor(Math.random() * 20) + 20,
  }));

  return { daily, weekly, monthly };
};

// Avaliações de clientes
const recentFeedback: ClientFeedback[] = [
  {
    id: "1",
    clientName: "Maria Silva",
    clientAvatar: "https://i.pravatar.cc/150?img=1",
    rating: 5,
    comment: "Excelente profissional! Muito atencioso e pontual.",
    date: "2025-05-14",
  },
  {
    id: "2",
    clientName: "João Pereira",
    clientAvatar: "https://i.pravatar.cc/150?img=3",
    rating: 4,
    comment: "Bom trabalho, recomendo o serviço.",
    date: "2025-05-12",
  },
  {
    id: "3",
    clientName: "Ana Rodrigues",
    clientAvatar: "https://i.pravatar.cc/150?img=5",
    rating: 5,
    comment: "Profissional muito qualificado e prestativo.",
    date: "2025-05-10",
  },
];

// Serviços mais populares
const popularServices: ServiceMetric[] = [
  {
    id: "1",
    name: "Instalação Elétrica",
    views: 245,
    conversions: 48,
    trend: "up",
    trendPercentage: 12,
  },
  {
    id: "2",
    name: "Manutenção Preventiva",
    views: 180,
    conversions: 35,
    trend: "stable",
    trendPercentage: 2,
  },
  {
    id: "3",
    name: "Reparos Emergenciais",
    views: 220,
    conversions: 42,
    trend: "up",
    trendPercentage: 8,
  },
  {
    id: "4",
    name: "Consultoria",
    views: 120,
    conversions: 15,
    trend: "down",
    trendPercentage: 5,
  },
];

// Notificações
const notifications: Notification[] = [
  {
    id: "1",
    type: "milestone",
    title: "200+ visualizações!",
    description: "Seu perfil ultrapassou 200 visualizações este mês. Continue assim!",
    date: "2025-05-14",
    read: false,
  },
  {
    id: "2",
    type: "review",
    title: "Nova avaliação 5 estrelas",
    description: "Você recebeu uma nova avaliação 5 estrelas de Maria Silva.",
    date: "2025-05-13",
    read: true,
  },
  {
    id: "3",
    type: "tip",
    title: "Dica de otimização",
    description: "Adicione mais fotos de seus trabalhos para aumentar sua taxa de conversão.",
    date: "2025-05-12",
    read: false,
  },
];

export const mockProfileData: ProfileMetrics = {
  profileViews: {
    total: 1458,
    trends: generateViewsTrends(),
  },
  conversionRate: {
    rate: 22.5,
    trends: generateConversionTrends(),
  },
  ratings: {
    average: 4.7,
    total: 48,
    recent: recentFeedback,
  },
  services: popularServices,
  benchmarking: {
    profileViews: { user: 1458, average: 1200 },
    conversionRate: { user: 22.5, average: 18.2 },
    ratings: { user: 4.7, average: 4.2 },
  },
  notifications: notifications,
};

export const mockUserData = {
  name: "Carlos Oliveira",
  profession: "Eletricista Profissional",
  avatar: "https://i.pravatar.cc/300?img=8",
  description: "Eletricista com mais de 10 anos de experiência em instalações residenciais e comerciais. Especializado em sistemas elétricos eficientes e sustentáveis.",
  location: "Rio de Janeiro, RJ",
  coverPhoto: "https://images.unsplash.com/photo-1621905252507-b35492cc74b4?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2089&q=80",
};
