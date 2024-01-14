export type User = {
    displayName: string
    personID: number
    team: Team
    role: Role
}

export enum Team {
    Red = "Red",
    Green = "Green",
    Blue = "Blue",
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