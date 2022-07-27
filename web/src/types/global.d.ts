declare global {
  interface Window {
    restServer: any;
    __REDUX_DEVTOOLS_EXTENSION_COMPOSE__?: (traceOptions: object) => Function;
  }
}

declare module "yup" {
  interface StringSchema<
    TType extends Maybe<string> = string | undefined,
    TContext extends AnyObject = AnyObject,
    TOut extends TType = TType
  > {
    password(message?: string): this;
    phone(message?: string): this;
  }
}

export {};
