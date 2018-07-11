<template>
  <div ref="container" class="spark-container">
    <div class="chemical-label">{{data.chemical}}</div>
    <svg class="chart"></svg>
    <div class="name-label">{{data.name}}</div>
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
    const data = this.data.data;

    const yScale = d3.scaleLinear().range([H, 0]).domain([0, 10]);
    const xScale = d3.scaleLinear().range([W, 0]).domain([0, 10]);

    const valueFn= d3.line()
      .x(function(d, i) { return xScale(i); })
      .y(function(d) { return yScale(d); });


    const type = this.data.type;
    if (type === 'water') {
      chart.append("path")
        .data([data])
        .attr("class", "line-water")
        .attr("d", valueFn);
    } else {
      chart.append("path")
        .data([data])
        .attr("class", "line-facility")
        .attr("d", valueFn);
    }

  }

}
</script>

<style>
.spark-container {
  display: flex;
  flex-direction: row;
}

.line-facility {
  fill: none;
  stroke: red;
  stroke-width: 2px;
}

.line-water {
  fill: none;
  stroke: steelblue;
  stroke-width: 2px;
}

.chemical-label {
  width: 80px;
}

.name-label {
  width: 120px;
}
</style>
