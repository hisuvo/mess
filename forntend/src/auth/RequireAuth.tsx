
import { redirect } from "react-router"
import { getCurrentUser} from "./auth.service"
import type { UserRole } from "../types/auth.type"


// This function checks:
// Is the user authenticated?
// Does the user have the required role?
// If not, redirect them to the appropriate page.

export const RequireAuth =(allowedRoles:UserRole[]) => {
  return async () => {
    const user = await getCurrentUser()

    if (!user) {
      throw redirect('/login')
    }

    if (!allowedRoles.includes(user.role)) {
      throw redirect("/unauthorized");
    }

    return user
  }
}
