export default class {
  static PCC(x, y) {
    let n = x.length;
    let sum_xy = 0;
    let sum_x2 = 0;
    let sum_y2 = 0;
    let meanx = 0, meany = 0;
    for (let i=0; i < n; i++) {
      sum_xy += x[i]*y[i];
      sum_x2 += x[i]*x[i];
      sum_y2 += y[i]*y[i];
      meanx += x[i];
      meany += y[i];
    }
    meanx /= n;
    meany /= n;
    let n_meanx2 = n * (meanx*meanx);
    let n_meany2 = n * (meany*meany);
    let xdenom = 0;
    let ydenom = 0;
    xdenom = Math.sqrt(sum_x2 - n_meanx2);
    ydenom = Math.sqrt(sum_y2 - n_meany2);
    if (xdenom === 0 || ydenom === 0) return 0;
    return (sum_xy - n * meanx * meany) / (xdenom * ydenom);
  }
}
