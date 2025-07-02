import { z } from "zod";

// ======================
// Constantes Compartilhadas
// ======================
const ROLES = ["client", "professional"] as const;
const DATE_SCHEMA = z.date().nullable();
export const MAX_CERTIFICATIONS = 5;
export const MAX_PROJECTS = 3;
export const MAX_PROJECT_IMAGES = 3;
export const MAX_SERVICES = 3;
export const MAX_SERVICE_IMAGES = 1;
export const MAX_JOB_DESCRIPTION_CHARS = 80;

// ======================
// Schemas Base
// ======================
const BaseUserSchema = z.object({
  id: z.string(),
  email: z.string().email(),
  role: z.enum(ROLES),
  createdAt: z.date(),
  updatedAt: DATE_SCHEMA,
  deletedAt: DATE_SCHEMA,
});

const SocialLinksSchema = z
  .object({
    instagram: z.string().url().optional(),
    linkedin: z.string().url("Insira uma URL válida"),
  })
  .optional();

// ======================
// Schemas Principais
// ======================
export const UserProfileSchema = BaseUserSchema.extend({
  fullName: z.string().min(2),
  profileImage: z.union([z.string().url(), z.instanceof(FileList)]).optional(),
  jobDescription: z.string().max(MAX_JOB_DESCRIPTION_CHARS),
  phone: z.string().min(14, "Telefone inválido"),
  socialLinks: SocialLinksSchema,
});

// ======================
// Schemas de Projetos e Certificações
// ======================
export const ProjectImageSchema = z.object({
  id: z.string().optional(),
  url: z.string().url(),
  ordering: z.number(),
  file: z.instanceof(File).optional(),
});

export const ProjectSchema = z.object({
  id: z.string().optional(),
  name: z.string(),
  description: z.string(),
  images: z.array(ProjectImageSchema).max(MAX_PROJECT_IMAGES),
});

export const CertificationSchema = z.object({
  id: z.string().optional(),
  institution: z.string().min(3),
  courseName: z.string().min(3),
  startDate: z.coerce.date().nullable(),
  endDate: z.coerce.date().optional().nullable(),
});

export const ServiceImageSchema = z.object({
  id: z.string().optional(),
  url: z.string().url(),
  file: z.instanceof(File).optional(),
});

export const ServiceSchema = z.object({
  name: z.string(),
  description: z.string(),
  price: z.number(),
  ownLocationPrice: z.number().optional(),
  images: z.array(ServiceImageSchema).max(MAX_SERVICE_IMAGES),
});

export const LocationSchema = z.object({
  street: z.string(),
  number: z.string(),
  complement: z.string(),
  neighborhood: z.string(),
});

// ======================
// Schemas de Autenticação
// ======================
export const AuthSchema = z.object({
  email: z.string().email(),
  password: z.string().min(8),
});

export const SignupSchema = AuthSchema.extend({
  name: z.string().min(2),
  confirmPassword: z.string(),
  role: z.enum(ROLES),
}).refine((data) => data.password === data.confirmPassword, {
  message: "As senhas não coincidem",
  path: ["confirmPassword"],
});

export const LoginSchema = AuthSchema.extend({
  rememberMe: z.boolean().default(false),
});

// ======================
// Schemas de Resposta
// ======================
export const ProfessionalProfileSchema = UserProfileSchema.extend({
  subcategoryID: z.string(),
  rating: z.number().min(0).max(5),
  hasOwnLocation: z.boolean(),
  location: LocationSchema,
  projects: z.array(ProjectSchema).max(MAX_PROJECTS),
  certifications: z.array(CertificationSchema).max(MAX_CERTIFICATIONS),
  services: z.array(ServiceSchema).max(MAX_SERVICES),
});

// ======================
// Tipos Exportados
// ======================
export type User = z.infer<typeof BaseUserSchema>;
export type UserProfile = z.infer<typeof UserProfileSchema>;
export type ProjectImage = z.infer<typeof ProjectImageSchema>;
export type Project = z.infer<typeof ProjectSchema>;
export type Certification = z.infer<typeof CertificationSchema>;
export type Service = z.infer<typeof ServiceSchema>;
export type ServiceImage = z.infer<typeof ServiceImageSchema>;
export type SignUpValues = z.infer<typeof SignupSchema>;
export type LoginValues = z.infer<typeof LoginSchema>;
export type ProfessionalProfile = z.infer<typeof ProfessionalProfileSchema>;
