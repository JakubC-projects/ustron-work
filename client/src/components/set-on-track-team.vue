
<script setup lang="ts">
import { PropType } from 'vue';
import { Team } from '../domain';
import Helmet from './helmet.vue';

defineProps({
    team: {type: String as PropType<Team>, required: true}
})

const model = defineModel<number>({required: true})

</script>

<template>
    <div class="flex gap-4 items-center justify-end">
        <Helmet class="w-16" :team="team"/>
        <input type="number" v-model="model"/>
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