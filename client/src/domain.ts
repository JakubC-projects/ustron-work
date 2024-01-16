export type User = {
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
    SuperAdmin = "SuperAdmin",
}

export type CreateRegistration = {
    type: RegistrationType

    hourlyWage: number
    hours: number

    paidSum: number
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

export type Status = Record<Team, number>

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