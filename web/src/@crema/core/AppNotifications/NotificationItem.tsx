import React from "react";
import ListItemAvatar from "@mui/material/ListItemAvatar";
import Avatar from "@mui/material/Avatar";
import { Box, ListItem, Typography } from "@mui/material";
import { Fonts } from "../../shared/constants/AppEnums";

interface NotificationItemProps {
  item: {
    image: string;
    name: string;
    message: string;
  };
}

const NotificationItem: React.FC<NotificationItemProps> = ({ item }) => {
  return (
    <ListItem
      sx={{
        padding: "8px 20px",
      }}
      className="item-hover"
    >
      <ListItemAvatar
        sx={{
          minWidth: 0,
          mr: 4,
        }}
      >
        <Avatar
          sx={{
            width: 48,
            height: 48,
          }}
          alt="Remy Sharp"
          src={item.image}
        />
      </ListItemAvatar>
      <Box
        sx={{
          fontSize: 14,
          color: (theme) => theme.palette.text.secondary,
        }}
      >
        <Typography>
          <Box
            component="span"
            sx={{
              fontSize: 14,
              fontWeight: Fonts.MEDIUM,
              mb: 0.5,
              color: (theme) => theme.palette.text.primary,
              mr: 1,
              display: "inline-block",
            }}
          >
            {item.name}
          </Box>
          {item.message}
        </Typography>
      </Box>
    </ListItem>
  );
};

export default NotificationItem;
