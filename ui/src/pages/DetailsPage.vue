<template>
  <div class="q-pa-md" style="max-width: 400px">

    <q-card>
      <q-card-section>
<q-form
  @submit="onSubmit"
  class="q-gutter-md"
>
  <q-input
    filled
    v-model="deviceId"
    label="Device Id"
    hint=""
    lazy-rules
    :rules="[ val => val && val.length > 0 || 'Please type something']"
  />

  <q-input
    filled
    v-model="deviceId"
    label="Pseudo Id"
    hint=""
    lazy-rules
    :rules="[ val => val && val.length > 0 || 'Please type something']"
  />
  <q-input
    filled
    v-model="deviceId"
    label="Name"
    hint=""
    lazy-rules
    :rules="[ val => val && val.length > 0 || 'Please type something']"
  />
  <q-input
    filled
    v-model="deviceId"
    label="Device type"
    hint=""
    lazy-rules
    :rules="[ val => val && val.length > 0 || 'Please type something']"
  />
  <q-input
    filled
    v-model="deviceId"
    label="Hardware version"
    hint=""
    lazy-rules
    :rules="[ val => val && val.length > 0 || 'Please type something']"
  />

  <div>
    <q-btn label="Submit" type="submit" color="primary"/>
  </div>
</q-form>
</q-card-section>
</q-card>
</div>  
</template>
  
<script lang="ts">
import router from '../main'
import { ref } from 'vue'
  
  export default {
    data (){
      let deviceId
      return {
        deviceId
      }
    },
    setup () {


        return {
        }
    },
    methods:{
        
    }
    
  }
  </script>
  
  