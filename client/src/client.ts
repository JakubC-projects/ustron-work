import { User } from "./domain"

export async function getMe(): Promise<User> {
    const res = await fetch("/api/me")
    const me = await res.json() as User

    return me
}