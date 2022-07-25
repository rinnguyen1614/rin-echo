import englishMessages from "ra-language-english";
import { app } from "./app";
import { layout } from "./layout";

const lang = {
  messages: {
    ...englishMessages,
    ...layout,
    ...app,
  },
  locale: "en-US",
};

export default lang;
