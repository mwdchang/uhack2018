<template>
  <div>
    <h4>Where is the Water Quality Worst?</h4>
    <div>{{currentLocation}} - {{currentFacility}}</div>
  </div>
</template>

<script>
import {  mapActions, mapGetters } from 'vuex';

export default {
  name: 'HeaderBar',
  computed: {
    ...mapGetters({
      currentLocation: 'currentLocation',
      currentFacility: 'currentFacility'
    })
  },
}
</script>

<style>
</style>
