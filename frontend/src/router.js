import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    path: "/",
    redirect: () => {
      return { path: "/info" };
    }
  },
  {
    path: "/info",
    component: () => import('./views/InfoView.vue'),
  },
  {
    path: "/container",
    component: () => import('./views/ContainerView.vue'),
  },
  {
    path: "/image",
    component: () => import('./views/ImageView.vue'),
  },
  {
    path: "/log",
    component: () => import('./views/LogView.vue'),
  },
  {
    path: "/volume",
    component: () => import('./views/VolumeView.vue'),
  },
];

export const router = createRouter({ history: createWebHashHistory(), routes });
