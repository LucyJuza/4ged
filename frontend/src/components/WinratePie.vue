<template>
  <div :key="id" class="h-100 w-100">
    <Pie
    id="piechart-winrate"
    :options="options"
    :data="data"
    />
  </div>
</template>
<script setup>
import { useAppStore } from '@/stores/app';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js'
import { ref, watch } from 'vue';
import { Pie } from 'vue-chartjs'
import { useDisplay } from 'vuetify';
const appStore = useAppStore()
const stats = appStore.getStatistics()
const id = ref(0)
const {height,width} = useDisplay()
// refresh chart on page resize because lib doesn't work properly by default
watch([height,width], () =>{
  id.value++
})
ChartJS.register(ArcElement, Tooltip, Legend)
const options = {
  responsive:true,
  maintainAspectRatio: false
}
const data = ref({
  labels: [
    'Gagnées',
    'Perdues',
  ],
  datasets: [{
    label: 'nombre de parties',
    data: [stats.personalWinrate.won, stats.personalWinrate.lost],
    backgroundColor: [
      '#B76F91',
      '#9B26DE',
    ],
    hoverOffset: 4
  }]
})
</script>
<style>
/*
Weird as fuck Chart.js issue
see: https://github.com/chartjs/Chart.js/issues/11005
for reason for this workaround
*/
canvas {
  width: 100% !important;
  height: 100% !important;
}
</style>