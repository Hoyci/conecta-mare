
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ServiceMetric } from "@/types/metrics";
import { ArrowUp, ArrowDown, Minus } from "lucide-react";

interface ServicesPerformanceProps {
  services: ServiceMetric[];
}

export function ServicesPerformance({ services }: ServicesPerformanceProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-lg">Desempenho dos Serviços</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="space-y-6">
          {services.map((service) => (
            <div key={service.id} className="flex items-center justify-between">
              <div className="space-y-1">
                <p className="font-medium">{service.name}</p>
                <div className="flex text-sm text-muted-foreground">
                  <p>{service.views} visualizações</p>
                  <span className="mx-2">•</span>
                  <p>{service.conversions} conversões</p>
                </div>
              </div>
              <div className="flex items-center">
                {service.trend === "up" && (
                  <div className="flex items-center text-conecta-green">
                    <ArrowUp className="h-4 w-4 mr-1" />
                    <span>{service.trendPercentage}%</span>
                  </div>
                )}
                {service.trend === "down" && (
                  <div className="flex items-center text-destructive">
                    <ArrowDown className="h-4 w-4 mr-1" />
                    <span>{service.trendPercentage}%</span>
                  </div>
                )}
                {service.trend === "stable" && (
                  <div className="flex items-center text-gray-500">
                    <Minus className="h-4 w-4 mr-1" />
                    <span>{service.trendPercentage}%</span>
                  </div>
                )}
              </div>
            </div>
          ))}
        </div>
      </CardContent>
    </Card>
  );
}
