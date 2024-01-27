import { defineStore } from "pinia";
import { GenderStatus, Person, Registration, Status, newGenderStatus, newStatus } from "../domain";
import { getMe, getMyRegistrations, getOnTrackGenderStatus, getOnTrackStatus, getStatus } from "../client";
import { useRouter } from "vue-router";

type State = {
    me?: Person,
    myRegistrations: Registration[]
    status: Status,
    onTrack: Status
    onTrackGender: GenderStatus
}

const router = useRouter()

export const useStore = defineStore('main', {
    state: (): State => {
        return {
            myRegistrations: [],
            status: newStatus(),
            onTrack: newStatus(),
            onTrackGender: newGenderStatus(),
        }
    },

    actions: {
        async loadAll() {
            return Promise.all([this.loadMe(), this.loadMyRegistrations(), this.loadStatus(), this.loadOnTrackStatus(), this.loadOnTrackGenderStatus()])
        },

        async loadMe() {
            try {
                this.me = await getMe()
            } catch(err: unknown) {
                router.replace({
                    path: '/error',
                    query: {
                    error: "cannot find me"
                    }
                })
            throw err
            }
        },
        async loadMyRegistrations() {
            this.myRegistrations = await getMyRegistrations()
        },
        async loadStatus() {
            this.status = await getStatus()
        },
        async loadOnTrackStatus() {
            this.onTrack = await getOnTrackStatus()
        },
        async loadOnTrackGenderStatus() {
            this.onTrackGender = await getOnTrackGenderStatus()
        }
    }
})