//SipEDA 0.1 schematic file
version 0.1;
type schematic;

add_cmp_lib "example.cmp.sip" // add a schematic library

rail GND; // create a rail
rail INOUT;

component "Common:resistor"; // add a component
ref 1;                       // set the reference of the last added component
footprint default;           // set the footprint to be the default

component "Common:test_point";
ref 2;
desc "GND"; // add a description to the component(will be put on the pcb)
footprint default;

component "Common:test_point";
ref 3;
desc "INOUT";
footprint default;

point 1.0; // add one point to connect
point GND; // add another
connect;   // connect the points

point 1.1;
point INOUT;
connect;

point GND;
point 2.0;
connect;

point INOUT;
point 3.0;
connect;