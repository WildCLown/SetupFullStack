<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <v-layout row wrap class="d-flex align-center">
        <v-flex>
          <v-layout row wrap class="align-center">
            <v-img
              alt="Vuetify Logo"
              class="shrink mr-2 d-inline-block"
              contain
              :src="require('./assets/icons/GT.png')"
              width="60"
            />
            <label class="primaryText--text text-h4 font-weight-bold">
              Gabriel Teixeira's Portifolio
            </label>
            <div class="salve">Salve</div>
          </v-layout>
        </v-flex>
        <v-flex> </v-flex>
      </v-layout>
    </v-app-bar>

    <v-main>
      <HelloWorld />
    </v-main>
  </v-app>
</template>

<script>
import HelloWorld from "./components/HelloWorld";

export default {
  name: "App",

  components: {
    HelloWorld,
  },

  data: () => ({
    //
  }),
};
</script>
<style lang="scss" scoped>
.salve {
  color: $base-color;
}
</style>