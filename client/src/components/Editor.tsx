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

        if (startdata.status == "new") {
            id = startdata.id
            startdata.status = "active";
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
            <AceEditor
                mode="text"
                theme="monokai"
                name="editableText"
                //onBlur={() => handleBlur(text)}
                onChange={(value) => setStartdata({ ...startdata, ["data"]: value || "" })}
                fontSize={10}
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
                style={{ border: '1px solid #ccc', padding: '8px', position: 'relative' }}
            />
            <div style={{ position: 'absolute', bottom: '8px', right: '8px' }}>
                <button onClick={handleSubmit}>SAVE</button>
                <button onClick={handleClose}>CLOSE</button>
                <button onClick={handleAddToStart}>Додати "Start here"</button>
                <button onClick={handleAddToEnd}>Додати "The End"</button>
            </div>
        </>
    );
};

export default EditableText;
