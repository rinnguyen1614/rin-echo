import {
  CartItems,
  CustomersData,
  FilterData,
  ProductData,
  RecentOrders,
} from "../models/ecommerce/EcommerceApp";
import { VIEW_TYPE } from "../../redux/reducers/Ecommerce";

export const GET_CART_ITEM = "GET_CART_ITEM";
export const ADD_CART_ITEM = "ADD_CART_ITEM";
export const GET_CUSTOMERS = "GET_CUSTOMERS";
export const GET_ECOMMERCE_LIST = "GET_ECOMMERCE_LIST";
export const GET_RECENT_ORDER = "GET_RECENT_ORDER";
export const REMOVE_CART_ITEM = "REMOVE_CART_ITEM";
export const SET_FILTER_DATA = "SET_FILTER_DATA";
export const SET_PRODUCT_DATA = "SET_PRODUCT_DATA";
export const SET_PRODUCT_VIEW_TYPE = "SET_PRODUCT_VIEW_TYPE";
export const UPDATE_CART_ITEM = "UPDATE_CART_ITEM";

export interface GetEcommerceAction {
  type: typeof GET_ECOMMERCE_LIST;
  payload: ProductData[];
}

export interface SetProductDetailAction {
  type: typeof SET_PRODUCT_DATA;
  payload: ProductData;
}

export interface GetRecentOrdersAction {
  type: typeof GET_RECENT_ORDER;
  payload: {
    orders: RecentOrders[];
    orderCount: number;
  };
}

export interface GetCustomersAction {
  type: typeof GET_CUSTOMERS;
  payload: { customers: CustomersData[]; customerCount: number };
}

export interface GetCartItemsAction {
  type: typeof GET_CART_ITEM;
  payload: CartItems[];
}

export interface AddCartItemsAction {
  type: typeof ADD_CART_ITEM;
  payload: ProductData;
}

export interface RemoveCartItemsAction {
  type: typeof REMOVE_CART_ITEM;
  payload: CartItems;
}

export interface UpdateCartItemsAction {
  type: typeof UPDATE_CART_ITEM;
  payload: CartItems;
}

export interface SetProductViewTypeAction {
  type: typeof SET_PRODUCT_VIEW_TYPE;
  payload: VIEW_TYPE;
}

export interface SetFilterDataAction {
  type: typeof SET_FILTER_DATA;
  payload: FilterData;
}

export type EcommerceActionTypes =
  | GetEcommerceAction
  | SetProductDetailAction
  | GetRecentOrdersAction
  | GetCustomersAction
  | GetCartItemsAction
  | AddCartItemsAction
  | RemoveCartItemsAction
  | UpdateCartItemsAction
  | SetProductViewTypeAction
  | SetFilterDataAction;
