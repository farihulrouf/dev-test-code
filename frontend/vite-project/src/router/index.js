import { createRouter, createWebHistory } from 'vue-router';
import DataTable from '../views/DataTable.vue';
import Calculator from '../views/Calculator.vue';
import MapPreview from '../components/Map.vue';
import Login from '../views/LoginPage.vue'; // Import the Login component

const routes = [
  {
    path: '/login', // Define the route for the login page
    name: 'Login',
    component: Login
  },
  {
    path: '/data-table',
    name: 'DataTable',
    component: DataTable
  },
  {
    path: '/calculator',
    name: 'Calculator',
    component: Calculator
  },
  {
    path: '/map',
    name: 'MapPreview',
    component: MapPreview
  },
  // Other routes...
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
