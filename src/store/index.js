import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export const storeConfig = {
  state: {
    currentScenario: null,
    currentLocation: null,
    currentFacility: null,
    currentChemical: null,
    waters: null,
    facilities: null
  },
  getters: {
    currentLocation: state => state.currentLocation,
    currentFacility: state => state.currentFacility,
    currentChemical: state => state.currentChemical,
    waters: state => state.waters,
    facilities: state => state.facilities
  },
  actions: {
    setCurrentFacility({ commit }, o) {
      commit('setCurrentFacility', o);
    },
    setCurrentLocation({ commit }, o) {
      commit('setCurrentLocation', o);
    },
    setCurrentChemical({ commit }, o) {
      commit('setCurrentChemical', o);
    },
    setWaters({ commit }, o) {
      commit('setWaters', o);
    },
    setFacilities({ commit }, o) {
      commit('setFacilities', o);
    },
  },
  mutations: {
    setCurrentFacility(state, o) {
      state.currentFacility = o;
    },
    setCurrentLocation(state, o) {
      state.currentLocation = o;
    },
    setCurrentChemical(state, o) {
      state.currentChemical = o;
    },
    setWaters(state, o) {
      state.waters = o;
    },
    setFacilities(state, o) {
      state.facilities = o;
    }
  }
};


export default new Vuex.Store(storeConfig);
