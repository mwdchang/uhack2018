import Vue from 'vue';
import Vuex from 'vuex';
import L from 'leaflet';

Vue.use(Vuex);

export const storeConfig = {
  state: {
    currentScenario: null,
    currentLocation: null,
    currentFacility: null,
    currentChemical: null,
    waters: null,
    facilities: null,
    filterdFacilities: null
  },
  getters: {
    currentLocation: state => state.currentLocation,
    currentFacility: state => state.currentFacility,
    currentChemical: state => state.currentChemical,
    waters: state => state.waters,
    facilities: state => state.facilities,
    filterdFacilities: state => state.filterdFacilities
  },
  actions: {
    setCurrentFacility({ commit }, o) {
      commit('setCurrentFacility', o);
    },
    setCurrentLocation({ commit, state }, o) {
      commit('setCurrentLocation', o);

      // Auto compute filtered
      const DIST = 30 * 1000; // filter distance
      const r = [];
      const loc = state.currentLocation;
      const origin = new L.LatLng(loc.lat, loc.lon);
      state.facilities.forEach(facility => {
          const fLoc = new L.LatLng(facility.lat, facility.lon);
          if (fLoc.distanceTo(origin) < DIST) {
            r.push(facility);
          }
      });
      commit('setFilterdFacilities', r);
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
    },
    setFilterdFacilities(state, o) {
      state.filterdFacilities = o;
    }
  }
};


export default new Vuex.Store(storeConfig);
