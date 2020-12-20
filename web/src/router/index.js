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
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
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
