<template>
  <div class="login">
    <div class="row">
      <div class="col-sm-9 col-md-7 col-lg-5 mx-auto">
        <div class="card">
          <div class="card-header">
            <h3 class="card-title text-center">Login</h3>
          </div>
          <div class="card-body">
            <form>

              <div class="form-group">

                <input
                  type="text"
                  name="username"
                  v-validate="'required'"
                  v-model="form.username"
                  class="form-control"
                  v-bind:placeholder="$t('user')"
                >

                <div
                  v-show="errors.has('username')"
                  class="alert alert-danger alert-dismissible fade show"
                  role="alert"
                >
                  {{ errors.first('username') }}
                </div>
              </div>

              <br>
              <div class="form-group">
                <input
                  type="text"
                  name="email"
                  v-validate="'required|email'"
                  v-model="form.email"
                  class="form-control"
                  v-bind:placeholder="$t('user')"
                >
                <!-- <span
                  v-show="errors.has('email')"
                  class="alert alert-danger"
                >{{ errors.first('email') }}</span>-->

                <div
                  v-show="errors.has('email')"
                  class="alert alert-danger alert-dismissible fade show"
                  role="alert"
                >
                  {{ errors.first('email') }}
                  <!-- <button
                    type="button"
                    class="close"
                    data-dismiss="alert"
                    aria-label="Close"
                  >
                    <span aria-hidden="true">&times;</span>
                  </button> -->
                </div>
              </div>

              <br>
              <div class="form-group">
                <input
                  type="password"
                  name="password"
                  v-validate="'required'"
                  v-model="form.password"
                  class="form-control"
                  v-bind:placeholder="$t('password')"
                >
                <!-- <span
                  v-show="errors.has('password')"
                  class="alert alert-danger"
                >{{ errors.first('password') }}</span>-->

                <div
                  v-show="errors.has('password')"
                  class="alert alert-danger alert-dismissible fade show"
                  role="alert"
                >
                  {{ errors.first('password') }}
                  <!-- <button
                    type="button"
                    class="close"
                    data-dismiss="alert"
                    aria-label="Close"
                  >
                    <span aria-hidden="true">&times;</span>
                  </button> -->
                </div>
              </div>

              <br>
              <button type="button" class="btn btn-outline-info" @click="enviar()">aceptar</button>
              <br>
              <!-- <span v-if="isError===true" class="alert alert-danger">{{ $t(messages) }}</span> -->
              <div
                v-if="isError===true"
                class="alert alert-danger alert-dismissible fade show"
                role="alert"
              >
                {{ $t(messages) }}
                <!-- <button
                  type="button"
                  class="close"
                  data-dismiss="alert"
                  aria-label="Close"
                >
                  <span aria-hidden="true">&times;</span>
                </button> -->
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import axios from "axios";
import store from "../store";
import Global from "../global";

@Component
export default class Register extends Vue {
  public form = {
    username: "",
    email: "",
    password: ""
  };

  public isError = false;
  public messages = "";

  private headers: any = {
    "Content-Type": "application/json"
  };

  public enviar() {
    this.isError = false;
    this.messages = "";

    const endpoint: any = Global.const.REGISTER;
    axios
      .post(endpoint, this.form, {
        headers: this.headers
      })
      .then(response => {
        if (response.data.success === true) {

          console.log(response.data);
          //store.commit("loginUser");
          //localStorage.setItem("token", response.data.data.token);
          //location.replace("/login");
        } else {
          // this.errors.push(response.data.message);
          this.isError = true;
          this.messages = response.data.message;
          setTimeout(() => {
            this.isError = false;
          }, 1000);
        }
      });
  }
}
</script>
