<script setup lang="ts">
import { ref, watchEffect } from 'vue';
import { CreateRegistration, RegistrationType, Goal } from '../domain';
import { createMyRegistration } from '../client';
import { useRouter } from 'vue-router';
import {useToast} from 'vue-toast-notification';
import InputRadio from '../components/input-radio.vue';
import ButtonTheme from '../components/button-theme.vue';
import Rules from '../components/rules.vue';

const registration = ref<CreateRegistration>({
    type:RegistrationType.Work,
    date: new Date().toISOString().split("T")[0],
    hourlyWage: 0,
    hours: 0,
    paidSum: 0,
    goal: Goal.BUK,
    description: "",
})

const router = useRouter()
const toast = useToast()

async function createNewRegistration() {
    try {
        await createMyRegistration(registration.value)
        toast.success("dodano rejestrację")
        router.push({path: "/"})
    } catch (err) {
        const msg = err instanceof Error ? err.message : 'Unknown error'
        toast.error(msg)
    }
}

watchEffect(() => {
    if(registration.value.type === RegistrationType.Work ) {
        registration.value.paidSum = 0
        registration.value.goal = Goal.BUK
    }
    if (registration.value.type === RegistrationType.Money ) {
        registration.value.hourlyWage = 0
        registration.value.hours = 0
        registration.value.goal = Goal.BUK
    }
})

</script>

<template>
    <div class="px-5">
        <router-link to="/" class="p-6">
            <img class="mx-auto mb-3" src="../assets/logo.svg" />
        </router-link>
        <Rules />
        <div class=" py-6">
            <p class="input-label">Dodaj</p>
            <div class="grid grid-cols-2 gap-4">
                <InputRadio name="type" v-model="registration.type" :value="RegistrationType.Work" label="Praca"/>
                <InputRadio name="type" v-model="registration.type" :value="RegistrationType.Money" label="Wpłata"/>
            </div>
        </div>
        <div class="py-6">
            <p class="input-label">Data</p>
            <input type="date" v-model="registration.date">
        </div>
        <template v-if="registration.type == RegistrationType.Work">
            <div class=" py-6">
                <p class="input-label">Ilość godzin (bez przerw)</p>
                <input type="number" v-model="registration.hours">
            </div>
            <div class=" py-6">
                <p class="input-label">Stawka godzinowa</p>
                <input type="number" v-model="registration.hourlyWage">
                <p class="opacity-75 text-base pt-2" >
                    *Standardowe prace w zborze:<br>
                    U18 - 20 zł/h<br>
                    O18 - 25 zł/h<br>
                    <span class="font-bold">Wyjątek: praca na zmywaku ma +5 zł/h</span>
                </p>
            </div>
        </template>
        <template v-else>
            <div class="py-6">
                <p class="input-label">Kwota wpłaty</p>
                <input type="number" v-model="registration.paidSum">
                <p class="opacity-75 text-base pt-2">*Sama liczba bez przecinków w PLN</p>
            </div>
        </template>

        <div class="py-6">
            <p class="input-label">Cel wpłaty</p>
            <div class="grid grid-cols-2 gap-4">
                <InputRadio name="goal" v-model="registration.goal" :value="Goal.BUK" label="BUK"/>
                <InputRadio v-if="registration.type === RegistrationType.Work" name="goal" v-model="registration.goal" :value="Goal.Maintenance" label="Utrzymanie Zboru"/>
                <InputRadio v-else name="goal" v-model="registration.goal" :value="Goal.Samvirk" label="Samvirk"/>
            </div>
        </div>
        <div class="py-6 ">
            <p class="input-label">{{registration.type == RegistrationType.Work ? "Opis Pracy":"Komentarz"}}</p>
            <input v-model="registration.description">
        </div>
        <div class="flex justify-center py-6">
            <ButtonTheme @click="createNewRegistration" label="Dodaj"/>
        </div>
    </div>

</template>

<style scoped>
.input-label {
    @apply font-bold text-lg pb-2
}
</style>