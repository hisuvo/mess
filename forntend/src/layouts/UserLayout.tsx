// import React from 'react'

import { Outlet } from "react-router"

const UserLayout = () => {
  return (
    <div>
      <main className="flex-1 p-6">
        <Outlet/>
      </main>
    </div>
  )
}

export default UserLayout