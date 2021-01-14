import React, { FC,useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
} from '@material-ui/core';
import { makeStyles, } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { EntYear } from '../../api/models/EntYear'; // import interfa Year
import { EntTerm } from '../../api/models/EntTerm'; // import interfa Term
import { EntSubject } from '../../api/models/EntSubject'; // import interface Subjects
import { EntDegree } from '../../api/models/EntDegree'; // import interface Degree
import Swal from 'sweetalert2'; // alert

const useStyles = makeStyles(theme => ({
   
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

  

interface  EntSubjectsOffered  {
  Subject :    number;
  Year  :    number;
  Degree  :    number;
  Term :       number;
  AMOUNT :      string;
  STATUS :      string;
}

const EntSubjectsOffered: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [subjectsoffered, setSubjectsOffered] = React.useState<Partial<EntSubjectsOffered>>({});

  const [subject, setSubject] = React.useState<EntSubject[]>([]);
  const [degree, setDegree] = React.useState<EntDegree[]>([]);
  const [year, setYear] = React.useState<EntYear[]>([]);
  const [term, setTerm] = React.useState<EntTerm[]>([]);
    // alert setting
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
  const getSubject = async () => {
    const res = await http.listSubject({ limit: 10, offset: 0 });
    setSubject(res);
  };

  const getDegree = async () => {
    const res = await http.listDegree({ limit: 10, offset: 0 });
    setDegree(res);
  };

  const getYear = async () => {
    const res = await http.listYear({ limit: 10, offset: 0 });
    setYear(res);
  };
  const getTerm = async () => {
    const res = await http.listTerm({ limit: 10, offset: 0 });
    setTerm(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getSubject();
    getYear();
    getDegree();
    getTerm();
  }, []);

  // set data to object so
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof EntSubjectsOffered;
    const { value } = event.target;
    setSubjectsOffered({ ...subjectsoffered, [name]: value });
    console.log(subjectsoffered);
  };
  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/Subjectsoffereds';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(subjectsoffered),
    };

    console.log(subjectsoffered); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.state === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ระบบบันทึกรายวิชาที่เปิดสอน`}
       subtitle=""
     >
       <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/"
               variant="contained"color="primary"
             >
               ออกจากระบบ
       </Button>
     </Header>
     <Content>
       <ContentHeader title="บันทึกข้อมูล"></ContentHeader>
       <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>รายวิชา</div>
            </Grid>
            <Grid item xs={9}>
      <FormControl fullWidth  variant="outlined" 
        className={classes.formControl}>
        <InputLabel>เลือกรายวิชา</InputLabel>
        <Select
          name="Subject"
          type="number"
          value={subjectsoffered.Subject || ''}
          onChange={handleChange}
        >
          {subject.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.nAME}
                      </MenuItem>
                    );
                  })}
        </Select>
      </FormControl>
      </Grid>
            

            <Grid item xs={3}>
              <div className={classes.paper}>ปีการศึกษา</div>
            </Grid>
            <Grid item xs={9}>
      <FormControl fullWidth  variant="outlined" 
        className={classes.formControl}>
        <InputLabel>เลือกปีการศึกษา</InputLabel>
        <Select
          name="Year"
          type="number"
          value={subjectsoffered.Year || ''}
          onChange={handleChange}
        >
          {year.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.yEAR}
                      </MenuItem>
                    );
                  })}
        </Select>
        
      </FormControl>
      </Grid>

      <Grid item xs={3}>
              <div className={classes.paper}>ภาคการศึกษา</div>
            </Grid>
            <Grid item xs={9}>
      <FormControl fullWidth  variant="outlined" 
        className={classes.formControl}>
        <InputLabel>เลือกภาคการศึกษา</InputLabel>
        <Select
          name="Term"
          type="number"
          value={subjectsoffered.Term || ''}
          onChange={handleChange}
        >
          {term.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.tERM}
                      </MenuItem>
                    );
                  })}
        </Select>
        
      </FormControl>
      </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ระดับการศึกษา</div>
            </Grid>
            <Grid item xs={9}>
      <div>
      <FormControl fullWidth  variant="outlined"
        className={classes.formControl}>
        <InputLabel id="demo-simple-select-label">เลือกระดับการศึกษา</InputLabel>
        <Select
          name="Degree"
          type="number"
          value={subjectsoffered.Degree || ''}
          onChange={handleChange}
        >
          {degree.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.educationLevel}
                      </MenuItem>
                    );
                  })}
        </Select>
      </FormControl>
      </div>
      </Grid>
      <Grid item xs={3}>
              <div className={classes.paper}>จำนวนที่รับ</div>
            </Grid>
            <Grid item xs={9}>
      <form className={classes.root} noValidate autoComplete="off" >
        <TextField label="กรอกจำนวนที่รับ"name ="AMOUNT" type="string" 
        value={subjectsoffered.AMOUNT } onChange={handleChange} className={classes.textField}
                  />
      </form>
      </Grid>
      <Grid item xs={3}>
              <div className={classes.paper}>สถานะ</div>
            </Grid>
            <Grid item xs={9}>
            <form className={classes.root} noValidate autoComplete="off">
      <TextField  label="สถานะ" name ="STATUS" type="string" 
      value={subjectsoffered.STATUS } onChange={handleChange} className={classes.textField}/>
    </form>      
      </Grid>
          <Grid item xs={3}></Grid>
          <Grid item xs={9}>
             <Button
               onClick={save}
               startIcon={<SaveIcon />}
               variant="contained"
               color="primary"
             >
               บันทึกข้อมูล
             </Button>
           
           </Grid>
          </Grid>
        </Container>
        </Content>
   </Page>
 );
}
export default EntSubjectsOffered;