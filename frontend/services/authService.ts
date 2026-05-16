import api from "./api";

export interface UserResponse {
    email: string;
    role: string;
    token: string;
}

export const login = async (email: string, password: string): Promise<UserResponse> => {
  try {
    const response = await api.post('/dashboard/v1/auth/login', { email, password });
    const userData: UserResponse = response.data;
    localStorage.setItem('token', userData.token);
    localStorage.setItem('role', userData.role);
    localStorage.setItem('email', userData.email);
    return userData;
  } catch (error: any) {
    const errorMessage = error.response?.data?.message || 'Something went wrong';
    throw new Error(errorMessage);
  }
};


export const getToken = (): string | null => {
    if (typeof window !== 'undefined') {
        return localStorage.getItem('token');
    }
    return null;
};

export const getRole = (): string | null => {
    if (typeof window !== 'undefined') {
        return localStorage.getItem('role');
    }
    return null;
};

export const isAuthenticated = (): boolean => {
    const token = getToken();
    return !!token;
};

export const logout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('role');
    localStorage.removeItem('email');
    window.location.href = '/login';
};