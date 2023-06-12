import React from 'react';
//import logo from './logo.svg';
import './App.css';
import Cell from './components/Cell'
import Synapse from './components/Synapse'
import { XY } from './models/XY'

const tmpdata = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed mollis mollis mi ut ultricies. Nullam magna ipsum, porta vel dui convallis, rutrum
imperdiet eros. Aliquam erat volutpat.

https://www.youtube.com/embed/wHPaGn5Q5ug
`;

const tmpdata1 = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed mollis mollis
mi ut ultricies. Nullam magna ipsum, porta vel dui convallis, rutrum
imperdiet eros. Aliquam erat volutpat.

    - lsdjflsdjf
    - lsdjflsdjf

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed mollis mollis
mi ut ultricies. Nullam magna ipsum, porta vel dui convallis, rutrum
imperdiet eros. Aliquam erat volutpat.

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed mollis mollis
mi ut ultricies. Nullam magna ipsum, porta vel dui convallis, rutrum
imperdiet eros. Aliquam erat volutpat.

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed mollis mollis
mi ut ultricies. Nullam magna ipsum, porta vel dui convallis, rutrum
imperdiet eros. Aliquam erat volutpat.
`;

const tmpdata2 = `https://cdna.artstation.com/p/assets/images/images/053/956/262/medium/sentron-edgerunner-copy.jpg`;

function App() {

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
                            console.log(res);
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
            <div style={{ color: "white", position: "fixed", top: "10px", right: "10px" }}>
                <p>
                    Mouse positioned at:{' '}
                    <b>
                        ({coords.x}, {coords.y})
                    </b>
                </p>
            </div>
            <svg width={1000} height={1000} xmlns="http://www.w3.org/2000/svg">
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
                    height={300}
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
            </svg>
        </div>
    );
}

export default App;
