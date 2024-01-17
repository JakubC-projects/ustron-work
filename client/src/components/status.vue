<script setup lang="ts">
import { PropType, computed } from 'vue';
import {Status, Team, fillColors} from '../domain'
import Helmet from './helmet.vue';

const props =defineProps({
    status: {type: Object as PropType<Status>, required: true}
})

const maxHeight = 50

const heightMultiplier = computed(() => {
    let highestStatus = 0;

    for(const t of Object.values(Team)) {
        if (props.status[t] > highestStatus) {
            highestStatus = props.status[t]
        }
    }

    return maxHeight / highestStatus
})

</script>

<template>
    <div class="px-5 pb-12">
        <h3 class="text-center text-2xl mb-6">Suma wpłat oraz godzin</h3>
        <div class="flex justify-between items-end">
            <div v-for="team in Team">
                <div class="flex flex-col items-center justify-end">
                    <p class="pb-1">{{ status[team] }}</p>
                    <div class="w-14 mb-3 transition-[height]" :style="`background-color: ${fillColors[team]}; height: ${status[team] * heightMultiplier}px;`" ></div>
                    <Helmet class="w-12" :team="team" />
                </div>
            </div>
        </div>
    </div>

</template>