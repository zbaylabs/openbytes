<template>
    <!-- <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js + TypeScript App"/>
  </div> -->
    <!-- <Chart type="line" :data="chartData" :options="chartOptions" /> -->
    <Bar v-if="chartData != null" :options="chartOptions" :data="chartData" />
</template>

<script setup lang="ts">

import {
    Chart as ChartJS,
    Title,
    Tooltip,
    Legend,
    BarElement,
    CategoryScale,
    LinearScale
} from 'chart.js'
import { Bar } from 'vue-chartjs'
//import * as chartConfig from './chartConfig.js'

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend)

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

    // stream = apiService.captureClient.traffic(new Capture());
    // stream.on('data', response => {
    //     // i++;
    //     ///console.log(i, 'here-1', response.toObject());
    //     labels.push(response.toObject().label);
    //     // console.log(i, 'here0', labels);

    //     tcpCounts = tcpCounts + response.toObject().tcpCount + ','
    //     udpCounts = udpCounts + response.toObject().udpCount + ','
    //     otherCounts = otherCounts + response.toObject().otherCount + ','
    //     let abc = {
    //         labels: labels,
    //         datasets: [
    //             {
    //                 label: 'TCP',
    //                 data: tcpCounts.split(",").map(Number),
    //                 fill: true,
    //                 tension: 0.4,
    //                 //borderColor: documentStyle.getPropertyValue('--orange-500')
    //                 borderColor: 'orange',
    //                 backgroundColor: 'rgba(255,167,38,0.2)',
    //             },
    //             {
    //                 label: 'UDP',
    //                 data: udpCounts.split(",").map(Number),
    //                 fill: false,
    //                 tension: 0.4,
    //                 borderDash: [5, 5],
    //                 borderColor: 'rgb(75, 192, 192)'//documentStyle.getPropertyValue('--teal-500')
    //             }, {
    //                 label: 'Others',
    //                 data: otherCounts.split(",").map(Number),
    //                 fill: false,
    //                 tension: 0.4,
    //                 borderColor: 'blue'//documentStyle.getPropertyValue('--blue-500')
    //             }
    //         ]
    //     };
    //     // console.log(i, 'here2', abc);
    //     chartData.value = abc;
    // });
    // stream.on('error', err => {
    //     console.log(err);
    // });

    chartData.value = {
        labels: ['a.com/1', 'b.com/2', 'c.com/3', 'd.com/4', 'e.com/5'],
        datasets: [
            {
                label: 'Top 5 latency',
                data: [540, 325, 702, 620, 200],
                backgroundColor: ['rgba(255, 159, 64, 0.2)', 'rgba(75, 192, 192, 0.2)', 'rgba(54, 162, 235, 0.2)', 'rgba(153, 102, 255, 0.2)', 'yellow'],
                borderColor: ['rgb(255, 159, 64)', 'rgb(75, 192, 192)', 'rgb(54, 162, 235)', 'rgb(153, 102, 255)'],
                borderWidth: 1
            }
        ]
    };

    chartOptions.value = {
        animation: false,
        responsive: true
    };
});
</script>
