<script setup lang="ts">
import { PropType } from 'vue';
import {Gender, GenderStatus} from '../domain'

defineProps({
    status: {type: Object as PropType<GenderStatus>, required: true}
})

</script>

<template>
    <div class="background px-5 py-8">
        <h3 class="text-center text-2xl mb-6">Dziewczyny vs Chłopaki</h3>
        <div class="grid grid-cols-2 items-end mb-8 text-2xl font-bold">
            <div class=" text-center mx-auto">
                <img class="w-16" src="/src/assets/woman.png" />
                <p><span class="value" :style="`--num: ${status[Gender.Female]}`"></span>%</p>
            </div>
            <div class="text-center mx-auto">
                <img class="w-16" src="/src/assets/man.png" />
                <p><span class="value" :style="`--num: ${status[Gender.Male]}`"></span>%</p>
            </div>
        </div>
    </div>
</template>

<style scoped>
.background {
    background-image: linear-gradient(110deg, #1E302D 0 50%, #193A36 50% 100%)
}


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