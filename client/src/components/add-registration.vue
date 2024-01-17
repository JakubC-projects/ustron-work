<script setup lang="ts">
import { CreateRegistration, RegistrationType } from '../domain';

const model = defineModel<CreateRegistration>({default: undefined})

defineEmits(["confirm"])
</script>

<template>
    <div class="w-full bg-black fixed bottom-0 left-0 rounded-t-xl">
        <div class="pb-48">
            <h3>New Registration</h3>
            {{ model }}
            <label>Registration Type: </label>
            <select v-model="model.type">
                <option :value="RegistrationType.Work">Work</option>
                <option :value="RegistrationType.Money">Money</option>
            </select>
            <br>
            <template v-if="model.type === RegistrationType.Work">
                <label>Hourly wage: </label>
                <input type="number" v-model="model.hourlyWage" />
                <br>

                <label>Hours: </label>
                <input type="number" v-model="model.hours" />
                <br>
            </template>

            <template v-if="model.type === RegistrationType.Money">

                <label>Paid amount: </label>
                <input type="number" v-model="model.paidSum" />
            </template>
        </div>
        <button @click="$emit('confirm')">Confirm</button>
    </div>
</template>