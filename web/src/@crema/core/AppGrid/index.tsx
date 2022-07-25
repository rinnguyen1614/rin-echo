import React, { CSSProperties, ReactNode } from "react";
import GridView from "./GridView";
import GridFooter from "./GridFooter";

interface AppCardProps {
  loading?: boolean;
  border?: boolean;
  footerProps?: {
    loading: boolean;
    footerText: string;
  };
  containerStyle?: CSSProperties;
  ListEmptyComponent?: ReactNode;
  ListFooterComponent?: ReactNode;
  data: any[];
  onEndReached?: () => void;
  renderRow: (item: any, index: number) => ReactNode;

  [x: string]: any;
}

const AppGrid: React.FC<AppCardProps> = ({ footerProps, ...rest }) => {
  return (
    <GridView
      {...rest}
      ListFooterComponent={
        footerProps ? (
          <GridFooter
            loading={footerProps.loading}
            footerText={footerProps.footerText}
          />
        ) : null
      }
    />
  );
};

export default AppGrid;
