//import React from 'react';
//import logo from './logo.svg';
import './App.css';
import Cell from './components/Cell'
import Synapse from './components/Synapse'

function App() {
    return (
        <div className="App">
            <Cell top={1000} left={10} />
            <Synapse x1={200} y1={150} x2={300} y2={300} size={4} color="tomato" />
            <Synapse x1={200} y1={150} x2={300} y2={20} size={4} color="green"  />
            <Synapse x1={200} y1={150} x2={50} y2={50} size={4} color="lightblue" />
            <Synapse x1={200} y1={150} x2={50} y2={300} size={4} color="gold" />
            <Synapse x1={200} y1={150} x2={200} y2={300} size={4} color="red" />
            <Synapse x1={200} y1={150} x2={200} y2={50} size={4} color="pink" />
            <Synapse x1={200} y1={150} x2={400} y2={150} size={4} color="purple" />
            <Synapse x1={200} y1={0} x2={50} y2={0} size={1} color="gold" />
            <Cell top={1000} left={500} />
        </div>
    );
}

export default App;
