<script setup lang="ts">
import { ref } from 'vue';
import { getMe } from './client';
import { User } from './domain';

const me = ref<User>()

async function load() {
  me.value = await getMe()
}

load()

</script>

<template>
  <div>HelloWorld
    {{ me?.displayName }}
</div>
</template>
