<template>
  <div class="modal" style="opacity:0.9" ref="modal"> 
    <div class="modal-dialog modal-lg" v-if="currentFacility"> 
      <div class="modal-content">
        <div class="modal-header">
          <strong>Facility {{currentFacility.name}} </strong>
        </div>
        <div class="modal-body">
        </div>
        <div class="modal-footer">
          <button class="btn btn-primary btn-sm" v-on:click="cancel()">Cancel</button>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import {  mapActions, mapGetters } from 'vuex';
import * as d3 from 'd3';

export default {
  name: 'FacilityDialog',
  mounted() {
    this.showDialog = true;
  },
  computed: {
    ...mapGetters({
      currentLocation: 'currentLocation',
      currentFacility: 'currentFacility',
      currentChemical: 'currentChemical'
    })
  },
  watch: {
    currentFacility: function changed() {
      if (this.currentFacility) {
        d3.select(this.$refs.modal).style('display', 'block');
      }
    }
  },
  methods: {
    cancel() {
      d3.select(this.$refs.modal).style('display', 'none');
    }
  }

}
</script>



