import { useLocalAuthStore } from "@/store/auth";
import { useSessionAuthStore } from "@/store/auth";
import { getCurrentAuthStore } from "@/store/auth-store-utils";
import { Session } from "@/types/auth";

export const useAuth = () => {
  const store = getCurrentAuthStore()

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
