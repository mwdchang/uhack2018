import Vue from 'vue'
import App from './App.vue'
import store from './store'

Vue.config.productionTip = false

Vue.filter('chem', function (value) {
  if (value === 'ARSENIC') return 'Ar';
  if (value === 'CHROMIUM') return 'Cr';
  if (value === 'LEAD') return 'Pb';
  return '';
});

new Vue({
  store,
  render: h => h(App)
}).$mount('#app')
