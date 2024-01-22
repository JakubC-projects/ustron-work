<script setup lang="ts">
import { PropType } from 'vue';
import {Status, Team} from '../domain'
import OnTrackTeam from './on-track-team.vue';

defineProps({
    status: {type: Object as PropType<Status>, required: true}
})

</script>

<template>
    <div class="on-track px-5 py-8">
        <h3 class="text-center text-2xl mb-6">Procent bycia <span class="text-[#FF0] italic">ON TRACK</span></h3>
        <div class="flex justify-between items-end mb-8">
            <div v-for="team in Team">
                <OnTrackTeam :team="team" :value="status[team]"/>
            </div>
        </div>
        <div class="flex justify-center">
            <div class="border-2 rounded-md border-[#FF0] text-[#FF0] flex gap-3 justify-center items-center px-2 py-1">
            <img src="../assets/warn.svg"/>
            <p class="text-xs">Polska nie jest ON TRACK!</p>
        </div>
        </div>
    </div>
</template>

<style>

.on-track {
    background-image: url("/src/assets/background.svg")
}
</style>