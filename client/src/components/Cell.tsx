export default function Test({data, x, y, width, height}: {data:any, x:number, y:number, width:number, height:number}) {
    return (
        <g fill="white" stroke="green" stroke-width="5">
            <rect
                rx="10"
                width={width}
                height={height}
                x={x}
                y={y}
                fill="#282c34"
                stroke="cadetblue"
                stroke-width={2}
                onClick={() => { alert("lsdfjlsdf") }}
            />
            <foreignObject x={x+10} y={y+10} width={width-20} height={height-20}>
                <div style={{color: "pink"}} dangerouslySetInnerHTML={{__html: data }} />
            </foreignObject>
        </g>
    );

}
