
import { Outlet, NavLink } from "react-router";

export default function AdminLayout() {
  return (
    <div className="flex min-h-screen bg-amber-300">
      <aside className="bg-green-400">
        <h2>Admin Panel</h2>

        <NavLink to="/admin" >Dashboard</NavLink>
        <br />
        <NavLink to="/admin/users" >Users</NavLink>
      </aside>

      <main className="flex-1 p-6">
        <Outlet />
      </main>
    </div>
  );
}