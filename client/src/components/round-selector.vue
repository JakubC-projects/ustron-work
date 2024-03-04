<template>
    <div class="w-full flex justify-center">
        <div class="flex items-center py-1 px-3 bg-[#4b5957] rounded-full">
            <button v-for="round of rounds" class="py-[2px] px-4 rounded-full font-bold text-lg" :class="round.id == currentRound.id ? 'bg-white text-black' : ''" @click="currentRound = round">
                Runda {{ round.id }}
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { PropType } from 'vue';
import { Round } from '../domain';

const currentRound = defineModel({
    type: Object as PropType<Round>,
    required: true
})

defineProps({
    rounds: {type: Array as PropType<Round[]>, required: true}
})
</script>