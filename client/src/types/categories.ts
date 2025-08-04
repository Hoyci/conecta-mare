import { z } from "zod";

export const CategorySchema = z.object({
  id: z.string(),
  name: z.string(),
  icon: z.string(),
});

export const SubcategorySchema = z.object({
  id: z.string(),
  name: z.string(),
  categoryId: z.string(),
});

export const SubcategoryWithUserCountSchema = CategorySchema.extend({
  user_count: z.number(),
});

export const CategoryWithUserCountResponseSchema = z.object({
  categories: z.array(SubcategoryWithUserCountSchema),
});

export const CategoryWithSubs = CategorySchema.extend({
  subcategories: z.array(SubcategorySchema),
});

export const CategoryWithSubsResponseSchema = z.object({
  categories: z.array(CategoryWithSubs),
});

export type Category = z.infer<typeof CategorySchema>;
export type Subcategory = z.infer<typeof SubcategorySchema>;
export type CategoryWithUserCount = z.infer<
  typeof SubcategoryWithUserCountSchema
>;
export type CategoryWithUserCountResponse = z.infer<
  typeof CategoryWithUserCountResponseSchema
>;
export type CategoryWithSubsResponse = z.infer<
  typeof CategoryWithSubsResponseSchema
>;
