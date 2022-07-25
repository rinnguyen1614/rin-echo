import { Box, Icon, ListItemIcon, ListItemText, MenuItem } from "@mui/material";
import { ReactElement } from "react";
import { ListProps } from "react-admin";
import ListBase from "@app/auth/ListBase";
import { RequestMethodField } from "../../components/requestMethods/RequestMethodField";
import { TreeView } from "@app/tree/TreeView";
import MenuCreate from "./MenuCreate";
import MenuEdit from "./MenuEdit";
import { AppAnimate, AppGridContainer } from "@crema";
import FormattedMessage from "@crema/utility/FormattedMessage";
import { Fonts } from "@crema/shared/constants/AppEnums";

const MenuList = (props: ListProps): ReactElement => {
  return (
    <AppAnimate animation="transition.slideUpIn" delay={200}>
      <Box>
        <Box
          component="h2"
          sx={{
            color: "text.primary",
            fontWeight: Fonts.BOLD,
            mb: 6,
            fontSize: 16,
          }}
        >
          <FormattedMessage id="menus" />
        </Box>
        <Box>
          <ListBase
            resource="admin/menus/trees"
            filter={{ select: "id,slug,name,path,parent_id,icon" }}
            pagination={false}
            perPage={1000}
            {...props}
          >
            <TreeView
              resource="admin/menus"
              resourceTree="admin/menus/trees"
              create={<MenuCreate />}
              edit={<MenuEdit />}
              addRootButton={true}
              nodeText={(node) => (
                <MenuItem component="div">
                  {node.icon && (
                    <ListItemIcon>
                      <Icon fontSize="small">{node.icon}</Icon>
                    </ListItemIcon>
                  )}
                  <ListItemText
                    primary={node.name}
                    secondary={
                      (node.method || node.path) && (
                        <>
                          <RequestMethodField
                            record={node.method}
                            small={true}
                          />{" "}
                          <Box component="span" ml={1}>
                            {node.path}
                          </Box>
                        </>
                      )
                    }
                  />
                </MenuItem>
              )}
            ></TreeView>
          </ListBase>
        </Box>
      </Box>
    </AppAnimate>
  );
};

export default MenuList;
