export default class {
  static getWater() {
    return fetch('http://localhost:3333/watera');
  }

  static facility() {
    return fetch('http://localhost:3333/facilities');
  }
}
