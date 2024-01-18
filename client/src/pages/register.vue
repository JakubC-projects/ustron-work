<script setup lang="ts">
import { ref } from 'vue';
import { CreateRegistration, RegistrationType, Goal } from '../domain';
import { createMyRegistration } from '../client';
import { useRouter } from 'vue-router';
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
    comment: "",
})

const router = useRouter()

async function createNewRegistration() {

  await createMyRegistration(registration.value)
  
  router.push({path: "/"})
}

</script>

<template>
    <div>
        <router-link to="/" class="p-6">
            <img class="mx-auto mb-3" src="../assets/logo.svg" />
        </router-link>
        <Rules />
        <div class="px-5 py-6">
            <p class="input-label">Dodaj</p>
            <div class="grid grid-cols-2 gap-4">
                <InputRadio v-model="registration.type" :value="RegistrationType.Work" label="Praca"/>
                <InputRadio v-model="registration.type" :value="RegistrationType.Money" label="Wpłata"/>
            </div>
        </div>
        <div class="px-5 py-6">
            <p class="input-label">Data</p>
            <input type="date" v-model="registration.date">
        </div>
        <template v-if="registration.type == RegistrationType.Work">
            <div class="px-5 py-6">
                <p class="input-label">Ilość godzin (bez przerw)</p>
                <input type="number" v-model="registration.hours">
            </div>
            <div class="px-5 py-6">
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
            <div class="px-5 py-6">
                <p class="input-label">Kwota wpłaty</p>
                <input type="number" v-model="registration.paidSum">
                <p class="opacity-75 text-base pt-2">*Sama liczba bez przecinków w PLN</p>
            </div>
        </template>

        <div class="px-5 py-6">
            <p class="input-label">Cel wpłaty</p>
            <div class="grid grid-cols-2 gap-4">
                <InputRadio v-model="registration.goal" :value="Goal.BUK" label="BUK"/>
                <InputRadio v-model="registration.goal" :value="Goal.Samvirk" label="Samvirk"/>
            </div>
        </div>
        <div class="px-5 py-6 ">
            <p class="input-label">Komentarz</p>
            <input v-model="registration.comment">
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