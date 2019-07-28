<template>
  <div class="off-canvas-sidebar">
    <!-- Navbar -->
    <nav class="navbar navbar-expand-lg navbar-transparent navbar-absolute fixed-top text-white">
      <div class="container">
        <div class="navbar-wrapper">
          <a class="navbar-brand" href="#pablo">Login</a>
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
          <ul class="navbar-nav">
            <li class="nav-item">
              <a href="register.html" class="nav-link">
                <i class="material-icons">person_add</i> Register
              </a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
    <!-- End Navbar -->
    <div class="wrapper wrapper-full-page">
      <div
        class="page-header login-page header-filter"
        filter-color="black"
        style="background-image: url('../../assets/img/intro.jpg'); background-size: cover; background-position: top center;"
      >
        <!--   you can change the color of the filter page using: data-color="blue | purple | green | orange | red | rose " -->
        <div class="container">
          <div class="row">
            <div class="col-lg-4 col-md-6 col-sm-8 ml-auto mr-auto">
              <form class="form" method action="#">
                <div class="card card-login">
                  <div class="card-header card-header-rose text-center">
                    <h4 class="card-title">Login</h4>
                    <div class="social-line">
                      <a href="#pablo" class="btn btn-just-icon btn-link btn-white">
                        <i class="fa fa-facebook-square"></i>
                      </a>
                      <a href="#pablo" class="btn btn-just-icon btn-link btn-white">
                        <i class="fa fa-twitter"></i>
                      </a>
                      <a href="#pablo" class="btn btn-just-icon btn-link btn-white">
                        <i class="fa fa-google-plus"></i>
                      </a>
                    </div>
                  </div>
                  <div class="card-body">
                    <span class="bmd-form-group">
                      <div class="input-group">
                        <div class="input-group-prepend">
                          <span class="input-group-text">
                            <i class="material-icons">email</i>
                          </span>
                        </div>
                        <input
                          type="text"
                          name="email"
                          v-validate="'required|email'"
                          v-model="form.email"
                          class="form-control"
                          v-bind:placeholder="$t('user')"
                        />
                        <div
                          v-show="errors.has('email')"
                          class="alert alert-danger alert-dismissible fade show"
                          role="alert"
                        >{{ errors.first('email') }}</div>
                      </div>
                    </span>
                    <span class="bmd-form-group">
                      <div class="input-group">
                        <div class="input-group-prepend">
                          <span class="input-group-text">
                            <i class="material-icons">lock_outline</i>
                          </span>
                        </div>

                        <input
                          type="password"
                          name="password"
                          v-validate="'required'"
                          v-model="form.password"
                          class="form-control"
                          v-bind:placeholder="$t('password')"
                        />
                        <div
                          v-show="errors.has('password')"
                          class="alert alert-danger alert-dismissible fade show"
                          role="alert"
                        >{{ errors.first('password') }}</div>
                      </div>
                    </span>
                  </div>

                  <div class="card-footer justify-content-center">
                    <button
                      type="button"
                      class="btn btn-rose btn-link btn-lg"
                      v-on:click="enviar()"
                    >Login</button>
                    <br>
                    <img
                      v-if="data===undefined"
                      src="../assets/loading2.gif"
                      class="rounded"
                      width="30"
                      height="30"
                      alt
                    />
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
        <footer class="footer">
          <div class="container">
            <!-- <nav class="float-left">
              <ul>
                <li>
                  <a href="https://www.creative-tim.com/">Creative Tim</a>
                </li>
                <li>
                  <a href="https://creative-tim.com/presentation">About Us</a>
                </li>
                <li>
                  <a href="http://blog.creative-tim.com/">Blog</a>
                </li>
                <li>
                  <a href="https://www.creative-tim.com/license">Licenses</a>
                </li>
              </ul>
            </nav>-->
          </div>
        </footer>
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
export default class Login extends Vue {
  public form = {
    email: "",
    password: ""
  };

  public data: any = [];

  public isError = false;
  public messages = "";

  private headers: any = {
    "Content-Type": "application/json"
  };

  public enviar() {
    // console.log('form: ');
    // console.log(this.form);
    this.data = undefined;
    this.isError = false;
    this.messages = "";
    // console.log("##hola.....");
    const endpoint: any = Global.const.LOGIN;
    // const endpoint:any = "https://disobey-api.herokuapp.com/api/v1/login"
    axios
      .post(endpoint, this.form, {
        headers: this.headers
      })
      .then(response => {
        if (response.data.success) {
          store.commit("loginUser");
          this.data = response.data.data;
          localStorage.setItem("token", response.data.data.token);
          // console.log("hola....." + response.data.data.token);
          // location.replace("/login");
          location.reload();
        } else {
          // this.errors.push(response.data.message);
          this.isError = true;
          this.messages = response.data.message;
          setTimeout(() => {
            this.isError = false;
            this.data = [];
          }, 1000);
        }
      });
  }
}
</script>
