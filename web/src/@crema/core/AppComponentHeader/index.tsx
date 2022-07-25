import React from "react";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import LinkIcon from "@mui/icons-material/Link";
import Box from "@mui/material/Box";
import AppAnimate from "../AppAnimate";
import { Fonts } from "../../shared/constants/AppEnums";

interface AppComponentHeaderProps {
  title: string;
  description?: string;
  refUrl?: string;
}

const AppComponentHeader: React.FC<AppComponentHeaderProps> = ({
  title,
  description,
  refUrl,
}) => {
  return (
    <AppAnimate animation="transition.slideDownIn" delay={300}>
      <Box
        sx={{
          display: "flex",
          flexDirection: { xs: "column", sm: "row" },
          justifyContent: { sm: "space-between" },
          pb: 4,
        }}
      >
        <Box sx={{ mb: 3, pr: { sm: 3 }, flex: { sm: 1 } }}>
          <Typography
            component="h3"
            sx={{
              color: (theme) => theme.palette.text.primary,
              fontWeight: Fonts.MEDIUM,
              fontSize: { xs: 18, sm: 20 },
            }}
          >
            {title}
          </Typography>
          {description ? (
            <Typography
              variant="h6"
              sx={{
                fontSize: 15,
                fontWeight: Fonts.REGULAR,
                color: (theme) => theme.palette.text.secondary,
              }}
            >
              {description}
            </Typography>
          ) : null}
        </Box>
        {refUrl ? (
          <Box sx={{ height: 40 }}>
            <Button
              variant="outlined"
              color="primary"
              href={refUrl}
              target="_blank"
            >
              Reference <LinkIcon sx={{ pl: 1 }} />
            </Button>
          </Box>
        ) : null}
      </Box>
    </AppAnimate>
  );
};

export default AppComponentHeader;
