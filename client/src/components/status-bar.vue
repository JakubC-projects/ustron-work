
<script setup lang="ts">
import { PropType, computed } from 'vue';
import { Team, fillColors } from '../domain';
import Helmet from './helmet.vue';

const props = defineProps({
    value: {type: Number, required: true},
    height: {type: Number, required: true},
    team: {type: String as PropType<Team>, required: true},
    hasWon: {type: Boolean, default: false},
    isBlurred: {type: Boolean, default: false}
})


const heightInPx = computed(() => {
    return maxHeight * props.height
})

const maxHeight = 50

</script>

<template>
    <div class="flex flex-col items-center justify-end" :class="isBlurred ? 'blur-[2px]' : ''">
        <div class="w-1 transition-[height] duration-500" :style="`height: ${maxHeight - heightInPx}px;`" ></div>
        <p class="value pb-1 font-bold" :class="hasWon ? 'text-xl': ''">{{ value.toLocaleString("pl-PL", {maximumFractionDigits: 0}) }}</p>
        <div class="mb-3 transition-[height] duration-500" :class="hasWon ? 'w-[92px]': 'w-[72px] '" :style="`background-color: ${fillColors[team]}; height: ${heightInPx}px;`" ></div>
        <Helmet class="w-12" :team="team" />
    </div>
</template>
