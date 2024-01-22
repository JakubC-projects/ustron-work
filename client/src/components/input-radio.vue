<script setup lang="ts">
defineProps({
    value: {type: String, required: true},
    label: {type: String, required: true},
    name: {type: String, required: false}
})

const model = defineModel({
    type: String,
    required: true,
})

</script>

<template>
    <label class="px-5 py-4 flex rounded-lg text-base items-center"
        :class="model == value ? 'bg-white': 'bg-[#21413E]'">
        <input class="w-8 h-8 accent-[#208070]" type="radio" :name="name" :value="value" v-model="model"/>
        <p class="w-full text-center"
            :class="model == value ? 'text-black': 'text-white'"
        >
            {{label}}
        </p>
    </label>
</template>