import React from 'react';
//import logo from './logo.svg';
import './App.css';
import Cell from './components/Cell'
import { XY } from './models/XY'
import { MindModel, CellModel } from './models/Mind'
import Slider from './components/Slider'
import CellForm from './components/CellForm'
import Fab from '@mui/material/Fab';
import AddIcon from '@mui/icons-material/Add';
import MoveToInboxIcon from '@mui/icons-material/MoveToInbox';

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
  handleSynapses,
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
  handleSynapses: (id: string) => void,
}) => {
  return (
    <g>
      {data != null && data.map(cell => {
        return (
          <g>
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
              handleSynapses={handleSynapses}
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
              handleSynapses={handleSynapses}
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

  const [isViewMoved, setIsViewMoved] = React.useState(false);
  const [viewX, setViewX] = React.useState(0);
  const [viewY, setViewY] = React.useState(0);

  const [formopenid, setFormopenId] = React.useState("");

  const [layout, setLayout] = React.useState("main");

  const [startdata, setStartdata] = React.useState<CellModel>({
    id: "0",
    data: "",
    position: [0, 0, 0],
    size: [0, 0],
    status: "",
  });

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

  const handleSynapses = (id: string) => {
    console.log("RRRRRRRR", id, coords.movX * scaleIndex, coords.movY * scaleIndex);
    let tmpdata = data;
    let element: CellModel;
    for (let i in id.split(" ")) {
      let chain = id.split(" ")[i];
      if (i == "0") {
        for (let j in tmpdata) {
          if (tmpdata[j].id == chain) {
            element = tmpdata[j];
          }
        }
      } else {
        for (let j in element!.cells) {
          if (element!.cells[Number(j)].id.slice(-4) == chain) {
            element = element!.cells[Number(j)];
          }
        }
      }
    }

    if (id.split(" ").length > 1) {
      element!.synapses![0].points![0][0] += coords.movX * scaleIndex;
      element!.synapses![0].points![0][1] += coords.movY * scaleIndex
    }

    for (let i in element!.cells!) {
      let child = element!.cells![i];
      child.synapses![0].points![1][0] += coords.movX * scaleIndex;
      child.synapses![0].points![1][1] += coords.movY * scaleIndex;
    }

    setData(tmpdata);

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
      <Slider scaleIndex={scaleIndex} setValue={setScaleIndex} />
      <Fab
        style={{ position: "absolute", bottom: "2em", right: "2em" }}
        size="large"
        color={"primary"}
        aria-label="add"
        onClick={() => setStartdata({ ...startdata, ["status"]: "active" })}
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
      <svg
        onClick={() => {
          console.log("click back");
          setSelected("")
        }}
        onMouseDown={() => setIsViewMoved(true)}
        onMouseUp={() => setIsViewMoved(false)}
        viewBox={`${viewX} ${viewY} ${scaleIndex * width} ${scaleIndex * height}`} xmlns="http://www.w3.org/2000/svg">
        <CellForm
          coords={coords}
          scaleIndex={scaleIndex}
          setDataChange={setDataChange}
          viewX={viewX}
          viewY={viewY}
          startdata={startdata}
          setStartdata={setStartdata}
          formopenid={formopenid}
          setFormopenId={setFormopenId}
        />
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
          handleSynapses={handleSynapses}
        />
      </svg>
    </div>
  );
}

export default App;
