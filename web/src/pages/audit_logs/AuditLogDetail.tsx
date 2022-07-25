import { AppAnimate, AppCard } from "@crema";
import FormattedMessage from "@crema/utility/FormattedMessage";
import { Box, Divider, Grid, Paper } from "@mui/material";
import React from "react";
import { useRecordContext } from "react-admin";

const AuditLogDetail = (props) => {
  const record = useRecordContext(props);

  if (!record) {
    return null;
  }

  const informations = [
    { id: 1, title: "audit_logs.fields.username", desc: record.username },
    { id: 2, title: "audit_logs.fields.ip_address", desc: record.ip_address },
    { id: 3, title: "audit_logs.fields.device_name", desc: record.device_name },
    { id: 4, title: "audit_logs.fields.user_agent", desc: record.user_agent },
    { id: 5, title: "audit_logs.fields.location", desc: record.location },
  ];

  const action_infos = [
    {
      id: 1,
      title: "audit_logs.fields.operation_name",
      desc: record.operation_name,
    },
    {
      id: 2,
      title: "audit_logs.fields.operation_method",
      desc: record.operation_method,
    },
    { id: 3, title: "audit_logs.fields.request_id", desc: record.request_id },
    {
      id: 4,
      title: "audit_logs.fields.request_url",
      desc: record.request_method + " - " + record.request_url,
    },
    {
      id: 5,
      title: "audit_logs.fields.request_body",
      desc: record.request_body,
    },
    {
      id: 6,
      title: "audit_logs.fields.response_body",
      desc: record.response_body,
    },
  ];

  return (
    <AppAnimate animation="transition.slideUpIn" delay={200}>
      <AppCard>
        <Box
          component="h3"
          sx={{
            color: "text.primary",
            fontSize: 16,
            mt: 4,
            mb: 3,
          }}
        >
          <FormattedMessage id="audit_logs.user_info" />
        </Box>
        <Grid container spacing={3}>
          {informations.map((data) => (
            <React.Fragment key={data.id}>
              <Grid item xs={4}>
                <Box
                  sx={{
                    color: "text.secondary",
                  }}
                >
                  <FormattedMessage id={data.title} />
                </Box>
              </Grid>
              <Grid item xs={8}>
                <Box> {data.desc}</Box>
              </Grid>
            </React.Fragment>
          ))}
        </Grid>
        <Divider style={{ marginTop: 15, marginBottom: 15 }} />
        <Box
          component="h3"
          sx={{
            color: "text.primary",
            fontSize: 16,
            mt: 4,
            mb: 3,
          }}
        >
          <FormattedMessage id="audit_logs.action_info" />
        </Box>
        <Grid container spacing={3}>
          {action_infos.map((data) => (
            <React.Fragment key={data.id}>
              <Grid item xs={4}>
                <Box
                  sx={{
                    color: "text.secondary",
                  }}
                >
                  <FormattedMessage id={data.title} />
                </Box>
              </Grid>
              <Grid item xs={8}>
                <Box> {data.desc}</Box>
              </Grid>
            </React.Fragment>
          ))}
        </Grid>

        <Divider style={{ marginTop: 15, marginBottom: 15 }} />

        <Box
          component="h3"
          sx={{
            color: "text.primary",
            fontSize: 16,
            mt: 4,
            mb: 3,
          }}
        >
          <FormattedMessage id="audit_logs.other_info" />
        </Box>

        <Grid container spacing={3}>
          <Grid item xs={4}>
            <Box
              sx={{
                color: "text.secondary",
              }}
            >
              <FormattedMessage id={"audit_logs.fields.error"} />
            </Box>
          </Grid>
          <Grid item xs={8}>
            <Box> {record.error}</Box>
          </Grid>
        </Grid>

        <Grid container spacing={3}>
          <Grid item xs={4}>
            <Box
              sx={{
                color: "text.secondary",
              }}
            >
              <FormattedMessage id={"audit_logs.fields.remark"} />
            </Box>
          </Grid>
          <Grid item xs={8}>
            <Box> {record.remark}</Box>
          </Grid>
        </Grid>
      </AppCard>
    </AppAnimate>
  );
};

export default AuditLogDetail;
