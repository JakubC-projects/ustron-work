export type Person = {
    displayName: string
    personID: number
    team: Team
    role: Role
}

export enum Team {
    Blue = "Blue",
    Red = "Red",
    Green = "Green",
    Orange = "Orange",
}

export enum Role {
    Base = "Base",
    Admin = "Admin",
}

export type CreateRegistration = {
    type: RegistrationType
    date: string
    hourlyWage: number
    hours: number

    paidSum: number

    goal: Goal
    comment: string
}

export type Registration = CreateRegistration & {
    uid: string
    personID: number
    team: Team
}

export enum RegistrationType {
    Money = "Money",
    Work = "Work"
}

export enum Goal {
    Samvirk = "Samvirk",
    BUK = "BUK"
}

export type Status = Record<Team, number>

export function newStatus(): Status {
    return {
        Blue: 0,
        Green: 0,
        Orange: 0,
        Red: 0
    }
}

export const fillColors: Record<Team, string> = {
    [Team.Blue]:
         "#646CFF",
    [Team.Red]:
         "#FF414C",
    [Team.Green]:
         "#00DB90",
    [Team.Orange]:
         "#FF6D00"
}

export const strokeColors: Record<Team, string> = {
    [Team.Blue]:
         "#0617B7",
    [Team.Red]:
         "#B70042",
    [Team.Green]:
         "#006B3F",
    [Team.Orange]:
         "#CC2B00"
}