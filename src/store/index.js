import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);


export const storeConfig = {
  state: {
    currentScenario: null,
    currentLocation: null,
    currentFacility: null
  },
  getters: {
    currentLocation: state => state.currentLocation,
    currentFacility: state => state.currentFacility
  },
  actions: {
    setCurrentFacility({ commit }, o) {
      commit('setCurrentFacility', o);
    },
    setCurrentLocation({ commit }, o) {
      commit('setCurrentLocation', o);
    },
  },
  mutations: {
    setCurrentFacility(state, o) {
      state.currentFacility = o;
    },
    setCurrentLocation(state, o) {
      state.currentLocation = o;
    }
  }
};


export default new Vuex.Store(storeConfig);
