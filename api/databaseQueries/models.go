package databaseQueries

// busRoute is a type that is designed to read from the trips_n_stops collection
// from MongoDB following passing the aggregation pipeline defined in the
// busRouteQueries file. The route short name is used as the id feature
// for each busRoute while this structure contains arrays of nested structures.
// The Stops array is made of type BusStop while the Shapes array is made of type
// Shape.
type busRoute struct {
	Id        []byte    `bson:"_id" json:"_id"`
	Direction string    `bson:"direction_id" json:"direction_id"`
	Stops     []BusStop `bson:"stops" json:"stops"`
	Shapes    []Shape   `bson:"shapes" json:"shapes"`
}

type RouteId struct {
	RouteNum  string `bson:"route_num" json:"route_num"`
	Direction string `bson:"direction" json:"direction"`
}

// busRouteV2 is a data model that is very similar to the busRoute model used to
// read in the data found in the Mongo collection but it also contains the origin
// and destination stop numbers as separate fields. Given the query structure for
// querying by stops within an array, these fields are necessary to later have the
// origin and destination stops for filtering out the correct length of the trip
// for several elements of the route finding functionality later
type busRouteV2 struct {
	Id                    string    `bson:"_id" json:"_id"`
	Direction             string    `bson:"direction_id" json:"direction_id"`
	Stops                 []BusStop `bson:"stops" json:"stops"`
	Shapes                []Shape   `bson:"shapes" json:"shapes"`
	OriginStopNumber      string    `bson:"origin_stop_number" json:"origin_stop_number"`
	DestinationStopNumber string    `bson:"destination_stop_number" json:"destination_stop_number"`
}

// busRouteJSON is designed in a very similar fashion to the busRoute structure.
// The ID field mirrors that of the busRoute struct and the Shapes array is exactly
// the same also. The main difference between these structures is in the Stops array.
// In the busRouteJSON this array is made of type RouteStop which as a key difference
// returns the coordinates of each bus stop as type float as opposed to strings.
type busRouteJSON struct {
	RouteNum   string               `bson:"route_num" json:"route_num"`
	Stops      []RouteStop          `bson:"stops" json:"stops"`
	Shapes     []ShapeJSON          `bson:"shapes" json:"shapes"`
	Fares      busFares             `bson:"fares" json:"fares"`
	TravelTime TravelTimePrediction `bson:"travel_time,omitempty" json:"travel_time,omitempty"`
	Direction  string               `bson:"direction" json:"direction"`
}

// RouteStop represents the stop information contained within the trips_n_stops
// collection in MongoDB. The information contains the StopId that can be used
// to identify each stop uniquely, the name of that stop, the stop number used
// by consumers of the Dublin Bus service, the coordinates of
// the stop that can be used to mark the stop on the map, a sequence
// number that can be used to sort the stops to ensure that they are in the
// correct order on a given route and finally arrival and departure times for
// when a bus arrived and departed that particular stop for a given trip.
// All fields are returned as strings from the database, apart from the
// coordinates of the stop that come back as a float each
type RouteStop struct {
	StopId            string  `bson:"stop_id" json:"stop_id"`
	StopName          string  `bson:"stop_name" json:"stop_name"`
	StopNumber        string  `bson:"stop_number" json:"stop_number"`
	StopLat           float64 `bson:"stop_lat" json:"stop_lat"`
	StopLon           float64 `bson:"stop_lon" json:"stop_lon"`
	StopSequence      string  `bson:"stop_sequence" json:"stop_sequence"`
	ArrivalTime       string  `bson:"arrival_time" json:"arrival_time"`
	DepartureTime     string  `bson:"departure_time" json:"departure_time"`
	DistanceTravelled float64 `bson:"shape_dist_traveled" json:"shape_dist_traveled"`
}

// Shape is struct that contains the coordinates for each turn in a bus
// line as it travels its designated route that combined together allow
// the bus route to be drawn on a map matching the road network of Dublin.
// All fields map to type string from the database
type Shape struct {
	ShapePtLat      string `bson:"shape_pt_lat" json:"shape_pt_lat"`
	ShapePtLon      string `bson:"shape_pt_lon" json:"shape_pt_lon"`
	ShapePtSequence string `bson:"shape_pt_sequence" json:"shape_pt_sequence"`
	ShapeDistTravel string `bson:"shape_dist_traveled" json:"shape_dist_traveled"`
}

// ShapeJSON is similar in design to the Shape model in but this new
// ShapeJSON model converts the latitudes and longitudes of each point
// into floats to be passed back to the front-end in this format
type ShapeJSON struct {
	ShapePtLat      float64 `bson:"shape_pt_lat" json:"shape_pt_lat"`
	ShapePtLon      float64 `bson:"shape_pt_lon" json:"shape_pt_lon"`
	ShapePtSequence string  `bson:"shape_pt_sequence" json:"shape_pt_sequence"`
	ShapeDistTravel string  `bson:"shape_dist_travel" json:"shape_dist_travel"`
}

// BusStop contains all the necessary information from the mongo collection
// trips_n_stops to provide information on each stop on a given bus route
// for a certain trip, including the stop id, its name and number, its
// coordinates, its number in the sequence of bus stops on the trip and the
// arrival and departure times for the bus making the given trip. All fields
// are returned as strings from the database
type BusStop struct {
	StopId            string `bson:"stop_id,omitempty" json:"stop_id,omitempty"`
	StopName          string `bson:"stop_name" json:"stop_name"`
	StopNumber        string `bson:"stop_number" json:"stop_number"`
	StopLat           string `bson:"stop_lat" json:"stop_lat"`
	StopLon           string `bson:"stop_lon" json:"stop_lon"`
	StopSequence      string `bson:"stop_sequence" json:"stop_sequence"`
	ArrivalTime       string `bson:"arrival_time" json:"arrival_time"`
	DepartureTime     string `bson:"departure_time" json:"departure_time"`
	DistanceTravelled string `bson:"shape_dist_traveled" json:"shape_dist_traveled"`
}

// StopWithCoordinates contains the fields necessary to map out a route
// on a map by including identifying information for each stop (its id,
// name and number) as well as the coordinates for that stop as floating
// point numbers
type StopWithCoordinates struct {
	StopID     string  `bson:"stop_id,omitempty" json:"stop_id,omitempty"`
	StopName   string  `bson:"stop_name" json:"stop_name"`
	StopNumber string  `bson:"stop_number" json:"stop_number"`
	StopLat    float64 `bson:"stop_lat" json:"stop_lat"`
	StopLon    float64 `bson:"stop_lon" json:"stop_lon"`
}

// findByAddressResponse is a simple structure that just contains two arrays
// of nested structures - the StopWithCoordinates structure. It separates the
// response into two different arrays - an array of stops that were sourced
// by matching the search keyword in our database and then an array of
// stops that were sourced by finding the coordinates of the search keyword
// using Google Maps' geocoding service and then finding stops with nearby
// coordinates in our database
type findByAddressResponse struct {
	Matched []StopWithCoordinates `bson:"matched" json:"matched"`
	Nearby  []StopWithCoordinates `bson:"nearby" json:"nearby"`
}

// busFares is an object that is used to map the fares for a respective
// route for each given demographic. These fare values are all floating
// point numbers and are calculated using the CalculateFare function
type busFares struct {
	AdultLeap   float64 `bson:"adult_leap" json:"adult_leap"`
	AdultCash   float64 `bson:"adult_cash" json:"adult_cash"`
	StudentLeap float64 `bson:"student_leap" json:"student_leap"`
	ChildLeap   float64 `bson:"child_leap" json:"child_leap"`
	ChildCash   float64 `bson:"child_cash" json:"child_cash"`
}

// TravelTimePredictionString contains the three different travel time
// fields (the initial prediction time for the whole route, that time plus
// the mean average error and that whole route time minus the mean average error)
// as strings in the format that they are returned from the flask application
// handling unpickling the pickle files containing the prediction models
type TravelTimePredictionString struct {
	TransitTime         string `bson:"transit_time" json:"transit_time"`
	TransitTimePlusMAE  string `bson:"transit_time_plus_mae" json:"transit_time_plus_mae"`
	TransitTimeMinusMAE string `bson:"transit_time_minus_mae" json:"transit_time_minus_mae"`
}

// TravelTimePredictionFloat contains the exact same three fields as TravelTimePredictionString
// but these fields are all converted into floating point numbers to facilitate additional
// calculcations in the back-end based on these values
type TravelTimePredictionFloat struct {
	TransitTime         float64 `bson:"transit_time" json:"transit_time"`
	TransitTimePlusMAE  float64 `bson:"transit_time_plus_mae" json:"transit_time_plus_mae"`
	TransitTimeMinusMAE float64 `bson:"transit_time_minus_mae" json:"transit_time_minus_mae"`
}

// TravelTimePrediction is the data model holding all the necessary information for the
// travel time prediction for a route object. It contains a source field determining if
// the prediction was generated statically from the timetable or dynamically using predictive
// models; three fields for the travel time as per TravelTimePredictionFloat but these fields
// are now integers rounded to the nearest minute; three fields with the travel time added to
// the static scheduled departure from the first stop to find the estimated arrival time for the
// given destination; and the scheduled departure time from the static timetable for the
// given origin
type TravelTimePrediction struct {
	Source                   string `bson:"source" json:"source"`
	TransitTime              int    `bson:"transit_time" json:"transit_time"`
	TransitTimePlusMAE       int    `bson:"transit_time_plus_mae" json:"transit_time_plus_mae"`
	TransitTimeMinusMAE      int    `bson:"transit_time_minus_mae" json:"transit_time_minus_mae"`
	EstimatedArrivalTime     string `bson:"estimated_arrival_time" json:"estimated_arrival_time"`
	EstimatedArrivalHighTime string `bson:"estimated_arrival_high_time" json:"estimated_arrival_high_time"`
	EstimatedArrivalLowTime  string `bson:"estimated_arrival_low_time" json:"estimated_arrival_low_time"`
	ScheduledDepartureTime   string `bson:"scheduled_departure_time" json:"scheduled_departure_time"`
}

// RouteByStop contains the id of a given route as its route number and the slice
// of stops along that route, each of which is type BusStop
type RouteByStop struct {
	Id    string    `bson:"_id" json:"_id"`
	Stops []BusStop `bson:"stops" json:"stops"`
}

// MatchedRoute is a data model that contains the origin stop number,
// the destination stop number and the route number for a route that has been
// matched in the database all as strings
type MatchedRoute struct {
	Id    []string  `bson:"_id" json:"_id"`
	Stops []BusStop `bson:"stops" json:"stops"`
}

type MatchedRouteWithOAndD struct {
	Id                    []string  `bson:"_id" json:"_id"`
	Stops                 []BusStop `bson:"stops" json:"stops"`
	OriginStopNumber      string    `bson:"origin_stop_number" json:"origin_stop_number"`
	DestinationStopNumber string    `bson:"destination_stop_number" json:"destination_stop_number"`
}

// GeolocatedStop is a data model that contains the stop information for a given
// bus stop in string format (i.e. its stop id, stop name, stop number and coordinates
// on a map) as well as the internal Mongo id for the data entry
type GeolocatedStop struct {
	ID         string `bson:"_id,omitempty" json:"_id,omitempty"`
	StopId     string `bson:"stop_id" json:"stop_id"`
	StopName   string `bson:"stop_name" json:"stop_name"`
	StopNumber string `bson:"stop_number" json:"stop_number"`
	StopLat    string `bson:"stop_lat" json:"stop_lat"`
	StopLon    string `bson:"stop_lon" json:"stop_lon"`
}
