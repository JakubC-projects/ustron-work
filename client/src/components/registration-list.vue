<script setup lang="ts">
import { PropType } from 'vue';
import { Registration } from '../domain';

defineProps({
    registrations: {type: Array as PropType<Registration[]>, default: []}
})
</script>

<template>
    <div>
        <h3>My registrations</h3>
        <div v-for="registration of registrations">{{ registration }}</div>
    </div>
</template>