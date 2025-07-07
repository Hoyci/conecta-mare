import { useState, useEffect } from "react";
import { useLocalAuthStore } from "@/store/auth";
import { useSessionAuthStore } from "@/store/auth";
import type { AuthState } from "@/store/auth";
import { Session } from "@/types/auth";

export const useAuth = () => {
  const [store, setStore] = useState<AuthState | null>(null);

  useEffect(() => {
    import("@/store/auth-store-utils").then(({ getCurrentAuthStore }) => {
      setStore(getCurrentAuthStore());
    });
  }, []);

  if (!store) {
    return {
      user: null,
      isAuthenticated: false,
      login: () => { },
      logout: () => { },
    };
  }

  const login = (session: Session, rememberMe: boolean) => {
    if (rememberMe) {
      useLocalAuthStore.getState().login(session);
    } else {
      useSessionAuthStore.getState().login(session);
    }
  };

  const logout = () => {
    useLocalAuthStore.getState().logout();
    useSessionAuthStore.getState().logout();
  };

  return {
    user: store.user,
    accessToken: store.accessToken,
    isAuthenticated: store.isAuthenticated,
    login,
    logout,
    updateAccessToken: store.updateAccessToken,
  };
};
