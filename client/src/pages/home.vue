<script setup lang="ts">
import StatusVue from '../components/status.vue';
import OnTrack from '../components/on-track.vue';
import MyRegistrations from '../components/my-registrations.vue';
import AddRegistrationButton from '../components/add-registration-button.vue';
import {useStore} from '../store/store'
import OnTrackGender from '../components/on-track-gender.vue';
import Rules from '../components/rules.vue';

const state = useStore()

state.loadAll()

// function reload() {
//   state.status = newStatus();
//   state.onTrack = newStatus();
//   setTimeout(state.loadAll, 1000)
// }

</script>

<template>
  <div class="">
    <div class="p-6">
      <img class="mx-auto mb-3" src="../assets/logo.svg" />
      <h1 class="text-white font-bold text-center text-5xl italic">Sezon 2!</h1>
    </div>
    <StatusVue :status="state.status"/>
    <OnTrack :status="state.onTrack"/>
    <OnTrackGender :status="state.onTrackGender"/>
    <div class="px-5 py-12">
      <MyRegistrations :registrations="state.myRegistrations" class="mb-3"/>
      <Rules class="mb-3" />
      <a href="/logout" >
        <div class="bg-white text-black text-center w-full py-4 mb-3 text-xl rounded-lg font-bold">
          Wyloguj
        </div>
      </a>
      
      <!-- <div class="bg-white text-black text-center w-full py-4 mb-3 text-xl rounded-lg" @click="reload">
          Reload
        </div> -->
    </div>
    <RouterLink to="/register">
      <AddRegistrationButton class="fixed right-5 bottom-5 hover:rotate-90 transition-transform"/>
    </RouterLink>
  </div>
</template>