import { z } from "zod";

const envSchema = z.object({
  VITE_PORT: z.coerce.number().default(3000),
  VITE_API_URL: z.string().url().default("http://localhost:8080"),

  VITE_ENV: z
    .enum(["development", "production", "test"])
    .default("development"),
});

export const env = envSchema.safeParse(import.meta.env);

if (!env.success) {
  console.error("❌ Invalid environment variables:", env.error.format());
  throw new Error("Invalid environment variables");
}

export type Env = z.infer<typeof envSchema>;
