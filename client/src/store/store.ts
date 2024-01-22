import { defineStore } from "pinia";
import { Person, Registration, Status, newStatus } from "../domain";
import { getMe, getMyRegistrations, getOnTrackStatus, getStatus } from "../client";
import { useRouter } from "vue-router";

type State = {
    me?: Person,
    myRegistrations: Registration[]
    status: Status,
    onTrack: Status
}

const router = useRouter()

export const useStore = defineStore('main', {
    state: (): State => {
        return {
            myRegistrations: [],
            status: newStatus(),
            onTrack: newStatus(),
        }
    },

    actions: {
        async loadAll() {
            return Promise.all([this.loadMe(), this.loadMyRegistrations(), this.loadStatus(), this.loadOnTrackStatus()])
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
        }
    }
})