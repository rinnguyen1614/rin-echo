declare global {
  interface Window {
    restServer: any;
    __REDUX_DEVTOOLS_EXTENSION_COMPOSE__?: (traceOptions: object) => Function;
  }
}

export {};
