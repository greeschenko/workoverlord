import React from 'react';

export default function Synapse(
    { x1, y1, x2, y2, size, color }: {
        x1: number,
        y1: number,
        x2: number,
        y2: number,
        size: number,
        color: string
    }) {
    const [datachange] = React.useState(0);
    const [divTop, setDivTop] = React.useState(0);
    const [divLeft, setDivLeft] = React.useState(0);
    const [width, setWidth] = React.useState(0);
    const [height, setHeight] = React.useState(0);
    const [mX1, setMX1] = React.useState(0);
    const [mX2, setMX2] = React.useState(0);
    const [mY1, setMY1] = React.useState(0);
    const [mY2, setMY2] = React.useState(0);

    React.useEffect(() => {
        if (x1 < x2 && y1 < y2) {
            setDivTop(y1);
            setDivLeft(x1);
            setWidth(x2-x1+size);
            setHeight(y2-y1+size);
            setMX1(size/2);
            setMY1(size/2);
            setMX2(x2-x1+size);
            setMY2(y2-y1+size);
        } else if (x1 < x2 && y1 > y2) {
            setDivTop(y2);
            setDivLeft(x1);
            setWidth(x2-x1+size);
            setHeight(y1-y2+size);
            setMX1(size/2);
            setMY1(y1-y2+size);
            setMX2(x2-x1+size);
            setMY2(size/2);
        } else if (x1 > x2 && y1 > y2) {
            setDivTop(y2);
            setDivLeft(x2);
            setWidth(x1-x2+size);
            setHeight(y1-y2+size);
            setMX1(x1-x2+size);
            setMY1(y1-y2+size);
            setMX2(size/2);
            setMY2(size/2);
        } else if (x1 > x2 && y1 < y2) {
            setDivTop(y1);
            setDivLeft(x2);
            setWidth(x1-x2+size);
            setHeight(y2-y1+size);
            setMX1(size/2);
            setMY1(y2-y1+size);
            setMX2(x1-x2+size);
            setMY2(size/2);
        } else if (x1 === x2 && y1 < y2) {
            setDivTop(y1);
            setDivLeft(x1);
            setWidth(size);
            setHeight(y2-y1+size);
            setMX1(size/2);
            setMY1(size/2);
            setMX2(size/2);
            setMY2(y2-y1+size);
        } else if (x1 === x2 && y1 > y2) {
            setDivTop(y2);
            setDivLeft(x2);
            setWidth(size);
            setHeight(y1-y2+size);
            setMX1(size/2);
            setMY1(y1-y2+size);
            setMX2(size/2);
            setMY2(size/2);
        } else if (x1 < x2 && y1 === y2) {
            setDivTop(y1);
            setDivLeft(x1);
            setWidth(x2-x1+size);
            setHeight(15);
            setMX1(size/2);
            setMY1(size/2);
            setMX2(x2-x1+size);
            setMY2(size/2);
        } else if (x1 > x2 && y1 === y2) {
            setDivTop(y2);
            setDivLeft(x2);
            setWidth(x1-x2+size);
            setHeight(15);
            setMX1(x1-x2+size);
            setMY1(size/2);
            setMX2(size/2);
            setMY2(size/2);
        }
    }, [datachange]);

    return (
        <div style={{ position: 'absolute', top: divTop, left: divLeft }}>
            <svg width={width} height={height} xmlns="http://www.w3.org/2000/svg">
                <line x1={mX1} y1={mY1} x2={mX2} y2={mY2} stroke={color}
                    stroke-width={size}
                />
            </svg>
        </div>
    );
}
