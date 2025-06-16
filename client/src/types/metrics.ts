
export type TimeRange = "daily" | "weekly" | "monthly";

export type ProfileMetrics = {
  profileViews: {
    total: number;
    trends: { [key in TimeRange]: { date: string; value: number }[] };
  };
  conversionRate: {
    rate: number;
    trends: { [key in TimeRange]: { date: string; value: number }[] };
  };
  ratings: {
    average: number;
    total: number;
    recent: ClientFeedback[];
  };
  services: ServiceMetric[];
  benchmarking: {
    profileViews: { user: number; average: number };
    conversionRate: { user: number; average: number };
    ratings: { user: number; average: number };
  };
  notifications: Notification[];
};

export type ClientFeedback = {
  id: string;
  clientName: string;
  clientAvatar?: string;
  rating: number;
  comment: string;
  date: string;
};

export type ServiceMetric = {
  id: string;
  name: string;
  views: number;
  conversions: number;
  trend: "up" | "down" | "stable";
  trendPercentage: number;
};

export type Notification = {
  id: string;
  type: "milestone" | "review" | "tip";
  title: string;
  description: string;
  date: string;
  read: boolean;
};
