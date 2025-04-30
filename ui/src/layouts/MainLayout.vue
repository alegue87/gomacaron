<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          @click="toggleLeftDrawer"
          icon="menu"
          aria-label="Menu"
        />
        <q-toolbar-title>
          {{ title }}<q-badge>v {{ version }}</q-badge>
        </q-toolbar-title>
        <q-space/>
        <div class="q-gutter-sm row items-center no-wrap">
          <span style="padding-top:5px">{{ username }}</span>      
          <q-btn round dense flat icon="account_circle"></q-btn>
          <q-btn round dense flat color="white" icon="logout" type="a" to="/login" @click="logout"></q-btn>
        </div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
      class="bg-primary text-white"
    >
      <q-list>
        <q-item to="/" active-class="q-item-no-link-highlighting">
          <q-item-section avatar>
            <q-icon name="devices"/>
          </q-item-section>
          <q-item-section>
            <q-item-label>Devices Overview</q-item-label>
          </q-item-section>
        </q-item>

        <q-item to="/scheduledjobs" active-class="q-item-no-link-highlighting">
          <q-item-section avatar>
            <q-icon name="event"/>
          </q-item-section>
          <q-item-section>
            <q-item-label>Scheduled jobs</q-item-label>
          </q-item-section>
        </q-item>
    
        <q-item to="/mermaid" active-class="q-item-no-link-highlighting">
          <q-item-section avatar>
            <q-icon name="mediation"/>
          </q-item-section>
          <q-item-section>
            <q-item-label>Mermaid flow editor</q-item-label>
          </q-item-section>
        </q-item>

        <q-item to="/univer" active-class="q-item-no-link-highlighting">
          <q-item-section avatar>
            <q-icon name="functions"/>
          </q-item-section>
          <q-item-section>
            <q-item-label>Univer editor</q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>

    <q-page-container class="bg-grey-2">
      <router-view/>
    </q-page-container>
  </q-layout>
</template>

<script>

  import {defineComponent, ref, computed} from 'vue'
  import { mapActions, storeToRefs } from 'pinia';
  import { useAuthStore } from '@/store';

  export default defineComponent({
    name: 'MainLayout',

    components: {
    },

    setup() {

      const store = useAuthStore()
      const leftDrawerOpen = ref(false)
      const env = import.meta.env

      return {
        username: computed( () => store._username ),
        title: env.VITE_APP_TITLE,
        version: env.VITE_APP_VERSION,
        leftDrawerOpen,
        toggleLeftDrawer() {
          leftDrawerOpen.value = !leftDrawerOpen.value
        }
      }
    },
    methods:{
      ...mapActions(useAuthStore, [
            'logout',
        ]),
    },

    mounted(){
    }
})
</script>

<style>

/* FONT AWESOME GENERIC BEAT */
.fa-beat {
  animation: fa-beat 5s ease infinite;
}

@keyframes fa-beat {
  0% {
    transform: scale(1);
  }
  5% {
    transform: scale(1.25);
  }
  20% {
    transform: scale(1);
  }
  30% {
    transform: scale(1);
  }
  35% {
    transform: scale(1.25);
  }
  50% {
    transform: scale(1);
  }
  55% {
    transform: scale(1.25);
  }
  70% {
    transform: scale(1);
  }
}

</style>
