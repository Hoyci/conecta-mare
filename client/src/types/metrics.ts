import { z } from "zod";

const DailyVisitSchema = z.object({
  date: z.string().datetime(),
  visits: z.number().int().nonnegative(),
});

const MetricsSchema = z.object({
  currentWeekVisits: z.number().int().nonnegative(),
  percentageChange: z.number(),
  currentWeekData: z.array(DailyVisitSchema),
});

export const UserProfileViewsResponseSchema = z.object({
  metrics: MetricsSchema,
});

export type DailyVisit = z.infer<typeof DailyVisitSchema>;
export type Metrics = z.infer<typeof MetricsSchema>;
export type UserProfileViewsResponse = z.infer<
  typeof UserProfileViewsResponseSchema
>;
