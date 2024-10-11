//SipEDA 0.1 footprint file
version 0.1;
type footprint;

footprint; // start a new footprint, have multiple "footprint;...endfootprint;" blocks
           // for multiple footprints in one file

// if the component is set to the other side, all layers will
// be automatically flipped(9 minus layer #)
lib "SMD";     // library of the footprint
name "0603";   // name of the footprint
pads 2;        // number of pads
pads smd;      // smd or tht
units mm;      // use millimeters(supports mm, mil, and in)
font openSans; // use the font open sans for text

layer 9cl; // top copper layer(0cl is bottom, for more than two-layer pcbs it's numbered 1-8)

pad 0;    // pad zero
    draw;     // begin draw
        x -1.275; // x pos
        y -0.5;   // y pos
        w 0.95;   // width
        h 1;      // height
        rect;     // create a rect from (x,y) to (x+w,y+h)
    enddraw;  // end draw

pad 1;
    draw;
        x 0.325;
        y 0.5;
        w 0.95;
        h 1;
        rect;
    enddraw;

deselect all; // deselect everything

layer 9ss; // top silkscreen(only 0ss and 9ss, named for consistency)

draw;

    x 0;
    y 1.15;
    align center;
    text "$ref";

    x 0;
    y -1.15;
    align center;
    text "$comp";

    x 0;
    y -2;
    align center;
    text "$desc";

enddraw;

layer 9mg; // top margin

draw;
    x -2;
    y -1.05;
    w 4;
    h 2.1;
    rect;
enddraw;

    from 9mg; // copy from
    to 9ss;   // copy to
    to 9fb;   // 9fb is the top fabrication layer
copy;     // actually copy

endfootprint;

footprint;

lib "TP";
name "1.0mm";
pads 1;
pads smd;
units mm;
font openSans;

layer 9cl;

pad 0;
    draw;
        x 0;
        y 0;
        d 1; // diameter
        circle; // draw a circle
    enddraw;

deselect all;

layer 9ss;

draw;

    x 0;
    y 2;
    align center;
    text "$ref";

    x 0;
    y -2;
    align center;
    text "$desc";

enddraw;

layer 9mg;

draw;
    x 0;
    y 0;
    d 1.5;
    circle;
enddraw;

    from 9mg;
    to 9ss;
    to 9fb;
copy;

endfootprint;