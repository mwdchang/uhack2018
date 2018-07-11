<template>
  <div class="side-panel">
    <!-- Show waters -->
    <div v-if="!currentLocation">
      <div>Water Measures</div>
      <div v-for="(spark, idx) of waterData" :key="idx">
        <div @click="switchLocation(spark)">
          <spark-line :data="spark"/>
        </div>
      </div>
    </div>

    <!-- Show facilities -->
    <div v-if="currentLocation">
      <div>Facility Measures</div>
      <div v-for="(spark, idx) of facilityData" :key="idx">
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
import API from '../util/api.js';

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
    API.getWater().then(d=>d.json()).then( waters => {
      waters.forEach (w => {
        const lastIdx = (w.data.length - 1);
        const delta = w.data[lastIdx] - w.data[0];
        w.delta = delta;
      });

      this.waterData = _.sortBy(waters, (d) => -d.delta);
    });

    API.getFacilities().then(d=>d.json()).then( facilities => {
      // TODO: FAKE
      facilities.forEach (f => {
        f.data = [];
        for (let i=0; i < 10; i++) {
          f.data.push( Math.random() * 10);
        }
        const lastIdx = f.data.length - 1;
        const delta = f.data[lastIdx] - f.data[0];
        f.delta = delta;
      });

      this.facilityData = _.sortBy(facilities, (d) => -d.delta);
    });
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
</style>
