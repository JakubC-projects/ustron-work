<script setup lang="ts">
defineProps({isOpen: {type: Boolean, default: false}})
</script>

<template>
    <button class="bg-black p-4 rounded-full">
        <img src="../assets/plus.svg" class="transition-transform" :class="isOpen ? 'rotate-45' : ''"/>
    </button>
</template>