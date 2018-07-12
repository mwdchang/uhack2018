<template>
  <div v-if="waterData && facilityData">
    <header-bar/>
    <div class="container">
      <side-panel/>
      <map-panel/>
    </div>
    <facility-dialog/>
  </div>
</template>

<script>
import _ from 'lodash';
import {  mapActions, mapGetters } from 'vuex';

import SidePanel from './side-panel.vue';
import HeaderBar from './header-bar.vue';
import MapPanel from './map-panel.vue';
import FacilityDialog from './facility-dialog.vue';
 
import API from '../util/api.js';


/* template */
export default {
  name: 'home',
  components: {
    SidePanel, HeaderBar, MapPanel, FacilityDialog
  },
  mounted() {
    API.getWater().then(d=>d.json()).then( waters => {
      waters.forEach (w => {
        const lastIdx = (w.data.length - 1);
        const delta = w.data[lastIdx] - w.data[0];
        w.delta = delta;
      });
      this.setWaters(_.sortBy(waters, (d) => -d.delta));
    });

    API.getFacilities().then(d=>d.json()).then( facilities => {
      // TODO: FAKE
      facilities.forEach (f => {
        const lastIdx = f.data.length - 1;
        const delta = f.data[lastIdx] - f.data[0];
        f.delta = delta;
      });
      this.setFacilities( _.sortBy(facilities, (d) => -d.delta));
    });

  },
  computed: {
    ...mapGetters({
      currentLocation: 'currentLocation',
      currentFacility: 'currentFacility',
      waterData: 'waters',
      facilityData: 'facilities'
    })
  },
  methods: {
    ...mapActions({
      setCurrentLocation: 'setCurrentLocation',
      setCurrentFacility: 'setCurrentFacility',
      setCurrentChemical: 'setCurrentChemical',
      setWaters: 'setWaters',
      setFacilities: 'setFacilities'
    })

  }
}
</script>

<style>

.container {
  display: flex;
  flex-direction: row;
  height: 100%;
  width: 100%;
}

</style>
