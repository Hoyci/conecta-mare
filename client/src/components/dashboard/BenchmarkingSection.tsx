
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ProfileMetrics } from "@/types/metrics";

interface BenchmarkingSectionProps {
  benchmarking: ProfileMetrics["benchmarking"];
}

export function BenchmarkingSection({ benchmarking }: BenchmarkingSectionProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-lg">Comparativo de Desempenho</CardTitle>
        <CardDescription>
          Seus resultados em comparação com outros profissionais da mesma categoria
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="space-y-6">
          {/* Visualizações do Perfil */}
          <div className="space-y-2">
            <div className="flex items-center justify-between">
              <p className="text-sm font-medium">Visualizações do Perfil</p>
              <div className="flex items-center">
                <p className="font-medium">{benchmarking.profileViews.user}</p>
                <span className="mx-2 text-muted-foreground">vs</span>
                <p className="text-muted-foreground">
                  {benchmarking.profileViews.average} (média)
                </p>
              </div>
            </div>
            <div className="h-2 w-full rounded-full bg-gray-100">
              <div
                className="h-2 rounded-full bg-conecta-blue"
                style={{
                  width: `${(benchmarking.profileViews.user / (benchmarking.profileViews.average * 1.5)) * 100}%`,
                }}
              />
            </div>
            <p className="text-xs text-muted-foreground">
              {benchmarking.profileViews.user > benchmarking.profileViews.average
                ? `Você está ${Math.round((benchmarking.profileViews.user / benchmarking.profileViews.average - 1) * 100)}% acima da média`
                : `Você está ${Math.round((1 - benchmarking.profileViews.user / benchmarking.profileViews.average) * 100)}% abaixo da média`}
            </p>
          </div>

          {/* Taxa de Conversão */}
          <div className="space-y-2">
            <div className="flex items-center justify-between">
              <p className="text-sm font-medium">Taxa de Conversão</p>
              <div className="flex items-center">
                <p className="font-medium">{benchmarking.conversionRate.user}%</p>
                <span className="mx-2 text-muted-foreground">vs</span>
                <p className="text-muted-foreground">
                  {benchmarking.conversionRate.average}% (média)
                </p>
              </div>
            </div>
            <div className="h-2 w-full rounded-full bg-gray-100">
              <div
                className="h-2 rounded-full bg-conecta-green"
                style={{
                  width: `${(benchmarking.conversionRate.user / (benchmarking.conversionRate.average * 1.5)) * 100}%`,
                }}
              />
            </div>
            <p className="text-xs text-muted-foreground">
              {benchmarking.conversionRate.user > benchmarking.conversionRate.average
                ? `Você está ${Math.round((benchmarking.conversionRate.user / benchmarking.conversionRate.average - 1) * 100)}% acima da média`
                : `Você está ${Math.round((1 - benchmarking.conversionRate.user / benchmarking.conversionRate.average) * 100)}% abaixo da média`}
            </p>
          </div>

          {/* Avaliações */}
          <div className="space-y-2">
            <div className="flex items-center justify-between">
              <p className="text-sm font-medium">Avaliação Média</p>
              <div className="flex items-center">
                <p className="font-medium">{benchmarking.ratings.user}</p>
                <span className="mx-2 text-muted-foreground">vs</span>
                <p className="text-muted-foreground">
                  {benchmarking.ratings.average} (média)
                </p>
              </div>
            </div>
            <div className="h-2 w-full rounded-full bg-gray-100">
              <div
                className="h-2 rounded-full bg-yellow-400"
                style={{
                  width: `${(benchmarking.ratings.user / 5) * 100}%`,
                }}
              />
            </div>
            <p className="text-xs text-muted-foreground">
              {benchmarking.ratings.user > benchmarking.ratings.average
                ? `Você está ${Math.round((benchmarking.ratings.user / benchmarking.ratings.average - 1) * 100)}% acima da média`
                : `Você está ${Math.round((1 - benchmarking.ratings.user / benchmarking.ratings.average) * 100)}% abaixo da média`}
            </p>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
