<template>
<router-view/>
</template>

<script lang="ts">
import { RouterLink, RouterView } from 'vue-router';
import { ref } from 'vue'

export default {
  data (){
    let password,username
    return {
      username,
      password
    }
  },
  setup () {
    const leftDrawerOpen = ref(false)

    function buttonProps (name, { href, route, isActive, isExactActive }) {
      const props = {
        color: 'black',
        noCaps: true,
        label: `${name}`,
        outline: true,
        to: href
      }

      if (isActive === true) {
        props.color = isExactActive === true ? 'primary' : 'amber-9'
      }
      else {
        props.color = 'black'
      }

      return props
    }
    return {
      leftDrawerOpen,
      buttonProps,
      toggleLeftDrawer () {
        leftDrawerOpen.value = !leftDrawerOpen.value
      }
    }
  },
  computed: {
    isAdmin(){

      if (!sessionStorage.getItem('isAdmin')){
        return true
      }
      else {
        return false
      }
    }
  }
}
</script>

