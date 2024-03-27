// router/index.js

import { createRouter, createWebHistory } from 'vue-router';
import DataTable from '../views/DataTable.vue';
import Calculator from '../views/Calculator.vue';
import MapPreview from '../components/Map.vue'; // Import the MapPreview component

const routes = [
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
    path: '/map', // Define the route for the map
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
