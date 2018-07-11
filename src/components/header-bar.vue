<template>
  <div>
    <div v-if="currentLocation === null">
      <h4 style="display:flex; flex-direction:row; align-items:center; justify-content: center">
        <span>Where is the Water Quality Worst?</span>
      </h4>
    </div>

    <div v-if="currentLocation !== null">
      <h4 style="display:flex; flex-direction:row; align-items:center; justify-content: center">
        <span>
          <i class="fa fa-caret-left" @click="reset()"></i>
          Back
        </span>

        <span>Who could be responsible for {{currentLocation.chemical}}</span>
        &nbsp; 
        <spark-line :data="currentLocation" :summary="true" style="width: 140px"/>
        &nbsp;
        <span> increase in {{ currentLocation.name }} ?</span>
      </h4>
    </div>

    <!--
    <div>
      <span>
        <span v-if="currentLocation">{{currentLocation.name}}</span> - 
        <span v-if="currentChemical">{{currentChemical}}</span> - 
        <span v-if="currentFacility">{{currentFacility.name}}</span>
      </span>
      <button v-if="currentLocation" class="btn btn-primary btn-sm" @click="reset()">Back</button>
    </div>
    -->
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
</style>
