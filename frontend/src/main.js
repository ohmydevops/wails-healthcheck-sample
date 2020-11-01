import "core-js/stable";
import "regenerator-runtime/runtime";
import Vue from "vue";
import { BootstrapVue, BootstrapVueIcons } from "bootstrap-vue";
import { BSkeletonTable } from "bootstrap-vue";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import VueSplash from "vue-splash";
import App from "./App.vue";

Vue.config.productionTip = false;
Vue.config.devtools = true;

Vue.use(BootstrapVue);
Vue.use(BootstrapVueIcons);
Vue.component("b-skeleton-table", BSkeletonTable);
Vue.use(VueSplash);

import * as Wails from "@wailsapp/runtime";

Wails.Init(() => {
  new Vue({
    render: (h) => h(App),
  }).$mount("#app");
});
