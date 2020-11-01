<template>
  <b-container class="bv-example-row">
    <b-row class="row justify-content-center">
      <b-col cols="12" md="10">
        <Table />
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
import Table from "./components/Table.vue";

export default {
  components: {
    Table,
  },
};
</script>

<style>
@import url("https://fonts.cdnfonts.com/css/calling-code");
body {
  font-family: "Calling Code", sans-serif;
  background: rgb(1, 9, 11);
  background: linear-gradient(
    90deg,
    rgba(1, 9, 11, 1) 0%,
    rgba(2, 31, 38, 1) 64%,
    rgba(0, 67, 83, 1) 100%
  );
}
td {
  color: black;
}
</style>
