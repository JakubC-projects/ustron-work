export type Person = {
    displayName: string
    personID: number
    team: Team
}

export enum Team {
    Blue = "Blue",
    Green = "Green",
    Orange = "Orange",
    Red = "Red",
}

export enum Gender {
    Male = "Male",
    Female = "Female",
}


export type CreateRegistration = {
    type: RegistrationType
    date: string
    hourlyWage: number
    hours: number

    paidSum: number

    goal: Goal
    description: string
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
    BUK = "BUK",
    Maintenance = "Maintenance"
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

export type GenderStatus = Record<Gender, number>

export function newGenderStatus(): GenderStatus {
    return {
        [Gender.Male]: 0,
        [Gender.Female]: 0,
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