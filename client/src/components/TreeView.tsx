import * as React from 'react';
import Box from '@mui/material/Box';
import Drawer from '@mui/material/Drawer';
import Typography from '@mui/material/Typography';
import { CellModel } from '../models/Mind'
import { SimpleTreeView } from '@mui/x-tree-view/SimpleTreeView';
import { TreeItem } from '@mui/x-tree-view/TreeItem';
import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';

function TreeElement({ cell }: { cell: CellModel }) {
    return (
        <TreeItem itemId={cell.id} label={"[" + cell.id + "] " + cell.data.split("\n")[0]}>
            {cell.cells?.map(chcell => {
                return (
                    <TreeElement cell={chcell} />
                );
            })}
        </TreeItem>
    );
}

export default function TemporaryDrawer({
    data,
    treeview,
    setTreeview,
    setDataChange,
}: {
    data: CellModel[],
    treeview: boolean,
    setTreeview: React.Dispatch<React.SetStateAction<boolean>>
    setDataChange: React.Dispatch<React.SetStateAction<number>>,
}) {
    const [archiveopen, setArchiveopen] = React.useState(false);
    const [lastSelectedItem, setLastSelectedItem] = React.useState<string | null>(
        null,
    );

    const handleItemSelectionToggle = (
        event: React.SyntheticEvent,
        itemId: string,
        isSelected: boolean,
    ) => {
        if (isSelected) {
            setLastSelectedItem(itemId);
        }
    };

    const toggleDrawer = (newOpen: boolean) => () => {
        setTreeview(newOpen);
    };

    const handleDeleteStart = (event: any) => {
        event.stopPropagation();
        setArchiveopen(true);
    }

    const handleDeleteClose = () => {
        setArchiveopen(false);
    }

    const handleDeleteSubmit = () => {
        const url = "http://localhost:2222/cells/" + lastSelectedItem;
        fetch(url, {
            method: "DELETE",
            //mode: "no-cors",
            //body: JSON.stringify(formdata),
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
                            setArchiveopen(false);
                            setDataChange(Date.now());
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

    const DrawerList = (
        <Box sx={{ height: 220, flexGrow: 1, maxWidth: 400 }}>
            <Typography>
                {lastSelectedItem == null
                    ? 'No item selection recorded'
                    : `Last selected item: ${lastSelectedItem}`}
            </Typography>
            {lastSelectedItem != null && <Stack sx={{ padding: '1em' }} spacing={2} direction="row" justifyContent="flex-end">
                <Button size="small" variant="outlined">Edit</Button>
                <Button size="small" variant="outlined" color="error" onClick={handleDeleteStart}>Delete</Button>
            </Stack>}
            <SimpleTreeView onItemSelectionToggle={handleItemSelectionToggle}>
                {data.map(cell => {
                    return (
                        <TreeElement cell={cell} />
                    );
                })}
            </SimpleTreeView>
        </Box>
    );

    return (
        <div>
            <Dialog fullWidth={false} open={archiveopen} onClose={handleDeleteClose}>
                <DialogContent>
                    Delete permanently. You are shure?
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleDeleteClose}>Cancel</Button>
                    <Button variant="outlined" color="error" onClick={handleDeleteSubmit}>DELETE</Button>
                </DialogActions>
            </Dialog>
            <Drawer open={treeview} anchor="left" onClose={toggleDrawer(false)}>
                {DrawerList}
            </Drawer>
        </div>
    );
}

