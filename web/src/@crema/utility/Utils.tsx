import moment from "moment";
import { authRole } from "../shared/constants/AppConst";
import { Breakpoint } from "@mui/system";

export const getBreakPointsValue = (valueSet: any, breakpoint: Breakpoint) => {
  if (typeof valueSet === "number") return valueSet;
  switch (breakpoint) {
    case "xs":
      return valueSet.xs;
    case "sm":
      return valueSet.sm || valueSet.xs;
    case "md":
      return valueSet.md || valueSet.sm || valueSet.xs;
    case "lg":
      return valueSet.lg || valueSet.md || valueSet.sm || valueSet.xs;
    default:
      return (
        valueSet.xl || valueSet.lg || valueSet.md || valueSet.sm || valueSet.xs
      );
  }
};

export const getFileSize = (bytes: number) => {
  if (bytes === 0) return "0 Bytes";
  let k = 1024,
    dm = 2,
    sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"],
    i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + " " + sizes[i];
};

export const multiPropsFilter = (
  products: any[],
  filters: any,
  stringKey = "title"
) => {
  const filterKeys = Object.keys(filters);
  return products.filter((product) => {
    return filterKeys.every((key) => {
      if (!filters[key].length) return true;
      // Loops again if product[key] is an array (for material attribute).
      if (Array.isArray(product[key])) {
        return product[key].some((keyEle: any) =>
          filters[key].includes(keyEle)
        );
      }
      console.log("key", key, filters[key], product[key]);
      if (key === stringKey) {
        return product[key].toLowerCase().includes(filters[key].toLowerCase());
      }
      return filters[key].includes(product[key]);
    });
  });
};

export const getCustomDateTime = (
  value = 0,
  unit = "days",
  format = "YYYY-MM-DD"
) => {
  if (value === 0) {
    return moment().format(format);
  } else {
    // @ts-ignore
    return moment().add(value, unit).format(format);
  }
};

export const timeFromNow = (date: any) => {
  const timestamp = moment(date).format("X");
  const newDate = moment.unix(Number(timestamp));
  return moment(newDate).fromNow();
};

export const checkPermission = (routeAuth: any, userRole: any) => {
  if (routeAuth === null || routeAuth === undefined) {
    return true;
  }

  if (userRole && Array.isArray(userRole)) {
    return routeAuth.some((r: any) => userRole.indexOf(r) >= 0);
  }

  if (routeAuth.length === 0) {
    return !userRole || userRole.length === 0;
  }
  if (userRole && Array.isArray(userRole) && Array.isArray(routeAuth)) {
    return routeAuth.some((r) => userRole.indexOf(r) >= 0);
  }
  return routeAuth.indexOf(userRole) >= 0;
};

export const generateUniqueID = () => {
  return `v1-${Date.now()}-${Math.floor(Math.random() * (9e12 - 1)) + 1e12}`;
};

export const getUserFromAuth0 = (user: any) => {
  if (user)
    return {
      id: 1,
      uid: user.sub,
      displayName: user.name,
      email: user.email,
      avatar_path: user.picture,
      role: authRole.user,
    };
  return user;
};

export const getUserFromFirebase = (user: any) => {
  console.log("user: ", user);
  if (user)
    return {
      id: 1,
      uid: user.uid,
      displayName: user.displayName ? user.displayName : "Crema User",
      email: user.email,
      avatar_path: user.avatar_path,
      role: authRole.user,
    };
  return user;
};
export const getUserFromAWS = (user: any) => {
  if (user)
    return {
      id: 1,
      uid: user.username,
      displayName: user.attributes.name ? user.attributes.name : "Crema User",
      email: user.attributes.email,
      avatar_path: user.avatar_path,
      role: authRole.user,
    };
  return user;
};

export const getUserFromJwtAuth = (user: any) => {
  if (user)
    return {
      id: 1,
      uid: user._id,
      displayName: user.name,
      email: user.email,
      avatar_path: user.avatar,
      role: authRole.user,
    };
  return user;
};
