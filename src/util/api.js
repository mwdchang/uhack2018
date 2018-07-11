export default class {
  static getWater() {
    return fetch('http://localhost:3333/water');
  }

  static getFacility() {
    return fetch('http://localhost:3333/facilities');
  }
}
