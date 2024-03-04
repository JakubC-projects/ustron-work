<template>
    <StatusVue :status="status"/>
    <OnTrack :status="onTrack"/>
    <OnTrackGender :status="onTrackGender"/>
    <div class="px-5 py-12">
      <MyRegistrations :registrations="myRegistrations" class="mb-3"/>
    </div>
</template>

<script setup lang="ts">
import StatusVue from '../components/status.vue';
import OnTrack from '../components/on-track.vue';
import MyRegistrations from '../components/my-registrations.vue';
import OnTrackGender from '../components/on-track-gender.vue';
import { PropType, ref } from 'vue';
import { GenderStatus, Registration, Round, Status, newGenderStatus, newStatus } from '../domain';
import { getMyRegistrations, getOnTrackGenderStatus, getOnTrackStatus, getStatus } from '../client';

const myRegistrations = ref<Registration[]>([])
const status = ref<Status>(newStatus())
const onTrack = ref<Status>(newStatus())
const onTrackGender = ref<GenderStatus>(newGenderStatus())

const props = defineProps({
  round: {type: Object as PropType<Round>, required: true}
})

async function load() {
  const roundId= props.round.id
  const [resRegs, resStatus, resOnTrack, resOnTrackGender] =
    await Promise.all([getMyRegistrations(roundId), getStatus(roundId), getOnTrackStatus(roundId), getOnTrackGenderStatus(roundId)])

  myRegistrations.value = resRegs
  status.value = resStatus
  onTrack.value = resOnTrack
  onTrackGender.value = resOnTrackGender
}

load()


</script>