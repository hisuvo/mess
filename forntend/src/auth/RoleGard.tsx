import { Navigate, Outlet } from "react-router";
import type { UserRole } from "../types/auth.type";
import { useAuth } from "../hooks/useAuth";


interface RoleGuardProps {
  allowedRoles: UserRole[];
}

export default function RoleGuard({allowedRoles}: RoleGuardProps) {
  const { user, isAuthenticated, isLoading } = useAuth();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (!isAuthenticated || !user) {
    return <Navigate to="/login" replace />;
  }

  // Check whether the logged-in user has an allowed role.
  if (!allowedRoles.includes(user.role)) {
    return <Navigate to="/403" replace />;
  }

  return <Outlet />;
}
