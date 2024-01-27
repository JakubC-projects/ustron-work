<script setup lang="ts">
import { PropType, computed, ref } from 'vue';
import {Gender, GenderStatus} from '../domain'
import { useToast } from 'vue-toast-notification';
import { setOnTrackGenderStatus } from '../client';

const props = defineProps({
    status: {type: Object as PropType<GenderStatus>, required: true}
})

const toast = useToast()

const statusEditable= ref<GenderStatus>(copyStatus())

function copyStatus() {
    return JSON.parse(JSON.stringify(props.status))
}

const emit = defineEmits(["updated"])

async function save() {
    try {
        await setOnTrackGenderStatus(statusEditable.value)
        emit("updated")
        toast.success("zapisano")
    } catch (err) {
        const msg = err instanceof Error ? err.message : "nieznany błąd"
        toast.error(msg)
    }
}

const isChanged = computed(() => {
    return JSON.stringify(statusEditable.value) != JSON.stringify(props.status)
})


</script>

<template>
    <div class="px-5 py-8">
        <h3 class="text-center text-2xl mb-6">Dziewczyny vs Chłopaki</h3>
        <div class="mb-3 flex gap-4 items-center">
            <img class="w-12" src="/src/assets/woman.png" />
            <input type="number" v-model="statusEditable[Gender.Female]"> 
        </div>
        <div class="mb-3 flex gap-4 items-center">
            <img class="w-12" src="/src/assets/man.png" />
            <input type="number" v-model="statusEditable[Gender.Male]"> 
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

<style>
.background {
    background-image: linear-gradient(110deg, #1E302D 0 50%, #193A36 50% 100%)
}
</style>
