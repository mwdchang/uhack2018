<template>
  <div class="side-panel">
    <div>Side panel stub</div>
    <div>
      <spark-line/>
      <spark-line/>
      <spark-line/>
    </div>
  </div>
</template>


<script>
import SparkLine from './spark-line.vue';

export default {
  name: 'side-panel',
  components: {
    SparkLine
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
