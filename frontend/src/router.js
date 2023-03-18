import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    path: "/",
    redirect: () => {
      return { path: "/container" };
    }
  },
  {
    path: "/container",
    component: () => import('./views/ContainerView.vue'),
  },
  {
    path: "/image",
    component: () => import('./views/ImageView.vue'),
  },
];

export const router = createRouter({ history: createWebHashHistory(), routes });
