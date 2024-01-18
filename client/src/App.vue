<script setup lang="ts">
import { ref } from 'vue';
import {  createMyRegistration, getMe, getMyRegistrations, getOnTrackStatus, getStatus } from './client';
import {  CreateRegistration, Registration,   RegistrationType,   Status,   User } from './domain';
import StatusVue from './components/status.vue';
import OnTrack from './components/on-track.vue';
import AddRegistrationButton from './components/add-registration-button.vue';
import AddRegistration from './components/add-registration.vue';
import MyRegistrations from './components/my-registrations.vue';
import AdditionalInformation from './components/additional-information.vue';

const me = ref<User>()
const newRegistration = ref<CreateRegistration>()
const myRegistrations = ref<Registration[]>([])
const status = ref<Status>()
const onTrack = ref<Status>()


async function loadData() {
  const promises = [getMe(), getMyRegistrations(), getStatus(), getOnTrackStatus()] as const

  [me.value, myRegistrations.value, status.value, onTrack.value] = await Promise.all(promises) 
}

loadData()

function startNewRegistration() {
  newRegistration.value = {
    type:RegistrationType.Work,
    hourlyWage: 0,
    hours: 0,
    paidSum: 0
  }
}

function toggleNewRegistration() {
  if(newRegistration.value) {
    newRegistration.value = undefined
  } else {
    startNewRegistration()
  }
}

async function createNewRegistration() {
  if(!newRegistration.value) {
    return
  }
  await createMyRegistration(newRegistration.value)
  await loadData()

  newRegistration.value = undefined
}

</script>

<template>
  <div class="relative max-w-sm mx-auto">
    <div class="p-6">
      <img class="mx-auto mb-3" src="./assets/logo.svg" />
      <h1 class="text-white text-center text-5xl italic">Sezon 2!</h1>
    </div>
    <StatusVue v-if="status" :status="status"/>
    <OnTrack v-if="onTrack" :status="onTrack"/>
    <div class="px-5 py-6">
      <MyRegistrations :registrations="myRegistrations" class="mb-3"/>
      <AdditionalInformation class="mb-3" />
      <a href="/logout" >
        <div class="bg-white text-black text-center w-full py-4 text-xl rounded-lg">
          Wyloguj
        </div>
      </a>
    </div>
    <div class="h-48"></div>
    <AddRegistration v-if="newRegistration" v-model="newRegistration" @confirm="createNewRegistration"/>
    <AddRegistrationButton class="fixed right-5 bottom-5" :is-open="newRegistration != undefined" @click="toggleNewRegistration"/>
    

  </div>
</template>
