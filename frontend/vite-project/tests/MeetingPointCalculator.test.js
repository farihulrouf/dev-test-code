// MeetingPointCalculator.js

// Calculate meeting point and time required for two riders to meet
export function calculateMeetingPoint(routeA, routeB) {
    // Calculate average latitude and longitude for each route
    const avgLatitudeA = routeA.reduce((sum, point) => sum + point.latitude, 0) / routeA.length;
    const avgLongitudeA = routeA.reduce((sum, point) => sum + point.longitude, 0) / routeA.length;
    const avgLatitudeB = routeB.reduce((sum, point) => sum + point.latitude, 0) / routeB.length;
    const avgLongitudeB = routeB.reduce((sum, point) => sum + point.longitude, 0) / routeB.length;
  
    // Calculate meeting point coordinates (average of average latitudes and longitudes)
    const meetingPoint = {
      latitude: (avgLatitudeA + avgLatitudeB) / 2,
      longitude: (avgLongitudeA + avgLongitudeB) / 2
    };
  
    // Calculate time required for riders to meet (assuming same speed)
    const distanceA = calculateDistance(routeA);
    const distanceB = calculateDistance(routeB);
    const totalTime = (distanceA + distanceB) / 2; // Average of total distances
  
    return { meetingPoint, timeRequired: totalTime };
  }
  
  // Helper function to calculate distance between consecutive points in a route
  function calculateDistance(route) {
    let totalDistance = 0;
    for (let i = 1; i < route.length; i++) {
      const lat1 = route[i - 1].latitude;
      const lon1 = route[i - 1].longitude;
      const lat2 = route[i].latitude;
      const lon2 = route[i].longitude;
      totalDistance += calculateDistanceBetweenPoints(lat1, lon1, lat2, lon2);
    }
    return totalDistance;
  }
  
  // Helper function to calculate distance between two points using Haversine formula
  function calculateDistanceBetweenPoints(lat1, lon1, lat2, lon2) {
    const R = 6371; // Earth's radius in kilometers
    const dLat = toRadians(lat2 - lat1);
    const dLon = toRadians(lon2 - lon1);
    const a =
      Math.sin(dLat / 2) * Math.sin(dLat / 2) +
      Math.cos(toRadians(lat1)) * Math.cos(toRadians(lat2)) *
      Math.sin(dLon / 2) * Math.sin(dLon / 2);
    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
    const distance = R * c; // Distance in kilometers
    return distance;
  }
  
  // Helper function to convert degrees to radians
  function toRadians(degrees) {
    return degrees * (Math.PI / 180);
  }
  