<template>
  <header class="header">
    <div class="logo">
      <img src="logo.png" alt="Logo">
    </div>
    <nav class="nav">
      <ul>
        <li><router-link to="/dashboard">Dashboard</router-link></li>
        <li><router-link to="/map">Map</router-link></li>
        <li><router-link to="/calculator">Calculator</router-link></li>
        <!-- Add more menu items as needed -->
      </ul>
    </nav>
  </header>
</template>

<script>
export default {
  name: 'HeaderWeb'
};
</script>

<style scoped>
.header {
  background-color: #333;
  color: #fff;
  padding: 10px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo img {
  height: 40px;
  width: auto;
}

.nav ul {
  list-style: none;
  margin: 0;
  padding: 0;
}

.nav ul li {
  display: inline-block;
  margin-right: 20px;
}

.nav ul li:last-child {
  margin-right: 0;
}

.nav ul li a {
  color: #fff;
  text-decoration: none;
}

.nav ul li a:hover {
  text-decoration: underline;
}
</style>
