import Cookies from "vue-cookies";
import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./App.vue";
import router from "@/router";
import "./tailwind.css";

const app = createApp(App);
app.use(createPinia());
app.use(router);
app.use(Cookies);
app.mount("#app");
