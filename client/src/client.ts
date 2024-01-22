import { CreateRegistration, Registration, Status, Person } from "./domain"

export async function getMe(): Promise<Person> {
    const res = await fetch("/api/me")
    const me = await res.json() as Person

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

export async function setOnTrackStatus(s: Status): Promise<void> {
  await fetch("/api/on-track", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(s),
  });
}