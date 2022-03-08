import * as React from 'react';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';



export default function FilterForm(props) {

  const theme = createTheme();

  const handleChange = (event) => {
    props.onSaveFilterFormData(event.target.files)

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
            <input
              type="file"
              onChange={handleChange}
            />
        </Grid>
        </Grid>
      </React.Fragment>
    </ThemeProvider>
  );
}