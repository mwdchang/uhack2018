<template>
  <div ref="container" class="spark-container">
    <div>
      <div class="chemical-label">{{data.chemical}}</div>
      <svg class="chart"></svg>
    </div>
    <div class="name-label" style="display: flex; flex-direction:column">
      <div>{{data.name}}</div>
      <div class="name-label" :class="{'positive': data.delta > 0, 'negative': data.delta <=0 }">{{deltaStr}}</div>
    </div>

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
  data: () => ({
    deltaStr: ''
  }),
  mounted() {
    const W = 200;
    const H = 50;
    const chart = d3.select(this.$refs.container).select('.chart');
    chart.attr('width', W +'px');
    chart.attr('height', H +'px');

    // FIXME
    const data = this.data.data.map(d => {
      return d < 0 ? 0 : d;
    });

    const ymax = d3.max(data);
    const len = data.length-1;


    this.deltaStr = this.data.delta.toFixed(2);

    const yScale = d3.scaleLinear().range([H, 0]).domain([-1, ymax+1]);
    const xScale = d3.scaleLinear().range([0, W]).domain([0, len]);

    const valueFn= d3.line()
      .x(function(d, i) { return xScale(i); })
      .y(function(d) { return yScale(d); });


    chart.append('rect')
      .attr('x', 0)
      .attr('y', yScale(data[0]))
      .attr('width', W)
      .attr('height', 1)
      .style('stroke', 'none')
      .style('fill', '#f80');


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
  padding: 2px;
  border: 1px solid #ccc;
  box-sizing: border-box;
  cursor: pointer;
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
  text-align: left;
  width: 120px;
  font-size: 80%;
}

.name-label {
  width: 120px;
  font-size: 80%;
}

.positive {
  color: #a00;
}

.negative {
  color: #0a0;
}
</style>
