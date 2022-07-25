import React from "react";
import List from "@mui/material/List";
import { RouterConfig } from "../../../../types/models/RouterConfig";
import NavVerticalGroup from "./VerticalNavGroup";
import VerticalCollapse from "./VerticalCollapse";
import VerticalItem from "./VerticalItem";
import useMenus from "@app/auth/useMenus";

const VerticalNav = () => {
  const { menus, isLoading } = useMenus();
  return (
    !isLoading && (
      <List
        sx={{
          position: "relative",
          padding: 0,
        }}
        component="div"
      >
        {menus?.map((item: RouterConfig) => (
          <React.Fragment key={item.id}>
            {item.type === "group" && (
              <NavVerticalGroup item={item} level={0} />
            )}

            {item.type === "collapse" && (
              <VerticalCollapse item={item} level={0} />
            )}

            {item.type === "item" && <VerticalItem item={item} level={0} />}
          </React.Fragment>
        ))}
      </List>
    )
  );
};

export default VerticalNav;
