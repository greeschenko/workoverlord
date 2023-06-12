export type Mind = Cell[]

export type Cell = {
    id: string;
    data: string;
    status: string;
    tags?: string;
    size?: number[];
    position?: number[];
    cells?: Cell[];
    synapses?: Synapse[];
}

export type Synapse = {
    points?: number[][];
    size?: number;
    color?: string;
    linetype?: string;
    endtype?: string;
}
