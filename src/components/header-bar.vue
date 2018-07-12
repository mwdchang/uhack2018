<template>
  <div class="header-bar">
    <div v-if="currentLocation === null">
      <h4 style="display:flex; flex-direction:row; align-items:flex-start">
        <span>Where is the Water Quality Worst?</span>
      </h4>
    </div>

    <div v-if="currentLocation !== null">
      <h4 style="display:flex; flex-direction:row; align-items:flex-start">
        <span>
          <i class="fa fa-arrow-circle-left" style="font-size:125%" @click="reset()"></i>
        </span>
        &nbsp; 
        &nbsp; 
        <span>Who could be responsible for {{currentLocation.chemical}}</span>
        &nbsp; 
        <spark-line :data="currentLocation" :summary="true" style="width: 140px"/>
        &nbsp;
        <span> increase in {{ currentLocation.name }} ?</span>
      </h4>
    </div>
  </div>
</template>

<script>
import {  mapActions, mapGetters } from 'vuex';
import SparkLine from './spark-line.vue';

export default {
  name: 'HeaderBar',
  components: {
    SparkLine
  },
  computed: {
    ...mapGetters({
      currentLocation: 'currentLocation',
      currentFacility: 'currentFacility',
      currentChemical: 'currentChemical'
    })
  },
  methods: {
    ...mapActions({
      setCurrentLocation: 'setCurrentLocation',
      setCurrentFacility: 'setCurrentFacility',
      setCurrentChemical: 'setCurrentChemical'
    }),

    reset() {
      this.setCurrentChemical(null);
      this.setCurrentFacility(null);
      this.setCurrentLocation(null);
    }
  }
}
</script>

<style>
.header-bar {
  background: #353752;
  color: #EFEFEF;
}

h4 {
  margin: 0;
  padding: 5px;
  height: 45px;
}
</style>
