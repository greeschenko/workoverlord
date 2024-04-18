import React, { useState } from 'react';
import AceEditor from 'react-ace';
import { CellModel } from '../models/Mind';
import 'ace-builds/src-noconflict/mode-text';
import 'ace-builds/src-noconflict/theme-monokai';
import 'ace-builds/src-noconflict/keybinding-vim';

interface EditableTextProps {
    startdata: CellModel,
    setStartdata: React.Dispatch<React.SetStateAction<CellModel>>,
    setDataChange: React.Dispatch<React.SetStateAction<number>>,
    setFormopenId: React.Dispatch<React.SetStateAction<string>>,
    initialState: CellModel,
}

const EditableText: React.FC<EditableTextProps> = ({
    startdata,
    setStartdata,
    setDataChange,
    setFormopenId,
    initialState,
}) => {

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

        setStartdata({ ...startdata, ["data"]: startdata.data + '\n<img width="100%" src="' + base64Content + '"/>' || "" })
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

    const handleAddToEnd = () => {
        // Додає "The End" в кінець тексту
        setStartdata({ ...startdata, ["data"]: startdata.data + ' The End' || "" })
    };

    const handleAddToStart = () => {
        // Додає "Start here" на початок тексту
        setStartdata({ ...startdata, ["data"]: 'Start here ' + startdata.data || "" })
    };

    const handleSubmit = () => {
        let method = "POST";
        let id = "0";

        if (startdata.status == "updated") {
            method = "PATCH";
            id = startdata.id
            startdata.status = "active";
        }

        if (startdata.status == "active") {
            id = startdata.id
        }

        const url = "http://localhost:2222/cells/" + id;
        fetch(url, {
            method: method,
            //mode: "no-cors",
            body: JSON.stringify(startdata),
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
                            setStartdata(initialState);
                            setDataChange(Date.now());
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
        setStartdata(initialState);
    };

    return (
        <>
            <input
                type="file"
                ref={inputRef}
                style={{ display: "none" }}
                onChange={handleFileChange}
            />
            <AceEditor
                mode="text"
                theme="monokai"
                name="editableText"
                //onBlur={() => handleBlur(text)}
                onChange={(value) => setStartdata({ ...startdata, ["data"]: value || "" })}
                fontSize={14}
                showPrintMargin={true}
                showGutter={true}
                highlightActiveLine={true}
                value={startdata.data}
                setOptions={{
                    useWorker: false,
                    showLineNumbers: true,
                    tabSize: 2,
                    keyboardHandler: "ace/keyboard/vim", // Додаємо підтримку Vim
                }}
                style={{ border: '1px solid #ccc', padding: '8px', position: 'relative', width: "80vw" }}
            />
            <div style={{ position: 'absolute', bottom: '8px', right: '8px' }}>
                <button onClick={handleClickFile}>ADD IMG</button>
                <button onClick={handleSubmit}>SAVE</button>
                <button onClick={handleClose}>CLOSE</button>
            </div>
        </>
    );
};

export default EditableText;
