import { z } from "zod";
import { CategorySchema, SubcategorySchema } from "./categories";

const rolesValues = ["client", "professional"] as const;

const rolesEnumSchema = z.enum(rolesValues);

export const UserSchema = z.object({
  id: z.string(),
  email: z.string().email(),
  role: rolesEnumSchema,
  createdAt: z.date(),
  updatesdAt: z.date().nullable(),
  deletedAt: z.date().nullable(),
});

export const UserProfileSchema = z.object({
  id: z.string(),
  userId: z.string(),
  fullName: z.string(),
  categoryId: z.string(),
  subcategoryId: z.string(),
  profileImage: z.string().url(),
  jobDescription: z.string(),
  phone: z.string(),
  socialLinks: z.object({
    github: z.string().optional(),
    instagram: z.string().optional(),
    linkedin: z.string().optional(),
  }),
  createdAt: z.date(),
  updatedAt: z.date().optional(),
});

export const UserProject = z.object({
  id: z.string(),
  userProfileId: z.string(),
  name: z.string(),
  description: z.string(),
  images: z.array(
    z.object({
      id: z.string(),
      url: z.string(),
      ordering: z.number(),
    }),
  ),
});

export const UserCertification = z.object({
  id: z.string(),
  userProfileId: z.string(),
  institution: z.string(),
  courseName: z.string(),
  startDate: z.date(),
  endDate: z.date().optional(),
});

export const SignupSchema = z
  .object({
    name: z.string().min(1, "Nome é obrigatório"),
    email: z.string().min(1, "E-mail é obrigatório").email("E-mail inválido"),
    password: z.string().min(8, "Senha deve ter no mínimo 8 caracteres"),
    confirmPassword: z.string(),
    role: rolesEnumSchema,
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "As senhas não conferem",
    path: ["confirmPassword"],
  });

export const LoginSchema = z.object({
  email: z.string().min(1, "E-mail é obrigatório").email("E-mail inválido"),
  password: z.string().min(8, "Senha deve ter no mínimo 8 caracteres"),
  rememberMe: z.coerce.boolean().default(false),
});

export const ProfessionalUserSchema = z.object({
  userId: UserSchema.shape.id,
  email: UserSchema.shape.email,
  role: UserSchema.shape.role,
  fullName: UserProfileSchema.shape.fullName,
  profileImage: UserProfileSchema.shape.profileImage,
  jobDescription: UserProfileSchema.shape.jobDescription,
  phone: UserProfileSchema.shape.phone,
  socialLinks: UserProfileSchema.shape.socialLinks,
  categoryName: CategorySchema.shape.name,
  subcategoryName: SubcategorySchema.shape.name,
  rating: z.number(),
  location: z.string(),
  projects: z.array(UserProject),
  certifications: z.array(UserCertification),
});

export const ProfessilnaUsersResponseSchema = z.object({
  professionals: z.array(ProfessionalUserSchema),
});

export const ProfessionalUserResponseSchema = z.object({
  data: ProfessionalUserSchema,
});
export type User = z.infer<typeof UserSchema>;
export type SignUpValues = z.infer<typeof SignupSchema>;
export type LoginValues = z.infer<typeof LoginSchema>;
export type ProfessionalUser = z.infer<typeof ProfessionalUserSchema>;
export type ProfessionalUsersResponse = z.infer<
  typeof ProfessilnaUsersResponseSchema
>;
export type ProfessionalUserResponse = z.infer<
  typeof ProfessionalUserResponseSchema
>;

export type ProjectValues = z.infer<typeof UserProject>;
