import React from 'react';
import { XY } from '../models/XY'
import { CellModel } from '../models/Mind'
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';

export default function Cell(
    {
        data,
        selected,
        setSelected,
        setDataChange,
        mousePosition,
        scaleIndex,
        startdata,
        setStartdata,
        formopenid,
        setFormopenId,
        layout,
    }:
        {
            data: CellModel,
            selected: boolean,
            setSelected: any,
            setDataChange: any,
            mousePosition: XY,
            scaleIndex: number,
            startdata: CellModel,
            setStartdata: React.Dispatch<React.SetStateAction<CellModel>>,
            formopenid: string,
            setFormopenId: React.Dispatch<React.SetStateAction<string>>,
            layout: string,
        }) {

    const [archiveopen, setArchiveopen] = React.useState(false);
    const [isMoved, setIsMoved] = React.useState(false);
    const [isResized, setIsResized] = React.useState(false);
    const [cX, setCX] = React.useState(data.position[0]);
    const [cY, setCY] = React.useState(data.position[1]);
    const [cW, setCW] = React.useState(data.size[0]);
    const [cH, setCH] = React.useState(data.size[1]);

    const btnWstart = 20
    const btnHstart = 20

    const [btnW, setBtnW] = React.useState(btnWstart);
    const [btnH, setBtnH] = React.useState(btnHstart);

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

        //altitude index = 4.3546

        //https://www.google.com/maps/embed?pb=!1m10!1m8!1m3!1d24338.25750225145!2d30.503572899999973!3d50.51767350000001!3m2!1i1024!2i768!4f13.1!5e1!3m2!1suk!2sua!4v1702584679328!5m2!1suk!2sua
        //https://www.google.com/maps/embed?pb=!1m10!1m8!1m3!1d5000.09420383183!2d27.5035729!3d49.5176735!3m2!1i1024!2i768!4f13.1!5e0!3m2!1suk!2sua!4v1702582143755!5m2!1suk!2sua

        //https://www.google.com/maps/@50.5176735,30.5035729,14z?authuser=0&entry=ttu

        const gmaptmpl = `<iframe src="https://www.google.com/maps/embed?pb=!1m10!1m8!1m3!1d5000.09420383183!2d27.5035729!3d49.5176735!3m2!1i1024!2i768!4f13.1!5e1!3m2!1suk!2sua!4v1702582143755!5m2!1suk!2sua" width="600" height="450" style="border:0;margin:1em;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`;

        const youtubetmpl = `<iframe width="100%" style="aspect-ratio: 16 / 9" src="https://www.youtube.com/embed/$2" title="YouTube video player"></iframe>`;
        //data = data.replace(/(https:\/\/www\.youtube\.com\/embed\/(\S+))/, youtubetmpl);
        data = data.replace(/(https:\/\/www\.youtube\.com\/watch\?v=(\S{11}))/g, youtubetmpl);
        data = data.replace(/((http(s?):)([/|.|\w|\s|-])*\.(?:jpg|gif|png))/g, '<img width="100%" src="$1"/>');
        data = data.replace(/(#\S+)/g, '<b style="color: lightblue">$1</b>');
        data = data.replace(/(gggg)/g, gmaptmpl);

        return data
    };

    const handleAdd = (event: any) => {
        event.stopPropagation();
        const initialState = {
            id: data.id,
            data: "",
            position: [0, 0, 0],
            size: [0, 0],
            status: "new",
        }
        console.log("TTTTTTTTT", initialState);
        setStartdata(initialState);
    }

    const saveGeometry = () => {
        var tmpdata = data
        tmpdata.position = [cX, cY];
        tmpdata.size = [cW, cH];
        const url = "http://localhost:2222/cells/" + data.id;
        fetch(url, {
            method: "PATCH",
            //mode: "no-cors",
            body: JSON.stringify(tmpdata),
            headers: {
                //"Content-Type": "application/json",
                //Authorization: "Bearer " + "sdflsdjfl",
            },
            //credentials: "same-origin",
        }).then(
            function(response) {
                if (response.status === 200) {
                    response.json().then(function(res) {
                        console.log(res);
                        if (res.errors != null) {
                            console.log(res.errors);
                        } else {
                            setDataChange(Date.now());
                        }
                    });
                } else {
                    console.log(response);
                    console.log("ERROR: " + response.status + " - " + response.statusText);
                }
            },
            function(error) {
                console.log(error.message);
            }
        );
    }

    const handleDone = (event: any) => {
        event.stopPropagation();
        var tmpdata = data
        tmpdata.status = "done"
        const url = "http://localhost:2222/cells/" + data.id;
        fetch(url, {
            method: "PATCH",
            //mode: "no-cors",
            body: JSON.stringify(tmpdata),
            headers: {
                //"Content-Type": "application/json",
                //Authorization: "Bearer " + "sdflsdjfl",
            },
            //credentials: "same-origin",
        }).then(
            function(response) {
                if (response.status === 200) {
                    response.json().then(function(res) {
                        console.log(res);
                        if (res.errors != null) {
                            console.log(res.errors);
                        } else {
                            setDataChange(Date.now());
                        }
                    });
                } else {
                    console.log(response);
                    alert("ERROR: " + response.status + " - " + response.statusText);
                }
            },
            function(error) {
                alert(error.message);
            }
        );
    }

    const handleActive = (event: any) => {
        event.stopPropagation();
        var tmpdata = data
        tmpdata.status = "active"
        const url = "http://localhost:2222/cells/" + data.id;
        fetch(url, {
            method: "PATCH",
            //mode: "no-cors",
            body: JSON.stringify(tmpdata),
            headers: {
                //"Content-Type": "application/json",
                //Authorization: "Bearer " + "sdflsdjfl",
            },
            //credentials: "same-origin",
        }).then(
            function(response) {
                if (response.status === 200) {
                    response.json().then(function(res) {
                        console.log(res);
                        if (res.errors != null) {
                            console.log(res.errors);
                        } else {
                            setDataChange(Date.now());
                        }
                    });
                } else {
                    console.log(response);
                    alert("ERROR: " + response.status + " - " + response.statusText);
                }
            },
            function(error) {
                alert(error.message);
            }
        );
    }

    const handleArchive = (event: any) => {
        event.stopPropagation();
        var tmpdata = data
        tmpdata.status = "archive"
        const url = "http://localhost:2222/cells/" + data.id;
        fetch(url, {
            method: "PATCH",
            //mode: "no-cors",
            body: JSON.stringify(tmpdata),
            headers: {
                //"Content-Type": "application/json",
                //Authorization: "Bearer " + "sdflsdjfl",
            },
            //credentials: "same-origin",
        }).then(
            function(response) {
                if (response.status === 200) {
                    response.json().then(function(res) {
                        console.log(res);
                        if (res.errors != null) {
                            console.log(res.errors);
                        } else {
                            setDataChange(Date.now());
                        }
                    });
                } else {
                    console.log(response);
                    alert("ERROR: " + response.status + " - " + response.statusText);
                }
            },
            function(error) {
                alert(error.message);
            }
        );
    }

    const handleDeleteStart = (event: any) => {
        event.stopPropagation();
        setArchiveopen(true);
    }

    const handleDeleteClose = () => {
        setArchiveopen(false);
    }

    const handleDeleteSubmit = () => {
        const url = "http://localhost:2222/cells/" + data.id;
        fetch(url, {
            method: "DELETE",
            //mode: "no-cors",
            //body: JSON.stringify(formdata),
            headers: {
                //"Content-Type": "application/json",
                //Authorization: "Bearer " + "sdflsdjfl",
            },
            //credentials: "same-origin",
        }).then(
            function(response) {
                if (response.status === 200) {
                    response.json().then(function(res) {
                        console.log(res);
                        if (res.errors != null) {
                            console.log(res.errors);
                        } else {
                            setDataChange(Date.now());
                        }
                    });
                } else {
                    console.log(response);
                    alert("ERROR: " + response.status + " - " + response.statusText);
                }
            },
            function(error) {
                alert(error.message);
            }
        );
    }

    const isVisible = () => {
        if (formopenid == data.id) {
            return false;
        }
        if (data.status == "archive" && layout != "archive") {
            return false;
        }
        return true;
    }

    return (
        <g
            display={isVisible() ? "inherit" : "none"}
            style={{ userSelect: selected ? "auto" : "none" }}
            fill="white"
            stroke="green"
            stroke-width="5"
        >
            <Dialog fullWidth={false} open={archiveopen} onClose={handleDeleteClose}>
                <DialogContent>
                    Delete permanently. You are shure?
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleDeleteClose}>Cancel</Button>
                    <Button variant="outlined" color="error" onClick={handleDeleteSubmit}>DELETE</Button>
                </DialogActions>
            </Dialog>
            {data.synapses != null && data.synapses.map(synapse => {
                return (
                    <g>
                        {synapse.points != undefined && synapse.points![0] != undefined && synapse.points![1] != undefined && <line
                            x1={synapse.points![0][0]}
                            y1={synapse.points![0][1]}
                            x2={synapse.points![1][0]}
                            y2={synapse.points![1][1]}
                            stroke={synapse.color}
                            stroke-width={synapse.size}
                        />}
                        {synapse.points != undefined && synapse.points![2] != undefined && <line
                            x1={synapse.points![1][0]}
                            y1={synapse.points![1][1]}
                            x2={synapse.points![2][0]}
                            y2={synapse.points![2][1]}
                            stroke={synapse.color}
                            stroke-width={synapse.size}
                        />}
                    </g>
                );
            })}
            <rect
                rx="6"
                width={cW}
                height={cH}
                x={cX}
                y={cY}
                fill="#282c34"
                stroke={selected ? "tomato" : data.status == "done" ? "gray" : "cadetblue"}
                stroke-width={2}
            />
            <foreignObject
                x={cX + 10}
                y={cY + 10}
                width={cW - 20}
                height={cH - 20}
                onClick={(event) => {
                    event.stopPropagation();
                    setSelected(data.id);
                }}
                onDoubleClick={() => {
                    var tmpdata = data
                    tmpdata.status = "updated";
                    setStartdata(tmpdata);
                    setFormopenId(tmpdata.id);
                }}
                onMouseDown={(event) => {
                    event.stopPropagation();
                    setIsMoved(true)
                }}
                onMouseUp={() => {
                    setIsMoved(false)
                    saveGeometry();
                }}
            >
                <div
                    style={
                        data.status == "done"
                            ? { color: "gray", whiteSpace: "pre-wrap", textDecoration: "line-through", fontFamily: "monospace" }
                            : { color: "pink", whiteSpace: "pre-wrap", fontFamily: "monospace" }
                    }
                    dangerouslySetInnerHTML={{ __html: handleLinks(data.data) }}
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
                {
                    data.status != "done" ?
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
                        :
                        <text
                            className="svgbtn"
                            fill="tomato"
                            stroke="none"
                            font-size="14"
                            font-family="monospace"
                            x={cX + cW + 4 + 5}
                            y={cY + 8 + 16}
                            onClick={handleActive}
                        >
                            ACTIVE
                        </text>
                }
                {
                    data.status != "archive" ?
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
                        :
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
                            ACTIVE
                        </text>
                }
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
                        saveGeometry();
                    }}
                />
                <line x1={cX + cW + 4 + 5} y1={cY + cH + 4 + btnW - 5} x2={cX + cW + 4 + btnW - 5} y2={cY + cH + 4 + btnW - 5} stroke="pink" stroke-width="1" />
                <line x1={cX + cW + 4 + btnW - 5} y1={cY + cH + 4 + btnW - 5} x2={cX + cW + 4 + btnW - 5} y2={cY + cH + 4 + 5} stroke="pink" stroke-width="1" />
            </g>
        </g>
    );

}
