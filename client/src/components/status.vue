<script setup lang="ts">
import { PropType, computed } from 'vue';
import {Round, Status, Team} from '../domain'
import StatusBar from './status-bar.vue';

const props =defineProps({
    status: {type: Object as PropType<Status>, required: true},
    round: {type: Object as PropType<Round>, required: true},
    userTeam: {type: String as PropType<Team>, required: true},
})

const isFrozen = computed(() => {
    const now = new Date().toISOString()
    if(now > props.round.freezeStartDate && now < props.round.endDate) {
        return true
    }
    return false 
})

const isDone = computed(() => {
    const now = new Date().toISOString()
    if(now > props.round.endDate) {
        return true
    }
    return false 
})

const frozenTime = computed(() => {
    return new Date(props.round.freezeStartDate).toLocaleString("pl", {dateStyle: "long", timeStyle:'medium'})
})

const winningTeam = computed(() => {
    let bestScore = 0
    let bestTeam = Team.Green
    for(const t of Object.values(Team)) {
        if (props.status[t] > bestScore) {
            bestScore = props.status[t]
            bestTeam = t
        }
    }

    return bestTeam

})

const highestStatus = computed(() => {
    if(props.status[winningTeam.value]) return props.status[winningTeam.value]
    return 1
})

</script>

<template>
    <div class="px-5 pt-3 pb-12">
        <h3 class="text-center text-2xl font-bold">Suma wpłat oraz godzin</h3>
        <p v-if="isFrozen" class="opacity-75 text-center">Widzisz aktualny wynik swojej drużyny i wyniki innych drużyn z <span class="font-bold">{{frozenTime}}</span></p>
        <p v-if="isDone" class="opacity-75 text-center">Gratulacje dla drużyny <span class="font-bold">{{ winningTeam }}</span></p>
        <div class="flex justify-between items-end mt-4">
            <div v-for="team in Team">
                <StatusBar :value="status[team]" :team="team" :height="status[team]/highestStatus" :is-blurred="isFrozen && team != userTeam" :has-won="isDone && team == winningTeam"/>
            </div>
        </div>
    </div>
</template>