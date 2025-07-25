import { env } from "@/config/env";
import { RudderAnalytics } from "@rudderstack/analytics-js";

let analyticsInstance: RudderAnalytics | null = null;

export const getAnalytics = (): RudderAnalytics => {
  if (!analyticsInstance) {
    analyticsInstance = new RudderAnalytics();
    analyticsInstance.load(
      env.data.VITE_RUDDER_WRITE_KEY,
      env.data.VITE_DATA_PLANE_URL,
      {},
    );
  }
  return analyticsInstance;
};
