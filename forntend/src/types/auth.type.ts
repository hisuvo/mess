export type UserRole = "OWNER" | "ADMIN" | "MANAGER" | "MEMBER";

export interface AuthUser {
  id: number;
  name: string;
  email: string;
  role: UserRole;
  tenant_id?: number;
}

export interface AuthContextType {
    user : AuthUser | null;
    isAuthenticated : boolean;
    isLoading: boolean;
    // here Promise means this not at the end and void say this function not some return
    logout:() => Promise<void> 
}