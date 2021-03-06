[![Build Status][drone-image]][drone-url]
[![Coverage Status][coveralls-image]][coveralls-url]

**This project is discontinued. My interest in EVE has dropped again and based on experience, it'll take some years until I might resub. Furthermore interest in this library was low, which is why I keep it as a project for experience.**

# everoute - Go library

This Go library provides a route calculator and optimizer for ship travel in [EVE Online](https://www.eveonline.com/).
The calculator can use and combine travel capabilities of any kind (jump gates, jump drive, wormholes, ...) and restricts the path finding by rules (length, security ratings, jump distances, ...). The library can find the best route between two systems and also optimize the order of waypoints of a larger route. The waypoint optimizer is using a genetic algorithm to provide decent routes within short time.

An interface based design allows extensions to be made regarding rules and search criteria; For example, a rule could be written, combining live API data, to avoid systems with certain SOV.

## Use
This library can be used directly by integrating it into another Go process - or by using its functionality through a web-service. The web-service is implemented in a dedicated project: [everoute-web](https://github.com/dertseha/everoute-web).

## Origin
This library is based on the same design as [eve-route.js](https://github.com/dertseha/eve-route.js) and has since superseded the JavaScript library. This Go implementation also surpasses performance in terms of memory usage due to strict and native types, as well as performance due to the parallelization capability.

## License

The project is available under the terms of the **New BSD License** (see LICENSE file).

[drone-url]: https://drone.io/github.com/dertseha/everoute/latest
[drone-image]: https://drone.io/github.com/dertseha/everoute/status.png
[coveralls-url]: https://coveralls.io/r/dertseha/everoute
[coveralls-image]: https://coveralls.io/repos/dertseha/everoute/badge.png
