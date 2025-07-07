import { Navigate, Outlet, useLocation } from "react-router-dom";
import { useAuth } from "@/hooks/use-auth";
import { useToast } from "@/hooks/use-toast";
import { useEffect } from "react";

const FullScreenLoader = () => (
  <div className="flex h-screen w-full items-center justify-center bg-background">
    <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-conecta-blue"></div>
  </div>
);

const ProtectedLayout = () => {
  const { isAuthenticated, isLoading } = useAuth();
  const location = useLocation();
  const { toast } = useToast();

  useEffect(() => {
    if (location.state?.from) {
      toast({
        title: "Acesso restrito",
        description: "Faça login para acessar esta página.",
        variant: "destructive",
      });
    }
  }, [location.state, toast]);

  if (isLoading) {
    return <FullScreenLoader />;
  }

  if (!isAuthenticated) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return <Outlet />;
};

export default ProtectedLayout;
