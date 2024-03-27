<template>
  <div class="dashboard">
    <h2>Welcome to Dashboard</h2>
    <div class="dashboard-content">
      <!-- Add your dashboard content here -->
      <div class="card">
        <h3>Recent Activity</h3>
        <ul>
          <li>Activity 1</li>
          <li>Activity 2</li>
          <li>Activity 3</li>
          <!-- Add more activities as needed -->
        </ul>
      </div>
      <div class="card">
        <h3>Statistics</h3>
        <p>Some statistics here...</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dashboard'
};
</script>

<style scoped>
.dashboard {
  padding: 20px;
}

.dashboard h2 {
  margin-bottom: 20px;
}

.dashboard-content {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  grid-gap: 20px;
}

.card {
  background-color: #f0f0f0;
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.card h3 {
  margin-bottom: 10px;
}

ul {
  list-style: none;
  padding: 0;
}

ul li {
  margin-bottom: 5px;
}
</style>
