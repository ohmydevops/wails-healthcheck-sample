<template>
  <div>
    <h3 class="mt-3 mb-3 text-white text-bold">$ ping 8.8.8.8</h3>
    <ol class="text-white" style="width: 100%; height: 500px; overflow: auto" v-autoscroll="{ smooth: true, always: true }">
      <li v-for="pingLog in pingLogs" :key="pingLog.key" >
        {{ pingLog }}
      </li>
    </ol>
  </div>
</template>

<script>
import * as Wails from "@wailsapp/runtime";

export default {
  data() {
    return {
      pingLogs: [],
    };
  },
  mounted() {
    Wails.Events.On("initialised", (m) => {
      this.pingLogs.push(m);
    });
  },
};
</script>
