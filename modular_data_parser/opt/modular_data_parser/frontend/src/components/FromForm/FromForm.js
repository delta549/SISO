import * as React from 'react';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import { Select, InputLabel, MenuItem } from '@mui/material';
import FormControl from '@mui/material/FormControl';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';

export default function AddressForm(props) {

  const theme = createTheme();

  let [dataIn, setDataIn] = React.useState('');

  const handleChangeSelection = (event) => {
    setDataIn(event.target.value);
  };

  const handleChangeFile = (event) => {
    const formData = {
      dataIn: dataIn,
      fromFiles: event.target.files[0]
    }
    props.onSaveFromFormData(formData);
  }

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <React.Fragment>
        <Typography variant="h6" gutterBottom>
          Choose the data type input e.g: JSON, TSV, CSV
        </Typography>
        <Grid container spacing={3}>
          <Grid item xs={12} sm={6}>
            <FormControl variant="filled" sx={{ m: 1, minWidth: 120 }}>
              <InputLabel id="demo-simple-select-label">Data Type</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={dataIn}
                label="Select Data"
                onChange={handleChangeSelection}
              >
                <MenuItem value={"JSON"}>JSON</MenuItem>
                <MenuItem value={"CSV"}>CSV</MenuItem>
                <MenuItem value={"TSV"}>TSV</MenuItem>
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6} sm={3}>
            Upload File
            <input
              type="file"
              onChange={handleChangeFile}
            />
        </Grid>
        </Grid>
      </React.Fragment>
    </ThemeProvider>

  );
}