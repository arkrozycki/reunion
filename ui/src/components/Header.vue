<script setup lang="ts">
import { useRoute } from 'vue-router'
import { Popover, PopoverButton, PopoverPanel } from '@headlessui/vue'
import {
  ChevronDownIcon,
  UserCircleIcon,
  LoginIcon,
} from '@heroicons/vue/solid'

import {
  ChartBarIcon,
  CursorClickIcon,
  PhoneIcon,
  PlayIcon,
  RefreshIcon,
  ShieldCheckIcon,
  ViewGridIcon,
} from '@heroicons/vue/outline'
const route = useRoute()
const storeId = route.params.store
</script>
<template>
</template>
