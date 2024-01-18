<script setup lang="ts">

import { ref } from 'vue';
import { getMyRegistrations, getOnTrackStatus, getStatus } from '../client';
import { Registration, Status } from '../domain';
import StatusVue from '../components/status.vue';
import OnTrack from '../components/on-track.vue';
import MyRegistrations from '../components/my-registrations.vue';
import AdditionalInformation from '../components/additional-information.vue';
import AddRegistrationButton from '../components/add-registration-button.vue';

const myRegistrations = ref<Registration[]>([])
const status = ref<Status>()
const onTrack = ref<Status>()


async function loadData() {
  const promises = [getMyRegistrations(), getStatus(), getOnTrackStatus()] as const

  [myRegistrations.value, status.value, onTrack.value] = await Promise.all(promises) 
}

loadData()

</script>

<template>
  <div class="font-bold">
    <div class="p-6">
      <img class="mx-auto mb-3" src="../assets/logo.svg" />
      <h1 class="text-white text-center text-5xl italic">Sezon 2!</h1>
    </div>
    <StatusVue v-if="status" :status="status"/>
    <OnTrack v-if="onTrack" :status="onTrack"/>
    <div class="px-5 py-12">
      <MyRegistrations :registrations="myRegistrations" class="mb-3"/>
      <AdditionalInformation class="mb-3" />
      <a href="/logout" >
        <div class="bg-white text-black text-center w-full py-4 text-xl rounded-lg">
          Wyloguj
        </div>
      </a>
    </div>
    <RouterLink to="/register">
      <AddRegistrationButton class="fixed right-5 bottom-5"/>
    </RouterLink>
    
  </div>
</template>