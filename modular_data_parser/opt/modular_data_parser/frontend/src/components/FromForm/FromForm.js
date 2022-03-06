import * as React from 'react';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import { Select, InputLabel, MenuItem } from '@mui/material';
import FormControl from '@mui/material/FormControl';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';

export default function AddressForm() {

  const theme = createTheme();

  const [age, setDataIn] = React.useState('');

  const handleChange = (event) => {
    setDataIn(event.target.value);
    console.log(event.target.value)
    //var lang = this.dropdown.value;
    //this.props.onSelectLanguage(lang);
  };

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
                value={age}
                label="Select Data"
                onChange={handleChange}
              >
                <MenuItem value={"JSON"}>JSON</MenuItem>
                <MenuItem value={"CSV"}>CSV</MenuItem>
                <MenuItem value={"TSV"}>TSV</MenuItem>
              </Select>
            </FormControl>
          </Grid>
        </Grid>
      </React.Fragment>
    </ThemeProvider>

  );
}