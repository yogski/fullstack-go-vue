// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import About from '../views/About.vue';
import Orders from '../views/Orders.vue'

const routes = [
  {
    path: '/',
    component: Home,
  },
  {
    path: '/about',
    component: About,
  },
  {
    path: '/orders',
    component: Orders
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
