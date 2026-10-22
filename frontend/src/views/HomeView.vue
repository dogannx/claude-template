<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { api } from '../api/client'

const status = ref<string>('…')

onMounted(async () => {
  try {
    const res = await api<{ status: string }>('/api/v1/health')
    status.value = res.status
  } catch {
    status.value = 'unreachable'
  }
})
</script>

<template>
  <section>
    <p>Backend health: {{ status }}</p>
  </section>
</template>
