//import React from 'react';
//import logo from './logo.svg';
import './App.css';
import Cell from './components/Cell'
import Synapse from './components/Synapse'



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
    return (
        <div className="App">
            <svg width={1000} height={1000} xmlns="http://www.w3.org/2000/svg">
                <Cell x={100} y={20} width={320} height={300} data={tmpdata} />
                <Cell x={100} y={400} width={320} height={300} data={tmpdata1} />
                <Cell x={500} y={250} width={320} height={300} data={tmpdata2} />
                <Synapse x1={420} y1={50} x2={600} y2={250} size={2} />
                <Synapse x1={250} y1={320} x2={250} y2={400} size={2} />
                <Synapse x1={420} y1={620} x2={650} y2={620} size={2} />
                <Synapse x1={650} y1={620} x2={650} y2={550} size={2} />
            </svg>
        </div>
    );
}

export default App;
