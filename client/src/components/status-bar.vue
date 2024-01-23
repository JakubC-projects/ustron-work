
<script setup lang="ts">
import { PropType, computed } from 'vue';
import { Team, fillColors } from '../domain';
import Helmet from './helmet.vue';

const props =defineProps({
    value: {type: Number, required: true},
    height: {type: Number, required: true},
    team: {type: String as PropType<Team>, required: true}
})

const valueInt = computed(() => {
  return Math.round(props.value)
})

const heightInPx = computed(() => {
    return maxHeight * props.height
})

const maxHeight = 50

</script>

<template>
    <div class="flex flex-col items-center justify-end">
        <div class="w-1 transition-[height] duration-500" :style="`height: ${maxHeight - heightInPx}px;`" ></div>
        <p class="value pb-1" :style="`--num: ${valueInt}`"></p>
        <div class="w-14 mb-3 transition-[height] duration-500" :style="`background-color: ${fillColors[team]}; height: ${heightInPx}px;`" ></div>
        <Helmet class="w-12" :team="team" />
    </div>
</template>

<style scoped>

@property --num {
  syntax: '<integer>';
  initial-value: 0;
  inherits: false;
}

.value {
  transition: --num 500ms;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  counter-reset: num var(--num);
}

.value::after {
  content: counter(num);
}

</style>