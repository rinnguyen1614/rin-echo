import React from "react";

export const useUserMenu = () => React.useContext(UserMenuContext);

export const UserMenuContext =
  React.createContext<UserMenuContextValue>(undefined);

export type UserMenuContextValue = {
  /**
   * Closes the user menu
   * @see UserMenu
   */
  onClose: () => void;
};

export const UserMenuContextProvider = ({ children, value }) => (
  <UserMenuContext.Provider value={value}>{children}</UserMenuContext.Provider>
);

export type UserMenuContextProviderProps = {
  children: React.ReactNode;
  value: UserMenuContextValue;
};
