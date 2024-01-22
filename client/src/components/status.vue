<script setup lang="ts">
import { PropType, computed } from 'vue';
import {Status, Team} from '../domain'
import StatusBar from './status-bar.vue';

const props =defineProps({
    status: {type: Object as PropType<Status>, required: true}
})


const highestStatus = computed(() => {
    let highestStatus = 1;

    for(const t of Object.values(Team)) {
        if (props.status[t] > highestStatus) {
            highestStatus = props.status[t]
        }
    }

    return highestStatus
})

</script>

<template>
    <div class="px-5 pb-12">
        <h3 class="text-center text-2xl mb-6">Suma wpłat oraz godzin</h3>
        <div class="flex justify-between items-end">
            <div v-for="team in Team">
                <StatusBar :value="status[team]" :team="team" :height="status[team]/highestStatus" />
            </div>
        </div>
    </div>

</template>