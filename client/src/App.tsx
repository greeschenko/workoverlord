import React from 'react';
//import logo from './logo.svg';
import './App.css';
import Cell from './components/Cell'
import { XY } from './models/XY'
import { MindModel, CellModel } from './models/Mind'
import Slider from './components/Slider'
import CellForm from './components/CellForm'

const Cells = ({ data, selected, coords, setSelected, setDataChange, scaleIndex }:
  { data: CellModel[], selected: string, coords: XY, setSelected: any, setDataChange: any, scaleIndex: number }) => {
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
            />
            <Cells
              data={cell.cells || []}
              selected={selected} coords={coords}
              setSelected={setSelected}
              scaleIndex={scaleIndex}
              setDataChange={setDataChange}
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
      <Slider scaleIndex={scaleIndex} setValue={setScaleIndex} />
      <CellForm
        coords={coords}
        scaleIndex={scaleIndex}
        setDataChange={setDataChange}
        viewX={viewX}
        viewY={viewY}
      />
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
        <Cells
          data={data}
          selected={selected}
          coords={coords}
          setSelected={setSelected}
          scaleIndex={scaleIndex}
          setDataChange={setDataChange}
        />
      </svg>
    </div>
  );
}

export default App;
