import React, { ReactNode } from "react";
import {
  alpha,
  Box,
  Card,
  CardActions,
  CardContent,
  CardMedia,
} from "@mui/material";
import { Fonts } from "../../shared/constants/AppEnums";

interface AppCardProps {
  sxStyle?: any;
  contentStyle?: any;
  cardMedia?: string;
  sxCardMedia?: any;
  cardMediaAction?: any;
  sxCardMediaAction?: any;
  footer?: any;
  footerStyle?: any;
  children: ReactNode;

  [x: string]: any;
}

const AppCardMedia: React.FC<AppCardProps> = ({
  children,
  sxStyle,
  contentStyle,
  cardMedia,
  sxCardMedia,
  cardMediaAction,
  sxCardMediaAction,
  footer,
  footerStyle,
  ...rest
}) => {
  return (
    <Card
      sx={{ display: "flex", flexDirection: "column", ...sxStyle }}
      {...rest}
    >
      <Box sx={{ position: "relative" }}>
        <CardMedia
          component="img"
          image={cardMedia}
          alt="card media"
          sx={{
            minHeight: 140,
            width: "100%",
            maxHeight: 500,
            ...sxCardMedia,
          }}
        />
        {cardMediaAction ? (
          <Box
            component="span"
            sx={{
              backgroundColor: (theme) =>
                alpha(theme.palette.text.primary, 0.7),
              color: (theme) => theme.palette.background.paper,
              pt: 0.75,
              pb: 1.75,
              px: 4,
              borderRadius: 7.5,
              position: "absolute",
              right: 15,
              top: 15,
              zIndex: 1,
              fontSize: 12,
              fontWeight: Fonts.BOLD,
              ...sxCardMediaAction,
            }}
          >
            {cardMediaAction}
          </Box>
        ) : null}
      </Box>

      <CardContent
        style={{
          paddingLeft: 20,
          paddingRight: 20,
          height: "100%",
          ...contentStyle,
        }}
      >
        {children}
      </CardContent>
      {footer ? (
        <CardActions
          style={{
            paddingLeft: 20,
            paddingRight: 20,
            paddingBottom: 16,
            ...footerStyle,
          }}
        >
          {footer}
        </CardActions>
      ) : null}
    </Card>
  );
};

export default AppCardMedia;
