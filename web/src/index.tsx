/** @jsxRuntime classic */
import "react-app-polyfill/ie11";
import "react-app-polyfill/stable";
import "proxy-polyfill";
// IE11 needs "jsxRuntime classic" for this initial file which means that "React" needs to be in scope
// https://github.com/facebook/create-react-app/issues/9906

import * as React from "react";
import ReactDOM from "react-dom/client";
import reportWebVitals from "./reportWebVitals";

import App from "./App";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
