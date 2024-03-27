<template>
  <aside class="sidebar">
    <div class="logo">
      <img src="logo.png" alt="Logo">
    </div>
    <nav class="menu">
      <ul>
        <li><router-link to="/dashboard">Dashboard</router-link></li>
        <li><router-link to="/data-table">Data Table</router-link></li>
        <li><router-link to="/calculator">Calculator</router-link></li>
        <li><router-link to="/map">Map</router-link></li>
        <!-- Add more menu items as needed -->
      </ul>
    </nav>
  </aside>
</template>

<script>
export default {
  name: 'SidebarMenu'
};
</script>

<style scoped>
.sidebar {
  background-color: #333;
  color: #fff;
  width: 250px;
  height: 100vh;
  padding: 20px;
  display: flex;
  flex-direction: column;
}

.logo img {
  height: 40px;
  width: auto;
  margin-bottom: 20px;
}

.menu ul {
  list-style: none;
  padding: 0;
}

.menu ul li {
  margin-bottom: 15px;
}

.menu ul li a {
  color: #fff;
  text-decoration: none;
  font-size: 18px;
}

.menu ul li a:hover {
  text-decoration: underline;
}
</style>
