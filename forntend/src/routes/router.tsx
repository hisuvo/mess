import { createBrowserRouter } from "react-router";
// import { RequireAuth } from "../auth/RequireAuth";
import AdminDashboard from "../pages/admin/Dashboard";
import MangeUsers from "../pages/admin/MangeUsers";
import ManagerDashboard from "../pages/manager/Dashboard";
import ManagerLayout from "../layouts/ManagerLayout";
import Reports from "../pages/manager/Reports";
import AdminLayout from "../layouts/AdminLayout";


const router = createBrowserRouter([
    // Publice
    {path:"/", element:<h1>Hello World</h1> },
    
    // Admin
    {
        path:"/admin",
        // loader: RequireAuth(["ADMIN"]),
        Component: AdminLayout,
        children: [
                { index: true, Component: AdminDashboard },
                { path: "users", Component: MangeUsers },
            ]

    },

    // Manager 
    {
        path:"/manager",
        Component: ManagerLayout,
        children: [
            { index:true, Component:ManagerDashboard },
            { path:"report", Component:Reports },
        ]
    },

    // User
    {
        path:"/users",
        Component: ManagerLayout,
        children: [
            {index:true, Component:ManagerDashboard},
            {path:"report", Component:Reports},
        ]
    }
]);

export default router