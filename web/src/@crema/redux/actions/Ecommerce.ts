import { Dispatch } from "redux";
import jwtAxios from "../../services/auth/jwt-auth";
import { AppActions } from "../../types";
import {
  ADD_CART_ITEM,
  GET_CUSTOMERS,
  GET_ECOMMERCE_LIST,
  GET_RECENT_ORDER,
  REMOVE_CART_ITEM,
  SET_FILTER_DATA,
  SET_PRODUCT_DATA,
  SET_PRODUCT_VIEW_TYPE,
  UPDATE_CART_ITEM,
} from "../../types/actions/Ecommerce.action";
import {
  CartItems,
  FilterData,
  ProductData,
} from "../../types/models/ecommerce/EcommerceApp";
import { VIEW_TYPE } from "../reducers/Ecommerce";
import { fetchError, fetchStart, fetchSuccess } from "./Common";

export const onGetEcommerceData = (filterData: FilterData) => {
  return (dispatch: Dispatch<AppActions>) => {
    dispatch(fetchStart());
    jwtAxios
      .get("/api/ecommerce/list", {
        params: { filterData },
      })
      .then((data) => {
        if (data.status === 200) {
          dispatch(fetchSuccess());
          dispatch({ type: GET_ECOMMERCE_LIST, payload: data.data.list });
        } else {
          dispatch(fetchError("Something went wrong, Please try again!"));
        }
      })
      .catch((error) => {
        dispatch(fetchError(error.message));
      });
  };
};
export const getProductDetail = (id: number | string) => {
  return (dispatch: Dispatch<AppActions>) => {
    dispatch(fetchStart());
    jwtAxios
      .get("/api/ecommerce/get", {
        params: { id: id },
      })
      .then((data) => {
        if (data.status === 200) {
          dispatch(fetchSuccess());
          dispatch({ type: SET_PRODUCT_DATA, payload: data.data });
        } else {
          dispatch(fetchError("Something went wrong, Please try again!"));
        }
      })
      .catch((error) => {
        dispatch(fetchError(error.message));
      });
  };
};
export const getRecentOrders = (search: string, page: number) => {
  return (dispatch: Dispatch<AppActions>) => {
    dispatch(fetchStart());
    jwtAxios
      .get("/api/ecommerce/orders", {
        params: { search, page },
      })
      .then((data) => {
        if (data.status === 200) {
          dispatch(fetchSuccess());
          dispatch({ type: GET_RECENT_ORDER, payload: data.data });
        } else {
          dispatch(fetchError("Something went wrong, Please try again!"));
        }
      })
      .catch((error) => {
        dispatch(fetchError(error.message));
      });
  };
};
export const getCustomers = (search: string, page: number) => {
  return (dispatch: Dispatch<AppActions>) => {
    dispatch(fetchStart());
    jwtAxios
      .get("/api/ecommerce/customers", {
        params: { search, page },
      })
      .then((data) => {
        if (data.status === 200) {
          dispatch(fetchSuccess());
          dispatch({ type: GET_CUSTOMERS, payload: data.data });
        } else {
          dispatch(fetchError("Something went wrong, Please try again!"));
        }
      })
      .catch((error) => {
        dispatch(fetchError(error.message));
      });
  };
};

export const getCartItems = () => {
  return (dispatch: Dispatch<AppActions>) => {
    dispatch(fetchStart());
    //   jwtAxios.get('/api/cart/get')
    //     .then((data) => {
    //       if (data.status === 200) {
    dispatch(fetchSuccess());
    //         dispatch({type: SET_CART_ITEMS, payload: data.data});
    //       } else {
    //         dispatch({
    //           type: FETCH_ERROR,
    //           payload: 'Something went wrong, Please try again!',
    //         });
    //       }
    //     })
    //     .catch((error) => {
    //         dispatch(fetchError(error.message));
    //     });
  };
};

export const setViewType = (viewType: VIEW_TYPE) => {
  return (dispatch: Dispatch<AppActions>) => {
    dispatch({ type: SET_PRODUCT_VIEW_TYPE, payload: viewType });
  };
};

export const removeCartItem = (data: CartItems) => {
  return (dispatch: Dispatch<AppActions>) => {
    dispatch({ type: REMOVE_CART_ITEM, payload: data });
  };
};

export const updateCartItem = (data: CartItems) => {
  return (dispatch: Dispatch<AppActions>) => {
    dispatch({ type: UPDATE_CART_ITEM, payload: data });
  };
};

export const addItemToCart = (data: ProductData) => {
  return (dispatch: Dispatch<AppActions>) => {
    dispatch({ type: ADD_CART_ITEM, payload: data });
  };
};

export const setCurrentProduct = (product: ProductData) => {
  return (dispatch: Dispatch<AppActions>) => {
    dispatch({ type: SET_PRODUCT_DATA, payload: product });
  };
};

export const setFilters = (filters: any) => {
  return (dispatch: Dispatch<AppActions>) => {
    dispatch({ type: SET_FILTER_DATA, payload: filters });
  };
};
