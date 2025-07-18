/* eslint-disable react-refresh/only-export-components */
import React, { createContext, useState, useEffect } from 'react';
import { get } from '../utils/api';

export interface AuthContextType {
  isAuthenticated: boolean;
  setIsAuthenticated: (value: boolean) => void;
  checkAuthStatus: () => Promise<void>;
  refreshAuthStatus: () => void;
  login: (token?: string) => void;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextType | null>(null);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  const checkAuthStatus = async () => {
    try {
      const response = await get('/v1/auth/check', true);
      
      if (response.ok) {
        setIsAuthenticated(true);
      } else if (response.status === 401) {
        // 401は認証されていない状態なので、正常な動作
        setIsAuthenticated(false);
      } else {
        // その他のエラーは一時的な問題の可能性があるため、現在の状態を保持
        setIsAuthenticated(false);
      }
    } catch {
      // ネットワークエラーなどの場合は認証されていない状態とする
      setIsAuthenticated(false);
    }
  };

  const login = (token?: string) => {
    if (token) {
      localStorage.setItem("token", token);
    }
    setIsAuthenticated(true);
  };

  const logout = () => {
    localStorage.removeItem("token");
    setIsAuthenticated(false);
  };

  // 初期状態をlocalStorageベースで判定
  const checkInitialAuthStatus = () => {
    const token = localStorage.getItem("token");
    setIsAuthenticated(!!token);
  };

  useEffect(() => {
    checkInitialAuthStatus();
  }, []);

  return (
    <AuthContext.Provider value={{ 
      isAuthenticated, 
      setIsAuthenticated, 
      checkAuthStatus, 
      refreshAuthStatus: checkInitialAuthStatus,
      login,
      logout
    }}>
      {children}
    </AuthContext.Provider>
  );
};