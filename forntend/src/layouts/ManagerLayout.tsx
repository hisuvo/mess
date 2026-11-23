// import React from 'react'

import { NavLink, Outlet } from "react-router"

const ManagerLayout = () => {
  return <div className="flex min-h-screen">
      <aside className="bg-green-400">
        <h2>Manager Panel</h2>

        <p><NavLink to="/admin" >Dashboard</NavLink></p>
        <p><NavLink to="/admin/users" >Users1</NavLink></p>
        <p><NavLink to="/admin/users" >Users2</NavLink></p>
        <p><NavLink to="/admin/users" >Users3</NavLink></p>
        <p><NavLink to="/admin/users" >Users4</NavLink></p>
      </aside>

      <main className="flex-1 p-6">
        <Outlet />
      </main>
    </div>
}

export default ManagerLayout
