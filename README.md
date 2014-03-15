## everoute - Go library

This Go library provides a route calculator and optimizer for ship travel in [EVE Online](https://www.eveonline.com/).
The calculator can use and combine travel capabilities of any kind (jump gates, jump drive, wormholes, ...) and restricts the path finding by rules (length, security ratings, jump distances, ...). The waypoint optimizer is using a genetic algorithm to provide decent routes within short time.

An interface based design allows extensions to be made regarding rules and search criteria; For example, a rule could be written, combining live API data, to avoid systems with certain SOV.

This library mirrors the implementation of [eve-route.js](https://github.com/dertseha/eve-route.js) and surpasses performance in terms of memory usage due to strict and native types, as well as performance due to the parallelization capability.

## License

The project is available under the terms of the **New BSD License** (see LICENSE file).
