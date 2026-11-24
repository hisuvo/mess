import { Navigate } from "react-router"
import { useAuth } from "../hooks/useAuth"


export const DashboardRedriect = () => {
  const {user, isLoading} = useAuth()

  if(isLoading) return <div>loading...</div>
  
  if (!user) return <Navigate to="/login" replace />;

    const dashboardRoutes = {
        OWNER: "/owner/dashboard",
        ADMIN: "/admin/dashboard",
        MANAGER: "/manager/dashboard",
        MEMBER: "/member/dashboard",
    };

  return <Navigate to={dashboardRoutes[user.role]} replace/>
}
