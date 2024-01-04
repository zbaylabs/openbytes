<script setup lang="ts">
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import ScrollPanel from 'primevue/scrollpanel';
import * as grpcWeb from 'grpc-web';

import { ref, onMounted, onBeforeUnmount } from 'vue';
import { apiService } from '../api.service';
import { Packet, Capture } from '../sdk/capture_pb';

const packets = ref<Packet.AsObject[]>([]);
let stream: grpcWeb.ClientReadableStream<Packet>;

onMounted(() => {
  packets.value = [];
  stream = apiService.captureClient.list(new Capture());
  stream.on('data', response => {
    console.log(response.toObject());
    packets.value.push(response.toObject());
  });
  stream.on('error', err => {
    console.log(err);
  });
});

onBeforeUnmount(() => {
  console.log('begin cancel...');
  stream.cancel;
})

</script>
<template>
  <DataTable :value="packets" tableStyle="min-width: 50rem;overflow: auto;">
    <Column field="protocol" header="Protocol"></Column>
    <Column field="layers" header="Layers"></Column>
    <Column field="src" header="Source">
      <template #body="slotProps">
        {{ slotProps.data.src.ip }}:{{ slotProps.data.src.port }}[{{ slotProps.data.src.mac }}]
      </template>
    </Column>
    <Column field="dst" header="Destination">
      <template #body="slotProps">
        {{ slotProps.data.dst.ip }}:{{ slotProps.data.dst.port }}[{{ slotProps.data.dst.mac }}]
      </template>
    </Column>
    <!-- <Column field="payload" header="Payload"></Column> -->
    <Column field="data" header="Data"></Column>
  </DataTable>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
