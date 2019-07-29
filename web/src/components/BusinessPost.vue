<template>
  <div id="business-post">
    <div class="card">
      <div class="card-body">
        <img
          v-if="data===undefined"
          src="../assets/loading0.gif"
          class="rounded"
          width="30"
          height="30"
          alt
        />
        <div class="row">
          <div class="col-xl-6" v-for="k in data" :key="k.id">
            <div class="card">
              <div class="card-body">
                <h5 class="card-title">
                  <div class="row">
                    <div class="col">
                      <button
                        class="btn btn-primary"
                        type="button"
                        data-toggle="collapse"
                        :data-target="'#collapseExample_'+k._id"
                        aria-expanded="false"
                        :aria-controls="'collapseExample_'+k._id"
                      >Number of Order: {{k.id}}</button>
                    </div>
                    <div class="col">
                      <button
                        class="btn btn-primary"
                        type="button"
                        data-toggle="modal"
                        v-if="k.confirmArticleByActor===true"
                        data-target="#editModal"
                        @click="edit(k._id,k._rev)"
                      >
                        <i class="icon ion-md-create"></i>
                      </button>
                    </div>
                  </div>
                </h5>

                <div class="collapse" :id="'collapseExample_'+k._id">
                  <div class="card" v-for="i in k.menuDetailPost" :key="i.id">
                    <div class="card-header" :id="'headingq_'+k.id+'_'+i.id" v-if="i.active===true">
                      <span
                        class="btn btn-link collapsed"
                        data-toggle="collapse"
                        :data-target="'#collapseqx_'+k.id+'_'+i.id"
                        aria-expanded="true"
                        :aria-controls="'collapseqx_'+k.id+'_'+i.id"
                        @click="mapInit1(k.post_data,k.id+'_'+i.id)"
                      >{{i.name}}</span>
                      <div
                        :id="'collapseqx_'+k.id+'_'+i.id"
                        class="collapse"
                        :aria-labelledby="'headingq_'+k.id+'_'+i.id"
                      >
                        <div v-if="i.id==1" :id="'map1_'+k.id+'_'+i.id" style="height: 400px"></div>
                        <div v-if="i.id==2">
                          <div
                            class="row"
                            v-for="(fi,index) in k.post_data.atachments"
                            :key="'atta'+index"
                          >
                            <div v-if="fi.type=='image'" class="row">
                              <div class="col" v-for="(im,index2) in fi.name" :key="'atta'+index2">
                                <img
                                  v-if="fi.name!=null"
                                  :src="im"
                                  class="rounded"
                                  :id="'zoomx_'+index2+'_'+k.id+'_'+i.id"
                                  @mouseover="zoomOver('zoomx_',index2+'_'+k.id+'_'+i.id)"
                                  @mouseleave="zoomOut('zoomx_',index2+'_'+k.id+'_'+i.id)"
                                  width="50"
                                  height="50"
                                  alt
                                />
                              </div>
                            </div>
                          </div>
                        </div>
                        <div v-if="i.id==3">
                          <fieldset>
                            <legend>Range Damage</legend>
                            <div class="form-group">
                              <div class="slidecontainer">
                                <input
                                  type="range"
                                  id="rangeInput"
                                  class="slider"
                                  name="rangeInput"
                                  disabled
                                  step="1"
                                  min="1"
                                  max="5"
                                  :value="k.post_data.range_damage"
                                  @change="changeDamageRange($event.target.value,k._id)"
                                />
                              </div>

                              <label :id="'rangeTextRead_'+k.id" />
                            </div>
                          </fieldset>
                          <fieldset>
                            <legend>{{k.post_data.full_text}}</legend>
                          </fieldset>
                        </div>
                        <div v-if="i.id==4">
                          <table>
                            <thead>
                              <tr>
                                <th>Name</th>
                                <th class="text text-right">Quantity ask</th>
                                <th class="text text-right">Quantity left</th>
                              </tr>
                            </thead>
                            <tbody>
                              <tr v-for="(j,index) in k.articles_data" :key="`itmlc-${index}`">
                                <td>{{j.name}}</td>
                                <td class="text text-right">{{j.quantity_ask}}</td>
                                <td class="text text-right">{{j.quantity_left}}</td>
                              </tr>
                            </tbody>
                          </table>
                        </div>
                        <div v-if="i.id==5">
                          <table>
                            <thead>
                              <tr>
                                <th>status</th>
                                <th>date</th>
                              </tr>
                            </thead>
                            <tbody>
                              <tr v-for="(l,index) in k.status_post" :key="`itml-${index}`">
                                <td>{{l.status_id}}</td>
                                <td>{{l.created_at | formatDate1 }}</td>
                              </tr>
                            </tbody>
                          </table>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
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
            <img
              v-if="dataEdit===undefined"
              src="../assets/loading0.gif"
              class="rounded"
              width="30"
              height="30"
              alt
            />

            <div
              class="row"
              v-for="(i,indexx2) in dataEdits.articles_data"
              :key="`itemsearc-${indexx2}`"
            >
              <div class="col" v-if="i.quantity_left!=0">{{i.name}}</div>
              <!-- <div class="col">{{i.name}}</div> -->
              <div class="col-sm-3">
                <div class="row" v-if="i.quantity_left!=0">
                  <!-- <div class="row"> -->
                  <div class="col">
                    <input
                      type="number"
                      class="form-control"
                      v-model="i.quantity_ask"
                      disabled
                      placeholder="quantity"
                      @oninput="subsQuantityAsk(i)"
                    />
                  </div>
                  <div class="col">
                    <input
                      type="number"
                      class="form-control"
                      v-model="i.quantity_left"
                      disabled
                      placeholder="quantity left"
                      @oninput="subsQuantityAsk(i)"
                    />
                  </div>
                  <div class="col">
                    <input
                      type="number"
                      class="form-control"
                      v-model="i.quantity_delivery"
                      placeholder="quantity send"
                      @keyup="validateFieldQuantity(i,$event)"
                      @change="validateFieldQuantity(i,$event)"
                    />
                  </div>
                </div>
              </div>
              <div class="col-sm-3" v-if="i.quantity_left!=0">
                <!-- <div class="col-sm-3"> -->
                <span class="btn btn-outline" @click="addArticleByActor(i)">
                  <i class="icon ion-md-add"></i>
                </span>
              </div>
            </div>

            <div class="row" v-for="(i,index3) in arrayArtcleByActors" :key="`itemart-${index3}`">
              <div class="col">
                <span class="btn btn-outline">
                  <i class="icon ion-md-pricetag"></i>
                  {{i.name}}
                </span>
              </div>
              <div class="col-sm-3">
                <input type="number" class="form-control" readonly :value="i.quantity_delivery" />
              </div>
              <div class="col-sm-3">
                <span @click="delItemArticleByActor(index3,i)">
                  <i class="icon ion-md-trash"></i>
                </span>
              </div>
            </div>
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
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import store from "../store";
import axios from "axios";
import Global from "../global";

declare var L: any;
declare var $: any;
@Component
export default class BusinessPost extends Vue {
  public fread: any = "";
  public fcreate: any = "";
  public fedit: any = "";
  public fupdate: any = "";
  public fdelete: any = "";

  menuAddPost: any = [
    {
      id: 1,
      name: "map",
      active: true
    },
    {
      id: 2,
      name: "attachment",
      active: true
    },
    {
      id: 3,
      name: "post_detail",
      active: true
    },
    {
      id: 4,
      name: "article",
      active: true
    },
    {
      id: 5,
      name: "status",
      active: false
    }
  ];

  rangeValues: any = {
    "1": "bajo",
    "2": "bajo-intermedio",
    "3": "intermedio",
    "4": "alto",
    "5": " extremo"
  };

  objDetail: any = {
    range_damage: 1,
    quantity_people: {
      child: 0,
      adult: 0,
      mayor: 0
    },
    fulltext: ""
  };

  dataEdit: any = [];

  form: any;

  arrayCate: any = [];

  arrayCatePostID: any = [];

  arrayArtcle: any = [];

  typeFile: any = ["image", "video"];

  attachment: any = [
    {
      type: "image",
      name: []
    },
    {
      type: "video",
      name: []
    }
  ];

  img: any = [];
  vid: any = [];
  limitImg = 5;
  limitVid = 1;
  limitSizeVid = 3501330;
  latLong: any = {
    latitude: "",
    longitude: ""
  };

  quantity_ask = 1;
  public skip = 0;
  public limit = 50;
  public limitArt = 5;

  private headers: any = {
    "Content-Type": "application/json"
  };

  browserSupportFlag = false;
  options: any = {
    enableHighAccuracy: true,
    timeout: 1000,
    maximumAge: 2000
  };
  watchId: any = "";
  dataCategories: any = [];
  dataArticles: any = [];
  columns: any = [];
  data: any = [];
  ID: any;
  NameDelete: any;
  REV: any;
  arrayArtcleByActors: any = [];
  dataEdits: any = {
    articles_data: []
  };
  formEdit: any;

  mapInit1(pos: any, id: any) {
    try {
      let container = L.DomUtil.get("map1_" + id);
      if (container != null) {
        container._leaflet_id = null;
      }
      let lat = pos.lat_long.latitude;
      let lng = pos.lat_long.longitude;
      // let radius = pos.accuracy;
      this.latLong = {
        latitude: lat,
        longitude: lng
      };
      console.log("lat: " + lat + " long: " + lng);

      // let map1 = L.map("map").setView([lat, lng], 13);

      let map1 = L.map("map1_" + id, {
        center: [lat, lng],
        zoom: 15
      });

      L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
        attribution:
          '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      }).addTo(map1);

      L.marker([lat, lng])
        .addTo(map1)
        .bindPopup("Here...")
        .openPopup();

      // L.circle([lat, lng], radius).addTo(map1);
      L.circle([lat, lng], 400, {
        color: "yellow",
        opacity: 1,
        fillColor: "blue",
        fillOpacity: 0.4
      }).addTo(map1);
    } catch (e) {
      console.log("error__map1: " + e.message);
    }
  }

  getAll() {
    const endpoint: any =
      Global.const.POST_BUSINESS_PAGINATE + this.skip + "/" + this.limit;
    this.columns = undefined;
    this.data = undefined;
    axios
      .get(endpoint, {
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
                    key !== "_id" &&
                    key !== "_rev"
                  ) {
                    keys.push(key);
                  }
                }
              });
            }

            let menuAddPost: any = [];
            this.data.map((x: any, k: any) => {
              x.articles_data.map((k1: any) => {
                // if (k1.articles_actors !== null) {
                //   x.confirmArticleByActor = true;
                // }

                x.confirmArticleByActor = true;
                if (k1.quantity_left === 0) {
                  x.confirmArticleByActor = false;
                }

                return k1;
              });

              menuAddPost = this.menuAddPost;
              x.menuDetailPost = menuAddPost.map((z: any) => {
                if (z.name === "status") {
                  z.active = true;
                }

                return z;
              });

              console.log(x.menuDetailPost);

              let y: any = x.post_data.atachments.map((i: any) => {
                let name = i.name.map((j: any) => {
                  // console.log(x);
                  return Global.DOMAIN_FILE + j;
                });
                i.name = name;
                return i;
              });

              x.post_data.atachments = y;

              console.log(x);
              // x.post_data.atachments.map((i: any) => {
              //   return i.name.map((j: any) => {
              //     // console.log(x);
              //     return Global.DOMAIN_FILE + j;
              //   });
              // });
            });
            this.columns = keys;
          }
        } else {
          // this.errors.push(response.data.message);
          this.data = [];
        }
      });
  }

  validateFieldQuantity(payload: any, event: any) {
    let val = event.target.value;
    if (
      payload.quantity_delivery < 1 ||
      payload.quantity_left < payload.quantity_delivery
    ) {
      event.target.value = val.substring(0, val.length - 1);
    }
  }

  addArticleByActor(payload: any) {
    let g = new Global();
    try {
      let val: any;
      val = payload.article_id;

      if (
        payload.quantity_delivery == null ||
        payload.quantity_delivery === undefined ||
        payload.quantity_delivery < 1
      ) {
        console.log("quantity_delivery should be major that  zero");
        return;
      }

      // if (payload.quantity_ask < payload.quantity_delivery) {
      if (payload.quantity_left < payload.quantity_delivery) {
        console.log("quantity_ask should be major that quantity_delivery");
        return;
      }

      let key: any;
      key = g.objectFindByKey(this.arrayArtcleByActors, "article_id", val);
      console.log("ke " + key);
      let keys: any;
      keys = JSON.parse(key);
      if (keys === null) {
        this.arrayArtcleByActors.push({
          article_id: val,
          name: payload.name,
          quantity_ask: parseInt(payload.quantity_ask),
          quantity_delivery: parseInt(payload.quantity_delivery)
        });

        this.dataEdits.articles_data.map((i: any, k: any) => {
          if (i.article_id == val) {
            this.dataEdits.articles_data.splice(k, 1);
          }
        });
      }
    } catch (e) {
      console.log("lol " + e.message);
    }
  }

  delItemArticleByActor(i: any, payload: any) {
    this.dataEdits.articles_data.push({
      article_id: payload.article_id,
      name: payload.name,
      quantity_ask: parseInt(payload.quantity_ask)
    });
    this.arrayArtcleByActors.splice(i, 1);
  }

  subsQuantityAsk(i: any) {
    // let qask = parseInt(i.quantity_ask) - parseInt(i.quantity_left);
    // i.quantity_ask = qask;
    console.log(i);
    return i;
  }

  changeDamageRange(val: any, id: any) {
    let d: any = document.getElementById("rangeTextRead_" + id);
    d.innerText = this.rangeValues[val];
  }

  zoomOver(el: any, index: any) {
    let c: any = document.getElementById(el + index);
    let arr: any = c.className.split(" ");
    let name: any = "transition";
    if (arr.indexOf(name) == -1) {
      c.className += " " + name;
      c.style.width = "450px";
      c.style.height = "450px";
    }
  }

  zoomOut(el: any, index: any) {
    let c: any = document.getElementById(el + index);
    c.classList.remove("transition");
    c.style.width = "50px";
    c.style.height = "50px";
  }

  public edit(id: any, _rev: any) {
    this.dataEdit = [];
    const endpoint: any = Global.const.POST_BUSINESS_EDIT + id;
    axios
      .get(endpoint, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success === true) {
          console.log(response.data.data);
          const data: any = response.data.data;
          this.dataEdit = data;
          this.dataEdits.articles_data = data.articles_data;
          // this.formEdit.name = data.name;
          // this.formEdit.description = data.description;
          this.ID = id;
          this.REV = _rev;
          this.NameDelete = data.name;
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  public update() {
    console.log(this.arrayArtcleByActors);
    let arrayArtcleByActors: any = [];
    this.arrayArtcleByActors.map((i: any) => {
      arrayArtcleByActors.push({
        article_id: i.article_id,
        articles_actors: [
          {
            quantity_delivery: parseInt(i.quantity_delivery)
          }
        ]
      });
    });

    if (arrayArtcleByActors.length == 0) {
      return;
    }
    this.dataEdit = undefined;
    const endpoint: any =
      Global.const.POST_BUSINESS_UPDATE + this.ID + "/" + this.REV;
    // return;
    this.formEdit = {
      articles_data: arrayArtcleByActors
    };
    axios
      .put(endpoint, this.formEdit, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success) {
          this.dataEdit = [];
          $("#editModal").modal("hide");
          this.dataEdits.articles_data = [];
          this.arrayArtcleByActors = [];
          this.getAll();
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  mounted() {
    this.initPermission();
    this.getAll();
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
<style scope>
.slidecontainer {
  width: 100%;
}

.slider {
  -webkit-appearance: none;
  width: 100%;
  height: 25px;
  background: #d3d3d3;
  outline: none;
  opacity: 0.7;
  -webkit-transition: 0.2s;
  transition: opacity 0.2s;
}

.slider:hover {
  opacity: 1;
}

.slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 25px;
  height: 25px;
  background: rgb(175, 76, 114);
  cursor: pointer;
}

.slider::-moz-range-thumb {
  width: 25px;
  height: 25px;
  background: rgb(175, 76, 114);
  cursor: pointer;
}
</style>

