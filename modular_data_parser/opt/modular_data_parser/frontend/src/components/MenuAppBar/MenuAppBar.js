import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
//import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import Icons8 from './icons8-ship-64.png'

export default function MenuAppBar() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <IconButton>
            <img  src={Icons8} alt=""/>
          </IconButton>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            SISO Ship-In Ship-Out
          </Typography>
        </Toolbar>
      </AppBar>
    </Box>
  );
}