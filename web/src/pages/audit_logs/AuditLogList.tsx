import { Chip } from "@mui/material";
import { makeStyles } from "@mui/styles";
import classnames from "classnames";
import { FC, memo, ReactElement } from "react";
import {
  BooleanField,
  Datagrid,
  DateField,
  DateTimeInput,
  ListProps,
  TextField,
  TextInput,
  useTranslate,
  FunctionField,
} from "react-admin";
import List from "@app/auth/List";
import AuditLogShow from "./AuditLogShow";
import { StatusField } from "./StatusField";
import AppsContent from "@crema/core/AppsContainer/AppsContent";
import AppsContainer from "@crema/core/AppsContainer";
import AppAnimate from "@crema/core/AppAnimate";

const auditLogFilters = [
  <TextInput
    label="Username"
    source="q=username:like"
    alwaysOn
    variant="outlined"
  />,
  <DateTimeInput
    source="start_time:>="
    label="Start time from"
    alwaysOn
    variant="outlined"
  />,
  <DateTimeInput
    source="start_time:<="
    label="Start time to"
    alwaysOn
    variant="outlined"
  />,
];

const AuditLogList = (props: ListProps): ReactElement => {
  const translate = useTranslate();

  return (
    <AppAnimate animation="transition.slideUpIn" delay={200}>
      <AppsContainer title="Audit logs" fullView>
        <AppsContent>
          <List
            {...props}
            filters={auditLogFilters}
            filter={{
              select:
                "id,created_at,username,request_method,request_url,start_time, ip_address,status_code",
            }}
            perPage={10}
          >
            <Datagrid
              optimized
              rowClick="expand"
              size="medium"
              expand={<AuditLogShow />}
              bulkActionButtons={false}
              sx={{
                "& .column-status_code": { textAlign: "center" },
                "& .column-request_url": {
                  maxWidth: 250,
                  overflow: "hidden !important",
                  textOverflow: "ellipsis",
                },
              }}
            >
              <TextField source="username" />
              <TextField source="ip_address" />
              <FunctionField
                source="request_url"
                label="Request URL"
                render={({ request_url, request_method }) =>
                  `${request_method} - ${request_url}`
                }
              />
              <StatusField source="status_code" />
              <DateField
                source="created_at"
                locales="us-US"
                showTime={true}
                label={translate("model.start_time")}
              />
              <DateField
                source="created_at"
                locales="us-US"
                showTime={true}
                label={translate("model.created_at")}
              />
            </Datagrid>
          </List>
        </AppsContent>
      </AppsContainer>
    </AppAnimate>
  );
};

export default AuditLogList;
