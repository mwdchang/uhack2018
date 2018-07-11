<template>
  <div id="mapid" class="map-panel">
  </div>
</template>

<script>
  import L from 'leaflet';
  import 'leaflet/dist/leaflet.css';
  import {  mapActions, mapGetters } from 'vuex';
  import Mock from '../util/mock.js';
  import API from '../util/api.js';

  /* template */
  export default {
    name: 'map-panel',

    mounted() {
      this.cluster = L.layerGroup();

      this.createMap();
      this.addLocationMarkers();
    },
    computed: {
      ...mapGetters({
        currentLocation: 'currentLocation',
        currentFacility: 'currentFacility',
        currentChemical: 'currentChemical',
        waters: 'waters',
        facilities: 'facilities',
      })
    },
    watch: {
      currentLocation: function changed() {
        if (this.currentLocation !== null) {
          this.locationChanged();
        } else {
          this.cluster.clearLayers();
        }
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

      addLocationMarkers() {
        API.getWater().then(d=>d.json()).then(waters => {
          waters.forEach( water => {
            let divIcon =L.divIcon({
              className:'water-marker-div-icon',
              html:'<i class="fa fa-tint fa-2x"></i><span class="location-marker-text">' + water.name + '</span>',
              iconAnchor:[14,14],
              iconSize:null,
              popupAnchor:[0,0]
            });

            L.marker([water.lat, water.lon], {
              icon: divIcon
            }).addTo(this.map);
          });
        });

        API.getFacilities().then(d=>d.json()).then(facilities => {
          facilities.forEach( facility => {
            let divIcon =L.divIcon({
              className:'facility-marker-div-icon',
              //html:'<i class="fa fa-industry fa-2x"></i><span class="location-marker-text">' + facility.name + '</span>',
              html:'<i class="fa fa-industry fa-2x"></i>',
              iconAnchor:[14,14],
              iconSize:null,
              popupAnchor:[0,0]
            });

            L.marker([facility.lat, facility.lon], {
              icon: divIcon,
              opacity: 0.3,
            }).addTo(this.map);
          });
        });
      },

      /* User click a water location spark line */
      locationChanged() {
        const DIST = 30 * 1000; // 100 km

        this.cluster.clearLayers();

        const loc = this.currentLocation;
        const origin = new L.LatLng(loc.lat, loc.lon);

        this.map.panTo(origin);

        // TODO: filtering by distance and checmical
        this.facilities.forEach( facility => {
          const fLoc = new L.LatLng(facility.lat, facility.lon);
          if (fLoc.distanceTo(origin) < DIST) {
            const pointList = [origin, fLoc];
            const line = new L.polyline(pointList, {
              color: '#F0F',
              weight: 2,
              opacity: 0.5,
              smoothFactor: 1
            });
            this.cluster.addLayer(line);
          }
        })
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

  height: 800px;
  width: 400px;
}

.water-marker-div-icon {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  color: blue;
  white-space: nowrap;
}

.facility-marker-div-icon {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  color: brown;
  white-space: nowrap;
}

.location-marker-text {
  padding-left: 5px;
  color: black;
  font-size: 14px;
  font-weight: 900;
}

</style>
