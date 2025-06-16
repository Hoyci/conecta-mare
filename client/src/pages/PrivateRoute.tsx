
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";
import { useToast } from "@/hooks/use-toast";
import { useAuth } from "@/hooks/use-auth";

interface PrivateRouteProps {
  children: JSX.Element;
  requireAuth?: boolean;
}

const PrivateRoute = ({ children, requireAuth = false }: PrivateRouteProps) => {
  const navigate = useNavigate();
  const { toast } = useToast();
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (requireAuth && !isAuthenticated) {
      toast({
        title: "Acesso restrito",
        description: "Faça login para acessar esta página",
        variant: "destructive",
      });
      navigate("/login");
    }
  }, [requireAuth, isAuthenticated, navigate, toast]);

  return children;
};

export default PrivateRoute;
