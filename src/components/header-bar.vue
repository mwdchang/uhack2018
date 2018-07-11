<template>
  <div>
    <h4>Where is the Water Quality Worst?</h4>
    <div>
      <span>{{currentLocation}} - {{currentChemical}} - {{currentFacility}}</span>
      <span @click="reset()"> **Back** </span>
    </div>
  </div>
</template>

<script>
import {  mapActions, mapGetters } from 'vuex';

export default {
  name: 'HeaderBar',
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
