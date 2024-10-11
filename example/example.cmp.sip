//SipEDA 0.1 component file
version 0.1;
type component;

add_fp_lib example.fop.sip; // add a footprint library

component; // start a new component, have multiple component;...endcomponent; blocks
           // for multiple components in one file

lib Common;    // library of the component
name resistor; // name of the component
pins 2;          // number of pins
ref R#;        // reference(# is replaced with the number)

footprint SMD:0603;         // repeat these sort of commands to allow multiple footprints
default_footprint SMD:0603; // default footprint

pin 0;        // pin zero
type passive; // pin type(passive, input, output, bidirectional, pwrinput, pwroutput)
name 0;     // we don't want a name for resistor pins

pin 1;
type passive;
name 1;

endcomponent;

component;

lib Common;
name test_point;
pins 1;
ref TP#;

footprint TP:1.0mm;
default_footprint TP:1.0mm;

pin 0;
type passive;
name 0;

endcomponent;
