import { createRouter, createWebHistory } from 'vue-router';
//import Dashboard from '../views/Dashboard.vue';
//import Dashboard from '../views/Dashboard.vue';

import Dashboard from '../views/DashboardView.vue'
import DataTable from '../views/DataTablePage.vue';
import Calculator from '../views/CalculatorView.vue';
import Map from '../views/MapView.vue';

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
