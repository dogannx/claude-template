import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

// Add routes here. Use lazy imports for non-critical pages:
//   { path: '/x', component: () => import('../views/X.vue') }
export const router = createRouter({
  history: createWebHistory(),
  routes: [{ path: '/', name: 'home', component: Home }],
})
