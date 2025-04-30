<template>
   <iframe :src="grafana_host+ '/d/b1a1d34c-622a-405e-ac2d-ec81bf1c1d74/new-dashboard-pg?kiosk=true&refresh=5s'" style="height:100vh;width:100vw"></iframe>
</template>
  
<script>

import {defineComponent, ref} from 'vue'

export default defineComponent({
    name: 'DashboardPage',
  
    components: {
    },
  
    setup() {
      const env = import.meta.env

      return {
        grafana_host: env.VITE_APP_GRAFANA_HOST
      }
    }

    
  })
</script>
  
  