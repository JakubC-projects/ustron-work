import { CreateRegistration, Registration, Status, User } from "./domain"

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

export async function getStatus(): Promise<Status> {
  const response = await fetch("/api/status");
  const status = await response.json()
  return status
}

export async function getOnTrackStatus(): Promise<Status> {
  const response = await fetch("/api/on-track");
  const status = await response.json()
  return status
}