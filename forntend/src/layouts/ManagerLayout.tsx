// import React from 'react'

import { Outlet } from "react-router"

const ManagerLayout = () => {
  return <div className="flex min-h-screen">
      <main className="flex-1 p-6">
        <Outlet />
      </main>
    </div>
}

export default ManagerLayout
