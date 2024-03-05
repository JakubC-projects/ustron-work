<script setup lang="ts">
import { ref } from 'vue';
import AddRegistrationButton from '../components/add-registration-button.vue';
import roundSelector from '../components/round-selector.vue';
import Rules from '../components/rules.vue';
import { Person, Round } from '../domain';
import { getMe, getRounds } from '../client';
import roundView from '../components/round-view.vue';

const rounds = ref<Round[]>([])
const me = ref<Person>()
const currentRound = ref<Round>()

async function load() {
  rounds.value = await getRounds()
  me.value = await getMe()

  currentRound.value = selectRound(rounds.value)
}

function selectRound(rounds: Round[]):Round {
  const now = new Date().toISOString()
  for(const round of rounds) {
    const finishShowingDate = new Date(round.endDate)
    finishShowingDate.setDate(finishShowingDate.getDate() + 1);

    if (now < finishShowingDate.toISOString()) {
      return round
    }
  }
  throw Error("no round found")
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

    <roundSelector v-if="currentRound" :rounds="rounds" v-model="currentRound" class="pb-2"/>
    <roundView v-if="currentRound && me" :round="currentRound" :me="me"/>

    <div class="px-5 pb-12">
      <Rules class="mb-7" />
      <a href="/logout" >
        <div class="border-2 text-center w-full py-4 mb-3 text-base rounded-lg font-bold">
          Wyloguj
        </div>
      </a>
    </div>
  </div>
</template>
