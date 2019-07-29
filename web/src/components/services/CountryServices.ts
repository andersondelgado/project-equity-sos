import Vue from 'vue';
import Vuex from 'vuex';
import Global from "../../global"
import axios from "axios";

Vue.use(Vuex);

export default class CountryService {

    static data_list:any = []

    public static getListCountry(){
        let http = Global.const.GET_COUNTRY_LIST;
        return axios.get(http).then(response =>{
            return this.evalReponseDomain(response)
        });
    }

    public static getCountryById(){
        return this.data_list;
    }    
    
    public putCountryById(params:object){
    }

    public static postSaveCountry(params:object){

    }

    public static evalReponseDomain(params:any){
        if(params.data.success) {
            return params.data.data;
        }else{
            console.log(params.message)
        }
    }
}