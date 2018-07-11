<template>
  <div>
    <h4>Where is the Water Quality Worst?</h4>

    <div>
      <span>
        <span v-if="currentLocation">{{currentLocation.name}}</span> - 
        <span v-if="currentChemical">{{currentChemical}}</span> - 
        <span v-if="currentFacility">{{currentFacility.name}}</span>
      </span>
      <button v-if="currentLocation" class="btn btn-primary btn-sm" @click="reset()">Back</button>
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
