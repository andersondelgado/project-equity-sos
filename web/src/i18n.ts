import Vue from 'vue'
import VueI18n, { LocaleMessages } from 'vue-i18n'
import VeeValidate from 'vee-validate';

// locale= process.env.VUE_APP_I18N_LOCALE || 'es';

// import validationMessages from 'vee-validate/dist/locale/es';
Vue.use(VueI18n)

// const i18n = new VueI18n();
// i18n.locale = "es";

// Vue.use(VeeValidate, {
//   i18nRootKey: 'validations',
//   i18n
//   ,
//   dictionary: {
//     es: validationMessages
//   }
// });


function loadLocaleMessages(): LocaleMessages {
  // const locales = require.context('locales', true, /[A-Za-z0-9-_,\s]+\.json$/i)
  const locales = require.context('./locales', true, /[A-Za-z0-9-_,\s]+\.json$/i)
  const messages: LocaleMessages = {}
  locales.keys().forEach(key => {
    const matched = key.match(/([A-Za-z0-9-_]+)\./i)
    if (matched && matched.length > 1) {
      const locale = matched[1]
      messages[locale] = locales(key)
    }
  })
  return messages
}

export default new VueI18n({
  locale: process.env.VUE_APP_I18N_LOCALE || 'es',
  fallbackLocale: process.env.VUE_APP_I18N_FALLBACK_LOCALE || 'en',
  messages: loadLocaleMessages()
})
