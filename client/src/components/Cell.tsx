import React from 'react';
import { XY } from '../models/XY'
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';

export default function Test(
    { id, data, x, y, width, height, selected, setSelected, mousePosition }:
        {
            id: string,
            data: any,
            x: number,
            y: number,
            width: number,
            height: number,
            selected: boolean,
            setSelected: any,
            mousePosition: XY,
        }) {

    const [archiveopen, setArchiveopen] = React.useState(false);
    const [isMoved, setIsMoved] = React.useState(false);
    const [isResized, setIsResized] = React.useState(false);
    const [cX, setCX] = React.useState(x);
    const [cY, setCY] = React.useState(y);
    const [cW, setCW] = React.useState(width);
    const [cH, setCH] = React.useState(height);

    React.useEffect(() => {
        if (isMoved) {
            setCX(cX + mousePosition.movX);
            setCY(cY + mousePosition.movY)
        }
        if (isResized) {
            setCW(cW + mousePosition.movX);
            setCH(cH + mousePosition.movY);
        }
    }, [mousePosition]);

    const handleLinks = (data: string): string => {

        const youtubetmpl = `<iframe width="350" height="200" src="https://www.youtube.com/embed/$2" title="YouTube video player" ></iframe>`;
        //data = data.replace(/(https:\/\/www\.youtube\.com\/embed\/(\S+))/, youtubetmpl);
        data = data.replace(/(https:\/\/www\.youtube\.com\/watch\?v=(\S{11}))/, youtubetmpl);
        data = data.replace(/((http(s?):)([/|.|\w|\s|-])*\.(?:jpg|gif|png))/, '<img width="300" src="$1"/>');

        return data
    };

    const handleAdd = () => {
        console.log("add add add");
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
                width={cW}
                height={cH}
                x={cX}
                y={cY}
                fill="none"
                stroke={selected ? "tomato" : "cadetblue"}
                stroke-width={2}
            />
            <foreignObject
                x={cX + 10}
                y={cY + 10}
                width={cW - 20}
                height={cH - 20}
                onClick={(event)=>{
                    event.stopPropagation();
                    setSelected(id);
                }}
                onDoubleClick={()=>alert("lsdjfldsfj")}
                onMouseDown={(event) => {
                    event.stopPropagation();
                    setIsMoved(true)
                }}
                onMouseUp={() => {
                    setIsMoved(false)
                }}
            >
                <div
                    style={{ color: "pink", whiteSpace: "pre-wrap" }}
                    dangerouslySetInnerHTML={{ __html: handleLinks(data) }}
                />
            </foreignObject>
            <g display={selected ? "inherit" : "none"} >
                <rect
                    rx="3"
                    width={20}
                    height={20}
                    x={cX - 24}
                    y={cY - 24}
                    fill="#282c34"
                    stroke={selected ? "tomato" : "cadetblue"}
                    stroke-width={2}
                />
                <line x1={cX - 24 + 10} y1={cY - 24 + 5} x2={cX - 24 + 10} y2={cY - 4 - 5} stroke="pink" stroke-width="1" />
                <line x1={cX - 24 + 5} y1={cY - 24 + 10} x2={cX - 24 + 20 - 5} y2={cY - 24 + 10} stroke="pink" stroke-width="1" />
            </g>
            <g display={selected ? "inherit" : "none"}>
                <rect
                    rx="3"
                    width={20}
                    height={20}
                    x={cX + cW + 4}
                    y={cY - 24}
                    fill="#282c34"
                    stroke={selected ? "tomato" : "cadetblue"}
                    stroke-width={2}
                    onClick={handleArchiveStart}
                />
                <line x1={cX + cW + 4 + 5} y1={cY - 24 + 5} x2={cX + cW + 4 + 20 - 5} y2={cY - 24 + 20 - 5} stroke="pink" stroke-width="1" />
                <line x1={cX + cW + 4 + 5} y1={cY - 24 + 20 - 5} x2={cX + cW + 4 + 20 - 5} y2={cY - 24 + 5} stroke="pink" stroke-width="1" />
            </g>
            <g display={selected ? "inherit" : "none"}>
                <rect
                    rx="3"
                    width={20}
                    height={20}
                    x={cX + cW + 4}
                    y={cY + cH + 4}
                    fill="#282c34"
                    stroke={selected ? "tomato" : "cadetblue"}
                    stroke-width={2}
                    onMouseDown={(event) => {
                        event.stopPropagation();
                        setIsResized(true)
                    }}
                    onMouseUp={() => {
                        setIsResized(false)
                    }}
                />
                <line x1={cX + cW + 4 + 5} y1={cY + cH + 4 + 20 - 5} x2={cX + cW + 4 + 20 - 5} y2={cY + cH + 4 + 20 - 5} stroke="pink" stroke-width="1" />
                <line x1={cX + cW + 4 + 20 - 5} y1={cY + cH + 4 + 20 - 5} x2={cX + cW + 4 + 20 - 5} y2={cY + cH + 4 + 5} stroke="pink" stroke-width="1" />
            </g>
        </g>
    );

}
