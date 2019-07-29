<template>
  <div class="menu">
    <div
      class="sidebar"
      data-color="rose"
      data-background-color="black"
      data-image="../../assets/img/intro.jpg"
    >
      <div class="logo">
        <img src="../assets/logo.png" alt="materialize logo" />
      </div>

      <div
        class="sidebar-wrapper ps-container ps-theme-default ps-active-y"
        data-ps-id="8992c927-ea82-ea0d-11b2-19825cd62e84"
      >
        <div class="user">
          <div class="photo">
            <img v-if="dataUser.avatar!==undefined" :src="dataUser.avatar" />
          </div>
          <div class="user-info">
            <a data-toggle="collapse" href="#collapseExample" class="username">
              <span>
                {{dataUser.username}}
                <!-- <b class="caret"></b> -->
              </span>
            </a>
            <div class="collapse" id="collapseExample">
              <ul class="nav">
                <!-- <li class="nav-item">
                  <a class="nav-link" href="#">
                    <span class="sidebar-mini">MP</span>
                    <span class="sidebar-normal">My Profile</span>
                  </a>
                </li>
                <li class="nav-item">
                  <a class="nav-link" href="#">
                    <span class="sidebar-mini">EP</span>
                    <span class="sidebar-normal">Edit Profile</span>
                  </a>
                </li>-->
                <!-- <li class="nav-item">
                  <a class="nav-link" href="#" @click="logout()">                     
                    <span class="sidebar-mini">L</span>
                    <span class="sidebar-normal">Logout</span>
                  </a>
                </li>-->
              </ul>
            </div>
          </div>
        </div>

        <ul class="nav">
          <img
            v-if="data===undefined"
            src="../assets/loadingv.gif"
            class="rounded"
            width="30"
            height="30"
            alt
          />
          <li class="nav-item" v-for="k in data" :key="k.url">
            <router-link :to="{ path: k.url}" class="nav-link">
              <i class="material-icons">dashboard</i>
              <!-- <p data-i18n>{{k.name}}</p> -->
              <p>{{$t(k.lang_property)}}</p>
            </router-link>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Emit, Prop } from "vue-property-decorator";
import axios from "axios";
import store from "../store";
import { messageService } from "../global";
import Global from "../global";
// import Vuex from "vue";

@Component
export default class Menu extends Vue {
  private data: any = [];
  private srtSearch: string = "";
  private headers: any = {
    "Content-Type": "application/json"
  };
  dataUser: any = {
    username: "",
    avatar: ""
  };

  public search(evt: any) {
    console.log("#search patern: " + evt);
    // let bus = new Vuex({});
    // bus.$emit("searchs", evt);
    // this.$emit('searchs',evt);
    messageService.setSearchs(evt);
  }

  public menu() {
    this.data = undefined;
    const endpoint: any = Global.const.MENU;
    axios
      .get(endpoint, {
        headers: this.headers
      })
      .then(response => {
        if (response.data.success) {
          let data: any = response.data.data;
          data.map((i: any) => {
            if (i.lang_property !== "kyc_admin") {
              return i;
            }
          });
          this.data = data;
          localStorage.setItem("permission", JSON.stringify(this.data));
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  public userInfo() {
    const endpoint: any = Global.const.USER_INFO;
    axios
      .get(endpoint, {
        headers: this.headers
      })
      .then(response => {
        if (response.data.success === true) {
          // this.dataUser = response.data.data;
          let data = response.data.data;
          this.dataUser.username = data.username;
          // this.dataUser.avatar = data.avatar;
          this.dataUser.avatar = Global.DOMAIN_FILE + data.avatar;
          localStorage.setItem("users", JSON.stringify(this.dataUser));
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  public logout() {
    localStorage.removeItem("token");
    localStorage.removeItem("permission");
    localStorage.removeItem("users");
    store.commit("logoutUser");
    // location.replace("/login");
    this.$router.replace({ path: "/login" });
    // location.reload();
  }

  public mounted() {
    this.userInfo();
    this.menu();
  }
}
</script>