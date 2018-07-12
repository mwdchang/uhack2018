<template>
  <div class="side-panel">
    <!-- Show waters -->
    <div v-if="!currentLocation">
      <div class="header"><i class="fa fa-tint"></i> Water Measures</div>
      <div v-for="(spark, idx) of waterData500" :key="idx">
        <div @click="switchLocation(spark)">
          <spark-line :data="spark"/>
        </div>
      </div>
    </div>

    <!-- Show facilities -->
    <div v-if="currentLocation">
      <div class="header"><i class="fa fa-industry"></i> Facility Measures</div>
      <div v-for="(spark, idx) of facilityData500" :key="idx">
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
      waterData: 'waters',
      facilityData: 'facilities'
    }),
    waterData500: function() {
      return _.take(this.waterData, 500);
    },
    facilityData500: function() {
      return _.take(this.facilityData, 500);
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
  width: 350px;
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
}
</style>
