<template>
  <div>
    <header-bar/>
    <div class="container">
      <side-panel/>
      <map-panel/>
    </div>
  </div>
</template>

<script>

import SidePanel from './side-panel.vue';
import HeaderBar from './header-bar.vue';
import MapPanel from './map-panel.vue';

/* template */
export default {
  name: 'home',
  components: {
    SidePanel, HeaderBar, MapPanel
  },
  data: () => ({
    label: 'hello world'
  }),
  mounted() {
  },
  watch: {
  },
  method: {
  }
}
</script>

<style>

.container {
  display: flex;
  flex-direction: row;
  height: 100%;
  width: 100%;
}

</style>
