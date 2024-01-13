export type User = {
    displayName: string
    uid: string
    team: Team
}

export enum Team {
    Red = "Red",
    Green = "Green",
    Blue = "Blue",
    Orange = "Orange",
}