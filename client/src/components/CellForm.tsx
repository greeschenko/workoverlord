import * as React from 'react';

import { CellModel } from '../models/Mind';
import { XY } from '../models/XY';

export default function FormDialog(
    {
        coords,
        scaleIndex,
        viewX,
        viewY,
        startdata,
        setStartdata,
        formopenid,
        setFormopenId,
    }: {
        coords: XY,
        scaleIndex: number,
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

    const inputRef = React.useRef<HTMLInputElement | null>(null);

    const handleClickFile = () => {
        inputRef.current?.click();
    };

    const handleFileChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
        const fileObj = event.target.files && event.target.files[0];
        if (!fileObj) {
            return;
        }

        console.log('fileObj is', fileObj);

        // 👇️ reset file input
        if (inputRef.current) {
            inputRef.current.value = '';
        }

        // 👇️ is now empty
        console.log(event.target.files);

        // 👇️ can still access file object here
        console.log(fileObj);
        console.log(fileObj.name);

        const base64Content = await readFileAsBase64(fileObj);
        console.log('File content in base64:', base64Content);
    };

    // Function to read file content as base64
    const readFileAsBase64 = (file: File): Promise<string> => {
        return new Promise((resolve, reject) => {
            const reader = new FileReader();
            reader.onload = () => {
                if (reader.result) {
                    resolve(reader.result.toString());
                } else {
                    reject(new Error('Failed to read file as base64.'));
                }
            };
            reader.onerror = (error) => reject(error);
            reader.readAsDataURL(file);
        });
    };

    const [cursorposition, setCursorposition] = React.useState(0);

    React.useEffect(() => {
        if (startdata.status == "active") {
            setPendingPosition(true);
            console.log("start here!!!");
        } else if (startdata.status == "new") {
            setPendingPosition(true);
            setStartdata({ ...startdata, ["status"]: "new" });
        } else if (startdata.status == "updated") {
            console.log("need to update the form");
            setStartdata(startdata);
            setCX(viewX + startdata.position[0] * scaleIndex);
            setCY(viewY + startdata.position[1] * scaleIndex);
            setCW(startdata.size[0] * scaleIndex);
            setCH(startdata.size[1] * scaleIndex);
            //setformtext(startdata.data);
            setFormopenId(startdata.id);
        }
    }, [startdata]);

    React.useEffect(() => {
        console.log("CURSOR", cursorposition);
    }, [cursorposition]);

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
                            setStartdata({
                                ...startdata,
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
                            setStartdata({
                                ...startdata,
                                ['size']: [
                                    Math.ceil((viewX + coords.x * scaleIndex) - startdata.position[0]),
                                    Math.ceil((viewY + coords.y * scaleIndex) - startdata.position[1]),
                                ]
                            });
                            setCW((viewX + coords.x * scaleIndex) - startdata.position[0]);
                            setCH((viewY + coords.y * scaleIndex) - startdata.position[1]);
                            setFormopenId("0");
                        }
                    }}
                ></div>
            </foreignObject>
            <rect
                display={pendingSize ? "inherit" : "none"}
                rx="6"
                width={(viewX + coords.x * scaleIndex) - startdata.position[0]}
                height={(viewY + coords.y * scaleIndex) - startdata.position[1]}
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

                    {/*
                    <div
                        id="forminput"
                        ref={textareaEl}
                        contentEditable={true}
                        suppressContentEditableWarning={true}
                        style={{
                            color: "pink",
                            whiteSpace: "pre-wrap",
                            fontFamily: "monospace",
                        }}
                        dangerouslySetInnerHTML={{ __html: textcontent }}
                        onInput={e => {
                            setFormdata({ ...formdata, ["data"]: e.currentTarget.innerText || "" });
                            console.log(window.getSelection());
                            setCursorposition(window.getSelection()!.anchorOffset);
                        }}
                    />
                        */}
                    <input
                        type="file"
                        ref={inputRef}
                        style={{ display: "none" }}
                        onChange={handleFileChange}
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
                    onClick={handleClickFile}
                >
                    ADD IMG
                </text>
            </g>
        </g>
    );
}
