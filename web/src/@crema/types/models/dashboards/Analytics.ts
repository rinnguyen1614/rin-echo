export interface WelcomeCardData {
  id: number;
  type: string;
  counts: number;
  icon: string;
}

export interface RevenueCards {
  id: number;
  type: string;
  value: string;
  growth: number;
  icon: string;
  strokeColor: string;
  graphData: { month: string; number: number }[];
}

export interface SalesStateData {
  id: number;
  amount: string;
  type: string;
  icon: string;
}

export interface ChartData {
  name: string;
  AS: number;
  Rev: number;
  amt: number;
}

export interface VisitorsPageView {
  name: string;
  Page: number;
  Visitor: number;
}

export interface ActiveVisitorsProps {
  growth: number;
  value: number;
  slug: string;
  graphData: { time: string; value: number }[];
}

export interface TopSellingProduct {
  id: number;
  icon: string;
  name: string;
  description: string;
  price: number;
  rate: number;
  color: string;
}

export interface EarningData {
  id: number;
  color: string;
  amount: number;
  country: string;
}

export interface Tickets {
  id: number;
  name: string;
  opened: number;
  overAllPercentage: {
    open: number;
    close: number;
    hold: number;
  };
}

export interface PageVisit {
  id: number;
  page: string;
  pageView: number;
  visitors: number;
}

export interface TransactionData {
  id: string;
  customer: string;
  date: string;
  paymentType: string;
  status: string;
}

export interface InfoWidgets {
  id: number;
  icon: string;
  count: number | string;
  details: string;
  bgColor?: string;
}

export interface TrafficData {
  id: number;
  title: string;
  value: number;
  session: number;
}

export interface Analytics {
  welcomeCard: WelcomeCardData[];
  revenueCards: RevenueCards[];
  salesState: SalesStateData[];
  salesChartData: ChartData[];
  visitorsPageView: VisitorsPageView[];
  activeVisitors: ActiveVisitorsProps;
  topSellingProduct: TopSellingProduct[];
  earningData: EarningData[];
  tickets: Tickets[];
  pageVisits: PageVisit[];
  transactionData: TransactionData[];
  infoWidgets: InfoWidgets[];
  trafficData: TrafficData[];
}
