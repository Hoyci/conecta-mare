
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ChartContainer } from "@/components/ui/chart";
import { ProfileMetrics, TimeRange } from "@/types/metrics";
import { useState } from "react";
import { BarChart, Bar, LineChart, Line, XAxis, YAxis, ResponsiveContainer, Tooltip, TooltipProps } from "recharts";
import { ArrowUp, ArrowDown, TrendingUp, Eye, BarChart as BarChartIcon } from "lucide-react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

interface MetricsOverviewProps {
  metrics: ProfileMetrics;
}

// Custom tooltip component for the charts
const CustomTooltip = ({ active, payload, label }: TooltipProps<number, string>) => {
  if (active && payload && payload.length) {
    return (
      <div className="bg-white p-2 border border-gray-200 rounded shadow-sm text-xs">
        <p className="font-medium">{label}</p>
        <p className="text-conecta-blue">{`Valor: ${payload[0].value}`}</p>
      </div>
    );
  }
  return null;
};

export function MetricsOverview({ metrics }: MetricsOverviewProps) {
  const [viewsTimeRange, setViewsTimeRange] = useState<TimeRange>("weekly");
  const [conversionTimeRange, setConversionTimeRange] = useState<TimeRange>("weekly");

  const viewsData = metrics.profileViews.trends[viewsTimeRange];
  const conversionData = metrics.conversionRate.trends[conversionTimeRange];

  return (
    <div className="grid grid-cols-1 gap-6 lg:grid-cols-2">
      {/* Visualizações */}
      <Card>
        <CardHeader className="pb-2">
          <div className="flex items-center justify-between">
            <div>
              <CardTitle className="text-lg flex items-center gap-2">
                <Eye className="h-5 w-5 text-conecta-blue" />
                Visualizações do Perfil
              </CardTitle>
              <CardDescription>
                Total: {metrics.profileViews.total} visualizações
              </CardDescription>
            </div>
            <Tabs
              defaultValue="weekly"
              value={viewsTimeRange}
              onValueChange={(v) => setViewsTimeRange(v as TimeRange)}
              className="w-[230px]"
            >
              <TabsList className="grid w-full grid-cols-3">
                <TabsTrigger value="daily">Diário</TabsTrigger>
                <TabsTrigger value="weekly">Semanal</TabsTrigger>
                <TabsTrigger value="monthly">Mensal</TabsTrigger>
              </TabsList>
            </Tabs>
          </div>
        </CardHeader>
        <CardContent className="pt-4">
          <ChartContainer
            config={{
              views: {
                color: "#0056b3",
                label: "Visualizações",
              },
            }}
            className="h-[220px]"
          >
            <ResponsiveContainer width="100%" height="100%">
              <LineChart data={viewsData}>
                <XAxis
                  dataKey="date"
                  tick={{ fontSize: 12 }}
                  tickLine={false}
                  axisLine={false}
                />
                <YAxis
                  tick={{ fontSize: 12 }}
                  tickLine={false}
                  axisLine={false}
                  width={40}
                />
                <Tooltip content={<CustomTooltip />} />
                <Line
                  type="monotone"
                  dataKey="value"
                  stroke="var(--color-views)"
                  strokeWidth={2}
                  activeDot={{ r: 6 }}
                />
              </LineChart>
            </ResponsiveContainer>
          </ChartContainer>
        </CardContent>
      </Card>

      {/* Taxa de Conversão */}
      <Card>
        <CardHeader className="pb-2">
          <div className="flex items-center justify-between">
            <div>
              <CardTitle className="text-lg flex items-center gap-2">
                <TrendingUp className="h-5 w-5 text-conecta-green" />
                Taxa de Conversão
              </CardTitle>
              <CardDescription>
                Média: {metrics.conversionRate.rate}% de conversões
              </CardDescription>
            </div>
            <Tabs
              defaultValue="weekly"
              value={conversionTimeRange}
              onValueChange={(v) => setConversionTimeRange(v as TimeRange)}
              className="w-[230px]"
            >
              <TabsList className="grid w-full grid-cols-3">
                <TabsTrigger value="daily">Diário</TabsTrigger>
                <TabsTrigger value="weekly">Semanal</TabsTrigger>
                <TabsTrigger value="monthly">Mensal</TabsTrigger>
              </TabsList>
            </Tabs>
          </div>
        </CardHeader>
        <CardContent className="pt-4">
          <ChartContainer
            config={{
              conversion: {
                color: "#28a745",
                label: "Conversões",
              },
            }}
            className="h-[220px]"
          >
            <ResponsiveContainer width="100%" height="100%">
              <BarChart data={conversionData}>
                <XAxis
                  dataKey="date"
                  tick={{ fontSize: 12 }}
                  tickLine={false}
                  axisLine={false}
                />
                <YAxis
                  tick={{ fontSize: 12 }}
                  tickLine={false}
                  axisLine={false}
                  width={40}
                />
                <Tooltip content={<CustomTooltip />} />
                <Bar
                  dataKey="value"
                  fill="var(--color-conversion)"
                  radius={[4, 4, 0, 0]}
                />
              </BarChart>
            </ResponsiveContainer>
          </ChartContainer>
        </CardContent>
      </Card>
    </div>
  );
}
