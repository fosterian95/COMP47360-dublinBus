import 'package:dublin_bus_diy/models/app_model.dart';
import 'package:flutter/material.dart';
import 'package:geolocator/geolocator.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:provider/provider.dart';

import '../api/location_service.dart';
import '../models/map_polylines.dart';

// Adding Google Maps to a Flutter app
// https://codelabs.developers.google.com/codelabs/google-maps-in-flutter#0
class GoogleMapComponent extends StatelessWidget {
  const GoogleMapComponent({Key? key}) : super(key: key);

  // The initial point that will be centered in the map
  final LatLng _center = const LatLng(53.34571963981868, -6.264174663517609);

  @override
  Widget build(BuildContext context) {
    print('build Polyline');
    // The Consumer widget make it possible for us to use the Model from the state conveniently
    // https://docs.flutter.dev/development/data-and-backend/state-mgmt/simple#consumer
    return Consumer<PolylinesModel>(
      builder: (context, polylinesModel, child) => Stack(
        children: [
          // Use SomeExpensiveWidget here, without rebuilding every time.
          if (child != null) child,
          GoogleMap(
            onMapCreated: (controller) {
              Provider.of<AppModel>(context, listen: false).mapController = controller;
            },
            initialCameraPosition: CameraPosition(
              target: _center,
              zoom: 11.0,
            ),
            polylines: Set<Polyline>.of(polylinesModel.itemsOfPolylines),
            markers: Set<Marker>.of(polylinesModel.itemsOfMarkers),
            myLocationEnabled: true,
          ),
        ],
      ),
      // Build the expensive widget here.
      // child: const SomeExpensiveWidget(),
    );
  }

  // animate camera refer to
  // https://github.com/flutter/plugins/blob/main/packages/google_maps_flutter/google_maps_flutter/example/lib/animate_camera.dart
  // https://levelup.gitconnected.com/how-to-add-google-maps-in-a-flutter-app-and-get-the-current-location-of-the-user-dynamically-2172f0be53f6
  updateMapLocation(GoogleMapController mapController, PolylinesModel polylinesModel) async {
    print('updating map location');
    Position here = await determinePosition();
    mapController.animateCamera(CameraUpdate.newCameraPosition(
        CameraPosition(target: LatLng(here.latitude, here.longitude), zoom: 14)));
    polylinesModel.showSingleMarkerPosition(here.latitude, here.longitude);
  }
}
