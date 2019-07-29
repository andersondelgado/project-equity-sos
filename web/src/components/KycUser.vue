<template>
  <div class="row">
    <!-- {{countryList}} -->
    <!-- User Info -->
    <div class="col-md-6">
      <div class="card">
        <div class="card-header card-header-icon card-header-rose">
          <div class="card-icon">
            <i class="material-icons">perm_identity</i>
          </div>
          <h4 class="card-title">
            Edit Profile -
            <small class="category">Complete your profile</small>
          </h4>
        </div>
        <div class="card-body">
          <form v-on:submit.prevent="onSubmitForm()">
            <div class="row">
              <div class="col-md-5">
                <div class="form-group bmd-form-group">
                  <label class="bmd-label-floating">Company Name</label>
                  <input
                    v-validate="'required|min:3|max:30'"
                    name="company_name"
                    v-model="form.company_name"
                    type="text"
                    class="form-control"
                  />
                  <span>{{ errors.first('company_name') }}</span>
                </div>
              </div>
              <div class="col-md-5">
                <div class="form-group bmd-form-group">
                  <label class="bmd-label-floating">Country Company</label>
                  <select
                    v-validate="'required'"
                    v-model="form.company_country_id"
                    name="company_country_id"
                    class="form-control"
                    data-size="7"
                    data-style="btn btn-primary btn-round"
                    title="Single Select"
                  >
                    <option disabled selected>Country Company</option>
                    <option v-for="i in countryList" :key="i.id">{{i.country}}</option>
                  </select>
                  <span>{{ errors.first('company_country_id') }}</span>
                </div>
              </div>
            </div>
            <div class="row">
              <div class="col-md-4">
                <div class="form-group bmd-form-group">
                  <label class="bmd-label-floating">Fist Name</label>
                  <input
                    v-validate="'required|min:3|max:20'"
                    name="name"
                    v-model="form.name"
                    type="text"
                    class="form-control"
                  />
                  <span>{{ errors.first('name') }}</span>
                </div>
              </div>
              <div class="col-md-4">
                <div class="form-group bmd-form-group">
                  <label class="bmd-label-floating">Last Name</label>
                  <input
                    v-validate="'required|min:3|max:20'"
                    name="last_name"
                    v-model="form.last_name"
                    type="text"
                    class="form-control"
                  />
                  <span>{{ errors.first('last_name') }}</span>
                </div>
              </div>
              <div class="col-md-4">
                <label class="bmd-label-floating">Born Date</label>
                <input
                  type="text"
                  v-validate="'required|date_format:dd/MM/yyyy'"
                  name="dob"
                  v-model="form.dob"
                  class="form-control datepicker"
                />
                <span>{{ errors.first('dob') }}</span>
              </div>
            </div>
            <div class="row">
              <div class="col-md-4">
                <div class="form-group bmd-form-group">
                  <img
                    v-if="countryList===undefined"
                    src="../assets/loading0.gif"
                    class="rounded"
                    width="30"
                    height="30"
                    alt
                  />
                  <select
                    v-validate="'required'"
                    v-model="form.country_id"
                    name="country_id"
                    class="form-control"
                    data-size="7"
                    data-style="btn btn-primary btn-round"
                    title="Single Select"
                  >
                    <option disabled selected>Country</option>
                    <option v-for="k in countryList" :key="k.id">{{k.country}}</option>
                  </select>
                  <span>{{ errors.first('country_id') }}</span>
                </div>
              </div>

              <div class="col-md-7">
                <div class="form-group bmd-form-group">
                  <label class="bmd-label-floating">Adress</label>
                  <input
                    v-validate="'required|min:10|max:200'"
                    name="address"
                    v-model="form.address"
                    type="text"
                    class="form-control"
                  />
                  <span>{{ errors.first('address') }}</span>
                </div>
              </div>
            </div>
            <div class="row">
              <div class="col-md-4">
                <div class="form-group bmd-form-group">
                  <label class="bmd-label-floating">City</label>
                  <input
                    v-validate="'required|min:3|max:30'"
                    name="city"
                    v-model="form.city"
                    type="text"
                    class="form-control"
                  />
                  <span>{{ errors.first('city') }}</span>
                </div>
              </div>
              <div class="col-md-4">
                <div class="form-group bmd-form-group">
                  <label class="bmd-label-floating">address</label>
                  <input
                    v-validate="'required|min:8|max:200'"
                    name="address"
                    v-model="form.address"
                    type="text"
                    class="form-control"
                  />
                  <span>{{ errors.first('address') }}</span>
                </div>
              </div>
              <div class="col-md-4">
                <div class="form-group bmd-form-group">
                  <label class="bmd-label-floating">Postal Code</label>
                  <input
                    v-validate="'required|min:2|max:6'"
                    name="postal_code"
                    v-model="form.postal_code"
                    type="text"
                    class="form-control"
                  />
                  <span>{{ errors.first('postal_code') }}</span>
                </div>
              </div>
            </div>
            <div class="row">
              <div class="col-md-12">
                <div class="form-group">
                  <label>About Me</label>
                  <div class="form-group bmd-form-group">
                    <label class="bmd-label-floating">Sobre mi, Breve Explicación</label>
                    <textarea
                      v-validate="'required|min:10|max:250'"
                      v-model="form.about_person"
                      name="about_person"
                      class="form-control"
                      rows="5"
                    ></textarea>
                    <span>{{ errors.first('about_person') }}</span>
                  </div>
                </div>
              </div>
            </div>
            <div class="row" v-if="userType == 'admin'">
              <div class="col-md-12">
                <h4 class="title">User Actions</h4>
                <div class="col-md-6">
                  <div class="togglebutton">
                    <label>
                      <input
                        type="checkbox"
                        v-model="form.status_user_info"
                        name="status_user_info"
                        @change="changeStateElement($event)"
                      />
                      <span class="toggle"></span>
                      Confirm Kyc Info
                    </label>
                  </div>
                </div>
                <div class="col-md-6">
                  <div class="togglebutton">
                    <label>
                      <input
                        type="checkbox"
                        v-model="form.disabled_user"
                        name="disabled_user"
                        @change="changeStateElement($event)"
                      />
                      <span class="toggle"></span>
                      Disabled User
                    </label>
                  </div>
                </div>
              </div>
            </div>
            <div style="text-align: center;">
              <button
                v-if="userType != 'admin'"
                type="submit"
                class="btn btn-rose pull-right"
              >Update Profile</button>
              <button
                v-if="userType == 'admin'"
                type="submit"
                class="btn btn-rose pull-right"
              >Validate User</button>
              <div class="clearfix"></div>
            </div>
          </form>
        </div>
      </div>
    </div>
    <!-- /User Info -->

    <!-- Kyc Documents -->
    <div class="col-md-5">
      <div class="card">
        <div class="card-header card-header-icon card-header-rose">
          <div class="card-icon">
            <i class="material-icons">assignment_ind</i>
          </div>
          <h4 class="card-title">
            Documents Verifications -
            <small
              class="category"
            >Please Upload your Documents to verify your Identity</small>
          </h4>
        </div>
        <div class="card-body">
          <img
            v-if="kycList===undefined"
            src="../assets/loading0.gif"
            class="rounded"
            width="30"
            height="30"
            alt
          />
          <form name="kyc_documents" v-on:submit.prevent="onSubmitFormDocuments()">
            <div id="accordionKyc" role="tablist">
              <div class="card-collapse" v-for="value in kycList" :key="value.id">
                <div class="card-header" role="tab" id="headingOne">
                  <h5 class="mb-0">
                    <a
                      data-toggle="collapse"
                      :href="'#collapseKyc_'+value.id"
                      aria-expanded="false"
                      :aria-controls="'collapseKyc_'+value.id"
                      class="collapsed"
                    >
                      {{value.name}}
                      <small
                        v-if="value.user_attachments != undefined"
                      >- {{value.user_attachments.status}}</small>
                      <!--<i class="material-icons"></i>-->
                      <i
                        class="material-icons"
                        v-html="verifyStatusAttachmentIcon('iconval'+value.id, value)"
                        :id="'iconval'+value.id"
                      ></i>
                    </a>
                  </h5>
                </div>
                <div
                  :id="'collapseKyc_'+value.id"
                  class="collapse"
                  role="tabpanel"
                  aria-labelledby="headingOne"
                  data-parent="#accordionKyc"
                  style
                >
                  <div class="card-body">
                    <div class="col-md-4 col-sm-4">
                      <div class="fileinput fileinput-new text-center" data-provides="fileinput">
                        <div class="fileinput-new thumbnail">
                          <img
                            v-if="value.user_attachments == undefined"
                            src="../assets/placeholder.jpg"
                            alt
                          />

                          <img
                            v-if="value.user_attachments != undefined"
                            :src="value.user_attachments.attachment_name"
                            class="rounded"
                            :id="'attachment_name'+value.id"
                            @mouseover="zoomOver('zoomx_',value.id)"
                            @mouseleave="zoomOut('zoomx_',value.id)"
                            width="50"
                            height="50"
                            alt
                          />
                        </div>
                        <div class="fileinput-preview fileinput-exists thumbnail"></div>
                        <div v-if="verifyUploadAttachment(value)">
                          <span class="btn btn-rose btn-round btn-file" v-if="userType != 'admin'">
                            <span class="fileinput-new">Select image</span>
                            <span class="fileinput-exists">Change</span>
                            <input
                              type="file"
                              accept="image/*"
                              :name="'document_'+value.id"
                              @change="setAttachmentImage(value.id, $event)"
                            />
                          </span>
                          <a
                            href="#pablo"
                            class="btn btn-danger btn-round fileinput-exists"
                            data-dismiss="fileinput"
                            @click="deletedChangeAttachmentsRepeated(value.id, formAttachments)"
                          >
                            <i class="fa fa-times"></i> Remove
                          </a>
                        </div>
                      </div>
                      <div v-if="userType == 'admin'">
                        <div
                          v-if="value.user_attachments != undefined"
                          class="col-sm-10 checkbox-radios"
                        >
                          <div class="form-check">
                            <div class="form-check">
                              <label class="form-check-label">
                                <input
                                  class="form-check-input"
                                  type="radio"
                                  name="exampleRadios"
                                  @change="chageStatusAttachment(value.id, $event)"
                                  value="approved"
                                /> Approved
                                <span class="circle">
                                  <span class="check"></span>
                                </span>
                              </label>
                              <label class="form-check-label">
                                <input
                                  class="form-check-input"
                                  type="radio"
                                  name="exampleRadios"
                                  @change="chageStatusAttachment(value.id, $event)"
                                  value="rejected"
                                /> Rejected
                                <span class="circle">
                                  <span class="check"></span>
                                </span>
                              </label>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
  <!-- /Kyc Documents -->
</template>


<script lang="ts">
import KycService from "./services/kycServices";
import CountryService from "./services/CountryServices";
import { Component, Prop, Vue } from "vue-property-decorator";
import Global from "../global";

@Component({
  components: {}
})
export default class KycUser extends Vue {
  // fs
  limitImg = 1;
  img: any = [];
  typeFile: any = ["image", "video"];
  extFile: RegExp = /.+(pdf|png|jpe?g)$/;
  // fs

  date: Date = new Date();
  //userType = 'Bussines'
  // userType = "Bussines";
  // userType = "admin";
  userType = "";
  countryList: any = [];
  selectedCountry: any = "";

  kycList: any = [];
  selectedKyc: string = "";

  formAttachments: any = [];
  public form: any = {
    id: "",
    company_name: "",
    company_country_id: "",
    country_id: "",
    name: "",
    last_name: "",
    city: "",
    address: "",
    about_person: "",
    postal_code: "",
    dob: "",
    attachment: "",
    status_user_info: false,
    disabled_user: false
  };

  created() {
    this.getCountrys();
    this.getKycList();
  }

  mounted() {
    this.isKYCAdmin();
  }

  getKycList() {
    this.kycList = undefined;
    KycService.getKycUser().then(payload => {
      console.log(payload);
      if (payload.id != undefined) {
        this.form = payload;
      }
      this.kycList = payload.attachment;
    });
  }

  getCountrys() {
    this.countryList = undefined;
    this.selectedCountry = null;
    CountryService.getListCountry().then(response => {
      this.countryList = response;
    });
  }

  onSubmitForm(event: any) {
    let vm = this;
    let serviceKyc: any = null;

    this.$validator.validate().then(valid => {
        if (valid) {           
            let formData:any = vm.form
                formData.attachment = this.formAttachments
            
            if(this.userType != 'admin'){
                if(formData.id != ""){
                    serviceKyc = KycService.putKycUser(formData)
                }else{
                    serviceKyc = KycService.postKycUser(formData)
                }
            }else{
                if(formData.id){
                    serviceKyc = KycService.putKycUserAdminValidate(formData)
                }
            }
            
            serviceKyc.then((response:any)=>{
                if (response)
                    this.getKycList();
            });
        }else{
            return false;
        }
    });
  }

  setAttachmentImage(id: number, evt: any) {
    let g = new Global();

    let files = evt.target.files;
    console.log(files[0].name);

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

    g.getBase64ImageEncode(files[0]).then((res: any) => {
      let elementsData = {
        document_id: id,
        attachment_name: res,
        status: "pending",
        created_at: this.date
      };

      if (this.formAttachments.length > 0) {
        this.deletedChangeAttachmentsRepeated(id, this.formAttachments);
      }
      this.formAttachments.push(elementsData);
    });
  }

  public deletedChangeAttachmentsRepeated(id: number, arrayData: any) {
    arrayData.forEach((data: any, index: any) => {
      if (data.document_id == id) {
        this.formAttachments.splice(index, 1);
      }
    });
  }

  public verifyUploadAttachment(data: any) {
    let icon = "keyboard_arrow_down";
    if (data["user_attachments"] != undefined) {
      if (/(approved|pending)/.test(data.user_attachments.status)) {
        return false;
      }
    }
    return true;
  }

  public verifyStatusAttachmentIcon(idTarget: any, data: any) {
    let d: any = document.getElementById(idTarget);
    let icon = "keyboard_arrow_down";
    if (data["user_attachments"] != undefined) {
      if (/(approved)/.test(data.user_attachments.status)) {
        icon = "check_circle";
      } else if (/(refused)/.test(data.user_attachments.status)) {
        icon = "cancel";
      } else if (/(pending)/.test(data.user_attachments.status)) {
        icon = "warning";
      }
    }
    if (d) d.innerHTML = icon;
  }

  //Admin Functions
  public chageStatusAttachment(id: number, event: any) {
    let valueOption = event.target.value;
    let elementsData = {
      document_id: id,
      status: valueOption,
      created_at: this.date
    };
    if (this.formAttachments.length > 0) {
      this.deletedChangeAttachmentsRepeated(id, this.formAttachments);
    }
    this.formAttachments.push(elementsData);
  }

  public changeStateElement(event: any) {
    let valueOption = event.target.checked;
    let nameOption: any = event.target.value;

    this.form[nameOption] = event.target.checked;
    console.log(this.form);
  }

  public isKYCAdmin() {
    const h = window.location.pathname;
    let spl=h.split("/main/")
    console.log("url: " + spl[0]+ " 1: "+spl[1]);
    let spl1=spl[1].split("kyc-")
    console.log("##url: " + spl1[0]+ " ##1: "+spl1[1]);
    switch(spl1[1]){
      case "user":
        this.userType = "Bussines";
        break;
        default:
          this.userType = "admin";
          break;
    }
  }
}
</script>
