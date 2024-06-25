import httpClient from "../interceptor/interceptor";
import { jwtDecode } from 'jwt-decode';
import {useNavigate} from "react-router-dom";

class AuthService {

    async loginUser(email, password) {
      try {
        const response = await httpClient.post('http://localhost:8080/api/login',{
            email: email,
            password: password
        });
        await this.setToken(response.data['accessToken']);
       return response;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }

    async setToken(user) {
      localStorage.setItem('token', JSON.stringify(user));
    }

    getToken() {
        const token = localStorage.getItem('token');
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
            return null;
        }

        const decodedToken = jwtDecode(token);
        //todo change this
        // const roles = decodedToken.role;

        // return roles ? roles[0].name : null;
        return true
    }

    async logout() {
      localStorage.removeItem('token');
    }

    async register(user){
        try {
            console.log(user);
            const response = await httpClient.post('http://localhost:8080/api/register',{
                ...user
            });
            await this.setToken(response.data['accessToken']);
            return response;
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    }
      
    async getProfileData() {
        try {
          const response = await httpClient.get('http://localhost:8080/api');
          console.log(response);
          return response.data;
        } catch (error) {
          console.error('Error fetching data:', error);
        }
      }
    
}
  
  const authService = new AuthService();
  
  export default authService;