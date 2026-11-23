
import { redirect } from "react-router"
import { getCurrentUser} from "./auth.service"
import type { Role } from "./role"

// This function checks:
// Is the user authenticated?
// Does the user have the required role?
// If not, redirect them to the appropriate page.

export const RequireAuth =(allowedRoles:Role[]) => {
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
