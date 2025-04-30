<template>
  <div class="q-pa-md">
    <q-table
      flat bordered
      title="Parameter"
      :rows="rows"
      :columns="columns"
      :filter="filter"
      :visible-columns="visibleColumns"
      v-model:pagination="pagination"
      row-key="parameterId"
    >
      
      <template v-slot:top>
        <q-input  debounce="300" v-model="filter" placeholder="Search">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>
        <q-space />

        <q-select
          v-model="visibleColumns"
          multiple
          outlined
          dense
          options-dense
          :display-value="quasar.lang.table.columns"
          emit-value
          map-options
          :options="columns"
          option-value="name"
          options-cover
          style="min-width: 150px"
        />
      </template>
      <template v-slot:header="props">
        <q-tr :props="props">
          <q-th auto-width />
          <q-th
            v-for="col in props.cols"
            :key="col.name"
            :props="props"
          >
            {{ col.label }}
          </q-th>
        </q-tr>
      </template>
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td auto-width>
          </q-td>
          <q-td key="parameterId" :props="props">{{ props.row.parameterId }}</q-td>
          <q-td key="pseudoId" :props="props">{{ props.row.pseudoId }}</q-td>
          <q-td key="parameterCode" :props="props">{{ props.row.parameterCode }}</q-td>
          <q-td key="parameterName" :props="props">{{ props.row.parameterName }}</q-td>
          <q-td key="description" :props="props">
            {{ props.row.description.substr(0, 5) }}...
            <q-popup-edit v-model="props.row.description" title="Description"  v-slot="scope">
              <q-input type="textarea" v-model="scope.value" dense autofocus readonly />
            </q-popup-edit>
          </q-td> 
          <q-td key="minValue" :props="props">{{ props.row.minValue }}</q-td>
          <q-td key="parameterValue" :props="props">
            {{ props.row.parameterValue }}
            <q-popup-edit v-model="props.row.parameterValue" :title="'Parameter '+ props.row.parameterId + ' value:'" buttons v-slot="scope" @save="saveValue">
              <q-input type="text" v-model="scope.value" dense autofocus />
            </q-popup-edit>
          </q-td>
          
          <q-td key="maxValue" :props="props">{{ props.row.maxValue }}</q-td>
          <q-td key="scalingFactor" :props="props">{{ props.row.scalingFactor }}</q-td>

        </q-tr>
        <q-tr v-show="props.expand" :props="props">
          <q-td colspan="100%">
            <div class="text-left">Description: {{ props.row.description }}.</div>
          </q-td>
        </q-tr>
      </template>


    </q-table>
  </div>
</template>

<script lang="ts">
import { ref } from 'vue'
import ApiService from '@/services/api.service'
import { QTableProps, useQuasar } from 'quasar'


type Parameter = {
  parameterId: string
  pseudoId: string
  parameterCode: string
  description: string
  minValue: string
  parameterValue: string
  maxValue: string
  scalingFactor: number
}

const columns : QTableProps = [
  { name: 'parameterId', align: 'left', label: 'ID', field: 'parameterId', sortable: true },
  { name: 'pseudoId', align: 'center', label: 'Pseudo Id', field: 'pseudoId', sortable: true },
  { name: 'parameterCode', align: 'center', label: 'Code', field: 'parameterCode', sortable: true },
  { name: 'parameterName', align: 'left', label: 'Name', field: 'parameterName', sortable: true },
  { name: 'description', align: 'left', label: 'Description', field: 'description', sortable: true },
  { name: 'minValue', align: 'center', label: 'Min Value', field: 'minValue', sortable: true },
  { name: 'parameterValue', align: 'center', label: 'Value', field: 'parameterValue', sortable: true },
  { name: 'maxValue', align: 'center', label: 'Max Value', field: 'maxValue', sortable: true },
  { name: 'scalingFactor', align: 'center', label: 'Scaling Factor', field: 'Scaling Factor', sortable: true },
]


export default {
  setup () {
    const env = import.meta.env
    const rows = ref([])
    const filter = ref('')
    const pagination = ref({
      sortBy: 'desc',
      descending: false,
      rowsPerPage: 20,
    })

    const quasar = useQuasar()

    return {
      visibleColumns: ref([ 'parameterId', 'pseudoId','parameterCode', 'parameterName', 'minValue', 'parameterValue', 'maxValue']),
      columns,
      rows,
      filter,
      pagination,
      quasar,
      fetchRows: async function(){
        const prom = await ApiService.get(env.VITE_APP_MOQUI_API_ENDPOINT+'/parameterList?deviceId='+this.$route.params.id)

        console.log(prom)
        //const prom = await fetch(env.VITE_APP_MOQUI_API_ENDPOINT+"/deviceList")
        const res: Awaited<Array<Parameter>> = await prom.data.parameterList
        console.log(res)

        rows.value = res
      },
      saveValue: function(e){
        alert(e)
      }
    }
  },
  mounted: function(){
    this.fetchRows()
  }
}
</script>
