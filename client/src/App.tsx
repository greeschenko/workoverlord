//import React from 'react';
//import logo from './logo.svg';
import './App.css';

function App() {
    return (
        <div className="App">
            <div style={{ position: 'absolute', top: '0px', left: '10px' }}>
                <svg width="400" height="400" xmlns="http://www.w3.org/2000/svg">
                    <g fill="white" stroke="green" stroke-width="5">
                        <rect
                            rx="10"
                            width="396"
                            height="396"
                            x="1"
                            y="1"
                            fill="#282c34"
                            stroke="cadetblue"
                            stroke-width={2}
                            onClick={() => { alert("lsdfjlsdf") }}
                        />
                    </g>
                </svg>
                <div style={{ position: 'absolute', top: '20px', left: '20px', color: "tomato" }}>
                    Lorem ipsum dolor sit amet, consetetur sadipscing elitr,
                    sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat,
                    sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.
                    Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.
                    <br />
                    <br />
                    <iframe
                        width="350"
                        height="200"
                        src="https://www.youtube.com/embed/wHPaGn5Q5ug"
                        title="YouTube video player"
                    ></iframe>
                </div>
            </div>
            <div style={{ position: 'absolute', top: '200px', left: '400px' }}>
                <svg width="200" height="200" xmlns="http://www.w3.org/2000/svg">
                    <line x1="10" y1="20" x2="100" y2="100" stroke="cadetblue"
                        stroke-width={2}
                    />
                </svg>
            </div>
            <div style={{ position: 'absolute', top: '150px', left: '500px' }}>
                <svg width="400" height="400" xmlns="http://www.w3.org/2000/svg">
                    <g fill="white" stroke="green" stroke-width="5">
                        <rect
                            rx="10"
                            width="300"
                            height="200"
                            x="1"
                            y="1"
                            fill="#282c34"
                            stroke="cadetblue"
                            stroke-width={2}
                            onClick={() => { alert("lsdfjlsdf") }}
                        />
                        <text
                            x="50"
                            y="50"
                            font-size="20px"
                            font-variant="small-caps"
                            font-weight="light"
                            font-family="sans-serif"
                            stroke="none"
                            fill="tomato"
                        >
                            <tspan>Test</tspan>
                            <tspan x="50" dy="20px">Test</tspan>
                        </text>
                    </g>
                </svg>
            </div>
        </div>
    );
}

export default App;
