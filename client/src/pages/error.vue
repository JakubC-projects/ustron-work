<script setup lang="ts">
import { useRoute } from 'vue-router';


const route = useRoute()
</script>

<template>
    <div>
        <h1>A niech to!</h1>
        <p>{{ route.query.error }}</p>
    </div>
</template>