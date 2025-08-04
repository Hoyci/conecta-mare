import { z } from "zod";
import { CategorySchema, SubcategorySchema } from "./categories";

// =================================================================
// CONSTANTES E TIPOS BASE
// =================================================================

const FileListClass = typeof FileList !== "undefined" ? FileList : class {};
export const ROLES = ["client", "professional"] as const;
const DATE_SCHEMA = z.date().nullable();

export const MAX_CERTIFICATIONS = 5;
export const MAX_PROJECTS = 3;
export const MAX_PROJECT_IMAGES = 3;
export const MAX_SERVICES = 3;
export const MAX_SERVICE_IMAGES = 1;
export const MAX_JOB_DESCRIPTION_CHARS = 40;
export const MAX_PROJECT_DESCRIPTION_CHARS = 100;

// =================================================================
// SCHEMAS DE DADOS (MODELO)
// =================================================================

const BaseUserSchema = z.object({
  id: z.string(),
  email: z.string().email(),
  role: z.enum(ROLES),
  createdAt: z.date(),
  updatedAt: DATE_SCHEMA,
  deletedAt: DATE_SCHEMA,
});

const SocialLinksSchema = z.object({
  instagram: z.string().url().optional(),
  linkedin: z.string().optional(),
});

export const LocationSchema = z.object({
  street: z.string(),
  number: z.string(),
  complement: z.string().optional(),
  communityId: z.string(),
});

const FormImageSchema = z.object({
  file: z.instanceof(File),
  url: z.string().url(),
});

export const ProjectImageSchema = z.object({
  id: z.string().optional(),
  url: z.string().url(),
  ordering: z.number(),
  file: z.instanceof(File).optional(),
});

export const ProjectSchema = z.object({
  name: z.string().min(1, "O nome do projeto é obrigatório."),
  description: z.string().min(1, "A descrição do projeto é obrigatória."),
  images: z
    .array(FormImageSchema)
    .min(1, "Adicione pelo menos uma imagem ao projeto.")
    .max(MAX_PROJECT_IMAGES),
});

export const CertificationSchema = z
  .object({
    institution: z.string({ message: "O nome da instituição é obrigatório." }),
    courseName: z.string({ message: "O nome do curso é obrigatório." }),
    startDate: z.coerce.date({
      message: "A data de início do curso é obrigatória",
    }),
    endDate: z.coerce.date().optional().nullable(),
  })
  .refine(
    (data) => {
      if (data.startDate && data.endDate) {
        return data.startDate < data.endDate;
      }
      return true;
    },
    {
      message: "A data de início deve ser anterior à data de término",
      path: ["endDate"],
    },
  );

export const ServiceImageSchema = z.object({
  id: z.string().optional(),
  url: z.string().url(),
  file: z.instanceof(File).optional(),
});

export const ServiceSchema = z.object({
  name: z.string().min(1, "O nome do serviço é obrigatório."),
  description: z.string().min(1, "A descrição do serviço é obrigatória."),
  price: z.coerce.number().min(1, "O preço deve ser um valor positivo."),
  ownLocationPrice: z.coerce.number().optional().nullable(),
  images: z.array(FormImageSchema).max(MAX_SERVICE_IMAGES).optional(),
});

export const UserProfileSchema = BaseUserSchema.extend({
  fullName: z.string().min(2),
  profileImage: z
    .union([z.string().url(), z.instanceof(FileListClass)])
    .optional(),
  jobDescription: z.string().max(MAX_JOB_DESCRIPTION_CHARS),
  phone: z.string().min(14, "Telefone inválido"),
  socialLinks: SocialLinksSchema.optional(),
});

export const ProfessionalProfileSchema = UserProfileSchema.extend({
  subcategoryID: z.string(),
  rating: z.number().min(0).max(5),
  hasOwnLocation: z.boolean(),
  location: LocationSchema.optional(),
  projects: z.array(ProjectSchema).max(MAX_PROJECTS).optional(),
  certifications: z
    .array(CertificationSchema)
    .max(MAX_CERTIFICATIONS)
    .optional(),
  services: z.array(ServiceSchema).max(MAX_SERVICES).optional(),
});

// =================================================================
// SCHEMAS DE AUTENTICAÇÃO
// =================================================================

export const AuthSchema = z.object({
  email: z
    .string({
      required_error: "O e-mail é obrigatório.",
    })
    .email({
      message: "Por favor, insira um endereço de e-mail válido.",
    }),
  password: z
    .string({
      required_error: "A senha é obrigatória.",
    })
    .min(8, {
      message: "A senha deve ter no mínimo 8 caracteres.",
    }),
});

export const SignupSchema = AuthSchema.extend({
  fullName: z
    .string({
      required_error: "O nome é obrigatório.",
    })
    .min(2, {
      message: "O nome deve ter no mínimo 2 caracteres.",
    }),
  confirmPassword: z.string({
    required_error: "A confirmação de senha é obrigatória.",
  }),
  role: z.enum(ROLES, {
    errorMap: () => ({ message: "Por favor, selecione um cargo válido." }),
  }),
}).refine((data) => data.password === data.confirmPassword, {
  message: "As senhas não coincidem",
  path: ["confirmPassword"],
});

export const LoginSchema = AuthSchema.extend({
  rememberMe: z.boolean().default(false),
});

// =================================================================
// SCHEMAS PARA O FORMULÁRIO DE ONBOARDING
// =================================================================
export const OnboardingRequestSchema = z
  .object({
    profileImage: z
      .instanceof(FileListClass)
      .refine((files) => files?.length >= 1, "A foto de perfil é obrigatória."),
    jobDescription: z
      .string()
      .min(1, "Necessário inserir uma descrição sobre seu trabalho")
      .max(MAX_JOB_DESCRIPTION_CHARS),
    subcategoryID: z.string({
      required_error: "Você deve selecionar uma categoria.",
    }),
    phone: z.string().min(15, "O telefone é obrigatório e deve ser válido."),
    socialLinks: z
      .object({
        instagram: z.string().optional(),
        linkedin: z.string().optional(),
      })
      .optional(),
    certifications: z
      .array(CertificationSchema)
      .max(MAX_CERTIFICATIONS)
      .optional(),
    projects: z.array(ProjectSchema).max(MAX_PROJECTS).optional(),
    services: z
      .array(ServiceSchema)
      .min(1, "Você deve adicionar pelo menos um serviço.")
      .max(MAX_SERVICES),
    hasOwnLocation: z.boolean(),
    location: LocationSchema.optional(),
  })
  .refine(
    (data) => {
      if (data.hasOwnLocation) {
        return (
          data.location?.street &&
          data.location?.number &&
          data.location?.communityId
        );
      }
      return true;
    },
    {
      message:
        "O endereço completo é obrigatório se você atende em local próprio.",
      path: ["location.street"],
    },
  )
  .refine(
    (data) => {
      if (data.hasOwnLocation) {
        return data.services.every(
          (s) =>
            typeof s.ownLocationPrice === "number" && s.ownLocationPrice > 0,
        );
      }
      return true;
    },
    {
      message:
        "O preço para atendimento no local próprio é obrigatório para todos os serviços.",
      path: ["services"],
    },
  );

const { fullName, profileImage, jobDescription } = UserProfileSchema.shape;
const { name: subcategoryName } = SubcategorySchema.shape;

export const GetUserResponseSchema = BaseUserSchema.pick({
  id: true,
  email: true,
  role: true,
}).extend({
  fullName: fullName,
  profileImage: profileImage,
  jobDescription: jobDescription,
  subcategoryName: subcategoryName,
});

export const ProfessionalUserResponseSchema = BaseUserSchema.pick({
  email: true,
  role: true,
}).extend({
  userId: z.string(),
  fullName: z.string(),
  profileImage: z.string().url(),
  jobDescription: z.string(),
  socialLinks: SocialLinksSchema.pick({
    instagram: true,
    linkedin: true,
  }).optional(),
  category: CategorySchema.partial({ id: true, name: true }),
  subcategory: SubcategorySchema.partial({ id: true, name: true }),
  rating: z.number(),
  location: LocationSchema.pick({
    street: true,
    number: true,
    complement: true,
    communityId: true,
  }).extend({ communityName: z.string() }),
  certifications: z.array(CertificationSchema).optional(),
  projects: z.array(ProjectSchema).optional(),
  services: z.array(ServiceSchema).optional(),
});

// =================================================================
// TIPOS EXPORTADOS
// =================================================================
export type User = z.infer<typeof BaseUserSchema>;
export type UserProfile = z.infer<typeof UserProfileSchema>;
export type ProjectImage = z.infer<typeof ProjectImageSchema>;
export type Project = z.infer<typeof ProjectSchema>;
export type Certification = z.infer<typeof CertificationSchema>;
export type Service = z.infer<typeof ServiceSchema>;
export type ServiceImage = z.infer<typeof ServiceImageSchema>;
export type ProfessionalProfile = z.infer<typeof ProfessionalProfileSchema>;
export type Roles = (typeof ROLES)[number];

// --- Tipos para autenticação ---
export type SignUpValues = z.infer<typeof SignupSchema>;
export type LoginValues = z.infer<typeof LoginSchema>;

// --- Tipo para o formulário de onboarding ---
export type OnboardingRequestValues = z.infer<typeof OnboardingRequestSchema>;

// --- Tipo de resposta de request ---
export type GetUserResponse = z.infer<typeof GetUserResponseSchema>;
export type ProfessionalUserResponse = z.infer<
  typeof ProfessionalUserResponseSchema
>;
