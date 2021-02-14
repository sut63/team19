import React, { useState, useEffect, FC } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
//import DeleteIcon from '@material-ui/icons/Delete';
import SearchIcon from '@material-ui/icons/Search';
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
import {  Grid, TextField } from '@material-ui/core';
import Swal from 'sweetalert2';
  const useStyles = makeStyles((theme: Theme) =>

  createStyles({

    table: {
      minWidth: 650,
    },

  paper: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(2),
    fontSize: 20,
  },

  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
   
  textField: {
    width: 2000,
  },

}),
);

const SearchSubjectsOffered: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();
  const [searchstext, setSearchstext] = React.useState(String)
  const [subjectsoffered, setSubjectsOffered] = useState<EntSubjectsOffered[]>([]);
  const [subjectsofferedsearch, setSubjectsOfferedsearch] = useState<EntSubjectsOffered[]>([]);
  

  useEffect(() => {
    getSubjectsOffered();
  }, []);

  
  const getSubjectsOffered = async () => {
    const res = await api.listSubjectsOffered({ limit: 10 });
    setSubjectsOffered(res); 
  };

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

   const handleChange = (event: React.ChangeEvent<{ value: unknown }>,) =>{
    setSearchstext(event.target.value as string);
    console.log(searchstext);
  }

  function search(){
    subjectsoffered.map((item: any) => {  
      if (item.edges?.subject?.subjectName == searchstext ) {
    const getSubjectsOfferedsearch = async () => {
      const res = await api.getSubjectsOffered({id: item.id})
      setSubjectsOfferedsearch(res);
      alerts = res.length
      if (alerts > 0) {
        Toast.fire({
          icon: 'success',
          title: 'Search Success',
        })
      } else {
        Toast.fire({
          icon: 'error',
          title: 'Search Error',
        })
      }
     };
     getSubjectsOfferedsearch();
    }
    else {
      setSubjectsOfferedsearch([]);
      Toast.fire({
        icon: 'error',
        title: 'Search Error',
      })
    }
  });
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
                  <div className={classes.paper}> ค้นหารายวิชาที่เปิดสอน </div>
                </Grid>
                <Grid item xs={3}>
                  <form className={classes.container}>
                    <TextField
                      label="รายวิชาที่เปิดสอน"
                      name="name"
                      type="string"
                      value={searchstext || ''} 
                      className={classes.textField}
                      onChange={handleChange}
                    />
                  </form>
                </Grid>

                    <Grid item xs={2}>
                        <Button
                            variant="contained"
                            color="secondary"
                            size="large"
                            startIcon={<SearchIcon/>}
                            onClick={() =>{
                                  search();
                                }}
                        >    Search
                        </Button>
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
          {subjectsofferedsearch === undefined
            ? null
            : subjectsofferedsearch.map((item) => (
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