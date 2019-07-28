import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    isLoggedIn: !!localStorage.getItem('token'),
    search: '',
  },
  mutations: {
    loginUser(state) {
      state.isLoggedIn = true;
    },
    logoutUser(state) {
      state.isLoggedIn = false;
    },
    searching(state, search) {
      state.search = search
    },
  },
  actions: {
    searching(context,search){
      context.commit('searching',search)
    }
  },
  getters:{
    getSearch(state){
      return state.search
    }
  }
});
