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

    if (response.ok) {
      return await response.json()
    }

    throw Error(await response.text())
}


export async function getMyRegistrations(): Promise<Registration[]> {
  const response = await fetch("/api/my-registrations");
  if (response.ok) {
    return await response.json()
  }

  throw Error(await response.text())
}

export async function getStatus(): Promise<Status> {
  const response = await fetch("/api/status");
  if (response.ok) {
    return await response.json()
  }

  throw Error(await response.text())
}

export async function getOnTrackStatus(): Promise<Status> {
  const response = await fetch("/api/on-track");
  if (response.ok) {
    return await response.json()
  }

  throw Error(await response.text())
}

export async function setOnTrackStatus(s: Status): Promise<void> {
  const response = await fetch("/api/on-track", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(s),
  });
  if (response.ok) {
    return await response.json()
  }

  throw Error(await response.text())
}