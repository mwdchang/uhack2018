<template>
  <div class="side-panel">
    <!-- Show waters -->
    <div v-if="!currentLocation">
      <div class="header"><i class="fa fa-tint"></i> <strong>Water Measures</strong>
        <br>
        <small>Most increase in Arsenic, Chromium, Lead</small>
      </div>
      <div v-for="(spark) of waterData500" :key="spark.vid">
        <div @click="switchLocation(spark)">
          <spark-line :data="spark"/>
        </div>
      </div>
    </div>

    <!-- Show facilities -->
    <div v-if="currentLocation">
      <div class="header">
        <i class="fa fa-industry"></i> <strong>Facility Measures</strong>
        <br>
        <small>Most increase in {{currentLocation.chemical}} from facilities within {{distanceFilter}}km</small>
      </div>
      <div v-for="(spark) of facilityData500" :key="spark.vid">
        <div @click="switchFacility(spark)">
          <spark-line :data="spark"/>
        </div>
      </div>

    </div>
  </div>
</template>


<script>
import _ from 'lodash';
import {  mapActions, mapGetters } from 'vuex';
import SparkLine from './spark-line.vue';

export default {
  name: 'side-panel',
  components: {
    SparkLine
  },
  mounted() {
  },
  computed: {
    ...mapGetters({
      currentLocation: 'currentLocation',
      currentFacility: 'currentFacility',
      currentChemical: 'currentChemical',
      filterdFacilities: 'filterdFacilities',
      waterData: 'waters',
      facilityData: 'facilities',
      distanceFilter: 'distanceFilter'
    }),
    waterData500: function() {
      return _.take(this.waterData, 500);
    },
    facilityData500: function() {
      return _.take(this.filterdFacilities, 500);
    }
  },
  methods: {
    ...mapActions({
      setCurrentLocation: 'setCurrentLocation',
      setCurrentFacility: 'setCurrentFacility',
      setCurrentChemical: 'setCurrentChemical',
    }),
    switchLocation(spark) {
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
  width: 400px;
  height: 100%;
  margin: 2px;

  border: 1px solid #ccc;
  max-height: 800px;
  overflow-y: scroll;
}

.header {
  text-align: left;
  font-size: 120%;
  padding-left: 5px;
  background: #E5E5E5;
}

small {
  font-size: 80%;
}
</style>
