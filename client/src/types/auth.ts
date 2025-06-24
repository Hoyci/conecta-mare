import { User } from "./user";

export type UserPayload = {
  user: User;
  exp: number;
  iat: number;
  iss: string;
  jti: string;
};

export type Session = {
  access_token: string;
  refresh_token: string;
};

export type RenewAccessTokenResponse = {
  access_token: string;
};
