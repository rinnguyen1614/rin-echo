import { ReactNode } from "react";

export interface RouterConfig {
  id: string;
  title: string;
  slug: string;
  icon?: string | ReactNode;
  type: "item" | "group" | "collapse" | "divider";
  children?: RouterConfig[];
  //   permittedRole?: RoutePermittedRole;
  color?: string;
  path?: string;
  exact?: boolean;
  count?: number;
}
