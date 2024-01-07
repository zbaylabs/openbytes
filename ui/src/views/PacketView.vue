<template>
  <q-item>
    <q-item-section top class="col-2 gt-sm">
      <q-select dense v-model="model.iface" :options="options" label="iface" />
    </q-item-section>
    <q-item-section top>
      <q-input dense v-model="model.keyword" label="keyword" placeholder="TCP 80" />
    </q-item-section>
    <q-item-section side>
      <q-btn icon="search" dense push color="white" text-color="primary" round @click="start()" />
    </q-item-section>
  </q-item>

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

<script setup lang="ts">
import * as grpcWeb from 'grpc-web';
import * as empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as struct_pb from 'google-protobuf/google/protobuf/struct_pb';
import { ref, onMounted, onBeforeUnmount } from 'vue';
import { apiService } from '../api.service';
import { Packet, Capture } from '../../../api/js/capture_pb';
import { useQuasar } from 'quasar'

const $q = useQuasar();
const packets = ref<Packet.AsObject[]>([]);
let stream: grpcWeb.ClientReadableStream<Packet>;

const model = ref({ iface: '', keyword: '' })
const options = ref<string[]>([])

function start() {
  if (model.value.iface == '') {
    $q.notify({ message: 'select iface please', position: 'top-right', color: 'primary', icon: 'warning' });
  } else {
    $q.notify({ message: 'start capture', position: 'top-right', color: 'primary', icon: 'info' });
    packets.value = [];
    stream = apiService.captureClient.list(new Capture().setIface(model.value.iface));
    stream.on('data', response => {
      // console.log(response.toObject());
      packets.value.splice(0, 0, response.toObject());
      // items.value.push({});
    });
    stream.on('error', err => {
      console.log(err);
    });
  }
};

onMounted(() => {
  apiService.captureClient.device(new empty_pb.Empty(), {}, (error: grpcWeb.RpcError, respone: struct_pb.ListValue) => {
    if (error != null) {
      console.log(error);
    } else {
      respone.getValuesList().forEach((item: any) => {
        options.value.push(item.toObject().stringValue);
      });
    }
  })
})

onBeforeUnmount(() => {
  console.log('begin cancel...');
  if (stream != null) {
    stream.cancel();
  }
})
</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
