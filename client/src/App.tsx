import React from 'react';
//import logo from './logo.svg';
import './App.css';
import Cell from './components/Cell'
import { XY } from './models/XY'
import { MindModel, CellModel } from './models/Mind'
import Slider from './components/Slider'
import CellForm from './components/CellForm'
import Editor from './components/Editor';
import TreeView from './components/TreeView';
import Fab from '@mui/material/Fab';
import AddIcon from '@mui/icons-material/Add';
import MoveToInboxIcon from '@mui/icons-material/MoveToInbox';
import AccountTreeIcon from '@mui/icons-material/AccountTree';

const Connections = ({
    data,
}: {
    data: CellModel,
}) => {
    return (
        <g>
            {data != null && data.cells?.map(cell => {
                return (
                    <line
                        x1={data.position[0]}
                        y1={data.position[1]}
                        x2={cell.position[0]}
                        y2={cell.position[1]}
                        stroke={"tomato"}
                        stroke-width={1}
                    />
                );
            })}
        </g>
    );
}

const Cells = ({
    data,
    selected,
    coords,
    setSelected,
    setDataChange,
    scaleIndex,
    startdata,
    setStartdata,
    formopenid,
    setFormopenId,
    layout,
    moveToStart,
    setMoveToStart,
}: {
    data: CellModel[],
    selected: string,
    coords: XY,
    setSelected: any,
    setDataChange: any,
    scaleIndex: number,
    startdata: CellModel,
    setStartdata: React.Dispatch<React.SetStateAction<CellModel>>,
    formopenid: string,
    setFormopenId: React.Dispatch<React.SetStateAction<string>>,
    layout: string,
    moveToStart: string,
    setMoveToStart: React.Dispatch<React.SetStateAction<string>>,
}) => {
    return (
        <g>
            {data != null && data.map(cell => {
                return (
                    <g>
                        {/*<Connections data={cell} />*/}
                        <Cell
                            data={cell}
                            selected={selected == cell.id}
                            setSelected={setSelected}
                            mousePosition={coords}
                            scaleIndex={scaleIndex}
                            setDataChange={setDataChange}
                            startdata={startdata}
                            setStartdata={setStartdata}
                            formopenid={formopenid}
                            setFormopenId={setFormopenId}
                            layout={layout}
                            moveToStart={moveToStart}
                            setMoveToStart={setMoveToStart}
                        />
                        <Cells
                            data={cell.cells || []}
                            selected={selected} coords={coords}
                            setSelected={setSelected}
                            scaleIndex={scaleIndex}
                            setDataChange={setDataChange}
                            startdata={startdata}
                            setStartdata={setStartdata}
                            formopenid={formopenid}
                            setFormopenId={setFormopenId}
                            layout={layout}
                            moveToStart={moveToStart}
                            setMoveToStart={setMoveToStart}
                        />
                    </g>
                );
            })}
        </g>
    );
}

function App() {
    const [data, setData] = React.useState<MindModel>([]);
    const [selected, setSelected] = React.useState("");
    const [coords, setCoords] = React.useState<XY>({ x: 0, y: 0, movX: 0, movY: 0 });
    const [scaleIndex, setScaleIndex] = React.useState(1);
    const [lastScaleIndex, setLastScaleIndex] = React.useState(1);
    const [datachange, setDataChange] = React.useState(0);
    const [width, setWidth] = React.useState(window.innerWidth);
    const [height, setHeight] = React.useState(window.innerHeight);

    const [moveToStart, setMoveToStart] = React.useState("");
    const [isViewMoved, setIsViewMoved] = React.useState(false);
    const [viewX, setViewX] = React.useState(0);
    const [viewY, setViewY] = React.useState(0);

    const [formopenid, setFormopenId] = React.useState("");

    const [layout, setLayout] = React.useState("main");

    const [treeview, setTreeview] = React.useState(false);

    const initialState = {
        id: "0",
        data: "",
        position: [0, 0, 0],
        size: [0, 0],
        status: "",
    };

    const [startdata, setStartdata] = React.useState<CellModel>(initialState);

    React.useEffect(() => {
        console.log("MOVE TO START ID = ", moveToStart);
    }, [moveToStart]);

    React.useEffect(() => {
        console.log("START DATA", startdata);
    }, [startdata]);

    React.useEffect(() => {
        if (isViewMoved) {
            setViewX(viewX - coords.movX * scaleIndex);
            setViewY(viewY - coords.movY * scaleIndex)
        }
    }, [coords]);

    const updateDimensions = () => {
        setWidth(window.innerWidth);
        setHeight(window.innerHeight - 10);
    }

    React.useEffect(() => {
        window.addEventListener("resize", updateDimensions);
        return () => window.removeEventListener("resize", updateDimensions);
    }, []);

    React.useEffect(() => {
        const updateWheel = (event: any) => {
            if (event.deltaY > 0) {
                if (scaleIndex < 10) {
                    setScaleIndex(scaleIndex + 1);
                }
            } else {
                if (scaleIndex > 1) {
                    setScaleIndex(scaleIndex - 1);
                }
            }
        }
        window.addEventListener("wheel", updateWheel);
        return () => window.removeEventListener("wheel", updateWheel);
    }, [scaleIndex]);

    React.useEffect(() => {
        const url = "http://localhost:2222/cells";
        fetch(url, {
            method: "GET",
            //mode: "no-cors",
            //body: JSON.stringify(a),
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
                            setData(res.data);
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
    }, [datachange]);

    React.useEffect(() => {
        if (lastScaleIndex < scaleIndex) {
            setViewX(viewX - width / 2);
            setViewY(viewY - height / 2);
        } else if (lastScaleIndex > scaleIndex) {
            setViewX(viewX + width / 2);
            setViewY(viewY + height / 2);
        }
        setLastScaleIndex(scaleIndex);
    }, [scaleIndex]);

    React.useEffect(() => {
        console.log("ViewXY", viewX, viewY);
    }, [viewX]);

    React.useEffect(() => {
        console.log("DATA", data);
    }, [data]);

    React.useEffect(() => {
        const handleWindowMouseMove = (event: any) => {
            //console.log(event);
            setCoords({
                x: event.clientX,
                y: event.clientY,
                movX: event.movementX,
                movY: event.movementY,
            });
        };
        window.addEventListener('mousemove', handleWindowMouseMove);

        return () => {
            window.removeEventListener(
                'mousemove',
                handleWindowMouseMove,
            );
        };
    }, []);

    return (
        <div className="App">
            <div style={{ position: 'absolute', top: '8px', left: '8px', display: formopenid != "" ? "inherit" : "none" }}>
                <Editor
                    startdata={startdata}
                    setStartdata={setStartdata}
                    setDataChange={setDataChange}
                    setFormopenId={setFormopenId}
                    initialState={initialState}
                />
            </div>
            <Slider scaleIndex={scaleIndex} setValue={setScaleIndex} />
            <Fab
                style={{ position: "absolute", bottom: "2em", right: "2em" }}
                size="large"
                color={"primary"}
                aria-label="add"
                onClick={() => setStartdata({ ...startdata, ["status"]: "new" })}
            >
                <AddIcon />
            </Fab>
            <Fab
                style={{ position: "absolute", bottom: "2em", right: "8em" }}
                variant="extended"
                size="medium"
                color={layout != "archive" ? "default" : "primary"}
                aria-label="archived"
                onClick={() => layout != "archive" ? setLayout("archive") : setLayout("main")}
            >
                <MoveToInboxIcon sx={{ mr: 1 }} />
                ARCHIVED
            </Fab>
            <Fab
                style={{ position: "absolute", bottom: "2em", right: "20em" }}
                variant="extended"
                size="medium"
                color={"default"}
                aria-label="archived"
                onClick={() => setTreeview(true)}
            >
                <AccountTreeIcon sx={{ mr: 1 }} />
                TREE VIEW
            </Fab>
            <TreeView
                data={data}
                treeview={treeview}
                setTreeview={setTreeview}
                setDataChange={setDataChange}
            />
            {/*
            <div style={{ color: "white", position: "fixed", top: "10px", right: "10px" }}>
                <p>
                    Mouse positioned at:{' '}
                    <b>
                        ({coords.x}, {coords.y}, {coords.movX}, {coords.movY})
                    </b>
                </p>
                <p>
                    window size:{' '}
                    <b>
                        ({width}, {height})
                    </b>
                </p>
                <p>
                    scale index:{' '}
                    <b>
                        ({scaleIndex})
                    </b>
                </p>
            </div>
            */}
            <CellForm
                coords={coords}
                scaleIndex={scaleIndex}
                viewX={viewX}
                viewY={viewY}
                startdata={startdata}
                setStartdata={setStartdata}
                formopenid={formopenid}
                setFormopenId={setFormopenId}
            />
            <svg
                onClick={() => {
                    console.log("click back");
                    setSelected("")
                }}
                onMouseDown={() => setIsViewMoved(true)}
                onMouseUp={() => setIsViewMoved(false)}
                viewBox={`${viewX} ${viewY} ${scaleIndex * width} ${scaleIndex * height}`} xmlns="http://www.w3.org/2000/svg">
                <Cells
                    data={data}
                    selected={selected}
                    coords={coords}
                    setSelected={setSelected}
                    scaleIndex={scaleIndex}
                    setDataChange={setDataChange}
                    startdata={startdata}
                    setStartdata={setStartdata}
                    formopenid={formopenid}
                    setFormopenId={setFormopenId}
                    layout={layout}
                    moveToStart={moveToStart}
                    setMoveToStart={setMoveToStart}
                />
            </svg>
        </div>
    );
}

export default App;
