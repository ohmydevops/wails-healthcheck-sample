<template>
  <b-container class="bv-example-row mt-5">
    <b-row class="row justify-content-center">
      <b-col cols="12" md="10">
        <router-link to="/table" class="mr-2">[Show status table]</router-link>
        <router-link to="/ping">[Show realtime logs]</router-link>

        <keep-alive>
          <router-view></router-view>
        </keep-alive>
      </b-col>
    </b-row>
    <br /><br />
  </b-container>
</template>

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
.router-link-active {
  color: white;
}
</style>
