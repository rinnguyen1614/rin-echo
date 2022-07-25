import { AuthUser } from "../../types/models/AuthUser";

export const authRole = {
  admin: ["admin"],
  user: ["user", "admin"],
};

export enum RoutePermittedRole {
  Admin = "admin",
  User = "user",
}

export const defaultUser: AuthUser = {
  uid: "john-alex",
  displayName: "John Alex",
  email: "demo@example.com",
  token: "access-token",
  role: "user",
  avatar_path: "/assets/images/avatar/A11.jpg",
};
export const initialUrl = "/dashboards"; // this url will open after login
