export type MindModel = CellModel[]

export type CellModel = {
    id?: string;
    data: string;
    status?: string;
    tags?: string;
    size: number[];
    position: number[];
    cells?: CellModel[];
    synapses?: SynapseModel[];
}

export type SynapseModel = {
    points?: number[][];
    size?: number;
    color?: string;
    linetype?: string;
    endtype?: string;
}
