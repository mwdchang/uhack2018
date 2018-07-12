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
      this.addLocationMarkers();
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
          this.map.setView([43.6532, -79.3832], 9);
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
        const CHEMS = {
          'ARSENIC': 'Ar',
          'CHROMIUM': 'Cr',
          'LEAD': 'Pb'
        };
        let uniqueWaters = {};
        this.waters.forEach( water => {
          if (water.id in uniqueWaters) {
            uniqueWaters[water.id].push(water);
          } else {
            uniqueWaters[water.id] = [water]
          }
        });
        Object.keys(uniqueWaters).forEach((k) => {
          let label = ''
          uniqueWaters[k].forEach((chem) => {
            label += CHEMS[chem.chemical];
          });
          let divIcon = L.divIcon({
            className:'water-marker-div-icon',
            //html:'<i class="fa fa-tint fa-2x"></i><span class="location-marker-text">' + uniqueWaters[k][0].name + '</span>',
            html:'<i class="fa fa-tint fa-2x"></i><span class="location-marker-text">' + label + '</span>',
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

        const loc = this.currentLocation;
        const origin = new L.LatLng(loc.lat, loc.lon);


        this.map.setView(origin, 8);  // zoom of 8 gives us about 60 km radius
        const max = d3.max(this.filterdFacilities.map( f => _.last(f.data)));

        // TODO
        const weightFn = (f) => {
          return 1 + Math.log2(_.last(f.data));
        }

        this.filterdFacilities.forEach( facility => {
          const fLoc = new L.LatLng(facility.lat, facility.lon);
          const pointList = [origin, fLoc];
          const line = new L.polyline(pointList, {
            color: '#F0F',
            weight: weightFn(facility),
            opacity: 0.5,
            smoothFactor: 1
          });
          this.cluster.addLayer(line);

          let divIcon =L.divIcon({
            className:'facility-marker-div-icon',
            //html:'<i class="fa fa-industry fa-2x"></i><span class="location-marker-text">' + facility.name + '</span>',
            html:'<i class="fa fa-industry fa-2x"></i>',
            iconAnchor:[14,14],
            iconSize:null,
            popupAnchor:[0,0]
          });

          let marker = L.marker([facility.lat, facility.lon], {
            icon: divIcon,
            //opacity: 0.3,
          });
          this.cluster.addLayer(marker);

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
  padding-left: 5px;
  color: black;
  font-size: 14px;
  font-weight: 900;
}

</style>
