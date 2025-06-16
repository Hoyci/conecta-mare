import { Link, useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { LogOut, User, Gauge } from "lucide-react";
import { useAuth } from "@/hooks/use-auth";
import { useMutation } from "@tanstack/react-query";
import { useToast } from "@/hooks/use-toast";
import { logoutUser } from "@/services/auth-service";

const AccountDropdown = () => {
  const { toast } = useToast();
  const { logout } = useAuth();
  const navigate = useNavigate();

  const { mutate: handleLogout, isPending } = useMutation({
    mutationFn: logoutUser,
    onSuccess: () => {
      logout();
      navigate("/");
    },
    onError: () => {
      toast({
        variant: "destructive",
        title: "Erro ao sair",
        description:
          "Erro ao encerrar sua sessão. Tente novamente mais tarde ou entre em contato com o suporte.",
      });
    },
  });

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="ghost"
          className="flex items-center space-x-2 text-conecta-blue hover:text-conecta-blue-dark"
        >
          <User size={20} />
          <span>Minha Conta</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuItem asChild className="hover:cursor-pointer">
          <Link to="/dashboard" className="flex flex-row gap-2">
            <Gauge size={16} />
            Dashboard
          </Link>
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem
          onClick={() => handleLogout()}
          className="text-red-600 hover:bg-red-50 hover:cursor-pointer flex flex-row gap-2"
        >
          <LogOut size={14} />
          {isPending ? "Saindo..." : "Sair"}
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default AccountDropdown;
