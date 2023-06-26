import React from 'react';
//import logo from './logo.svg';
import './App.css';
import Cell from './components/Cell'
import Synapse from './components/Synapse'
import { XY } from './models/XY'
import { MindModel, CellModel, SynapseModel } from './models/Mind'

import Fab from '@mui/material/Fab';
import AddIcon from '@mui/icons-material/Add';

const Cells = ({ data, selected, coords, setSelected }: { data: CellModel[], selected: string, coords: XY, setSelected: any }) => {
    return (
        <g>
            {data != null && data.map(cell => {
                return (
                    <g>
                        <Cell
                            x={cell.position[0]}
                            y={cell.position[1]}
                            width={cell.size[0]}
                            height={cell.size[1]}
                            data={cell.data}
                            selected={selected == cell.id}
                            onClick={() => {
                                setSelected(cell.id);
                            }}
                            mousePosition={coords}
                        />
                        <Cells data={cell.cells || []} selected={selected} coords={coords} setSelected={setSelected} />
                    </g>
                );
            })}
        </g>
    );
}

function App() {
    const [data, setData] = React.useState<MindModel>([]);

    const [selected, setSelected] = React.useState("");

    const start: XY = { x: 0, y: 0 };

    const [coords, setCoords] = React.useState(start);

    const [datachange, setDataChange] = React.useState(0);

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
        console.log("DATA", data);
    }, [data]);

    React.useEffect(() => {
        const handleWindowMouseMove = (event: any) => {
            setCoords({
                x: event.clientX,
                y: event.clientY,
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
            <Fab style={{position: "absolute", top: "90%", left: "90%"}} size="large" color="primary" aria-label="add">
                <AddIcon />
            </Fab>
            <div style={{ color: "white", position: "fixed", top: "10px", right: "10px" }}>
                <p>
                    Mouse positioned at:{' '}
                    <b>
                        ({coords.x}, {coords.y})
                    </b>
                </p>
            </div>
            <svg width={10000} height={10000} xmlns="http://www.w3.org/2000/svg">
                <Cells data={data} selected={selected} coords={coords} setSelected={setSelected} />
                {/*
                <Cell
                    x={100}
                    y={20}
                    width={320}
                    height={300}
                    data={tmpdata}
                    selected={selected == "0"}
                    onClick={() => {
                        setSelected("0");
                    }}
                    mousePosition={coords}
                />
                <Cell
                    x={100}
                    y={400}
                    width={320}
                    height={550}
                    data={tmpdata1}
                    selected={selected == "1"}
                    onClick={() => {
                        setSelected("1");
                    }}
                    mousePosition={coords}
                />
                <Cell
                    x={500}
                    y={250}
                    width={320}
                    height={300}
                    data={tmpdata2}
                    selected={selected == "2"}
                    onClick={() => {
                        setSelected("2");
                    }}
                    mousePosition={coords}
                />
                <Synapse x1={420} y1={50} x2={600} y2={250} size={2} />
                <Synapse x1={250} y1={320} x2={250} y2={400} size={2} />
                <Synapse x1={420} y1={620} x2={650} y2={620} size={2} />
                <Synapse x1={650} y1={620} x2={650} y2={550} size={2} />
                */}
            </svg>
        </div>
    );
}

export default App;
