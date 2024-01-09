<template>
  <div class="q-gutter-md row items-start">
    <div class="q-pa-md" style="max-width: 40%;">
      <q-list style="height: 160px;">
        <q-item-label header>Source</q-item-label>
        <q-item>
          <q-item-section avatar>
            <q-icon color="primary" name="dvr" />
          </q-item-section>
          <q-item-section>
            <!-- <q-item-label>OVERLINE</q-item-label> -->
            <!-- <q-item-label>Single line item</q-item-label>
          <q-item-label caption>Secondary line text. Lorem ipsum dolor sit amet, consectetur adipiscit
            elit.</q-item-label> -->
            <q-select dense v-model="cr.iface" :options="options" label="iface" />
          </q-item-section>
          <q-item-section>
            <q-input dense v-model="cr.port" placeholder="port" />
          </q-item-section>
          <!-- <q-item-section side top>
          <q-item-label caption>5 min ago</q-item-label>
        </q-item-section> -->
        </q-item>


      </q-list>
    </div>
    <div class="q-pa-md" style="max-width: 30%">
      <q-list style="height: 160px;">
        <q-item-label header>Destination</q-item-label>

        <q-item v-for="item in cr.destinationsList">
          <q-item-section avatar>
            <q-icon color="primary" name="dvr" />
          </q-item-section>
          <q-item-section>
            <q-input dense v-model="item.ip" placeholder="ip" />
          </q-item-section>
          <q-item-section>
            <q-input dense v-model="item.port" placeholder="port" />
          </q-item-section>
        </q-item>

        <q-item clickable v-ripple @click="addDestination()">
          <q-item-section avatar>
            <q-icon color="primary" name="add_circle" />
          </q-item-section>
          <!-- <q-item-section avatar>
          <q-icon color="primary" name="remove_circle" @click="removeDestination()" />
        </q-item-section> -->
          <q-item-section></q-item-section>
        </q-item>

        <!-- <q-separator />
        <q-item>
          <q-item-section></q-item-section>
          <q-item-section side>
            <q-btn icon="done" dense push color="white" text-color="primary" round @click="start()" />
          </q-item-section>
          <q-item-section side>
            <q-btn icon="clear" dense push color="white" text-color="primary" round @click="start()" />
          </q-item-section>
        </q-item> -->
      </q-list>
    </div>
    <div class="q-pa-md" style="max-width: 40%">
      <q-list>
        <!-- <q-item-label header>Action</q-item-label> -->

        <!-- <q-separator /> -->
        <q-item style="height: 160px;">
          <q-item-section></q-item-section>
          <q-item-section side>
            <q-btn icon="done" dense push color="white" text-color="primary" round @click="start()" />
          </q-item-section>
          <q-item-section side>
            <q-btn icon="clear" dense push color="white" text-color="primary" round @click="cancel()" />
          </q-item-section>
        </q-item>
      </q-list>
    </div>
  </div>
</template>

<script setup lang="ts">
import * as grpcWeb from 'grpc-web';
import * as empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as struct_pb from 'google-protobuf/google/protobuf/struct_pb';
import * as wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';
import { ref, onMounted, onUnmounted } from 'vue';
import { apiService } from '../api.service';
import { CopyRequest } from '../../../api/js/capture_pb';
import { useQuasar } from 'quasar';

const $q = useQuasar();
const options = ref<string[]>([]);
const cr = ref(new CopyRequest().toObject())

function start() {
  if (cr.value.iface == '') {
    $q.notify({ message: 'select iface', position: 'top-right', color: 'primary', icon: 'warning' });
  } else {
    $q.notify({ message: 'start mirror', position: 'top-right', color: 'primary', icon: 'info' });
    let copyRequest = new CopyRequest().setIface(cr.value.iface).setPort(cr.value.port);
    for (var item of cr.value.destinationsList) {
      copyRequest.addDestinations(new CopyRequest.Destination().setIp(item.ip).setPort(item.port));
    }
    apiService.captureClient.copy(copyRequest, {}, (err: grpcWeb.RpcError, respone: wrappers_pb.StringValue) => {
      if (err != null) {
        console.log(err);
        $q.notify({ message: err.message, position: 'top-right', color: 'red', icon: 'error' });
      } else {
        console.log(respone.getValue());
      }
    })
  };
}

function cancel() {
  cr.value = new CopyRequest().toObject();
  cr.value.destinationsList.push(new CopyRequest.Destination().toObject());
}

function addDestination() {
  cr.value.destinationsList.push(new CopyRequest.Destination().toObject());
}

onMounted(() => {
  cr.value.destinationsList.push(new CopyRequest.Destination().toObject());
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

onUnmounted(() => {
  console.log('begin cancel...');
})
</script>