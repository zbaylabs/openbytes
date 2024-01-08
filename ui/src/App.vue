<template>
  <q-layout view="hHh lpR fFf">

    <q-header bordered class="bg-primary text-white">
      <q-toolbar style="height: 60px;">
        <q-btn dense flat round icon="menu" @click="toggleLeftDrawer" />

        <q-toolbar-title>
          <q-avatar>
            <img src="https://cdn.quasar.dev/logo-v2/svg/logo-mono-white.svg">
          </q-avatar>
          Openbytes
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-drawer show-if-above v-model="leftDrawerOpen" side="left" :width="230">
      <!-- drawer content -->
      <div class="q-pa-md">
        <q-list>
          <q-expansion-item group="somegroup" label="Overview" default-opened header-class="text-primary">
            <q-card>
              <q-card-section>
                <router-link to="/" style="padding-left: 5px;">Traffic</router-link>
              </q-card-section>
            </q-card>
          </q-expansion-item>

          <q-separator />

          <q-expansion-item group="somegroup" label="Capture" header-class="text-teal">
            <q-card>
              <q-card-section>
                <router-link to="/packet" style="padding-left: 5px;">Packets</router-link>
              </q-card-section>
            </q-card>
          </q-expansion-item>

          <q-separator />

          <q-expansion-item group="somegroup" label="Replay" header-class="text-purple">
            <q-card>
              <q-card-section>
                <router-link to="/mirror" style="padding-left: 5px;">Mirror</router-link>
              </q-card-section>
            </q-card>
          </q-expansion-item>

          <q-separator />

          <!-- <q-expansion-item
        group="somegroup"
        icon="bluetooth"
        label="Fourth"
        header-class=""
        expand-icon-class="text-white"
      >
        <q-card class="">
          <q-card-section>
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Quidem, eius reprehenderit eos corrupti
            commodi magni quaerat ex numquam, dolorum officiis modi facere maiores architecto suscipit iste
            eveniet doloribus ullam aliquid.
          </q-card-section>
        </q-card>
      </q-expansion-item> -->
        </q-list>
      </div>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>

  </q-layout>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const leftDrawerOpen = ref(false)

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value
}

// export default {
//   setup () {
//     const leftDrawerOpen = ref(false)

//     return {
//       leftDrawerOpen,
//       toggleLeftDrawer () {
//         leftDrawerOpen.value = !leftDrawerOpen.value
//       }
//     }
//   }
// }
</script>

<style>
.my-menu-link {
  color: white;
  background: #F2C037
}
</style>