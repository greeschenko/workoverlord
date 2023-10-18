import * as React from 'react';

import { CellModel } from '../models/Mind'
import { XY } from '../models/XY'

export default function FormDialog(
  {
    coords,
    scaleIndex,
    setDataChange,
    viewX,
    viewY,
    startdata,
    setStartdata,
    formopenid,
    setFormopenId,
  }: {
    coords: XY,
    scaleIndex: number,
    setDataChange: React.Dispatch<React.SetStateAction<number>>
    viewX: number,
    viewY: number,
    startdata: CellModel,
    setStartdata: React.Dispatch<React.SetStateAction<CellModel>>,
    formopenid: string,
    setFormopenId: React.Dispatch<React.SetStateAction<string>>,
  }) {
  const [pendingPosition, setPendingPosition] = React.useState(false);
  const [pendingSize, setPendingSize] = React.useState(false);
  const [cX, setCX] = React.useState(0);
  const [cY, setCY] = React.useState(0);
  const [cW, setCW] = React.useState(0);
  const [cH, setCH] = React.useState(0);

  const [cursorposition, setCursorposition] = React.useState(0);

  const initialState = {
    id: "0",
    data: "",
    position: [0, 0, 0],
    size: [0, 0],
    status: "",
  }

  const [formdata, setFormdata] = React.useState<CellModel>(initialState);
  const [textcontent, setContent] = React.useState("write something here...");

  const textareaEl = React.useRef<HTMLDivElement>(null)


  React.useEffect(() => {
    if (startdata.status == "new") {
      setPendingPosition(true);
      console.log("start here!!!");
    } else if (startdata.status == "updated") {
      console.log("need to update the form");
      setFormdata(startdata);
      setCX(viewX + startdata.position[0] * scaleIndex);
      setCY(viewY + startdata.position[1] * scaleIndex);
      setCW(startdata.size[0] * scaleIndex);
      setCH(startdata.size[1] * scaleIndex);
      setformtext(startdata.data);
      setFormopenId(startdata.id);
    }
  }, [startdata]);

  React.useEffect(() => {
    console.log("CURSOR", cursorposition);
  }, [cursorposition]);

  React.useEffect(() => {
    console.log("FORMDATA", formdata);
  }, [formdata]);

  const handleSubmit = () => {
    let method = "POST";
    let id = "0";

    if (formdata.status == "updated") {
      method = "PATCH";
      id = formdata.id
      formdata.status = "new";
    }

    const url = "http://localhost:2222/cells/" + id;
    fetch(url, {
      method: method,
      //mode: "no-cors",
      body: JSON.stringify(formdata),
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
              setFormdata(initialState);
              setStartdata(initialState);
              setDataChange(Date.now());
              clearform();
              setFormopenId("");
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

  const handleClose = () => {
    setFormopenId("");
    setFormdata(initialState);
    setStartdata(initialState);
  };

  const clearform = () => {
    var el = document.getElementById('forminput');
    el!.innerText = textcontent;
  }

  const setformtext = (text: string) => {
    var el = document.getElementById('forminput');
    el!.innerText = text;
  }

  return (
    <g fill="white" stroke="green" stroke-width="5">
      <foreignObject
        x={0}
        y={0}
        width={"100%"}
        height={"100%"}
        display={pendingPosition || pendingSize ? "inherit" : "none"}
      >
        <div
          style={{
            //display: pendingPosition || pendingSize ? "block" : "none",
            position: "fixed",
            //backgroundColor: "red",
            top: 0,
            left: 0,
            width: "100%",
            height: "100%",
            cursor: "crosshair",
          }}
          onClick={(e) => {
            console.log(e);
            if (pendingPosition) {
              setPendingPosition(false);
              setPendingSize(true);
              setFormdata({
                ...formdata,
                ['position']:
                  [
                    Math.ceil(viewX + coords.x * scaleIndex),
                    Math.ceil(viewY + coords.y * scaleIndex),
                    0,
                  ]
              });
              setCX(viewX + coords.x * scaleIndex);
              setCY(viewY + coords.y * scaleIndex);
            } else if (pendingSize) {
              setPendingPosition(false);
              setPendingSize(false);
              setFormdata({
                ...formdata,
                ['size']: [
                  Math.ceil((viewX + coords.x * scaleIndex) - formdata.position[0]),
                  Math.ceil((viewY + coords.y * scaleIndex) - formdata.position[1]),
                ]
              });
              setCW((viewX + coords.x * scaleIndex) - formdata.position[0]);
              setCH((viewY + coords.y * scaleIndex) - formdata.position[1]);
              setFormopenId("0");
            }
          }}
        ></div>
      </foreignObject>
      <rect
        display={pendingSize ? "inherit" : "none"}
        rx="6"
        width={(viewX + coords.x * scaleIndex) - formdata.position[0]}
        height={(viewY + coords.y * scaleIndex) - formdata.position[1]}
        x={cX}
        y={cY}
        fill="none"
        stroke={"cyan"}
        stroke-width={1}
        stroke-dasharray={"5,5"}
      />
      <g display={formopenid != "" ? "inherit" : "none"} >
        <rect
          rx="6"
          width={cW}
          height={cH}
          x={cX}
          y={cY}
          fill="none"
          stroke={"cyan"}
          stroke-width={1}
          stroke-dasharray={"5,5"}
        />
        <foreignObject
          x={cX + 10}
          y={cY + 10}
          width={cW - 20}
          height={cH - 40}
        //onClick={(event) => {
        //event.stopPropagation();
        //setSelected(data.id);
        //}}
        //onDoubleClick={() => alert("lsdjfldsfj")}
        //        onMouseDown={(event) => {
        //          event.stopPropagation();
        //          setIsMoved(true)
        //        }}
        //        onMouseUp={() => {
        //          setIsMoved(false)
        //        }}
        >

          <div
            id="forminput"
            ref={textareaEl}
            contentEditable={true}
            suppressContentEditableWarning={true}
            style={{
              color: "pink",
              whiteSpace: "pre-wrap",
            }}
            dangerouslySetInnerHTML={{ __html: textcontent }}
            onInput={e => {
              setFormdata({ ...formdata, ["data"]: e.currentTarget.innerText || "" });
              console.log(window.getSelection());
              setCursorposition(window.getSelection()!.anchorOffset);
            }}
          />

        </foreignObject>
        <text
          className="svgbtn"
          fill="tomato"
          stroke="none"
          font-size="14"
          font-family="monospace"
          x={cX + cW + 4 + 5}
          y={cY + 8}
          onClick={handleSubmit}
        >
          SAVE
        </text>
        <text
          className="svgbtn"
          fill="tomato"
          stroke="none"
          font-size="14"
          font-family="monospace"
          x={cX + cW + 4 + 5}
          y={cY + 8 + 16}
          onClick={handleClose}
        >
          CANCEL
        </text>
      </g>
    </g>
  );
}
