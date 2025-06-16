
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ClientFeedback } from "@/types/metrics";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { StarFilledIcon } from "@radix-ui/react-icons";

interface FeedbackSectionProps {
  average: number;
  total: number;
  feedback: ClientFeedback[];
}

export function FeedbackSection({ average, total, feedback }: FeedbackSectionProps) {
  return (
    <Card>
      <CardHeader>
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between">
          <CardTitle className="text-lg">Avaliações dos Clientes</CardTitle>
          <div className="flex items-center mt-2 sm:mt-0">
            <div className="flex">
              {[...Array(5)].map((_, i) => (
                <StarFilledIcon
                  key={i}
                  className={`h-4 w-4 ${
                    i < Math.floor(average)
                      ? "text-yellow-400"
                      : "text-gray-300"
                  }`}
                />
              ))}
            </div>
            <p className="ml-2 text-sm font-medium">
              {average} de 5 ({total} avaliações)
            </p>
          </div>
        </div>
      </CardHeader>
      <CardContent>
        <div className="space-y-6">
          {feedback.map((item) => (
            <div key={item.id} className="space-y-3">
              <div className="flex items-center justify-between">
                <div className="flex items-center">
                  <Avatar className="h-8 w-8">
                    <AvatarImage src={item.clientAvatar} />
                    <AvatarFallback>{item.clientName.charAt(0)}</AvatarFallback>
                  </Avatar>
                  <div className="ml-3">
                    <p className="font-medium">{item.clientName}</p>
                    <div className="flex">
                      {[...Array(5)].map((_, i) => (
                        <StarFilledIcon
                          key={i}
                          className={`h-3 w-3 ${
                            i < item.rating ? "text-yellow-400" : "text-gray-300"
                          }`}
                        />
                      ))}
                    </div>
                  </div>
                </div>
                <p className="text-sm text-muted-foreground">
                  {new Date(item.date).toLocaleDateString("pt-BR")}
                </p>
              </div>
              <p className="text-muted-foreground">{item.comment}</p>
            </div>
          ))}
          <button className="w-full text-center text-sm font-medium text-conecta-blue hover:underline">
            Ver todas as avaliações
          </button>
        </div>
      </CardContent>
    </Card>
  );
}
