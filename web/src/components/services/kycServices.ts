import Vue from 'vue';
import Vuex from 'vuex';
import Global from "../../global"
import axios from "axios";

Vue.use(Vuex);

export default class KycService {

    public static getListKyc() {
        let http = Global.const.GET_KYC_LIST;
        return axios.get(http).then((response:any) => this.evalReponseDomain(response))
    }

    public static getKycListUser() {
        let http = Global.const.GET_KYC_USER;
        return axios.get(http).then((response:any) => this.evalReponseDomain(response))
    }

    public static getKycListUserByID(id: any) {
        let http = Global.const.GET_KYC_USER_LIST_ID + id;
        return axios.get(http).then((response:any) => this.evalReponseDomain(response))
    }

    public static getKycUser() {
        let data: any = []
        return this.getListKyc().then((payloadKycList:any) => {
            return KycService.getKycListUser().then((payloadKycUser:any) => {
                if (payloadKycUser) {
                    let documentsUser = payloadKycList.map((val: any) => {
                        let ind = payloadKycUser.attachment.findIndex((obj: any) => val.id == obj.document_id)
                        if (ind >= 0)
                            val.user_attachments = payloadKycUser.attachment[ind]
                        return val
                    })
                    payloadKycUser.attachment = documentsUser
                    data = payloadKycUser
                } else {
                    data = {
                        attachment: payloadKycList
                    }
                }
                console.log(data)
                return data
            })
        })
    }

    public static postKycUser(payload: object) {
        let http = Global.const.POST_KYC_USER_DOCUMENTS;
        let headers = Global.cors;
        return axios.post(http, payload, headers).then((response:any) => {
            return this.evalReponseDomain(response)
        });
    }

    public static putKycUser(payload: object) {
        let http = Global.const.PUT_KYC_USER_DOCUMENTS;
        let headers = Global.cors;
        return axios.put(http, payload, headers).then((response:any) => {
            return this.evalReponseDomain(response)
        });
    }

    public static putKycUserAdminValidate(payload: object) {
        let http = Global.const.PUT_KYC_USER_DOCUMENTS_VALIDATE;
        let headers = Global.cors;
        return axios.put(http, payload, headers).then((response:any) => {
            return this.evalReponseDomain(response)
        });
    }

    public static evalReponseDomain(payload: any) {
        let responseApi: any = false

        if (payload.data.success) {
            responseApi = payload.data.data;
        } else {
            console.log(payload.message)
        }
        return responseApi;
    }

}