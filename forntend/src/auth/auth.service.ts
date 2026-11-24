import type { AuthUser } from "../types/auth.type";

export async function getCurrentUser():Promise<AuthUser | any> {
    const response = await fetch("/api/v1/auth/me", {
        credentials: "include",
    });

    if (response.status === 401) {
        return null
    }

    if (!response.ok){
         throw new Error("Unable to fetch current user");
    }

    return response.json()
}