<template>
  <div class="q-pa-md" style="max-width: 400px">

    <q-card>
      <q-card-section>
<q-form
  @submit="onSubmit"
  class="q-gutter-md"
>
  <q-input
    v-model="deviceDetail.deviceId"
    label="Device Id"
    hint=""
    readonly
  />

  <q-input
    v-model="deviceDetail.pseudoId"
    label="Pseudo Id"
    hint=""
  />
  <q-input
    v-model="deviceDetail.deviceName"
    label="Name"
    hint=""
  />
  <q-select
    v-model="deviceDetail.deviceTypeEnumId"
    label="Type"
    hint=""
    :options="typeOptions"
  />
  <q-select
    v-model="deviceDetail.purposeEnumId"
    label="Purpose"
    hint=""
    :options="purposeOptions"
  />
  <q-select
    v-model="deviceDetail.statusFlowId"
    label="Flow ID"
    hint=""
    :options="statusFlowOptions"
    @update:model-value="fetchStatusOptionsOnUpdate"
  />
  <q-select
    v-model="deviceDetail.statusId"
    label="Status ID"
    hint=""
    :options="statusOptions"
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
  import { ref } from 'vue'
  import ApiService from '@/services/api.service'

  type Detail = {
    deviceId: string
    pseudoId: string
    deviceName: string
    deviceTypeEnumId: string
    purposeEnumId: string
    firmwareVersion: string
    statusId: string
    statusFlowId: string
  }

  type StatusFlow = {
    statusFlowId: string
    description: string
  }

  type Status = {
    statusId: string
    description: string
  }

  type Enum = {
    enumId: string
    description: string
  }

  export default {
    data (){
      let deviceId, statusFlowId
      return {
        deviceId,
        statusFlowId
      }
    },
    setup () {
        const env = import.meta.env
        const deviceDetail = ref<Detail | null>({
          deviceId: '',
          pseudoId: '',
          deviceName: '',
          deviceTypeEnumId: '',
          purposeEnumId: '',
          firmwareVersion: '',
          statusId: '',
          statusFlowId: ''
        })
        const statusFlowList = ref<Array<StatusFlow> | null>([])
        const statusList = ref([])
        const statusFlowOptions  = ref<Array<Map<string, string>> | null>([]) 
        const typeList = ref([])
        const typeOptions  = ref<Array<Map<string, string>> | null>([]) 
        const purposeList = ref([])
        const purposeOptions  = ref<Array<Map<string, string>> | null>([]) 
        const statusFlowSelect = ref<Map<string, string>>({})
        const statusOptions  = ref<Array<Map<string, string>> | null>([]) 

        return {
          deviceDetail,
          statusFlowSelect,
          statusFlowList,
          statusFlowOptions,
          statusList,
          statusOptions,
          typeOptions,
          purposeOptions,
          fetchData: async function(){
            const prom = await ApiService.get(env.VITE_APP_MOQUI_API_ENDPOINT+'/deviceDetail?deviceId='+this.$route.params.id)
            const res: Awaited<Detail> = await prom.data.deviceDetail
            console.log(res)

            deviceDetail.value = res
            this.fetchStatusOptions(deviceDetail.value.statusFlowId)
          },
          fetchTypeOptions: async function(){
            const prom = await ApiService.get(env.VITE_APP_MOQUI_API_ENDPOINT+'/enumeration?enumTypeId=DeviceType')
            const res: Awaited<Array<Enum>> = await prom.data.enumList
            console.log(res)

            typeList.value = res

            for( const i in typeList.value){
              typeOptions.value.push({
                label: typeList.value[i].description,
                value: typeList.value[i].enumId
              })
            }
          },
          fetchPurposeOptions: async function(){
            const prom = await ApiService.get(env.VITE_APP_MOQUI_API_ENDPOINT+'/enumeration?enumTypeId=DevicePurpose')
            const res: Awaited<Array<Enum>> = await prom.data.enumList
            console.log(res)

            purposeList.value = res

            for( const i in purposeList.value){
              purposeOptions.value.push({
                label: purposeList.value[i].description,
                value: purposeList.value[i].enumId
              })
            }
          },
          fetchStatusFlowOptions: async function(){
            const prom = await ApiService.get(env.VITE_APP_MOQUI_API_ENDPOINT+'/statusFlowList')
            const res: Awaited<Array<StatusFlow>> = await prom.data.statusFlowList
            console.log(res)

            statusFlowList.value = res

            for( const i in statusFlowList.value){
              statusFlowOptions.value.push({
                label: statusFlowList.value[i].description,
                value: statusFlowList.value[i].statusFlowId
              })
            }
          },
          fetchStatusOptions: async function(statusFlowId){
            console.log('status flow id ', statusFlowId)
            const prom = await ApiService.get(env.VITE_APP_MOQUI_API_ENDPOINT+'/statusList?statusFlowId='+statusFlowId)
            const res: Awaited<Array<Status>> = await prom.data.statusList
            console.log(res)

            statusList.value = res

            for( const i in statusList.value){
              statusOptions.value.push({
                label: statusList.value[i].statusId,
                value: statusList.value[i].statusId
              })
            }
          },
          fetchStatusOptionsOnUpdate: async function(e){
            console.log('status flow id ', e.value)
            const prom = await ApiService.get(env.VITE_APP_MOQUI_API_ENDPOINT+'/statusList?statusFlowId='+e.value)
            const res: Awaited<Array<Status>> = await prom.data.statusList
            console.log(res)

            statusList.value = res
            deviceDetail.value.statusId = ''
            statusOptions.value = []
            for( const i in statusList.value){
              statusOptions.value.push({
                label: statusList.value[i].statusId,
                value: statusList.value[i].statusId
              })
            }
          },
        }
    },
    methods:{
        
    },
    mounted() {
      this.fetchData()
      this.fetchTypeOptions()
      this.fetchPurposeOptions()
      this.fetchStatusFlowOptions()
      
    },
    
  }
</script>
  
  