<script setup lang="ts">
defineProps({
    label: {type: String}
})
</script>

<template>
    <div class="relative cursor-pointer">
        <img src="../assets/button.svg" />
        <p class="label">{{label}}</p>
    </div>
</template>

<style scoped>
.label {
    @apply absolute top-1/2 left-1/2 text-black font-bold text-2xl -translate-x-1/2 -translate-y-1/2 -rotate-3;
    top: calc(50% - 4px)
}

</style>