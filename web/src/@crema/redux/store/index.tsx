import { applyMiddleware, combineReducers, compose, createStore } from "redux";
import thunk from "redux-thunk";
import reducers from "../reducers";

const createBrowserHistory = require("history").createBrowserHistory;
const history = createBrowserHistory();
const rootReducer = combineReducers({
  ...reducers,
});

export type AppState = ReturnType<typeof rootReducer>;

export default function configureStore(initialState?: AppState) {
  const enhancers: any[] = [];

  return createStore(
    rootReducer,
    initialState,
    compose(applyMiddleware(thunk), ...enhancers)
  );
}

export { history };
