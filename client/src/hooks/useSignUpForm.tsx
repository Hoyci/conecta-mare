import { useToast } from "@/components/ui/use-toast";
import { useNavigate } from "react-router-dom";
import { useMutation } from "@tanstack/react-query";
import { signUpUser } from "@/services/auth-service";
import { useMemo } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { SignupSchema, SignUpValues } from "@/types/user";

export function useSignUpForm() {
  const { toast } = useToast();
  const navigate = useNavigate();
  const methods = useForm<SignUpValues>({
    resolver: zodResolver(SignupSchema),
    mode: "onChange",
  });

  const {
    handleSubmit,
    register,
    formState: { errors, isValid },
    watch,
    setValue,
  } = methods;

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

  const { mutate, isPending: isSignUpPending } = useMutation({
    mutationFn: (payload: SignUpValues) => signUpUser(payload),
    onSuccess: () => {
      toast({
        title: "Cadastro realizado com sucesso",
        description: "Seja bem-vindo ao ConectaMaré",
      });
      navigate("/login");
    },
    onError: (error: any) => {
      const errorMessages: Record<string, string> = {
        CONFLICT_ERROR: "Esse e-mail já está sendo usado.",
        UNPROCESSABLE_ENTITY_ERROR: "Dados inválidos ou incompletos.",
      };
      toast({
        variant: "destructive",
        title: "Erro ao fazer login",
        description:
          errorMessages[error.tag] || "Erro inesperado. Tente novamente.",
      });
    },
  });

  return {
    methods,
    handleSubmit,
    register,
    errors,
    setValue,
    isValid,
    mutate,
    isSignUpPending,
    passwordRequirements,
  };
}
