<template>
    <main class="flex flex-col justify-left items-left gap-8 py-12 px-8">
        <div v-if="data" class="q-pa-md row items-start q-gutter-md">
            <q-input
                ref="filterRef"
                v-model="filter"
                :onchange="filt"
                style='display:inline-block;width:320px;'
                label="Filter by device ID"
            >
                <template v-slot:append>
                    <q-icon v-if="filter !== ''" name="clear" class="cursor-pointer" @click="resetFilter" />
                </template>
            </q-input>
            <q-select
                style='display:inline-block;width:320px;margin-left:10px'
                ref="filterErrorRef"
                @update:model-value="filtError"
                :options="['none', 'hasError', 'connectionFault']"
                v-model="filterError"
                label="Filter by errors"
            >
            </q-select>
                <q-toggle
                    v-model="autoload"
                    label="Enable autoload"
                    @update:model-value="toggleAutoload"
                />
                
            
            <q-linear-progress v-if="loading" query />
            <div v-else style="margin:10px;width:100%"></div>
            <q-card v-for="item in dataFilter.deviceList" 
                v-ripple
                :class="(item.hasError == 'Y' || item.connectionFault == 'Y' ? 'bg-red-6' : item.type == 'plc' ? 'bg-teal-6' : 'bg-green-6')" 
                class="my-card text-white " 
                style="min-width:320px;min-height:100px">
                <q-card-section>
                    <div class="text-h6"></div>
                    <div class="text-subtitle2"></div>
                </q-card-section>
                <!-- span class="q-focus-helper"></span-->
                <q-expansion-item
                    dark
                    expand-separator
                    :icon="item.hasError == 'Y' || item.connectionFault == 'Y' ? 'warning' : 'devices'"
                    :label="item.id "
                    :caption="item.name || 'no-name'"
                >
                <q-card-section style="height:65%" 
                    :class="(item.hasError == 'Y' || item.connectionFault == 'Y' ? 'bg-red-8' : item.type == 'plc' ? 'bg-teal-8' : 'bg-green-8')">
                  <q-list dense  padding q-m-s class="rounded-borders">
                      <q-item> Type: {{ item.type }}</q-item>
                      <q-item> Location: {{ item.location }}</q-item>
                      <q-item> Status ID: {{ item.statusId }} </q-item>
                      <q-item> Flow ID: {{ item.statusFlowId }}</q-item>
                      <q-item> Connection: {{ item.connectionFault }}</q-item>
                      <q-item> Has errors: {{ item.hasError }}</q-item>

                  </q-list>
                </q-card-section>
                    <q-separator dark />
                    
                    <q-card-actions class="row justify-center">                    
                        <q-btn flat :to="'/device/'+item.id+'/dashboard'">View</q-btn>

                    </q-card-actions>
                </q-expansion-item>             
               
            </q-card>
        </div>
    </main>
</template>

<script lang="ts">
import ApiService from '../services/api.service'

import { ref } from 'vue';

type Device = {
    type: string
    id: number
    name: string
    location: string
    statusId: string
    statusFlowId: string
    hasError: string
    connectionFault: string
}

type DeviceList = {
    list: Array<Device>
}

type PostPersonName = {
    name: string
}

export default {
    setup() {

        let data = ref<Array<Device> | null>(null)
        const dataFilter = ref<Array<Device> | null>(null)
        let loading = false
        const name = <string>("fetching..")
        const filterError = ref('none')
        const filter: string = ''
        const autoload = ref(false)
        const env = import.meta.env
        let intervalHandle
        return {
            data,
            dataFilter,
            name,
            filter,
            filterError,
            loading,
            autoload,
            intervalHandle,
            fetchData: async function () {
                
                const prom = await ApiService.get(env.VITE_APP_MOQUI_API_ENDPOINT+'/deviceList')

                console.log(prom)
                //const prom = await fetch(env.VITE_APP_MOQUI_API_ENDPOINT+"/deviceList")
                const res: Awaited <DeviceList> = await prom.data
                console.log(res)
                
                dataFilter.value = res
                data.value = {...dataFilter.value} // Effettua una copia
                //name.value = res.name
            },
        }
    },
    methods: {
        filt: function(e){
            console.log(e.target.value)
            let filter : string = e.target.value

            this.filterError = 'none'
            if(filter == '' || filter == '*'){
                this.fetchData()
                return
            }
            
            let deviceList : Array<Device> = []
            let device : Device;
            
            for ( const i  in this.data.deviceList) {
                device = this.data.deviceList[i]
                console.log(device.id)
                if ( (device.id as string).indexOf(filter) >= 0 ) deviceList.push(device)
            }
            this.dataFilter.deviceList = deviceList
        },
        resetFilter: function(){
            this.fetchData()
            this.filter = ''
        },
        filtError: function(value){
            
            this.filter = ''
            let filter : string = value

            if(filter == 'none' ){
                this.fetchData()
                return
            }
            
            let deviceList : Array<Device> = []
            let device : Device;
            
            for ( const i  in this.data.deviceList) {
                device = this.data.deviceList[i]
                console.log(device.id)

                if ( filter == 'hasError' && device.hasError ) deviceList.push(device)
                if ( filter == 'connectionFault' && device.connectionFault ) deviceList.push(device)

            }
            this.dataFilter.deviceList = deviceList
        },
        toggleAutoload: function(){
            this.autoload = this.autoload!
            console.log(this.autoload)
            if(this.autoload) {
                this.loading = true
                this.intervalHandle = setInterval(() => {
                    this.loading = !this.loading
                    this.fetchData()
                }    
                , 5000)
            } else {
                clearInterval(this.intervalHandle)
                this.loading = false
            }
        },
    },
    mounted: function(){
        this.fetchData()
    }
}
</script>
