import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './bootstrap';
import './registerServiceWorker';
import i18n from './i18n'
import VeeValidate from 'vee-validate';
import moment from 'moment';

Vue.use(VeeValidate);

Vue.config.productionTip = false;
Vue.filter('formatDate', (value: any) => {
  if (value) {
    return moment(String(value)).format('MM/DD/YYYY hh:mm');
  }
});

Vue.filter('formatDate1', (value: any) => {
  if (value) {
    return moment(String(value)).format('MM-DD-YYYY hh:mm');
  }
});

new Vue({
  router,
  store,
  i18n,
  render: (h) => h(App)
}).$mount('#app');
