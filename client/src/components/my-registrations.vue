<script setup lang="ts">
import { PropType } from 'vue';
import { Registration } from '../domain';
import Dropdown from './dropdown.vue';

defineProps({
    registrations: {type: Array as PropType<Registration[]>, default: []}
})

</script>

<template>
    <Dropdown label="Mój wkład">
        <div class="px-8 py-6 ">
            <div v-for="registration of registrations">{{ registration }}</div>
        </div>
    </Dropdown>
</template>