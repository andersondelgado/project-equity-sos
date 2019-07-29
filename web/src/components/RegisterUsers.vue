<template>
  <div id="register_users">
    <div class="card">
      <div class="card-body">
        <button
          type="button"
          class="btn btn-outline-primary"
          data-toggle="modal"
          data-target="#createModal"
          v-if="fcreate!=0"
        >
          <i class="icon ion-md-document"></i>
        </button>
        <div v-if="fread!=0">
          <img
            v-if="data===undefined"
            src="../assets/loading0.gif"
            class="rounded"
            width="30"
            height="30"
            alt
          />
          <div class="table-responsive">
            <table v-if="data!==undefined">
              <thead>
                <tr v-if="columns!==undefined">
                  <th v-for="k in columns" :key="k">{{$t(k)}}</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="k in data" :key="k.id">
                  <!-- <td>{{k.avatar}}</td> -->
                  <td>{{k.username}}</td>
                  <td>{{k.email}}</td>
                  <td>
                    <button
                      class="btn btn-outline-info"
                      v-if="fedit!=0"
                      data-toggle="modal"
                      data-target="#editModal"
                      @click="edit(k.id,k._rev)"
                    >
                      <i class="icon ion-md-create"></i>
                    </button>
                    <button
                      class="btn btn-outline-info"
                      v-if="fdelete!=0"
                      data-toggle="modal"
                      data-target="#deleteModal"
                      @click="edit(k.id,k._rev)"
                    >
                      <i class="icon ion-md-trash"></i>
                    </button>

                    <router-link
                      class="btn btn-outline-info"
                      :to="{ name: 'kyc-admin', params: { id: k._id }}"
                    >kyc</router-link>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>

    <div
      class="modal fade"
      id="createModal"
      tabindex="-1"
      role="dialog"
      aria-labelledby="myModal"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="createModal"></h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form>
              <img
                v-if="dataEdit===undefined"
                src="../assets/loading0.gif"
                class="rounded"
                width="30"
                height="30"
                alt
              />
              <span class="btn btn-outline btn-file">
                <i class="icon ion-md-image"></i>
                <input type="file" class="form-control-file" @change="setAvatar($event)" />
              </span>
              <div class="photo">
                <img
                  v-if="form.avatar!==undefined || form.avatar!==''"
                  :src="form.avatar"
                  width="100"
                  height="100"
                  class
                  alt
                />
              </div>
              <div class="form-group">
                <input
                  type="text"
                  v-model="form.username"
                  name="username"
                  v-validate="'required'"
                  class="form-control"
                  placeholder="username"
                />
                <br />
                <div
                  v-if="errors.has('username')"
                  class="alert alert-danger alert-dismissible fade show"
                  role="alert"
                >{{ errors.first('username') }}</div>
              </div>

              <div class="form-group">
                <input
                  type="text"
                  v-model="form.email"
                  name="email"
                  v-validate="'required'"
                  class="form-control"
                  placeholder="email"
                />
                <div
                  v-if="errors.has('email')"
                  class="alert alert-danger alert-dismissible fade show"
                  role="alert"
                >{{ errors.first('email') }}</div>
                <br />
              </div>
              <div class="form-group">
                <input
                  type="password"
                  v-model="form.password"
                  name="password"
                  v-validate="'required'"
                  class="form-control"
                  placeholder="password"
                />
                <div
                  v-if="errors.has('password')"
                  class="alert alert-danger alert-dismissible fade show"
                  role="alert"
                >{{ errors.first('password') }}</div>
                <br />
              </div>
              <div class="form-group">
                <select name="rol" v-model="form.rol" class="form-control">
                  <option>Select rol</option>
                  <option v-for="i in data_select" :key="i._id" v-bind:value="i.rol">{{i.rol}}</option>
                </select>
              </div>
            </form>
          </div>
          <div class="modal-footer">
            <button
              type="button"
              id="save"
              class="btn btn-outline-info"
              v-if="fcreate!=0"
              @click="send()"
            >
              <i class="icon ion-md-save"></i>
            </button>
          </div>
        </div>
      </div>
    </div>
    <div
      class="modal fade"
      id="editModal"
      tabindex="-1"
      role="dialog"
      aria-labelledby="myModal"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="editModal"></h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form>
              <img
                v-if="dataEdit===undefined"
                src="../assets/loading0.gif"
                class="rounded"
                width="30"
                height="30"
                alt
              />
              <span class="btn btn-outline btn-file">
                <i class="icon ion-md-image"></i>
                <input type="file" class="form-control-file" @change="setAvatarEdit($event)" />
              </span>
              <div class="photo">
                <img
                  v-if="formEdit.avatar!==undefined || form.avatar!==''"
                  :src="formEdit.avatar"
                  width="100"
                  height="100"
                  class
                  alt
                />
              </div>
              <div class="form-group">
                <input
                  type="text"
                  v-model="formEdit.username"
                  name="username"
                  v-validate="'required'"
                  class="form-control"
                  placeholder="username"
                />
                <br />
                <div
                  v-if="errors.has('username')"
                  class="alert alert-danger alert-dismissible fade show"
                  role="alert"
                >{{ errors.first('username') }}</div>
              </div>

              <div class="form-group">
                <input
                  type="text"
                  v-model="formEdit.email"
                  name="email"
                  v-validate="'required'"
                  class="form-control"
                  placeholder="email"
                />
                <div
                  v-if="errors.has('email')"
                  class="alert alert-danger alert-dismissible fade show"
                  role="alert"
                >{{ errors.first('email') }}</div>
                <br />
              </div>
              <div class="form-group">
                <input
                  type="password"
                  v-model="formEdit.password"
                  name="password"
                  v-validate="'required'"
                  class="form-control"
                  placeholder="password"
                />
                <div
                  v-if="errors.has('password')"
                  class="alert alert-danger alert-dismissible fade show"
                  role="alert"
                >{{ errors.first('password') }}</div>
                <br />
              </div>
              <div class="form-group">
                <select name="rol" v-model="formEdit.rol" class="form-control">
                  <option>Select rol</option>
                  <option v-for="i in data_select" :key="i._id" v-bind:value="i.rol">{{i.rol}}</option>
                </select>
              </div>
            </form>
          </div>
          <div class="modal-footer">
            <button
              type="button"
              id="update"
              class="btn btn-outline-info"
              v-if="fupdate!=0"
              v-on:click="update()"
            >
              <i class="icon ion-md-sync"></i>
            </button>
          </div>
        </div>
      </div>
    </div>

    <div
      class="modal fade"
      id="deleteModal"
      tabindex="-1"
      role="dialog"
      aria-labelledby="myModal"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="deleteModal"></h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <img
            v-if="dataEdit===undefined"
            src="../assets/loading0.gif"
            class="rounded"
            width="30"
            height="30"
            alt
          />
          <div class="modal-body">¿Esta seguro de eliminar {{NameDelete}}?</div>
          <div class="modal-footer">
            <button type="button" class="btn btn-primary" @click="deletes()">
              <i class="icon ion-md-trash"></i>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import store from "../store";
import axios from "axios";
import { messageService } from "../global";
import Global from "../global";

declare var $: any;

@Component
export default class RegisterUsers extends Vue {
  public columns: any = [];
  public data: any = [];
  public data_select: any = [];
  public dataEdit: any = [];
  public srtSearch: string = "";
  forms = {
    avatar: "",
    username: "",
    email: "",
    password: "",
    rol: ""
  };

  public form = this.forms;
  public formEdit = this.forms;

  public fread: any = "";
  public fcreate: any = "";
  public fedit: any = "";
  public fupdate: any = "";
  public fdelete: any = "";

  public skip = 0;
  public limit = 50;

  private headers: any = {
    "Content-Type": "application/json"
  };

  private ID: any;
  private REV: any;
  public NameDelete: any = "";
  subscription: any;
  public getAll() {
    const endpoint: any =
      Global.const.USERS_PAGINATE + this.skip + "/" + this.limit;
    this.columns = undefined;
    this.data = undefined;
    axios
      .get(endpoint, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success === true) {
          // console.log(response.data.data);
          this.data = response.data.data;
          const keys: any = [];
          const jsonData: any = this.data;
          for (let i = 0; i < jsonData.length; i++) {
            Object.keys(jsonData[i]).forEach((key: any) => {
              if (keys.indexOf(key) === -1) {
                if (
                  key !== "id" &&
                  key !== "created_at" &&
                  key !== "updated_at" &&
                  key !== "password" &&
                  key !== "avatar" &&
                  key !== "_id" &&
                  key !== "_rev"
                ) {
                  keys.push(key);
                }
              }
            });
          }

          this.columns = keys;
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  public setAvatar(evt: any) {
    let g = new Global();
    g.getBase64ImageEncode(evt.target.files[0]).then((res: any) => {
      this.form.avatar = res;
    });
  }

  public setAvatarEdit(evt: any) {
    let g = new Global();
    g.getBase64ImageEncode(evt.target.files[0]).then((res: any) => {
      this.formEdit.avatar = res;
    });
  }

  public permission() {
    this.data_select = undefined;
    const endpoint: any = Global.const.PERMISSION_LIST;
    axios
      .get(endpoint, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success === true) {
          this.data_select = response.data.data;
        }
      });
  }

  public searchAll(payload: any) {
    const endpoint: any =
      Global.const.USERS_SEARCH_PAGINATE + this.skip + "/" + this.limit;
    this.columns = undefined;
    this.data = undefined;
    axios
      .post(endpoint, payload, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success === true) {
          // console.log(response.data.data);
          this.data =
            response.data.data !== undefined ? response.data.data : [];
          const keys: any = [];
          const jsonData: any = this.data;
          if (jsonData.length !== 0) {
            for (let i = 0; i < jsonData.length; i++) {
              Object.keys(jsonData[i]).forEach((key: any) => {
                if (keys.indexOf(key) === -1) {
                  if (
                    key !== "id" &&
                    key !== "created_at" &&
                    key !== "updated_at" &&
                    key !== "password" &&
                    key !== "avatar" &&
                    key !== "_id" &&
                    key !== "_rev"
                  ) {
                    keys.push(key);
                  }
                }
              });
            }
            this.columns = keys;
          }
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  private observable() {
    console.log("#search child init ");
    this.subscription = messageService
      .getSearchs()
      .subscribe((message: any) => {
        let rr = JSON.stringify(message);
        let obj = JSON.parse(rr);
        let r = message["text"];
        this.srtSearch = r;
        console.log("#search child: " + r);
        console.log("##obj: " + obj.text);
        let payload = {
          name: obj.text
        };

        this.searchAll(payload);
      });
  }

  public edit(id: any, _rev: any) {
    const endpoint: any = Global.const.USERS_EDIT + id;
    this.dataEdit = undefined;
    axios
      .get(endpoint, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success === true) {
          // console.log(response.data.data);
          const data: any = response.data.data;
          this.dataEdit = data;
          this.formEdit.avatar = data.avatar;
          this.formEdit.username = data.username;
          this.formEdit.email = data.email;

          // this.formEdit.type_article = data.type_article;
          this.ID = id;
          this.REV = _rev;
          this.NameDelete = data.username;
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  public send() {
    this.dataEdit = undefined;
    const endpoint: any = Global.const.USERS_SAVE;
    axios
      .post(endpoint, this.form, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success === true) {
          this.dataEdit = [];
          $("#createModal").modal("hide");
          this.initComponent();
          this.getAll();
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  public update() {
    this.dataEdit = undefined;
    const endpoint: any = Global.const.USERS_UPDATE + this.ID + "/" + this.REV;
    axios
      .put(endpoint, this.formEdit, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success === true) {
          this.dataEdit = [];
          $("#editModal").modal("hide");
          this.initComponent();
          this.getAll();
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  public deletes() {
    this.dataEdit = undefined;
    const endpoint: any = Global.const.USERS_DELETE + this.ID + "/" + this.REV;
    axios
      .get(endpoint, {
        headers: this.headers
      })
      .then((response: any) => {
        this.dataEdit = [];
        $("#deleteModal").modal("hide");
        this.NameDelete = "";
        this.initComponent();
        this.getAll();
      });
  }

  initComponent() {}

  created() {
    this.observable();
  }

  beforeDestroy() {
    this.subscription.unsubscribe();
  }

  mounted() {
    this.initPermission();
    this.getAll();
    this.permission();
  }

  initPermission() {
    let global = new Global();
    const h = window.location.pathname;
    const url = h;
    // console.log("#url: " + url);
    this.fread = global.isRead(url);
    this.fcreate = global.isCreate(url);
    this.fedit = global.isEdit(url);
    this.fupdate = global.isUpdate(url);
    this.fdelete = global.isDelete(url);
  }
}
</script>
<style lang="scss" scoped>
.btn-file {
  position: relative;
  overflow: hidden;
}
.btn-file input[type="file"] {
  position: absolute;
  top: 0;
  right: 0;
  min-width: 100%;
  min-height: 100%;
  font-size: 100px;
  text-align: right;
  filter: alpha(opacity=0);
  opacity: 0;
  outline: none;
  background: white;
  cursor: inherit;
  display: block;
}

.btn-outline {
  background-color: transparent;
  border-color: transparent;
  box-shadow: none;
}
</style>
