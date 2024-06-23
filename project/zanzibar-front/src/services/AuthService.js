import { jwtDecode } from 'jwt-decode';

class AuthService {

    async loginUser(user, password) {
      try {
        await this.setToken(user);
        return true;
      } catch (error) {
        console.error('Error fetching data:', error);
        return false;
      }
    }

    async setToken(user) {
      localStorage.setItem('user', JSON.stringify(user));
    }

    getToken() {
        const token = localStorage.getItem('user');
        let user = null;
        try {
            user = JSON.parse(token);
        } catch (error) {
            console.log('Error parsing token:', error);
        }
        return user;
    }

    validateUser() {
        const token = this.getToken();
        if (!token) {
            return false;
        }

        const decodedToken = jwtDecode(token);
        return true;
    }

    async logout() {
      localStorage.removeItem('user');
    }
    
}
  
  const authService = new AuthService();
  
  export default authService;