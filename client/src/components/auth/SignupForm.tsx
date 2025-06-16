import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { User, Mail, Lock, Award } from "lucide-react";

import { Controller, FormProvider } from "react-hook-form";

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { useSignUpForm } from "@/hooks/useSignUpForm";
import { SignUpValues } from "@/types/user";
import ProfilePhotoUpload from "../ui/profile-photo-upload";

const SignupForm = () => {
  const {
    methods,
    handleSubmit,
    register,
    errors,
    setValue,
    isValid,
    mutate,
    isSignUpPending,
    categories,
    isCategoriesLoading,
    userRole,
    passwordRequirements,
  } = useSignUpForm();

  const { watch } = methods;

  const [password, picture] = watch(["password", "picture"]);

  console.log("errors", errors);
  console.log("isValid", isValid);
  console.log("formState", methods.formState);

  const onSubmit = (data: SignUpValues) => {
    const formData = new FormData();
    formData.append("name", data.name);
    formData.append("email", data.email);
    if (userRole !== "client") {
      formData.append("subcategory_id", data.subcategoryId);
    }
    formData.append("password", data.password);
    formData.append("confirm_password", data.confirmPassword);
    formData.append("user_role", data.userRole);
    formData.append("picture", data.picture[0]);

    mutate(formData);
  };

  return (
    <div className="w-full max-w-md mx-auto">
      <div className="text-center mb-6">
        <h2 className="text-2xl font-bold text-gray-900 mb-2">
          Crie sua conta
        </h2>
        <p className="text-gray-600">
          Junte-se ao ConectaMaré para encontrar profissionais qualificados ou
          oferecer seus serviços
        </p>
      </div>

      <FormProvider {...methods}>
        <Controller
          control={methods.control}
          name="userRole"
          defaultValue="client"
          render={({ field }) => (
            <Tabs
              value={field.value.toLowerCase()}
              onValueChange={(value) => field.onChange(value)}
            >
              <TabsList className="grid w-full grid-cols-2 mb-6">
                <TabsTrigger value="client">Sou Cliente</TabsTrigger>
                <TabsTrigger value="professional">Sou Profissional</TabsTrigger>
              </TabsList>
            </Tabs>
          )}
        />
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <ProfilePhotoUpload
            value={picture as File[]}
            onChange={(files) =>
              setValue("picture", files, { shouldValidate: true })
            }
          />
          <div className="space-y-2">
            <Label htmlFor="fullName">Nome</Label>
            <div className="relative">
              <User
                size={18}
                className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
              />
              <Input
                {...register("name")}
                id="name"
                placeholder="Seu nome"
                className="pl-10"
              />
            </div>
            {errors.name && (
              <p className="text-red-500 text-xs mt-1">{errors.name.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="email">E-mail</Label>
            <div className="relative">
              <Mail
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

          {userRole !== "client" && (
            <div className="space-y-2">
              <Label htmlFor="occupation">Especialidade</Label>
              <Controller
                control={methods.control}
                name="subcategoryId"
                defaultValue=""
                render={({ field }) => (
                  <Select
                    onValueChange={field.onChange}
                    value={field.value}
                    disabled={isCategoriesLoading}
                  >
                    <SelectTrigger className="data-[placeholder]:text-gray-400">
                      <Award size={18} className="text-gray-400 shrink-0" />
                      <SelectValue
                        placeholder="Selecione sua especialidade"
                        className="truncate"
                      />
                    </SelectTrigger>
                    <SelectContent>
                      {categories?.map((category) => (
                        <SelectGroup key={category.id}>
                          <SelectLabel>{`${category.name} ${category.icon}`}</SelectLabel>
                          {category.subs.map((subcategory) => (
                            <SelectItem
                              key={subcategory.id}
                              value={subcategory.id}
                            >
                              {subcategory.name}
                            </SelectItem>
                          ))}
                        </SelectGroup>
                      ))}
                    </SelectContent>
                  </Select>
                )}
              />
              {errors.subcategoryId && (
                <p className="text-red-500 text-xs mt-1">
                  {errors.subcategoryId.message}
                </p>
              )}
            </div>
          )}

          <div className="space-y-2">
            <Label htmlFor="password">Senha</Label>
            <div className="relative">
              <Lock
                size={18}
                className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
              />
              <Input
                {...register("password")}
                id="password"
                type="password"
                placeholder="Crie uma senha forte"
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

          <div className="space-y-2">
            <Label htmlFor="confirmPassword">Confirmar senha</Label>
            <div className="relative">
              <Lock
                size={18}
                className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
              />
              <Input
                {...register("confirmPassword")}
                id="confirmPassword"
                type="password"
                placeholder="Confirme sua senha"
                className="pl-10"
              />
            </div>
            {errors.confirmPassword && (
              <p className="text-red-500 text-xs mt-1">
                {errors.confirmPassword.message}
              </p>
            )}
          </div>

          <div className="pt-4">
            <Button
              type="submit"
              className="w-full bg-conecta-blue hover:bg-conecta-blue-dark"
              disabled={isSignUpPending || !isValid}
            >
              {isSignUpPending ? "Cadastrando..." : "Criar conta"}
            </Button>
          </div>
        </form>
      </FormProvider>

      <div className="mt-6 text-center">
        <p className="text-sm text-gray-600">
          Já tem uma conta?{" "}
          <Link
            to="/login"
            className="text-conecta-blue font-medium hover:underline"
          >
            Fazer login
          </Link>
        </p>
      </div>
    </div>
  );
};

export default SignupForm;
