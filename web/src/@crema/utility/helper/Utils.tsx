import { IntlShape, useIntl } from "react-intl";
import { useMediaQuery, useTheme } from "@mui/material";
import { ReactNode } from "react";

export const useWidth = () => {
  const theme = useTheme();
  const keys = [...theme.breakpoints.keys].reverse();
  // @ts-ignore
  return (
    keys.reduce((output, key: any) => {
      // eslint-disable-next-line react-hooks/rules-of-hooks
      const matches = useMediaQuery(theme.breakpoints.up(key));
      return !output && matches ? key : output;
    }, null) || "xs"
  );
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

// 'intl' service singleton reference
let intl: IntlShape;

export const IntlGlobalProvider = ({
  children,
}: {
  children: ReactNode;
}): any => {
  intl = useIntl();
  // Keep the 'intl' service reference
  return children;
};

export const appIntl = () => {
  return intl;
};
