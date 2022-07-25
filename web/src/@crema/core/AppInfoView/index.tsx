import React from "react";

import { useSelector } from "react-redux";
import AppMessageView from "@crema/core/AppMessageView";
import AppLoader from "@crema/core/AppLoader";
import { AppState } from "../../redux/store";

const AppInfoView = () => {
  // @ts-ignore
  const { error, loading, message } = useSelector<AppState, AppState["common"]>(
    ({ common }) => common
  );

  const showMessage = () => {
    return <AppMessageView variant="success" message={message.toString()} />;
  };

  const showError = () => {
    return <AppMessageView variant="error" message={error.toString()} />;
  };
  console.log("loader", loading);
  return (
    <>
      {loading && <AppLoader />}

      {message && showMessage()}
      {error && showError()}
    </>
  );
};

export default AppInfoView;
