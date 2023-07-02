import React from 'react';
import { XY } from '../models/XY'
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';

export default function Test(
    { id, data, x, y, width, height, selected, onClick, mousePosition }:
        {
            id: string,
            data: any,
            x: number,
            y: number,
            width: number,
            height: number,
            selected: boolean,
            onClick: (event: any) => void,
            mousePosition: XY,
        }) {

    const [archiveopen, setArchiveopen] = React.useState(false);
    const [isMoved, setIsMoved] = React.useState(false);
    const [isResized, setIsResized] = React.useState(false);
    const [cX, setCX] = React.useState(x);
    const [cY, setCY] = React.useState(y);

    React.useEffect(() => {
        if (isMoved) {
            setCX(mousePosition.x + 24 / 2);
            setCY(mousePosition.y + 24 / 2)
        }
    }, [mousePosition]);

    const handleLinks = (data: string): string => {

        const youtubetmpl = `<iframe width="350" height="200" src="https://www.youtube.com/embed/$2" title="YouTube video player" ></iframe>`;
        //data = data.replace(/(https:\/\/www\.youtube\.com\/embed\/(\S+))/, youtubetmpl);
        data = data.replace(/(https:\/\/www\.youtube\.com\/watch\?v=(\S{11}))/, youtubetmpl);
        data = data.replace(/((http(s?):)([/|.|\w|\s|-])*\.(?:jpg|gif|png))/, '<img width="300" src="$1"/>');

        return data
    };

    const handleMove = () => {
        console.log("move move move");
    }

    const handleAdd = () => {
        console.log("add add add");
    }

    const handleResize = () => {
        console.log("resize resize");
    }

    const handleArchiveStart = (event: any) => {
        event.stopPropagation();
        setArchiveopen(true);
    }

    const handleArchiveClose = () => {
        setArchiveopen(false);
    }

    const handleArchiveSubmit = () => {
        console.log(id);
    }

    return (
        <g style={{ userSelect: selected ? "auto" : "none" }} fill="white" stroke="green" stroke-width="5">
            <Dialog fullWidth={false} open={archiveopen} onClose={handleArchiveClose}>
                <DialogContent>
                    Put to the archive. You are shure?
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleArchiveClose}>Cancel</Button>
                    <Button onClick={handleArchiveSubmit}>Submit</Button>
                </DialogActions>
            </Dialog>
            <rect
                rx="6"
                width={width}
                height={height}
                x={cX}
                y={cY}
                fill="none"
                stroke={selected ? "tomato" : "cadetblue"}
                stroke-width={2}
            />
            <foreignObject
                x={cX + 10}
                y={cY + 10}
                width={width - 20}
                height={height - 20}
                onClick={onClick}
            >
                <div
                    style={{ color: "pink", whiteSpace: "pre-wrap" }}
                    dangerouslySetInnerHTML={{ __html: handleLinks(data) }}
                />
            </foreignObject>
            <rect
                rx="3"
                width={20}
                height={20}
                x={cX - 24}
                y={cY - 24}
                fill="#282c34"
                stroke={selected ? "tomato" : "cadetblue"}
                stroke-width={2}
                display={selected ? "inherit" : "none"}
                onMouseDown={() => setIsMoved(true)}
                onMouseUp={() => setIsMoved(false)}
            />
            <g display={selected ? "inherit" : "none"}>
                <rect
                    rx="3"
                    width={20}
                    height={20}
                    x={cX + width + 4}
                    y={cY - 24}
                    fill="#282c34"
                    stroke={selected ? "tomato" : "cadetblue"}
                    stroke-width={2}
                    onClick={handleArchiveStart}
                />
                <line x1={cX + width + 4} y1={cY - 24} x2={cX + width + 4 + 20} y2={cY - 24 + 20} stroke="pink" stroke-width="1" />
                <line x1={cX + width + 4} y1={cY - 24 + 20} x2={cX + width + 4 + 20} y2={cY - 24} stroke="pink" stroke-width="1" />
            </g>
            <rect
                rx="3"
                width={20}
                height={20}
                x={cX + width + 4}
                y={cY + height + 4}
                fill="#282c34"
                stroke={selected ? "tomato" : "cadetblue"}
                stroke-width={2}
                display={selected ? "inherit" : "none"}
                onClick={handleResize}
            />
        </g>
    );

}
