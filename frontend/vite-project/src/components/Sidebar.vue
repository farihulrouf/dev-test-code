<template>
    <aside class="sidebar bg-gray-800 text-white w-64 flex-shrink-0">
      <div class="p-4">
        <h1 class="text-2xl font-semibold">Sidebar</h1>
        <ul class="mt-4">
          <li><router-link to="/data-table" class="block py-2 px-4 hover:bg-gray-700 transition duration-200">Data Table</router-link></li>
          <li><router-link to="/calculator" class="block py-2 px-4 hover:bg-gray-700 transition duration-200">Calculator</router-link></li>
          <li><router-link to="/map" class="block py-2 px-4 hover:bg-gray-700 transition duration-200">Map</router-link></li>
        </ul>
      </div>
    </aside>
  </template>
  
  <script>
  export default {
    name: 'Sidebar'
  };
  </script>
  
  <style>
  /* Add custom styles here if needed */
  </style>
  