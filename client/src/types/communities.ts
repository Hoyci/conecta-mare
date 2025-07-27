import { z } from "zod";

export const CommunitySchema = z.object({
  id: z.string(),
  name: z.string(),
  censo_id: z.string()
});

export const GetCommunitiesResponseSchema = z.object({
  communities: z.array(CommunitySchema),
});

export type Community = z.infer<typeof CommunitySchema>;
export type GetCommunitiesResponse = z.infer<
  typeof GetCommunitiesResponseSchema
>;
