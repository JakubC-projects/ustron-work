<script setup lang="ts">
import { PropType, computed, ref } from 'vue';
import {Status, Team} from '../domain'
import {setOnTrackStatus} from '../client'
import SetOnTrackTeam from './set-on-track-team.vue';

const props = defineProps({
    status: {type: Object as PropType<Status>, required: true}
})

const statusEditable= ref<Status>(copyStatus())

function copyStatus() {
    return JSON.parse(JSON.stringify(props.status))
}

const emit = defineEmits(["updated"])

async function save() {
    await setOnTrackStatus(statusEditable.value)
    emit("updated")
}

const isChanged = computed(() => {
    return JSON.stringify(statusEditable.value) != JSON.stringify(props.status)
})

</script>

<template>
    <div class="px-5 py-8">
        <h3 class="text-center text-2xl mb-6">Procent bycia <span class="text-[#FF0] italic">ON TRACK</span></h3>
        <div v-for="team in Team" class="mb-3">
            <SetOnTrackTeam :team="team" v-model="statusEditable[team]"/>
        </div>
        <div class="flex gap-4"  :class="isChanged ? '' : 'opacity-50'">
            <button  @click="statusEditable = copyStatus()" :disabled="!isChanged" class="bg-white text-black text-center w-full py-4 text-xl rounded-lg">
            Anuluj
            </button>
            <button @click="save" :disabled="!isChanged" class="bg-white text-black text-center w-full py-4 text-xl rounded-lg">
            Zapisz
            </button>
        </div>
    </div>
</template>
