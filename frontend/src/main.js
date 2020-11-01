import "core-js/stable";
import "regenerator-runtime/runtime";
import Vue from "vue";
import { BootstrapVue, BootstrapVueIcons } from "bootstrap-vue";
import { BSkeletonTable } from "bootstrap-vue";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import VueSplash from "vue-splash";
import VueRouter from "vue-router";
import VueAutoscroll from "@codekraft-studio/vue-autoscroll";
import App from "./App.vue";

Vue.config.productionTip = false;
Vue.config.devtools = true;

Vue.use(BootstrapVue);
Vue.use(BootstrapVueIcons);
Vue.component("b-skeleton-table", BSkeletonTable);
Vue.use(VueSplash);
Vue.use(VueRouter);
Vue.use(VueAutoscroll);

import * as Wails from "@wailsapp/runtime";
import Table from "./components/Table.vue";
import Ping from "./components/Ping.vue";

const routes = [
  { path: "/table", component: Table },
  { path: "/ping", component: Ping },
];

const router = new VueRouter({
  routes,
});

Wails.Init(() => {
  new Vue({
    render: (h) => h(App),
    router,
  }).$mount("#app");
});
