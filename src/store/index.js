import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export const storeConfig = {
  state: {
    currentScenario: null,
    currentLocation: null,
    currentFacility: null,
    currentChemical: 'lead'
  },
  getters: {
    currentLocation: state => state.currentLocation,
    currentFacility: state => state.currentFacility,
    currentChemical: state => state.currentChemical
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
    }
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
    }
  }
};


export default new Vuex.Store(storeConfig);
