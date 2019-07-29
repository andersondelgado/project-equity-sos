<template>
  <div id="app">
    <!--
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet" />-->
    <div id="leaf"></div>
    <!-- <div class="jumbotron text-center" style="margin-bottom:0">
      <img alt="Vue logo" src="./assets/logo.png"  style="width: 80px" height="100" class="rounded">
      <h3>Equity SOS</h3>
    </div>
    -->
    <!-- <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div>-->
    <!-- <Menu></Menu> -->
    <!--
    <div class="container-fluid">
      
    </div>
    -->
    <router-view />
  </div>
</template>
<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
//import Menu from '@/views/Menu.vue';

@Component({
  components: {
    //Menu,
  }
})
export default class App extends Vue {}
</script>

<script>
export default {
  mounted() {
    // var arr_css = [
    //   "/bower_components/bootstrap/dist/css/bootstrap.min.css",
    //   "/bower_components/bootstrap/dist/css/bootstrap-grid.min.css",
    //   "/bower_components/bootstrap/dist/css/bootstrap-reboot.min.css",
    //   "/bower_components/ionicons/docs/css/ionicons.min.css",
    //   "/bower_components/bootstrap-material-design/assets/css/docs.css"
    // ];

    // var arr_lib = [
    //   "/bower_components/jquery/dist/jquery.min.js",
    //   "/bower_components/jquery/dist/jquery.slim.min.js",
    //   "/bower_components/popper.js/dist/umd/popper.min.js",
    //   "/bower_components/bootstrap/dist/js/bootstrap.min.js",
    //   "/bower_components/bootstrap-material-design/assets/js/docs.min.js"
    // ];

    // var arr_css = [
    //   "/assets/css/custom/custom.css",
    //   "/assets/css/themes/vertical-modern-menu-template/materialize.css",
    //   "/assets/css/themes/vertical-modern-menu-template/style.css",

    //   "/assets/css/vendors/vendors.min.css",
    //   "/assets/css/vendors/animate-css/animate.css",
    //   "/assets/css/pages/login.css",

    //   "/assets/vendors/chartist-js/chartist.min.css",
    //   "/assets/vendors/chartist-js/chartist-plugin-tooltip.css",
    //   "/assets/css/pages/dashboard-modern.css"

    // ];

    // var arr_lib = [
    //   "/assets/js/vendors.min.js",
    //   "/assets/js/scripts/dashboard-modern.js",
    //   "/assets/js/scripts/intro.js",
    //   "/assets/js/plugins.js",
    //   "/assets/js/scripts/customizer.js"
    // ];

    var arr_css = [
      "/bower_components/google-material-icons/google-material-icons.min.css",
      "/bower_components/ionicons/docs/css/ionicons.min.css",
      //"/bower_components/bootstrap-material-design/assets/css/docs.css",
      "/assets/css/material-dashboard.min.css",
      "/assets/demo/demo.css",
      "/leaflet/leaflet.css"
     //, "/leaflet-search/dist/leaflet-search.min.css",
     // "/leaflet-routing-machine/dist/leaflet-routing-machine.css"
    ];

    var arr_lib = [
      "/assets/js/core/jquery.min.js",
      "/assets/js/core/popper.min.js",
      "/assets/js/core/bootstrap-material-design.min.js",
      "/assets/js/plugins/perfect-scrollbar.jquery.min.js",
      "/assets/js/plugins/moment.min.js",
      "/assets/demo/demo.js",
      "/assets/js/material-dashboard.min.js?v=2.1.0",
      // Plugins
      "/assets/js/plugins/jquery.dataTables.min.js",
      "/assets/js/plugins/jquery.validate.min.js",
      "/assets/js/plugins/bootstrap-datetimepicker.min.js",
      "/assets/js/plugins/bootstrap-selectpicker.js",
      "/assets/js/plugins/fullcalendar.min.js",
      "/assets/js/plugins/chartist.min.js",
      "/assets/js/plugins/jasny-bootstrap.min.js",
      //Custom Exports
      "/assets/js/custom.js"
    ];

    var leaflet_lib = [
      "/leaflet/leaflet.js"
      // ,"leaflet-search/dist/leaflet-search.min.js",
      // "/leaflet-routing-machine/dist/leaflet-routing-machine.min.js"
    ];

    for (var i = 0; i < arr_css.length; i++) {
      var loaderTag = document.createElement("link");
      loaderTag.rel = "stylesheet";
      loaderTag.href = arr_css[i];

      var s0 = document.getElementsByTagName("head")[0];
      s0.parentNode.insertBefore(loaderTag, s0);
    }

    for (let i = 0; i < leaflet_lib.length; i++) {
      var loaderTag1 = document.createElement("script");
      loaderTag1.async = true;
      loaderTag1.src = leaflet_lib[i];

      // var s01 = document.getElementsByTagName("head")[0];
      var s01 = document.getElementById("leaf");
      s01.parentNode.insertBefore(loaderTag1, s01);
      // document.write(
      //   "<scr" +
      //     'ipt type="text/javascript" async src="' +
      //     leaflet_lib[i] +
      //     '" ></scr' +
      //     "ipt>"
      // );
    }

    

    for (var i = 0; i < arr_lib.length; i++) {
      document.write(
        "<scr" +
          'ipt type="text/javascript" src="' +
          arr_lib[i] +
          '" ></scr' +
          "ipt>"
      );

      // var loaderTag = document.createElement("script");
      // loaderTag.async = true;
      // loaderTag.src = arr_lib[i];

      // var s0 = document.getElementsByTagName("head")[0];
      // s0.parentNode.insertBefore(loaderTag, s0);
    }
  }
};
</script>


<style lang="scss">
// #app {
//   font-family: "Avenir", Helvetica, Arial, sans-serif;
//   -webkit-font-smoothing: antialiased;
//   -moz-osx-font-smoothing: grayscale;
//   text-align: center;
//   color: #2c3e50;
// }
#nav {
  padding: 30px;
  a {
    font-weight: bold;
    color: #2c3e50;
    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
