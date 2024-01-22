
<script setup lang="ts">
import { PropType } from 'vue';
import { Team } from '../domain';
import Helmet from './helmet.vue';

defineProps({
    value: {type: Number, required: true},
    team: {type: String as PropType<Team>, required: true}
})

</script>

<template>
    <div class="flex flex-col items-center justify-end">
        <Helmet class="w-16" :team="team" :fill="value / 100"/>
        <p class="pb-1"><span class="value" :style="`--num: ${value}`"></span> %</p>
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