<template>
  <div id="mapid" class="map-panel">
  </div>
</template>

<script>
  import L from 'leaflet';
  import 'leaflet/dist/leaflet.css';
  import {  mapActions, mapGetters } from 'vuex';
  import Mock from '../util/mock.js';


  /* template */
  export default {
    name: 'map-panel',

    mounted() {
      this.cluster = L.layerGroup();

      this.createMap();
    },
    computed: {
      ...mapGetters({
        currentLocation: 'currentLocation',
        currentFacility: 'currentFacility',
        currentChemical: 'currentChemical'
      })
    },
    watch: {
      currentLocation: function changed() {
        this.locationChanged();
      }
    },
    methods: {
      ...mapActions({
        setCurrentLocation: 'setCurrentLocation',
        setCurrentFacility: 'setCurrentFacility',
        setCurrentChemical: 'setCurrentChemical'
      }),
      createMap() {
        this.map = L.map('mapid').setView([43.6532, -79.3832], 8);
        this.tileLayer = L.tileLayer(
          'https://cartodb-basemaps-{s}.global.ssl.fastly.net/rastertiles/voyager/{z}/{x}/{y}.png',
          {
              maxZoom: 18,
              attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>, &copy; <a href="https://carto.com/attribution">CARTO</a>',
          }
        );
        this.tileLayer.addTo(this.map);
      },

      /* User click a water location spark line */
      locationChanged() {
        this.cluster.clearLayers();

        const loc = this.currentLocation;
        const origin = new L.LatLng(loc.lat, loc.lon);

        // TODO: filtering by distance and checmical
        const mockF = Mock.mockF();
        mockF.forEach( facility => {
          const fLoc = new L.LatLng(facility.lat, facility.lon);
          const pointList = [origin, fLoc];

          const line = new L.polyline(pointList, {
            color: '#F80',
            weight: Math.random()*15,
            opacity: 0.5,
            smoothFactor: 1
          });
          // firstpolyline.addTo(this.map);
          this.cluster.addLayer(line);
        });
        this.cluster.addTo(this.map);

      }
    }
  }
</script>

<style>
.map-panel {
  flex: 1;
  display: block;
  box-sizing: border-box;
  border: 1px solid #ccc;
  margin: 2px;

  height: 400px;
  width: 400px;
}

</style>
