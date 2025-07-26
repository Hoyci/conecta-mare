import { Link, useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { ArrowLeft } from "lucide-react";

export const OnboardingLayout = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const navigate = useNavigate();

  return (
    <div className="min-h-screen bg-gradient-to-br from-conecta-blue/5 to-conecta-green/5">
      <div className="bg-white shadow-sm border-b">
        <div className="max-w-4xl mx-auto px-4 py-4">
          <div className="flex items-center justify-between">
            <Link to="/" className="flex-shrink-0 flex items-center">
              <span className="text-conecta-blue text-xl font-bold">
                Conecta<span className="text-conecta-green">Maré</span>
              </span>
            </Link>
            <Button
              variant="ghost"
              onClick={() => navigate("/")}
              className="text-gray-500 hover:text-gray-700"
            >
              <ArrowLeft className="w-4 h-4 mr-2" />
              Voltar
            </Button>
          </div>
        </div>
      </div>

      <div className="max-w-6xl mx-auto px-4 py-8">{children}</div>
    </div>
  );
};
