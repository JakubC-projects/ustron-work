<script setup lang="ts">
import {  ref } from 'vue';

defineProps({
    label: {type: String, default: "Open"}
})

const isOpen = ref(false)

</script>

<template>
    <div class=" bg-[#21413E] rounded-xl ">
        <div class="px-8 py-6  flex justify-between cursor-pointer" @click="isOpen = !isOpen">
            <h3 class="font-bold text-xl">{{label}}</h3>
            <img src="../assets/chevron.svg" class="transition-transform" :class="isOpen ? 'rotate-180' : ''">
        </div>
        <div class="transition-[max-height] overflow-y-scroll" :class="isOpen ? 'max-h-[512px]' : 'max-h-0'">
            <slot></slot>
        </div>
    </div>
</template>