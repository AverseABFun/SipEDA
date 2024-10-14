//SipEDA 0.1 PCB file
version 0.1;
type pcb;

add_sch example.sch.sip;    // add a schematic to check against
add_fp_lib example.fop.sip; // add a footprint library

footprint SMD:0603; // add a footprint
ref 1;              // set a reference number(connects to schematic)
x 0;                // set the x...
y 0;                // and the y
rotate 90;          // rotate 90 degrees

footprint TP:1.0mm;
ref 2;
x 3;
y 0;

footprint TP:1.0mm;
ref 3;
x 3;
y 3;

p 1.0; // get x and y from a pad
       // you can also use x and y commands to set position
       // then use "point" to create the point
p 2.0;
// add more points to create lines between them
track; // create a track

p 1.1;
p 3.0;
track;

layer bol; // set the layer to the board outline layer

// see the footprint file for info on these commands:
draw;
    x -5;
    y -5;
    w 20;
    h 10;
    rect;
enddraw;

