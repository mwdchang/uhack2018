<template>
  <div ref="container" class="spark-container">
    <div>Name</div>
    <svg class="chart"></svg>
  </div>
</template>

<script>
import * as d3 from 'd3';

export default {
  name: 'SparkLine',
  props: {
    data: {
      type: Object,
    }
  },
  mounted() {
    const W = 200;
    const H = 50;
    const chart = d3.select(this.$refs.container).select('.chart');
    chart.attr('width', W +'px');
    chart.attr('height', H +'px');
    const data = [1, 3, 5, 7, 3, 2, 5, 6, 8, 8];

    const yScale = d3.scaleLinear().range([H, 0]).domain([0, 10]);
    const xScale = d3.scaleLinear().range([W, 0]).domain([0, 10]);

    const valueFn= d3.line()
      .x(function(d, i) { return xScale(i); })
      .y(function(d) { return yScale(d); });

    chart.append("path")
      .data([data])
      .attr("class", "line")
      .attr("d", valueFn);

  }

}
</script>

<style>
.spark-container {
  display: flex;
  flex-direction: row;
}

.line {
  fill: none;
  stroke: steelblue;
  stroke-width: 2px;
}
</style>
