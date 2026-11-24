import { useAuth } from '../hooks/useAuth'
import { Navigate, Outlet } from 'react-router';

const ProtectedRoute = () => {
    const {user, isLoading, isAuthenticated} = useAuth()
   // Wait until the authentication check finishes.
    if (isLoading) {
        return <div>Loading...</div>;
    }

    // Redirect unauthenticated users to login.
    if (!isAuthenticated || !user) {
        return <Navigate to="/login" replace />;
    }

    // Allow authenticated users to access nested routes.
    return <Outlet />;
}

export default ProtectedRoute