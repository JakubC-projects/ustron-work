<script setup lang="ts">
import { ref } from 'vue';
import { createMyRegistration, getMe, getMyRegistrations } from './client';
import { CreateRegistration, Registration, RegistrationType, User } from './domain';
import RegistrationForm from './components/registration-form.vue';
import RegistrationList from './components/registration-list.vue';

const me = ref<User>()
const newRegistration = ref<CreateRegistration>()
const myRegistrations = ref<Registration[]>([])

async function load() {
  me.value = await getMe()
  myRegistrations.value = await getMyRegistrations()
}

load()

function startNewRegistration() {
  newRegistration.value = {
    type:RegistrationType.Work,
    hourlyWage: 0,
    hours: 0,
    paidSum: 0
  }
}

async function createNewRegistration() {
  if(!newRegistration.value) {
    return
  }
  await createMyRegistration(newRegistration.value)
  myRegistrations.value = await getMyRegistrations()
  newRegistration.value = undefined
}

</script>

<template>
  <div>
    <div>Hello  {{ me?.displayName }}
    </div>

    <a href="/logout">Logout</a><br>

    <button @click="startNewRegistration">New Registration</button>

    <div  v-if="newRegistration">
      <RegistrationForm v-model="newRegistration" />
      <button @click="createNewRegistration">Create</button>
    </div>
    <RegistrationList :registrations="myRegistrations" />
  </div>
</template>
