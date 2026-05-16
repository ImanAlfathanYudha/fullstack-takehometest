import axios from 'axios';

// Create an instance with your backend's base URL
const api = axios.create({
  baseURL: 'http://localhost:8080', //make it global in case backend URL changes
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to add the token to headers.
api.interceptors.request.use(
  (config) => {
    // Check token in local storage
    const token = localStorage.getItem('token');
    
    // Check the token, add it to header
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default api;
