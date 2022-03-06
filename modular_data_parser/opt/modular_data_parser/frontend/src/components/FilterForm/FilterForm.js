import * as React from 'react';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import { Select, InputLabel, MenuItem } from '@mui/material';
import FormControl from '@mui/material/FormControl';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import Button from '@mui/material/Button';



export default function FilterForm(props) {

  const theme = createTheme();

  const [age, setDataIn] = React.useState('');

  const handleChange = (event, props) => {
    //setDataIn(event.target.value);
    console.log(event.target.value)
    console.log(props.name)
    //var lang = event.target.value;
    //this.props.onSelectLanguage(lang);
  };

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <React.Fragment>
        <Typography variant="h6" gutterBottom>
          Upload your filter file if no filter to be used just skip...
        </Typography>
        <Grid container spacing={3}>
          <Grid item xs={6} sm={12}>
          <Button
            variant="contained"
            component="label"
          >
            Upload File
            <input
              type="file"
              hidden
            />
          </Button>
        </Grid>
        </Grid>
      </React.Fragment>
    </ThemeProvider>
  );
}