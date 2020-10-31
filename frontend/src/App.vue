<template>
  <b-container class="bv-example-row">
    <b-row class="row justify-content-center">
      <b-col cols="12" md="8">
        <Table/>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
import Table from "./components/Table.vue";

export default {
  components: {
    Table
  }
}
</script>
