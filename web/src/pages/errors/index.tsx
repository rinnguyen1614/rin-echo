import React from "react";

const Error401 = React.lazy(() => import("./Error401"));
const Error403 = React.lazy(() => import("./Error403"));
const Error404 = React.lazy(() => import("./Error404"));
const Error500 = React.lazy(() => import("./Error500"));
const Error503 = React.lazy(() => import("./Error503"));
const Maintenance = React.lazy(() => import("./Maintenance"));

export const errorPagesConfigs = [
  {
    path: "/error-pages/error-401",
    element: <Error401 />,
  },
  {
    path: "/error-pages/error-403",
    element: <Error403 />,
  },
  {
    path: "/error-pages/error-404",
    element: <Error404 />,
  },
  {
    path: "/error-pages/error-500",
    element: <Error500 />,
  },
  {
    path: "/error-pages/error-503",
    element: <Error503 />,
  },
  {
    path: "/error-pages/maintenance",
    element: <Maintenance />,
  },
];
