<template>
    <q-tabs
        style="margin-top:1px"
        v-model="tab"
        inline-label
        class="bg-primary text-white shadow-2"
    >
        <span style="margin-right:20%"><q-icon name='devices' style="margin-right:5px"/>DEVICE {{ $route.params.id }}</span>
        <q-route-tab name="dashboard" icon="dashboard" label="Dashboard"    :to="'/device/'+$route.params.id+'/dashboard'"/>
        <q-route-tab name="details" icon="article" label="Details"          :to="'/device/'+$route.params.id+'/details'"/>
        <q-route-tab name="parameters" icon="list" label="Parameters"       :to="'/device/'+$route.params.id+'/parameters'"/>
        <!--q-route-tab name="schedule" icon="event" label="Schedule"          :to="'/device/'+$route.params.id+'/schedule'"/-->
    </q-tabs>
    <q-page class="w-full">

        <router-view/>
    </q-page>
  </template>
  
<script lang="ts">
import { ref } from 'vue'

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

            function login(){
                
            }

            return {
                tab: ref('dashboard')
            }
        },
        methods:{
        }
        
    }
</script>
  
  