import Navbar from "@/components/layout/Navbar";
import { useState } from "react";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import { Card, CardContent } from "@/components/ui/card";
import { Avatar } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import {
  LineChart,
  Line,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  BarChart,
  Bar,
  RadarChart,
  Radar,
  PolarGrid,
  PolarAngleAxis,
  PolarRadiusAxis,
  ResponsiveContainer,
  Legend,
} from "recharts";
import { AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Calendar, Edit, StarIcon } from "lucide-react";
import { useAuth } from "@/hooks/use-auth";
import { startCase, camelCase } from "lodash-es";
import DashboardSkeleton from "@/components/dashboard/DashboardSkeleton";
import { getSignedUser } from "@/services/user-service";
import { useQuery } from "@tanstack/react-query";

const mockProfile = {
  name: "Juliana Silva",
  profession: "Designer Gráfico",
  description: "Especializada em identidade visual e branding.",
};

const profileViews = [
  { date: "Seg", value: 20 },
  { date: "Ter", value: 40 },
  { date: "Qua", value: 80 },
  { date: "Qui", value: 55 },
  { date: "Sex", value: 90 },
];

const services = [
  { name: "Logo Design", views: 150, conversions: 30 },
  { name: "Cartão de Visita", views: 100, conversions: 15 },
  { name: "Manual da Marca", views: 70, conversions: 25 },
];

const benchmarks = [
  { metric: "Visualizações", você: 300, média: 250 },
  { metric: "Conversões", você: 75, média: 50 },
  { metric: "Retenção", você: 40, média: 35 },
];

const Dashboard = () => {
  const [tab, setTab] = useState("overview");


  const { data: userData = {} } = useQuery({
    queryKey: ['userData'],
    queryFn: getSignedUser
  })

  if (!userData.id) {
    return <DashboardSkeleton />;
  }


  return (

    <div className="min-h-screen flex flex-col bg-conecta-gray">
      <Navbar />
      <main className="flex-grow py-6">
        <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
          <Card className="border-none shadow-md ">
            <CardContent className="p-0">
              <div className="flex flex-col py-4 md:flex-row md:items-end px-6 relative">
                <div className="w-full flex flex-col gap-4">
                  <div className="flex flex-row">
                    <Avatar className="w-24 h-24 border-2 rounded-full shadow-lg">
                      {userData.profileImage ?
                        <AvatarImage src={userData.profileImage} />
                        :
                        <AvatarFallback className="bg-conecta-blue text-white text-2xl font-bold">
                          {userData.fullName.charAt(0)}
                        </AvatarFallback>
                      }
                    </Avatar>

                    <div className="mt-4 md:ml-6 md:mt-0 pb-4">
                      <div className="flex items-center">
                        <h1 className="text-2xl font-bold">
                          {startCase(camelCase(userData.fullName))}
                        </h1>
                        <span className="ml-3 bg-conecta-green-light text-conecta-green-dark text-xs px-2 py-1 rounded-full font-medium">
                          Verificado
                        </span>
                      </div>
                      <p className="text-conecta-blue font-semibold flex items-center">
                        {userData.subcategoryName}
                      </p>
                      <p className="text-gray-600 mt-1">
                        {userData.jobDescription}
                      </p>
                    </div>

                    <div className="md:ml-auto mt-4 md:mt-0 pb-4">
                      <Button className="bg-conecta-blue hover:bg-conecta-blue-dark text-white shadow-md flex items-center">
                        <Edit className="h-4 w-4 mr-2" /> Editar Perfil
                      </Button>
                    </div>
                  </div>

                  <div className="w-full flex gap-4 justify-between">
                    <div className="p-4 bg-conecta-blue-light text-conecta-blue rounded-lg w-full">
                      <p className="text-sm font-medium">Visualizações</p>
                      <p className="text-2xl text-conecta-blue-dark font-bold">
                        285
                      </p>
                      <p className="text-xs text-green-600">+12% esta semana</p>
                    </div>
                    <div className="p-4 bg-conecta-green-light text-conecta-green rounded-lg w-full">
                      <p className="text-sm font-medium">Projetos</p>
                      <p className="text-2xl text-conecta-green-dark font-bold">
                        27
                      </p>
                      <p className="text-xs text-gray-500">3 ativos</p>
                    </div>
                    <div className="p-4 bg-conecta-yellow-light text-conecta-yellow rounded-lg w-full">
                      <p className="text-sm font-medium">Avaliação</p>
                      <div className="flex gap-2 items-center">
                        <p className="text-2xl text-conecta-yellow-dark font-bold">
                          4.8
                        </p>
                        <StarIcon
                          fill="currentColor"
                          stroke="none"
                          className="text-conecta-yellow"
                        />
                      </div>
                      <p className="text-xs text-gray-500">52 avaliações</p>
                    </div>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>

          <Tabs value={tab} onValueChange={setTab} className="space-y-6">
            <TabsList className="bg-white p-1 shadow-sm rounded-lg border border-gray-200">
              <TabsTrigger
                value="overview"
                className="data-[state=active]:bg-conecta-blue data-[state=active]:text-white px-6 py-2 rounded-md transition-all"
              >
                Visão Geral
              </TabsTrigger>
              <TabsTrigger
                value="services"
                className="data-[state=active]:bg-conecta-blue data-[state=active]:text-white px-6 py-2 rounded-md transition-all"
              >
                Serviços
              </TabsTrigger>
              <TabsTrigger
                value="comparativo"
                className="data-[state=active]:bg-conecta-blue data-[state=active]:text-white px-6 py-2 rounded-md transition-all"
              >
                Comparativo
              </TabsTrigger>
            </TabsList>

            <TabsContent value="overview" className="space-y-6">
              <div className="grid md:grid-cols-2 gap-6">
                <Card className="border-none shadow-md">
                  <CardContent className="p-6">
                    <div className="flex justify-between items-center mb-4">
                      <h3 className="font-semibold text-gray-800 flex items-center">
                        <Calendar className="h-5 w-5 mr-2 text-conecta-blue" />
                        Visualizações do Perfil
                      </h3>
                      <select className="text-sm border-gray-200 rounded-md text-gray-500">
                        <option>Esta semana</option>
                        <option>Este mês</option>
                      </select>
                    </div>
                    <ResponsiveContainer width="100%" height={220}>
                      <LineChart data={profileViews}>
                        <defs>
                          <linearGradient
                            id="colorViews"
                            x1="0"
                            y1="0"
                            x2="0"
                            y2="1"
                          >
                            <stop
                              offset="5%"
                              stopColor="#0070f3"
                              stopOpacity={0.3}
                            />
                            <stop
                              offset="95%"
                              stopColor="#0070f3"
                              stopOpacity={0}
                            />
                          </linearGradient>
                        </defs>
                        <Line
                          type="monotone"
                          dataKey="value"
                          stroke="#0070f3"
                          strokeWidth={3}
                          dot={{
                            stroke: "#0070f3",
                            strokeWidth: 2,
                            r: 4,
                            fill: "white",
                          }}
                          activeDot={{
                            r: 6,
                            stroke: "#0070f3",
                            strokeWidth: 2,
                            fill: "white",
                          }}
                        />
                        <CartesianGrid stroke="#eee" vertical={false} />
                        <XAxis
                          dataKey="date"
                          axisLine={false}
                          tickLine={false}
                        />
                        <YAxis axisLine={false} tickLine={false} />
                        <Tooltip
                          contentStyle={{
                            borderRadius: "8px",
                            border: "none",
                            boxShadow: "0 4px 12px rgba(0,0,0,0.1)",
                          }}
                        />
                      </LineChart>
                    </ResponsiveContainer>
                    <div className="mt-2 flex items-center justify-between text-sm">
                      <div className="text-gray-500">
                        Total de visualizações:{" "}
                        <span className="font-bold text-conecta-blue">285</span>
                      </div>
                      <div className="flex items-center text-green-600 font-medium">
                        +28%{" "}
                        <span className="text-gray-500 ml-1 font-normal">
                          vs período anterior
                        </span>
                      </div>
                    </div>
                  </CardContent>
                </Card>

                <div className="grid gap-6">
                  <Card className="border-none shadow-md">
                    <CardContent className="p-6">
                      <div className="flex justify-between">
                        <div>
                          <h3 className="font-semibold text-gray-800 mb-1">
                            Taxa de Conversão
                          </h3>
                          <p className="text-gray-500 text-sm">
                            Projetos confirmados / visualizações
                          </p>
                        </div>
                        <div className="text-right">
                          <div className="text-4xl font-bold text-conecta-green">
                            18.5%
                          </div>
                          <p className="text-green-600 text-sm">
                            +2.7% vs média
                          </p>
                        </div>
                      </div>
                      <div className="mt-4 bg-gray-100 rounded-full h-2.5 w-full">
                        <div
                          className="bg-conecta-green h-2.5 rounded-full"
                          style={{ width: "65%" }}
                        ></div>
                      </div>
                    </CardContent>
                  </Card>

                  <Card className="border-none shadow-md">
                    <CardContent className="p-6">
                      <div className="flex justify-between items-center mb-2">
                        <h3 className="font-semibold text-gray-800">
                          Feedback dos Clientes
                        </h3>
                        <div className="text-lg bg-yellow-50 px-2 py-1 rounded-md font-semibold text-yellow-600">
                          ⭐ 4.8
                        </div>
                      </div>
                      <p className="text-gray-500 text-sm mb-3">
                        Baseado em 52 avaliações
                      </p>
                      <div className="space-y-2">
                        <div className="bg-gray-50 border border-gray-100 rounded-lg p-3 text-sm italic text-gray-600">
                          "Excelente trabalho e comunicação durante todo o
                          projeto de identidade visual."
                        </div>
                        <div className="bg-gray-50 border border-gray-100 rounded-lg p-3 text-sm italic text-gray-600">
                          "Voltaria a contratar com certeza. Profissional
                          extremamente talentosa!"
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                </div>
              </div>
            </TabsContent>

            <TabsContent value="services" className="space-y-6">
              <Card className="border-none shadow-md">
                <CardContent className="p-6">
                  <div className="flex justify-between items-center mb-6">
                    <h3 className="font-semibold text-gray-800">
                      Serviços Mais Populares
                    </h3>
                    <Button
                      variant="outline"
                      className="text-conecta-blue border-conecta-blue hover:bg-blue-50"
                    >
                      Adicionar Serviço
                    </Button>
                  </div>
                  <ResponsiveContainer width="100%" height={320}>
                    <BarChart
                      data={services}
                      margin={{ top: 10, right: 30, left: 0, bottom: 20 }}
                    >
                      <CartesianGrid strokeDasharray="3 3" vertical={false} />
                      <XAxis
                        dataKey="name"
                        axisLine={false}
                        tickLine={false}
                        padding={{ left: 30, right: 30 }}
                      />
                      <YAxis axisLine={false} tickLine={false} />
                      <Tooltip
                        contentStyle={{
                          borderRadius: "8px",
                          border: "none",
                          boxShadow: "0 4px 12px rgba(0,0,0,0.1)",
                        }}
                      />
                      <Legend verticalAlign="top" height={36} />
                      <Bar
                        dataKey="views"
                        fill="#0070f3"
                        name="Visualizações"
                        radius={[4, 4, 0, 0]}
                        barSize={35}
                      />
                      <Bar
                        dataKey="conversions"
                        fill="#00c853"
                        name="Conversões"
                        radius={[4, 4, 0, 0]}
                        barSize={35}
                      />
                    </BarChart>
                  </ResponsiveContainer>

                  <div className="mt-6 grid grid-cols-3 gap-4">
                    <Card className="bg-blue-50 border-none">
                      <CardContent className="p-4">
                        <h4 className="font-medium mb-2">Logo Design</h4>
                        <div className="flex items-center justify-between">
                          <p className="text-3xl font-bold text-conecta-blue">
                            150
                          </p>
                          <p className="bg-blue-100 text-conecta-blue text-sm rounded-full px-2 py-1">
                            + 12%
                          </p>
                        </div>
                        <p className="text-gray-500 text-xs mt-1">
                          Visualizações este mês
                        </p>
                      </CardContent>
                    </Card>
                    <Card className="bg-green-50 border-none">
                      <CardContent className="p-4">
                        <h4 className="font-medium mb-2">Cartão de Visita</h4>
                        <div className="flex items-center justify-between">
                          <p className="text-3xl font-bold text-conecta-green">
                            100
                          </p>
                          <p className="bg-green-100 text-conecta-green text-sm rounded-full px-2 py-1">
                            + 8%
                          </p>
                        </div>
                        <p className="text-gray-500 text-xs mt-1">
                          Visualizações este mês
                        </p>
                      </CardContent>
                    </Card>
                    <Card className="bg-purple-50 border-none">
                      <CardContent className="p-4">
                        <h4 className="font-medium mb-2">Manual da Marca</h4>
                        <div className="flex items-center justify-between">
                          <p className="text-3xl font-bold text-purple-600">
                            70
                          </p>
                          <p className="bg-purple-100 text-purple-600 text-sm rounded-full px-2 py-1">
                            + 5%
                          </p>
                        </div>
                        <p className="text-gray-500 text-xs mt-1">
                          Visualizações este mês
                        </p>
                      </CardContent>
                    </Card>
                  </div>
                </CardContent>
              </Card>
            </TabsContent>

            {/* Comparativo  */
            }
            <TabsContent value="comparativo" className="space-y-6">
              <Card className="border-none shadow-md">
                <CardContent className="p-6">
                  <div className="flex justify-between items-center mb-6">
                    <h3 className="font-semibold text-gray-800">
                      Benchmark com outros profissionais
                    </h3>
                    <select className="text-sm border-gray-200 rounded-md text-gray-500">
                      <option>Últimos 30 dias</option>
                      <option>Este trimestre</option>
                    </select>
                  </div>
                  <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 items-center">
                    <ResponsiveContainer width="100%" height={350}>
                      <RadarChart outerRadius={110} data={benchmarks}>
                        <PolarGrid stroke="#e2e8f0" />
                        <PolarAngleAxis
                          dataKey="metric"
                          tick={{ fill: "#64748b", fontSize: 12 }}
                        />
                        <PolarRadiusAxis
                          angle={30}
                          domain={[0, 400]}
                          tick={false}
                          axisLine={false}
                        />
                        <Radar
                          name="Você"
                          dataKey="você"
                          stroke="#0070f3"
                          fill="#0070f3"
                          fillOpacity={0.6}
                        />
                        <Radar
                          name="Média do setor"
                          dataKey="média"
                          stroke="#8884d8"
                          fill="#8884d8"
                          fillOpacity={0.3}
                        />
                        <Tooltip
                          contentStyle={{
                            borderRadius: "8px",
                            border: "none",
                            boxShadow: "0 4px 12px rgba(0,0,0,0.1)",
                          }}
                        />
                        <Legend />
                      </RadarChart>
                    </ResponsiveContainer>

                    <div className="space-y-4">
                      <div className="bg-blue-50 p-4 rounded-xl">
                        <div className="flex justify-between">
                          <div>
                            <p className="text-sm font-medium text-gray-500">
                              Visualizações
                            </p>
                            <h4 className="text-2xl font-bold text-conecta-blue">
                              300
                            </h4>
                          </div>
                          <div className="bg-blue-100 h-fit text-conecta-blue px-2 py-1 rounded-md text-sm font-medium">
                            +20% acima da média
                          </div>
                        </div>
                      </div>

                      <div className="bg-green-50 p-4 rounded-xl">
                        <div className="flex justify-between">
                          <div>
                            <p className="text-sm font-medium text-gray-500">
                              Conversões
                            </p>
                            <h4 className="text-2xl font-bold text-conecta-green">
                              75
                            </h4>
                          </div>
                          <div className="bg-green-100 h-fit text-conecta-green px-2 py-1 rounded-md text-sm font-medium">
                            +50% acima da média
                          </div>
                        </div>
                      </div>

                      <div className="bg-purple-50 p-4 rounded-xl">
                        <div className="flex justify-between">
                          <div>
                            <p className="text-sm font-medium text-gray-500">
                              Retenção
                            </p>
                            <h4 className="text-2xl font-bold text-purple-600">
                              40
                            </h4>
                          </div>
                          <div className="bg-purple-100 h-fit text-purple-600 px-2 py-1 rounded-md text-sm font-medium">
                            +14% acima da média
                          </div>
                        </div>
                      </div>

                      <p className="text-sm text-gray-500 italic">
                        Você está se destacando em todas as métricas principais
                        quando comparado com outros profissionais da sua área.
                        Continue mantendo a qualidade dos seus serviços!
                      </p>
                    </div>
                  </div>
                </CardContent>
              </Card>
            </TabsContent>
          </Tabs>
        </div>
      </main>
    </div>
  );
};

export default Dashboard;
