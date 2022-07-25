import { PaletteColorOptions } from "@mui/material";

export interface DealsTableData {
  id: number;
  name: string;
  progress: string;
  type: string;
  amount: string;
  created: string;
  logo: string;
}

export interface EarningGraphData {
  name: string;
  value: number;
  color: string;
  colorName: PaletteColorOptions;
}

export interface ProgressGraphData {
  name: string;
  actual: number;
  progress: number;
}

export interface QuickStatsData {
  clientsData: {
    count: string;
  };
  invoiceData: {
    count: string;
  };
  totalProjectsData: {
    count: string;
  };
  openProjectsData: {
    count: string;
  };
}

export interface RevenueData {
  ytdRevenue: string;
  clients: number;
  countries: string;
  revenueGraphData: ReviewGraphData[];
}

export interface ReviewGraphData {
  name: string;
  value: number;
}

export interface SocialMediaData {
  id: number;
  name: string;
  revenue: number;
  change: number;
  color: string;
}

export interface StatisticData {
  month: string;
  number: number;
}

export interface StatisticsGraph {
  projectData: StatisticData[];
  clientsData: StatisticData[];
  incomeData: StatisticData[];
}

export interface TicketSupportDataProps {
  id: number;
  name: string;
  ticketId: string;
  created: string;
  contact: string;
  image: string;
}

export interface TodayTaskData {
  id: number;
  task: string;
  date: string;
  isChecked: boolean;
}

export interface WebsiteTrafficData {
  month: number;
  users: number;
}

export interface CRM {
  dealsTableData: DealsTableData[];
  earningGraphData: EarningGraphData[];
  progressGraphData: ProgressGraphData[];
  quickStatsData: QuickStatsData;
  revenueData: RevenueData;
  reviewGraphData: ReviewGraphData[];
  socialMediaData: SocialMediaData[];
  statisticsGraph: StatisticsGraph;
  ticketSupportData: TicketSupportDataProps[];
  todayTaskData: TodayTaskData[];
  websiteTrafficData: WebsiteTrafficData[];
}
