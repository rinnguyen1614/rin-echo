import { PaletteColorOptions } from "@mui/material";

export interface IncomeLastYear {
  value: string;
  graphData: { name: string; revenue: number }[];
}

export interface WebsiteTrafficData {
  value: string;
  graphData: { name: string; traffic: number }[];
}

export interface RevenueGrowthData {
  value: string;
  graphData: { name: string; growth: number }[];
}

export interface IncrementActiveUsers {
  value: string;
  graphData: { name: string; activeUsers: number }[];
}

export interface ExtraRevenue {
  value: string;
  graphData: { name: string; revenue: number }[];
}

export interface TrafficRaise {
  value: string;
  graphData: { name: string; traffic: number }[];
}

export interface LessOrders {
  value: string;
  graphData: { name: string; orders: number }[];
}

export interface EarningInMonthData {
  id: number;
  name: string;
  value: number;
  color: string;
  colorName: PaletteColorOptions | string;
}

export interface SubscriptionData {
  dataOne: { number: number; value: number }[];
  dataTwo: { number: number; value: number }[];
  dataThree: { number: number; value: number }[];
}

export interface MetricsLineGraphData {
  value: string;
  difference: string;
  differencePercent: string;
  graphData: { number: string; value: number }[];
}

export interface SalesData {
  salesToday: string;
  salesYesterday: string;
  salesGraphData: {
    day: number;
    number: number;
  }[];
}

export interface MetricsFloatingChildData {
  value: string;
  change: number;
  strokeColor: string;
  areaColor: string;
  graphData: {
    number: string;
    value: number;
  }[];
}

export interface MetricsFloatingGraphData {
  salesData: MetricsFloatingChildData;
  clientsData: MetricsFloatingChildData;
  revenueData: MetricsFloatingChildData;
  newUserData: MetricsFloatingChildData;
}

export interface VisitsData {
  new: number;
  returning: number;
  graphData: {
    dataOne: { number: string; value: number }[];
    dataTwo: { number: string; value: number }[];
    dataThree: { number: string; value: number }[];
  };
}

export interface OrdersData {
  revenue: number;
  orders: number;
  graphData: {
    dataOne: { number: string; value: number }[];
    dataTwo: { number: string; value: number }[];
    dataThree: { number: string; value: number }[];
  };
}

export interface ProfileViewsData {
  views: string;
  graphData: { day: number; number: number }[];
}

export interface WorkViewsData {
  views: string;
  graphData: { name: string; value: number }[];
}

export interface SocialData {
  likes: number;
  comments: number;
}

export interface StateGraphData {
  number: number;
  value: number;
}

export interface StatsGraphData {
  dataOne: {
    stats1: StateGraphData[];
    stats2: StateGraphData[];
  };
  dataTwo: {
    stats1: StateGraphData[];
    stats2: StateGraphData[];
  };
  dataThree: {
    stats1: StateGraphData[];
    stats2: StateGraphData[];
  };
}

export interface SocialVisitorsData {
  id: number;
  name: string;
  visitors: number;
  change: number;
  color: string;
}

export interface AccountData {
  name: string;
  complete: number;
  week: number;
}

export interface ShareData {
  icon: string;
  color: string;
  value: number;
}

export interface Metrics {
  ordersThisYear: string;
  queryIcon: string;
  revenueThisYear: string;
  visitsThisYear: string;
  queriesThisYear: string;
  websiteTraffic: string;
  incomeLastYear: IncomeLastYear;
  websiteTrafficData: WebsiteTrafficData;
  revenueGrowthData: RevenueGrowthData;
  incrementActiveUsers: IncrementActiveUsers;
  extraRevenue: ExtraRevenue;
  trafficRaise: TrafficRaise;
  lessOrders: LessOrders;
  salesData: SalesData;
  earningInMonth: EarningInMonthData[];
  subscriptionData: SubscriptionData;
  metricsLineGraphData: MetricsLineGraphData;
  metricsFloatingGraphData: MetricsFloatingGraphData;
  visitsData: VisitsData;
  ordersData: OrdersData;
  profileViewsData: ProfileViewsData;
  workViewsData: WorkViewsData;
  socialData: SocialData;
  statsGraph: StatsGraphData;
  socialVisitorsData: SocialVisitorsData[];
  accountData: AccountData[];
  shareData: ShareData[];
}
