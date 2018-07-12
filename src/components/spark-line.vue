<template>
  <div ref="container" class="spark-container">
    <div v-if="summary === false" class="name-label">{{data.name}}</div>
    <div v-if="summary === false" class="chemical-label">{{data.chemical}}</div>
    <svg class="chart"></svg>
    <div class="delta-label">
      <div :class="{'positive': data.delta > 0, 'negative': data.delta <=0 }">{{deltaStr}}</div>
    </div>
    <div v-if="summary === false && currentLocation" class="delta-label">
      {{pcc}}
    </div>
  </div>
</template>

<script>
import * as d3 from 'd3';
import {  mapActions, mapGetters } from 'vuex';
import Util from '../util/util';

export default {
  name: 'SparkLine',
  props: {
    data: {
      type: Object,
    },
    summary: {
      type: Boolean,
      default: false
    }
  },
  data: () => ({
    deltaStr: ''
  }),
  computed: {
    ...mapGetters({
      currentLocation: 'currentLocation'
    }),
    pcc: function() {
      const s1 = this.data.data;
      const s2 = this.currentLocation.data;
      const pcc = Util.PCC(s1, s2).toFixed(2);
      if (pcc > 0.6) return '+++';
      if (pcc > 0.2) return '+';
      if (pcc < -0.6) return '---';
      if (pcc < -0.2) return '-';
      return '';
    }
  },
  mounted() {
    const W = 80;
    const H = 30;
    const chart = d3.select(this.$refs.container).select('.chart');
    chart.attr('width', W +'px');
    chart.attr('height', H +'px');

    // FIXME
    const data = this.data.data.map(d => {
      return d < 0 ? 0 : d;
    });

    const ymax = d3.max(data);
    const len = data.length-1;

    const formatter = d3.format(".2s");

    this.deltaStr = formatter(this.data.delta);

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
      .style('fill', '#ccc');


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
  background: #FFF;
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
  width: 110px;
  font-size: 80%;
}

.delta-label {
  text-align: left;
  width: 50px;
  font-size: 80%;
  padding-left: 10px;
}


.name-label {
  width: 180px;
  font-size: 80%;
  text-align: left;
}

.positive {
  color: #a00;
}

.negative {
  color: #0a0;
}
</style>
