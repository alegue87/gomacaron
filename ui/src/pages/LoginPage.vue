<template>
    <q-layout view="hHh lpR fFf">
  
        <div class="bg-white window-height window-width row justify-center items-center">
          <div class="column">
              <div class="row">
                  <q-card square bordered class="q-pa-lg shadow-1">
                    <div class="row justify-center items-center"><h6>Login moqui</h6></div>
                    <div v-if="authenticating">
                        <q-linear-progress indeterminate />
                    </div>
                    <q-card-section>
                        <q-form class="q-gutter-md">
                        <q-input square filled clearable v-model="username" type="text" label="Username" />
                        <q-input square filled clearable v-model="password" type="password" label="Password" @keydown.enter.prevent="handleSubmit"/>
                        </q-form>
                    </q-card-section>
                    <q-card-actions class="q-px-md">
                        <q-btn unelevated color="light-blue-7" size="lg" class="full-width" label="Login" @click="handleSubmit"/>
                    </q-card-actions>

                    <div class="row justify-center" v-if="authenticationError">Login error: {{ authenticationErrorCode }}</div>
                  </q-card>
              </div>
          </div>
        </div>

    </q-layout>
  </template>
  
<script lang="ts">
import { useAuthStore } from '@/store'
import { mapActions } from 'pinia'
import { computed } from 'vue'

export default {
    data (){
        let password: string,
            username: string
        return {
            username,
            password
        }
    },
    setup () {
        const store = useAuthStore()

        return {
            authenticating: computed( () => store._authenticating),
            authenticationError: computed( () => store._authenticationError),
            authenticationErrorCode: computed( () => store._authenticationErrorCode)
        }
    },
    computed: {
    },
    methods:{
        ...mapActions(useAuthStore, [
            'login',
        ]),

        handleSubmit() {
            // Perform a simple validation that email and password have been typed in
            if (this.username != '' && this.password != '') {
                this.login(this.username, this.password)
                this.password = ""
            }
        }

    }
}
</script>

