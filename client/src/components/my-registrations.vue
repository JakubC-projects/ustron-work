<script setup lang="ts">
import { PropType, computed } from 'vue';
import { Registration, RegistrationType } from '../domain';
import Dropdown from './dropdown.vue';

const props = defineProps({
    registrations: {type: Array as PropType<Registration[]>, default: []}
})

type MonthRegistrations = {
    month: number
    year: number
    registrations: Registration[]
}

const registrationsByMonth = computed<MonthRegistrations[]>(() => {
    const monthsMap = new Map<string, Registration[]>()
    for(const reg of props.registrations) {
        const workDate = new Date(reg.date)
        const month = `${workDate.getMonth()}-${workDate.getFullYear()}`
        
        const monthRegistration = monthsMap.get(month) ?? []
        monthRegistration.push(reg)
        monthsMap.set(month, monthRegistration)
    }

    const monthRegistrations: MonthRegistrations[] = []

    for(const [monthRaw, registrations] of monthsMap) {
        const [monthStr, yearStr] = monthRaw.split("-")
        monthRegistrations.push({
            month: parseInt(monthStr),
            year: parseInt(yearStr),
            registrations
        })
    }

    return monthRegistrations
})

const months = [
    "Styczeń",
    "Luty",
    "Marzec",
    "Kwiecień",
    "Maj",
    "Czerwiec",
    "Lipiec",
    "Sierpień",
    "Wrzesień",
    "Październik",
    "Listopad",
    "Grudzień",
]

function formatDate(date: string):string {
    const d = new Date(date)
    return `${d.getDate()}. ${months[d.getMonth()]}`
}

</script>

<template>
    <Dropdown label="Mój wkład">
        <div class="px-8 pt-2 pb-6 text-base font-normal flex flex-col gap-8">
            <div v-for="{year, month, registrations} of registrationsByMonth">
                <div class="flex justify-between opacity-60 mb-3 ">
                    <p>{{ months[month].toUpperCase() }}</p>
                    <p>{{ year }}</p>
                </div>
                <div v-for="reg of registrations" class="py-2 border-b border-white">
                    <div class="flex justify-between font-bold gap-2">
                        <div class="text-ellipsis overflow-hidden whitespace-nowrap min-w-0">{{ reg.type === RegistrationType.Work ? reg.description : "Wpłata" }}</div>
                        <div>{{(reg.hourlyWage * reg.hours + reg.paidSum).toFixed(2)}}</div>
                    </div>
                    <p>{{ formatDate(reg.date) }}</p>
                </div>
            </div>
        </div>
    </Dropdown>
</template>