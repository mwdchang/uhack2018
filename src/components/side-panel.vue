<template>
  <div class="side-panel">
    <div>Side panel stub</div>
    <div>
      <div @click="switchLocation()">
        <spark-line/>
      </div>
    </div>
  </div>
</template>


<script>
import {  mapActions, mapGetters } from 'vuex';

import SparkLine from './spark-line.vue';

export default {
  name: 'side-panel',
  components: {
    SparkLine
  },
  methods: {
    ...mapActions({
      setCurrentLocation: 'setCurrentLocation',
      setCurrentFacility: 'setCurrentFacility'
    }),
    switchLocation() {
      this.setCurrentLocation('xyz');
    }
  }
}
</script>

<style>
.side-panel {
  box-sizing: border-box;
  width: 250px;
  height: 100%;
  margin: 2px;

  border: 1px solid #ccc;
}
</style>
