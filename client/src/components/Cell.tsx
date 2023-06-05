
export default function Cell({top, left}:{top: number, left: number}) {
    return (
        <div style={{ position: 'absolute', top: top, left: left }}>
            <svg width={400} height={400} xmlns="http://www.w3.org/2000/svg">
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
    );
}
