<template>
  <div id="mapid" class="map-panel">
  </div>
</template>

<script>
  import L from 'leaflet';
  import _ from 'lodash';
  import * as d3 from 'd3';
  import 'leaflet/dist/leaflet.css';
  import {  mapActions, mapGetters } from 'vuex';

  /* template */
  export default {
    name: 'map-panel',

    mounted() {
      this.cluster = L.layerGroup();
      this.waterMarkers = L.layerGroup();

      this.createMap();
      this.addLocationMarkers(this.waters, 'water', 'tint');
    },
    computed: {
      ...mapGetters({
        currentLocation: 'currentLocation',
        currentFacility: 'currentFacility',
        currentChemical: 'currentChemical',
        waters: 'waters',
        facilities: 'facilities',
        filterdFacilities: 'filterdFacilities'
      })
    },
    watch: {
      currentLocation: function changed() {
        if (this.currentLocation !== null) {
          this.locationChanged();
        } else {
          this.map.setView([43.6532, -79.3832], 8);
          this.cluster.clearLayers();
          this.waterMarkers.clearLayers();
          this.addLocationMarkers(this.waters, 'water', 'tint');
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

      addLocationMarkers(waters, type, icon) {
        const CHEMS = {
          'ARSENIC': 'Ar',
          'CHROMIUM': 'Cr',
          'LEAD': 'Pb'
        };
        let uniqueWaters = {};
        let curMaxVals = {
          'ARSENIC': 0,
          'CHROMIUM': 0,
          'LEAD': 0
        };
        let deltaMaxVals = {
          'ARSENIC': 0,
          'CHROMIUM': 0,
          'LEAD': 0
        };
        waters.forEach( water => {
          if (water.data[water.data.length - 1] > curMaxVals[water.chemical]) {
            curMaxVals[water.chemical] = water.data[water.data.length - 1];
          }
          if (water.data[water.data.length - 1] - water.data[water.data.length - 2] > deltaMaxVals[water.chemical]) {
            deltaMaxVals[water.chemical] = water.data[water.data.length - 1] - water.data[water.data.length - 2];
          }
          if (water.id in uniqueWaters) {
            uniqueWaters[water.id].push(water);
          } else {
            uniqueWaters[water.id] = [water]
          }
        });
        Object.keys(uniqueWaters).forEach((k) => {
          let label = '';
          uniqueWaters[k].forEach((chem) => {
            let weight = chem.data[chem.data.length - 1] / (curMaxVals[chem.chemical]);
            weight = weight < 0 ? 0 : weight;
            weight = weight * 900;

            let size = (chem.data[chem.data.length - 1] - chem.data[chem.data.length - 2]) / (deltaMaxVals[chem.chemical]);
            size = size < 0 ? 0 : size;
            size = size * 8 + 12;
            label += '<span style="font-weight: ' + weight + '; font-size: '+size+'px" class="location-marker-text">' + CHEMS[chem.chemical] + '</span>';
          });
          let divIcon = L.divIcon({
            className: type+'-marker-div-icon',
            //html:'<i class="fa fa-tint fa-2x"></i><span class="location-marker-text">' + uniqueWaters[k][0].name + '</span>',
            html:'<i class="fa fa-'+icon+' fa-2x"></i> ' + label,
            iconAnchor:[14,14],
            iconSize:null,
            popupAnchor:[0,0]
          });

          let marker = L.marker([uniqueWaters[k][0].lat, uniqueWaters[k][0].lon], {
            icon: divIcon
          }).bindPopup(uniqueWaters[k][0].name);

          this.waterMarkers.addLayer(marker);

        });
        this.waterMarkers.addTo(this.map);


      },

      /* User click a water location spark line */
      locationChanged() {
        this.cluster.clearLayers();
        this.waterMarkers.clearLayers();

        const loc = this.currentLocation;
        const origin = new L.LatLng(loc.lat, loc.lon);

        this.map.setView(origin, 10);  // zoom of 8 gives us about 60 km radius

        this.addLocationMarkers([loc], 'water', 'tint');

        // TODO
        const weightFn = (f) => {
          return 1 + Math.log2(_.last(f.data));
        }

        // Show all close by facilities but only draw lines if the chemical match
        this.filterdFacilities.forEach( facility => {
          const fLoc = new L.LatLng(facility.lat, facility.lon);

          if (facility.chemical == this.currentLocation.chemical) {
            const pointList = [origin, fLoc];
            const line = new L.polyline(pointList, {
              color: '#555',
              weight: 2, //weightFn(facility),
              opacity: 0.5,
              smoothFactor: 1
            });
            this.cluster.addLayer(line);

            this.addLocationMarkers([facility], 'facility', 'industry');
          }
        })
        this.cluster.addTo(this.map);
      }
    }
  }
</script>

<style>
  .map-panel {
    position: relative;
    display: block;
    box-sizing: border-box;
    border: 1px solid #ccc;
    margin: 2px;

    min-height: 600px;
    width: 100%;
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
    //padding-left: 5px;
    color: #555;
    font-size: 14px;
    font-weight: 600;
  }

</style>
