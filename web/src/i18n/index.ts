import polyglotI18nProvider from "ra-i18n-polyglot";
import enLang from "./en";

const locales: any = {
  en: enLang,
};

const i18nProvider = polyglotI18nProvider((locale) => {
  const def = locales.en;
  const found = locales[locale];
  if (found) {
    return found.messages;
  }

  return def;
}, "en");

export default i18nProvider;
