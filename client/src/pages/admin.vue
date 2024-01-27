<script setup lang="ts">
import SetOnTrackGender from '../components/set-on-track-gender.vue';
import SetOnTrack from '../components/set-on-track.vue';
import { useStore } from '../store/store';

const state = useStore()

state.loadAll()

</script>

<template>
    <h1 class="text-3xl m-6 text-center font-bold">Admin Panel</h1>
    <SetOnTrack @updated="state.loadAll" :status="state.onTrack" :key="JSON.stringify(state.onTrack)"/>
    <SetOnTrackGender @updated="state.loadAll" :status="state.onTrackGender" :key="JSON.stringify(state.onTrackGender)"/>
</template>