<template>
  <div>
    <b-table
      striped
      hover
      :items="websitesList"
      :fields="fields"
      :tbody-tr-class="statusColor"
    ></b-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      websitesList: [],
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
  methods: {
    startLoading() {
      this.updateList();
      this.updateHealthCheck();
    },
    updateHealthCheck() {
      setInterval(() => {
        this.updateList();
      }, 3000);
    },
    updateList() {
      window.backend.updateListHealthCheck().then((list) => {
        this.websitesList = list;
      });
    },
    statusColor(item, type) {
      if (!item || type !== "row") return;
      if (item.Status.startsWith("2")) return "table-success";
      if (item.Status.startsWith("5")) return "table-danger";
      if (item.Status.startsWith("4")) return "table-warning";
      return "table-info";
    },
  },
};
</script>
