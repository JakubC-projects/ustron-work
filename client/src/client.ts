import { CreateRegistration, Registration, User } from "./domain"

export async function getMe(): Promise<User> {
    const res = await fetch("/api/me")
    const me = await res.json() as User

    return me
}

export async function createMyRegistration(r: CreateRegistration): Promise<Registration> {
    const response = await fetch("/api/my-registrations", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(r),
      });

    const registration = await response.json()

    return registration
}


export async function getMyRegistrations(): Promise<Registration[]> {
  const response = await fetch("/api/my-registrations");

  const registrations = await response.json()

  return registrations
}