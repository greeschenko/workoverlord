import React from 'react';
import { XY } from '../models/XY'
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';

export default function Cell(
    { id, data, x, y, width, height, selected, setSelected, mousePosition, scaleIndex }:
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
            scaleIndex: number,
        }) {

    const [archiveopen, setArchiveopen] = React.useState(false);
    const [isMoved, setIsMoved] = React.useState(false);
    const [isResized, setIsResized] = React.useState(false);
    const [cX, setCX] = React.useState(x);
    const [cY, setCY] = React.useState(y);
    const [cW, setCW] = React.useState(width);
    const [cH, setCH] = React.useState(height);

    const btnWstart = 20
    const btnHstart = 20

    const [btnW, setBtnW] = React.useState(btnWstart * scaleIndex);
    const [btnH, setBtnH] = React.useState(btnHstart * scaleIndex);

    {/*
      *React.useEffect(() => {
      *    setBtnW(btnWstart * scaleIndex);
      *    setBtnH(btnHstart * scaleIndex);
      *}, [scaleIndex]);
      */}

    React.useEffect(() => {
        if (isMoved) {
            setCX(cX + mousePosition.movX * scaleIndex);
            setCY(cY + mousePosition.movY * scaleIndex)
        }
        if (isResized) {
            setCW(cW + mousePosition.movX * scaleIndex);
            setCH(cH + mousePosition.movY * scaleIndex);
        }
    }, [mousePosition]);

    React.useEffect(() => {
        if (!selected && isMoved) {
            setIsMoved(false);
        }
        if (!selected && isResized) {
            setIsResized(false);
        }
    }, [selected]);

    const handleLinks = (data: string): string => {

        const youtubetmpl = `<iframe width="350" height="200" src="https://www.youtube.com/embed/$2" title="YouTube video player" ></iframe>`;
        //data = data.replace(/(https:\/\/www\.youtube\.com\/embed\/(\S+))/, youtubetmpl);
        data = data.replace(/(https:\/\/www\.youtube\.com\/watch\?v=(\S{11}))/, youtubetmpl);
        data = data.replace(/((http(s?):)([/|.|\w|\s|-])*\.(?:jpg|gif|png))/, '<img width="300" src="$1"/>');

        return data
    };

    const handleAdd = (event: any) => {
        event.stopPropagation();
        console.log("add add add");
    }

    const handleDone = (event: any) => {
        event.stopPropagation();
        console.log("done done");
    }

    const handleArchive = (event: any) => {
        event.stopPropagation();
        console.log("archive archive");
    }

    const handleDeleteStart = (event: any) => {
        event.stopPropagation();
        setArchiveopen(true);
    }

    const handleDeleteClose = () => {
        setArchiveopen(false);
    }

    const handleDeleteSubmit = () => {
        console.log("permanently delete", id);
    }

    return (
        <g style={{ userSelect: selected ? "auto" : "none" }} fill="white" stroke="green" stroke-width="5">
            <Dialog fullWidth={false} open={archiveopen} onClose={handleDeleteClose}>
                <DialogContent>
                    Delete permanently. You are shure?
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleDeleteClose}>Cancel</Button>
                    <Button variant="outlined" color="error" onClick={handleDeleteSubmit}>DELETE</Button>
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
                onClick={(event) => {
                    event.stopPropagation();
                    setSelected(id);
                }}
                onDoubleClick={() => alert("lsdjfldsfj")}
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
            <g display={selected ? "inherit" : "none"}>
                <text
                    className="svgbtn"
                    fill="tomato"
                    stroke="none"
                    font-size="14"
                    font-family="monospace"
                    x={cX + cW + 4 + 5}
                    y={cY + 8}
                    onClick={handleAdd}
                >
                    ADD SUBCELL
                </text>
                <text
                    className="svgbtn"
                    fill="tomato"
                    stroke="none"
                    font-size="14"
                    font-family="monospace"
                    x={cX + cW + 4 + 5}
                    y={cY + 8 + 16}
                    onClick={handleDone}
                >
                    DONE
                </text>
                <text
                    className="svgbtn"
                    fill="tomato"
                    stroke="none"
                    font-size="14"
                    font-family="monospace"
                    x={cX + cW + 4 + 5}
                    y={cY + 8 + 16 * 2}
                    onClick={handleArchive}
                >
                    ARCHIVE
                </text>
                <text
                    className="svgbtn"
                    fill="tomato"
                    stroke="none"
                    font-size="14"
                    font-family="monospace"
                    x={cX + cW + 4 + 5}
                    y={cY + 8 + 16 * 3}
                    onClick={handleDeleteStart}
                >
                    DELETE
                </text>
            </g>
            {/*
              *<g display={selected ? "inherit" : "none"} >
              *    <rect
              *        rx="3"
              *        width={btnW}
              *        height={btnH}
              *        x={cX - btnW - 4}
              *        y={cY - btnH - 4}
              *        fill="#282c34"
              *        stroke={selected ? "tomato" : "cadetblue"}
              *        stroke-width={2}
              *    />
              *    <line x1={cX - btnW - 4 + 10} y1={cY - btnW - 4 + 5} x2={cX - btnW - 4 + 10} y2={cY - 4 - 5} stroke="pink" stroke-width="1" />
              *    <line x1={cX - btnW - 4 + 5} y1={cY - btnW - 4 + 10} x2={cX - 4 - 5} y2={cY - btnW - 4 + 10} stroke="pink" stroke-width="1" />
              *</g>
              */}
            {/*
              *<g display={selected ? "inherit" : "none"}>
              *    <rect
              *        rx="3"
              *        width={btnW}
              *        height={btnH}
              *        x={cX + cW + 4}
              *        y={cY - btnW - 4}
              *        fill="#282c34"
              *        stroke={selected ? "tomato" : "cadetblue"}
              *        stroke-width={2}
              *        onClick={handleArchiveStart}
              *    />
              *    <line x1={cX + cW + 4 + 5} y1={cY - btnW - 4 + 5} x2={cX + cW + 4 + btnW - 5} y2={cY - 4 - 5} stroke="pink" stroke-width="1" />
              *    <line x1={cX + cW + 4 + 5} y1={cY - 4 - 5} x2={cX + cW + 4 + btnW - 5} y2={cY - btnW - 4 + 5} stroke="pink" stroke-width="1" />
              *</g>
              */}
            <g display={selected ? "inherit" : "none"}>
                <rect
                    rx="3"
                    width={btnW}
                    height={btnW}
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
                <line x1={cX + cW + 4 + 5} y1={cY + cH + 4 + btnW - 5} x2={cX + cW + 4 + btnW - 5} y2={cY + cH + 4 + btnW - 5} stroke="pink" stroke-width="1" />
                <line x1={cX + cW + 4 + btnW - 5} y1={cY + cH + 4 + btnW - 5} x2={cX + cW + 4 + btnW - 5} y2={cY + cH + 4 + 5} stroke="pink" stroke-width="1" />
            </g>
        </g>
    );

}
