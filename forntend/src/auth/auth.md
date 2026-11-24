আপনি React.js + Vite + React Router দিয়ে একটি **Role-Based Dashboard & Layout System (RBAC)** তৈরি করতে চাচ্ছেন।

আপনার সিস্টেমে চারটি Role থাকবে:

1. **OWNER** – পুরো Coaching Management System নিয়ন্ত্রণ করবে।
2. **ADMIN** – নিজের Coaching-এর administrative কাজ পরিচালনা করবে।
3. **MANAGER** – নির্দিষ্ট management-related কাজ পরিচালনা করবে।
4. **MEMBER** – সাধারণ user হিসেবে নিজের অনুমোদিত কাজ করবে।

আপনার মূল requirement হলো:

- OWNER login করলে শুধু Owner Layout এবং Owner Dashboard দেখবে।
- ADMIN login করলে শুধু Admin Layout এবং Admin Dashboard দেখবে।
- MANAGER login করলে শুধু Manager Layout এবং Manager Dashboard দেখবে।
- MEMBER login করলে শুধু Member Layout এবং Member Dashboard দেখবে।
- অন্য Role-এর Dashboard বা Layout access করতে পারবে না।
- URL সরাসরি লিখেও অন্য Role-এর Dashboard access করতে পারবে না।

এটি করার জন্য React Router-এর **Nested Routes + Protected Routes + Role-Based Authorization** ব্যবহার করবেন।

একটি গুরুত্বপূর্ণ বিষয়: শুধু React frontend দিয়ে security নিশ্চিত করা যায় না। আপনার Go backend-এও প্রতিটি protected API-তে role এবং permission যাচাই করতে হবে।

## **1. Recommended project architecture**

আপনার React project এ নিচের folder structure ব্যবহার করতে পারেন।

```
src/
├── app/
│   ├── router.tsx
│   └── providers.tsx
│
├── features/
│   └── auth/
│       ├── AuthProvider.tsx
│       ├── auth.service.ts
│       ├── ProtectedRoute.tsx
│       ├── RoleGuard.tsx
│       └── auth.types.ts
│
├── layouts/
│   ├── OwnerLayout.tsx
│   ├── AdminLayout.tsx
│   ├── ManagerLayout.tsx
│   └── MemberLayout.tsx
│
├── pages/
│   ├── auth/
│   │   └── LoginPage.tsx
│   │
│   ├── owner/
│   │   ├── OwnerDashboard.tsx
│   │   ├── CoachingPage.tsx
│   │   └── UsersPage.tsx
│   │
│   ├── admin/
│   │   ├── AdminDashboard.tsx
│   │   ├── StudentsPage.tsx
│   │   └── TeachersPage.tsx
│   │
│   ├── manager/
│   │   ├── ManagerDashboard.tsx
│   │   └── AttendancePage.tsx
│   │
│   ├── member/
│   │   ├── MemberDashboard.tsx
│   │   └── ProfilePage.tsx
│   │
│   └── errors/
│       ├── ForbiddenPage.tsx
│       └── NotFoundPage.tsx
│
└── main.tsx
```

এই architecture-এ প্রত্যেক Role-এর নিজস্ব Layout এবং Dashboard থাকবে।

## **2. Define user roles and types**

প্রথমে Role-এর জন্য একটি TypeScript type তৈরি করুন।

`src/features/auth/auth.types.ts`

```

export type UserRole =
  | "OWNER"
  | "ADMIN"
  | "MANAGER"
  | "MEMBER";

export interface AuthUser {
  id: number;
  name: string;
  email: string;
  role: UserRole;
  tenant_id?: number;
}

export interface AuthContextType {
  user: AuthUser | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  logout: () => Promise<void>;
}

```

এখানে `role` নির্ধারণ করবে user কোন dashboard ব্যবহার করতে পারবে।

## **3. Create AuthProvider**

এটি আপনার authentication system-এর কেন্দ্রীয় অংশ।

Login করার পর backend থেকে authenticated user-এর তথ্য নিতে হবে।

আমি recommend করব:

- Access token / session secure HttpOnly cookie-তে রাখবেন।
- Application reload হলে `/api/v1/auth/me` call করবেন।
- Backend authenticated user-এর তথ্য return করবে।
- React Context সেই তথ্য পুরো application-এ share করবে।

`src/features/auth/AuthProvider.tsx`

```

import {
  createContext,
  useContext,
  useEffect,
  useState,
  type ReactNode,
} from "react";

import type {
  AuthContextType,
  AuthUser,
} from "./auth.types";

const AuthContext = createContext<AuthContextType | null>(null);

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

// Reusable hook for accessing authentication state.
export function useAuth() {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error(
      "useAuth must be used inside AuthProvider"
    );
  }

  return context;
}

```

এখন application-এর যেকোনো component থেকে `useAuth()` ব্যবহার করে login করা user-এর information পাওয়া যাবে।

## **4. Create ProtectedRoute**

এটি নিশ্চিত করবে login ছাড়া কেউ dashboard access করতে পারবে না।

`src/features/auth/ProtectedRoute.tsx`

```

import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "./AuthProvider";

export default function ProtectedRoute() {
  const { user, isAuthenticated, isLoading } = useAuth();

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

```

এখানে `Outlet` হচ্ছে React Router-এর nested route rendering system।

## **5. Create RoleGuard**

এটাই আপনার system-এর সবচেয়ে গুরুত্বপূর্ণ অংশ।

এটি যাচাই করবে user-এর Role এবং requested Dashboard-এর Role একই কি না।

`src/features/auth/RoleGuard.tsx`

```

import {
  Navigate,
  Outlet,
} from "react-router-dom";

import { useAuth } from "./AuthProvider";
import type { UserRole } from "./auth.types";

interface RoleGuardProps {
  allowedRoles: UserRole[];
}

export default function RoleGuard({
  allowedRoles,
}: RoleGuardProps) {
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

```

উদাহরণ:

```
<RoleGuard allowedRoles={["ADMIN"]} />
```

এখানে শুধু ADMIN access পাবে।

```
<RoleGuard allowedRoles={["OWNER", "ADMIN"]} />
```

এখানে OWNER এবং ADMIN দুজনই access পাবে।

## **6. Create separate Layouts**

প্রত্যেক Role-এর জন্য আলাদা Layout তৈরি করুন।

[Shiptrack: Shipment Management Admin Dashboard Saas - Sidebar by Keitoto  on Dribbble](https://images.openai.com/static-rsc-4/p7G4Ny8P4oTwKZjZ63_k0V0HLaopuTdQsWSih3w_4Vfbyhep4zUaT9r53tlvk5gs3J915ZotmKyVnrhoHYYrJtUuw2kY1Q3bymLcGCReurFnNh74ZMQpj1fARCRIXw8BWTwoxN-AF0iyhEwP4fp61GwfMz34ZfBzVx_VYY7R0E0?purpose=inline)

**Admin Layout**

Students, Teachers, Payments

**Manager Layout**

Attendance, Reports, Batches

[Smart Agency OS — The Operating System for Modern Agencies](https://images.openai.com/static-rsc-4/EvxWVFkJCfQl7jn6F_4a9ZGYcoiS9hGOu2qV2WW8kQnVoLgyYuEWeUfMQ3-UjtAxcoi1MFWZ_c4sfdSES_CM286TeR8dBJr8worOzh0nfTEY0u7BybhypzUfG8jJOqmy7gTMFVjxRnanITpAnuuvjgOBAsdzXqZI7S8YQPXwLG4?purpose=inline)

**Owner Layout**

Coaching, Users, Analytics

[Student Portal designs, themes, templates and downloadable graphic elements on Dribbble](https://images.openai.com/static-rsc-4/opXlTukfv8TL6B_ffTRyuDo69jiuxk4uOUg67FkHkyPagZYoFvBlY5PDWqVgDzk5jFTAG3fHN_w-drJD_jHUfZPxK9CqLDcLEdqkNwRTlEYVXtikXLtcg9ItJPkZ_wcADZL76pkHsKaKyUVMBfUOpzs47xy92cvliLShCF0V7CM?purpose=inline)

**Member Layout**

Profile, Classes, Results

উদাহরণ হিসেবে Admin Layout:

`src/layouts/AdminLayout.tsx`

```

import { Outlet, NavLink } from "react-router-dom";
import { useAuth } from "../features/auth/AuthProvider";

export default function AdminLayout() {
  const { user, logout } = useAuth();

  return (
    <div className="flex min-h-screen">

      {/* Admin sidebar navigation */}
      <aside className="w-64 bg-gray-900 p-5 text-white">
        <h2 className="mb-6 text-xl font-bold">
          Admin Panel
        </h2>

        <nav className="flex flex-col gap-4">
          <NavLink to="/admin/dashboard">
            Dashboard
          </NavLink>

          <NavLink to="/admin/students">
            Students
          </NavLink>

          <NavLink to="/admin/teachers">
            Teachers
          </NavLink>
        </nav>
      </aside>

      {/* Admin main content */}
      <main className="flex-1 bg-gray-100">

        <header className="flex justify-between bg-white p-5 shadow">
          <h3>Admin Dashboard</h3>

          <div className="flex items-center gap-4">
            <span>{user?.name}</span>

            <button onClick={logout}>
              Logout
            </button>
          </div>
        </header>

        {/* Render only admin child pages */}
        <section className="p-6">
          <Outlet />
        </section>

      </main>
    </div>
  );
}

```

একইভাবে `OwnerLayout`, `ManagerLayout` এবং `MemberLayout` তৈরি করবেন। প্রতিটি Layout-এর Sidebar এবং Navigation আলাদা থাকবে।

## **7. Configure React Router**

এখন সবকিছু একসঙ্গে connect করবেন।

আপনার route structure হবে:

```

/
├── /login
│
├── /owner
│   ├── /dashboard
│   ├── /coachings
│   └── /users
│
├── /admin
│   ├── /dashboard
│   ├── /students
│   └── /teachers
│
├── /manager
│   ├── /dashboard
│   └── /attendance
│
├── /member
│   ├── /dashboard
│   └── /profile
│
└── /403

```

`src/app/router.tsx`

```

import {
  createBrowserRouter,
  Navigate,
} from "react-router-dom";

import ProtectedRoute from "../features/auth/ProtectedRoute";
import RoleGuard from "../features/auth/RoleGuard";
import { useAuth } from "../features/auth/AuthProvider";

import OwnerLayout from "../layouts/OwnerLayout";
import AdminLayout from "../layouts/AdminLayout";
import ManagerLayout from "../layouts/ManagerLayout";
import MemberLayout from "../layouts/MemberLayout";

import LoginPage from "../pages/auth/LoginPage";

import OwnerDashboard from "../pages/owner/OwnerDashboard";
import CoachingPage from "../pages/owner/CoachingPage";
import UsersPage from "../pages/owner/UsersPage";

import AdminDashboard from "../pages/admin/AdminDashboard";
import StudentsPage from "../pages/admin/StudentsPage";
import TeachersPage from "../pages/admin/TeachersPage";

import ManagerDashboard from "../pages/manager/ManagerDashboard";
import AttendancePage from "../pages/manager/AttendancePage";

import MemberDashboard from "../pages/member/MemberDashboard";
import ProfilePage from "../pages/member/ProfilePage";

import ForbiddenPage from "../pages/errors/ForbiddenPage";

// Redirect the authenticated user to their own dashboard.
function DashboardRedirect() {
  const { user, isLoading } = useAuth();

  if (isLoading) return <div>Loading...</div>;

  if (!user) return <Navigate to="/login" replace />;

  const dashboardRoutes = {
    OWNER: "/owner/dashboard",
    ADMIN: "/admin/dashboard",
    MANAGER: "/manager/dashboard",
    MEMBER: "/member/dashboard",
  };

  return (
    <Navigate
      to={dashboardRoutes[user.role]}
      replace
    />
  );
}

export const router = createBrowserRouter([
  {
    path: "/login",
    element: <LoginPage />,
  },

  {
    path: "/",
    element: <DashboardRedirect />,
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
                path: "coachings",
                element: <CoachingPage />,
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
                path: "students",
                element: <StudentsPage />,
              },
              {
                path: "teachers",
                element: <TeachersPage />,
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
                element: <AttendancePage />,
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

```

এখানে লক্ষ্য করুন:

```
element: <RoleGuard allowedRoles={["ADMIN"]} />
```

ADMIN ছাড়া অন্য কেউ `/admin/dashboard` access করতে পারবে না।

আবার:

```
element: <RoleGuard allowedRoles={["OWNER"]} />
```

এখানে শুধু OWNER access করতে পারবে।

## **8. Connect Router with main.tsx**

`src/main.tsx`

```

import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider } from "react-router-dom";

import { AuthProvider } from "./features/auth/AuthProvider";
import { router } from "./app/router";

import "./index.css";

ReactDOM.createRoot(
  document.getElementById("root")!
).render(
  <React.StrictMode>
    <AuthProvider>
      <RouterProvider router={router} />
    </AuthProvider>
  </React.StrictMode>
);

```

## **9. Create separate Dashboard pages**

এখন প্রত্যেক Role-এর Dashboard আলাদা হবে।

**Owner Dashboard**

`src/pages/owner/OwnerDashboard.tsx`

```
export default function OwnerDashboard() {
  return (
    <div>
      <h1 className="text-2xl font-bold">
        Owner Dashboard
      </h1>

      <div className="grid grid-cols-3 gap-5 mt-6">
        <div className="rounded-lg bg-white p-5 shadow">
          Total Coachings: 25
        </div>

        <div className="rounded-lg bg-white p-5 shadow">
          Total Users: 500
        </div>

        <div className="rounded-lg bg-white p-5 shadow">
          Total Revenue: 100000
        </div>
      </div>
    </div>
  );
}
```

**Admin Dashboard**

`src/pages/admin/AdminDashboard.tsx`

```
export default function AdminDashboard() {
  return (
    <div>
      <h1 className="text-2xl font-bold">
        Admin Dashboard
      </h1>

      <div className="grid grid-cols-3 gap-5 mt-6">
        <div className="bg-white p-5 shadow rounded-lg">
          Total Students: 150
        </div>

        <div className="bg-white p-5 shadow rounded-lg">
          Total Teachers: 20
        </div>

        <div className="bg-white p-5 shadow rounded-lg">
          Total Batches: 15
        </div>
      </div>
    </div>
  );
}
```

Manager এবং Member-এর জন্যও একইভাবে নিজস্ব dashboard component তৈরি করবেন।

## **10. Create Forbidden Page**

কেউ অন্য Role-এর Dashboard access করতে চাইলে এই page দেখাবেন।

`src/pages/errors/ForbiddenPage.tsx`

```
import { Link } from "react-router-dom";

export default function ForbiddenPage() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center">
      <h1 className="text-6xl font-bold">403</h1>

      <p className="mt-3">
        You are not authorized to access this page.
      </p>

      <Link
        to="/"
        className="mt-5 text-blue-600 underline"
      >
        Go to your Dashboard
      </Link>
    </div>
  );
}
```

## **11. Login-এর পর কীভাবে কাজ করবে?**

ধরুন আপনার backend এই response return করছে:

```
{
  "success": true,
  "data": {
    "user": {
      "id": 10,
      "name": "Suvo",
      "email": "suvo@example.com",
      "role": "ADMIN",
      "tenant_id": 5
    }
  }
}
```

Frontend login success হওয়ার পর `/` route-এ navigate করবে।

`DashboardRedirect` user-এর role দেখে redirect করবে।

**Role-based navigation flow**

**User Login**

**Backend Authentication**

Verify session / JWT and retrieve user role

**Read User Role**

OWNER

**/owner/dashboard**

ADMIN

**/admin/dashboard**

MANAGER

**/manager/dashboard**

MEMBER

**/member/dashboard**

**Role-specific Layout + Dashboard**

## **12. Very important: Secure your Go Backend**

আপনার React frontend-এ RoleGuard ব্যবহার করলেই security সম্পূর্ণ হবে না।

কেউ Postman, Apidog অথবা browser DevTools দিয়ে সরাসরি API call করতে পারে।

তাই Go + Echo backend-এও Role Authorization implement করতে হবে।

উদাহরণ:

| API                  | OWNER | ADMIN | MANAGER | MEMBER |
| -------------------- | ----- | ----- | ------- | ------ |
| `/api/v1/coachings`  | Yes   | No    | No      | No     |
| `/api/v1/students`   | No    | Yes   | Limited | No     |
| `/api/v1/attendance` | No    | Yes   | Yes     | No     |
| `/api/v1/profile`    | Yes   | Yes   | Yes     | Yes    |
| `/api/v1/reports`    | Yes   | Yes   | Limited | No     |

আপনার Go backend-এ middleware ব্যবহার করুন।

উদাহরণ:

```
// Only ADMIN can access this endpoint.
adminGroup := e.Group("/api/v1/admin")

adminGroup.Use(
    middleware.JWTMiddleware(),
    middleware.RequireRole("ADMIN"),
)

adminGroup.GET("/students", studentHandler.List)
```

এটি illustrative structure; আপনার বর্তমান Echo v5 middleware implementation অনুযায়ী function signature adjust করবেন।

**Multi-tenant security-ও গুরুত্বপূর্ণ।**

আপনার Coaching Management SaaS-এ:

- OWNER: platform-level access।
- ADMIN: নিজের Coaching-এর data।
- MANAGER: নিজের Coaching-এর অনুমোদিত data।
- MEMBER: নিজের অনুমোদিত data।

শুধু Role যাচাই যথেষ্ট নয়। প্রতিটি database query-তেও `tenant_id` scope যাচাই করতে হবে।

যেমন:

```
// Restrict student records to the authenticated tenant.
db.Where("tenant_id = ?", tenantID).
   Find(&students)
```

কোনো user request থেকে পাঠানো `tenant_id`-কে অন্ধভাবে বিশ্বাস করবেন না। Authenticated session/JWT থেকে server-side tenant identity নেবেন।

## **13. Recommended technology stack**

আপনার React + Vite + Go project-এর জন্য:

| Technology          | Purpose                       |
| ------------------- | ----------------------------- |
| React + Vite        | Frontend                      |
| TypeScript          | Type safety                   |
| React Router        | Routing                       |
| React Context       | Authentication state          |
| TanStack Query      | API data fetching and caching |
| Axios / Fetch       | HTTP requests                 |
| Zod                 | Data validation               |
| Tailwind CSS        | UI styling                    |
| Go + Echo           | Backend                       |
| PostgreSQL          | Database                      |
| JWT / Secure Cookie | Authentication                |

## **14. Test your Role-Based System**

এই test cases অবশ্যই করবেন।

**Security testing checklist**

0/6 tested

**ADMIN login**

Expected: Admin Dashboard

**ADMIN opens Owner Dashboard**

Expected: 403 Forbidden

**MEMBER opens Admin Dashboard**

Expected: 403 Forbidden

**MANAGER opens Member Dashboard**

Expected: 403 Forbidden

**OWNER opens Owner Dashboard**

Expected: Owner Dashboard

**Unauthenticated user opens dashboard**

Expected: Redirect to Login

Reset checklist

## **My architectural recommendation**

আপনার Coaching Management SaaS-এর জন্য আমি এই architecture recommend করব:

**Authentication + Role-Based Access Control + Permission-Based Access Control**

কারণ ভবিষ্যতে শুধু চারটি Role থাকলেই হবে না। যেমন:

- ADMIN: Students Create, Update, Delete করতে পারবে।
- MANAGER: Students View এবং Attendance Manage করতে পারবে।
- MEMBER: শুধু নিজের Profile এবং Results দেখতে পারবে।

এজন্য ভবিষ্যতে `roles`, `permissions`, `role_permissions` এবং `user_roles` structure ব্যবহার করতে পারেন।

আরেকটি বিষয়: Role হলো user-এর পরিচয়/দায়িত্ব, Permission হলো user কী কী কাজ করতে পারবে।

**আমার পরামর্শ:** শুরুতে RoleGuard দিয়ে চারটি Dashboard তৈরি করুন। তারপর Permission-Based Access Control যুক্ত করুন। এতে আপনার application maintainable এবং scalable হবে।

### **পরবর্তী ধাপ**

আমি আপনার existing Go + Echo + GORM + PostgreSQL backend-এর সঙ্গে মিলিয়ে একটি complete production-style RBAC frontend তৈরি করতে সাহায্য করতে পারি, যেখানে থাকবে separate sidebar, responsive layout, login/logout, role-based navigation, permission-based menu এবং protected API integration।