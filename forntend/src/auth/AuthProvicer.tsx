
import {
  createContext,
  useEffect,
  useState,
  type ReactNode,
} from "react";
import type { AuthUser } from "../types/auth.type";
import { AuthContext } from "../context/AuthContext";




interface AuthProviderProps {
  children: ReactNode;
}

export function AuthProvider({ children }: AuthProviderProps) {
  const [user, setUser] = useState<AuthUser | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  // Restore the authenticated session when the app starts.
  useEffect(() => {
    async function loadUser() {
      try {
        const response = await fetch(
          "http://localhost:5000/api/v1/auth/me",
          {
            credentials: "include",
          }
        );

        if (!response.ok) {
          setUser(null);
          return;
        }

        const result = await response.json();

        setUser(result.data.user);
      } catch (error) {
        console.error("Authentication failed:", error);
        setUser(null);
      } finally {
        setIsLoading(false);
      }
    }

    loadUser();
  }, []);

  // Clear the session through the backend during logout.
  async function logout() {
    try {
      await fetch(
        "http://localhost:5000/api/v1/auth/logout",
        {
          method: "POST",
          credentials: "include",
        }
      );
    } finally {
      setUser(null);
    }
  }

  return (
    <AuthContext.Provider
      value={{
        user,
        isAuthenticated: Boolean(user),
        isLoading,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

