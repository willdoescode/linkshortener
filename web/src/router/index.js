import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Direct from "@/views/Direct";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Create",
    component: Home
  },
  {
    path: "/:id",
    name: "",
    component: Direct
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
