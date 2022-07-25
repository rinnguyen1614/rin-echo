import React, { CSSProperties, ReactNode, useEffect, useState } from "react";
import { useBottomScrollListener } from "react-bottom-scroll-listener";
import { Box, Theme, useTheme } from "@mui/material";
import AppAnimateGroup from "../AppAnimateGroup";
import { useWidth } from "../../utility/helper/Utils";
import { SxProps } from "@mui/system";

interface GridViewProps {
  sx?: SxProps<Theme>;
  width?: string;
  responsive?: {
    xs: number;
    sm: number;
    md: number;
    lg: number;
    xl: number;
  };
  itemPadding?: number;
  renderRow: (item: any, index: number) => ReactNode;
  border?: boolean;
  column?: number;
  animation?: string;
  containerStyle?: CSSProperties;
  ListEmptyComponent?: ReactNode;
  ListFooterComponent?: ReactNode;
  data?: any[];
  onEndReached?: () => void;

  [x: string]: any;
}

const getEmptyContainer = (ListEmptyComponent: any) => {
  if (ListEmptyComponent)
    return React.isValidElement(ListEmptyComponent) ? (
      ListEmptyComponent
    ) : (
      <ListEmptyComponent />
    );
  return null;
};

const getFooterContainer = (ListFooterComponent: any) => {
  if (ListFooterComponent)
    return React.isValidElement(ListFooterComponent) ? (
      ListFooterComponent
    ) : (
      <ListFooterComponent />
    );
  return null;
};

const GridView: React.FC<GridViewProps> = ({
  sx,
  column = 3,
  responsive,
  itemPadding = 12,
  animation = "transition.expandIn",
  renderRow,
  onEndReached,
  data = [],
  containerStyle,
  border = false,
  ListFooterComponent,
  ListEmptyComponent,
}) => {
  const theme = useTheme();
  const width = useWidth();
  const borderStyle: CSSProperties = {
    border: `1px solid ${theme.palette.divider}`,
    backgroundColor: theme.palette.background.paper,
    borderRadius: 4,
    overflow: "hidden",
  };

  const [displayColumn, setColumn] = useState<number>(column);
  if (!onEndReached) {
    onEndReached = () => {};
  }

  useEffect(() => {
    setColumn(column);
  }, [column]);

  useEffect(() => {
    const getColumnCount = () => {
      if (responsive) {
        if (width === "xs") {
          return responsive.xs || column;
        } else if (width === "sm") {
          return responsive.sm || responsive.xs || column;
        } else if (width === "md") {
          return responsive.md || responsive.sm || responsive.xs || column;
        } else if (width === "lg") {
          return (
            responsive.lg ||
            responsive.md ||
            responsive.sm ||
            responsive.xs ||
            column
          );
        } else if (width === "xl") {
          return (
            responsive.xl ||
            responsive.lg ||
            responsive.md ||
            responsive.sm ||
            responsive.xs ||
            column
          );
        }
      } else {
        return column;
      }
    };
    setColumn(getColumnCount() as number);
  }, [width, column, responsive]);

  let style = containerStyle;
  if (border) {
    style = { ...style, ...borderStyle };
  }
  useBottomScrollListener(onEndReached, { debounce: 200 });
  return (
    <Box
      sx={{
        width: "100%",
        ...sx,
      }}
    >
      <AppAnimateGroup
        enter={{
          animation,
        }}
        style={{
          display: "flex",
          flexDirection: "row",
          flexWrap: "wrap",
          margin: -itemPadding,
          ...style,
        }}
      >
        {data.length > 0 ? (
          data.map((item, index) => (
            <Box
              style={{
                flexGrow: 0,
                maxWidth: `${100 / displayColumn}%`,
                flexBasis: `${100 / displayColumn}%`,
                padding: itemPadding,
                boxSizing: "border-box",
              }}
              key={"grid-" + index}
            >
              {renderRow(item, index)}
            </Box>
          ))
        ) : (
          <div />
        )}
      </AppAnimateGroup>
      {data.length === 0 ? getEmptyContainer(ListEmptyComponent) : null}
      {getFooterContainer(ListFooterComponent)}
    </Box>
  );
};

export default GridView;
