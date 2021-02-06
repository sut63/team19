import React, { useState, useEffect, FC } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
//import DeleteIcon from '@material-ui/icons/Delete';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntSubjectsOffered } from '../../api/models/EntSubjectsOffered'; // import interface SubjectsOffered
import { Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, Grid, InputLabel, MenuItem, Select } from '@material-ui/core';
import Swal from 'sweetalert2';
  const useStyles = makeStyles((theme: Theme) =>

  createStyles({

    table: {
      minWidth: 650,
    },
   
    root: {
      '& .MuiTextField-root': {
        marginLeft: theme.spacing(9),
        width: '30ch',
      },
    },


text0: {  
      marginLeft: theme.spacing(0),
      marginRight: theme.spacing(0),
      marginTop: theme.spacing(1),
      fontSize: 20,
    },

text4: {
  marginLeft: theme.spacing(3), 
  marginRight: theme.spacing(124), 
  fontSize: 30,
  color: "white",
},


text7: {
  marginLeft: theme.spacing(10),
  marginTop: theme.spacing(0),
  color: "white",
},


box1: {
  marginLeft: theme.spacing(0),
  marginTop: theme.spacing(0),
  marginBottom: theme.spacing(2),
  width: '40ch',  
},


  button: {
    marginLeft: theme.spacing(4),
  marginTop: theme.spacing(1),
  marginBottom: theme.spacing(2),
    //display: 'flex',
    //flexWrap: 'wrap',
    
  },

  paper: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(2),
  },
}),
);

const SearchSubjectsOffered: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  const [subjectsofferedshow, setSubjectsOfferedShow] = useState<EntSubjectsOffered[]>([]);
  const [subjectsofferedid, setSubjectsOfferedId] = useState<number>(0);

  

  useEffect(() => {
    getSubjectsOffered();
  }, []);

  const [subjectsoffered, setSubjectsOffered] = useState<EntSubjectsOffered[]>([]);
  const getSubjectsOffered = async () => {
    const res = await api.listSubjectsOffered({ limit: 10 });
    setSubjectsOffered(res); 
  };

  //const deleteSubjectsOffered = async (id: number) => {
  //  const res = await http.deleteSubjectsOffered({ id: id });
  //};
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
      },
    });
  var alerts : number;

  const getSubjectsOfferedShow = async () => {
    const res = await api.getSubjectsOffered({id: subjectsofferedid});
    setSubjectsOfferedShow(res);
    alerts = res.length ;
    if(alerts > 0 ){
      Toast.fire({
        icon: 'success',
        title: 'Search Success',
      })
    }else {
      Toast.fire({
        icon: 'error',
        title: 'Search Error',
      })
    }
   };

   const handleChange = (
    event: React.ChangeEvent<{name?: string ;value: any }>,
  ) => {
    const name = event.target.name as keyof typeof SearchSubjectsOffered;
    const { value } = event.target;
    setSubjectsOfferedId(value);
    console.log(subjectsofferedid);
  };

  function search(){

    getSubjectsOfferedShow();
  }
  return (
    <Page theme={pageTheme.home}>
    <Header
      title={`ระบบค้นหารายวิชาที่เปิดสอน`}
      subtitle=""
    >
    </Header>
    <Content>
    <Grid container spacing={1}>  
     <Grid item xs={2}>
       <FormControl className={classes.text0} > <text>ชื่อรายวิชาที่ต้องการ</text> </FormControl>   
      </Grid>
      <Grid item xs={3}>
       <FormControl fullWidth  variant="outlined" 
        className={classes.box1}>
        <InputLabel>เลือก</InputLabel>
        <Select
          value={ subjectsofferedid || ''}
          onChange={handleChange}
        >
        {subjectsoffered.map(item => {
            return (
              <MenuItem  value = {item.id}>
                {item.edges?.subject?.subjectName}
                </MenuItem>
            );
          }  
            )}
          
        </Select>
        </FormControl>
      </Grid> 
      <Grid>
     <Button size="large"  className={classes.button}  variant="contained" color="primary" onClick = {search}> ค้นหารายวิชาที่เปิดสอน </Button>
     </Grid>

  </Grid>
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">ID</TableCell>
            <TableCell align="center">Subject</TableCell>
            <TableCell align="center">Degree</TableCell>
            <TableCell align="center">Year</TableCell>
            <TableCell align="center">Term</TableCell>
            <TableCell align="center">Amount</TableCell>
            <TableCell align="center">Status</TableCell>
            <TableCell align="center">Remain</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {subjectsofferedshow === undefined
            ? null
            : subjectsofferedshow.map((item) => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.id}</TableCell>
              <TableCell align="center">{item.edges?.subject?.subjectName}</TableCell>
              <TableCell align="center">{item.edges?.degree?.degreeName}</TableCell>
              <TableCell align="center">{item.edges?.year?.yEAR}</TableCell>
              <TableCell align="center">{item.edges?.term?.tERM}</TableCell>
              <TableCell align="center">{item.aMOUNT}</TableCell>
              <TableCell align="center">{item.sTATUS}</TableCell>
              <TableCell align="center">{item.remain}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
    </Content>
   </Page>
  );
}
export default SearchSubjectsOffered;