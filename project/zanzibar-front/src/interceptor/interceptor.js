import axios from 'axios';
import authService from '../services/AuthService';

const httpClient = axios.create();

httpClient.interceptors.request.use(
    async (config) => {
        const skipUrls = ['/api/register', '/api/login'];

        // Only add the token if the URL is not in the skipUrls array
        if (!skipUrls.includes(config.url)) {
            const token = await authService.getToken();
            config.headers['Authorization'] = `Bearer ${token}`;
        }

        // Add CORS headers
        config.headers['Access-Control-Allow-Origin'] = '*'; // Allow all origins
        config.headers['Access-Control-Allow-Methods'] = 'GET, POST, PUT, DELETE, OPTIONS';
        config.headers['Access-Control-Allow-Headers'] = 'Origin, X-Requested-With, Content-Type, Accept, Authorization';
        config.headers['Access-Control-Allow-Credentials'] = 'true';
        return config;
    },
    (error) => {
        console.error('Error setting Authorization header:', error);
        return Promise.reject(error);
    }
);

httpClient.interceptors.response.use((response) => {
    return response;
}, function (error) {
    if (error.response && error.response.status === 401 || error.response.status === 403) {
        const event = new CustomEvent('unauthorized');
        console.error('Unauthorized event dispatched');
        window.dispatchEvent(event);
    }
    return Promise.reject(error);
});

export default httpClient;