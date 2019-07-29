<template>
  <div id="add-post">
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

        <button class="btn btn-primary" type="button" @click="getAll()">
          <i class="icon ion-md-sync"></i>
        </button>
        <br />

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
                        v-if="k.confirmArticleByActor===true"
                        type="button"
                        data-toggle="modal"
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

                            <div v-if="fi.type=='video'" class="row">
                              <div
                                class="col"
                                v-for="(fi1,index1) in fi.name"
                                :key="`itemll-${index1}`"
                              >
                                <div class="embed-responsive embed-responsive-4by3">
                                  <video
                                    class="embed-responsive-item"
                                    v-if="vid!=null"
                                    :src="fi1"
                                    width="250"
                                    height="250"
                                    controls
                                  />
                                </div>
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
            <div class="card" v-for="i in menuAddPost" :key="i.id">
              <div class="card-header" :id="'heading_'+i.id" v-if="i.id!=5 && i.active===true">
                <span
                  class="btn btn-link collapsed"
                  data-toggle="collapse"
                  :data-target="'#collapsex_'+i.id"
                  aria-expanded="true"
                  :aria-controls="'collapsex_'+i.id"
                >{{i.name}}</span>
                <div :id="'collapsex_'+i.id" class="collapse" :aria-labelledby="'heading_'+i.id">
                  <div v-if="i.id==1" id="map"></div>
                  <div v-if="i.id==1" id="status"></div>
                  <div v-if="i.id==2">
                    <div class>
                      <span class="btn btn-outline btn-file">
                        <i class="icon ion-md-image"></i>
                        <input
                          type="file"
                          name="file[]"
                          class="form-control-file"
                          accept="image/*"
                          @change="setAttachmentImage($event)"
                          multiple
                        />
                      </span>
                      <span class="btn btn-outline btn-file">
                        <i class="icon ion-md-videocam"></i>
                        <input
                          type="file"
                          name="files[]"
                          class="form-control-file"
                          accept="video/*"
                          @change="setAttachmentVideo($event)"
                          multiple
                        />
                      </span>
                    </div>
                    <ul class="list-group list-group-flush" v-for="(fi,index) in img" :key="index">
                      <li class="list-group-item">
                        <img
                          v-if="img!=null"
                          :src="fi"
                          class="rounded zoom"
                          :id="'zoom_'+index"
                          @mouseover="zoomOver('zoom_',index)"
                          @mouseleave="zoomOut('zoom_',index)"
                          width="50"
                          height="50"
                          alt
                        />
                        <div class="zoom"></div>
                        <button
                          class="btn btn-outline"
                          v-if="img!==undefined"
                          @click="delItemImage(index)"
                        >x</button>
                      </li>
                    </ul>

                    <ul
                      class="list-group list-group-flush"
                      v-for="(fi1,index1) in vid"
                      :key="`item-${index1}`"
                    >
                      <li class="list-group-item">
                        <div class="embed-responsive embed-responsive-4by3">
                          <video
                            class="embed-responsive-item"
                            v-if="vid!=null"
                            :src="fi1"
                            width="80"
                            height="80"
                            controls
                          />
                        </div>
                        <button
                          class="btn btn-outline"
                          v-if="vid!==undefined"
                          @click="delItemVideo(index1)"
                        >x</button>
                      </li>
                    </ul>
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
                            v-model="objDetail.range_damage"
                            step="1"
                            min="1"
                            max="5"
                            value="1"
                            @change="changeRange($event)"
                          />
                        </div>

                        <label id="rangeText" />
                      </div>
                    </fieldset>

                    <fieldset>
                      <legend>categories</legend>
                      <div class="form-group">
                        <input
                          type="search"
                          class="form-control"
                          placeholder="categories"
                          v-on:keypress.enter="searchAllCategories($event)"
                        />
                      </div>
                      <img
                        v-if="dataCategories===undefined"
                        src="../assets/loading2.gif"
                        class="rounded"
                        width="30"
                        height="30"
                        alt
                      />
                      <ul
                        class="list-group list-group-flush"
                        v-for="(i,index2) in dataCategories"
                        :key="`itemsear-${index2}`"
                      >
                        <li class="list-group-item">
                          <span class="btn btn-outline" @click="addCategories(i)">
                            {{i.name}}
                            <i class="icon ion-md-add"></i>
                          </span>
                        </li>
                      </ul>
                      <ul
                        class="list-group list-group-flush"
                        v-for="(i,inde) in arrayCate"
                        :key="`itemsr-${inde}`"
                      >
                        <li class="list-group-item">
                          <span class="btn btn-outline">
                            <i class="icon ion-md-pricetag"></i>
                            {{i.name}}
                          </span>
                          <span @click="delItemCategories(i)">
                            <i class="icon ion-md-trash"></i>
                          </span>
                        </li>
                      </ul>
                    </fieldset>

                    <fieldset>
                      <legend>peoples</legend>
                      <div class="form-group">
                        <div class="row">
                          <div class="col">
                            <input
                              type="text"
                              v-model="objDetail.quantity_people.child"
                              class="form-control"
                              placeholder="quantity childs"
                            />
                          </div>
                          <div class="col">
                            <input
                              type="text"
                              v-model="objDetail.quantity_people.adult"
                              class="form-control"
                              placeholder="quantity adult"
                            />
                          </div>
                          <div class="col">
                            <input
                              type="text"
                              v-model="objDetail.quantity_people.mayor"
                              class="form-control"
                              placeholder="quantity mayor"
                            />
                          </div>
                        </div>
                      </div>
                    </fieldset>

                    <fieldset>
                      <legend>Explain</legend>
                      <div class="form-group">
                        <textarea
                          v-model="objDetail.fulltext"
                          placeholder="explain"
                          class="form-control"
                          cols="30"
                          rows="10"
                        ></textarea>
                      </div>
                    </fieldset>
                  </div>
                  <div v-if="i.id==4">
                    <fieldset>
                      <legend>Article</legend>

                      <div class="row">
                        <div class="col">
                          <div class="form-group">
                            <input
                              type="text"
                              class="form-control"
                              name="article"
                              @keypress.enter="searchAllArtcles($event)"
                              placeholder="article"
                            />
                          </div>
                          <img
                            v-if="dataArticles===undefined"
                            src="../assets/loading2.gif"
                            class="rounded"
                            width="30"
                            height="30"
                            alt
                          />

                          <div
                            class="row"
                            v-for="(i,indexx2) in dataArticles"
                            :key="`itemsearc-${indexx2}`"
                          >
                            <div class="col">{{i.name}}</div>
                            <div class="col-sm-3">
                              <input
                                type="number"
                                class="form-control"
                                v-model="i.quantity_ask"
                                placeholder="quantity"
                              />
                            </div>
                            <div class="col-sm-3">
                              <span class="btn btn-outline" @click="addArticle(i)">
                                <i class="icon ion-md-add"></i>
                              </span>
                            </div>
                          </div>
                        </div>
                      </div>
                      <div class="row" v-for="(i,index3) in arrayArtcle" :key="`itemart-${index3}`">
                        <div class="col">
                          <span class="btn btn-outline">
                            <i class="icon ion-md-pricetag"></i>
                            {{i.name}}
                          </span>
                        </div>
                        <div class="col-sm-3">
                          <div class="form-group">
                            <input
                              type="number"
                              class="form-control"
                              readonly
                              :value="i.quantity_ask"
                            />
                          </div>
                        </div>
                        <div class="col-sm-3">
                          <span @click="delItemArticle(index3)">
                            <i class="icon ion-md-trash"></i>
                          </span>
                        </div>
                      </div>
                    </fieldset>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <img
              v-if="dataEdit===undefined"
              src="../assets/loading0.gif"
              class="rounded"
              width="30"
              height="30"
              alt
            />
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
            <img
              v-if="dataEdit===undefined"
              src="../assets/loading0.gif"
              class="rounded"
              width="30"
              height="30"
              alt
            />

            <div v-if="dataEdits.articles_data!=undefined" class="table-responsive">
              <table class="table table-striped">
                <thead>
                  <tr>
                    <th>Name</th>
                    <th class="text text-right">Quantity ask</th>
                    <th class="text text-right">Quantity left</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="(j,index) in dataEdits.articles_data" :key="`itmlcq-${index}`">
                    <td>
                      {{j.name}}
                      <table class="table table-striped">
                        <thead>
                          <tr>
                            <th>email</th>
                            <th>quantity_delivery</th>
                            <th>status_delivery_sender</th>
                            <th>status_delivery_reciever</th>
                            <th>status_order</th>
                            <th>created_at</th>
                          </tr>
                        </thead>
                        <tbody>
                          <tr v-for="(y,index1) in j.articles_actors" :key="`itmlcxq-${index1}`">
                            <td v-if="y.quantity_delivery!=0">{{y.users.email}}</td>
                            <td v-if="y.quantity_delivery!=0">-{{y.quantity_delivery}}</td>
                            <td v-if="y.quantity_delivery!=0">
                              <label class="switch">
                                <input type="checkbox" disabled :checked="y.status_delivery_sender" />
                                <span class="slider1 round"></span>
                              </label>
                            </td>
                            <td v-if="y.quantity_delivery!=0">
                              <label
                                class="switch"
                                v-if="y.quantity_delivery!=0 && y.status_order!=='completed'"
                              >
                                <!-- <input
                                  type="checkbox"
                                  :id="'chk_receiver_'+y.sequences_id"
                                  :disabled="y.quantity_left===0"
                                  :checked="y.status_delivery_reciever"
                                  @change="changeStatusDeliveryReceiver('chk_receiver_'+y.sequences_id,y)"
                                />-->
                                <input
                                  type="checkbox"
                                  :id="'chk_receiver_'+y.sequences_id"
                                  :disabled="y.status_order==='finalize'"
                                  :checked="y.status_delivery_reciever"
                                  @change="changeStatusDeliveryReceiver('chk_receiver_'+y.sequences_id,y)"
                                />
                                <span class="slider1 round"></span>
                              </label>
                              <label class="switch" v-if="y.status_order==='completed'">
                                <input
                                  type="checkbox"
                                  :id="'chk_receivers_'+y.sequences_id"
                                  disabled
                                  :checked="y.status_delivery_reciever"
                                />
                                <span class="slider1 round"></span>
                              </label>
                            </td>
                            <!-- <td v-if="y.quantity_delivery!=0">{{y.status_order}}</td> -->
                            <td v-if="y.quantity_delivery!=0">{{y.status_order}}</td>
                            <td v-if="y.quantity_delivery!=0">{{y.created_at | formatDate1 }}</td>
                          </tr>
                        </tbody>
                      </table>
                    </td>
                    <td class="text text-right">{{j.quantity_ask}}</td>
                    <td class="text text-right">{{j.quantity_left}}</td>
                  </tr>
                </tbody>
              </table>
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
import Global from "../global";
// import store from "../store";
import axios from "axios";
// import L from 'leaflet';
declare var L: any;
declare var $: any;
@Component
export default class AddPost extends Vue {
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
  dataEdits: any = {
    articles_data: []
  };
  ID: any;
  REV: any;
  NameDelete: any;
  formEdit: any;

  setAttachmentImage(evt: any) {
    let g = new Global();

    let files = evt.target.files;

    for (let x = 0; x < files.length; x++) {
      let f = files[x];
      if (files.length > this.limitImg) {
        console.log(
          "excedd limit length file: " +
            this.limitImg +
            " current: " +
            files.length
        );
        return;
      }
    }

    let f: any = g.multiFileBase64(evt);
    console.log("image\n");
    console.log(f);
    this.img = f;
    // this.img = g.multiFileBase64(evt);
    // let argsFile: any = {
    //   type: this.typeFile[0],
    //   name: this.img
    // };

    this.attachment.map((i: any) => {
      if (i.type == this.typeFile[0]) {
        i.name = this.img;
      }
    });

    // this.attachment.push(argsFile);
  }

  setAttachmentVideo(evt: any) {
    let g = new Global();
    let files = evt.target.files;

    for (let x = 0; x < files.length; x++) {
      let f = files[x];
      if (files.length > this.limitVid) {
        console.log(
          "excedd limit length file: " +
            this.limitVid +
            " current: " +
            files.length
        );
        return;
      }

      if (f.size > this.limitSizeVid) {
        console.log(
          "file limit excedd: " + this.limitSizeVid + " f: " + f.size
        );
        return;
      }
    }

    let f: any = g.multiFileBase64(evt);
    console.log("video\n");
    console.log(f);
    this.vid = f;
    // let argsFile: any = {
    //   type: this.typeFile[1],
    //   name: this.vid
    // };

    this.attachment.map((i: any) => {
      if (i.type == this.typeFile[1]) {
        i.name = this.vid;
      }
    });

    // this.attachment.push(argsFile);
  }

  addArticle(payload: any) {
    let g = new Global();
    try {
      let val: any;
      val = payload._id;
      let key: any;

      if (
        payload.quantity_ask == null ||
        payload.quantity_ask === undefined ||
        payload.quantity_ask < 1
      ) {
        console.log("quantity_ask should be major that  zero");
        return;
      }

      key = g.objectFindByKey(this.arrayArtcle, "_id", val);
      console.log("ke " + key);
      let keys: any;
      keys = JSON.parse(key);
      if (keys === null) {
        this.arrayArtcle.push({
          article_id: val,
          name: payload.name,
          quantity_left: parseInt(payload.quantity_ask),
          quantity_ask: parseInt(payload.quantity_ask)
        });

        this.dataArticles.map((i: any, k: any) => {
          if (i.id == val) {
            this.dataArticles.splice(k, 1);
          }
        });
      }
    } catch (e) {
      console.log("lol " + e.message);
    }
  }

  addCategories(payload: any) {
    let g = new Global();
    try {
      let val: any;
      val = payload._id;
      let key: any;
      key = g.objectFindByKey(this.arrayCate, "_id", val);
      // console.log("ke " + key);
      let keys: any;
      keys = JSON.parse(key);
      if (keys === null) {
        this.arrayCate.push({
          _id: val,
          name: payload.name
        });

        this.arrayCatePostID.push(val);

        this.dataCategories.map((i: any, k: any) => {
          if (i.id == val) {
            this.dataCategories.splice(k, 1);
          }
        });
      }
    } catch (e) {
      console.log("lol " + e.message);
    }
  }

  delItemArticle(i: any) {
    this.arrayArtcle.splice(i, 1);
  }

  delItemCategories(i: any) {
    this.arrayCate.splice(i, 1);
  }

  delItemImage(i: any) {
    this.img.splice(i, 1);
  }

  delItemVideo(i: any) {
    this.vid.splice(i, 1);
  }

  public searchAllCategories(evt: any) {
    let payload: any = {};
    payload["name"] = evt.target.value;
    const endpoint: any =
      Global.const.CATEGORY_SEARCH_PAGINATE + this.skip + "/" + this.limit;

    this.dataCategories = undefined;
    axios
      .post(endpoint, payload, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success) {
          this.dataCategories =
            response.data.data !== undefined ? response.data.data : [];

          let arrcat = this.arrayCate;
          let arrCatSearch = this.dataCategories;

          let arr0 = arrCatSearch.filter((i: any) => {
            let array_client = arrcat.filter((j: any) => {
              return j._id == i._id;
            });

            if (array_client.length == 0) {
              return i;
            }
          });

          this.dataCategories = arr0;
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  public searchAllArtcles(evt: any) {
    let payload: any = {};
    payload["name"] = evt.target.value;
    const endpoint: any =
      Global.const.ARTICLE_SEARCH_PAGINATE + this.skip + "/" + this.limitArt;

    this.dataArticles = undefined;
    axios
      .post(endpoint, payload, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success) {
          this.dataArticles =
            response.data.data !== undefined ? response.data.data : [];

          let arrcat = this.arrayArtcle;
          let arrCatSearch = this.dataArticles;

          let arr0 = arrCatSearch.filter((i: any) => {
            let array_client = arrcat.filter((j: any) => {
              return j.article_id == i._id;
            });

            if (array_client.length == 0) {
              return i;
            }
          });

          this.dataArticles = arr0;
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  errorMapCallBack(error: any) {
    var strMessage = "";
    // Check for known errors
    switch (error.code) {
      case error.PERMISSION_DENIED:
        strMessage =
          "El acceso a su ubicación está desactivado. " +
          "Cambia tu configuración para volver a encenderla.";
        break;
      case error.POSITION_UNAVAILABLE:
        strMessage =
          "Los datos de los servicios de ubicación son " +
          "actualmente no disponible.";
        break;
      case error.TIMEOUT:
        strMessage =
          "La ubicación no pudo determinarse " +
          "dentro de un período de tiempo de espera especificado.";
        break;
      default:
        strMessage = "";
        break;
    }
    let d: any = document.getElementById("status");
    d.innerHTML = strMessage;
  }

  sucessMapCallBack(pos: any) {
    this.mapInit(pos);
  }

  mapInit(pos: any) {
    try {
      let container = L.DomUtil.get("map");
      if (container != null) {
        container._leaflet_id = null;
      }
      let lat = pos.coords.latitude;
      let lng = pos.coords.longitude;
      // let radius = pos.coords.accuracy;
      this.latLong = {
        latitude: lat,
        longitude: lng
      };
      // console.log("lat: " + lat + " long: " + lng + " radius: " + radius);

      // let map1 = L.map("map").setView([lat, lng], 13);

      let map1 = L.map("map", {
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
      console.log("error__map: " + e.message);
    }
  }

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

  send() {
    let promedy_peoples = [
      {
        people: "child",
        quantity: parseInt(this.objDetail.quantity_people.child)
      },
      {
        people: "adult",
        quantity: parseInt(this.objDetail.quantity_people.adult)
      },
      {
        people: "mayor",
        quantity: parseInt(this.objDetail.quantity_people.mayor)
      }
    ];

    if (this.arrayArtcle.length == 0) {
      return;
    }

    this.form = {
      post_data: {
        range_damage: this.objDetail.range_damage,
        atachments: this.attachment,
        promedy_peoples: promedy_peoples,
        full_text: this.objDetail.fulltext,
        category_post_id: this.arrayCatePostID,
        category_post: this.arrayCate,
        lat_long: this.latLong
      },
      articles_data: this.arrayArtcle
    };

    console.log(this.form);
    this.dataEdit = undefined;
    const endpoint: any = Global.const.POST_SAVE;
    axios
      .post(endpoint, this.form, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success === true) {
          this.dataEdit = [];
          $("#createModal").modal("hide");
          // this.initComponent();
          this.getAll();
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  getAll() {
    const endpoint: any =
      Global.const.POST_PAGINATE + this.skip + "/" + this.limit;
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
                if (k1.articles_actors !== null) {
                  x.confirmArticleByActor = true;
                }

                // k1.quantity_left
                if (k1.quantity_left === 0) {
                  x.confirmArticleByActor = false;
                }

                return k1;
              });

              // x.post_data.atachments = y;

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
          // this.columns = undefined;
          this.data = [];
        }
      });
  }

  public edit(id: any, _rev: any) {
    const endpoint: any = Global.const.POST_EDIT + id;
    this.dataEdit = undefined;
    this.formEdit = undefined;
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
          this.formEdit = data;
          // this.formEdit.description = data.description;
          this.ID = id;
          this.REV = _rev;
          this.NameDelete = data.name;
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  changeStatusDeliveryReceiver(id: any, payload: any) {
    console.log("#ELEMENT ID: " + id);
    let d: any = document.getElementById(id);
    let evt = d.checked;
    console.log("#evt: " + evt);
    // console.log(payload);
    // payload.articles_actors.map((i: any) => {
    // i.status_delivery_reciever = evt;
    // return i;
    // });
    payload.status_delivery_reciever = evt;
    console.log("---#payload: ");
    console.log(payload);
  }

  update() {
    console.log("send....");
    // console.log(this.dataEdits);
    let dataEdit: any = this.formEdit;
    this.formEdit.articles_data = this.dataEdits.articles_data;
    console.log(dataEdit);

    this.dataEdit = undefined;
    console.log(this.formEdit);

    const endpoint: any = Global.const.POST_UPDATE + this.ID + "/" + this.REV;
    // return;
    axios
      .put(endpoint, this.formEdit, {
        headers: this.headers
      })
      .then((response: any) => {
        if (response.data.success) {
          this.dataEdit = [];
          $("#editModal").modal("hide");
          this.dataEdits.articles_data = [];
          this.getAll();
        } else {
          // this.errors.push(response.data.message);
        }
      });
  }

  mapRender() {
    if (navigator.geolocation) {
      this.browserSupportFlag = true;
      navigator.geolocation.getCurrentPosition(
        this.sucessMapCallBack,
        this.errorMapCallBack,
        this.options
      );
    } else {
      alert(
        "Tu navegador no soporta la geolocalización, actualiza tu navegador."
      );
    }
  }

  changeRange(evt: any) {
    let d: any = document.getElementById("rangeText");
    d.innerText = this.rangeValues[evt.target.value];
  }

  defaulRange() {
    let d: any = document.getElementById("rangeText");
    d.innerText = this.rangeValues[1];
  }

  changeDamageRange(val: any, id: any) {
    let d: any = document.getElementById("rangeTextRead_" + id);
    d.innerText = this.rangeValues[val];
  }

  zoomImage() {}

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

  created() {
    //  this.zoomImage();
  }

  beforeDestroy() {
    // this.zoomImage();
  }

  mounted() {
    this.getAll();
    this.defaulRange();
    this.mapRender();
    this.initPermission();
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

<style scoped>
#map {
  /* width: 960px; */
  height: 400px;
}

#map1 {
  /* width: 960px; */
  height: 400px;
}

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

/* img.zoom {
   width: 350px;
  height: 200px; 
  transition: all 0.2s ease-in-out;
  -webkit-transition: all 0.2s ease-in-out;
  -moz-transition: all 0.2s ease-in-out;
  -o-transition: all 0.2s ease-in-out;
  -ms-transition: all 0.2s ease-in-out;
} */

.transition {
  -webkit-transform: scale(1.8);
  -moz-transform: scale(1.8);
  -o-transform: scale(1.8);
  transform: scale(1.8);
}

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

.switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 34px;
}

/* input {
  opacity: 0;
  width: 0;
  height: 0;
} */

.switch input[type="checkbox"] {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider1 {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  -webkit-transition: 0.4s;
  transition: 0.4s;
}

.slider1:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  -webkit-transition: 0.4s;
  transition: 0.4s;
}

input:checked + .slider1 {
  background-color: #2196f3;
}

input:focus + .slider1 {
  box-shadow: 0 0 1px #2196f3;
}

input:checked + .slider1:before {
  -webkit-transform: translateX(26px);
  -ms-transform: translateX(26px);
  transform: translateX(26px);
}

/* Rounded sliders */
.slider1.round {
  border-radius: 34px;
}

.slider1.round:before {
  border-radius: 50%;
}
</style>
