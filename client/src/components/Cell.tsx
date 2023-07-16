import React from 'react';
import { XY } from '../models/XY'
import { CellModel } from '../models/Mind'
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';

export default function Cell(
  { data, selected, setSelected, setDataChange, mousePosition, scaleIndex }:
    {
      data: CellModel,
      selected: boolean,
      setSelected: any,
      setDataChange: any,
      mousePosition: XY,
      scaleIndex: number,
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
          setSelected(data.id);
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
          dangerouslySetInnerHTML={{ __html: handleLinks(data.data) }}
        />
        <div
          style={{ color: "pink", whiteSpace: "pre-wrap" }}
          dangerouslySetInnerHTML={{ __html: handleLinks(data.status || "") }}
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
