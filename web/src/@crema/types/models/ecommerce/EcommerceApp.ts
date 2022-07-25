export interface BrandData {
  id: number;
  name: string;
}

export interface ProductColors {
  BLUE: string;
  GREY: string;
  PARROT: string;
  LIGHT_PINK: string;
  PINK: string;
}

export interface IdealFor {
  id: number;
  name: string;
}

export interface Addresses {
  id: number;
  name: string;
  mobile: string;
  addressLine: string;
  city: string;
  pin: string;
}

export interface DiscountList {
  id: number;
  name: string;
}

export interface CartItems {
  id: number;
  title: string;
  mrp: string;
  discount: string;
  brand: number | string;
  image: string;
  count: number;
}

export interface ProductData {
  id: number;
  title: string;
  description: string;
  mrp: string;
  discount: string;
  rating: number;
  ideaFor: number;
  brand: number;
  color: string;
  reviews: number;
  image: {
    id: number;
    src: string;
  }[];
}

export interface RecentOrders {
  id: string;
  customer: string;
  product: string;
  date: string;
  paymentType: string;
  price: string;
  status: string;
}

export interface FilterData {
  title: string;
  page?: string | number;
  brand: number[];
  ideaFor: number[];
  discount: number[];
  color: any[];
  rating: number[];
}

export interface RecentOrders {
  id: string;
  customer: string;
  product: string;
  date: string;
  paymentType: string;
  price: string;
  status: string;
}

export interface CustomersData {
  id: number;
  name: string;
  email: string;
  lastItem: string;
  lastOrder: string;
  rating: string;
  balance: string;
  address: string;
  joinDate: string;
}
