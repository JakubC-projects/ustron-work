import { CreateRegistration, Registration, Status, Person, GenderStatus } from "./domain"

export async function getMe(): Promise<Person> {
    const response = await fetch("/api/me")
    if (!response.ok) {
      throw Error(await response.text())
    }
  
    return await response.json()
}

export async function createMyRegistration(r: CreateRegistration): Promise<void> {
    const response = await fetch("/api/my-registrations", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(r),
      });

    if (!response.ok) {
      throw Error(await response.text())
    }

}


export async function getMyRegistrations(): Promise<Registration[]> {
  const response = await fetch("/api/my-registrations");
  if (!response.ok) {
    throw Error(await response.text())
  }

  return await response.json()
}

export async function getStatus(): Promise<Status> {
  const response = await fetch("/api/status");
  if (!response.ok) {
    throw Error(await response.text())
  }

  return await response.json()
}

export async function getOnTrackStatus(): Promise<Status> {
  const response = await fetch("/api/on-track");
  if (!response.ok) {
    throw Error(await response.text())
  }

  return await response.json()
}

export async function setOnTrackStatus(s: Status): Promise<void> {
  const response = await fetch("/api/on-track", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(s),
  });
  if (!response.ok) {
    throw Error(await response.text())
    ;
  }
}

export async function getOnTrackGenderStatus(): Promise<GenderStatus> {
  const response = await fetch("/api/on-track-gender");
  if (!response.ok) {
    throw Error(await response.text())
  }

  return await response.json()
}

export async function setOnTrackGenderStatus(s: GenderStatus): Promise<void> {
  const response = await fetch("/api/on-track-gender", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(s),
  });
  if (!response.ok) {
    throw Error(await response.text())
    ;
  }

}