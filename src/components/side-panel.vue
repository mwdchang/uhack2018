<template>
  <div class="side-panel">

    <!-- Show waters -->
    <div v-if="!currentLocation">
      <div>Water Measures</div>
      <div v-for="spark of waterData" :key="spark.id">
        <div @click="switchLocation(spark)">
          <spark-line :data="spark"/>
        </div>
      </div>
    </div>

    <!-- Show facilities -->
    <div v-if="currentLocation">
      <div>Facility Measures</div>
      <div v-for="spark of facilityData" :key="spark.id">
        <div @click="switchFacility(spark)">
          <spark-line :data="spark"/>
        </div>
      </div>

    </div>
  </div>
</template>


<script>
import {  mapActions, mapGetters } from 'vuex';

import SparkLine from './spark-line.vue';
import Mock from '../util/mock.js';

export default {
  name: 'side-panel',
  components: {
    SparkLine
  },
  data: () => ({
    waterData: [],
    facilityData: []
  }),
  mounted() {
    this.waterData = Mock.mockW();
    this.facilityData = Mock.mockF();
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
    switchLocation(spark) {
      console.log('switching');
      this.setCurrentLocation(spark);
      this.setCurrentChemical(spark.chemical);
    },
    switchFacility(spark) {
      this.setCurrentFacility(spark);
    }
  }
}
</script>

<style>
.side-panel {
  box-sizing: border-box;
  width: 300px;
  height: 100%;
  margin: 2px;

  border: 1px solid #ccc;
}
</style>
