export interface AuthUser {
  id?: number;
  uid?: string;
  displayName?: string;
  email?: string;
  avatar_path?: string;
  token?: string;
  role?: string[] | string;
}
