<script setup lang="ts">
import * as grpcWeb from 'grpc-web';

import { ref, onMounted, onBeforeUnmount } from 'vue';
import { apiService } from '../api.service';
import { Packet, Capture } from '../../../api/js/capture_pb';

const packets = ref<Packet.AsObject[]>([]);
let stream: grpcWeb.ClientReadableStream<Packet>;

onMounted(() => {
  packets.value = [];
  stream = apiService.captureClient.list(new Capture());
  stream.on('data', response => {
    console.log(response.toObject());
    //packets.value.push(response.toObject());
    packets.value.splice(0, 0, response.toObject());
    //items.value.push({});
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
  <div class="q-pa-md">
    <q-scroll>
      <!-- <template v-slot:loading>
        <div class="row justify-center q-my-md">
          <q-spinner color="primary" name="dots" size="40px" />
        </div>
      </template> -->

      <div v-for="(item, index) in packets" :key="index" class="caption q-py-sm">
        <q-badge class="shadow-1">
          {{ packets.length - index }}
        </q-badge>
        {{ item.data }}
        <!-- Lorem ipsum dolor sit amet consectetur adipisicing elit. Rerum repellendus sit voluptate voluptas eveniet porro.
        Rerum blanditiis perferendis totam, ea at omnis vel numquam exercitationem aut, natus minima, porro labore. -->
      </div>
    </q-scroll>
  </div>
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
