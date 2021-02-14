import React, { useState, useEffect, FC } from 'react';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
//import DoneOutlineIcon from '@material-ui/icons/DoneOutline';
import SearchIcon from '@material-ui/icons/Search';
import Swal from 'sweetalert2';
import {
    Grid,
    TextField,
} from '@material-ui/core';
import { DefaultApi } from '../../api';
import { EntCourse } from '../../api/models/EntCourse'; // import interface EntCourse
const useStyles = makeStyles(theme => ({
  table: {
    minWidth: 650,
  },
  paper: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(2),
  },
  textField: {
    width: 300,
  },
  formControl: {
    width: 300,
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  }));

const FindCourse: FC<{}> = ()  => {
  const classes = useStyles();
  const api = new DefaultApi();
  const [searchstext, setSearchstext] = React.useState(String)
  const [coursesearch, setCoursesearch] = React.useState<EntCourse[]>([])

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'center',
    showConfirmButton: false,
    timer: 4000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
      },
    });

  //courses
  const [courses, setCourse] = React.useState<EntCourse[]>([])
  const getCourse = async () => {
   const res = await api.listCourse({limit: 100, offset: 0 });
   setCourse(res);
  };

  //coursesshow
  const [coursesshow, setCourseshow] = React.useState<EntCourse[]>([])
  const getCourseshow = async () => {
   const res = await api.listCourse({limit: 100, offset: 0 });
   setCourseshow(res);
  };
    
  // Lifecycle Hooks
  useEffect(() => {
    getCourse();
    getCourseshow();
  }, []);

  function clearshow(){
    setCourseshow([])
  }

  var alerts : number

  function Searchs() {
    courses.map((item: any) => {  
      if (item.courseName == searchstext ) {
    const getCoursesearch = async () => {
      const res = await api.getCourse({id: item.id})
      setCoursesearch(res);
      alerts = res.length
      if (alerts > 0) {
        Toast.fire({
          icon: 'success',
          title: 'พบหลักสูตรที่คุณค้นหา',
        })
      } else {
        Toast.fire({
          icon: 'error',
          title: 'ไม่พบหลักสูตรที่คุณค้นหา',
        })
      }
     };
    getCoursesearch();
    }
    else {
      setCoursesearch([]);
      Toast.fire({
        icon: 'error',
        title: 'ไม่พบหลักสูตรที่คุณค้นหา',
      })
    }
  });
  }

  const handleChange = (event: React.ChangeEvent<{ value: unknown }>,) =>{
    setSearchstext(event.target.value as string);
    console.log(searchstext);
  }

  return (
    <Content>
    <Grid container spacing={1}>
                    <Grid item xs={1}>
                  <div className={classes.paper}>Find Course </div>
                </Grid>
                <Grid item xs={3}>
                  <form className={classes.container}>
                    <TextField
                      label="Name Course"
                      name="name"
                      type="string"
                      value={searchstext || ''} 
                      className={classes.textField}
                      onChange={handleChange}
                    />
                  </form>
                </Grid>

                    <Grid item xs={1}>
                        <Button
                            variant="contained"
                            color="secondary"
                            size="large"
                            startIcon={<SearchIcon/>}
                            onClick={() =>{
                                  Searchs();
                                  clearshow()
                                }}
                        >    Search
                        </Button>
                    </Grid>
      </Grid>
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
          <TableCell align="center">id</TableCell>
           <TableCell align="center">year</TableCell>
           <TableCell align="center">name</TableCell>
           <TableCell align="center">Teacher_id</TableCell>
           <TableCell align="center">Degree</TableCell>
           <TableCell align="center">Department</TableCell>
           <TableCell align="center">Subject</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {coursesshow.map((item : EntCourse) => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.id}</TableCell>  
             <TableCell align="center">{item.courseYear}</TableCell>     
             <TableCell align="center">{item.courseName}</TableCell>     
             <TableCell align="center">{item.teacherId}</TableCell>     
             <TableCell align="center">{item.edges?.degreeID?.degreeName}</TableCell>     
             <TableCell align="center">{item.edges?.departmentID?.dEPARTMENT}</TableCell>     
             <TableCell align="center">{item.edges?.subjectID?.subjectName}</TableCell> 
             <TableCell align="center"></TableCell>   
            </TableRow>
          ))}
        </TableBody>
        <TableBody>
          {coursesearch.map((item : EntCourse) => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.id}</TableCell>  
             <TableCell align="center">{item.courseYear}</TableCell>     
             <TableCell align="center">{item.courseName}</TableCell>     
             <TableCell align="center">{item.teacherId}</TableCell>     
             <TableCell align="center">{item.edges?.degreeID?.degreeName}</TableCell>     
             <TableCell align="center">{item.edges?.departmentID?.dEPARTMENT}</TableCell>     
             <TableCell align="center">{item.edges?.subjectID?.subjectName}</TableCell> 
             <TableCell align="center"></TableCell>   
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
    </Content>
  );
}
export default FindCourse;
