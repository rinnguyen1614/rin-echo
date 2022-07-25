import ShowBase from "@app/auth/Show";

import React from "react";
import AuditLogDetail from "./AuditLogDetail";

const AuditLogShow = (props) => {
  return (
    <ShowBase {...props}>
      <AuditLogDetail />
    </ShowBase>
  );
};

export default AuditLogShow;
