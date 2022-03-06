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

  const [filterData, setFilterData] = React.useState('');

  const handleChange = (event) => {
    setFilterData(event.target.files);
    const formData = {
      files: filterData
    }
    props.onSaveFilterFormData(formData)
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