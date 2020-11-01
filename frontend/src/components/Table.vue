<template>
  <div>
    <vue-splash
      :show="this.loading"
      :style="{ zIndex: '999' }"
      title="loading ..."
      color="#00bfa5"
      :size="300"
      :fixed="true"
    />

    <div class="mt-5">
      <b-table
        striped
        hover
        dark
        bordered
        :items="websitesList"
        :fields="fields"
        :tbody-tr-class="statusColor"
      ></b-table>
    </div>

    <p class="text-white">Next check at: {{ this.nextCheckAt }} sec</p>
    <a
      v-on:click="
        openURL('https://github.com/amirbagh75/wails-healthcheck-sample')
      "
      href="#"
      >Github repo</a
    >
  </div>
</template>

<script>
export default {
  data() {
    return {
      websitesList: [],
      checkInterval: 5,
      nextCheckAt: 5, // bad pattern! shame
      interval: null,
      loading: true,
      fields: [
        {
          key: "ID",
          label: "ID",
          sortable: true,
        },
        {
          key: "Name",
          label: "Name",
          sortable: true,
        },
        {
          key: "Status",
          label: "Status",
          sortable: true,
        },
        {
          key: "LastCheckedDatetime",
          label: "Last check at",
          sortable: true,
        },
      ],
    };
  },
  mounted() {
    this.startLoading();
  },
  watch: {
    loading: function(newValue) {
      if (!newValue) this.updateTimer();
    },
  },
  methods: {
    startLoading() {
      this.initList();
      this.updateHealthCheck();
    },
    updateHealthCheck() {
      this.interval = setInterval(() => {
        this.initList();
      }, this.checkInterval * 1000);
    },
    updateTimer() {
      setInterval(() => {
        if (this.nextCheckAt > 1) {
          this.nextCheckAt--;
        } else {
          this.nextCheckAt = this.checkInterval;
        }
      }, 1000);
    },
    initList() {
      window.backend.updateListHealthCheck().then((list) => {
        this.websitesList = list;
        this.loading = false;
      });
    },
    statusColor(item, type) {
      if (!item || type !== "row") return;
      if (item.Status.startsWith("2")) return "table-success";
      if (item.Status.startsWith("5")) return "table-danger";
      if (item.Status.startsWith("4")) return "table-warning";
      return "table-info";
    },
    openURL(url) {
      window.backend.openURL(url).then((r) => {
        console.log(r);
      });
    },
  },
};
</script>
