<template>
  <v-layout row wrap class="header">
    <v-flex xs3>
      <img :src="require('../assets/icons/GT.png')" style="max-width: 50px" />
    </v-flex>
    <v-flex style="color: white" class="ml-3"> Dale </v-flex>
  </v-layout>
</template>

<script>
export default {
  name: "HeaderTest",
  data() {
    return {};
  },
};
</script>
<style scoped>
.header {
  background-color: #000d41;
}
</style>
