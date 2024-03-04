<script setup lang="ts">
import { ref } from 'vue';
import AddRegistrationButton from '../components/add-registration-button.vue';
import roundSelector from '../components/round-selector.vue';
import Rules from '../components/rules.vue';
import { Round } from '../domain';
import { getRounds } from '../client';
import round from '../components/round.vue';

const rounds = ref<Round[]>([])
const currentRound = ref<Round>()

async function load() {
  rounds.value = await getRounds()

  currentRound.value = selectRound(rounds.value)
}

function selectRound(rounds: Round[]):Round {
  return rounds[0]
}

load()

</script>

<template>
  <div>
    <div class="p-6">
      <img class="mx-auto mb-3" src="../assets/logo.svg" />
      <h1 class="text-white font-bold text-center text-5xl italic">Sezon 2!</h1>
    </div>
    <div class="px-5 py-6">
      <RouterLink to="/register" >
        <AddRegistrationButton />
      </RouterLink>
    </div>

    <roundSelector :rounds="rounds" v-model="currentRound"/>
    
    <round :round="currentRound" />

    <Rules class="mb-7" />
      <a href="/logout" >
        <div class="border-2 text-center w-full py-4 mb-3 text-base rounded-lg font-bold">
          Wyloguj
        </div>
      </a>
    
  </div>
</template>