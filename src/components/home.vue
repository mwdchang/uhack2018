<template>
  <div>
    <header-bar/>
    <div class="container">
      <side-panel/>
      <map-panel/>
    </div>
    <facility-dialog/>
  </div>
</template>

<script>
import {  mapActions, mapGetters } from 'vuex';

import SidePanel from './side-panel.vue';
import HeaderBar from './header-bar.vue';
import MapPanel from './map-panel.vue';
import FacilityDialog from './facility-dialog.vue';


/* template */
export default {
  name: 'home',
  components: {
    SidePanel, HeaderBar, MapPanel, FacilityDialog
  },
  mounted() {


  },
  computed: {
    ...mapGetters({
      currentLocation: 'currentLocation',
      currentFacility: 'currentFacility'
    })
  },
  method: {
    ...mapActions({
      setCurrentLocation: 'setCurrentLocation',
      setCurrentFacility: 'setCurrentFacility',
      setCurrentChemical: 'setCurrentChemical'
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
