import { CreateRegistration, Registration, Status, Person, GenderStatus, Round } from "./domain"

export async function getMe(): Promise<Person> {
    const response = await fetch("/api/me")
    if (!response.ok) {
      throw Error(await response.text())
    }
  
    return await response.json()
}

export async function getRounds(): Promise<Round[]> {
  const response = await fetch(`/api/rounds`);
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

export async function getMyRegistrations(roundId: number): Promise<Registration[]> {
  const response = await fetch(`/api/my-registrations?roundId=${roundId}`);
  if (!response.ok) {
    throw Error(await response.text())
  }

  return await response.json()
}

export async function getStatus(roundId: number): Promise<Status> {
  const response = await fetch(`/api/status?roundId=${roundId}`);
  if (!response.ok) {
    throw Error(await response.text())
  }

  return await response.json()
}

export async function getOnTrackStatus(roundId: number): Promise<Status> {
  const response = await fetch(`/api/on-track?roundId=${roundId}`);
  if (!response.ok) {
    throw Error(await response.text())
  }

  return await response.json()
}


export async function getOnTrackGenderStatus(roundId: number): Promise<GenderStatus> {
  const response = await fetch(`/api/on-track-gender?roundId=${roundId}`);
  if (!response.ok) {
    throw Error(await response.text())
  }

  return await response.json()
}
