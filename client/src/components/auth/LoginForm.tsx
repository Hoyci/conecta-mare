import { Link, useNavigate } from "react-router-dom";
import { useToast } from "@/components/ui/use-toast";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Checkbox } from "@/components/ui/checkbox";
import { UserIcon, Lock } from "lucide-react";
import { loginUser } from "@/services/auth-service";
import { useMutation } from "@tanstack/react-query";
import { useAuth } from "@/hooks/use-auth";
import { LoginSchema, LoginValues } from "@/types/user";
import { zodResolver } from "@hookform/resolvers/zod";
import { FormProvider, useForm } from "react-hook-form";
import { useMemo } from "react";

const LoginForm = () => {
  const { toast } = useToast();
  const navigate = useNavigate();
  const { login } = useAuth();

  const methods = useForm<LoginValues>({
    resolver: zodResolver(LoginSchema),
    mode: "onChange",
    defaultValues: {
      rememberMe: false,
    },
  });

  const {
    handleSubmit,
    register,
    formState: { errors, isValid },
    watch,
  } = methods;

  const { mutate, isPending } = useMutation({
    mutationFn: (payload: LoginValues) =>
      loginUser(payload.email, payload.password),
    onSuccess: (session, { rememberMe }) => {
      login(session, rememberMe);
      toast({
        title: "Login realizado com sucesso",
        description: "Seja bem-vindo ao ConectaMaré",
      });
      navigate("/dashboard");
    },
    onError: (error: Error) => {
      const errorMessages: Record<string, string> = {
        UNAUTHORIZED_ERROR:
          "E-mail ou senha incorretos. Verifique e tente novamente.",
        DISABLED_USER_ERROR:
          "Esta conta foi desativada. Entre em contato com o suporte.",
      };

      toast({
        variant: "destructive",
        title: "Erro ao fazer login",
        description:
          errorMessages[error.tag] ||
          "Ocorreu um erro inesperado. Tente novamente.",
      });
    },
  });

  const onSubmit = async (data: LoginValues) => {
    mutate(data);
  };

  const password = watch("password");

  const passwordRequirements = useMemo(() => {
    if (!password) return null;

    return {
      length: password.length >= 8,
      uppercase: /[A-Z]/.test(password),
      number: /[0-9]/.test(password),
      symbol: /[^A-Za-z0-9]/.test(password),
    };
  }, [password]);

  return (
    <div className="w-full max-w-md mx-auto">
      <div className="text-center mb-8">
        <h2 className="text-2xl font-bold text-gray-900 mb-2">
          Entrar no ConectaMaré
        </h2>
        <p className="text-gray-600">
          Acesse sua conta para encontrar profissionais ou gerenciar seus
          serviços
        </p>
      </div>

      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
          <div className="space-y-2">
            <Label htmlFor="email">E-mail</Label>
            <div className="relative">
              <UserIcon
                size={18}
                className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
              />
              <Input
                {...register("email")}
                id="email"
                type="email"
                placeholder="Seu e-mail"
                className="pl-10"
              />
            </div>
            {errors.email && (
              <p className="text-red-500 text-xs mt-1">
                {errors.email.message}
              </p>
            )}
          </div>

          <div className="space-y-2">
            <div className="flex justify-between items-center">
              <Label htmlFor="password">Senha</Label>
              <Link
                to="/forgot-password"
                className="text-sm text-conecta-blue hover:underline"
              >
                Esqueceu sua senha?
              </Link>
            </div>
            <div className="relative">
              <Lock
                size={18}
                className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
              />
              <Input
                {...register("password")}
                id="password"
                type="password"
                placeholder="Sua senha"
                className="pl-10"
              />
            </div>
            {password && (
              <div className="mt-1">
                <div className="mt-2 text-xs text-red-600">
                  <p
                    className={
                      passwordRequirements?.length ? "text-conecta-green" : ""
                    }
                  >
                    ✓ Mínimo 8 caracteres
                  </p>
                  <p
                    className={
                      passwordRequirements?.uppercase
                        ? "text-conecta-green"
                        : ""
                    }
                  >
                    ✓ Letra maiúscula
                  </p>
                  <p
                    className={
                      passwordRequirements?.number ? "text-conecta-green" : ""
                    }
                  >
                    ✓ Número
                  </p>
                  <p
                    className={
                      passwordRequirements?.symbol ? "text-conecta-green" : ""
                    }
                  >
                    ✓ Símbolo especial
                  </p>
                </div>
              </div>
            )}
          </div>

          <div className="flex items-center space-x-2">
            <Checkbox
              id="rememberMe"
              {...register("rememberMe")}
              onCheckedChange={(checked) => {
                methods.setValue("rememberMe", Boolean(checked));
                methods.trigger("rememberMe");
              }}
            />
            <Label htmlFor="rememberMe" className="text-sm">
              Lembrar de mim
            </Label>
          </div>

          <Button
            type="submit"
            className="w-full bg-conecta-blue hover:bg-conecta-blue-dark text-white"
            disabled={isPending || !isValid}
          >
            {isPending ? "Entrando..." : "Entrar"}
          </Button>
        </form>
      </FormProvider>

      <div className="mt-6 text-center">
        <p className="text-sm text-gray-600">
          Não tem uma conta?{" "}
          <Link
            to="/signup"
            className="text-conecta-blue font-medium hover:underline"
          >
            Cadastre-se
          </Link>
        </p>
      </div>
    </div>
  );
};

export default LoginForm;
