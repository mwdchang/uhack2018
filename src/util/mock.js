export default class {
  static mockW() {
    return [
      {
        'lat': 43.69,
        'lon': -79.59,
        'id': 1,
        'name': 'One',
        'data': [1, 3, 3, 4, 5, 6, 7, 8, 2, 9],
        'type': 'water',
        'chemical': 'lead'
      },
      {
        'lat': 43.69,
        'lon': -79.59,
        'id': 2,
        'name': 'Two',
        'data': [1, 3, 1, 4, 5, 6, 5, 8, 9, 9],
        'type': 'water',
        'chemical': 'arsenic'
      },
    ];

  }


  // facilaties
  static mockF() {
    return [
      {
        'lat': 44.09,
        'lon': -80.59,
        'id': 3,
        'name': 'One',
        'data': [2, 8, 9, 4, 5, 6, 7, 8, 2, 9],
        'type': 'facility',
        'chemical': 'lead'
      },
      {
        'lat': 44.12,
        'lon': -78.59,
        'id': 4,
        'name': 'Two',
        'data': [3, 3, 8, 4, 5, 6, 5, 3, 3, 2],
        'type': 'facility',
        'chemical': 'arsenic'
      },
    ];

  }

}


