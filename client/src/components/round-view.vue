<template>
    <div class="text-center text-base mx-auto opacity-75 pb-4">{{roundStatus}}</div>
    <StatusVue :round="round" :status="status" :user-team="me.team"/>
    <OnTrack :status="onTrack"/>
    <OnTrackGender :status="onTrackGender"/>
    <div class="px-5 pt-12">
      <MyRegistrations :registrations="myRegistrations" class="mb-3"/>
    </div>
</template>

<script setup lang="ts">
import StatusVue from '../components/status.vue';
import OnTrack from '../components/on-track.vue';
import MyRegistrations from '../components/my-registrations.vue';
import OnTrackGender from '../components/on-track-gender.vue';
import { PropType, computed, ref, watchEffect } from 'vue';
import { GenderStatus, Person, Registration, Round, Status, newGenderStatus, newStatus } from '../domain';
import { getMyRegistrations, getOnTrackGenderStatus, getOnTrackStatus, getStatus } from '../client';

const myRegistrations = ref<Registration[]>([])
const status = ref<Status>(newStatus())
const onTrack = ref<Status>(newStatus())
const onTrackGender = ref<GenderStatus>(newGenderStatus())

const props = defineProps({
  round: {type: Object as PropType<Round>, required: true},
  me: {type: Object as PropType<Person>, required: true}
})

async function load(roundId: number) {
  const [resRegs, resStatus, resOnTrack, resOnTrackGender] =
    await Promise.all([getMyRegistrations(roundId), getStatus(roundId), getOnTrackStatus(roundId), getOnTrackGenderStatus(roundId)])

  myRegistrations.value = resRegs
  status.value = resStatus
  onTrack.value = resOnTrack
  onTrackGender.value = resOnTrackGender
}

watchEffect(() => {
  load(props.round.id)
})

const roundStatus = computed(() => {
  const now = new Date().toISOString()
  if (now < props.round.startDate) {
    return `Runda ${props.round.id} zaczyna się ${formatDate(props.round.startDate)}`
  }
  if (now < props.round.endDate) {
    return `Runda ${props.round.id} kończy się ${formatDate(props.round.endDate)}`
  }
  return ""
})

function formatDate(date: string) {
  return new Date(date).toLocaleString("pl", {dateStyle:  'long', timeStyle:'medium'})
}

</script>