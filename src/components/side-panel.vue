<template>
  <div class="side-panel">
    <div>Side panel stub</div>
    <div v-for="spark of sparkData">
      <div @click="switchLocation(spark)">
        <spark-line :data="spark"/>
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
    sparkData: []
  }),
  mounted() {
    this.sparkData = Mock.mockdata();
  },
  methods: {
    ...mapActions({
      setCurrentLocation: 'setCurrentLocation',
      setCurrentFacility: 'setCurrentFacility',
      setCurrentChemical: 'setCurrentChemical'
    }),
    switchLocation(spark) {
      this.setCurrentLocation(spark.name);
      this.setCurrentChemical(spark.chemical);
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
