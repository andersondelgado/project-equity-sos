<template>
  <div id="index" class="wrapper">
    <Menu></Menu>
    <div
      class="main-panel ps-container ps-theme-default ps-active-y"
      data-ps-id="e450480f-5853-93a0-ec30-b393a23daab0"
    >
      <!-- Navbar -->
      <nav class="navbar navbar-expand-lg navbar-transparent navbar-absolute fixed-top">
        <div class="container-fluid">
          <div class="navbar-wrapper">
            <!-- <div class="navbar-minimize">
              <button id="minimizeSidebar" class="btn btn-just-icon btn-white btn-fab btn-round">
                <i class="material-icons text_align-center visible-on-sidebar-regular">more_vert</i>
                <i class="material-icons design_bullet-list-67 visible-on-sidebar-mini">view_list</i>
              </button>
            </div>-->
            <!-- <a class="navbar-brand" href="#"></a> -->
          </div>
          <button
            class="navbar-toggler"
            type="button"
            data-toggle="collapse"
            aria-controls="navigation-index"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span class="sr-only">Toggle navigation</span>
            <span class="navbar-toggler-icon icon-bar"></span>
            <span class="navbar-toggler-icon icon-bar"></span>
            <span class="navbar-toggler-icon icon-bar"></span>
          </button>
          <div class="collapse navbar-collapse justify-content-end">
            <div class="navbar-form">
              <span class="bmd-form-group">
                <div class="input-group no-border">
                  <input
                    class="form-control"
                    type="text"
                    name="Search"
                    placeholder="Buscar..."
                    v-on:keypress.enter="search($event.target.value)"
                  />
                  <button type="button" class="btn btn-white btn-round btn-just-icon">
                    <i class="material-icons">search</i>
                    <div class="ripple-container"></div>
                  </button>
                </div>
              </span>
            </div>

            <ul class="navbar-nav">
              <!-- <li class="nav-item">
                <a class="nav-link" href="#pablo">
                  <i class="material-icons">dashboard</i>
                  <p class="d-lg-none d-md-block">Stats</p>
                </a>
              </li>-->
              <!-- <li class="nav-item dropdown">
                <a
                  class="nav-link"
                  href="http://example.com/"
                  id="navbarDropdownMenuLink"
                  data-toggle="dropdown"
                  aria-haspopup="true"
                  aria-expanded="true"
                >
                  <i class="material-icons">notifications</i>
                  <span class="notification">5</span>
                  <p class="d-lg-none d-md-block">Some Actions</p>
                  <div class="ripple-container"></div>
                </a>
                <div
                  class="dropdown-menu dropdown-menu-right hide"
                  aria-labelledby="navbarDropdownMenuLink"
                >
                  <a class="dropdown-item" href="#">Mike John responded to your email</a>
                  <a class="dropdown-item" href="#">You have 5 new tasks</a>
                  <a class="dropdown-item" href="#">You're now friend with Andrew</a>
                  <a class="dropdown-item" href="#">Another Notification</a>
                  <a class="dropdown-item" href="#">Another One</a>
                </div>
              </li>-->
              <li class="nav-item dropdown">
                <a
                  class="nav-link"
                  href="#pablo"
                  id="navbarDropdownProfile"
                  data-toggle="dropdown"
                  aria-haspopup="true"
                  aria-expanded="false"
                >
                  <i class="material-icons">person</i>
                  <p class="d-lg-none d-md-block">Account</p>
                  <div class="ripple-container"></div>
                </a>
                <div
                  class="dropdown-menu dropdown-menu-right"
                  aria-labelledby="navbarDropdownProfile"
                >
                  <!-- <a class="dropdown-item" href="#">Profile</a> -->
                  <!-- <a class="dropdown-item" href="#">Settings</a>  -->
                  <!-- <div class="dropdown-divider"></div> -->
                  <a
                    class="dropdown-item"
                    href="javascript:void(0)"
                    @click="logout()"
                  >{{$t('logout')}}</a>
                </div>
              </li>
            </ul>
          </div>
        </div>
      </nav>
      <!-- End Navbar -->

      <!-- content -->
      <div class="content">
        <div class="container-fluid">
          <router-view></router-view>
        </div>
      </div>
      <!-- /content -->
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import Menu from "@/views/Menu.vue";
// import Menu from "@/components/Menu.vue";
import axios from "axios";
import store from "../store";
import Global from "../global";
import { messageService } from "../global";

@Component({
  components: {
    Menu
  }
})
export default class Index extends Vue {
  private data: any = [];

  private headers: any = {
    "Content-Type": "application/json"
  };

  dataUser: any = {
    username: "",
    avatar: ""
  };

  /**/
  public menu() {
    console.log("Menu de Componente Index");
    // console.log('form: ');
    // console.log(this.form);
    const endpoint: any = Global.const.MENU;
    axios
      .get(endpoint, {
        headers: this.headers
      })
      .then(response => {
        if (response.data.success === true) {
          this.data = response.data.data;
          localStorage.setItem("permission", JSON.stringify(this.data));
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }
  /**/
  public userInfo() {
    // console.log('form: ');
    // console.log(this.form);
    const endpoint: any = Global.const.USER_INFO;
    axios
      .get(endpoint, {
        headers: this.headers
      })
      .then(response => {
        if (response.data.success === true) {
          this.dataUser = response.data.data;
          localStorage.setItem("users", JSON.stringify(this.dataUser));
        } else {
          // this.errors.push(response.data.message);
        }
      });
  } /**/

  public search(evt: any) {
    console.log("#search patern: " + evt);
    // let bus = new Vuex({});
    // bus.$emit("searchs", evt);
    // this.$emit('searchs',evt);

    messageService.setSearchs(evt);
  }

  public logout() {
    localStorage.removeItem("token");
    localStorage.removeItem("permission");
    store.commit("logoutUser");
    // location.replace("/login");
    this.$router.replace({ path: "/login" });
    // location.reload();
  }

  public mounted() {
    this.menu();
    this.userInfo();
  }
}
</script>

