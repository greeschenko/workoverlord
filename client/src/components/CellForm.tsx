import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import Fab from '@mui/material/Fab';
import AddIcon from '@mui/icons-material/Add';
import Alert from '@mui/material/Alert';
import Stack from '@mui/material/Stack';

import { CellModel } from '../models/Mind'
import { XY } from '../models/XY'



export default function FormDialog(
  { coords, scaleIndex, setDataChange }: {
    coords: XY,
    scaleIndex: number,
    setDataChange: React.Dispatch<React.SetStateAction<number>>
  }) {
  const [pendingPosition, setPendingPosition] = React.useState(false);
  const [pendingSize, setPendingSize] = React.useState(false);
  const [open, setOpen] = React.useState(false);

  const [formdata, setFormdata] = React.useState<CellModel>({
    id: "0",
    data: "",
    tags: "",
    position: [0, 0, 0],
    size: [0, 0],
    status: "new",
  });

  React.useEffect(() => {
    console.log("FORMDATA", formdata);
  }, [formdata]);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setFormdata({ ...formdata, [event.target.name]: event.target.value });
  }

  const handleSubmit = () => {
    const url = "http://localhost:2222/cells/0";
    fetch(url, {
      method: "POST",
      //mode: "no-cors",
      body: JSON.stringify(formdata),
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
              setOpen(false);
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

  const handleClickPending = () => {
    setPendingPosition(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div>
      <div
        style={{
          display: pendingPosition || pendingSize ? "block" : "none",
          position: "absolute",
          top: 0,
          left: 0,
          width: "100%",
          height: "100%",
          cursor: "crosshair",
        }}
        onClick={(e) => {
          console.log(e);
          if (pendingPosition) {
            setPendingPosition(false);
            setPendingSize(true);
            setFormdata({ ...formdata, ['position']: [coords.x * scaleIndex, coords.y * scaleIndex, 0] });
          } else if (pendingSize) {
            setPendingPosition(false);
            setPendingSize(false);
            setFormdata({
              ...formdata,
              ['size']: [coords.x * scaleIndex - formdata.position[0], coords.y * scaleIndex - formdata.position[1]]
            });
            setOpen(true);
          }
        }}
      ></div>
      <Stack sx={{
        position: "absolute",
        top: "2em",
        left: "2em",
        display: pendingPosition || pendingSize ? "block" : "none",
      }} spacing={2}>
        <Alert variant="filled" severity="info">
          {pendingPosition
            ? "Select the element position on screen"
            : "Select the element size on screen"}
        </Alert>
      </Stack>
      <Fab
        style={{ position: "absolute", bottom: "2em", right: "2em" }}
        size="large"
        color={pendingPosition || pendingSize ? "default" : "primary"}
        aria-label="add"
        onClick={handleClickPending}
      >
        <AddIcon />
      </Fab>
      <Dialog fullWidth={true} open={open} onClose={handleClose}>
        <DialogContent>
          <TextField
            autoFocus
            multiline
            rows={3}
            maxRows={6}
            margin="dense"
            id="data"
            name="data"
            label="Data"
            fullWidth
            variant="standard"
            value={formdata.data}
            onChange={handleChange}
          />
          <TextField
            margin="dense"
            id="tags"
            name="tags"
            label="Tags"
            fullWidth
            variant="standard"
            value={formdata.tags}
            onChange={handleChange}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleSubmit}>Save</Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}
