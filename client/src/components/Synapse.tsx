//import React from 'react';

export default function Synapse({x1, y1, x2, y2, color, size}: {
    x1: number,
    y1: number,
    x2: number,
    y2: number,
    color?: string,
    size: number,
}) {

    return (
        <line x1={x1} y1={y1} x2={x2} y2={y2} stroke={color || "cadetblue"} stroke-width={size} />
    );
}
