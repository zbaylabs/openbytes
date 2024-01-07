<template>
    <!-- <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js + TypeScript App"/>
  </div> -->
    <!-- <Chart type="line" :data="chartData" :options="chartOptions" /> -->
    <Line v-if="chartData != null" :options="chartOptions" :data="chartData" />
</template>

<script setup lang="ts">
import { defineComponent } from 'vue';
import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src

import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend, Filler } from 'chart.js'
import { Line } from 'vue-chartjs';
ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend, Filler)

import * as grpcWeb from 'grpc-web';
import { ref, onMounted } from "vue";
import { apiService } from '../api.service';
import { Point, Capture } from '../../../api/js/capture_pb';

const points = ref<Point.AsObject[]>([]);
let stream: grpcWeb.ClientReadableStream<Point>;

const chartData = ref();
const chartOptions = ref();

onMounted(() => {
    let labels: string[] = [];
    let tcpCounts = '';
    let udpCounts = '';
    let otherCounts = '';
    let n = 1;

    stream = apiService.captureClient.traffic(new Capture());
    stream.on('data', response => {
        // i++;
        ///console.log(i, 'here-1', response.toObject());
        labels.push(response.toObject().label);
        // console.log(i, 'here0', labels);

        tcpCounts = tcpCounts + response.toObject().tcpCount + ','
        udpCounts = udpCounts + response.toObject().udpCount + ','
        otherCounts = otherCounts + response.toObject().otherCount + ','
        let abc = {
            labels: labels,
            datasets: [
                {
                    label: 'TCP',
                    data: tcpCounts.split(",").map(Number),
                    fill: true,
                    tension: 0.4,
                    //borderColor: documentStyle.getPropertyValue('--orange-500')
                    borderColor: 'orange',
                    backgroundColor: 'rgba(255,167,38,0.2)',
                },
                {
                    label: 'UDP',
                    data: udpCounts.split(",").map(Number),
                    fill: false,
                    tension: 0.4,
                    borderDash: [5, 5],
                    borderColor: 'rgb(75, 192, 192)'//documentStyle.getPropertyValue('--teal-500')
                }, {
                    label: 'Others',
                    data: otherCounts.split(",").map(Number),
                    fill: false,
                    tension: 0.4,
                    borderColor: 'blue'//documentStyle.getPropertyValue('--blue-500')
                }
            ]
        };
        // console.log(i, 'here2', abc);
        chartData.value = abc;
    });
    stream.on('error', err => {
        console.log(err);
    });

    chartOptions.value = {
        animation: false,
        responsive: true
    };
});
</script>
