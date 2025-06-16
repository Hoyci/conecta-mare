import { useLocalAuthStore, useSessionAuthStore } from "./auth"

export const getCurrentAuthStore = () => {
  const local = useLocalAuthStore.getState()
  const session = useSessionAuthStore.getState()

  const isLocalActive = local.isAuthenticated
  return isLocalActive ? local : session
}
