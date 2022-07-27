import { Box, ListItemText, MenuItem, Paper } from "@mui/material";
import { ReactElement } from "react";
import { ListProps, NullableBooleanInput, TextInput } from "react-admin";
import { RequestMethodField } from "../../components/requestMethods/RequestMethodField";
import { TreeView } from "@app/tree/TreeView";
import ResourceCreate from "./ResourceCreate";
import ResourceEdit from "./ResourceEdit";
import ListBase from "@app/auth/ListBase";
import AppsContainer from "@crema/core/AppsContainer";
import AppsContent from "@crema/core/AppsContainer/AppsContent";
import FormattedMessage from "@crema/utility/FormattedMessage";
import AppAnimate from "@crema/core/AppAnimate";
import { Fonts } from "@crema/shared/constants/AppEnums";

const resourceFilters = [
  <TextInput label="Search" source="q=name:like,slug:like" alwaysOn />,
  <NullableBooleanInput label="Default" source="is_default" />,
];

const ResourceList = (props: ListProps): ReactElement => {
  return (
    // <AppAnimate animation="transition.slideUpIn" delay={200}>
    <Box>
      <Box
        sx={{
          // component: 'h2',
          color: "text.primary",
          fontWeight: Fonts.BOLD,
          mb: 6,
          fontSize: 16,
        }}
      >
        {/* <FormattedMessage id="resources" /> */}
      </Box>

      <ListBase
        resource="admin/resources/trees"
        filters={resourceFilters}
        filter={{ select: "id,slug,name,object, action, parent_id" }}
        pagination={false}
        perPage={1000}
        {...props}
      >
        <TreeView
          component={Box}
          resource="admin/resources"
          resourceTree="admin/resources/trees"
          create={<ResourceCreate />}
          edit={<ResourceEdit />}
          addRootButton={true}
          nodeText={(node) => (
            <MenuItem component="div">
              <ListItemText
                primary={node.name}
                secondary={
                  (node.method || node.path) && (
                    <>
                      <RequestMethodField record={node.method} small={true} />{" "}
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
    // </AppAnimate>
  );
};

export default ResourceList;
