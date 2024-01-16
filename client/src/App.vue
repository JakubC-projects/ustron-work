<script setup lang="ts">
import { ref } from 'vue';
import {  getMe, getMyRegistrations } from './client';
import {  Registration,   User } from './domain';
import Status from './components/status.vue';
import OnTrack from './components/on-track.vue';

const me = ref<User>()
// const newRegistration = ref<CreateRegistration>()
const myRegistrations = ref<Registration[]>([])

async function load() {
  me.value = await getMe()
  myRegistrations.value = await getMyRegistrations()
}

load()

// function startNewRegistration() {
//   newRegistration.value = {
//     type:RegistrationType.Work,
//     hourlyWage: 0,
//     hours: 0,
//     paidSum: 0
//   }
// }

// async function createNewRegistration() {
//   if(!newRegistration.value) {
//     return
//   }
//   await createMyRegistration(newRegistration.value)
//   myRegistrations.value = await getMyRegistrations()
//   newRegistration.value = undefined
// }

</script>

<template>
  <div class="max-w-sm  mx-auto">
    <div class="p-6">
      <img class="mx-auto mb-3" src="./assets/logo.svg" />
      <h1 class="text-white text-center text-5xl italic">Sezon 2!</h1>
    </div>
    <Status :status="{Blue: 100, Red: 200, Green: 300, Orange: 400}"/>
    <OnTrack :status="{Blue: 0.3, Red: 0.8, Green: 0.1, Orange: 0.2}"/>
  </div>
</template>

<style>

body{
    color: #fff;
    font-family: 'Poppins', sans-serif;
    background: #1E302D;
}

</style>