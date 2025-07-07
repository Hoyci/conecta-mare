import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";
import { User } from "@/types/user";
import { Session, UserPayload } from "@/types/auth";
import { jwtDecode } from "jwt-decode";
import { toCamelCase } from "@/lib/utils.ts";

export type AuthState = {
  user: User | null;
  accessToken: string | null;
  isAuthenticated: boolean;
  login: (session: Session) => void;
  logout: () => void;
  updateAccessToken: (accesToken: string) => void;
};

const createAuthStore = (storage: Storage, name: string) =>
  create<AuthState>()(
    persist(
      (set) => ({
        user: null,
        isAuthenticated: false,
        accessToken: null,
        login: (session: Session) => {
          const decodedUser = jwtDecode<UserPayload>(session.access_token);
          const user = toCamelCase(decodedUser);

          set({
            user,
            accessToken: session.access_token,
            isAuthenticated: true,
          });
        },
        logout: () =>
          set({ user: null, accessToken: null, isAuthenticated: false }),
        updateAccessToken: (accessToken: string) =>
          set((state) => ({ ...state, accessToken })),
      }),
      {
        name,
        storage: createJSONStorage(() => storage),
        partialize: (state) => ({
          user: state.user,
          accessToken: state.accessToken,
          isAuthenticated: state.isAuthenticated,
        }),
      },
    ),
  );


const isBrowser = typeof window !== "undefined";

const noopStorage: Storage = {
  getItem: (_: string) => null,
  setItem: (_: string, _value: string) => { },
  removeItem: (_: string) => { },
  clear: () => { },
  key: (_index: number) => null,
  length: 0,
};

const localStorageOrNoop = isBrowser ? window.localStorage : noopStorage;
const sessionStorageOrNoop = isBrowser ? window.sessionStorage : noopStorage;

export const useLocalAuthStore = createAuthStore(
  localStorageOrNoop,
  "auth-local",
);

export const useSessionAuthStore = createAuthStore(
  sessionStorageOrNoop,
  "auth-session",
);
