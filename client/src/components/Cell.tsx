import React from 'react';

export default function Test({ data, x, y, width, height }: { data: any, x: number, y: number, width: number, height: number }) {

    const [selected, setSelected] = React.useState(false);

    const handleLinks = (data: string): string => {

        const youtubetmpl = `<iframe width="350" height="200" src="$1" title="YouTube video player" ></iframe>`;
        data = data.replace(/(https:\/\/www\.youtube\.com\/embed\/(\S+))/, youtubetmpl);
        data = data.replace(/((http(s?):)([/|.|\w|\s|-])*\.(?:jpg|gif|png))/, '<img width="300" src="$1"/>');

        return data
    };

    return (
        <g fill="white" stroke="green" stroke-width="5">
            <rect
                rx="6"
                width={width}
                height={height}
                x={x}
                y={y}
                fill="#282c34"
                stroke={selected ? "tomato" : "cadetblue"}
                stroke-width={2}
            />
            <foreignObject
                x={x + 10}
                y={y + 10}
                width={width - 20}
                height={height - 20}
                onClick={() => { setSelected(!selected) }}
            >
                <div
                    style={{ color: "pink", whiteSpace: "pre-wrap" }}
                    dangerouslySetInnerHTML={{ __html: handleLinks(data) }}
                />
            </foreignObject>
            <rect
                rx="3"
                width={20}
                height={20}
                x={x - 24}
                y={y - 24}
                fill="#282c34"
                stroke={selected ? "tomato" : "cadetblue"}
                stroke-width={2}
                display={selected ? "inherit" : "none"}
            />
            <rect
                rx="3"
                width={20}
                height={20}
                x={x + width + 4}
                y={y - 24}
                fill="#282c34"
                stroke={selected ? "tomato" : "cadetblue"}
                stroke-width={2}
                display={selected ? "inherit" : "none"}
            />
            <rect
                rx="3"
                width={20}
                height={20}
                x={x + width + 4}
                y={y + height + 4}
                fill="#282c34"
                stroke={selected ? "tomato" : "cadetblue"}
                stroke-width={2}
                display={selected ? "inherit" : "none"}
            />
        </g>
    );

}
