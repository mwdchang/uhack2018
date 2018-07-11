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
        const circleMarkerRadius = 3000;
        let circleMarkers = [];
        const mockF = Mock.mockF();
        mockF.forEach( facility => {
          const circleMarker = L.circle([facility.lat, facility.lon], {
            color: 'red',
            fillColor: 'red',
            fillOpacity: 0.5,
            radius: circleMarkerRadius
          }).addTo(this.map);
          circleMarkers.push(circleMarker);
        });

        API.getWater().then(d=>d.json()).then(waters => {
          waters.forEach( water => {
            const circleMarker = L.circle([water.lat, water.lon], {
              color: 'red',
              fillColor: 'red',
              fillOpacity: 0.5,
              radius: circleMarkerRadius
            }).addTo(this.map);
            circleMarkers.push(circleMarker);
          });



        })

        const mockW = Mock.mockW();
        mockW.forEach( water => {
          const circleMarker = L.circle([water.lat, water.lon], {
            color: 'blue',
            fillColor: 'blue',
            fillOpacity: 0.5,
            radius: circleMarkerRadius
          }).addTo(this.map);
          circleMarkers.push(circleMarker);
        });

        /* adjust circle marker radius depending on zoom level */
        let myZoom = {
          start:  this.map.getZoom(),
          end: this.map.getZoom()
        };
        this.map.on('zoomstart', () => {
          myZoom.start = this.map.getZoom();
        });
        this.map.on('zoomend', () => {
          myZoom.end = this.map.getZoom();
          const diff = myZoom.start - myZoom.end;
          circleMarkers.forEach(circleMarker => {
            if (diff > 0) {
              circleMarker.setRadius(circleMarker.getRadius() * 2);
            } else if (diff < 0) {
              circleMarker.setRadius(circleMarker.getRadius() / 2);
            }
          })
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

</style>
