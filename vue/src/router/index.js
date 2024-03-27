import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '../views/Dashboard'
import DataTable from '../views/DataTable.vue';
import Calculator from '../views/Calculator.vue';
import Map from '../views/Map.vue';

const routes = [
  { path: '/dashboard', component: Dashboard },
  { path: '/data-table', component: DataTable },
  { path: '/calculator', component: Calculator },
  { path: '/map', component: Map }
  // Add more routes as needed
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
