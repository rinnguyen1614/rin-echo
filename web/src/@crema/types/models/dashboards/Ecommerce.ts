export interface SalesStateData {
  id: number;
  type: string;
  value: string;
  bgColor: string;
  icon: string;
}

export interface ReportCards {
  id: number;
  type: string;
  value: string;
  growth: number;
  icon: string;
  strokeColor: string;
  graphData: {
    month: string;
    number: number;
  }[];
}

export interface PopularProductData {
  id: number;
  icon: string;
  name: string;
  description: string;
  price: number;
  mrp: number;
}

export interface MarketingCampaignData {
  id: number;
  name: string;
  description: string;
  icon: string;
  graph: number;
  growth: boolean;
  spent: string;
  like?: string;
  share?: string;
}

export interface NotificationsData {
  id: number;
  image: string;
  name: string;
  type: string;
  message: string;
}

export interface NewCustomersData {
  id: number;
  image: string;
  name: string;
  orders: number;
  color: string;
  message: string;
}

export interface SiteVisitorData {
  id: number;
  color: string;
  value: number;
  icon: string;
  country: string;
}

export interface BrowserData {
  id: number;
  value: string;
  name: string;
  icon: string;
}

export interface RecentOrderData {
  id: string;
  customer: string;
  product: string;
  date: string;
  paymentType: string;
  price: string;
  status: string;
}

export interface Ecommerce {
  salesState: SalesStateData[];
  reportCards: ReportCards[];
  popularProducts: PopularProductData[];
  marketingCampaign: MarketingCampaignData[];
  notifications: NotificationsData[];
  newCustomers: NewCustomersData[];
  siteVisitors: SiteVisitorData[];
  browser: BrowserData[];
  recentOrders: RecentOrderData[];
}
