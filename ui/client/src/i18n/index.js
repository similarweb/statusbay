import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import en from './translations/en';

const resources = {
  en: {
    translation: en,
  },
};

i18n
  .use(initReactI18next)
  .init({
    resources,
    lng: 'en',

    keySeparator: false,

    interpolation: {
      escapeValue: false,
    },
  });
window.i18n = i18n;
export default i18n;
