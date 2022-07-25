import {
  DashboardActionTypes,
  GET_ACADEMY_DATA,
  GET_ANALYTICS_DATA,
  GET_CRM_DATA,
  GET_CRYPTO_DATA,
  GET_ECOMMERCE_DATA,
  GET_HEALTH_CARE_DATA,
  GET_METRICS_DATA,
  GET_WIDGETS_DATA,
} from "../../types/actions/Dashboard.action";
import { Metrics } from "../../types/models/dashboards/Metrics";
import { Analytics } from "../../types/models/dashboards/Analytics";
import { CRM } from "../../types/models/dashboards/CRM";
import { Crypto } from "../../types/models/dashboards/Crypto";
import { Widgets } from "../../types/models/dashboards/Widgets";
import { Academy } from "../../types/models/dashboards/Academy";
import { Ecommerce } from "../../types/models/dashboards/Ecommerce";
import { HealthCare } from "../../types/models/dashboards/HealthCare";

const initialState: {
  academyData: Academy | null;
  ecommerceData: Ecommerce | null;
  healthCare: HealthCare | null;
  analyticsData: Analytics | null;
  crmData: CRM | null;
  cryptoData: Crypto | null;
  metricsData: Metrics | null;
  widgetsData: Widgets | null;
} = {
  ecommerceData: null,
  healthCare: null,
  academyData: null,
  analyticsData: null,
  crmData: null,
  cryptoData: null,
  metricsData: null,
  widgetsData: null,
};

const Dashboard = (state = initialState, action: DashboardActionTypes) => {
  switch (action.type) {
    case GET_ANALYTICS_DATA:
      return {
        ...state,
        analyticsData: action.payload,
      };
    case GET_ECOMMERCE_DATA:
      return {
        ...state,
        ecommerceData: action.payload,
      };
    case GET_HEALTH_CARE_DATA:
      return {
        ...state,
        healthCare: action.payload,
      };
    case GET_CRM_DATA:
      return {
        ...state,
        crmData: action.payload,
      };

    case GET_CRYPTO_DATA:
      return {
        ...state,
        cryptoData: action.payload,
      };

    case GET_ACADEMY_DATA:
      return {
        ...state,
        academyData: action.payload,
      };

    case GET_METRICS_DATA:
      return {
        ...state,
        metricsData: action.payload,
      };

    case GET_WIDGETS_DATA:
      return {
        ...state,
        widgetsData: action.payload,
      };

    default:
      return state;
  }
};
export default Dashboard;
