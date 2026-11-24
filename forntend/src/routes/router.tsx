import { createBrowserRouter, Navigate } from "react-router";
// import { RequireAuth } from "../auth/RequireAuth";
import AdminDashboard from "../pages/admin/Dashboard";
import ManagerDashboard from "../pages/manager/Dashboard";
import ManagerLayout from "../layouts/ManagerLayout";
import AdminLayout from "../layouts/AdminLayout";
import ProtectedRoute from "../auth/ProtectedRoute";
import { DashboardRedriect } from "../auth/DashboardRedirect";
import RoleGuard from "../auth/RoleGard";
import OwnerLayout from "../layouts/OwnerLayout";
import { MemberLayout } from "../layouts/MemberLayout";
import OwnerDashboard from "../components/dashboard/OwnerDashboard";
import MassPage from "../pages/owner/MassPage";
import UsersPage from "../pages/owner/UsersPage";
import MangeUsers from "../pages/admin/MangeUsers";
import PaymentPage from "../pages/admin/PaymentPage";
import MealAttendancePage from "../pages/member/MealAttendancePage";
import MemberDashboard from "../components/dashboard/MemberDashboard";
import ProfilePage from "../pages/member/ProfilePage";
import ForbiddenPage from "../pages/errors/ForbiddenPage";
import LoginPage from "../pages/auth/LoginPage";

export const router = createBrowserRouter([
  {
    path: "/login",
    element: <LoginPage />,
  },

  {
    path: "/",
    element: <DashboardRedriect />,
  },

  // Require authentication before entering dashboards.
  {
    element: <ProtectedRoute />,
    children: [

      // OWNER ROUTES
      {
        element: <RoleGuard allowedRoles={["OWNER"]} />,
        children: [
          {
            path: "/owner",
            element: <OwnerLayout />,
            children: [
              {
                index: true,
                element: <Navigate to="dashboard" replace />,
              },
              {
                path: "dashboard",
                element: <OwnerDashboard />,
              },
              {
                path: "Mass",
                element: <MassPage />,
              },
              {
                path: "users",
                element: <UsersPage />,
              },
            ],
          },
        ],
      },

      // ADMIN ROUTES
      {
        element: <RoleGuard allowedRoles={["ADMIN"]} />,
        children: [
          {
            path: "/admin",
            element: <AdminLayout />,
            children: [
              {
                index: true,
                element: <Navigate to="dashboard" replace />,
              },
              {
                path: "dashboard",
                element: <AdminDashboard />,
              },
              {
                path: "users",
                element: <MangeUsers />,
              },
              {
                path: "payments",
                element: <PaymentPage />,
              },
            ],
          },
        ],
      },

      // MANAGER ROUTES
      {
        element: <RoleGuard allowedRoles={["MANAGER"]} />,
        children: [
          {
            path: "/manager",
            element: <ManagerLayout />,
            children: [
              {
                index: true,
                element: <Navigate to="dashboard" replace />,
              },
              {
                path: "dashboard",
                element: <ManagerDashboard />,
              },
              {
                path: "attendance",
                element: <MealAttendancePage />,
              },
            ],
          },
        ],
      },

      // MEMBER ROUTES
      {
        element: <RoleGuard allowedRoles={["MEMBER"]} />,
        children: [
          {
            path: "/member",
            element: <MemberLayout />,
            children: [
              {
                index: true,
                element: <Navigate to="dashboard" replace />,
              },
              {
                path: "dashboard",
                element: <MemberDashboard />,
              },
              {
                path: "profile",
                element: <ProfilePage />,
              },
            ],
          },
        ],
      },
    ],
  },

  {
    path: "/403",
    element: <ForbiddenPage />,
  },
]);

// const router = createBrowserRouter([
//     // Publice
//     {path:"/", element:<h1>Hello World</h1> },
    
//     // Admin
//     {
//         path:"/admin",
//         // loader: RequireAuth(["ADMIN"]),
//         Component: AdminLayout,
//         children: [
//                 { index: true, Component: AdminDashboard },
//                 { path: "users", Component: MangeUsers },
//             ]

//     },

//     // Manager 
//     {
//         path:"/manager",
//         // loader: RequireAuth(["MANAGER"]),
//         Component: ManagerLayout,
//         children: [
//             { index:true, Component:ManagerDashboard },
//             { path:"report", Component:Reports },
//         ]
//     },

//     // User
//     {
//         path:"/users",
//         // loader: RequireAuth(["MEMBER"]),
//         Component: ManagerLayout,
//         children: [
//             {index:true, Component:ManagerDashboard},
//             {path:"report", Component:Reports},
//         ]
//     }
// ]);

export default router