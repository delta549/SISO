import * as React from 'react';
import CssBaseline from '@mui/material/CssBaseline';
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import Paper from '@mui/material/Paper';
import Stepper from '@mui/material/Stepper';
import Step from '@mui/material/Step';
import StepLabel from '@mui/material/StepLabel';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import MenuAppBar from '../MenuAppBar/MenuAppBar';

import FromForm from '../FromForm/FromForm';
import FilterForm from '../FilterForm/FilterForm'
import ToForm from '../ToForm/ToForm';



const steps = ['Convert From', 'Choose Filter', 'Convert Too'];

const newObj = {}

function allData (allData) {
  console.log("here")
  console.log(newObj)
  
}

function saveFromFormData(FromFormData) {
  const newFromFormData = {
    ...FromFormData
  };
  newObj["FromFormData"] = newFromFormData;
  //console.log(newFromFormData)
};

function saveFilterFormData(FilterFormData) {
  const newFilterFormData = {
    ...FilterFormData
  };
  newObj["newFilterData"] = newFilterFormData;
};

function savetoFormData(toFormData) {
  const newToFormData = {
    ...toFormData
  };
  newObj["newToFormData"] = newToFormData;
  console.log("firing to form data")
  console.log(toFormData)
};

function getStepContent(step) {
  switch (step) {
    case 0:
      return <FromForm onSaveFromFormData={saveFromFormData}/>;
    case 1:
      return <FilterForm onSaveFilterFormData={saveFilterFormData}/>;
    case 2:
      return <ToForm onSaveToFormData={savetoFormData}/>;
    default:
      throw new Error('Unknown step');
  }
}

const theme = createTheme();

export default function Checkout() {
  const [activeStep, setActiveStep] = React.useState(0); 

  const handleNext = () => {
    setActiveStep(activeStep + 1);
  };

  const handleBack = () => {
    setActiveStep(activeStep - 1);
  };

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <MenuAppBar/>
      <Container component="main" maxWidth="sm" sx={{ mb: 4 }}>
        <Paper variant="outlined" sx={{ my: { xs: 3, md: 6 }, p: { xs: 2, md: 3 } }}>
          <Typography component="h1" variant="h4" align="center">
            Convert!
          </Typography>
          <Stepper activeStep={activeStep} sx={{ pt: 3, pb: 5 }}>
            {steps.map((label) => (
              <Step key={label}>
                <StepLabel>{label}</StepLabel>
              </Step>
            ))}
          </Stepper>
          <React.Fragment>
            {activeStep === steps.length ? (
              <React.Fragment>
                {allData(newObj)}
                <Typography variant="h5" gutterBottom>
                  Converting
                </Typography>
              </React.Fragment>
            ) : (
              <React.Fragment>
                {getStepContent(activeStep)}
                <p></p>
                <Box sx={{ display: 'flex', justifyContent: 'flex-end' }}>
                  {activeStep !== 0 && (
                    <Button onClick={handleBack} sx={{ mt: 3, ml: 1 }}>
                      Back
                    </Button>
                  )}

                  <Button
                    variant="contained"
                    onClick={handleNext}
                    sx={{ mt: 3, ml: 1 }}
                  >
                    {activeStep === steps.length - 1 ? 'Convert!': 'Next'}
                  </Button>
                </Box>
              </React.Fragment>
            )}
          </React.Fragment>
        </Paper>
      </Container>
    </ThemeProvider>
  );
}