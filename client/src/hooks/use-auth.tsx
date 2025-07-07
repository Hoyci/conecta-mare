import { useState, useEffect } from "react";
import type { AuthState } from "@/store/auth";
import { Session } from "@/types/auth";

export const useAuth = () => {
  const [store, setStore] = useState<AuthState | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    import("@/store/auth-store-utils").then(({ getCurrentAuthStore }) => {
      const authStore = getCurrentAuthStore();
      setStore(authStore);
      setIsLoading(false);
    });
  }, []);

  if (isLoading || !store) {
    return {
      user: null,
      accessToken: null,
      isAuthenticated: false,
      isLoading: true,
      login: (_session: Session, _rememberMe: boolean) => { },
      logout: () => { },
      updateAccessToken: (_accessToken: string) => { },
    };
  }

  return {
    user: store.user,
    accessToken: store.accessToken,
    isAuthenticated: store.isAuthenticated,
    isLoading: false,
    login: store.login,
    logout: store.logout,
    updateAccessToken: store.updateAccessToken,
  };
};
