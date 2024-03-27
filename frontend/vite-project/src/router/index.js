import { createRouter, createWebHistory } from 'vue-router';
import DataTable from '../views/DataTable.vue';
import Calculator from '../views/Calculator.vue'; // Import the Calculator component

const routes = [
  {
    path: '/data-table',
    name: 'DataTable',
    component: DataTable
  },
  {
    path: '/calculator', // Define the route for the calculator
    name: 'Calculator',
    component: Calculator
  },
  // Other routes...
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
